package backup

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"net/url"
	"strings"
	"time"
)

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
			response, err := session.Connect(config.Token()).GetIndexBackup(args[0])
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
				return ui.Error(i18n.C.ErrorIndexRequired)
			}
			response, err := session.Connect(config.Token()).GetIndexBackup(args[0])
			if err != nil {
				return err
			}

			file, err := extractFile(response.GetData()[0].URL)

			date := parseDate(response.GetData()[0].DateAdded)

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
