$ErrorActionPreference = 'Stop'

# Read the OS architecture (not the process architecture — PROCESSOR_ARCHITEW6432
# is set when a 32-bit PowerShell is running on a 64-bit OS).
$procArch = if ($env:PROCESSOR_ARCHITEW6432) { $env:PROCESSOR_ARCHITEW6432 } else { $env:PROCESSOR_ARCHITECTURE }
$arch = switch ($procArch.ToUpper()) {
    "AMD64" { "amd64" }
    "X86"   { "386" }
    "ARM64" { "arm64" }
    default {
        Write-Error "Unsupported Windows architecture: $procArch (supported: AMD64, X86, ARM64)"
        exit 1
    }
}
$latestRelease = Invoke-RestMethod -Uri "https://api.github.com/repos/vulncheck-oss/cli/releases/latest"
$asset = $latestRelease.assets | Where-Object { $_.name -like "*windows_$arch.zip" } | Select-Object -First 1

if (-not $asset) {
    Write-Error "No suitable release asset found for Windows $arch"
    exit 1
}

$downloadUrl = $asset.browser_download_url
$zipPath = Join-Path $env:TEMP "vulncheck-cli.zip"
$extractPath = Join-Path $env:TEMP "vulncheck-cli"

Invoke-WebRequest -Uri $downloadUrl -OutFile $zipPath

Expand-Archive -Path $zipPath -DestinationPath $extractPath -Force

$exePath = Get-ChildItem -Path $extractPath -Filter "vulncheck.exe" -Recurse | Select-Object -First 1 -ExpandProperty FullName

if (-not $exePath) {
    Write-Error "vulncheck.exe not found in the extracted files"
    exit 1
}

$installPath = "$env:LOCALAPPDATA\Programs\vulncheck"
New-Item -ItemType Directory -Path $installPath -Force | Out-Null

Move-Item -Path $exePath -Destination $installPath -Force

$envPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($envPath -notlike "*$installPath*") {
    [Environment]::SetEnvironmentVariable("Path", "$envPath;$installPath", "User")
}

Write-Host "Vulncheck CLI has been installed to $installPath"
Write-Host "You may need to restart your shell for the changes to take effect"
