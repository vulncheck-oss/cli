package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vulncheck-oss/sdk"
	"strings"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#6667ab"))

type tableModel struct {
	table  table.Model
	action func(index string) error
}

func (m tableModel) Init() tea.Cmd { return nil }

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			m.action(m.table.SelectedRow()[0])
			return m, tea.Quit
			/*
				return m, tea.Batch(
					m.action(m.table.SelectedRow()[0]),
				)
			*/
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func IndicesRows(indices []sdk.IndicesMeta, search string) []table.Row {
	var rows []table.Row
	for _, index := range indices {
		if search != "" && !strings.Contains(index.Name, search) && !strings.Contains(index.Description, search) {
			continue
		}
		rows = append(rows, table.Row{
			index.Name,
			index.Description,
			index.Href,
		})
	}
	return rows
}

func Indices(indices []sdk.IndicesMeta, search string, action func(index string) error) error {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "Description", Width: 40},
		{Title: "URL", Width: 20},
	}

	rows := IndicesRows(indices, search)

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6667ab")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("#000000")).
		Background(lipgloss.Color("#34d399")).
		Bold(false)
	t.SetStyles(s)

	m := tableModel{t, action}
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running program: %v", err)
	}

	return nil
}
