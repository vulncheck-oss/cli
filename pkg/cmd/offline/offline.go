package offline

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "offline <command>",
		Short: "Offline commands",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(sync.Command())

	return cmd
}
