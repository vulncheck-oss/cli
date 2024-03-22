package backup

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/environment"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk"
	"net/url"
	"strings"
	"time"
)

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "backup <command>",
		Short: "Download a backup of a specified index",
	}

	cmdUrl := &cobra.Command{
		Use:   "url <index>",
		Short: "Get the temporary signed URL of the backup of an index",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("index name is required")
			}
			response, err := sdk.Connect(environment.Env.API, config.Token()).GetIndexBackup(args[0])
			if err != nil {
				return err
			}
			ui.Json(response.GetData()[0])
			return nil
		},
	}

	cmdDownload := &cobra.Command{
		Use:   "download <index>",
		Short: "Download the backup of an index",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error("index name is required")
			}
			response, err := sdk.Connect(environment.Env.API, config.Token()).GetIndexBackup(args[0])
			if err != nil {
				return err
			}

			file, err := extractFile(response.GetData()[0].URL)

			date := parseDate(response.GetData()[0].DateAdded)

			ui.Info(fmt.Sprintf("Backup of %s found, created on %s", args[0], date))
			ui.Info(fmt.Sprintf("Downloading backup as %s ", file))
			if err := ui.Download(response.GetData()[0].URL, file); err != nil {
				return err
			}
			ui.Success("Backup downloaded successfully")
			return nil
		},
	}

	cmd.AddCommand(cmdUrl)
	cmd.AddCommand(cmdDownload)

	return cmd
}

func parseDate(date string) string {
	dateAdded, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s, %s, %s",
		dateAdded.Format("January 2, 2006"), // Format date
		dateAdded.Format("3:04:05 am"),      // Format time
		dateAdded.Format("MST"),             // Format timezone
	)
}

func extractFile(urlStr string) (string, error) {

	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		return "", err
	}

	path := strings.TrimPrefix(parsedUrl.Path, "/")
	if !strings.HasSuffix(path, ".zip") {
		return "", fmt.Errorf("invalid file format")
	}

	return path, nil
}
