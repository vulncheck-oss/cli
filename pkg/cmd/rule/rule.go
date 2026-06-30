package rule

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type Options struct {
	Table bool
}

func Command() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:     "rule <rule>",
		Short:   i18n.C.RuleShort,
		Example: fmt.Sprintf(i18n.C.RuleExample, "snort", "suricata"),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.RuleErrorRuleNameRequired)
			}

			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetRule(args[0])
			if err != nil {
				return err
			}

			rulesList := strings.Split(response, "\n")

			if r.IsJSON() {
				return r.JSON(rulesList)
			}

			if opts.Table {
				return ui.SingleColumnResults(rulesList, "Results")
			}

			r.Println(response)
			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Table, "table", "t", false, "Output as Table")

	return cmd
}
