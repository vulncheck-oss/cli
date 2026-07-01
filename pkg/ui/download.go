package ui

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/vulncheck-oss/cli/pkg/sdk"
)

var p *tea.Program

type progressWriter struct {
	total      int
	downloaded int
	file       *os.File
	reader     io.Reader
	onProgress func(float64)
}

func (pw *progressWriter) Start() {
	// TeeReader calls pw.Write() each time a new response is received
	_, err := io.Copy(pw.file, io.TeeReader(pw.reader, pw))
	if err != nil {
		p.Send(progressErrMsg{err})
	}
}

func (pw *progressWriter) Write(p []byte) (int, error) {
	pw.downloaded += len(p)
	if pw.total > 0 && pw.onProgress != nil {
		pw.onProgress(float64(pw.downloaded) / float64(pw.total))
	}
	return len(p), nil
}

func getResponse(url string) (*http.Response, error) {
	if err := sdk.EnforceHTTPS(url); err != nil {
		return nil, err
	}
	resp, err := http.Get(url) // nolint:gosec
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("receiving status of %d for url: %s", resp.StatusCode, url)
	}
	return resp, nil
}

// DownloadHeadless streams url to filename using ctx for cancellation and
// emits coarse progress lines to progressOut (typically the renderer's
// stderr). It is the non-TUI counterpart to Download: no bubbletea, no TTY
// requirement, no panic on missing content length. Use when --json,
// --no-interactive, or a non-TTY stdout makes the bubbletea path unsafe.
func DownloadHeadless(ctx context.Context, url, filename string, progressOut io.Writer) error {
	if err := sdk.EnforceHTTPS(url); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}
	resp, err := http.DefaultClient.Do(req) // nolint:gosec
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint:errcheck
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP %d while downloading %s", resp.StatusCode, url)
	}

	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	tmp := filename + ".part"
	file, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	// Belt-and-braces: if anything errors below, remove the partial file
	// so a re-run starts clean.
	defer func() {
		_ = file.Close()
		if _, statErr := os.Stat(tmp); statErr == nil {
			_ = os.Remove(tmp)
		}
	}()

	// Periodic progress to stderr (machine-readable enough for scripts to
	// grep on, brief enough to not flood logs).
	total := resp.ContentLength
	buf := make([]byte, 64*1024)
	var written int64
	lastReport := time.Now()
	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		n, rerr := resp.Body.Read(buf)
		if n > 0 {
			if _, werr := file.Write(buf[:n]); werr != nil {
				return fmt.Errorf("write: %w", werr)
			}
			written += int64(n)
		}
		if progressOut != nil && time.Since(lastReport) > 500*time.Millisecond {
			if total > 0 {
				fmt.Fprintf(progressOut, "downloading %s: %d/%d bytes\n", filepath.Base(filename), written, total)
			} else {
				fmt.Fprintf(progressOut, "downloading %s: %d bytes\n", filepath.Base(filename), written)
			}
			lastReport = time.Now()
		}
		if rerr == io.EOF {
			break
		}
		if rerr != nil {
			return fmt.Errorf("read: %w", rerr)
		}
	}

	// Close before rename so the file is fully flushed.
	if err := file.Close(); err != nil {
		return fmt.Errorf("close: %w", err)
	}
	if err := os.Rename(tmp, filename); err != nil {
		return fmt.Errorf("finalise: %w", err)
	}
	return nil
}

func Download(url string, filename string) error {
	resp, err := getResponse(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // nolint:errcheck

	// Don't add TUI if the header doesn't include content size
	// it's impossible see progress without total
	if resp.ContentLength <= 0 {
		return fmt.Errorf("can't parse content length, aborting download")
	}

	// Create directory path if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return fmt.Errorf("could not create directory path: %w", err)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file: %w", err)
	}
	defer file.Close() // nolint:errcheck

	pw := &progressWriter{
		total:  int(resp.ContentLength),
		file:   file,
		reader: resp.Body,
		onProgress: func(ratio float64) {
			p.Send(progressMsg(ratio))
		},
	}

	m := downloadModel{
		pw:       pw,
		progress: progress.New(progress.WithScaledGradient("#6667AB", "#34D399")),
	}
	// Start Bubble Tea
	p = tea.NewProgram(m)

	// Start the download
	go pw.Start()

	if _, err := p.Run(); err != nil {
		return fmt.Errorf("error running program: %w", err)
	}

	return nil
}
