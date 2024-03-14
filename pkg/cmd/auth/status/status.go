package status

import "github.com/spf13/cobra"

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "status",
		Short: "Check authentication status",
		Long:  "Check if you're currently authenticated and if so, display the account information".
		RunE: func(cmd *cobra.Command, args []string) error {
			config.Init()
		},
	}

	return cmd
}
