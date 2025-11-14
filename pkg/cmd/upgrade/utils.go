package upgrade

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func getLatestRelease() (*Release, error) {
	resp, err := httpClient.Get("https://api.github.com/repos/vulncheck-oss/cli/releases/latest")
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := in.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close source file: %v\n", closeErr)
		}
	}()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := out.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close destination file: %v\n", closeErr)
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return out.Sync()
}

func getSpecificRelease(version string) (*Release, error) {
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}

	url := fmt.Sprintf("https://api.github.com/repos/vulncheck-oss/cli/releases/tags/%s", version)
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("release %s not found", version)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
	}

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return nil, err
	}

	return &release, nil
}

func downloadAndInstall(downloadURL, filename, currentVersion string) error {
	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "vulncheck-upgrade-*")
	if err != nil {
		return err
	}
	defer func() {
		if removeErr := os.RemoveAll(tempDir); removeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to remove temp directory: %v\n", removeErr)
		}
	}()

	// Download file
	fmt.Printf("📥 Downloading %s...\n", filename)
	resp, err := httpClient.Get(downloadURL)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := resp.Body.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close response body: %v\n", closeErr)
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %d", resp.StatusCode)
	}

	tempFile := filepath.Join(tempDir, filename)
	out, err := os.Create(tempFile)
	if err != nil {
		return err
	}
	defer func() {
		if closeErr := out.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close temp file: %v\n", closeErr)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("📦 Extracting %s...\n", filename)
	binaryPath, err := extractArchive(tempFile, tempDir)
	if err != nil {
		return err
	}

	// Get current executable path
	currentExe, err := os.Executable()
	if err != nil {
		return err
	}

	// Get the real path (resolve symlinks)
	currentExe, err = filepath.EvalSymlinks(currentExe)
	if err != nil {
		return err
	}

	// Create backup of current binary
	backupFilename := fmt.Sprintf("vulncheck.backup.v%s.%s",
		currentVersion,
		time.Now().Format("20060102.150405"))
	backupPath := filepath.Join(filepath.Dir(currentExe), backupFilename)
	if err := copyFile(currentExe, backupPath); err != nil {
		return fmt.Errorf("failed to create backup: %v", err)
	}

	fmt.Printf("💾 Created backup at %s\n", backupPath)

	// Replace current binary
	fmt.Printf("🔧 Installing new binary...\n")

	if runtime.GOOS == "windows" {
		// On Windows, use the atomic rename approach since file locking is different
		tempBinary := currentExe + ".tmp"
		if err := copyFile(binaryPath, tempBinary); err != nil {
			return fmt.Errorf("failed to copy new binary to temp location: %v", err)
		}

		if err := os.Rename(tempBinary, currentExe); err != nil {
			os.Remove(tempBinary)
			if restoreErr := copyFile(backupPath, currentExe); restoreErr != nil {
				fmt.Fprintf(os.Stderr, "Critical: failed to restore backup: %v\n", restoreErr)
			}
			return fmt.Errorf("failed to install new binary: %v", err)
		}

		// Remove backup file after successful installation
		if removeErr := os.Remove(backupPath); removeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to remove backup file: %v\n", removeErr)
		}

		return nil
	} else {
		// On Unix systems, use a replacement script to avoid self-modification issues
		if err := createReplacementScript(binaryPath, currentExe, backupPath); err != nil {
			return fmt.Errorf("failed to create replacement script: %v", err)
		}

		// The script will handle the replacement and this process will exit
		fmt.Printf("✅ Upgrade process initiated. The binary will be replaced momentarily.\n")
		os.Exit(0)
	}

	return nil
}

func extractArchive(archivePath, destDir string) (string, error) {
	var binaryName string
	if runtime.GOOS == "windows" {
		binaryName = "vulncheck.exe"
	} else {
		binaryName = "vulncheck"
	}

	if strings.HasSuffix(archivePath, ".zip") {
		return extractZip(archivePath, destDir, binaryName)
	} else if strings.HasSuffix(archivePath, ".tar.gz") {
		return extractTarGz(archivePath, destDir, binaryName)
	}

	return "", fmt.Errorf("unsupported archive format")
}

func extractZip(zipPath, destDir, binaryName string) (string, error) {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := r.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close zip reader: %v\n", closeErr)
		}
	}()

	for _, f := range r.File {
		if filepath.Base(f.Name) == binaryName {
			rc, err := f.Open()
			if err != nil {
				return "", err
			}
			defer func() {
				if closeErr := rc.Close(); closeErr != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to close zip entry: %v\n", closeErr)
				}
			}()

			binaryPath := filepath.Join(destDir, binaryName)
			outFile, err := os.OpenFile(binaryPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return "", err
			}
			defer func() {
				if closeErr := outFile.Close(); closeErr != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to close output file: %v\n", closeErr)
				}
			}()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return "", err
			}

			return binaryPath, nil
		}
	}

	return "", fmt.Errorf("binary %s not found in zip archive", binaryName)
}

func extractTarGz(tarPath, destDir, binaryName string) (string, error) {
	file, err := os.Open(tarPath)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close tar file: %v\n", closeErr)
		}
	}()

	gzr, err := gzip.NewReader(file)
	if err != nil {
		return "", err
	}
	defer func() {
		if closeErr := gzr.Close(); closeErr != nil {
			fmt.Fprintf(os.Stderr, "Warning: failed to close gzip reader: %v\n", closeErr)
		}
	}()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		if filepath.Base(header.Name) == binaryName && header.Typeflag == tar.TypeReg {
			binaryPath := filepath.Join(destDir, binaryName)
			outFile, err := os.Create(binaryPath)
			if err != nil {
				return "", err
			}
			defer func() {
				if closeErr := outFile.Close(); closeErr != nil {
					fmt.Fprintf(os.Stderr, "Warning: failed to close output file: %v\n", closeErr)
				}
			}()

			_, err = io.Copy(outFile, tr)
			if err != nil {
				return "", err
			}

			return binaryPath, nil
		}
	}

	return "", fmt.Errorf("binary %s not found in tar.gz archive", binaryName)
}

func createReplacementScript(newBinaryPath, targetPath, backupPath string) error {
	scriptPath := targetPath + ".upgrade.sh"

	// Create a shell script that will replace the binary after this process exits
	script := fmt.Sprintf(`#!/bin/bash
set -e

# Wait for the parent process to exit
sleep 1

# Function to restore backup on failure
restore_backup() {
    if [ -f "%s" ]; then
        echo "🔄 Restoring backup..."
        cp "%s" "%s" || {
            echo "Critical: failed to restore backup" >&2
            exit 1
        }
        rm -f "%s"
    fi
}

# Set trap to restore backup on any error
trap restore_backup ERR

echo "🔧 Finalizing binary replacement..."

# Replace the binary
cp "%s" "%s"
chmod +x "%s"

# Remove backup on success
rm -f "%s"
rm -f "%s"

echo "✅ Successfully upgraded to new version!"
`, backupPath, backupPath, targetPath, backupPath,
		newBinaryPath, targetPath, targetPath,
		backupPath, scriptPath)

	if err := os.WriteFile(scriptPath, []byte(script), 0755); err != nil {
		return fmt.Errorf("failed to write replacement script: %v", err)
	}

	// Execute the script in the background
	if err := exec.Command("/bin/bash", scriptPath).Start(); err != nil {
		os.Remove(scriptPath)
		return fmt.Errorf("failed to start replacement script: %v", err)
	}

	return nil
}
