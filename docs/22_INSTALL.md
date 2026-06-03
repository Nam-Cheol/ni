# Install ni

`ni` is usable from source, from a locally built binary, from the verified
v0.5.0 GitHub Release archives, or through the verified curl installer.

## Prerequisites

Release binary and curl installer paths do not require Go.

Source and local build paths require:

- Go 1.22 or newer.
- Git, if you want builds to include a git-derived version string.

## Install Path Status

Every public install path has exactly one status:

| Path | Status | Notes |
| --- | --- | --- |
| Source | Available | Run `go run ./cmd/ni ...` from this checkout. |
| Local binary | Available | Build or install locally from this checkout. |
| Release binary | Available | Use the verified v0.5.0 GitHub Release archives and checksums. |
| Curl installer | Available | Use the verified v0.5.0 `install.sh` path after inspecting the script. |
| Homebrew | Planned | No tap or formula is published or tested. |
| Model workspaces | Experimental | Repo-local model assistance can draft docs; the CLI remains authority. |
| No-terminal method | Experimental | Assisted planning only; deterministic validation still requires CLI proof. |

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
~/.local/bin/ni version
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
"$tmpdir/bin/ni" --help
"$tmpdir/bin/ni" version
```

## Release binary

Release binary status: Available.

The v0.5.0 GitHub Release contains the expected OS/arch archives and
`ni_0.5.0_checksums.txt`:

<https://github.com/Nam-Cheol/ni/releases/tag/v0.5.0>

Use this matrix to choose the archive for your OS and architecture:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

Each valid release must include `ni_<version>_checksums.txt`. For v0.5.0,
download the archive and checksum file from the same release, verify the
archive, unpack the binary, and then run `ni --help` and `ni version`.

Linux example:

```bash
VERSION="0.5.0"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_linux_amd64.tar.gz"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_checksums.txt"
grep " ni_${VERSION}_linux_amd64.tar.gz$" "ni_${VERSION}_checksums.txt" | sha256sum -c -
tar -xzf "ni_${VERSION}_linux_amd64.tar.gz"
./ni --help
./ni version
```

macOS example:

```bash
VERSION="0.5.0"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_darwin_arm64.tar.gz"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_checksums.txt"
grep " ni_${VERSION}_darwin_arm64.tar.gz$" "ni_${VERSION}_checksums.txt" | shasum -a 256 -c -
tar -xzf "ni_${VERSION}_darwin_arm64.tar.gz"
./ni --help
./ni version
```

Windows PowerShell:

```powershell
$Version = "0.5.0"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_windows_amd64.zip" -OutFile "ni_$($Version)_windows_amd64.zip"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_checksums.txt" -OutFile "ni_$($Version)_checksums.txt"
Get-FileHash "ni_$($Version)_windows_amd64.zip" -Algorithm SHA256
Select-String "ni_$($Version)_windows_amd64.zip" "ni_$($Version)_checksums.txt"
Expand-Archive "ni_$($Version)_windows_amd64.zip" -DestinationPath "ni_$($Version)_windows_amd64"
.\ni_$($Version)_windows_amd64\ni.exe --help
.\ni_$($Version)_windows_amd64\ni.exe version
```

Compare the `Get-FileHash` output with the checksum line printed by
`Select-String` before trusting the extracted binary. The v0.5.0 Windows asset
and checksum are verified in
[`117_V0_5_0_POST_RELEASE_VERIFICATION.md`](117_V0_5_0_POST_RELEASE_VERIFICATION.md),
but execution on a real Windows host remains manually unverified.

## Curl installer

Curl installer status: Available for verified v0.5.0 release assets.

`install.sh` can install a release archive without requiring Go. It was
verified against the real v0.5.0 darwin/arm64 archive and checksum file on
2026-06-02. The installer downloads the selected archive and checksum file,
verifies the checksum when a local sha256 tool is available, installs only the
`ni` binary, and does not install model skills or run downstream work.

Download and inspect the script before any local install:

```bash
VERSION="0.5.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --version "$VERSION"
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

See [Curl Installer](install-curl.md) for `BINDIR`, checksum behavior, and the
manual verification path. The manual verification path is to download the
matching archive and `ni_0.5.0_checksums.txt` from the same release, verify the
archive checksum, extract it, and then run `ni --help` and `ni version`.

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
v0.5.0 GitHub Release assets and curl installer availability only for the
verified v0.5.0 installer path. It does not claim package distribution,
Homebrew support, Scoop support, no-terminal deterministic validation, runtime
execution behavior, or global model-pack installation.
