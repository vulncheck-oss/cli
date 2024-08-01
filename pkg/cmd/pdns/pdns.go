package pdns

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
	Json bool
}

func Command() *cobra.Command {

	opts := &Options{
		Json: false,
	}

	cmd := &cobra.Command{
		Use:     "pdns",
		Short:   i18n.C.PdnsShort,
		Example: i18n.C.PdnsExample,
		RunE: func(cmd *cobra.Command, args []string) error {

			response, err := session.Connect(config.Token()).GetPdns()
			if err != nil {
				return err
			}

			pdnsList := strings.Split(response, "\n")

			if opts.Json {
				ui.Json(pdnsList)
				return nil
			}

			// default output
			fmt.Println(response)

			return nil
		},
	}

	cmd.Flags().BoolVarP(&opts.Json, "json", "j", false, "Output as JSON")

	return cmd
}
