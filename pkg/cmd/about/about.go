package about

import (
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/session"
)

//go:embed about.txt
var ascii string

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:    "about",
		Hidden: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(ascii)
			fmt.Println(cmd.Root().Annotations["aboutInfo"])
		},
	}
	session.DisableAuthCheck(cmd)
	return cmd
}
