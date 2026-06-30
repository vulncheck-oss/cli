package token

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token <command>",
		Short: i18n.C.TokenShort,
	}

	cmd.AddCommand(List())
	cmd.AddCommand(Create())
	cmd.AddCommand(Remove())
	cmd.AddCommand(Browse())

	return cmd
}

func Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <label>",
		Short: i18n.C.CreateTokenShort,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) == 0 {
				return fmt.Errorf("%s", i18n.C.CreateTokenLabelRequired)
			}

			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).CreateToken(args[0])
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(response.Data)
			}

			r.Success(i18n.C.CreateTokenSuccess, args[0], response.Data.Token)
			return nil
		},
	}
	return cmd
}

func Remove() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove <id>",
		Short: i18n.C.RemoveTokenShort,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) == 0 {
				return fmt.Errorf("%s", i18n.C.RemoveTokenIDRequired)
			}

			_, err := session.ConnectWithContext(cmd.Context(), config.Token()).DeleteToken(args[0])
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(map[string]any{"id": args[0], "removed": true})
			}
			r.Success(i18n.C.RemoveTokenSuccess, args[0])
			return nil
		},
	}
	return cmd
}

func List() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list <search>",
		Short: i18n.C.ListTokensShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetTokens()
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(response.GetData())
			}

			r.Info(i18n.C.ListTokensFull, len(response.GetData()))
			return ui.TokensList(response.GetData())
		},
	}

	return cmd
}

func tokenFromId(tokens []sdk.TokenData, tokenId string) *sdk.TokenData {
	for _, token := range tokens {
		if token.ID == tokenId {
			return &token
		}
	}
	return nil
}

func Browse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "browse",
		Short: i18n.C.BrowseTokensShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if !r.Interactive() {
				return errs.Validation("token browse requires an interactive terminal; use `vulncheck token list --json` instead")
			}

			for {
				response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetTokens()
				if err != nil {
					return err
				}
				ui.ClearScreen()
				r.Info(i18n.C.BrowseTokens, len(response.GetData()))
				tokens := response.GetData()
				selectedID, err := ui.TokensBrowse(tokens)
				if err != nil {
					return err
				}

				if selectedID == "" {
					return nil
				}

				if selectedID == "createEntry" {
					if err := BrowseCreate(cmd.Context()); err != nil {
						return err
					}
					continue
				}

				token := tokenFromId(tokens, selectedID)
				if token != nil {
					if err := BrowseActions(cmd.Context(), *token, tokens); err != nil {
						return err
					}
				} else {
					return fmt.Errorf("selected token not found")
				}
			}
		},
	}

	return cmd
}

func BrowseCreate(ctx context.Context) error {
	var label string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter a label for the new token").
				Value(&label),
		),
	)

	err := form.Run()
	if err != nil {
		return err
	}

	if label == "" {
		return fmt.Errorf("%s", i18n.C.CreateTokenLabelRequired)
	}

	response, err := session.ConnectWithContext(ctx, config.Token()).CreateToken(label)
	if err != nil {
		return err
	}

	ui.ClearScreen()
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6667ab")).
		Padding(1, 1)

	tokenStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#34d399")).
		Bold(true)

	content := fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		"Token created",
		"Your new token (copy it now, it won't be shown again):",
		tokenStyle.Render(response.Data.Token),
	)

	fmt.Println(boxStyle.Render(content))
	fmt.Println("\nPress Enter to continue...")
	_, _ = fmt.Scanln()

	return nil
}

func BrowseActions(ctx context.Context, token sdk.TokenData, tokens []sdk.TokenData) error {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#6667ab")).
		Padding(0, 1)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6667ab")).
		Bold(true)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fff"))

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#34d39")).
		Bold(true).
		Padding(0, 1)

	content := fmt.Sprintf(
		"%s\n\n%s\t%s\n%s\t%s\n%s\t%s\n%s\t%s",
		titleStyle.Render("Token Details"),
		labelStyle.Render("ID:"),
		valueStyle.Render(token.ID),
		labelStyle.Render("Source:"),
		valueStyle.Render(token.GetSourceLabel()),
		labelStyle.Render("Location:"),
		valueStyle.Render(token.GetLocationString()),
		labelStyle.Render("Last Activity:"),
		valueStyle.Render(token.GetHumanUpdatedAt()),
	)

	labels := []string{"ID:", "Source:", "Location:", "Last Activity:"}
	maxLabelWidth := 0
	for _, label := range labels {
		if len(label) > maxLabelWidth {
			maxLabelWidth = len(label)
		}
	}

	var buf bytes.Buffer
	tabWriter := tabwriter.NewWriter(&buf, maxLabelWidth, 0, 1, ' ', 0)

	if _, err := fmt.Fprint(tabWriter, content); err != nil {
		return fmt.Errorf("failed to write to tabWriter: %w", err)
	}
	if err := tabWriter.Flush(); err != nil {
		return fmt.Errorf("failed to flush tabWriter: %w", err)
	}

	ui.ClearScreen()
	fmt.Println(boxStyle.Render(buf.String()))
	fmt.Println()

	var action string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose an action").
				Options(
					huh.NewOption("Go back to token list", "back"),
					huh.NewOption("Delete token", "delete"),
					huh.NewOption("Quit", "quit"),
				).
				Value(&action),
		),
	)

	err := form.Run()
	if err != nil {
		return err
	}

	switch action {
	case "delete":
		confirmed, err := BrowseActionConfirm(token.ID)
		if err != nil {
			return err
		}
		if confirmed {
			_, err := session.ConnectWithContext(ctx, config.Token()).DeleteToken(token.ID)
			if err != nil {
				return err
			}
			ui.Success(fmt.Sprintf(i18n.C.RemoveTokenSuccess, token.ID))
		}
		return nil
	case "back":
		return nil
	case "quit":
		os.Exit(0)
		return nil
	default:
		return fmt.Errorf("invalid action")
	}
}

func BrowseActionConfirm(tokenID string) (bool, error) {
	var confirmed bool

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("Are you sure you want to delete token %s?", tokenID)).
				Description("This action cannot be undone.").
				Value(&confirmed),
		),
	)

	err := form.Run()
	if err != nil {
		return false, err
	}

	return confirmed, nil
}
