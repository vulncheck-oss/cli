package cpe

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/batch"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	var fromFile string

	cmd := &cobra.Command{
		Use:     "cpe <scheme>",
		Short:   i18n.C.CpeShort,
		Example: fmt.Sprintf(i18n.C.CpeExample, "cpe:2.3:a:sap:businessobjects_business_intelligence_platform:4.2:-:*"),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			inputs, err := batch.CollectInputs(args, fromFile, stdinIfPiped())
			if err != nil {
				return err
			}

			switch {
			case len(inputs) == 0:
				return ui.Error(i18n.C.CpeErrorSchemeRequired)
			case len(inputs) > 1 || fromFile != "":
				if !r.IsJSON() {
					return errs.Validation("batch cpe lookups require --json output")
				}
				return batchCPE(cmd, inputs, r)
			default:
				return singleCPE(cmd, inputs[0], r)
			}
		},
	}

	cmd.Flags().StringVar(&fromFile, "from-file", "", "Read CPEs from a file, one per line (# comments and blank lines skipped)")

	return cmd
}

func singleCPE(cmd *cobra.Command, cpe string, r *output.Renderer) error {
	response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetCpe(cpe)
	if err != nil {
		return err
	}

	cves := response.GetData()
	if r.IsJSON() {
		return r.JSON(cves)
	}

	if err := ui.CpeMeta(response.CpeMeta()); err != nil {
		return err
	}
	if len(cves) == 0 {
		r.Info(i18n.C.CpeNoCves, cpe)
		return nil
	}
	r.Info(i18n.C.CpeCvesFound, len(cves), cpe)
	return r.JSON(cves)
}

func batchCPE(cmd *cobra.Command, inputs []string, r *output.Renderer) error {
	// No server-side batch endpoint for cpe; loop sequentially. Use the
	// same client across iterations so connection reuse kicks in.
	client := session.ConnectWithContext(cmd.Context(), config.Token())
	out := make([]batch.Result, 0, len(inputs))
	for _, in := range inputs {
		resp, err := client.GetCpe(in)
		if err != nil {
			out = append(out, batch.ResultFromError(in, err))
			continue
		}
		out = append(out, batch.ResultFromData(in, resp.GetData()))
	}
	return r.JSON(out)
}

func stdinIfPiped() io.Reader {
	if output.IsTTY(os.Stdin) {
		return nil
	}
	return os.Stdin
}
