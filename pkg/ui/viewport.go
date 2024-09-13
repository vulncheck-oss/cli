package ui

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vulncheck-oss/sdk"
	"strings"
)

const useHighPerformanceRenderer = false

var subtle = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}

var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.BorderStyle(b)
	}()
)

type model struct {
	index      string
	content    string
	ready      bool
	viewport   viewport.Model
	paginated  bool
	page       int
	totalPages int
	showHelp   bool
	loadPage   func(index string, page int) (*sdk.IndexResponse, error)
}

type newPageMsg struct {
	content    string
	page       int
	totalPages int
}

type errMsg struct{ error }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			if m.showHelp {
				m.showHelp = false
				return m, nil
			}
			return m, tea.Quit
		case "?":
			m.showHelp = !m.showHelp
			return m, nil
		}

		if m.paginated {
			switch msg.String() {
			case "left", "[":
				if m.page > 1 {
					m.page--
					return m, loadPage(m)
				}
			case "right", "]":
				if m.page < m.totalPages {
					m.page++
					return m, loadPage(m)
				}
			}
		}

	case newPageMsg:
		m.content = msg.content
		m.page = msg.page
		m.totalPages = msg.totalPages
		m.viewport.SetContent(m.content)
		return m, nil

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight
			m.viewport.HighPerformanceRendering = useHighPerformanceRenderer
			m.viewport.SetContent(m.content)
			m.ready = true

			// This is only necessary for high performance rendering, which in
			// most cases you won't need.
			//
			// Render the viewport one line below the header.
			m.viewport.YPosition = headerHeight + 1
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

		if useHighPerformanceRenderer {
			// Render (or re-render) the whole viewport. Necessary both to
			// initialize the viewport and when the window is resized.
			//
			// This is needed for high-performance rendering only.
			cmds = append(cmds, viewport.Sync(m.viewport))
		}
	}

	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func loadPage(m model) tea.Cmd {
	return func() tea.Msg {
		response, err := m.loadPage(m.index, m.page)
		if err != nil {
			return errMsg{err}
		}

		marshaled, err := json.MarshalIndent(response.GetData(), "", "  ")
		if err != nil {
			return errMsg{err}
		}

		var buf strings.Builder
		err = quick.Highlight(&buf, string(marshaled), "json", "terminal256", "nord")
		if err != nil {
			return errMsg{err}
		}

		return newPageMsg{
			content:    buf.String(),
			page:       response.Meta.Page,
			totalPages: response.Meta.TotalPages,
		}
	}
}

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}

	if m.showHelp {
		return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.helpView(), m.footerView())
	}

	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

func (m model) headerView() string {
	title := titleStyle.Render("Browsing index: " + m.index)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	var info string
	if m.paginated {
		info = infoStyle.Render(fmt.Sprintf("%3.f%% | Page %d of %d | ? for help", m.viewport.ScrollPercent()*100, m.page, m.totalPages))
	} else {
		info = infoStyle.Render(fmt.Sprintf("%3.f%% | ? for help", m.viewport.ScrollPercent()*100))
	}
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func (m model) helpView() string {
	helpContent := `
Hot Keys:

q, esc         : Close help / Quit
?              : Toggle help
left, [        : Previous page
right, ]       : Next page
`

	// Create a style for the help content
	helpStyle := lipgloss.NewStyle().
		Width(42). // Increased width to accommodate border
		Align(lipgloss.Left).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2)

	// Render the help content
	renderedHelp := helpStyle.Render(helpContent)

	// Center the rendered help in the viewport
	return lipgloss.Place(
		m.viewport.Width,
		m.viewport.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderedHelp,
	)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Viewport(index string, data interface{}) {

	marshaled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	var buf strings.Builder

	err = quick.Highlight(&buf, string(marshaled), "json", "terminal256", "nord")

	if err != nil {
		panic(err)
	}

	p := tea.NewProgram(
		model{index: index, content: buf.String()},
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
	}
}

func ViewportPaginated(index string, initialData interface{}, initialMeta sdk.IndexMeta, loadPageFunc func(index string, page int) (*sdk.IndexResponse, error)) {
	marshaled, err := json.MarshalIndent(initialData, "", "  ")
	if err != nil {
		panic(err)
	}

	var buf strings.Builder
	err = quick.Highlight(&buf, string(marshaled), "json", "terminal256", "nord")
	if err != nil {
		panic(err)
	}

	m := model{
		index:      index,
		content:    buf.String(),
		paginated:  true,
		page:       initialMeta.Page,
		totalPages: initialMeta.TotalPages,
		loadPage:   loadPageFunc,
	}

	p := tea.NewProgram(
		m,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
	}
}
