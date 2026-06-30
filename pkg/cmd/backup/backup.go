package backup

import (
	"context"
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

// validateIndex checks whether index exists. If it does not but close matches
// are found, an interactive select is presented so the user can pick one.
// Returns the confirmed index name, or an error if the name is unrecognised.
func validateIndex(ctx context.Context, index string, interactive bool) (string, error) {
	indicesResponse, err := session.ConnectWithContext(ctx, config.Token()).GetIndices()
	if err != nil {
		return "", err
	}

	var indexNames []string
	available := make(map[string]bool)
	for _, idx := range indicesResponse.GetData() {
		available[idx.Name] = true
		indexNames = append(indexNames, idx.Name)
	}

	if available[index] {
		return index, nil
	}

	suggestions := utils.SuggestFor(index, indexNames)
	if len(suggestions) == 0 {
		return "", fmt.Errorf("index '%s' does not exist", index)
	}

	if !interactive {
		joined := suggestions[0]
		for _, s := range suggestions[1:] {
			joined += ", " + s
		}
		return "", fmt.Errorf("index '%s' does not exist; did you mean: %s", index, joined)
	}

	options := make([]huh.Option[string], len(suggestions))
	for i, s := range suggestions {
		options[i] = huh.NewOption(s, s)
	}

	var selected string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(fmt.Sprintf("index '%s' does not exist. Did you mean one of these?", index)).
				Options(options...).
				Value(&selected),
		),
	)
	if err := form.Run(); err != nil {
		return "", fmt.Errorf("index '%s' does not exist", index)
	}

	return selected, nil
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "backup <command>",
		Short: i18n.C.BackupShort,
	}

	cmdUrl := &cobra.Command{
		Use:   "url <index>",
		Short: i18n.C.BackupUrlShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error("index name is required")
			}

			index := args[0]
			client := session.ConnectWithContext(cmd.Context(), config.Token())
			response, err := client.GetIndexBackup(index)

			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(cmd.Context(), index, r.Interactive())
					if validationErr != nil {
						return validationErr
					}
					response, err = client.GetIndexBackup(corrected)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}

			data := response.GetData()[0]
			if r.IsJSON() {
				return r.JSON(data)
			}

			r.Stat("Filename", data.Filename)
			r.Stat("SHA256", data.Sha256)
			r.Stat("Date Added", data.DateAdded)
			r.Stat("URL", data.URL)
			return nil
		},
	}

	cmdDownload := &cobra.Command{
		Use:   "download <index>",
		Short: i18n.C.BackupDownloadShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.IndexErrorRequired)
			}

			index := args[0]
			client := session.ConnectWithContext(cmd.Context(), config.Token())
			response, err := client.GetIndexBackup(index)

			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(cmd.Context(), index, r.Interactive())
					if validationErr != nil {
						return validationErr
					}
					index = corrected
					response, err = client.GetIndexBackup(index)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}

			file, err := utils.ExtractFileBasename(response.GetData()[0].URL)
			if err != nil {
				return err
			}

			date := utils.ParseDate(response.GetData()[0].DateAdded)

			r.Info(i18n.C.BackupDownloadInfo, index, date)
			r.Info(i18n.C.BackupDownloadProgress, file)
			if err := ui.Download(response.GetData()[0].URL, file); err != nil {
				return err
			}
			if r.IsJSON() {
				return r.JSON(map[string]any{
					"index":    index,
					"file":     file,
					"sha256":   response.GetData()[0].Sha256,
					"complete": true,
				})
			}
			r.Success("%s", i18n.C.BackupDownloadComplete)
			return nil
		},
	}

	cmd.AddCommand(cmdUrl)
	cmd.AddCommand(cmdDownload)

	return cmd
}
