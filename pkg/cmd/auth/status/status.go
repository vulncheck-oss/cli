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

func newAuthStatus() authStatus {
	return authStatus{SchemaVersion: output.SchemaVersion}
}

// authStatus is the JSON payload emitted by `auth status --json`.
// Agents can rely on .authenticated being a bool — never null — so a
// missing token cleanly maps to {"authenticated": false, ...}.
type authStatus struct {
	SchemaVersion int    `json:"schema_version"`
	Authenticated bool   `json:"authenticated"`
	TokenSource   string `json:"token_source,omitempty"` // "env" | "config" | ""
	// TokenEnvVar names the variable that supplied the token when TokenSource
	// is "env". Additive: callers that only know token_source keep working.
	TokenEnvVar string `json:"token_env_var,omitempty"` // "VC_TOKEN" | "VULNCHECK_API_TOKEN"
	// TokenShadowed reports that a token environment variable is overriding a
	// different token saved in the config file. Agents can key off this to
	// explain why a freshly saved credential appears to have no effect.
	TokenShadowed bool   `json:"token_shadowed,omitempty"`
	User          string `json:"user,omitempty"`
	Email         string `json:"email,omitempty"`
	Reason        string `json:"reason,omitempty"` // populated when authenticated=false
}

// tokenSourceLabel describes the active token's origin in human terms,
// naming the concrete location so the user knows where to go and change it.
func tokenSourceLabel(res config.Resolution) string {
	if res.FromEnv() {
		return res.EnvVar + " environment variable"
	}
	if dir, err := config.Dir(); err == nil {
		return dir + "/vulncheck.yaml"
	}
	return "config file"
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: i18n.C.AuthStatusShort,
		Long:  i18n.C.AuthStatusLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			config.Init()

			res := config.Resolve()

			if res.Token == "" {
				if r.IsJSON() {
					// No token = not an error from the agent's POV; emit the
					// status payload and exit 0 so callers can dispatch on
					// .authenticated rather than parsing error envelopes.
					s := newAuthStatus()
					s.Authenticated = false
					s.Reason = i18n.C.ErrorNoToken
					return r.JSON(s)
				}
				return errs.AuthRequired(i18n.C.ErrorNoToken)
			}

			// Verify against the API rather than trusting "token is present".
			resp, err := session.ConnectWithContext(cmd.Context(), res.Token).GetMe()
			if err != nil {
				if r.IsJSON() {
					reason := err.Error()
					if errors.Is(err, sdk.ErrorUnauthorized) {
						reason = "token rejected by server"
					}
					s := newAuthStatus()
					s.Authenticated = false
					s.TokenSource = string(res.Source)
					s.TokenEnvVar = res.EnvVar
					s.TokenShadowed = res.Shadowed
					s.Reason = reason
					return r.JSON(s)
				}
				return err
			}

			if r.IsJSON() {
				s := newAuthStatus()
				s.Authenticated = true
				s.TokenSource = string(res.Source)
				s.TokenEnvVar = res.EnvVar
				s.TokenShadowed = res.Shadowed
				s.User = resp.Data.Name
				s.Email = resp.Data.Email
				return r.JSON(s)
			}

			// The most useful fact when an unexpected account is in play.
			r.Stat("Token source", tokenSourceLabel(res))
			// Reported, not instructed: nothing failed, and an environment
			// token winning is usually deliberate.
			if res.Shadowed {
				r.Warn("%s is overriding a different token saved in your config file.", res.EnvVar)
			}

			return login.SaveToken(res.Token)
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}
