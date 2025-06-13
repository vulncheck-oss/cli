package ui

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/lipgloss"
)

var format = "%s %s\n"
var statFormat = "%s %s: %s\n"

var Pantone = lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))
var White = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffffff"))
var Gray = lipgloss.NewStyle().Foreground(lipgloss.Color("#cccccc"))
var Emerald = lipgloss.NewStyle().Foreground(lipgloss.Color("#34d399"))
var Red = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff0000"))

func ClearScreen() {
	cmd := exec.Command("clear") // for Linux and macOS
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	}
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func Success(str string) {
	fmt.Printf(
		format,
		Emerald.Render("✓"),
		White.Render(str),
	)
}

func Info(str string) {
	fmt.Printf(
		format,
		Pantone.Render("i"),
		White.Render(str),
	)
}

func Stat(label string, value string) {
	fmt.Printf(
		statFormat,
		Pantone.Render("i"),
		Gray.Render(label),
		White.Render(value),
	)
}

func Danger(str string) error {
	return fmt.Errorf(
		format,
		Red.Render("✗"),
		White.Render(str),
	)
}

func Json(data interface{}) {
	marshaled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(marshaled))
}

func JsonFile(data interface{}, filename string) error {
	marshaled, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, marshaled, 0644)
	if err != nil {
		return err
	}

	return nil
}

type FlagError struct {
	// Note: not struct{error}: only *FlagError should satisfy error.
	err error
}

func (fe *FlagError) Error() string {
	return fe.err.Error()
}

func (fe *FlagError) Unwrap() error {
	return fe.err
}

func Error(format string, args ...interface{}) error {
	// If no args are provided, treat format as a plain string
	// This avoids the non-constant format string linter error
	if len(args) == 0 {
		return FlagErrorWrap(fmt.Errorf("%s", format))
	}
	return FlagErrorWrap(fmt.Errorf(format, args...))
}

// FlagErrorWrap FlagError returns a new FlagError that wraps the specified error.
func FlagErrorWrap(err error) error { return &FlagError{err} }
