package login

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
)

func chooseAuthMethod() (string, error) {

	var choice string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an authentication method").
				Options(
					huh.NewOption("Login with a web browser", "web"),
					huh.NewOption("Paste an authentication token", "token"),
				).Value(&choice),
		),
	)

	err := form.Run()
	if err != nil {
		return "", ui.Error("Failed to select authentication method: %v", err)
	}

	return choice, nil
}

func existingToken() error {
	logoutChoice := true
	confirm := huh.NewForm(huh.NewGroup(huh.NewConfirm().
		Title("You currently have a token saved. Do you want to invalidate it first?").
		Affirmative("Yes").
		Negative("No").
		Value(&logoutChoice))).WithTheme(huh.ThemeDracula())
	confirm.Run()

	if logoutChoice {
		if _, err := session.InvalidateToken(config.Token()); err != nil {
			if err := config.RemoveToken(); err != nil {
				return ui.Error("Failed to remove token from config")
			}
			ui.Info("Token was not valid, removing from config")
		} else {
			if err := config.RemoveToken(); err != nil {
				return ui.Error("Failed to remove token from config")
			}
			ui.Success("Token invalidated successfully")
		}
	} else {
		return nil
	}

	return nil
}

func cmdToken(cmd *cobra.Command, args []string) error {

	var token string

	input := huh.
		NewInput().
		Title("Enter your authentication token").
		Password(true).
		Placeholder("vulncheck_******************").
		Value(&token)

	if err := input.Run(); err != nil {
		return ui.Error("Token verification failed: %v", err)
	}

	if !config.ValidToken(token) {
		return ui.Error("Invalid token specified")
	}

	return SaveToken(token)
}

func SaveToken(token string) error {

	var res *sdk.UserResponse
	var err error

	_ = spinner.New().
		Style(ui.Pantone).
		Title(" Verifying token...").Action(func() {
		res, err = session.CheckToken(token)
	}).Run()

	if err != nil {
		return ui.Error("Token verification failed: %v", err)
	}
	if err := config.SaveToken(token); err != nil {
		return ui.Error("Failed to save token: %v", err)
	}
	ui.Success(fmt.Sprintf("Authenticated as %s (%s)", res.Data.Name, res.Data.Email))
	return nil
}