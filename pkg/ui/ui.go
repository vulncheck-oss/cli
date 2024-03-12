package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func Success(strings ...string) {
	greenCheck := lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render("✓")
	success := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(strings...)
	fmt.Println(greenCheck, success)
}
