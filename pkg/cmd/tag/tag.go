package tag

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

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "tag",
		Short:   i18n.C.TagShort,
		Example: fmt.Sprintf(i18n.C.TagExample, "vulncheck-c2"),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.TagErrorTagNameRequired)
			}
			response, err := session.ConnectWithContext(cmd.Context(), config.Token()).GetTag(args[0])
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(strings.Split(response, "\n"))
			}

			r.Println(response)
			return nil
		},
	}

	return cmd
}
