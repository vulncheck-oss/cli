<p align="center">
    <img src="/logo-cli.png" align="center" alt="VulnCheck Logo" width="150" />
</p>

# The VulnCheck CLI
`vci` is access to the VulnCheck API on the command line. It brings index browsing, backup management, and vulnerability scanning to the terminal.

<p align="center">
    <img src="/vci-scan.gif" />
</p>

[![Release](https://img.shields.io/github/v/release/vulncheck-oss/cli)](https://github.com/vulncheck-oss/cli/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/vulncheck-oss/cli)](https://goreportcard.com/report/github.com/vulncheck-oss/cli)
[![Go Reference](https://pkg.go.dev/badge/github.com/vulncheck-oss/cli.svg)](https://pkg.go.dev/github.com/vulncheck-oss/cli)
[![Lint](https://github.com/vulncheck-oss/cli/actions/workflows/lint.yml/badge.svg)](https://github.com/vulncheck-oss/cli/actions/workflows/lint.yml)
[![Tests](https://github.com/vulncheck-oss/cli/actions/workflows/test.yml/badge.svg)](https://github.com/vulncheck-oss/cli/actions/workflows/test.yml)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/vulncheck-oss/cli/pulls)

## Installation 

`vci` is available for MacOS, Linux, and Windows. You can download precompiled binaries from our [releases page](https://github.com/vulncheck-oss/cli/releases/latest)

> [!NOTE]
> Support for package managers is coming soon.


## Configuration
* Run `vci auth login` to authenticate with your VulnCheck account.
* Alternatively `vci` will respect the `VC_TOKEN` environment variable.
* `vci auth` by itself will show other options like checking your status and logging out.


## Available commands

- [Browse/list indices](#browselist-indices)
- [Browse/list an index](#browselist-an-index)
- [Download a backup](#download-a-backup)
- [Request vulnerabilities related to a CPE](#request-vulnerabilities-related-to-a-cpe)
- [Request vulnerabilities related to a PURL](#request-vulnerabilities-related-to-a-purl)
- [Scan a repository for vulnerabilities](#scan-a-repository-for-vulnerabilities)

### Browse/list indices
You can browse all available indices interactively or output them as a list

```
vci indices browse|list <search> [flags]
```

You can search for a specific index by passing a search term.

> [!TIP]
> Pressing `[Enter]` on an index while browsing will begin browsing that particular index

#### Flags (list only)

| Flag | Description |
|------|-------------|
| --json | Output the list of indices in JSON format. |


### Browse/list an index

You can browse the contents of any index interactively or output some as JSON

```
vci index browse|list <index> [flags]
```

#### Flags
 
| Flag | Description |
| ---- | ----------- | 
|  --alias `string` |              Alias |
|  --botnet `string` |             Botnet |
|  --cve `string` |                Cve |
|  --iava `string` |               Iava |
|  --lastmodenddate `string` |     LastModEndDate |
|  --lastmodstartdate `string` |   LastModStartDate |
|  --mispid `string` |             MispId |
|  --mitreid `string` |            MitreId |
|  --pubenddate `string` |         PubEndDate |
|  --pubstartdate `string` |       PubStartDate |
|  --ransomware `string` |         Ransomware |
|  --threatactor `string` |        ThreatActor |


### Download a backup 

Download a backup of a specified index either interactively or retrieve a signed temporary URL

```
vci backup download|url <index>
```

#### Flags (url only)

| Flag | Description |
| ---- | ----------- |
| --json | Output the download URL in JSON format. |



### Request vulnerabilities related to a CPE

Based on the specified CPE (Common Platform Enumeration) URI string, this endpoint will return a list of vulnerabilities that are related to the package. We support v2.2 and v2.3

```
vci cpe <cpe>
```

### Request vulnerabilities related to a PURL

Based on the specified PURL, this command will return a list of vulnerabilities that are related to the package.
You can find a list of supported package managers [here](https://docs.vulncheck.com/products/exploit-and-vulnerability-intelligence/package-manager-support)

```
vci purl <purl>
```

### Scan a repository for vulnerabilities
This command will scan a directory for traces of packages via generating an SBOM and then check for vulnerabilities.

`vci scan <path>`

#### Flags
| Flag | Description |
| ---- | ----------- |
| -f | Save scan results to `output.json` |

```





