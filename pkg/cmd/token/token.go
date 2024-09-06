package token

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
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

type ListOptions struct {
	Json bool
}

func Create() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create <label>",
		Short: i18n.C.CreateTokenShort,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return fmt.Errorf(i18n.C.CreateTokenLabelRequired)
			}

			response, err := session.Connect(config.Token()).CreateToken(args[0])
			if err != nil {
				return err
			}
			ui.Success(fmt.Sprintf(i18n.C.CreateTokenSuccess, args[0], response.Data.Token))
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
			if len(args) == 0 {
				return fmt.Errorf(i18n.C.RemoveTokenIDRequired)
			}

			_, err := session.Connect(config.Token()).DeleteToken(args[0])
			if err != nil {
				return err
			}
			ui.Success(fmt.Sprintf(i18n.C.RemoveTokenSuccess, args[0]))
			return nil
		},
	}
	return cmd
}

func List() *cobra.Command {

	opts := &ListOptions{
		Json: false,
	}

	cmd := &cobra.Command{
		Use:   "list <search>",
		Short: i18n.C.ListTokensShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := session.Connect(config.Token()).GetTokens()
			if err != nil {
				return err
			}
			ui.Info(fmt.Sprintf(i18n.C.ListTokensFull, len(response.GetData())))
			if opts.Json {
				ui.Json(response.GetData())
				return nil
			}

			if err := ui.TokensList(response.GetData()); err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
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
			response, err := session.Connect(config.Token()).GetTokens()
			if err != nil {
				return err
			}
			ui.Info(fmt.Sprintf(i18n.C.ListTokensFull, len(response.GetData())))

			tokens := response.GetData()

			return BrowseTokens(tokens)
		},
	}

	return cmd
}

func BrowseTokens(tokens []sdk.TokenData) error {
	action := func(tokenId string) error {
		token := tokenFromId(tokens, tokenId)
		if token == nil {
			return fmt.Errorf("token not found")
		}
		return BrowseActions(*token, tokens)
	}

	return ui.TokensBrowse(tokens, action)
}

func clearScreen() {
	tea.EnterAltScreen()
}

func BrowseActions(token sdk.TokenData, tokens []sdk.TokenData) error {
	clearScreen()
	var action string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose an action").
				Options(
					huh.NewOption("Delete token", "delete"),
					huh.NewOption("Go back to token list", "back"),
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
		clearScreen()
		ui.Info("chosen to delete")
		return nil
	case "back":
		clearScreen()
		BrowseTokens(tokens)
		return nil // This will return to the token list
	default:
		return fmt.Errorf("invalid action")
	}
}
