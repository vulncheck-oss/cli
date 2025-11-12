package upgrade

import (
	"fmt"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/fumeapp/taskin"
	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/build"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func StatusCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "status",
		Short: i18n.C.UpgradeStatusShort,
		Long:  i18n.C.UpgradeStatusLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatus()
		},
	}

	session.DisableAuthCheck(cmd)
	return cmd
}

func runStatus() error {
	currentVersion := utils.TrimVersionString()

	var release *Release
	var latestVersion string
	var assetName string
	var assetExists bool
	var isDevelopmentBuild bool

	// Check if this is a development build
	if currentVersion == "DEV" || currentVersion == "" {
		isDevelopmentBuild = true
	}

	tasks := taskin.Tasks{}

	// Task 1: Display current version
	tasks = append(tasks, taskin.Task{
		Title: "Checking current version",

		Task: func(t *taskin.Task) error {
			if isDevelopmentBuild {
				t.Title = fmt.Sprintf("Current version: %s (development build)", build.Version)
			} else {
				t.Title = fmt.Sprintf("Current version: %s", currentVersion)
			}
			return nil
		},
	})

	// Task 2: Fetch latest release from GitHub
	tasks = append(tasks, taskin.Task{
		Title: "Fetching latest release from GitHub",
		Task: func(t *taskin.Task) error {
			var err error
			release, err = getLatestRelease()
			if err != nil {
				return fmt.Errorf("failed to fetch latest release: %v", err)
			}
			latestVersion = utils.TrimVersionString(release.TagName)
			t.Title = fmt.Sprintf("Latest version available: %s", latestVersion)
			return nil
		},
	})

	// Task 3: Check platform compatibility
	tasks = append(tasks, taskin.Task{
		Title: "Checking platform compatibility",
		Task: func(t *taskin.Task) error {
			var err error
			assetName, err = utils.GetPlatformAssetName(latestVersion)
			if err != nil {
				return fmt.Errorf("failed to determine platform asset: %v", err)
			}

			assetExists = false
			for _, asset := range release.Assets {
				if asset.Name == assetName {
					assetExists = true
					break
				}
			}

			if !assetExists {
				t.Title = fmt.Sprintf("Platform-specific release (%s) not available", assetName)
				return nil
			}

			t.Title = fmt.Sprintf("Platform asset found: %s", assetName)
			return nil
		},
	})

	// Run all tasks
	runners := taskin.New(tasks, taskin.Config{
		DisableUI: false,
		ProgressOptions: []progress.Option{
			progress.WithScaledGradient("#6667AB", "#34D399"),
			progress.WithWidth(20),
			progress.WithoutPercentage(),
		},
	})

	if err := runners.Run(); err != nil {
		return err
	}

	// Display final results after tasks complete
	fmt.Println() // Add some spacing

	if !assetExists {
		fmt.Printf("⚠️  Platform-specific release (%s) not available for version %s\n", assetName, latestVersion)
		return nil
	}

	if isDevelopmentBuild {
		fmt.Printf("💡 Latest stable release is %s\n", latestVersion)
		fmt.Printf("🔗 Release notes: https://github.com/vulncheck-oss/cli/releases/tag/v%s\n", latestVersion)
		fmt.Println("💡 Run 'vulncheck upgrade latest' to install the latest stable version")
		return nil
	}

	// Compare versions for release builds
	currentV, err := version.NewVersion(currentVersion)
	if err != nil {
		return fmt.Errorf("invalid current version: %v", err)
	}

	latestV, err := version.NewVersion(latestVersion)
	if err != nil {
		return fmt.Errorf("invalid latest version: %v", err)
	}

	if latestV.GreaterThan(currentV) {
		fmt.Printf("⬆️  Update available! You can upgrade from %s to %s\n", currentVersion, latestVersion)
		fmt.Printf("🔗 Release notes: https://github.com/vulncheck-oss/cli/releases/tag/v%s\n", latestVersion)
		fmt.Println("💡 Run 'vulncheck upgrade latest' to update to the latest version")
	} else if latestV.Equal(currentV) {
		fmt.Printf("✅ You are running the latest version (%s)\n", currentVersion)
	} else {
		fmt.Printf("🚧 You are running a newer version (%s) than the latest release (%s)\n", currentVersion, latestVersion)
	}

	return nil
}
