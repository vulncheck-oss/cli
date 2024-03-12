package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	spinner  spinner.Model
	quitting bool
	err      error
	copy     string
}

func (m model) Init() tea.Cmd {
	return m.spinner.Tick
}

func (m model) Quit() {
	m.quitting = true
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			m.Quit()
		}
	}

	var cmd tea.Cmd
	m.spinner, cmd = m.spinner.Update(msg)
	return m, cmd
}

func (m model) View() string {
	str := fmt.Sprintf("%s %s \n", m.spinner.View(), m.copy)
	if m.quitting {
		return str + "\r"
	}
	return str
}

func Spinner(copy string) *tea.Program {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	p := tea.NewProgram(model{spinner: s, copy: copy})

	go func() {
		p.Run()
	}()

	return p
}
