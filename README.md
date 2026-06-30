<p align="center">
    <img src="/logo-cli.png" align="center" alt="VulnCheck Logo" width="150" />
</p>

# The VulnCheck CLI
`vulncheck` is access to the VulnCheck API on the command line. It brings index browsing, backup management, and vulnerability scanning to the terminal.

<p align="center">
    <img src="/vulncheck-scan.gif" />
</p>

[![Release](https://img.shields.io/github/v/release/vulncheck-oss/cli)](https://github.com/vulncheck-oss/cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/vulncheck-oss/cli)](https://goreportcard.com/report/github.com/vulncheck-oss/cli)
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


## Agentic / scripted usage

The CLI is designed to be safe to drive from scripts and AI agents. This section is the contract — the surfaces below are intended to remain stable across releases (additions are not breaking changes; renames / removals are).

### Global flags

| Flag                | Effect |
|---------------------|--------|
| `--json`            | Emit JSON on stdout; route info/progress lines to stderr; errors emitted as a structured envelope. |
| `--quiet`           | Suppress informational output. Errors and payloads still render. |
| `--verbose` / `-v`  | Extra detail on stderr (debug-level). |
| `--no-color`        | Disable ANSI styling. Also honours the `NO_COLOR` env var. |
| `--no-interactive` | Refuse to block on TUI prompts; commands that need a prompt return an error instead. Implied by `--json`, non-TTY stdin/stdout, and any of the `CI` / `BUILD_NUMBER` / `RUN_ID` env vars. |

### Environment variables

| Variable                                   | Effect |
|--------------------------------------------|--------|
| `VC_TOKEN`                                 | API token. Takes precedence over `~/.config/vulncheck/vulncheck.yaml`. |
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
    "http_status": 401
  }
}
```

`code` is one of: `internal`, `validation`, `auth_required`, `auth_invalid`, `not_found`, `rate_limited`, `network`, `bad_request`, `cancelled`. `http_status` is omitted for non-HTTP errors.

### Probe commands

Use these to inspect the CLI itself before dispatching work:

```bash
vulncheck version --json
# {"schema_version": 1, "version": "...", "build_date": "...", "changelog_url": "..."}

vulncheck auth status --json
# {"schema_version": 1, "authenticated": true, "token_source": "env", "user": "...", "email": "..."}
# Exit 0 even when authenticated=false — agents dispatch on the bool.
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

- [Browse/list indices](#browselist-indices)
- [Browse/list an index](#browselist-an-index)
- [Download a backup](#download-a-backup)
- [Request vulnerabilities related to a CPE](#request-vulnerabilities-related-to-a-cpe)
- [Request vulnerabilities related to a PURL](#request-vulnerabilities-related-to-a-purl)
- [Scan a repository for vulnerabilities](#scan-a-repository-for-vulnerabilities)
- [Upgrade the VulnCheck CLI](#upgrade-the-vulncheck-cli)


### Browse/list indices
You can browse all available indices interactively or output them as a list

```
vulncheck indices browse|list <search> [flags]
```

You can search for a specific index by passing a search term.

> [!TIP]
> Pressing `[Enter]` on an index while browsing will begin browsing that particular index

#### Flags (list only)

| Flag   | Description                                |
|--------|--------------------------------------------|
| --json | Output the list of indices in JSON format. |



### Browse/list an index

You can browse the contents of any index interactively or output some as JSON

```
vulncheck index browse|list <index> [flags]
```

#### Flags
 
| Flag                   | Type   | Description           |
|------------------------|--------|-----------------------|
| --alias                | string | Alias                 |
| --asn                  | string | Asn                   |
| --botnet               | string | Botnet                |
| --cidr                 | string | Cidr                  |
| --country              | string | Country               |
| --country_code         | string | CountryCode           |
| --cursor               | string | Cursor                |
| --cve                  | string | Cve                   |
| --hostname             | string | Hostname              |
| --iava                 | string | Iava                  |
| --id                   | string | ID                    |
| --ilvn                 | string | Ilvn                  |
| --jvndb                | string | Jvndb                 |
| --kind                 | string | Kind                  |
| --lastModEndDate       | string | LastModEndDate        |
| --lastModStartDate     | string | LastModStartDate      |
| --limit                | string | Limit                 |
| --misp_id              | string | MispId                |
| --mitre_id             | string | MitreId               |
| --order                | string | Order                 |
| --page                 | string | Page                  |
| --pubEndDate           | string | PubEndDate            |
| --pubStartDate         | string | PubStartDate          |
| --ransomware           | string | Ransomware            |
| --sort                 | string | Sort                  |
| --start_cursor         | string | StartCursor           |
| --threat_actor         | string | ThreatActor           |
| --updatedAtEndDate     | string | UpdatedAtEndDate      |
| --updatedAtStartDate   | string | UpdatedAtStartDate    |
| --date                 | string | Date                  |
| --src_country          | string | SrcCountry            |
| --dst_country          | string | DstCountry            |
| --src_ip               | string | SrcIp                 |
| --src_asn              | string | SrcASN                |
| --help                 |        | Show help for command |


### Download a backup 

Download a backup of a specified index either interactively or retrieve a signed temporary URL

```
vulncheck backup download|url <index>
```

#### Flags (url only)

| Flag   | Description                             |
|--------|-----------------------------------------|
| --json | Output the download URL in JSON format. |




### Request vulnerabilities related to a CPE

Based on the specified CPE (Common Platform Enumeration) URI string, this endpoint will return a list of vulnerabilities that are related to the package. We support v2.2 and v2.3

```
vulncheck cpe <cpe>
```


### Request vulnerabilities related to a PURL

Based on the specified PURL, this command will return a list of vulnerabilities that are related to the package.
You can find a list of supported package managers [here](https://docs.vulncheck.com/products/exploit-and-vulnerability-intelligence/package-manager-support)

```
vulncheck purl <purl>
```


### Scan a repository for vulnerabilities
This command will scan a directory for traces of packages via generating an SBOM and then check for vulnerabilities.

```
vulncheck scan <path> [flags]
```

#### Flags
| Flag | Description                        |
|------|------------------------------------|
| -f   | Save scan results to `output.json` |


### Upgrade the VulnCheck CLI
To check for updates and upgrade to the latest version of the VulnCheck CLI, use the following commands:
```
vulncheck upgrade status
vulncheck upgrade latest
vulncheck upgrade --version X.X.X
```

To see if a new version is available, run `vulncheck upgrade status`. If an update is available, you can upgrade to the latest version by running `vulncheck upgrade latest`. 

You can use the `--force` flag with the `latest` command to reinstall the current version if needed.

If you want to install a specific version, you can use the `--version` flag followed by the desired version number.

* `vulncheck upgrade` - Shows help
* `vulncheck upgrade --version X.X.X` - Upgrades to specific version
* `vulncheck upgrade latest` - Upgrades to latest version
* `vulncheck upgrade latest --force` - Force upgrade to latest version
* `vulncheck upgrade status` - Check upgrade status


> [!TIP]
> Looking to plug this into your Github Repository? Check out our own [Action](https://github.com/vulncheck-oss/action)
