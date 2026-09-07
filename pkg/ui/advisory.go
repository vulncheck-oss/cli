package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
	ltable "github.com/charmbracelet/lipgloss/table"
	"github.com/vulncheck-oss/cli/pkg/sdk"
)

// AdvisoryFeedRow pairs a feed with its backup availability.
//
// The v4 feed list is only {name, href} -- there is no description, which is
// the column that makes the v3 index catalogue worth rendering -- so the
// backup columns are what give this table something to say.
type AdvisoryFeedRow struct {
	Name string `json:"name"`
	// Href is carried through from the API payload rather than dropped, so this
	// row stays a superset of what /v4/advisory/list returned.
	Href string `json:"href"`
	// Available is nil when the backup join was unavailable, so a consumer can
	// tell "no backup" apart from "we could not find out".
	// No omitempty on either: v4 backup is a separate entitlement from v4
	// advisory, so omitting them would make the JSON shape depend on the
	// caller's token tier. The keys are always present, and a null Available
	// means the backup list could not be read.
	Available       *bool  `json:"available"`
	BackupWrittenAt string `json:"backup_written_at"`
}

// AdvisoryFeeds joins the feed catalogue with backup availability.
//
// backups may be nil: v4 backup is a separate entitlement from v4 advisory (a
// trial grants the latter but not the former), so the join is best-effort and
// the caller passes nil when /v4/backup was refused. Rendering then degrades to
// the name column rather than failing a discovery command the user is entitled
// to run.
func AdvisoryFeeds(feeds []sdk.AdvisoryFeedMeta, backups []sdk.AdvisoryBackupMeta, search string) []AdvisoryFeedRow {
	available := make(map[string]sdk.AdvisoryBackupMeta, len(backups))
	for _, b := range backups {
		available[b.Name] = b
	}

	rows := make([]AdvisoryFeedRow, 0, len(feeds))
	for _, feed := range feeds {
		if search != "" && !strings.Contains(feed.Name, search) {
			continue
		}
		row := AdvisoryFeedRow{Name: feed.Name, Href: feed.Href}
		if backup, ok := available[feed.Name]; ok {
			hasBackup := backup.Available
			row.Available = &hasBackup
			row.BackupWrittenAt = sdk.ShortDate(backup.BackupWrittenAt)
		}
		rows = append(rows, row)
	}
	return rows
}

func AdvisoryFeedsList(rows []AdvisoryFeedRow, withBackups bool) error {
	headers := []string{"Name"}
	if withBackups {
		headers = append(headers, "Backup", "Written")
	}

	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers(headers...).Width(TermWidth())

	for _, row := range rows {
		if withBackups {
			t.Row(row.Name, yesNo(row.Available), row.BackupWrittenAt)
			continue
		}
		t.Row(row.Name)
	}

	fmt.Println(t)
	return nil
}

func AdvisoryList(records []sdk.AdvisoryRecordView) error {
	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("Feed", "CVE", "Title", "CVSS", "Affected", "Refs", "Updated").Width(TermWidth())

	for _, r := range records {
		t.Row(r.Feed, r.CveID, r.Title, r.CVSS, strconv.Itoa(r.Affected), strconv.Itoa(r.References), r.Updated)
	}

	fmt.Println(t)
	return nil
}

// yesNo renders a tri-state backup flag: unknown when the join was refused.
func yesNo(value *bool) string {
	switch {
	case value == nil:
		return ""
	case *value:
		return "yes"
	default:
		return "no"
	}
}

func AdvisoryBackupsList(backups []sdk.AdvisoryBackupMeta) error {
	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("Feed", "Available", "Written").Width(TermWidth())

	for _, b := range backups {
		available := "no"
		if b.Available {
			available = "yes"
		}
		t.Row(b.Name, available, sdk.ShortDate(b.BackupWrittenAt))
	}

	fmt.Println(t)
	return nil
}

// BackupsList renders the v3 backup catalogue (GET /v3/backup).
//
// Href is deliberately not a column: it is api.vulncheck.com/v3/backup/<name>
// for every row, so it spends a third of the terminal width restating the name.
// --json still carries it.
func BackupsList(backups []sdk.BackupsMeta) error {
	t := ltable.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(lipgloss.Color("#6667ab"))).
		Headers("Name", "Description").Width(TermWidth())

	for _, b := range backups {
		t.Row(b.Name, b.Description)
	}

	fmt.Println(t)
	return nil
}
