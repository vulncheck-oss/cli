package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type tokenModel struct {
	TextInput textinput.Model
	err       error
}

type (
	errMsg error
)

func initialModel() tokenModel {
	ti := textinput.New()
	ti.Placeholder = "vulncheck_***********"
	ti.Focus()
	ti.CharLimit = 74
	ti.Width = 74
	ti.EchoMode = textinput.EchoPassword

	return tokenModel{
		TextInput: ti,
		err:       nil,
	}
}

func (m tokenModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m tokenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.TextInput, cmd = m.TextInput.Update(msg)
	return m, cmd
}

func (m tokenModel) View() string {
	return fmt.Sprintf(
		"Paste your authentication token\n\n%s\n\n%s",
		m.TextInput.View(),
		"(esc to equit)",
	) + "\n"

}

func TokenPrompt() (string, error) {
	p := tea.NewProgram(initialModel())
	result, err := p.Run()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return result.(tokenModel).TextInput.Value(), nil
}
