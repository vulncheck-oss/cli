package index

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "index <command>",
		Short: i18n.C.IndexShort,
	}

	cmdList := &cobra.Command{
		Use:   "list <index>",
		Short: i18n.C.IndexListShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.ErrorIndexRequired)
			}
			response, err := session.Connect(config.Token()).GetIndex(args[0])
			if err != nil {
				return err
			}
			ui.Json(response.GetData())
			return nil
		},
	}

	cmdBrowse := &cobra.Command{
		Use:   "browse <index>",
		Short: i18n.C.IndexBrowseShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.ErrorIndexRequired)
			}
			response, err := session.Connect(config.Token()).GetIndex(args[0])
			if err != nil {
				return err
			}
			ui.Viewport(args[0], response.GetData())
			return nil
		},
	}

	cmd.AddCommand(cmdList)
	cmd.AddCommand(cmdBrowse)

	return cmd
}
