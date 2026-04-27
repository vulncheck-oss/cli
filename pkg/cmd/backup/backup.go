package backup

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
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
func validateIndex(index string) (string, error) {
	indicesResponse, err := session.Connect(config.Token()).GetIndices()
	if err != nil {
		return "", err
	}

	// Create a map of indices to compare against
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

	// If the index is not present in the map but close matches exist, present
	// an interactive select so the user can choose the intended index
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

type UrlOptions struct {
	Json bool
}

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "backup <command>",
		Short: i18n.C.BackupShort,
	}

	opts := &UrlOptions{
		Json: false,
	}

	cmdUrl := &cobra.Command{
		Use:   "url <index>",
		Short: i18n.C.BackupUrlShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("index name is required")
			}

			index := args[0]
			client := session.Connect(config.Token())
			response, err := client.GetIndexBackup(index)

			// If GetIndexBackup fails due to a HTTP request error fallback
			// and attempt to validate the index name argument provided if
			// there was a typo/spelling error it will suggest similar names
			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(index)
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

			if opts.Json {
				ui.Json(response.GetData()[0])
				return nil
			}

			ui.Stat("Filename", response.GetData()[0].Filename)
			ui.Stat("SHA256", response.GetData()[0].Sha256)
			ui.Stat("Date Added", response.GetData()[0].DateAdded)
			ui.Stat("URL", response.GetData()[0].URL)
			return nil
		},
	}
	cmdUrl.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")

	cmdDownload := &cobra.Command{
		Use:   "download <index>",
		Short: i18n.C.BackupDownloadShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.IndexErrorRequired)
			}

			index := args[0]
			client := session.Connect(config.Token())
			response, err := client.GetIndexBackup(index)

			// If GetIndexBackup fails due to a HTTP request error fallback
			// and attempt to validate the index name argument provided if
			// there was a typo/spelling error it will suggest similar names
			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(index)
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

			ui.Info(fmt.Sprintf(i18n.C.BackupDownloadInfo, index, date))
			ui.Info(fmt.Sprintf(i18n.C.BackupDownloadProgress, file))
			if err := ui.Download(response.GetData()[0].URL, file); err != nil {
				return err
			}
			ui.Success(i18n.C.BackupDownloadComplete)
			return nil
		},
	}

	cmd.AddCommand(cmdUrl)
	cmd.AddCommand(cmdDownload)

	return cmd
}
