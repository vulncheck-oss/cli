package upgrade

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

func Command() *cobra.Command {
	var version string

	cmd := &cobra.Command{
		Use:   "upgrade",
		Short: i18n.C.UpgradeShort,
		Long:  i18n.C.UpgradeLong,
		RunE: func(cmd *cobra.Command, args []string) error {
			// If a specific version is provided, upgrade to that version
			if version != "" {
				return runUpgradeToVersion(version)
			}
			// Otherwise, show help since no default action
			return cmd.Help()
		},
	}

	cmd.AddCommand(StatusCommand())
	cmd.AddCommand(LatestCommand())

	cmd.Flags().StringVarP(&version, "version", "v", "", "Upgrade to a specific version (e.g., 1.0.0)")

	session.DisableAuthCheck(cmd)

	return cmd
}

func runUpgradeToVersion(targetVersion string) error {
	currentVersion := utils.TrimVersionString()

	targetVersion = strings.TrimPrefix(targetVersion, "v")

	// Fetch the specific release
	release, err := getSpecificRelease(targetVersion)
	if err != nil {
		return fmt.Errorf("failed to fetch release for version %s: %v", targetVersion, err)
	}

	fmt.Printf("🔄 Upgrading from %s to %s...\n", currentVersion, targetVersion)

	// Determine platform-specific assets
	assetName, err := utils.GetPlatformAssetName(targetVersion)
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

	fmt.Printf("✅ Successfully upgraded to version %s!\n", targetVersion)
	fmt.Printf("🔗 Changelog: https://github.com/vulncheck-oss/cli/releases/tag/v%s\n", targetVersion)

	return nil
}
