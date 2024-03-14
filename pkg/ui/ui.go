package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func Success(str string) {
	fmt.Printf(
		"[ %s ] %s\n",
		lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render("✓"),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(str),
	)
}
