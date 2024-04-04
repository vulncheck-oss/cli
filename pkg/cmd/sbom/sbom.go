package sbom

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/sbom/list"
	"github.com/vulncheck-oss/cli/pkg/cmd/sbom/scan"
	"github.com/vulncheck-oss/cli/pkg/i18n"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sbom <command>",
		Short: i18n.C.SbomShort,
	}
	cmd.AddCommand(list.Command())
	cmd.AddCommand(scan.Command())
	return cmd
}
