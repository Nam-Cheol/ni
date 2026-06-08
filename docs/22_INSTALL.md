# Install ni

`ni` is usable from source, from a locally built binary, from the verified
v0.5.1 GitHub Release archives, or through the verified curl installer.

## Prerequisites

Release binary and curl installer paths do not require Go.

Source and local build paths require:

- Go 1.22 or newer.
- Git, if you want builds to include a git-derived version string.

## Primary README Paths

README intentionally shows only two primary first-success paths:

| Platform | Primary path | Verify | First project | Uninstall |
| --- | --- | --- | --- | --- |
| macOS | Install the latest release with `curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh \| sh -s -- --update-path`. | Open a new shell and run `ni --help` and `ni version`. | `mkdir my-project`, `cd my-project`, `ni init .` | `curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh \| sh -s -- --uninstall` |
| Windows | Download to a temp path with `$Installer = Join-Path $env:TEMP "ni-install.ps1"`, `irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile $Installer`, then run `powershell -NoProfile -ExecutionPolicy Bypass -File $Installer`. | Open a new PowerShell session and run `ni --help` and `ni version`. | `mkdir my-project`, `cd my-project`, `ni init .` | `$Installer = Join-Path $env:TEMP "ni-install.ps1"`, `irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile $Installer`, `powershell -NoProfile -ExecutionPolicy Bypass -File $Installer -Uninstall` |

These paths prove global command-name resolution first. They do not run agents,
execute generated prompts, or prove downstream implementation readiness.

## Install Path Status

Every public install path has exactly one status:

| Path | Status | Notes |
| --- | --- | --- |
| Source | Available | Run `go run ./cmd/ni ...` from this checkout. |
| Local binary | Available | Build or install locally from this checkout. |
| Release binary | Available | Use the verified v0.5.1 GitHub Release archives and checksums. |
| Curl installer | Available | Use the verified v0.5.1 `install.sh` path after inspecting the script. |
| Homebrew | Planned | No tap or formula is published or tested. |
| Model workspaces | Experimental | Repo-local model assistance can draft docs; the CLI remains authority. |
| No-terminal method | Experimental | Assisted planning only; deterministic validation still requires CLI proof. |

Homebrew: Planned / v0.5 candidate. Package-manager work is documented as a
planned path only; do not present Homebrew as an available README install path.

## Run from source

Use this mode when developing `ni` or trying the CLI without creating a binary:

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
go run ./cmd/ni status --dir .
```

Source runs use the default development version unless you pass linker flags
manually.

## Build a local binary

Build into `bin/ni`:

```bash
make build
./bin/ni --help
./bin/ni version
```

`make build` injects the version from:

```bash
git describe --tags --always --dirty
```

If git metadata is unavailable, the build falls back to `0.0.0-dev`.

## Install locally

Install to `~/.local/bin/ni` by default:

```bash
make install-local
PATH="$HOME/.local/bin:$PATH" ni version
```

To choose another install location, override `PREFIX` or `BINDIR`:

```bash
make install-local PREFIX=/usr/local
make install-local BINDIR="$HOME/bin"
```

Ensure the chosen directory is on your `PATH` before running `ni` by name.
For verification or tests, use a temporary `BINDIR` instead of a user-owned
install directory.

```bash
tmpdir="$(mktemp -d)"
make install-local BINDIR="$tmpdir/bin"
PATH="$tmpdir/bin:$PATH" sh -c 'ni --help && ni version'
```

## Release binary

Release binary status: Available.

The v0.5.1 GitHub Release contains the expected OS/arch archives and
`ni_0.5.1_checksums.txt`:

<https://github.com/Nam-Cheol/ni/releases/tag/v0.5.1>

Use this matrix to choose the archive for your OS and architecture:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

Each valid release must include `ni_<version>_checksums.txt`. For v0.5.1,
download the archive and checksum file from the same release, verify the
archive, unpack the binary into a directory on `PATH`, and then run `ni --help`
and `ni version` by command name.

Linux example:

```bash
VERSION="0.5.1"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_linux_amd64.tar.gz"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_checksums.txt"
grep " ni_${VERSION}_linux_amd64.tar.gz$" "ni_${VERSION}_checksums.txt" | sha256sum -c -
tar -xzf "ni_${VERSION}_linux_amd64.tar.gz"
install -m 0755 ni "$HOME/.local/bin/ni"
PATH="$HOME/.local/bin:$PATH" ni --help
PATH="$HOME/.local/bin:$PATH" ni version
```

macOS example:

```bash
VERSION="0.5.1"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_darwin_arm64.tar.gz"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_checksums.txt"
grep " ni_${VERSION}_darwin_arm64.tar.gz$" "ni_${VERSION}_checksums.txt" | shasum -a 256 -c -
tar -xzf "ni_${VERSION}_darwin_arm64.tar.gz"
install -m 0755 ni "$HOME/.local/bin/ni"
PATH="$HOME/.local/bin:$PATH" ni --help
PATH="$HOME/.local/bin:$PATH" ni version
```

Windows PowerShell:

```powershell
$Version = "0.5.1"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_windows_amd64.zip" -OutFile "ni_$($Version)_windows_amd64.zip"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_checksums.txt" -OutFile "ni_$($Version)_checksums.txt"
Get-FileHash "ni_$($Version)_windows_amd64.zip" -Algorithm SHA256
Select-String "ni_$($Version)_windows_amd64.zip" "ni_$($Version)_checksums.txt"
Expand-Archive "ni_$($Version)_windows_amd64.zip" -DestinationPath "ni_$($Version)_windows_amd64"
.\ni_$($Version)_windows_amd64\ni.exe --help
.\ni_$($Version)_windows_amd64\ni.exe version
```

Compare the `Get-FileHash` output with the checksum line printed by
`Select-String` before trusting the extracted binary. The v0.5.1 Windows asset
and checksum are verified in
[`132_V0_5_1_POST_RELEASE_VERIFICATION.md`](132_V0_5_1_POST_RELEASE_VERIFICATION.md),
but execution on a real Windows host remains manually unverified. Use the
PowerShell installer below for User PATH handling and global command setup.

## Curl installer

Curl installer status: Available for verified v0.5.1 release assets.

`install.sh` can install a release archive without requiring Go. It was
verified against the real v0.5.1 darwin/arm64 archive and checksum file on
2026-06-08. The installer downloads the selected archive and checksum file,
verifies the checksum when a local sha256 tool is available, installs only the
`ni` binary, and does not install model skills or run downstream work.

For the README latest-by-default path:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --update-path
```

If you omit `--version`, the installer asks GitHub for the latest release tag
during actual install. For reproducible installs, download and inspect the
script, then pin the release:

```bash
VERSION="0.5.1"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

The dry-run path is documented with `--version` because current dry-run output
does not resolve the latest GitHub release tag when `--version` is omitted.

Open a new shell so the PATH update is loaded. First, check that the global
command is available:

```bash
ni --help
```

Then check the installed version:

```bash
ni version
```

See [Curl Installer](install-curl.md) for `BINDIR`, checksum behavior, and the
manual verification path. The manual verification path is to download the
matching archive and `ni_0.5.1_checksums.txt` from the same release, verify the
archive checksum, extract it into a directory on `PATH`, and then run
`ni --help` and `ni version` by command name.

Uninstall the curl-installed binary and any ni-managed PATH block with the
README default path:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --uninstall
```

If you already downloaded the installer or installed into a custom `BINDIR`,
run uninstall from that local script and pass the same `BINDIR`:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

## Windows PowerShell installer

`install.ps1` installs `ni.exe` to `%LOCALAPPDATA%\ni\bin\ni.exe` by default
and adds that directory to User PATH only. It does not modify Machine PATH by
default and does not use `setx`.

PowerShell also has a built-in alias:

```powershell
ni -> New-Item
```

Without an alias fix, `ni --help` can invoke `New-Item` and create a file named
`--help` instead of running `ni.exe`. The installer therefore creates `$PROFILE`
only if needed, preserves existing profile content, and adds this ni-managed
block only once:

```powershell
# >>> ni installer >>>
Remove-Item Alias:ni -Force -ErrorAction SilentlyContinue
# <<< ni installer <<<
```

The block removes the PowerShell alias in new sessions so PATH can resolve
`ni.exe` when the user types `ni`. If the profile update fails, the installer
prints the exact block as the manual fix instead of silently claiming command
resolution is complete.

```powershell
$Installer = Join-Path $env:TEMP "ni-install.ps1"
```

Download the installer to that path:

```powershell
irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile $Installer
```

Run the downloaded installer:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File $Installer
```

If you omit `-Version`, the installer asks GitHub for the latest release tag
during actual install. For reproducible installs, inspect and dry-run the
script from the directory where you downloaded `install.ps1`, then pin the
release:

```powershell
$Version = "0.5.1"
irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile install.ps1
Get-Content .\install.ps1
.\install.ps1 -DryRun -Version $Version
.\install.ps1 -Version $Version
```

Open a new PowerShell session so the User PATH update and profile block are
loaded. First, check command resolution:

```powershell
Get-Command ni -All
```

`Get-Command ni` should resolve to `ni.exe` after the profile block loads. Then
check that the global command is available:

```powershell
ni --help
```

Then check the installed version:

```powershell
ni version
```

Uninstall removes `ni.exe`, removes the install directory if empty, removes
only the matching `ni` directory from User PATH, and removes only the
ni-managed PowerShell profile block. It preserves unrelated profile content and
unrelated PATH entries:

```powershell
$Installer = Join-Path $env:TEMP "ni-install.ps1"
```

Download a fresh copy of the installer:

```powershell
irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile $Installer
```

Run the installer in uninstall mode:

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File $Installer -Uninstall
```

Windows execution has not been verified on this macOS host. Do not claim it
verified until a real Windows PowerShell install, new-session help/version, and
uninstall transcript exists.

Package manager status: Planned. Do not use package manager instructions for
`ni` yet; Homebrew and Scoop packages are not published.

## Validation

The supported local validation commands are:

```bash
make test
make quality
make smoke
make build
make install-check
```

These map to the repository's existing Go tests, quality checks, smoke tests,
local binary build, and source/local-install verification. `make install-check`
uses a temporary install path and is also part of `bash scripts/release-check.sh`;
it is not required for every quality run.

For release configuration validation, run `goreleaser check` when GoReleaser is
installed locally. The repository release gate remains:

```bash
bash scripts/release-check.sh
```

## License

`ni` is licensed under the [MIT License](../LICENSE).

This install document claims release binary availability only for the verified
v0.5.1 GitHub Release assets and curl installer availability only for the
verified v0.5.1 installer path. It does not claim package distribution,
Homebrew support, Scoop support, no-terminal deterministic validation, runtime
execution behavior, Windows execution verification on macOS, or global
model-pack installation.
