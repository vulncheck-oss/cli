package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var format = "[ %s ] %s\n"

func Success(str string) error {
	fmt.Printf(
		format,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#34d399")).Render("✓"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(str),
	)
	return nil
}

func Info(str string) error {
	fmt.Printf(
		format,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab")).Render("i"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(str),
	)
	return nil
}

func Danger(sr string) error {
	return fmt.Errorf(
		format,
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000")).Render("✗"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(sr),
	)
}
