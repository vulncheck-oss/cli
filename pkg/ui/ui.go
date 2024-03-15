package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var format = "[ %s ] %s\n"

func Success(str string) {
	fmt.Printf(
		format,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render("✓"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(str),
	)
}

func Danger(sr string) error {
	return fmt.Errorf(
		format,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("✗"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(sr),
	)
}
