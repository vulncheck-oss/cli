package ui

import (
	"encoding/json"
	"sort"
	"strings"
	"testing"

	"github.com/vulncheck-oss/cli/pkg/sdk"
)

var testFeeds = []sdk.AdvisoryFeedMeta{
	{Name: "epss", Href: "http://api.vulncheck.com/v4/advisory?name=epss"},
	{Name: "ghsa", Href: "http://api.vulncheck.com/v4/advisory?name=ghsa"},
	{Name: "sigmahq-sigma-rules"},
}

// --json is the payload untouched everywhere else, so the row must stay a
// superset of what /v4/advisory/list returned rather than dropping href.
func TestAdvisoryFeedsCarriesHrefThrough(t *testing.T) {
	rows := AdvisoryFeeds(testFeeds, nil, "")
	if rows[0].Href != "http://api.vulncheck.com/v4/advisory?name=epss" {
		t.Errorf("Href = %q, want the API's value", rows[0].Href)
	}
}

// An empty result must marshal as [] like advisory list, not null.
func TestAdvisoryFeedsReturnsAnEmptySliceNotNil(t *testing.T) {
	rows := AdvisoryFeeds(testFeeds, nil, "zzzznomatch")
	if rows == nil {
		t.Fatal("rows must not be nil")
	}
	if len(rows) != 0 {
		t.Errorf("rows = %v, want empty", rows)
	}
}

// v4 backup is a separate entitlement from v4 advisory -- a trial grants the
// latter but not the former -- so the join must degrade rather than fail a
// discovery command the caller is entitled to run.
func TestAdvisoryFeedsDegradesWithoutBackups(t *testing.T) {
	rows := AdvisoryFeeds(testFeeds, nil, "")

	if len(rows) != 3 {
		t.Fatalf("rows = %d, want 3", len(rows))
	}
	for _, row := range rows {
		if row.Available != nil {
			t.Errorf("%s: Available = %v, want nil when the join was unavailable", row.Name, *row.Available)
		}
		if row.BackupWrittenAt != "" {
			t.Errorf("%s: BackupWrittenAt = %q, want empty", row.Name, row.BackupWrittenAt)
		}
	}
}

func TestAdvisoryFeedsJoinsBackupAvailability(t *testing.T) {
	backups := []sdk.AdvisoryBackupMeta{
		{Name: "epss", Available: true, BackupWrittenAt: "2026-09-02T16:17:49Z"},
		{Name: "ghsa", Available: false},
	}

	rows := AdvisoryFeeds(testFeeds, backups, "")

	if rows[0].Available == nil || !*rows[0].Available {
		t.Error("epss should report an available backup")
	}
	if rows[0].BackupWrittenAt != "2026-09-02" {
		t.Errorf("BackupWrittenAt = %q, want the date only", rows[0].BackupWrittenAt)
	}
	if rows[1].Available == nil || *rows[1].Available {
		t.Error("ghsa should report no available backup")
	}
	// A feed absent from the backup list is unknown, not unavailable.
	if rows[2].Available != nil {
		t.Error("a feed missing from the backup list should stay unknown")
	}
}

// v4 backup is a separate entitlement from v4 advisory, so the JSON key set
// must not depend on whether the join succeeded -- a script parsing available
// should see null, not a missing field.
func TestAdvisoryFeedsJSONShapeIsIndependentOfEntitlement(t *testing.T) {
	withJoin, err := json.Marshal(AdvisoryFeeds(testFeeds, []sdk.AdvisoryBackupMeta{{Name: "epss", Available: true}}, "")[0])
	if err != nil {
		t.Fatal(err)
	}
	withoutJoin, err := json.Marshal(AdvisoryFeeds(testFeeds, nil, "")[0])
	if err != nil {
		t.Fatal(err)
	}

	keys := func(raw []byte) []string {
		var m map[string]any
		if err := json.Unmarshal(raw, &m); err != nil {
			t.Fatal(err)
		}
		out := make([]string, 0, len(m))
		for k := range m {
			out = append(out, k)
		}
		sort.Strings(out)
		return out
	}

	a, b := keys(withJoin), keys(withoutJoin)
	if strings.Join(a, ",") != strings.Join(b, ",") {
		t.Errorf("key sets differ: with join %v, without %v", a, b)
	}
	if !strings.Contains(string(withoutJoin), `"available":null`) {
		t.Errorf("unknown availability should be null, got %s", withoutJoin)
	}
}

func TestAdvisoryFeedsFiltersBySearch(t *testing.T) {
	rows := AdvisoryFeeds(testFeeds, nil, "sigma")
	if len(rows) != 1 || rows[0].Name != "sigmahq-sigma-rules" {
		t.Errorf("rows = %v, want just sigmahq-sigma-rules", rows)
	}
}

func TestYesNoIsTriState(t *testing.T) {
	yes, no := true, false
	if got := yesNo(nil); got != "" {
		t.Errorf("yesNo(nil) = %q, want empty", got)
	}
	if got := yesNo(&yes); got != "yes" {
		t.Errorf("yesNo(true) = %q, want yes", got)
	}
	if got := yesNo(&no); got != "no" {
		t.Errorf("yesNo(false) = %q, want no", got)
	}
}
