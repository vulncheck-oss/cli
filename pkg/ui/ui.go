package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var format = "%s %s\n"

var Pantone = lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))
var White = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
var Emerald = lipgloss.NewStyle().Foreground(lipgloss.Color("#34d399"))
var Red = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))

func Success(str string) error {
	fmt.Printf(
		format,
		Emerald.Render("✓"),
		White.Render(str),
	)
	return nil
}

func Info(str string) error {
	fmt.Printf(
		format,
		Pantone.Render("i"),
		White.Render(str),
	)
	return nil
}

func Danger(str string, a ...any) error {
	return fmt.Errorf(
		format,
		Red.Render("✗"),
		White.Render(fmt.Sprintf(str, a)),
	)
}
