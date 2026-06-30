package ipintel

import (
	"fmt"
	"slices"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/cmd/offline/sync"
	"github.com/vulncheck-oss/cli/pkg/db"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func Command() *cobra.Command {
	var country, asn, cidr, countryCode, hostname, id string

	validTimeframes := []string{"3d", "10d", "30d"}

	cmd := &cobra.Command{
		Use:     "ipintel <3d|10d|30d> [flags]",
		Short:   "IP Intel offline search",
		Long:    "Search offline IP Intel data",
		Example: "vulncheck offline ipintel 3d --country=Sweden",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("requires a timeframe argument: the timeframe (3d, 10d, or 30d)\n\nUsage:\n  %s", cmd.UseLine())
			}
			if !slices.Contains(validTimeframes, args[0]) {
				return fmt.Errorf("invalid timeframe: %s. Must be one of %v\n\nUsage:\n  %s", args[0], validTimeframes, cmd.UseLine())
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)

			if country == "" && asn == "" && cidr == "" && countryCode == "" && hostname == "" && id == "" {
				return fmt.Errorf("requires at least one filter flag\n\nUsage:\n  %s", cmd.UseLine())
			}

			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			indexAvailable, err := sync.EnsureIndexSync(indices, fmt.Sprintf("ipintel-%s", args[0]), false)
			if err != nil {
				return err
			}

			if !indexAvailable {
				return fmt.Errorf("index %s is required to proceed", fmt.Sprintf("ipintel-%s", args[0]))
			}

			indices, err = cache.Indices()
			if err != nil {
				return err
			}

			index := indices.GetIndex(fmt.Sprintf("ipintel-%s", args[0]))
			if index == nil {
				return fmt.Errorf("index ipintel-%s is required for this command, and is not cached", args[0])
			}

			if !r.IsJSON() {
				r.Info("Searching index %s, last updated on %s", index.Name, utils.ParseDate(index.LastUpdated))
			}

			results, stats, err := db.IPIntelSearch(index.Name, country, asn, cidr, countryCode, hostname, id)
			if err != nil {
				return err
			}

			if r.IsJSON() {
				return r.JSON(results)
			}

			r.Stat("Results found", fmt.Sprintf("%d", len(results)))
			r.Stat("Search duration", stats.Duration.String())

			for i, result := range results {
				if i >= 10 {
					break
				}
				r.Info("%2d. IP: %-15s Country: %-15s ASN: %-10s ID: %s",
					i+1, result.IP, result.Country, result.ASN, result.Type.ID)
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&country, "country", "", "Filter by country")
	cmd.Flags().StringVar(&asn, "asn", "", "Filter by ASN")
	cmd.Flags().StringVar(&cidr, "cidr", "", "Filter by CIDR")
	cmd.Flags().StringVar(&countryCode, "country-code", "", "Filter by country code")
	cmd.Flags().StringVar(&hostname, "hostname", "", "Filter by hostname")
	cmd.Flags().StringVar(&id, "id", "", "Filter by ID")

	return cmd
}

func AliasCommands() []*cobra.Command {
	timeframes := []struct {
		short string
		long  string
	}{
		{"3d", "3 days"},
		{"10d", "10 days"},
		{"30d", "30 days"},
	}
	var commands []*cobra.Command

	for _, tf := range timeframes {
		shortTf := tf.short
		longTf := tf.long
		aliasCmd := &cobra.Command{
			Use:     fmt.Sprintf("ipintel-%s [flags]", shortTf),
			Short:   fmt.Sprintf("IP Intel offline search for %s timeframe", longTf),
			Long:    fmt.Sprintf("Search offline IP Intel data for %s timeframe", longTf),
			Example: fmt.Sprintf("vulncheck offline ipintel-%s --country=Sweden", shortTf),
			RunE: func(cmd *cobra.Command, args []string) error {
				mainCmd := Command()

				cmd.Flags().VisitAll(func(f *pflag.Flag) {
					if f.Changed {
						_ = mainCmd.Flags().Set(f.Name, f.Value.String())
					}
				})

				mainCmd.SetArgs(append([]string{shortTf}, args...))
				return mainCmd.Execute()
			},
		}

		mainCmd := Command()
		aliasCmd.Flags().AddFlagSet(mainCmd.Flags())

		commands = append(commands, aliasCmd)
	}

	return commands
}
