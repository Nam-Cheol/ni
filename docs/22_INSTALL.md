# Install Namba Intent

Namba Intent is usable from source, from a locally built current-tree binary,
and from the published v0.6.0 release as `namba-intent`.

The v0.6.0 macOS curl installer path has been published and verified for
darwin/arm64. Historical v0.5.1 GitHub Release evidence remains tied to the
older `ni` command.

## Prerequisites

Release binary and curl installer paths do not require Go.

Source and local build paths require:

- Go 1.22 or newer.
- Git, if you want builds to include a git-derived version string.

## Primary README Paths

README intentionally shows only two primary first-success paths:

| Platform | Primary path | Verify | First project | Uninstall |
| --- | --- | --- | --- | --- |
| macOS | `curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh \| sh -s -- --update-path --version 0.6.0` | Open a new shell and run `namba-intent --help` and `namba-intent version`. | `mkdir my-project`, `cd my-project`, `namba-intent init .` | `curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh \| sh -s -- --uninstall` |
| Windows | Download current-main `install.ps1` to `$Installer = Join-Path $env:TEMP "namba-intent-install.ps1"` and inspect it before use. | After install, open a new PowerShell session and run `namba-intent --help` and `namba-intent version`; real-host verification is still pending. | `mkdir my-project`, `cd my-project`, `namba-intent init .` | `$Installer = Join-Path $env:TEMP "namba-intent-install.ps1"`, `irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile $Installer`, `powershell -NoProfile -ExecutionPolicy Bypass -File $Installer -Uninstall` |

These paths prove global command-name resolution first. They do not run agents,
execute generated prompts, or prove downstream implementation readiness.

## Install Path Status

Every public install path has exactly one status:

| Path | Status | Notes |
| --- | --- | --- |
| Source | Available | Run `go run ./cmd/namba-intent ...` from this checkout. |
| Local binary | Available | Build or install `namba-intent` locally from this checkout. |
| Release binary | Available | Use the verified v0.6.0 GitHub Release archives and checksums for `namba-intent`. Historical v0.5.1 evidence remains `ni` only. |
| Curl installer | Available | `install.sh --version 0.6.0` retrieves and verifies published `namba-intent_<version>...` assets on macOS darwin/arm64. |
| Homebrew | Planned | No tap or formula is published or tested. |
| Model workspaces | Experimental | Repo-local model assistance can draft docs; the CLI remains authority. |
| No-terminal method | Experimental | Assisted planning only; deterministic validation still requires CLI proof. |

Homebrew: Planned / v0.5 candidate. Package-manager work is documented as a
planned path only; do not present Homebrew as an available README install path.

## Run from source

Use this mode when developing Namba Intent or trying the CLI without creating a
binary:

```bash
go run ./cmd/namba-intent --help
go run ./cmd/namba-intent version
go run ./cmd/namba-intent status --dir .
```

Source runs use the default development version unless you pass linker flags
manually.

## Build a local binary

Build into `bin/namba-intent`:

```bash
make build
./bin/namba-intent --help
./bin/namba-intent version
```

`make build` injects the version from:

```bash
git describe --tags --always --dirty
```

If git metadata is unavailable, the build falls back to `0.0.0-dev`.

## Install locally

Install to `~/.local/bin/namba-intent` by default:

```bash
make install-local
PATH="$HOME/.local/bin:$PATH" namba-intent version
```

To choose another install location, override `PREFIX` or `BINDIR`:

```bash
make install-local PREFIX=/usr/local
make install-local BINDIR="$HOME/bin"
```

Ensure the chosen directory is on your `PATH` before running `namba-intent` by
name.
For verification or tests, use a temporary `BINDIR` instead of a user-owned
install directory.

```bash
tmpdir="$(mktemp -d)"
make install-local BINDIR="$tmpdir/bin"
PATH="$tmpdir/bin:$PATH" sh -c 'namba-intent --help && namba-intent version'
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

Curl installer status: Available for verified v0.6.0 macOS `namba-intent`
retrieval.

Current-main `install.sh` can install a published `namba-intent` release archive
without requiring Go. It selects `namba-intent_<version>_<os>_<arch>` archives,
downloads `namba-intent_<version>_checksums.txt`, verifies the checksum when a
local sha256 tool is available, installs only the `namba-intent` binary, and
does not install model skills or run downstream work.

Historical v0.5.1 curl installer verification remains valid only for the
previous `ni` public release path recorded in
[`132_V0_5_1_POST_RELEASE_VERIFICATION.md`](132_V0_5_1_POST_RELEASE_VERIFICATION.md).
The public v0.6.0 macOS curl installer path is verified in
[`140_V0_6_0_POST_RELEASE_VERIFICATION.md`](140_V0_6_0_POST_RELEASE_VERIFICATION.md).

For the verified v0.6.0 README path:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --update-path --version 0.6.0
```

If you omit `--version`, the installer asks GitHub for the latest release tag
during actual install. Use `--version 0.6.0` when you want to reproduce the
verified Namba Intent release path.

For reproducible v0.6.0 checks, download and inspect the script, then pin the
verified release version:

```bash
VERSION="0.6.0"
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
namba-intent --help
```

Then check the installed version:

```bash
namba-intent version
```

See [Curl Installer](install-curl.md) for `BINDIR`, checksum behavior, and the
manual verification path. The manual verification path is to download the
matching archive and `namba-intent_<version>_checksums.txt` from the same
release, verify the archive checksum, extract it into a directory on `PATH`,
and then run `namba-intent --help` and `namba-intent version` by command name.

Uninstall the curl-installed binary and any namba-intent-managed PATH block
with the README default path:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --uninstall
```

If you already downloaded the installer or installed into a custom `BINDIR`,
run uninstall from that local script and pass the same `BINDIR`:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

## Windows PowerShell installer

`install.ps1` installs `namba-intent.exe` to
`%LOCALAPPDATA%\namba-intent\bin\namba-intent.exe` by default and adds that
directory to User PATH only. It does not modify Machine PATH by default and
does not use `setx`.

PowerShell alias cleanup for `ni -> New-Item` is legacy v0.5.x guidance and is
not required for the primary `namba-intent.exe` path.

```powershell
$Installer = Join-Path $env:TEMP "namba-intent-install.ps1"
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
during actual install. Windows host execution is still pending until a real
Windows transcript exists.

For reproducible v0.6.0 checks, inspect and dry-run the script from the
directory where you downloaded `install.ps1`, then pin the release:

```powershell
$Version = "0.6.0"
irm https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.ps1 -OutFile install.ps1
Get-Content .\install.ps1
.\install.ps1 -DryRun -Version $Version
.\install.ps1 -Version $Version
```

Open a new PowerShell session so the User PATH update and profile block are
loaded. First, check command resolution:

```powershell
Get-Command namba-intent -All
```

`Get-Command namba-intent` should resolve to `namba-intent.exe`. Then check
that the global command is available:

```powershell
namba-intent --help
```

Then check the installed version:

```powershell
namba-intent version
```

Uninstall removes `namba-intent.exe`, removes the install directory if empty,
and removes only the matching `namba-intent` directory from User PATH. It
preserves unrelated PATH entries:

```powershell
$Installer = Join-Path $env:TEMP "namba-intent-install.ps1"
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
Namba Intent yet; Homebrew and Scoop packages are not published.

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

Namba Intent is licensed under the [MIT License](../LICENSE).

This install document claims v0.6.0 `namba-intent` release binary and macOS curl
installer availability only where verified. Historical v0.5.1 evidence remains
tied to `ni`. It does not claim package distribution, Homebrew support, Scoop support, no-terminal
deterministic validation, runtime execution behavior, Windows execution
verification on macOS, or global model-pack installation.
