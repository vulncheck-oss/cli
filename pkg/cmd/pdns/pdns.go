package pdns

import (
	"fmt"
	"io"
	"os"
	"strings"

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
		Use:     "pdns <list>",
		Short:   i18n.C.PdnsShort,
		Example: fmt.Sprintf(i18n.C.PdnsExample, "vulncheck-c2"),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			inputs, err := batch.CollectInputs(args, fromFile, stdinIfPiped())
			if err != nil {
				return err
			}
			switch {
			case len(inputs) == 0:
				return ui.Error(i18n.C.PdnsErrorListNameRequired)
			case len(inputs) > 1 || fromFile != "":
				if !r.IsJSON() {
					return errs.Validation("batch pdns lookups require --json output")
				}
				client := session.ConnectWithContext(cmd.Context(), config.Token())
				out := make([]batch.Result, 0, len(inputs))
				for _, in := range inputs {
					resp, err := client.GetPdns(in)
					if err != nil {
						out = append(out, batch.ResultFromError(in, err))
						continue
					}
					out = append(out, batch.ResultFromData(in, strings.Split(resp, "\n")))
				}
				return r.JSON(out)
			default:
				resp, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetPdns(inputs[0])
				if err != nil {
					return err
				}
				if r.IsJSON() {
					return r.JSON(strings.Split(resp, "\n"))
				}
				r.Println(resp)
				return nil
			}
		},
	}

	cmd.Flags().StringVar(&fromFile, "from-file", "", "Read list names from a file, one per line")

	return cmd
}

func stdinIfPiped() io.Reader {
	if output.IsTTY(os.Stdin) {
		return nil
	}
	return os.Stdin
}
