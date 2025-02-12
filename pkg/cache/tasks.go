package cache

import (
	"fmt"
	"github.com/fumeapp/taskin"
	"github.com/vulncheck-oss/cli/pkg/db"
	"github.com/vulncheck-oss/cli/pkg/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// DownloadTask creates a task for downloading a file from the given URL.
func taskDownload(url string, index string, filename string) taskin.Task {
	return taskin.Task{
		Title: fmt.Sprintf("Download %s", index),
		Task: func(t *taskin.Task) error {
			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("failed to get URL: %w", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				return fmt.Errorf("received non-200 status code: %d", resp.StatusCode)
			}

			file, err := os.Create(filename)
			if err != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}
			defer file.Close()

			size := resp.ContentLength
			if size <= 0 {
				return fmt.Errorf("unknown content length")
			}

			written := int64(0)
			buffer := make([]byte, 32*1024)
			for {
				nr, er := resp.Body.Read(buffer)
				if nr > 0 {
					nw, ew := file.Write(buffer[0:nr])
					if nw > 0 {
						written += int64(nw)
					}
					if ew != nil {
						err = ew
						break
					}
					if nr != nw {
						err = io.ErrShortWrite
						break
					}
				}
				if er != nil {
					if er != io.EOF {
						err = er
					}
					break
				}

				progress := float64(written) / float64(size)
				t.Progress(int(progress*100), 100)
				t.Title = fmt.Sprintf("Downloading %s %.2f%%", index, progress*100)
			}

			if err != nil {
				return fmt.Errorf("error during download: %w", err)
			}

			return nil
		},
	}
}

func taskDB(index string, configDir string, filePath string, lastUpdated string, indexInfo *InfoFile) taskin.Task {
	return taskin.Task{
		Title: fmt.Sprintf("Index %s", index),
		Task: func(t *taskin.Task) error {
			lastProgress := -1
			eta := utils.NewETACalculator()

			progressCallback := func(progress int) {
				if progress != lastProgress {
					_, duration := eta.Update(progress)
					etaStr := utils.FormatETA(duration)

					t.Progress(progress, 100)
					t.Title = fmt.Sprintf("Indexing %s %d%% (ETA: %s)",
						index, progress, etaStr)
					lastProgress = progress
				}
			}

			indexDir := filepath.Join(configDir, index)
			if err := db.ImportIndex(filePath, indexDir, progressCallback); err != nil {
				return err
			}

			size, err := utils.GetDirectorySize(indexDir)
			if err != nil {
				return fmt.Errorf("failed to calculate size of index directory: %s", err)
			}

			updatedInfo := IndexInfo{
				Name:        index,
				LastSync:    time.Now(),
				Size:        size,
				LastUpdated: lastUpdated,
			}

			found := false
			for i, info := range indexInfo.Indices {
				if info.Name == index {
					indexInfo.Indices[i] = updatedInfo
					found = true
					break
				}
			}
			if !found {
				indexInfo.Indices = append(indexInfo.Indices, updatedInfo)
			}

			totalTime := eta.TotalTime()
			totalTimeStr := utils.FormatETA(totalTime)
			t.Title = fmt.Sprintf("Indexed %s (Size: %s, Time: %s)", index, utils.GetSizeHuman(size), totalTimeStr)
			return nil
		},
	}
}

// extractIndexTask creates a task for extracting the index file.
func taskExtract(index string, configDir string, filePath string) taskin.Task {
	return taskin.Task{
		Title: fmt.Sprintf("Extract %s", index),
		Task: func(t *taskin.Task) error {
			t.Title = fmt.Sprintf("Extracting %s", index)
			indexDir := filepath.Join(configDir, index)
			if err := os.MkdirAll(indexDir, 0755); err != nil {
				return fmt.Errorf("failed to create index directory: %w", err)
			}
			if err := utils.Unzip(filePath, indexDir); err != nil {
				return fmt.Errorf("failed to unzip index file: %w", err)
			}

			size, err := utils.GetDirectorySize(indexDir)
			if err != nil {
				return fmt.Errorf("failed to calculate size of index directory: %w", err)
			}

			t.Title = fmt.Sprintf("Extracted %s (Size: %s)", index, utils.GetSizeHuman(size))

			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("failed to remove downloaded zip file: %s", err)
			}

			return nil
		},
	}
}
