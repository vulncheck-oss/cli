package scan

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/sbom/scan/list"
	"github.com/vulncheck-oss/cli/pkg/i18n"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scan <command>",
		Short: i18n.C.SbomScanShort,
	}
	cmd.AddCommand(list.Command())
	return cmd
}
