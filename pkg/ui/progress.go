package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strings"
)

type progressMessage int

var progressModelInstance *progressModel
var progressChannel chan int = make(chan int)

type progressModel struct {
	progress progress.Model
	total    int
	current  int
	percent  float64
}

// Configure a new progressModel
func newModel(total int) *progressModel {
	// Create progress bar with color gradient
	p := progress.New(progress.WithScaledGradient("#6667AB", "#34D399"))

	return &progressModel{
		progress: p,
		total:    total,
		current:  0,
	}
}

// Start a new progress bar
func NewProgress(total int) {
	progressModelInstance = newModel(total)
	p := tea.NewProgram(progressModelInstance)

	go func() {
		for update := range progressChannel {
			p.Send(progressMessage(update))

			// Close progress when 100% is reached
			if update >= total {
				close(progressChannel)
				break
			}
		}
	}()

	go func() {
		if err := p.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Could not start progress: %v", err)
			os.Exit(1)
		}
	}()
}

func UpdateProgress(value int) {
	progressChannel <- value
}

func (m progressModel) Init() tea.Cmd {
	return nil
}

func (m progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case progressMessage:
		m.current = int(msg)
		if m.current > m.total {
			m.current = m.total // Prevent going over
		}
		m.percent = float64(m.current) / float64(m.total)
		if m.current >= m.total {
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m progressModel) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.progress.ViewAs(m.percent) + "\n\n" +
		pad + helpStyle("Press any key to quit")
}
