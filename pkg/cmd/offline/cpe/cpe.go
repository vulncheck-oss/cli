package cpe

import (
	"fmt"
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/cpeparse"
)

func Command() *cobra.Command {

	var jsonOutput bool

	cmd := &cobra.Command{
		Use:     "cpe <scheme>",
		Short:   "Offline CPE lookup",
		Long:    "Search offline package data via CPE schemes",
		Example: "vulncheck offline cpe \"pkg:hackage/aeson@0.3.2.8\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			attr, err := cpeparse.Parse(args[0])

			if err != nil {
				return err
			}

			fmt.Printf("Parsed CPE: %+v\n", attr)

			ray.Ray(attr.String())

			return nil
		},
	}
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON format")
	return cmd
}
