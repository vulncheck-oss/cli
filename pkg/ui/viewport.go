package ui

import (
	"encoding/json"
	"fmt"
	"github.com/alecthomas/chroma/v2/quick"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

const useHighPerformanceRenderer = false

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
	index       string
	content     string
	ready       bool
	searching   bool
	searchQuery string
	searchIndex int
	viewport    viewport.Model
}

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
			if m.searching || m.searchQuery != "" {
				m.searching = false
				m.searchQuery = ""
				m.searchIndex = 0

			} else {
				return m, tea.Quit
			}
		case "/":
			m.searching = true
			m.searchQuery = ""
			m.searchIndex = 0

		case "backspace":
			if m.searching && len(m.searchQuery) > 0 {
				m.searchQuery = m.searchQuery[:len(m.searchQuery)-1]
			}

		case "enter":
			if m.searching {
				m.searching = false
				if m.searchQuery != "" {
					m.searchIndex = strings.Index(m.content, m.searchQuery)
					if m.searchIndex != -1 {
						// Scroll to the search result
						lineNumber := strings.Count(m.content[:m.searchIndex], "\n")
						m.viewport.GotoTop()
						m.viewport.LineDown(lineNumber)
					}
				}
			}

		case "n":
			if m.searching {
				m.searchQuery += "n"
			} else if m.searchQuery != "" {
				nextIndex := strings.Index(m.content[m.searchIndex+1:], m.searchQuery)
				if nextIndex != -1 {
					m.searchIndex += nextIndex + 1
					// Scroll to the next search result
					lineNumber := strings.Count(m.content[:m.searchIndex], "\n")
					m.viewport.GotoTop()
					m.viewport.LineDown(lineNumber)
				} else {
					// If not found, wrap around to the beginning
					m.searchIndex = strings.Index(m.content, m.searchQuery)
					lineNumber := strings.Count(m.content[:m.searchIndex], "\n")
					m.viewport.GotoTop()
					m.viewport.LineDown(lineNumber)
				}
			}
		default:
			if m.searching {
				m.searchQuery += msg.String()
			}
		}

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

func (m model) View() string {
	if !m.ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s", m.headerView(), m.viewport.View(), m.footerView())
}

/*
func (m model) headerView() string {
	title := titleStyle.Render("Browsing index: " + m.index)
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}
*/

func (m model) headerView() string {
	var title string
	if m.searching {
		title = titleStyle.Render(fmt.Sprintf("Search: %s", m.searchQuery))
	} else if m.searchQuery != "" {
		occurrences := strings.Count(m.content, m.searchQuery)
		currentOccurrence := 0
		if m.searchIndex != -1 {
			currentOccurrence = strings.Count(m.content[:m.viewport.YOffset+m.searchIndex], m.searchQuery)
		}
		title = titleStyle.Render(fmt.Sprintf("Search: %s (%d/%d) - \"n\" = next, \"q\" = clear ", m.searchQuery, currentOccurrence, occurrences))
	} else {
		title = titleStyle.Render("Browsing index: " + m.index)
	}
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
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
