# Curl Installer

`install.sh` is release asset infrastructure for installing a released `ni`
binary without requiring Go. It downloads an archive, verifies the checksum when
the release provides one, copies `ni` into a local bin directory, and prints next
steps. It does not install model skills or run downstream work.

Status: release-gated. GitHub Release archives and checksums are available, but
the curl installer should not be presented as an available public install path
until `install.sh` is verified against the real release assets.

## Safer Script Path

Download and inspect the installer first:

```bash
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run
```

By default, the script installs to `~/.local/bin/ni`. Override the destination
with `BINDIR`:

```bash
BINDIR="$HOME/bin" sh install.sh --dry-run
```

If you omit `--version`, the installer asks GitHub for the latest release tag.
After the installer gate is cleared, the installed CLI should only be checked
with help or version commands:

```bash
~/.local/bin/ni --help
~/.local/bin/ni version
```

## Manual Verification Path

Manual install keeps every trust step visible.

1. Open the release page:
   <https://github.com/Nam-Cheol/ni/releases>
2. Pick the archive for your platform:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

3. Download the archive and `ni_<version>_checksums.txt` from the same release.
4. Verify the archive checksum.

Linux:

```bash
grep ' ni_0.3.0_linux_amd64.tar.gz$' ni_0.3.0_checksums.txt | sha256sum -c -
```

macOS:

```bash
grep ' ni_0.3.0_darwin_arm64.tar.gz$' ni_0.3.0_checksums.txt | shasum -a 256 -c -
```

5. Extract and install:

```bash
mkdir -p "$HOME/.local/bin"
tar -xzf ni_0.3.0_darwin_arm64.tar.gz
install -m 0755 ni "$HOME/.local/bin/ni"
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

Use the matching archive name for your platform. On Windows, expand the `.zip`
archive with PowerShell or another trusted unzip tool and place `ni.exe` on your
`PATH`.

## What The Script Does

- Detects `linux`, `darwin`, or Windows-compatible shells.
- Detects `amd64` or `arm64`.
- Selects the GoReleaser asset named
  `ni_<version>_<os>_<arch>.tar.gz` or `ni_<version>_windows_amd64.zip`.
- Downloads from GitHub Releases by default.
- Downloads `ni_<version>_checksums.txt` and verifies the archive if possible.
- Installs to `~/.local/bin`, unless `BINDIR` is set.
- Prints help/version commands as next steps.

It does not run `ni init`, `ni status`, `ni end`, `ni run`, shell commands,
agents, queues, or runtime execution.

## Test Release Validation

Repository validation uses a local fake release asset so the installer can be
tested without network access and without Go. This is not the same as real
GitHub Release verification:

```bash
bash scripts/test-install-sh.sh
```
