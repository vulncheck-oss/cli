package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

// A simple message to update progress
type progressMessage int

// Global instance and a channel to communicate progress updates
var (
	currentModel *progressModel
	updateCh     chan int = make(chan int)
)

type progressModel struct {
	progress progress.Model
	total    int
	current  int
}

func initModel(total int) *progressModel {
	return &progressModel{
		progress: progress.New(progress.WithScaledGradient("#6667AB", "#34D399")),
		total:    total,
		current:  0,
	}
}

func NewProgress(total int) {
	currentModel = initModel(total)

	p := tea.NewProgram(currentModel)

	// Handling updates through a channel within the Bubble Tea's update loop
	go func() {
		for update := range updateCh {
			p.Send(progressMessage(update))
		}
	}()

	go func() {
		if err := p.Start(); err != nil {
			fmt.Fprintf(os.Stderr, "Error running progress program: %v", err)
			os.Exit(1)
		}
	}()
}

func UpdateProgress(value int) {
	updateCh <- value
}

func (m *progressModel) Init() tea.Cmd {
	return nil
}

func (m *progressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case progressMessage:
		m.current = int(msg)
		if m.current > m.total {
			close(updateCh) // Close channel to signal completion
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *progressModel) View() string {
	percentComplete := float64(m.current) / float64(m.total)
	return "\n" + m.progress.ViewAs(percentComplete) + "\n"
}
