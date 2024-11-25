package cpe

import (
	"fmt"
	"github.com/octoper/go-ray"
	"github.com/spf13/cobra"
	"strings"
)
import "github.com/facebookincubator/nvdtools/wfn"

func Command() *cobra.Command {

	var jsonOutput bool

	cmd := &cobra.Command{
		Use:     "cpe <scheme>",
		Short:   "Offline CPE lookup",
		Long:    "Search offline package data via CPE schemes",
		Example: "vulncheck offline cpe \"pkg:hackage/aeson@0.3.2.8\"",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cpeString := args[0]
			var attr *wfn.Attributes
			var err error

			// Check if the CPE string is in 2.2 format
			if strings.HasPrefix(cpeString, "cpe:/") {
				// Convert 2.2 to 2.3 format
				attr, err = wfn.UnbindURI(cpeString)
			} else {
				// Assume 2.3 format
				attr, err = wfn.UnbindFmtString(cpeString)
			}

			if err != nil {
				return fmt.Errorf("invalid CPE string: %v", err)
			}

			// Print the parsed CPE attributes for debugging
			fmt.Printf("Parsed CPE: %+v\n", attr)

			ray.Ray(attr.String())

			return nil
		},
	}
	cmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "Output in JSON format")
	return cmd
}
