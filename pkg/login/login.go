package login

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
)

func ChooseAuthMethod() (string, error) {

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

func ExistingToken() error {
	logoutChoice := true
	confirm := huh.NewForm(huh.NewGroup(huh.NewConfirm().
		Title("You currently have a token saved. Do you want to invalidate it first?").
		Affirmative("Yes").
		Negative("No").
		Value(&logoutChoice))).WithTheme(huh.ThemeCatppuccin())
	if err := confirm.Run(); err != nil {
		return err
	}

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

	if !config.TokenFromEnv() {
		if err := config.SaveToken(token); err != nil {
			return ui.Error("Failed to save token: %v", err)
		}
	}
	ui.Success(fmt.Sprintf("Authenticated as %s (%s)", res.Data.Name, res.Data.Email))
	return nil
}
