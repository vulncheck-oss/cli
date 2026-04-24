package backup

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func validateIndex(index string) (*sdk.Client, error) {
	client := session.Connect(config.Token())
	indicesResponse, err := client.GetIndices()
	if err != nil {
		return nil, err
	}

	// Create a map of indices to compare against
	var indexNames []string
	available := make(map[string]bool)
	for _, idx := range indicesResponse.GetData() {
		available[idx.Name] = true
		indexNames = append(indexNames, idx.Name)
	}

	// If the index is not present in the map, print output and suggest
	// options for what the argument may have intended
	if !available[index] {
		var msg strings.Builder
		fmt.Fprintf(&msg, "index '%s' does not exist", index)
		if suggestions := utils.SuggestFor(index, indexNames); len(suggestions) > 0 {
			msg.WriteString("\n\nDid you mean this?\n")
			for _, s := range suggestions {
				fmt.Fprintf(&msg, "\t%s\n", s)
			}
		}
		return nil, fmt.Errorf("%s", msg.String())
	}
	return client, nil
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

			client, err := validateIndex(args[0])
			if err != nil {
				return err
			}

			response, err := client.GetIndexBackup(args[0])
			if err != nil {
				return err
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

			client, err := validateIndex(args[0])
			if err != nil {
				return err
			}

			response, err := client.GetIndexBackup(args[0])
			if err != nil {
				return err
			}

			file, err := utils.ExtractFileBasename(response.GetData()[0].URL)
			if err != nil {
				return err
			}

			date := utils.ParseDate(response.GetData()[0].DateAdded)

			ui.Info(fmt.Sprintf(i18n.C.BackupDownloadInfo, args[0], date))
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
