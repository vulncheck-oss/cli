package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func Success(strings ...string) string {
	greenCheck := lipgloss.NewStyle().Foreground(lipgloss.Color("#00ff00")).Render("✓")
	success := lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff")).Render(strings...)
	return fmt.Sprintf("[ %s ] %s", greenCheck, success)
}
