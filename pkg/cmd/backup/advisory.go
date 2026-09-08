package backup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

// validateFeed mirrors validateIndex for the v4 feed namespace. Unlike the
// advisory query endpoint, /v4/backup/{feed} does return 404 for an unknown
// feed, so this is only reached once a request has already failed.
func validateFeed(ctx context.Context, feed string, interactive bool) (string, error) {
	response, err := session.ConnectWithContext(ctx, config.Token()).GetAdvisoryBackups()
	if err != nil {
		return "", err
	}
	if response == nil {
		return "", errs.Validation("advisory feed '%s' does not exist", feed)
	}

	var names []string
	available := make(map[string]bool)
	for _, f := range response.GetData() {
		available[f.Name] = true
		names = append(names, f.Name)
	}

	if available[feed] {
		return feed, nil
	}

	suggestions := utils.SuggestFor(feed, names)
	if len(suggestions) == 0 {
		return "", errs.Validation("advisory feed '%s' does not exist", feed)
	}

	if !interactive {
		joined := suggestions[0]
		for _, s := range suggestions[1:] {
			joined += ", " + s
		}
		return "", errs.Validation("advisory feed '%s' does not exist; did you mean: %s", feed, joined)
	}

	options := make([]huh.Option[string], len(suggestions))
	for i, s := range suggestions {
		options[i] = huh.NewOption(s, s)
	}

	var selected string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(fmt.Sprintf("advisory feed '%s' does not exist. Did you mean one of these?", feed)).
				Options(options...).
				Value(&selected),
		),
	)
	if err := form.Run(); err != nil || selected == "" {
		return "", errs.Validation("advisory feed '%s' does not exist", feed)
	}

	return selected, nil
}

// fetchAdvisoryBackup resolves a feed to its backup, retrying once against a
// corrected feed name when the first attempt 404s.
func fetchAdvisoryBackup(cmd *cobra.Command, feed string, interactive bool) (*sdk.AdvisoryBackup, string, error) {
	client := session.ConnectWithContext(cmd.Context(), config.Token())

	backup, err := client.GetAdvisoryBackup(feed)
	if err != nil {
		// Only a 404 means "no such feed". Re-resolving an entitlement or
		// server error through the catalogue costs an extra request and cannot
		// change the outcome.
		reqErr, ok := err.(sdk.ReqError)
		if !ok || reqErr.StatusCode != http.StatusNotFound {
			return nil, feed, err
		}
		corrected, validationErr := validateFeed(cmd.Context(), feed, interactive)
		if validationErr != nil {
			return nil, feed, validationErr
		}
		feed = corrected
		if backup, err = client.GetAdvisoryBackup(feed); err != nil {
			return nil, feed, err
		}
	}

	if backup == nil {
		return nil, feed, errs.New(errs.KindInternal, "empty response from /v4/backup/%s", feed)
	}
	if !backup.Available || backup.URL == "" {
		return nil, feed, errs.NotFound(i18n.C.BackupAdvisoryUnavailable, feed)
	}

	return backup, feed, nil
}

// advisoryBackupOutput adds provenance the API does not send. A downloaded
// abbott.zip is indistinguishable on disk from its v3 namesake, so --json names
// the corpus and the archive format explicitly.
type advisoryBackupOutput struct {
	*sdk.AdvisoryBackup
	Corpus string `json:"corpus"`
	Format string `json:"format"`
}

const (
	advisoryCorpus = "v4-advisory"
	advisoryFormat = "zip/jsonl"
)

// advisoryCommand groups the v4 advisory feed backups.
//
// These sit under `backup` rather than beside the advisory query commands so
// that everything bulk-download-shaped is discoverable in one place, and they
// are named for the corpus rather than the API version because that is the real
// distinction: a v4 zip holds CVE 5.2 records, one per CVE, while the v3 zip for
// the same name holds that feed's native records including any without a CVE.
func advisoryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "advisory <command>",
		Short: i18n.C.BackupAdvisoryShort,
	}

	cmdList := &cobra.Command{
		Use:   "list",
		Short: i18n.C.BackupAdvisoryListShort,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetAdvisoryBackups()
			if err != nil {
				return err
			}
			if response == nil {
				return errs.New(errs.KindInternal, "empty response from /v4/backup")
			}

			if r.IsJSON() {
				return r.JSON(response.GetData())
			}

			r.Info(i18n.C.BackupAdvisoryListFull, len(response.GetData()))
			return ui.AdvisoryBackupsList(response.GetData())
		},
	}

	cmdUrl := &cobra.Command{
		Use:   "url <feed>",
		Short: i18n.C.BackupAdvisoryUrlShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.BackupAdvisoryErrorRequired)
			}

			backup, _, err := fetchAdvisoryBackup(cmd, args[0], r.Interactive())
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(advisoryBackupOutput{
					AdvisoryBackup: backup,
					Corpus:         advisoryCorpus,
					Format:         advisoryFormat,
				})
			}

			// v4 carries no filename or date_added -- backup_written_at lives
			// on the list endpoint -- so the checksum and expiry are what there
			// is to report. Corpus is named because the same feed name also
			// addresses a v3 index backup holding different data.
			r.Stat("Feed", backup.Feed)
			r.Stat("Corpus", advisoryCorpus)
			r.Stat("Format", advisoryFormat)
			r.Stat("SHA256", backup.SHA256)
			r.Stat("Expires", backup.URLExpires)
			r.Stat("URL", backup.URL)
			return nil
		},
	}

	cmdDownload := &cobra.Command{
		Use:   "download <feed>",
		Short: i18n.C.BackupAdvisoryDownloadShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.BackupAdvisoryErrorRequired)
			}

			backup, feed, err := fetchAdvisoryBackup(cmd, args[0], r.Interactive())
			if err != nil {
				return err
			}

			// The v4 object key carries no timestamp, so this resolves to
			// "<feed>.zip" and a repeat download overwrites the previous one.
			// That is distinct from the v3 name, which embeds a timestamp.
			file, err := utils.BackupFilename(backup.URL)
			if err != nil {
				return err
			}

			r.Info(i18n.C.BackupAdvisoryDownloadInfo, feed)
			r.Info(i18n.C.BackupDownloadProgress, file)
			r.Warn(i18n.C.BackupAdvisoryOverwrite, file)

			downloadErr := func() error {
				if r.Interactive() {
					return ui.Download(backup.URL, file)
				}
				return ui.DownloadHeadless(cmd.Context(), backup.URL, file, r.Stderr())
			}()
			if downloadErr != nil {
				return downloadErr
			}

			if r.IsJSON() {
				return r.JSON(map[string]any{
					"feed":     feed,
					"corpus":   advisoryCorpus,
					"format":   advisoryFormat,
					"file":     file,
					"sha256":   backup.SHA256,
					"complete": true,
				})
			}
			r.Success("%s", i18n.C.BackupDownloadComplete)
			return nil
		},
	}

	cmd.AddCommand(cmdList)
	cmd.AddCommand(cmdUrl)
	cmd.AddCommand(cmdDownload)

	return cmd
}
