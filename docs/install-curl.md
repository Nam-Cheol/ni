# Curl Installer

`install.sh` is release asset infrastructure for installing a released `ni`
binary without requiring Go. It downloads an archive, verifies the checksum when
the release provides one, copies `ni` into a local bin directory, and prints next
steps. With explicit opt-in, it can add a reversible zsh/bash PATH block. It
does not install model skills or run downstream work.

Status: Available for the verified v0.5.1 GitHub Release assets. `install.sh`
downloads the selected archive and matching `ni_<version>_checksums.txt`,
verifies the archive when a local sha256 tool is available, installs only the
`ni` binary, and does not run downstream work.

## Safer Script Path

For latest-by-default install, omit `--version` during actual install:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --update-path
```

When you want an inspect-first or reproducible pinned install, download the
installer before any local install:

```bash
VERSION="0.5.1"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

By default, the script installs to `~/.local/bin/ni`. Override the destination
with `BINDIR`:

```bash
BINDIR="$HOME/bin" sh install.sh --dry-run --version "$VERSION"
```

If you omit `--version`, the installer asks GitHub for the latest release tag.
Pin `VERSION="0.5.1"` when you want the verified release covered by
[v0.5.1 Post-Release Verification](132_V0_5_1_POST_RELEASE_VERIFICATION.md).
Current dry-run output does not resolve latest without `--version`, so the
dry-run example stays pinned instead of serving as the README primary path.
Open a new shell after installation, then check the global command with help or
version commands:

```bash
ni --help
ni version
```

After global command verification, start a planning workspace from the project
directory:

```bash
mkdir my-project
cd my-project
ni init .
```

Uninstall the binary and the ni-managed PATH block, if one was added:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --uninstall
```

If you installed from a downloaded script or used a custom `BINDIR`, run
uninstall from that local script and pass the same `BINDIR`:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

## Manual Verification Path

Manual install keeps every trust step visible.

1. Open the release page:
   <https://github.com/Nam-Cheol/ni/releases>
2. Confirm the release has assets, then pick the archive for your platform:

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
VERSION="<published-version-without-v>"
grep " ni_${VERSION}_linux_amd64.tar.gz$" "ni_${VERSION}_checksums.txt" | sha256sum -c -
```

macOS:

```bash
VERSION="<published-version-without-v>"
grep " ni_${VERSION}_darwin_arm64.tar.gz$" "ni_${VERSION}_checksums.txt" | shasum -a 256 -c -
```

5. Extract and install:

```bash
mkdir -p "$HOME/.local/bin"
VERSION="<published-version-without-v>"
tar -xzf "ni_${VERSION}_darwin_arm64.tar.gz"
install -m 0755 ni "$HOME/.local/bin/ni"
PATH="$HOME/.local/bin:$PATH" ni --help
PATH="$HOME/.local/bin:$PATH" ni version
```

Use the matching archive name for your platform. On Windows, expand the `.zip`
archive with PowerShell or another trusted unzip tool and place `ni.exe` on your
User PATH, or use `install.ps1` from the repository root.

## What The Script Does

- Detects `linux`, `darwin`, or Windows-compatible shells.
- Detects `amd64` or `arm64`.
- Selects the GoReleaser asset named
  `ni_<version>_<os>_<arch>.tar.gz` or `ni_<version>_windows_amd64.zip`.
- Downloads from GitHub Releases by default.
- Downloads `ni_<version>_checksums.txt` and verifies the archive if possible.
- Installs to `~/.local/bin`, unless `BINDIR` is set.
- Optionally adds a marked zsh/bash PATH block when `--update-path` is passed.
- Removes the installed binary and only the marked PATH block with
  `--uninstall`.
- Prints command-name help/version commands as next steps.

It does not run `ni init`, `ni status`, `ni end`, `ni run`, shell commands,
agents, queues, or runtime execution.

## Test Release Validation

Repository validation uses a local fake release asset so the installer can be
tested without network access and without Go:

```bash
bash scripts/test-install-sh.sh
```

For future releases, repeat the real release verification before changing
public availability language:

```bash
VERSION="0.5.1"
sh install.sh --dry-run --version "$VERSION"
BINDIR="$(mktemp -d)" sh install.sh --version "$VERSION"
```

The v0.5.1 verification passed on 2026-06-08. The installer printed
`Verified checksum for ni_0.5.1_darwin_arm64.tar.gz`, installed the binary into
a temporary `BINDIR`, and the installed binary returned `0.5.1` for
`ni version`. Global command-name verification is now covered by
`bash scripts/install-check.sh` with a temporary install directory and fresh
shell PATH context.
