$ErrorActionPreference = 'Stop'

$arch = if ([Environment]::Is64BitOperatingSystem) { "amd64" } else { "386" }
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
Write-Host "Please restart your terminal or run 'refreshenv' to update your PATH"
