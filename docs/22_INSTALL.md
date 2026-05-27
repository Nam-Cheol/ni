# Install ni

`ni` is currently distributed from source. Package publishing, Homebrew taps,
GoReleaser, and automated release tooling are intentionally excluded at this
stage.

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

## License

`ni` is licensed under the [MIT License](../LICENSE).

This install document does not claim package distribution, Homebrew support,
GoReleaser support, or a published binary release. Use source, local build, or
local install mode unless a future release process documents another supported
distribution channel.
