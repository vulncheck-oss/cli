package sdk

import (
	"fmt"
	stdurl "net/url"
)

// EnforceHTTPS refuses to follow a URL supplied by the upstream API
// unless it is https://. Guards against a compromised / MITM'd API
// tricking the CLI into downloading from an attacker-controlled host
// over cleartext, or being redirected to file:// / gopher:// / ftp://
// schemes that Go's default transport would happily fetch.
//
// This is the minimum viable defence: hostname pinning to a specific
// vulncheck-owned domain is intentionally NOT enforced here so backend
// operators can move backup storage between providers without a CLI
// release. If that policy tightens in the future, tighten this
// function; do not sprinkle host checks across call sites.
func EnforceHTTPS(rawURL string) error {
	u, err := stdurl.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid download URL: %w", err)
	}
	if u.Scheme != "https" {
		return fmt.Errorf("refusing to follow non-https download URL (scheme=%q); the VulnCheck API should only ever return https:// backup URLs — treat this as a MITM signal", u.Scheme)
	}
	if u.Host == "" {
		return fmt.Errorf("download URL is missing a host: %q", rawURL)
	}
	return nil
}
