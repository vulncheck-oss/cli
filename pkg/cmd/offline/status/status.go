package status

import (
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
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
			r := output.FromCmd(cmd)
			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(indices.Indices)
			}

			return ui.CacheResults(indices.Indices)
		},
	}
	return cmd
}
