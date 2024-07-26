package rule

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

type Options struct {
	Json  bool
	Table bool
}

func Command() *cobra.Command {

	opts := &Options{
		Json:  false,
		Table: false,
	}

	cmd := &cobra.Command{
		Use:     "rule <rule>",
		Short:   i18n.C.RuleShort,
		Example: fmt.Sprintf(i18n.C.RuleExample, "snort", "suricata"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.RuleErrorRuleNameRequired)
			}

			response, err := session.Connect(config.Token()).GetRule(args[0])
			if err != nil {
				return err
			}

			rulesList := strings.Split(response, "\n")

			if opts.Json {
				ui.Json(rulesList)
				return nil
			}

			if opts.Table {
				if err := ui.RuleResults(rulesList); err != nil {
					return err
				}

				return nil
			}

			// default output
			fmt.Println(response)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")
	cmd.Flags().BoolVarP(&opts.Table, "table", "t", false, "Output as Table")

	return cmd
}
