# v0.4.0 Release Verification

Date: 2026-05-29

Scope: Verify the published GitHub Release assets for `v0.4.0` after the
manual tag push and GitHub Actions release build. This release-asset
verification does not publish another release, push tags, change the release
version, mark Homebrew available, add package-manager availability claims, or
add runtime execution behavior.

Repository release:
`https://github.com/Nam-Cheol/ni/releases/tag/v0.4.0`

Release metadata:

- Tag: `v0.4.0`
- Name: `v0.4.0`
- Draft: `false`
- Prerelease: `false`
- Published: `2026-05-29T06:58:39Z`
- Target: `main`

## Asset List

The published release contains the expected OS/arch assets:

| Asset | Size | Digest |
| --- | ---: | --- |
| `ni_0.4.0_checksums.txt` | 471 | `sha256:b24746a824084b01ebcc5c706afe9e5e43ca8bdb2b8606a766edcaa1bb4a70ca` |
| `ni_0.4.0_darwin_amd64.tar.gz` | 1,209,092 | `sha256:b7b503cf998a963fc21174967617710bdb0d8efecd9d4202e6899f2e0e36a9a1` |
| `ni_0.4.0_darwin_arm64.tar.gz` | 1,149,041 | `sha256:da5e3b715a79e2a284b3095c2e2f5e813ecccc7b2ea89657e363af48d2603813` |
| `ni_0.4.0_linux_amd64.tar.gz` | 1,190,005 | `sha256:ca4733be1cc67417fbd51b2d6ef3866b6d3eb3021feaa1f049e39f5e6266dd6f` |
| `ni_0.4.0_linux_arm64.tar.gz` | 1,107,207 | `sha256:4e2bffb3505a11b6a0bd6ff2939994f7629aa715768a77a96357b2f062b22733` |
| `ni_0.4.0_windows_amd64.zip` | 1,249,396 | `sha256:95f9c09b815327e0040bd0263fe4db35949472e99bbc22d70f916789bb101f2f` |

Expected platform coverage is present:

- Checksums file
- macOS amd64
- macOS arm64
- Linux amd64
- Linux arm64
- Windows amd64

## Download

Assets were downloaded into a temporary local directory:

```text
/private/tmp/ni-v0.4.0-assets.leSaCK
```

Downloaded files:

```text
ni_0.4.0_checksums.txt
ni_0.4.0_darwin_amd64.tar.gz
ni_0.4.0_darwin_arm64.tar.gz
ni_0.4.0_linux_amd64.tar.gz
ni_0.4.0_linux_arm64.tar.gz
ni_0.4.0_windows_amd64.zip
```

## Checksum Verification

Command:

```bash
shasum -a 256 -c ni_0.4.0_checksums.txt
```

Output:

```text
ni_0.4.0_darwin_amd64.tar.gz: OK
ni_0.4.0_darwin_arm64.tar.gz: OK
ni_0.4.0_linux_amd64.tar.gz: OK
ni_0.4.0_linux_arm64.tar.gz: OK
ni_0.4.0_windows_amd64.zip: OK
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
ni_0.4.0_darwin_arm64.tar.gz
```

`ni --help` output:

```text
ni is a project intent compiler.

Usage:
  ni --help
  ni amend create --title <title> [--dir <path>]
  ni amend list [--dir <path>]
  ni amend show <id> [--dir <path>]
  ni amend apply <id> [--dir <path>]
  ni conflicts --base <path-or-lock> --head <path-or-lock> [--json]
  ni diff --base <path-or-lock> --head <path-or-lock> [--json]
  ni end --dir <path>
  ni export --target hyper-run|namba-ai|ouroboros|spec-kit --out <dir> [--dir <path>]
  ni feedback add --file <path> [--dir <path>]
  ni feedback list [--dir <path>] [--json]
  ni graph --dir <path> [--json]
  ni harness plan --dir <path> [--json]
  ni harness candidates [--dir <path>] [--json]
  ni harness propose --from-pressure <id> [--dir <path>]
  ni harness validate <candidate-id> --evidence <path> [--dir <path>]
  ni harness accept <candidate-id> [--dir <path>]
  ni harness retire <candidate-id> [--dir <path>]
  ni init --dir <path> [--profile concept|prototype|mvp|beta|production] [--product-type <type>] [--surface <surface>] [--interaction-mode <mode>]
  ni pressure status [--dir <path>] [--json]
  ni pressure promote <id> [--dir <path>]
  ni pressure retire <id> [--dir <path>]
  ni relock --dir <path>
  ni run --dir <path> [--target <target>] [--out <path>] [--max-chars N]
  ni status --dir <path> [--json] [--proof] [--next-questions]
  ni targets [--json]
  ni version

Commands:
  amend   Create, inspect, and apply explicit contract amendments.
  conflicts Detect semantic planning conflicts between two contracts or locked plans.
  diff     Show contract-level changes between two contracts or locked plans.
  end      Lock the accepted planning contract.
  export   Write locked-plan seed artifacts for a downstream target.
  feedback Record and list inert downstream feedback.
  graph    Propose a read-only work graph.
  harness  Manage inert generated harness proposals.
  init     Create planning docs and .ni skeleton.
  pressure Track inert planning pressure without changing readiness rules.
  relock   Create a new lock from an explicitly amended plan.
  run      Compile a goal prompt from the locked plan.
  status   Validate planning readiness.
  targets  List downstream prompt/export targets.
  version  Print the ni version.
```

`ni version` output:

```text
0.4.0
```

## Result

The published `v0.4.0` release assets are verified.

Checksums pass, all archives extract successfully, every non-current platform
archive has the expected binary naming, and the current platform binary runs
`ni --help` and `ni version`. The version output is `0.4.0`.

This verification does not mark Homebrew as available and does not add any
package-manager claims.
