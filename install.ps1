param(
    [string]$Version = "",
    [string]$Repo = "Nam-Cheol/ni",
    [string]$InstallDir = "",
    [string]$BaseUrl = "",
    [switch]$DryRun,
    [switch]$Uninstall
)

$ErrorActionPreference = "Stop"

if ([string]::IsNullOrWhiteSpace($InstallDir)) {
    $InstallDir = Join-Path $env:LOCALAPPDATA "ni\bin"
}

$InstallDir = [System.IO.Path]::GetFullPath($InstallDir)
$Target = Join-Path $InstallDir "ni.exe"

function Write-Step {
    param([string]$Message)
    Write-Host $Message
}

function Get-UserPathEntries {
    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ([string]::IsNullOrWhiteSpace($userPath)) {
        return @()
    }
    return @($userPath -split ";" | Where-Object { -not [string]::IsNullOrWhiteSpace($_) })
}

function Set-UserPathEntries {
    param([string[]]$Entries)
    [Environment]::SetEnvironmentVariable("Path", ($Entries -join ";"), "User")
}

function Test-PathEntryPresent {
    param([string[]]$Entries, [string]$Entry)
    foreach ($existing in $Entries) {
        if ([string]::Equals($existing.TrimEnd("\"), $Entry.TrimEnd("\"), [System.StringComparison]::OrdinalIgnoreCase)) {
            return $true
        }
    }
    return $false
}

function Add-UserPathEntry {
    param([string]$Entry)
    $entries = @(Get-UserPathEntries)
    if (Test-PathEntryPresent -Entries $entries -Entry $Entry) {
        Write-Step "User PATH already contains $Entry"
        return
    }
    Set-UserPathEntries -Entries @($entries + $Entry)
    Write-Step "Added $Entry to User PATH"
}

function Remove-UserPathEntry {
    param([string]$Entry)
    $entries = @(Get-UserPathEntries)
    $kept = @()
    $removed = $false
    foreach ($existing in $entries) {
        if ([string]::Equals($existing.TrimEnd("\"), $Entry.TrimEnd("\"), [System.StringComparison]::OrdinalIgnoreCase)) {
            $removed = $true
            continue
        }
        $kept += $existing
    }
    if ($removed) {
        Set-UserPathEntries -Entries $kept
        Write-Step "Removed $Entry from User PATH"
    }
}

function Resolve-LatestTag {
    $release = Invoke-RestMethod "https://api.github.com/repos/$Repo/releases/latest"
    return $release.tag_name
}

Write-Step "ni Windows installer"
Write-Step "  repository: $Repo"
Write-Step "  install to: $Target"
Write-Step "  PATH scope: User"

if ($Uninstall) {
    if (Test-Path $Target) {
        Remove-Item $Target -Force
        Write-Step "Removed $Target"
    } else {
        Write-Step "No installed ni.exe found at $Target"
    }

    if ((Test-Path $InstallDir) -and -not (Get-ChildItem $InstallDir -Force)) {
        Remove-Item $InstallDir -Force
        Write-Step "Removed empty directory $InstallDir"
    }

    Remove-UserPathEntry -Entry $InstallDir
    Write-Step "Uninstall complete. Open a new PowerShell session and verify ni is no longer found."
    exit 0
}

if ([string]::IsNullOrWhiteSpace($Version)) {
    if ($DryRun) {
        $Tag = "<latest>"
        $Version = "<latest>"
    } else {
        $Tag = Resolve-LatestTag
        $Version = $Tag.TrimStart("v")
    }
} else {
    $Tag = $Version
    if (-not $Tag.StartsWith("v")) {
        $Tag = "v$Tag"
    }
    $Version = $Version.TrimStart("v")
}

$Asset = "ni_${Version}_windows_amd64.zip"
$Checksums = "ni_${Version}_checksums.txt"
if ([string]::IsNullOrWhiteSpace($BaseUrl)) {
    $BaseUrl = "https://github.com/$Repo/releases/download/$Tag"
}
$AssetUrl = "$($BaseUrl.TrimEnd("/"))/$Asset"
$ChecksumUrl = "$($BaseUrl.TrimEnd("/"))/$Checksums"

Write-Step "  asset: $Asset"
Write-Step "  checksums: $Checksums"

if ($DryRun) {
    Write-Step "  mode: dry-run"
    Write-Step "Would download:"
    Write-Step "  $AssetUrl"
    Write-Step "  $ChecksumUrl"
    Write-Step "Would add install directory to User PATH if missing."
    exit 0
}

$TempRoot = Join-Path ([System.IO.Path]::GetTempPath()) ("ni-install-" + [System.Guid]::NewGuid().ToString("N"))
New-Item -ItemType Directory -Path $TempRoot | Out-Null

try {
    $ArchivePath = Join-Path $TempRoot $Asset
    $ChecksumPath = Join-Path $TempRoot $Checksums
    $ExtractDir = Join-Path $TempRoot "extract"

    Invoke-WebRequest $AssetUrl -OutFile $ArchivePath
    Invoke-WebRequest $ChecksumUrl -OutFile $ChecksumPath

    $actual = (Get-FileHash $ArchivePath -Algorithm SHA256).Hash.ToLowerInvariant()
    $line = Select-String -Path $ChecksumPath -Pattern ([regex]::Escape($Asset)) | Select-Object -First 1
    if ($null -eq $line) {
        throw "checksum file does not contain $Asset"
    }
    $expected = (($line.Line -split "\s+")[0]).ToLowerInvariant()
    if ($actual -ne $expected) {
        throw "checksum mismatch for $Asset"
    }
    Write-Step "Verified checksum for $Asset"

    Expand-Archive $ArchivePath -DestinationPath $ExtractDir -Force
    $Found = Get-ChildItem $ExtractDir -Recurse -Filter "ni.exe" | Select-Object -First 1
    if ($null -eq $Found) {
        throw "archive did not contain ni.exe"
    }

    New-Item -ItemType Directory -Path $InstallDir -Force | Out-Null
    Copy-Item $Found.FullName $Target -Force
    Write-Step "Installed ni.exe to $Target"

    Add-UserPathEntry -Entry $InstallDir

    Write-Step ""
    Write-Step "Open a new PowerShell session and verify the global command:"
    Write-Step "  ni --help"
    Write-Step "  ni version"
    Write-Step ""
    Write-Step "Uninstall:"
    Write-Step '  $Installer = Join-Path $env:TEMP "ni-install.ps1"'
    Write-Step '  irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile $Installer'
    Write-Step "  powershell -NoProfile -ExecutionPolicy Bypass -File `$Installer -Uninstall"
    Write-Step ""
    Write-Step "The installer does not install model skills or run downstream work."
}
finally {
    if (Test-Path $TempRoot) {
        Remove-Item $TempRoot -Recurse -Force
    }
}
