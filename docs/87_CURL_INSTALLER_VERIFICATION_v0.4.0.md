# v0.4.0 Curl Installer Verification

Date: 2026-05-29

Scope: Verify `install.sh` against the real `v0.4.0` GitHub Release assets
after the release assets and checksums were verified. This verification does
not publish a release, push tags, mark Homebrew available, add package-manager
claims, install model skills, or add runtime execution behavior.

Release under test:

`https://github.com/Nam-Cheol/ni/releases/tag/v0.4.0`

## Commands Run

Dry-run:

```bash
sh install.sh --dry-run --version v0.4.0
```

Real temporary install:

```bash
BINDIR="/var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy" sh install.sh --version v0.4.0
"/var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni" --help
"/var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni" version
```

Platform override dry-runs:

```bash
NI_INSTALL_OS=linux NI_INSTALL_ARCH=amd64 sh install.sh --dry-run --version v0.4.0
NI_INSTALL_OS=linux NI_INSTALL_ARCH=arm64 sh install.sh --dry-run --version v0.4.0
NI_INSTALL_OS=darwin NI_INSTALL_ARCH=amd64 sh install.sh --dry-run --version v0.4.0
NI_INSTALL_OS=darwin NI_INSTALL_ARCH=arm64 sh install.sh --dry-run --version v0.4.0
NI_INSTALL_OS=windows NI_INSTALL_ARCH=amd64 sh install.sh --dry-run --version v0.4.0
```

Checksum mismatch safety test:

```bash
bash scripts/test-install-sh.sh
```

The mismatch test uses a local fake release asset and intentionally bad local
checksum file. It does not mutate or spoof the published GitHub Release.

## Dry-Run Output

```text
ni installer
  repository: Nam-Cheol/ni
  platform:   darwin/arm64
  asset:      ni_0.4.0_darwin_arm64.tar.gz
  checksums:  ni_0.4.0_checksums.txt
  install to: /Users/namba/.local/bin/ni
  mode:       dry-run

Would download:
  https://github.com/Nam-Cheol/ni/releases/download/v0.4.0/ni_0.4.0_darwin_arm64.tar.gz
  https://github.com/Nam-Cheol/ni/releases/download/v0.4.0/ni_0.4.0_checksums.txt
```

## Platform Override Results

| Override | Selected asset |
| --- | --- |
| `linux/amd64` | `ni_0.4.0_linux_amd64.tar.gz` |
| `linux/arm64` | `ni_0.4.0_linux_arm64.tar.gz` |
| `darwin/amd64` | `ni_0.4.0_darwin_amd64.tar.gz` |
| `darwin/arm64` | `ni_0.4.0_darwin_arm64.tar.gz` |
| `windows/amd64` | `ni_0.4.0_windows_amd64.zip` |

Each override selected `ni_0.4.0_checksums.txt` from the same release URL.

## Real Temporary Install Output

```text
ni installer
  repository: Nam-Cheol/ni
  platform:   darwin/arm64
  asset:      ni_0.4.0_darwin_arm64.tar.gz
  checksums:  ni_0.4.0_checksums.txt
  install to: /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni
Downloading ni_0.4.0_darwin_arm64.tar.gz
Verified checksum for ni_0.4.0_darwin_arm64.tar.gz
Installed ni to /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni

Next steps:
  1. Ensure /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy is on your PATH.
  2. Check the installed CLI:
     /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni --help
     /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni version

The installer does not install model skills or run downstream work.
```

Installed binary checks:

```text
$ /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni --help
ni is a project intent compiler.

$ /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.1lOuglWASy/ni version
0.4.0
```

## Checksum Behavior

- The real install downloaded `ni_0.4.0_checksums.txt` from the v0.4.0 release.
- The installer found the `ni_0.4.0_darwin_arm64.tar.gz` checksum entry.
- The installer printed `Verified checksum for ni_0.4.0_darwin_arm64.tar.gz`.
- `bash scripts/test-install-sh.sh` confirmed checksum mismatch failure using a
  local fake release and bad checksum file.

## Result

Curl installer status: Available for the verified v0.4.0 release assets.

Homebrew status: Planned.

No package-manager availability is claimed by this verification.
