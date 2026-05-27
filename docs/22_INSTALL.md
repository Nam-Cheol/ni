# Install ni

`ni` is usable from source or from a locally built binary. Hosted release
archives and the curl installer are release-gated until a GitHub Release
actually publishes binary assets and checksums, and the installer is verified
against those assets.

## Prerequisites

- Go 1.22 or newer.
- Git, if you want builds to include a git-derived version string.

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

Release binary status: Release-gated.

Do not treat GitHub Release downloads as an available install path until the
release page contains the expected archives and `ni_<version>_checksums.txt`:

<https://github.com/Nam-Cheol/ni/releases>

When a release exists, use this matrix to choose the archive for your OS and
architecture:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

Each valid release must include `ni_<version>_checksums.txt`. Download the
archive and checksum file from the same release, verify the archive, unpack the
binary, and then run `ni --help` and `ni version`.

Linux example after a release exists:

```bash
VERSION="<published-version-without-v>"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_linux_amd64.tar.gz"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_checksums.txt"
grep " ni_${VERSION}_linux_amd64.tar.gz$" "ni_${VERSION}_checksums.txt" | sha256sum -c -
tar -xzf "ni_${VERSION}_linux_amd64.tar.gz"
./ni --help
./ni version
```

macOS example after a release exists:

```bash
VERSION="<published-version-without-v>"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_darwin_arm64.tar.gz"
curl -fSLO "https://github.com/Nam-Cheol/ni/releases/download/v${VERSION}/ni_${VERSION}_checksums.txt"
grep " ni_${VERSION}_darwin_arm64.tar.gz$" "ni_${VERSION}_checksums.txt" | shasum -a 256 -c -
tar -xzf "ni_${VERSION}_darwin_arm64.tar.gz"
./ni --help
./ni version
```

Windows PowerShell after a release exists:

```powershell
$Version = "<published-version-without-v>"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_windows_amd64.zip" -OutFile "ni_$($Version)_windows_amd64.zip"
Invoke-WebRequest "https://github.com/Nam-Cheol/ni/releases/download/v$Version/ni_$($Version)_checksums.txt" -OutFile "ni_$($Version)_checksums.txt"
Get-FileHash "ni_$($Version)_windows_amd64.zip" -Algorithm SHA256
Select-String "ni_$($Version)_windows_amd64.zip" "ni_$($Version)_checksums.txt"
Expand-Archive "ni_$($Version)_windows_amd64.zip" -DestinationPath "ni_$($Version)_windows_amd64"
.\ni_$($Version)_windows_amd64\ni.exe --help
.\ni_$($Version)_windows_amd64\ni.exe version
```

Compare the `Get-FileHash` output with the checksum line printed by
`Select-String` before trusting the extracted binary.

## Curl installer

Curl installer status: Release-gated.

`install.sh` can install a release archive without requiring Go, but it depends
on hosted release archives. Do not present it as available until release assets
exist and the script has been verified against the real archives and checksum
file.

After that gate is satisfied, download and inspect the script before any local
trial:

```bash
VERSION="<published-version-without-v>"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --version "$VERSION"
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

See [Curl Installer](install-curl.md) for `BINDIR`, checksum behavior, and the
manual verification path.

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

This install document does not claim published GitHub Release binary assets,
curl installer availability, package distribution, Homebrew support, Scoop
support, no-terminal deterministic validation, runtime execution behavior, or
global model-pack installation.
