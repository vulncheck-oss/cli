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
Option 1: Using PowerShell
Open PowerShell and run:

```
iex ((New-Object System.Net.WebClient).DownloadString('https://raw.githubusercontent.com/vulncheck-oss/cli/main/install.ps1'))
```

`vulncheck` binaries are also available for MacOS, Linux, and Windows. You can download precompiled binaries from our [releases page](https://github.com/vulncheck-oss/cli/releases/latest)


> [!NOTE]
> Support for package managers is coming soon.


## Configuration
* Run `vulncheck auth login` to authenticate with your VulnCheck account.
* Alternatively `vulncheck` will respect the `VC_TOKEN` environment variable.
* `vulncheck auth` by itself will show other options like checking your status and logging out.


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
| --src_country          | string | SrcCountry            |
| --dst_country          | string | DstCountry            |
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
