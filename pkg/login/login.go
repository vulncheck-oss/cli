package login

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

// GuardEnvToken refuses a login when a token environment variable is set.
// Every login path writes to the config file, but Resolve() prefers the
// environment — so the write would be dead on arrival. Previously these paths
// verified the pasted token, skipped the save, and still printed
// "Authenticated as ...", which is why a user with a stale VC_TOKEN could log
// in repeatedly and never change anything.
//
// Callers must invoke this *before* prompting, so the user isn't asked to paste
// a token or complete a browser flow that is going to be discarded.
func GuardEnvToken() error {
	res := config.Resolve()
	if !res.FromEnv() {
		return nil
	}

	e := errs.Validation(
		"%s is set, so the CLI already authenticates with it and `auth login` would have no effect",
		res.EnvVar)
	if res.Shadowed {
		return e.WithHint(
			"a different token is already saved in your config file. To use it, %s, or keep using the environment token as-is.",
			res.UnsetHint())
	}
	// Exported on purpose by anyone running an SDK or the MCP server too, so
	// lead with the fact that nothing is wrong.
	return e.WithHint(
		"nothing to do if that is the token you want. To store a different credential in the config file instead, %s first.",
		res.UnsetHint())
}

func ChooseAuthMethod() (string, error) {

	var choice string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select an authentication method").
				Options(
					huh.NewOption("Login with a web browser (Does not work with Safari)", "web"),
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
		// Revoke the token saved on disk specifically. Callers reach this only
		// after GuardEnvToken, so Resolve() would return the same value today —
		// but naming ConfigToken keeps it correct if that ordering ever changes,
		// rather than silently revoking the caller's environment credential.
		if _, err := session.InvalidateToken(config.Resolve().ConfigToken); err != nil {
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

	if config.IsCI() {
		res, err = session.CheckToken(token)
	} else {
		_ = spinner.New().
			Style(ui.Pantone).
			Title(" Verifying token...").Action(func() {
			res, err = session.CheckToken(token)
		}).Run()
	}

	if err != nil {
		return ui.Error("Token verification failed: %v", err)
	}

	// An env token still wins after this write, so don't imply the save
	// changed which credential is in use. Login paths are already blocked by
	// GuardEnvToken; this branch is reached by `auth status`, which re-verifies
	// the active token without wanting to persist an environment value to disk.
	if active := config.Resolve(); active.FromEnv() {
		ui.Info(fmt.Sprintf("Verified as %s (%s) using the token from %s; config file not modified",
			res.Data.Name, res.Data.Email, active.EnvVar))
		return nil
	}

	if err := config.SaveToken(token); err != nil {
		return ui.Error("Failed to save token: %v", err)
	}
	ui.Success(fmt.Sprintf("Authenticated as %s (%s)", res.Data.Name, res.Data.Email))
	return nil
}
