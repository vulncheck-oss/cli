package advisory

import (
	"context"
	"fmt"
	"strings"

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

// queryOptions holds the flags shared by `list` and `browse`.
//
// The flags are declared explicitly rather than reflected off the parameter
// struct, as pkg/cmd/index does. Three reasons: /v4/advisory needs int and bool
// parameters as well as strings, --feed has to be renamed off its "name" wire
// parameter, and the API answers an unrecognised parameter with an empty result
// set instead of an error -- so the set that reaches the wire wants to be
// written out and reviewed, not derived.
type queryOptions struct {
	feed            string
	cve             string
	vendor          string
	product         string
	platform        string
	version         string
	cpe             string
	purl            string
	packageName     string
	referenceURL    string
	referenceTag    string
	descriptionLang string
	updatedAfter    string
	updatedBefore   string
	limit           int
	page            int
	cursor          string
	startCursor     bool
	all             bool
	full            bool
}

func (o *queryOptions) register(cmd *cobra.Command) {
	cmd.Flags().StringVar(&o.feed, "feed", "", "Filter by advisory feed name (see 'vulncheck advisory feeds')")
	cmd.Flags().StringVar(&o.cve, "cve", "", "Filter by CVE ID, e.g. CVE-2019-12255")
	cmd.Flags().StringVar(&o.vendor, "vendor", "", "Filter by vendor name")
	cmd.Flags().StringVar(&o.product, "product", "", "Filter by product name")
	cmd.Flags().StringVar(&o.platform, "platform", "", "Filter by OS or platform")
	cmd.Flags().StringVar(&o.version, "version", "", "Filter by product version (semver aware)")
	cmd.Flags().StringVar(&o.cpe, "cpe", "", "Filter by CPE")
	cmd.Flags().StringVar(&o.purl, "purl", "", "Filter by package URL")
	cmd.Flags().StringVar(&o.packageName, "package-name", "", "Filter by package name")
	cmd.Flags().StringVar(&o.referenceURL, "reference-url", "", "Filter by reference URL")
	cmd.Flags().StringVar(&o.referenceTag, "reference-tag", "", "Filter by reference tag")
	cmd.Flags().StringVar(&o.descriptionLang, "description-lang", "", "Filter by description language")
	cmd.Flags().StringVar(&o.updatedAfter, "updated-after", "", "Updated after an RFC3339 date or date-math, e.g. now-30d")
	cmd.Flags().StringVar(&o.updatedBefore, "updated-before", "", "Updated before an RFC3339 date or date-math")
	cmd.Flags().IntVar(&o.limit, "limit", 0,
		fmt.Sprintf("Records per page (max %d, default %d); ignored by --all", sdk.MaxAdvisoryLimit, sdk.DefaultAdvisoryLimit))
	cmd.Flags().IntVar(&o.page, "page", 0, "Page number; page x limit may not exceed 10000")
	cmd.Flags().StringVar(&o.cursor, "cursor", "",
		"Continue from a next_cursor value emitted by --full; the only way past the result window without --all")
	cmd.Flags().BoolVar(&o.startCursor, "start-cursor", false,
		"Begin cursor pagination, so --full reports a next_cursor to continue from")
	cmd.Flags().BoolVar(&o.all, "all", false,
		fmt.Sprintf("Auto-paginate via cursor at the maximum page size (%d) and emit every matching record; cannot be combined with --page", sdk.MaxAdvisoryLimit))
	cmd.Flags().BoolVarP(&o.full, "full", "f", false, "Output the full response envelope, including _meta")
}

func (o *queryOptions) params() sdk.AdvisoryQueryParameters {
	return sdk.AdvisoryQueryParameters{
		Name:            o.feed,
		CveID:           o.cve,
		Vendor:          o.vendor,
		Product:         o.product,
		Platform:        o.platform,
		Version:         o.version,
		CPE:             o.cpe,
		PURL:            o.purl,
		PackageName:     o.packageName,
		ReferenceURL:    o.referenceURL,
		ReferenceTag:    o.referenceTag,
		DescriptionLang: o.descriptionLang,
		UpdatedAfter:    o.updatedAfter,
		UpdatedBefore:   o.updatedBefore,
		Limit:           o.limit,
		Page:            o.page,
		Cursor:          o.cursor,
		StartCursor:     o.startCursor,
	}
}

// validate turns the SDK's pre-flight checks into classified CLI errors so they
// exit 2 as validation problems rather than 1 as internal ones.
func (o *queryOptions) validate() error {
	params := o.params()
	if o.all && o.page > 0 {
		return errs.Validation("--all cannot be combined with --page; --all pages through everything via cursor")
	}
	if o.all && (o.cursor != "" || o.startCursor) {
		return errs.Validation("--all cannot be combined with --cursor or --start-cursor; it drives cursor pagination itself")
	}
	if o.all && o.limit > 0 && o.limit != sdk.MaxAdvisoryLimit {
		return errs.Validation(
			"--all cannot be combined with --limit %d; it always walks at the maximum page size of %d",
			o.limit, sdk.MaxAdvisoryLimit)
	}
	if o.all {
		// --all drives cursor pagination, so the offset ceiling does not apply.
		params.Page = 0
	}
	if err := params.Validate(); err != nil {
		return errs.Validation("%s", err)
	}
	return nil
}

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "advisory <command>",
		Short: i18n.C.AdvisoryShort,
	}

	cmd.AddCommand(Feeds())
	cmd.AddCommand(List())
	cmd.AddCommand(Browse())

	return cmd
}

// Feeds lists the advisory feed catalogue.
//
// This is GET /v4/advisory/list. It is spelled `feeds` rather than `list`
// because the API and the CLI use "list" in opposite senses: in the URL /list
// is the catalogue, while in the CLI `list` means records, as it does for
// `vulncheck index list`. Keeping the CLI self-consistent was judged the more
// useful of the two consistencies.
func Feeds() *cobra.Command {
	return &cobra.Command{
		Use:   "feeds [search]",
		Short: i18n.C.AdvisoryFeedsShort,
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			search := ""
			if len(args) > 0 {
				search = args[0]
			}

			client := session.ConnectWithContext(cmd.Context(), config.Token())
			response, err := client.GetAdvisoryFeeds()
			if err != nil {
				return err
			}
			if response == nil {
				return errs.New(errs.KindInternal, "empty response from /v4/advisory/list")
			}

			// Best-effort: v4 backup is a separate entitlement from v4
			// advisory, so a caller entitled to the feed list may still be
			// refused the backup list. Degrade to the name column rather than
			// failing a command they are entitled to run.
			var backups []sdk.AdvisoryBackupMeta
			withBackups := false
			if backupResponse, backupErr := client.GetAdvisoryBackups(); backupErr == nil && backupResponse != nil {
				backups = backupResponse.GetData()
				withBackups = true
			}

			rows := ui.AdvisoryFeeds(response.GetData(), backups, search)

			if r.IsJSON() {
				return r.JSON(rows)
			}

			if search != "" {
				r.Info(i18n.C.AdvisoryFeedsSearch, len(rows), search)
			} else {
				r.Info(i18n.C.AdvisoryFeedsFull, len(rows))
			}
			return ui.AdvisoryFeedsList(rows, withBackups)
		},
	}
}

func List() *cobra.Command {
	opts := &queryOptions{}

	cmd := &cobra.Command{
		Use:   "list [flags]",
		Short: i18n.C.AdvisoryListShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			if err := opts.validate(); err != nil {
				return err
			}

			client := session.ConnectWithContext(cmd.Context(), config.Token())
			response, pages, err := fetch(cmd.Context(), client, opts, r.Interactive())
			if err != nil {
				return err
			}

			if r.IsJSON() {
				if opts.full {
					return r.JSON(response)
				}
				return r.JSON(response.GetData())
			}

			views := sdk.ParseAdvisoryRecords(response.GetData())
			r.Info(i18n.C.AdvisoryFound, len(views), response.Meta.Total)
			if opts.full {
				// Without this --full would be a silent no-op outside --json.
				//
				// A cursor-mode response always echoes page 1, so after a
				// completed walk "Page 1 of 32" would say the opposite of what
				// happened; the terminal cursor is equally meaningless there.
				// Under --all these are facts about the request rather than
				// the response: the synthesised envelope reports one page of
				// everything, so echoing its limit here would just restate
				// the record count.
				if opts.all {
					r.Stat("Pages walked", fmt.Sprintf("%d", pages))
					r.Stat("Page size", fmt.Sprintf("%d", sdk.MaxAdvisoryLimit))
				} else {
					r.Stat("Page", fmt.Sprintf("%d of %d", response.Meta.Page, response.Meta.Pages))
					r.Stat("Limit", fmt.Sprintf("%d", response.Meta.Limit))
					if response.Meta.NextCursor != "" {
						r.Stat("Next cursor", response.Meta.NextCursor)
					}
				}
			}
			return ui.AdvisoryList(views)
		},
	}

	opts.register(cmd)
	return cmd
}

func Browse() *cobra.Command {
	opts := &queryOptions{}

	cmd := &cobra.Command{
		Use:   "browse [flags]",
		Short: i18n.C.AdvisoryBrowseShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			if err := opts.validate(); err != nil {
				return err
			}

			client := session.ConnectWithContext(cmd.Context(), config.Token())
			response, _, err := fetch(cmd.Context(), client, opts, r.Interactive())
			if err != nil {
				return err
			}

			var viewportOutput interface{} = response.GetData()
			if opts.full {
				viewportOutput = response
			}

			// browse is interactive; fall back to JSON whenever the TUI cannot
			// render, matching `vulncheck index browse`.
			if r.IsJSON() || !r.Interactive() {
				return r.JSON(viewportOutput)
			}

			ui.Viewport("advisory", viewportOutput)
			return nil
		},
	}

	opts.register(cmd)
	return cmd
}

// fetch runs the query and returns the whole response, so callers can emit
// either the records alone or the full envelope.
//
// An unknown feed name is not an error at the API -- /v4/advisory answers HTTP
// 200 with an empty page -- so a zero-result query with --feed set is checked
// against the catalogue, and retried if the user picks a correction.
func fetch(ctx context.Context, client *sdk.Client, opts *queryOptions, interactive bool) (*sdk.AdvisoryResponse, int, error) {
	params := opts.params()
	if opts.all {
		params.Page = 0
	}

	response, pages, err := runQuery(client, opts.all, params)
	if err != nil {
		return nil, 0, err
	}

	if len(response.GetData()) == 0 && opts.feed != "" {
		corrected, feedErr := checkFeed(ctx, client, opts.feed, interactive)
		if feedErr != nil {
			return nil, 0, feedErr
		}
		if corrected != "" && corrected != opts.feed {
			params.Name = corrected
			return runQuery(client, opts.all, params)
		}
	}

	return response, pages, nil
}

// runQuery issues one query, collapsing the paged and cursor-walked forms into
// the same response shape.
func runQuery(client *sdk.Client, all bool, params sdk.AdvisoryQueryParameters) (*sdk.AdvisoryResponse, int, error) {
	if all {
		records, pages, err := client.GetAllAdvisories(params)
		if err != nil {
			return nil, 0, err
		}
		// The last page's metadata would report page 1 of N with a filtered
		// count of that page alone, contradicting the data beside it. The
		// API's own total is an estimate -- it under-reports a completed walk
		// -- so the walked count is both self-consistent and more accurate.
		n := len(records)
		return &sdk.AdvisoryResponse{
			Data: records,
			Meta: sdk.AdvisoryMeta{Total: n, Page: 1, Pages: 1, Limit: n, Filtered: n},
		}, pages, nil
	}

	response, err := client.GetAdvisories(params)
	if err != nil {
		return nil, 0, err
	}
	if response == nil {
		return nil, 0, errs.New(errs.KindInternal, "empty response from /v4/advisory")
	}
	return response, 0, nil
}

// checkFeed validates a feed name once a query has already come back empty, so
// the common path costs no extra request.
//
// It returns a corrected name when the user picks one from the suggestions, so
// the caller can retry -- the same prompt-then-retry flow `backup` uses for
// index names, rather than telling the user to run the command again.
func checkFeed(ctx context.Context, client *sdk.Client, feed string, interactive bool) (string, error) {
	response, err := client.GetAdvisoryFeeds()
	if err != nil || response == nil {
		// The catalogue is only being used to improve an error message; if it
		// is unavailable, let the empty result stand.
		return "", nil
	}

	names := make([]string, 0, len(response.GetData()))
	for _, f := range response.GetData() {
		if f.Name == feed {
			return "", nil
		}
		names = append(names, f.Name)
	}

	suggestions := utils.SuggestFor(feed, names)
	if len(suggestions) == 0 {
		return "", errs.Validation("advisory feed '%s' does not exist", feed)
	}

	if !interactive {
		return "", errs.Validation("advisory feed '%s' does not exist; did you mean: %s",
			feed, strings.Join(suggestions, ", "))
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
