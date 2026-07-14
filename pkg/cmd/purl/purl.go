package purl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/batch"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	var fromFile string

	cmd := &cobra.Command{
		Use:     "purl <scheme>",
		Short:   i18n.C.PurlShort,
		Example: fmt.Sprintf(i18n.C.PurlExample, "pkg:hackage/aeson@0.3.2.8"),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			// Resolve inputs from positional args, --from-file, or stdin
			// (stdin is only safe when not a TTY — we let CollectInputs
			// read it; a TTY would block, which the user can fix with
			// --no-interactive or a real input).
			inputs, err := batch.CollectInputs(args, fromFile, stdinIfPiped())
			if err != nil {
				return err
			}

			switch {
			case len(inputs) == 0:
				return ui.Error(i18n.C.PurlErrorSchemeRequired)
			case len(inputs) > 1 || fromFile != "":
				if !r.IsJSON() {
					return errs.Validation("batch purl lookups require --json output")
				}
				return batchPurl(cmd, inputs, r)
			default:
				return singlePurl(cmd, inputs[0], r)
			}
		},
	}

	cmd.Flags().StringVar(&fromFile, "from-file", "", "Read PURLs from a file, one per line (# comments and blank lines skipped)")

	return cmd
}

func singlePurl(cmd *cobra.Command, purl string, r *output.Renderer) error {
	response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetPurl(purl)
	if err != nil {
		return fmt.Errorf("error fetching purl %s: %w", purl, err)
	}

	if r.IsJSON() {
		return r.JSON(response.GetData())
	}

	vulns := response.Vulnerabilities()
	if err := ui.PurlMeta(response.PurlMeta()); err != nil {
		return err
	}
	if len(vulns) == 0 {
		r.Info(i18n.C.PurlNoVulns, purl)
		return nil
	}
	if len(vulns) == 1 {
		r.Info(i18n.C.PurlVulnFound, purl)
	} else {
		r.Info(i18n.C.PurlVulnsFound, len(vulns), purl)
	}
	return ui.PurlVulns(vulns)
}

func batchPurl(cmd *cobra.Command, inputs []string, r *output.Renderer) error {
	resp, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetPurls(inputs)
	if err != nil {
		return err
	}

	// Map response rows by the purl string so we can render results in the
	// caller's input order (which is the documented contract).
	byPurl := make(map[string]*sdk.BatchPurlData, len(resp.PurlData))
	for i := range resp.PurlData {
		byPurl[resp.PurlData[i].Purl] = &resp.PurlData[i]
	}

	out := make([]batch.Result, 0, len(inputs))
	for _, in := range inputs {
		if d, ok := byPurl[in]; ok {
			out = append(out, batch.ResultFromData(in, d))
			continue
		}
		out = append(out, batch.ResultFromError(in, fmt.Errorf("no result returned for this purl")))
	}
	return r.JSON(out)
}
