package status

import (
	"errors"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/login"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
)

// authStatus is the JSON payload emitted by `auth status --json`.
// Agents can rely on .authenticated being a bool — never null — so a
// missing token cleanly maps to {"authenticated": false, ...}.
type authStatus struct {
	Authenticated bool   `json:"authenticated"`
	TokenSource   string `json:"token_source,omitempty"` // "env" | "config" | ""
	User          string `json:"user,omitempty"`
	Email         string `json:"email,omitempty"`
	Reason        string `json:"reason,omitempty"` // populated when authenticated=false
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: i18n.C.AuthStatusShort,
		Long:  i18n.C.AuthStatusLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			config.Init()

			if !config.HasToken() {
				if r.IsJSON() {
					// No token = not an error from the agent's POV; emit the
					// status payload and exit 0 so callers can dispatch on
					// .authenticated rather than parsing error envelopes.
					return r.JSON(authStatus{
						Authenticated: false,
						Reason:        i18n.C.ErrorNoToken,
					})
				}
				return errs.AuthRequired(i18n.C.ErrorNoToken)
			}

			source := "config"
			if config.TokenFromEnv() {
				source = "env"
			}

			// Verify against the API rather than trusting "token is present".
			resp, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetMe()
			if err != nil {
				if r.IsJSON() {
					reason := err.Error()
					if errors.Is(err, sdk.ErrorUnauthorized) {
						reason = "token rejected by server"
					}
					return r.JSON(authStatus{
						Authenticated: false,
						TokenSource:   source,
						Reason:        reason,
					})
				}
				return err
			}

			if r.IsJSON() {
				return r.JSON(authStatus{
					Authenticated: true,
					TokenSource:   source,
					User:          resp.Data.Name,
					Email:         resp.Data.Email,
				})
			}

			return login.SaveToken(config.Token())
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}
