package export

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/errs"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/db"
)

func Command() *cobra.Command {
	var outputFile string

	cmd := &cobra.Command{
		Use:   "export <index>",
		Short: "Export a synced offline index as JSONL",
		Long: "Dump every row of a locally synced offline index as newline-delimited JSON. " +
			"Writes to stdout by default; use -o to write a file. " +
			"Progress and stats go to stderr so stdout stays a clean JSONL stream.",
		Example: "vulncheck offline export vulncheck-nvd2 -o nvd2.jsonl\nvulncheck offline export npm | head",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			indexName := args[0]

			if !db.KnownOfflineIndex(indexName) {
				return errs.Validation("index %q is not a known offline schema", indexName)
			}

			out := cmd.OutOrStdout()
			closeOut := func() error { return nil }
			if outputFile != "" {
				file, err := os.Create(outputFile)
				if err != nil {
					return fmt.Errorf("unable to create %s: %w", outputFile, err)
				}
				out = file
				closeOut = file.Close
			}
			defer func() { _ = closeOut() }()

			stats, err := db.ExportIndexJSONL(indexName, out)
			if err != nil {
				return err
			}

			if outputFile != "" {
				r.Info("Exported %d rows from %s to %s in %s", stats.Rows, indexName, outputFile, stats.Duration)
			} else {
				// Keep stdout as pure JSONL; say so on the info stream.
				r.Info("Exported %d rows from %s in %s", stats.Rows, indexName, stats.Duration)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&outputFile, "output", "o", "", "Write JSONL to this file instead of stdout")

	return cmd
}
