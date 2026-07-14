package upgrade

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"os"
	"path/filepath"
	"testing"
)

// TestIsArchivedBinary covers the disambiguation contract: the extract
// code must accept ONLY the entry at bin/<binary>, never a same-basename
// file elsewhere in the archive (bash completion script, man page, etc.).
// This was the bug that clobbered installed binaries during the initial
// 1.0.0 rollout — the release archive ships bin/vulncheck AND
// share/bash-completion/completions/vulncheck.
func TestIsArchivedBinary(t *testing.T) {
	cases := []struct {
		name  string
		entry string
		want  bool
	}{
		{"real binary at bin/", "vulncheck_1.0.0_macOS_arm64/bin/vulncheck", true},
		{"real binary with backslashes", "vulncheck_1.0.0_macOS_arm64\\bin\\vulncheck", true},
		{"bash completion — same basename, different dir", "vulncheck_1.0.0_macOS_arm64/share/bash-completion/completions/vulncheck", false},
		{"windows binary at bin/", "vulncheck_1.0.0_windows_amd64/bin/vulncheck.exe", false /* wrong binaryName */},
		{"LICENSE", "vulncheck_1.0.0_macOS_arm64/LICENSE", false},
		{"man page", "vulncheck_1.0.0_macOS_arm64/share/man/man1/vulncheck.1", false},
	}
	for _, c := range cases {
		got := isArchivedBinary(c.entry, "vulncheck")
		if got != c.want {
			t.Errorf("%s: isArchivedBinary(%q, %q) = %v, want %v", c.name, c.entry, "vulncheck", got, c.want)
		}
	}

	// Cross-check the windows binary against the windows binaryName.
	if !isArchivedBinary("vulncheck_1.0.0_windows_amd64/bin/vulncheck.exe", "vulncheck.exe") {
		t.Error("windows binary should match when binaryName is vulncheck.exe")
	}
}

// TestExtractZipPicksBinaryNotCompletion reproduces the 1.0.0 collision:
// build a zip whose completion script (same basename `vulncheck`) sits
// BEFORE the real binary, and confirm the extractor picks the binary.
func TestExtractZipPicksBinaryNotCompletion(t *testing.T) {
	dir := t.TempDir()
	zipPath := filepath.Join(dir, "release.zip")

	buf := &bytes.Buffer{}
	w := zip.NewWriter(buf)
	// Order matters — put the completion FIRST, mirroring what
	// goreleaser produced for the buggy 1.0.0 archive.
	mustAddZip(t, w, "vulncheck_1.0.0_macOS_arm64/share/bash-completion/completions/vulncheck", "#!/bin/bash\necho this is a completion script\n")
	mustAddZip(t, w, "vulncheck_1.0.0_macOS_arm64/LICENSE", "MIT")
	mustAddZip(t, w, "vulncheck_1.0.0_macOS_arm64/bin/vulncheck", "REAL_BINARY_BYTES")
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(zipPath, buf.Bytes(), 0o600); err != nil {
		t.Fatal(err)
	}

	out, err := extractZip(zipPath, dir, "vulncheck")
	if err != nil {
		t.Fatalf("extractZip: %v", err)
	}

	got, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "REAL_BINARY_BYTES" {
		t.Fatalf("wrong entry extracted: got %q; expected the bin/vulncheck payload", got)
	}
}

// TestExtractTarGzPicksBinaryNotCompletion mirrors the zip test for the
// linux tar.gz archive path.
func TestExtractTarGzPicksBinaryNotCompletion(t *testing.T) {
	dir := t.TempDir()
	tarPath := filepath.Join(dir, "release.tar.gz")

	buf := &bytes.Buffer{}
	gz := gzip.NewWriter(buf)
	tw := tar.NewWriter(gz)
	mustAddTar(t, tw, "vulncheck_1.0.0_linux_amd64/share/bash-completion/completions/vulncheck", "#!/bin/bash\ncompletion\n")
	mustAddTar(t, tw, "vulncheck_1.0.0_linux_amd64/LICENSE", "MIT")
	mustAddTar(t, tw, "vulncheck_1.0.0_linux_amd64/bin/vulncheck", "REAL_LINUX_BINARY")
	if err := tw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := gz.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(tarPath, buf.Bytes(), 0o600); err != nil {
		t.Fatal(err)
	}

	out, err := extractTarGz(tarPath, dir, "vulncheck")
	if err != nil {
		t.Fatalf("extractTarGz: %v", err)
	}

	got, err := os.ReadFile(out)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != "REAL_LINUX_BINARY" {
		t.Fatalf("wrong entry extracted: got %q; expected the bin/vulncheck payload", got)
	}
}

func mustAddZip(t *testing.T, w *zip.Writer, name, body string) {
	t.Helper()
	f, err := w.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.Write([]byte(body)); err != nil {
		t.Fatal(err)
	}
}

func mustAddTar(t *testing.T, w *tar.Writer, name, body string) {
	t.Helper()
	hdr := &tar.Header{
		Name:     name,
		Mode:    0o755,
		Size:    int64(len(body)),
		Typeflag: tar.TypeReg,
	}
	if err := w.WriteHeader(hdr); err != nil {
		t.Fatal(err)
	}
	if _, err := w.Write([]byte(body)); err != nil {
		t.Fatal(err)
	}
}
