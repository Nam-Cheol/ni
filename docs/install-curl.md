# Curl Installer

`install.sh` is release asset infrastructure for installing a released Namba
Intent binary without requiring Go. Current main selects
`namba-intent_<version>` release assets, copies `namba-intent` into a local bin
directory, and prints next steps. With explicit opt-in, it can add a reversible
zsh/bash PATH block. It does not install model skills or run downstream work.

Status: Available for the verified v0.6.4 macOS `namba-intent` path. The
verified v0.5.1 GitHub Release assets use historical `ni_<version>` names and
remain documented in the v0.5.1 post-release verification record.

## Safer Script Path

For the verified v0.6.4 macOS install:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --update-path --version 0.6.4
```

When you want an inspect-first or reproducible pinned install, download the
installer before any local install:

```bash
VERSION="0.6.4"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

By default, the script installs to `~/.local/bin/namba-intent`. Override the
destination with `BINDIR`:

```bash
BINDIR="$HOME/bin" sh install.sh --dry-run --version "$VERSION"
```

If you omit `--version`, the installer asks GitHub for the latest release tag.
Use `--version 0.6.4` when you want to reproduce the verified Namba Intent
release path. Pin `VERSION="0.5.1"` only when using historical scripts or
evidence covered by
[v0.5.1 Post-Release Verification](132_V0_5_1_POST_RELEASE_VERIFICATION.md);
current-main `install.sh --version 0.5.1` selects `namba-intent_0.5.1...`
assets that were not published.
Current dry-run output does not resolve latest without `--version`, so the
dry-run example stays pinned instead of serving as the README primary path.
Open a new shell after installation, then check the global command with help or
version commands:

```bash
namba-intent --help
namba-intent version
```

After global command verification, start a planning workspace from the project
directory:

```bash
mkdir my-project
cd my-project
namba-intent init .
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
2. Confirm the release has assets, then pick the archive for your platform.
   For v0.6.0 and later Namba Intent releases, use `namba-intent_<version>`
   names. For historical v0.5.1 evidence, use `ni_<version>` names from the
   v0.5.1 release page.

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `namba-intent_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `namba-intent_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `namba-intent_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `namba-intent_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `namba-intent_<version>_windows_amd64.zip` |

3. Download the archive and `namba-intent_<version>_checksums.txt` from the
   same release.
4. Verify the archive checksum.

Linux:

```bash
VERSION="<published-version-without-v>"
grep " namba-intent_${VERSION}_linux_amd64.tar.gz$" "namba-intent_${VERSION}_checksums.txt" | sha256sum -c -
```

macOS:

```bash
VERSION="<published-version-without-v>"
grep " namba-intent_${VERSION}_darwin_arm64.tar.gz$" "namba-intent_${VERSION}_checksums.txt" | shasum -a 256 -c -
```

5. Extract and install:

```bash
mkdir -p "$HOME/.local/bin"
VERSION="<published-version-without-v>"
tar -xzf "namba-intent_${VERSION}_darwin_arm64.tar.gz"
install -m 0755 namba-intent "$HOME/.local/bin/namba-intent"
PATH="$HOME/.local/bin:$PATH" namba-intent --help
PATH="$HOME/.local/bin:$PATH" namba-intent version
```

Use the matching archive name for your platform. On Windows, expand the `.zip`
archive with PowerShell or another trusted unzip tool and place
`namba-intent.exe` on your User PATH, or use `install.ps1` from the repository
root. PowerShell alias cleanup for `ni -> New-Item` is legacy v0.5.x guidance
and is not required for `namba-intent.exe`. Use `Get-Command namba-intent -All`
if command-name resolution does not look right.

## What The Script Does

- Detects `linux`, `darwin`, or Windows-compatible shells.
- Detects `amd64` or `arm64`.
- Selects the GoReleaser asset named
  `namba-intent_<version>_<os>_<arch>.tar.gz` or
  `namba-intent_<version>_windows_amd64.zip`.
- Downloads from GitHub Releases by default.
- Downloads `namba-intent_<version>_checksums.txt` and verifies the archive if possible.
- Installs to `~/.local/bin`, unless `BINDIR` is set.
- Optionally adds a marked zsh/bash PATH block when `--update-path` is passed.
- Removes the installed binary and only the marked PATH block with
  `--uninstall`.
- Prints command-name help/version commands as next steps.

It does not run `namba-intent init`, `namba-intent status`, `namba-intent end`,
`namba-intent run`, shell commands, agents, queues, or runtime execution.

## Test Release Validation

Repository validation uses a local fake release asset so the installer can be
tested without network access and without Go:

```bash
bash scripts/test-install-sh.sh
```

The v0.6.4 macOS installer path has real release verification. For later
releases, repeat that verification before changing public availability language:

```bash
VERSION="0.6.4"
sh install.sh --dry-run --version "$VERSION"
BINDIR="$(mktemp -d)" sh install.sh --version "$VERSION"
```

The historical v0.5.1 verification passed on 2026-06-08 for the old `ni`
release assets. The v0.6.4 macOS verification is recorded in
[`143_V0_6_4_POST_RELEASE_VERIFICATION.md`](143_V0_6_4_POST_RELEASE_VERIFICATION.md).
Current tree command-name verification is covered by `bash scripts/install-check.sh`
with a temporary install directory and fresh shell PATH context.
