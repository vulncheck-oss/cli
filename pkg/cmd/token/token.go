package token

import (
	"bytes"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
	"os"
	"text/tabwriter"
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

			for {
				selectedID, err := ui.TokensBrowse(tokens)
				if err != nil {
					return err
				}

				if selectedID == "" {
					// User quit the browse view
					return nil
				}

				token := tokenFromId(tokens, selectedID)
				if token != nil {
					if err := BrowseActions(*token, tokens); err != nil {
						return err
					}
					// If BrowseActions returns without error, continue the loop to show the token list again
				} else {
					return fmt.Errorf("selected token not found")
				}
			}

		},
	}

	return cmd
}

func BrowseActions(token sdk.TokenData, tokens []sdk.TokenData) error {

	// Define styles
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

	// Calculate the widest label for alignment
	labels := []string{"ID:", "Source:", "Location:", "Last Activity:"}
	maxLabelWidth := 0
	for _, label := range labels {
		if len(label) > maxLabelWidth {
			maxLabelWidth = len(label)
		}
	}

	// Set tab stop for alignment
	var buf bytes.Buffer
	tabWriter := tabwriter.NewWriter(&buf, maxLabelWidth, 0, 1, ' ', 0)

	// Write content to tabWriter
	fmt.Fprint(tabWriter, content)
	tabWriter.Flush()

	// Render the box with all content
	fmt.Println(boxStyle.Render(buf.String()))
	fmt.Println() // Add a newline after the box

	var action string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose an action").
				Options(
					huh.NewOption("Go back to token list", "back"),
					huh.NewOption("Quit", "quit"),
					huh.NewOption("Delete token", "delete"),
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
		ui.Info("chosen to delete")
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
