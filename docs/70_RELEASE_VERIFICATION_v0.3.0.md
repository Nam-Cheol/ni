# v0.3.0 Release Verification

Date: 2026-05-28

Scope: Verify the published GitHub Release assets for `v0.3.0` before marking
release binaries as available. This verification does not mark the curl
installer, Homebrew, Scoop, or any package-manager distribution as available.

Repository release:
`https://github.com/Nam-Cheol/ni/releases/tag/v0.3.0`

Release metadata:

- Tag: `v0.3.0`
- Name: `v0.3.0`
- Draft: `false`
- Prerelease: `false`
- Published: `2026-05-27T15:20:01Z`
- Target: `main`

## Asset List

The published release contains the expected OS/arch assets:

| Asset | Size | Digest |
| --- | ---: | --- |
| `ni_0.3.0_darwin_amd64.tar.gz` | 1,166,535 | `sha256:b6d65b177f0a58e7c9457fc562494e8d6dfdc92655aa0b1bb4aa697a8da952e0` |
| `ni_0.3.0_darwin_arm64.tar.gz` | 1,101,915 | `sha256:a41a45afb0e1f11779b28d70f397430773d7ad5f23252771077cc8fafefe0f33` |
| `ni_0.3.0_linux_amd64.tar.gz` | 1,148,197 | `sha256:7032a70dbe8e3824b10c6fa83e315507d8d135c89fe1cf0cc1597ebab19896e9` |
| `ni_0.3.0_linux_arm64.tar.gz` | 1,063,788 | `sha256:e7401a78465f2401c1948a05c2a4c646dfc9e6f0be834e8f0b888a466e3b20f9` |
| `ni_0.3.0_windows_amd64.zip` | 1,206,050 | `sha256:068d3a9ad0a857bf773f4c522f1e1803cc3e11f0d0b49bbef71d3b183f1e1267` |
| `ni_0.3.0_checksums.txt` | 471 | `sha256:b961642164db1b751e62bda0c5d489e23f901c0c2838b5206611dc9fa1557f44` |

Expected platform coverage is present:

- Linux amd64
- Linux arm64
- macOS amd64
- macOS arm64
- Windows amd64
- Checksums file

## Download

Assets were downloaded into a temporary local directory:

```text
/private/tmp/ni-v0.3.0-assets.9wZywv
```

Downloaded files:

```text
ni_0.3.0_checksums.txt
ni_0.3.0_darwin_amd64.tar.gz
ni_0.3.0_darwin_arm64.tar.gz
ni_0.3.0_linux_amd64.tar.gz
ni_0.3.0_linux_arm64.tar.gz
ni_0.3.0_windows_amd64.zip
```

## Checksum Verification

Command:

```bash
shasum -a 256 -c ni_0.3.0_checksums.txt
```

Output:

```text
ni_0.3.0_darwin_amd64.tar.gz: OK
ni_0.3.0_darwin_arm64.tar.gz: OK
ni_0.3.0_linux_amd64.tar.gz: OK
ni_0.3.0_linux_arm64.tar.gz: OK
ni_0.3.0_windows_amd64.zip: OK
```

## Archive Verification

Each archive extracted successfully.

Archive contents:

```text
darwin amd64:  LICENSE, README.md, README.ko.md, ni
darwin arm64:  LICENSE, README.md, README.ko.md, ni
linux amd64:   LICENSE, README.md, README.ko.md, ni
linux arm64:   LICENSE, README.md, README.ko.md, ni
windows amd64: LICENSE, README.md, README.ko.md, ni.exe
```

Binary format checks:

```text
extract/darwin_amd64/ni:      Mach-O 64-bit executable x86_64
extract/darwin_arm64/ni:      Mach-O 64-bit executable arm64
extract/linux_amd64/ni:       ELF 64-bit LSB executable, x86-64, statically linked
extract/linux_arm64/ni:       ELF 64-bit LSB executable, ARM aarch64, statically linked
extract/windows_amd64/ni.exe: PE32+ executable (console) x86-64, for MS Windows
```

## Current Platform Binary

Current verification platform:

```text
Darwin arm64
```

Current platform asset:

```text
ni_0.3.0_darwin_arm64.tar.gz
```

`ni --help` ran successfully and printed the CLI help.

`ni version` output:

```text
0.3.0
```

## Result

The published `v0.3.0` release assets are usable for the expected OS/arch
matrix, checksums verify, all archives extract, and the current platform binary
runs successfully.

Release binaries can now be marked as available.

The curl installer, Homebrew, Scoop, and package-manager distribution remain
not available until separately verified.
