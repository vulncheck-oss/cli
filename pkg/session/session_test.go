package session

import (
	"testing"
)

func TestVersionFormat(t *testing.T) {
	expects := "vc version 1.4.0 (2020-12-15)\nhttps://github.com/vulncheck-oss/cli/releases/tag/v1.4.0\n"
	if got := VersionFormat("1.4.0", "2020-12-15"); got != expects {
		t.Errorf("Format() = %q, wants %q", got, expects)
	}
}

func TestChangelogURL(t *testing.T) {
	tag := "0.3.2"
	url := "https://github.com/vulncheck-oss/cli/releases/tag/v0.3.2"
	result := ChangelogURL(tag)
	if result != url {
		t.Errorf("expected %s to create url %s but got %s", tag, url, result)
	}

	tag = "v0.3.2"
	url = "https://github.com/vulncheck-oss/cli/releases/tag/v0.3.2"
	result = ChangelogURL(tag)
	if result != url {
		t.Errorf("expected %s to create url %s but got %s", tag, url, result)
	}

	tag = "0.3.2-pre.1"
	url = "https://github.com/vulncheck-oss/cli/releases/tag/v0.3.2-pre.1"
	result = ChangelogURL(tag)
	if result != url {
		t.Errorf("expected %s to create url %s but got %s", tag, url, result)
	}

	tag = "0.3.5-90-gdd3f0e0"
	url = "https://github.com/vulncheck-oss/cli/releases/latest"
	result = ChangelogURL(tag)
	if result != url {
		t.Errorf("expected %s to create url %s but got %s", tag, url, result)
	}

	tag = "deadbeef"
	url = "https://github.com/vulncheck-oss/cli/releases/latest"
	result = ChangelogURL(tag)
	if result != url {
		t.Errorf("expected %s to create url %s but got %s", tag, url, result)
	}
}
