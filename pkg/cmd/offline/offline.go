package offline

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/ipintel"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/purl"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offline <command>",
		Short: "Offline commands",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				panic(err)
			}
		},
	}

	cmd.AddCommand(sync.Command())
	cmd.AddCommand(ipintel.Command())
	cmd.AddCommand(ipintel.AliasCommands()...)
	cmd.AddCommand(purl.Command())

	return cmd
}
