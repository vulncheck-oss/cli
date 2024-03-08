package ascii

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
)

//go:embed ascii.txt
var ascii string

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "ascii",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(ascii)
		},
	}
	return cmd
}
