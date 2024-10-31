package ipintel

import (
	"fmt"
	"github.com/package-url/packageurl-go"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/vulncheck-oss/cli/pkg/cache"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/search"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"slices"
	"strings"
)

func Command() *cobra.Command {
	var country, asn, cidr, countryCode, hostname, id string

	var jsonOutput bool

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

			// we need to error if zero flags were specified
			if country == "" && asn == "" && cidr == "" && countryCode == "" && hostname == "" && id == "" {
				return fmt.Errorf("requires at least one filter flag\n\nUsage:\n  %s", cmd.UseLine())
			}

			indices, err := cache.Indices()
			if err != nil {
				return err
			}

			index := indices.GetIndex(fmt.Sprintf("ipintel-%s", args[0]))

			if index == nil {
				return fmt.Errorf("index ipintel-%s is required for this command, and is not cached", args[0])
			}

			query := buildQuery(country, asn, cidr, countryCode, hostname, id)

			if !jsonOutput && !config.IsCI() {
				ui.Info(fmt.Sprintf("Searching index %s, last updated on %s", index.Name, utils.ParseDate(index.LastUpdated)))
			}

			results, stats, err := search.Index(index.Name, query)
			if err != nil {
				return err
			}

			if jsonOutput || config.IsCI() {
				ui.Json(results)
				return nil
			}

			ui.Stat("Results found", fmt.Sprintf("%d", len(results)))
			ui.Stat("Files/Lines processed", fmt.Sprintf("%d/%d", stats.TotalFiles, stats.TotalLines))
			ui.Stat("Search duration", stats.Duration.String())

			for i, result := range results {
				if i >= 10 {
					break
				}
				ui.Info(fmt.Sprintf("%2d. IP: %-15s Country: %-15s ASN: %-10s ID: %s",
					i+1, result.IP, result.Country, result.ASN, result.Type.ID))
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
	cmd.Flags().BoolVar(&jsonOutput, "json", false, "Output results in JSON format")

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

				// Transfer all flags from the alias command to the main command
				cmd.Flags().VisitAll(func(f *pflag.Flag) {
					if f.Changed {
						_ = mainCmd.Flags().Set(f.Name, f.Value.String())
					}
				})

				// Set the timeframe argument
				mainCmd.SetArgs(append([]string{shortTf}, args...))
				return mainCmd.Execute()
			},
		}

		// Copy flags from the main command
		mainCmd := Command()
		aliasCmd.Flags().AddFlagSet(mainCmd.Flags())

		commands = append(commands, aliasCmd)
	}

	return commands
}

func buildQuery(country, asn, cidr, countryCode, hostname, id string) string {
	var conditions []string

	if country != "" {
		conditions = append(conditions, fmt.Sprintf(".country == %q", country))
	}
	if asn != "" {
		conditions = append(conditions, fmt.Sprintf(".asn == %q", asn))
	}
	if cidr != "" {
		// Note: CIDR matching would require additional logic
		conditions = append(conditions, fmt.Sprintf(".ip == %q", cidr))
		// conditions = append(conditions, fmt.Sprintf(".ip | startswith(%q)", strings.Split(cidr, "/")[0]))
	}
	if countryCode != "" {
		conditions = append(conditions, fmt.Sprintf(".country_code == %q", countryCode))
	}
	if hostname != "" {
		conditions = append(conditions, fmt.Sprintf(".hostnames | any(. == %q)", hostname))
	}
	if id != "" {
		conditions = append(conditions, fmt.Sprintf(".type.id == %q", id))
	}

	if len(conditions) == 0 {
		return "true"
	}
	return strings.Join(conditions, " and ")
}

func BuildPurlQuery(instance packageurl.PackageURL) string {
	seperator := "/"
	var conditions []string

	if instance.Type == "maven" {
		seperator = ":"
	}

	if instance.Namespace == "alpine" {
		conditions = append(conditions, fmt.Sprintf(".package_name == %q", instance.Name))
	} else {
		if instance.Namespace != "" {
			conditions = append(conditions, fmt.Sprintf(".name == %q", fmt.Sprintf("%s%s%s", instance.Namespace, seperator, instance.Name)))
		} else {
			conditions = append(conditions, fmt.Sprintf(".name == %q", instance.Name))
		}
	}

	if instance.Version != "" {
		conditions = append(conditions, fmt.Sprintf(".version == %q", instance.Version))
	}

	// purlStr := instance.ToString()
	// conditions = append(conditions, fmt.Sprintf(".purl | any(. == %q)", purlStr))

	return strings.Join(conditions, " and ")
}
