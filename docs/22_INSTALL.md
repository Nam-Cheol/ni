# Install ni

`ni` is currently usable from source or from a locally built binary. A
GoReleaser-based release binary pipeline is prepared for future GitHub Releases,
but hosted release assets are not available until the first tagged release
publishes them.

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

## Release binary after first assets exist

Release binaries are planned for GitHub Releases, but no release asset should be
treated as available until a published release page includes the matching file
and checksum.

When those assets exist, use this matrix:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

Each release is also expected to include
`ni_<version>_checksums.txt`. After downloading a published archive and the
checksum file from the same release, verify the archive before placing `ni` on
your `PATH`.

Linux or macOS:

```bash
sha256sum -c ni_<version>_checksums.txt --ignore-missing
tar -xzf ni_<version>_<os>_<arch>.tar.gz
./ni version
```

macOS also provides `shasum -a 256 -c` if `sha256sum` is not installed.

Windows PowerShell:

```powershell
Expand-Archive ni_<version>_windows_amd64.zip -DestinationPath ni_<version>_windows_amd64
.\ni_<version>_windows_amd64\ni.exe version
```

## Curl installer after first assets exist

`install.sh` can install a release archive without requiring Go, but it remains
release-gated: use it only after the GitHub Release contains the matching archive
and checksum file.

Download and inspect the script before running it:

```bash
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run
sh install.sh --version 0.2.0
```

See [Curl Installer](install-curl.md) for `BINDIR`, checksum behavior, and the
manual verification path.

Do not use package manager instructions for `ni` yet; Homebrew and Scoop
packages are not published.

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

This install document does not claim package distribution, Homebrew support,
Scoop support, or currently available release binaries. Use source, local build,
or local install mode unless a published GitHub Release contains verified
assets.
