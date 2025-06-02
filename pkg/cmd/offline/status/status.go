package status

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/ui"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: i18n.C.OfflineStatusShort,
		Long:  i18n.C.OfflineStatusLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			// Format and display the indices using the table UI
			if err := ui.CacheResults(indices.Indices); err != nil {
				return err
			}

			return nil
		},
	}
	return cmd
}
