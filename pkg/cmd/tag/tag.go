package tag

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
		Use:     "tag",
		Short:   i18n.C.TagShort,
		Example: fmt.Sprintf(i18n.C.TagExample, "vulncheck-c2"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.TagErrorTagNameRequired)
			}
			response, err := session.Connect(config.Token()).GetTag(args[0])
			if err != nil {
				return err
			}

			tagList := strings.Split(response, "\n")

			if opts.Json {
				ui.Json(tagList)
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
