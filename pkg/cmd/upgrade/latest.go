package upgrade

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func LatestCommand() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "latest",
		Short: i18n.C.UpgradeLatestShort,
		Long:  i18n.C.UpgradeLatestLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUpgradeLatest(force)
		},
	}

	cmd.Flags().BoolVarP(&force, "force", "f", false, "Force upgrade even if already on the latest version")

	session.DisableAuthCheck(cmd)
	return cmd
}

func runUpgradeLatest(force bool) error {
	currentVersion := utils.TrimVersionString()

	release, err := getLatestRelease()
	if err != nil {
		return fmt.Errorf("failed to fetch latest release: %v", err)
	}

	latestVersion := strings.TrimPrefix(release.TagName, "v")

	// Check if we need to upgrade (unless forced)
	if !force {
		// Handle development version
		if currentVersion == "DEV" || currentVersion == "" {
			fmt.Printf("🔄 Upgrading from development build to %s...\n", latestVersion)
		} else {
			currentV, err := version.NewVersion(currentVersion)
			if err != nil {
				return fmt.Errorf("invalid current version: %v", err)
			}

			latestV, err := version.NewVersion(latestVersion)
			if err != nil {
				return fmt.Errorf("invalid latest version: %v", err)
			}

			if !latestV.GreaterThan(currentV) {
				fmt.Printf("✅ Already on the latest version (%s)\n", currentVersion)
				return nil
			}

			fmt.Printf("🔄 Upgrading from %s to %s...\n", currentVersion, latestVersion)
		}
	} else {
		if currentVersion == "DEV" || currentVersion == "" {
			fmt.Printf("🔄 Force upgrading from development build to %s...\n", latestVersion)
		} else {
			fmt.Printf("🔄 Force upgrading from %s to %s...\n", currentVersion, latestVersion)
		}
	}

	assetName, err := utils.GetPlatformAssetName(latestVersion)
	if err != nil {
		return fmt.Errorf("failed to determine platform asset: %v", err)
	}

	var downloadURL string
	for _, asset := range release.Assets {
		if asset.Name == assetName {
			downloadURL = asset.BrowserDownloadURL
			break
		}
	}

	if downloadURL == "" {
		return fmt.Errorf("asset %s not found in release %s", assetName, release.TagName)
	}

	if err := downloadAndInstall(downloadURL, assetName, currentVersion); err != nil {
		return fmt.Errorf("failed to download and install: %v", err)
	}

	fmt.Printf("✅ Successfully upgraded to version %s!\n", latestVersion)
	fmt.Printf("🔗 Changelog: https://github.com/vulncheck-oss/cli/releases/tag/v%s\n", latestVersion)

	return nil
}
