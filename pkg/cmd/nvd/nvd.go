package nvd

import (
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
)

// The nvd commands mirror VulnCheck's NVD API 2.0 compatible endpoints and nothing more:
// one subcommand per endpoint, one flag per query parameter, values sent unchanged, and
// validation left to the API, which answers a bad value with a 400 naming the parameter.
// Each command makes exactly one request; callers page with --cursor (or --start-index
// on `nvd source`), the way the API itself does.
//
// The parameters are written out per endpoint rather than derived, for the same reason
// pkg/cmd/advisory gives: the set that reaches the wire wants to be read and reviewed.

type kind int

const (
	text kind = iota
	number
	flag
)

// param is one API query parameter and the flag that sends it.
type param struct {
	name  string // the API's parameter name, sent on the wire
	flag  string // the CLI flag, the parameter name in kebab-case
	kind  kind
	usage string
}

type endpoint struct {
	use     string
	short   func() string
	path    string
	dataKey string // the response key holding the record array
	params  []param
}

// The paging parameters every endpoint takes. /source/2.0 takes no cursor.
var (
	resultsPerPage = param{"resultsPerPage", "results-per-page", number, "Results per page (the endpoint's maximum is also its default)"}
	startIndex     = param{"startIndex", "start-index", number, "Zero-based offset of the first result"}
	cursor         = param{"cursor", "cursor", text, `Cursor paging: pass "" for the first page, then each response's nextCursor (see --full)`}
)

func lastModParams() []param {
	return []param{
		{"lastModStartDate", "last-mod-start-date", text, "Last modified on or after, ISO-8601. Requires --last-mod-end-date"},
		{"lastModEndDate", "last-mod-end-date", text, "Last modified on or before, ISO-8601. Requires --last-mod-start-date"},
	}
}

func endpoints() []endpoint {
	return []endpoint{
		{
			use:     "cves",
			short:   func() string { return i18n.C.NvdCvesShort },
			path:    sdk.NvdCvesPath,
			dataKey: "vulnerabilities",
			params: append(append([]param{
				{"cveId", "cve-id", text, "A single CVE ID, e.g. CVE-2021-44228"},
				{"cveIds", "cve-ids", text, "Comma-separated CVE IDs"},
				{"cpeName", "cpe-name", text, "A complete CPE 2.3 name; CVEs whose configurations match it"},
				{"isVulnerable", "is-vulnerable", flag, "Only CVEs where --cpe-name is marked vulnerable"},
				{"virtualMatchString", "virtual-match-string", text, "A CPE 2.3 name or prefix matched against configuration criteria"},
				{"versionStart", "version-start", text, "Lower version bound; requires --virtual-match-string and --version-start-type"},
				{"versionStartType", "version-start-type", text, "including or excluding"},
				{"versionEnd", "version-end", text, "Upper version bound; requires --virtual-match-string and --version-end-type"},
				{"versionEndType", "version-end-type", text, "including or excluding"},
				{"vulnStatuses", "vuln-statuses", text, "Comma-separated vulnStatus values, e.g. Analyzed,Modified"},
				{"cveTag", "cve-tag", text, "disputed, unsupported-when-assigned or exclusively-hosted-service"},
				{"noRejected", "no-rejected", flag, "Exclude rejected CVEs"},
				{"hasKev", "has-kev", flag, "Only CVEs in the CISA KEV catalog"},
				{"vcVulnerableCPEs", "vc-vulnerable-cpes", flag, "Include VulnCheck's flat vcVulnerableCPEs list on each record"},
				{"pubStartDate", "pub-start-date", text, "Published on or after, ISO-8601. Requires --pub-end-date"},
				{"pubEndDate", "pub-end-date", text, "Published on or before, ISO-8601. Requires --pub-start-date"},
			}, lastModParams()...),
				param{"sourceIdentifier", "source-identifier", text, "Only CVEs from this source, e.g. cve@mitre.org"},
				resultsPerPage, startIndex, cursor,
			),
		},
		{
			use:     "cvehistory",
			short:   func() string { return i18n.C.NvdCveHistoryShort },
			path:    sdk.NvdCveHistoryPath,
			dataKey: "cveChanges",
			params: []param{
				{"cveId", "cve-id", text, "Only changes to this CVE"},
				{"eventName", "event-name", text, `Only changes of this type, e.g. "CVE Modified"`},
				{"changeStartDate", "change-start-date", text, "Changed on or after, ISO-8601. Requires --change-end-date"},
				{"changeEndDate", "change-end-date", text, "Changed on or before, ISO-8601. Requires --change-start-date"},
				resultsPerPage, startIndex, cursor,
			},
		},
		{
			use:     "cpes",
			short:   func() string { return i18n.C.NvdCpesShort },
			path:    sdk.NvdCpesPath,
			dataKey: "products",
			params: append(append([]param{
				{"cpeNameId", "cpe-name-id", text, "A single dictionary entry by its NVD UUID"},
				{"cpeMatchString", "cpe-match-string", text, "A CPE 2.3 name or prefix, e.g. cpe:2.3:a:apache:log4j"},
				{"matchCriteriaId", "match-criteria-id", text, "The CPE names this match criteria UUID resolves to"},
			}, lastModParams()...),
				resultsPerPage, startIndex, cursor,
			),
		},
		{
			use:     "cpematch",
			short:   func() string { return i18n.C.NvdCpeMatchShort },
			path:    sdk.NvdCpeMatchPath,
			dataKey: "matchStrings",
			params: append(append([]param{
				{"cveId", "cve-id", text, "The match criteria in this CVE's configurations"},
				{"matchCriteriaId", "match-criteria-id", text, "A single match criteria by its NVD UUID"},
			}, lastModParams()...),
				resultsPerPage, startIndex, cursor,
			),
		},
		{
			use:     "source",
			short:   func() string { return i18n.C.NvdSourceShort },
			path:    sdk.NvdSourcePath,
			dataKey: "sources",
			params: append(append([]param{
				{"sourceIdentifier", "source-identifier", text, "Only the source with this identifier, e.g. cve@mitre.org"},
			}, lastModParams()...),
				resultsPerPage, startIndex,
			),
		},
	}
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "nvd <command>",
		Short: i18n.C.NvdShort,
	}
	for _, ep := range endpoints() {
		cmd.AddCommand(subcommand(ep))
	}
	return cmd
}

func subcommand(ep endpoint) *cobra.Command {
	texts := map[string]*string{}
	numbers := map[string]*int{}
	flags := map[string]*bool{}
	var full bool

	cmd := &cobra.Command{
		Use:   ep.use + " [flags]",
		Short: ep.short(),
		// Every filter is a flag. Accepting a positional would drop it and answer a
		// broader question than the one asked.
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			var params []sdk.NvdParam
			for _, p := range ep.params {
				// Changed rather than non-zero: an empty --cursor is a real request,
				// and an unset number must leave the API's default in place.
				if !cmd.Flags().Changed(p.flag) {
					continue
				}
				var value string
				switch p.kind {
				case number:
					value = strconv.Itoa(*numbers[p.flag])
				case flag:
					value = strconv.FormatBool(*flags[p.flag])
				default:
					value = *texts[p.flag]
				}
				params = append(params, sdk.NvdParam{Name: p.name, Value: value})
			}

			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetNvd(ep.path, params)
			if err != nil {
				return err
			}

			r.Info(i18n.C.NvdFound, response.ResultsPerPage, response.TotalResults)
			if response.NextCursor != "" {
				r.Info(i18n.C.NvdNextCursor, response.NextCursor)
			}

			if full {
				return r.JSON(response.Raw())
			}
			return r.JSON(response.Records(ep.dataKey))
		},
	}

	for _, p := range ep.params {
		switch p.kind {
		case number:
			numbers[p.flag] = cmd.Flags().Int(p.flag, 0, p.usage)
		case flag:
			flags[p.flag] = cmd.Flags().Bool(p.flag, false, p.usage)
		default:
			texts[p.flag] = cmd.Flags().String(p.flag, "", p.usage)
		}
	}
	cmd.Flags().BoolVarP(&full, "full", "f", false, i18n.C.NvdFlagFullResponse)

	return cmd
}
