# pkg/parity

Opt-in integration tests that diff the CLI's offline emulation of server-side
matching against the live VulnCheck API. Offline mode is meant to emulate the
API for air-gapped users, but nothing else in the tree asserts that `--offline`
and online produce the same CVE set for the same input — past drift between
them has only surfaced via customer reports. This suite closes that loop with a
small, seeded fixture set so regressions get caught at PR time.

## Running locally

```sh
# Auth (same env var the CLI itself reads via config.Token()).
export VC_TOKEN=...

# Sync every index the fixtures reference.
vulncheck offline sync --add cpecve
vulncheck offline sync --add alpine-purls
vulncheck offline sync --add ubuntu-purls
vulncheck offline sync --add debian-purls
vulncheck offline sync --add npm
vulncheck offline sync --add pypi
vulncheck offline sync --add cargo
vulncheck offline sync --add maven

go test -tags=parity ./pkg/parity/...
# or
make test-parity
```

Missing token or missing indices → `t.Skip`. The suite never fails a default
`go test ./...` run because the `parity` build tag hides it.

## What the tests do

- `TestPURLParity` — for each PURL in `testdata/purls.txt`, calls
  `bill.GetBatchVulns` (online, `/v3/purls` batch) and `bill.GetOfflineVulns`
  (offline sqlite), collapses each result to a CVE set, and asserts the sets
  are equal.
- `TestCPEParity` — for each CPE in `testdata/cpes.txt`, calls
  `session.Client.GetCpe` (online, `/v3/cpe`) and `bill.GetOfflineCpeVulns`
  (offline `cpecve` sqlite + client-side `cpeutils.Process`), and asserts the
  CVE sets are equal.

Failures print `only in online` / `only in offline` so the diff points at the
exact drift, not just "sets differ".

## Adding fixtures

Every fixture ships with a comment explaining why it's there — either a
regression seed tied to a shipped bug, or a coverage class (wildcards, version
collisions, negative case). Keep the sets small; this is a parity gate, not a
coverage test. When new drift is discovered, seed the case here so it can't
come back silently.

## What the tests deliberately do NOT do

- Do not compare vulnerability metadata (CVSS, KEV, description). Those live
  behind `MetaByCVE` / `GetIndexVulncheckNvd2` and are a CVE-id → row mapping
  — extremely unlikely to drift.
- Do not exercise the `scan` command end-to-end. That's a contract-test
  surface. Running the underlying lookups directly makes the failure signal
  precise.
- Do not validate flags that only exist on one side (e.g. `--include-cpes` is
  offline-only today; that asymmetry is a separate fix).
