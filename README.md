<p align="center">
    <img src="/logo-cli.png" align="center" alt="VulnCheck Logo" width="150" />
</p>

# The VulnCheck CLI
`vulncheck` is access to the VulnCheck API on the command line. It brings index browsing, backup management, and vulnerability scanning to the terminal.

<p align="center">
    <img src="/vulncheck-scan.gif" />
</p>

[![Release](https://img.shields.io/github/v/release/vulncheck-oss/cli)](https://github.com/vulncheck-oss/cli/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/vulncheck-oss/cli.svg)](https://pkg.go.dev/github.com/vulncheck-oss/cli)
[![Lint](https://github.com/vulncheck-oss/cli/actions/workflows/lint.yml/badge.svg)](https://github.com/vulncheck-oss/cli/actions/workflows/lint.yml)
[![Tests](https://github.com/vulncheck-oss/cli/actions/workflows/test.yml/badge.svg)](https://github.com/vulncheck-oss/cli/actions/workflows/test.yml)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/vulncheck-oss/cli/pulls)

## Installation 

### Provided install scripts

You can easily install vulncheck using an install script. Choose the script and method that matches your operating system:

### macOS and Linux

Open a terminal and run:

```bash
curl -sSL https://raw.githubusercontent.com/vulncheck-oss/cli/main/install.sh | bash
```

This will prompt you to choose between system-wide installation (requires sudo) or local user installation.

> [!NOTE]
> The install script also supports non-interactive installation options:
> - `--sudo` for system-wide installation without prompts
> - `--non-sudo` for local user installation without prompts
> - `--help` or `-h` to see all available options
> ```bash
> curl -sSL https://raw.githubusercontent.com/vulncheck-oss/cli/main/install.sh | bash -s -- --help
> ```

### Windows

Open PowerShell and run:

```
iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/vulncheck-oss/cli/main/install.ps1'))
```

`vulncheck` binaries are also available for MacOS, Linux, and Windows. You can download precompiled binaries from our [releases page](https://github.com/vulncheck-oss/cli/releases/latest)

### Verify the installation

Once installed, confirm the binary prints the desired version:

```bash
vulncheck version
```

You should see the version, build date, and changelog URL. If you get "command not found", reopen your shell so the new `PATH` is picked up, then try again.

## Configuration
* Run `vulncheck auth login` to authenticate with your VulnCheck account.
* Alternatively `vulncheck` will respect the `VC_TOKEN` environment variable.
* `vulncheck auth` by itself will show other options like checking your status and logging out.

`VC_TOKEN` wins over the saved config file. Because of that, `auth login` and
`auth logout` refuse while it is set — otherwise they would report success
having changed nothing that takes effect. Run `vulncheck auth status` to see
which of the two sources the active token came from.


## Agentic / scripted usage

The CLI is designed to be safe to drive from scripts and AI agents. This section is the contract — the surfaces below are intended to remain stable across releases (additions are not breaking changes; renames / removals are).

### Global flags

| Flag                | Effect |
|---------------------|--------|
| `--json`            | Emit JSON on stdout; route info/progress lines to stderr; errors emitted as a structured envelope. |
| `--quiet`           | Suppress informational output. Errors and payloads still render. |
| `--no-color`        | Disable ANSI styling. Also honours the `NO_COLOR` env var. |
| `--no-interactive` | Refuse to block on TUI prompts; commands that need a prompt return an error instead. Implied by `--json`, non-TTY stdin/stdout, and any of the `CI` / `BUILD_NUMBER` / `RUN_ID` env vars. |

### Environment variables

| Variable                                   | Effect |
|--------------------------------------------|--------|
| `VC_TOKEN`                                 | API token. Takes precedence over `~/.config/vulncheck/vulncheck.yaml`. While it is set, `auth login` and `auth logout` refuse rather than writing a config file that would be ignored — `unset VC_TOKEN` first. |
| `NO_COLOR`                                 | Any non-empty value disables ANSI styling. |
| `CI` / `BUILD_NUMBER` / `RUN_ID`           | Any of these set implies non-interactive mode (no prompts). |

### Exit codes

| Code | Meaning |
|------|---------|
| 0    | Success. |
| 1    | Generic / internal error. |
| 2    | Validation failure (bad args, missing required flag, malformed request). |
| 3    | Auth failure (no token, or the server rejected the token). |
| 4    | Resource not found (HTTP 404, no such index). |
| 5    | Rate limited (HTTP 429). |
| 6    | Network failure (DNS, connection refused, timeout). |
| 130  | Cancelled by SIGINT (POSIX `128 + 2`). |

### Error envelope

In `--json` mode, errors are emitted to **stdout** as:

```json
{
  "schema_version": 1,
  "error": {
    "code": "auth_required",
    "message": "...",
    "http_status": 401,
    "hint": "..."
  }
}
```

`code` is one of: `internal`, `validation`, `auth_required`, `auth_invalid`, `not_found`, `rate_limited`, `network`, `bad_request`, `cancelled`. `http_status` is omitted for non-HTTP errors. `hint` is optional remediation context — present only when the message alone isn't actionable (e.g. naming `VC_TOKEN` as the source of a rejected token) — and is rendered on stderr as `hint: ...` outside `--json` mode.

### Probe commands

Use these to inspect the CLI itself before dispatching work:

```bash
vulncheck version --json
# {"schema_version": 1, "version": "...", "build_date": "...", "changelog_url": "..."}

vulncheck auth status --json
# {"schema_version": 1, "authenticated": true, "token_source": "env", "user": "...", "email": "..."}
# Exit 0 even when authenticated=false — agents dispatch on the bool.
# token_shadowed: true is added when VC_TOKEN is overriding a *different*
# token saved in vulncheck.yaml — the usual cause of "I logged in but
# nothing changed". Omitted otherwise, so the CI shape (VC_TOKEN only,
# no config file) never reports shadowing.

vulncheck commands
# {"schema_version": 1, "root": {"name":"vulncheck", "subcommands":[...]}, ...}
# Machine-readable dump of the whole command tree — every subcommand,
# every flag (with type + default + usage), aliases, deprecation. Use
# this instead of parsing --help. Auth is not required.
```

### Pagination

```bash
vulncheck token list --json --limit 10 --page 2
vulncheck token list --json --all       # auto-paginate, single combined array
vulncheck index list <index> --json --all
```

### Batch input

`purl`, `cpe`, `tag`, `pdns` accept multiple inputs via positional args, stdin (when piped), or `--from-file <path>`. Blank lines and `#`-prefixed comments in the file are ignored. Batch mode requires `--json`.

```bash
# Stdin
cat purls.txt | vulncheck purl --json

# File
vulncheck cpe --from-file ./cpes.txt --json
```

The batch envelope is a stable array, one row per input, in input order:

```json
[
  {"input": "pkg:npm/lodash@4.0.0", "data": { ... }},
  {"input": "pkg:bad/string", "error": "no result returned for this purl"}
]
```

### Cancellation

`SIGINT` / `SIGTERM` cancel in-flight HTTP requests cleanly via context propagation. Long-running ops (`scan`, `offline sync`, `backup download`) honour cancellation; partial files are removed on the way out where applicable.

### JSON output discipline

When `--json` is set:
- stdout carries **only** the JSON payload (or the error envelope above).
- stderr carries every info line, progress bar, spinner, prompt, and warning.
- TUI elements (bubbletea, huh prompts) are automatically suppressed.

This means `vulncheck <cmd> --json | jq` always works — no `tail`/`sed` cleanup needed.


## Available commands

Every command below accepts the [global flags](#global-flags) (`--json`, `--quiet`, `--no-color`, `--no-interactive`, `--help`/`-h`). Per-command flag tables list only what is specific to that command.

- [`auth`](#auth) — log in / out, check status
- [`token`](#token) — API token management
- [`indices`](#indices) — list or browse the catalogue of indices
- [`index`](#index) — query one index
- [`backup`](#backup) — download or fetch a signed URL for an index backup
- [`cpe`](#cpe) — look up CVEs for a CPE (single or batch)
- [`purl`](#purl) — look up CVEs for a PURL (single or batch)
- [`tag`](#tag) — look up IP-intelligence tag membership
- [`pdns`](#pdns) — look up passive-DNS list membership
- [`rule`](#rule) — look up initial-access-intelligence rules
- [`scan`](#scan) — scan a directory (SBOM + vulnerability lookup)
- [`offline`](#offline) — sync indices locally and query them without hitting the API
- [`version`](#version) — print the CLI version
- [`upgrade`](#upgrade) — update the CLI in place


### auth

```
vulncheck auth login
vulncheck auth logout
vulncheck auth status [--json]
```

`login` walks you through a browser or paste-token flow — refuses when `--no-interactive`. `status --json` calls `/me` to actually verify the token; the payload always has `.authenticated: bool` and exits `0` regardless (see [Probe commands](#probe-commands)).


### token

```
vulncheck token list [--limit N] [--page N] [--all]
vulncheck token create <label>
vulncheck token remove <id>
vulncheck token browse
```

`list --json --all` auto-paginates and returns one combined JSON array.

`create --json` returns `{schema_version, id, label, token_on_stderr: true}` and prints the actual secret on a single line to **stderr**. That way a pipeline like `vulncheck token create ci-runner --json > token.json` never captures the secret in the JSON file. Pass `--allow-token-on-stdout` if you want the token embedded in the JSON payload instead (`{... "token": "vc_..."}`) — you're taking responsibility for the redirection.

`browse` is interactive; it refuses with exit `2` under `--no-interactive`.


### indices

```
vulncheck indices list [<search>]
vulncheck indices browse [<search>]
```

Lists (or interactively browses) the catalogue of available indices. `list` accepts a fuzzy search term.


### index

```
vulncheck index list <index> [--full] [--all] [query flags]
vulncheck index browse <index> [query flags]
```

`list --all` walks `next_cursor` end-to-end and emits one combined array. `browse` runs an interactive viewport; under `--json` (or `--no-interactive`) it falls back to `list` output. Query flags come from the index's schema (`--cve`, `--alias`, `--limit`, `--cursor`, etc.); see the [API docs](https://docs.vulncheck.com/api/indice) for the full set.


### backup

```
vulncheck backup url <index>       # signed temporary URL only
vulncheck backup download <index>  # download the archive
```

`download` picks the bubbletea progress bar for TTYs and a plain SIGINT-safe streaming download (writing to `<name>.part` and renaming on success) for headless callers. `url --json` returns `{filename, sha256, date_added, url}`.


### cpe

```
vulncheck cpe <cpe>
vulncheck cpe --from-file cpes.txt --json
echo cpe:2.3:… | vulncheck cpe --json
```

Batch mode (multiple positional args, `--from-file`, or piped stdin) requires `--json` and returns the [stable batch envelope](#batch-input).


### purl

```
vulncheck purl <purl>
vulncheck purl --from-file purls.txt --json
```

Same batch semantics as `cpe`. Batch requests are sent as a single `/v3/purls` POST rather than N GETs.


### tag

```
vulncheck tag <tag-name>
vulncheck tag --from-file tags.txt --json
```

Returns the newline-split list of matches for the given IP-intelligence tag. Batch input supported.


### pdns

```
vulncheck pdns <list-name>
vulncheck pdns --from-file lists.txt --json
```

Passive-DNS list membership. Batch input supported.


### rule

```
vulncheck rule <rule-name> [--table]
```

Look up an initial-access-intelligence rule. `--table` renders single-column table output; `--json` (global) supersedes.


### scan

```
vulncheck scan <path> [flags]
vulncheck scan --sbom-input-file <file> --json
```

Generates an SBOM for `<path>`, extracts PURLs, then either calls the vulncheck API (default) or queries local offline indices.

| Flag | Description |
|------|-------------|
| `-f`, `--file` | Save results to a file. |
| `-n`, `--file-name` | Custom output filename (default `output.json`). |
| `-o`, `--sbom-output-file` | Save the generated SBOM to a file. |
| `-i`, `--sbom-input-file` | Load an existing SBOM instead of generating one. |
| `-s`, `--sbom-only` | Generate the SBOM without running the vuln lookup. |
| `-c`, `--include-cpes` | Extract CPEs as well as PURLs (offline mode only, for now). |
| `--offline` | Use locally-synced indices instead of the API. |
| `--offline-meta` | Populate metadata (CVSS, KEV, description) from `vulncheck-nvd2` in offline mode. |
| `--warn-on-index` | Warn instead of failing when a required offline index isn't cached. |
| `--disable-ui` | Alias for `--no-interactive` (kept for backwards compatibility). |
| `--enrich` | Enrich the generated SBOM with metadata from `proxy.golang.org` / Maven Central / NPM / PyPI. Comma-separated scopes: `all`, `golang`, `java`, `javascript`, `python`; prefix with `-` to exclude (e.g. `all,-java`). Off by default; requires network — cannot be combined with `--offline` or `--sbom-input-file`. |

The progress TUI is auto-suppressed whenever the renderer can't safely draw it (`--json`, `--no-interactive`, non-TTY, CI).


### offline

```
vulncheck offline sync   [--add <name>|--remove <name>|--purge|--force|--choose] [--json]
vulncheck offline status [--json]
vulncheck offline purl   <purl> [--json]
vulncheck offline cpe    <cpe>  [--json] [--stats]
vulncheck offline ipintel <3d|10d|30d> [--country=...] [--asn=...] [--cidr=...] [--json]
```

Local queries against synced indices. `sync --json` returns `{schema_version, action, selected, elapsed_seconds}`; `status --json` returns the cached-index array. Under `--no-interactive` the sync command refuses without an explicit `--add` / `--remove` / `--purge` (no picker prompt).


### version

```
vulncheck version [--json]
```

`--json` returns `{schema_version, version, build_date, changelog_url}` — the canonical probe for agents.


### upgrade

```
vulncheck upgrade status
vulncheck upgrade latest [--force]
vulncheck upgrade --version X.X.X
```

- `upgrade status` — check whether a newer release is available.
- `upgrade latest` — install the newest release. `--force` reinstalls the current version.
- `upgrade --version X.X.X` — install a specific version.


> [!TIP]
> Looking to plug this into your Github Repository? Check out our own [Action](https://github.com/vulncheck-oss/action)
