package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
)

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
	copy     string
}

func (m *model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m *model) Quit() tea.Msg {
	m.quitting = true
	return tea.Quit
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m *model) View() string {
	return fmt.Sprintf("[ %s] %s", m.spinner.View(), m.copy)
}

func Spinner(copy string) *tea.Program {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))
	program := tea.NewProgram(&model{spinner: s, copy: copy})
	go func() {
		if _, err := program.Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()
	return program
}
