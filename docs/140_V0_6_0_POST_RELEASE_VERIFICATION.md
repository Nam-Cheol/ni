# v0.6.0 Post-Release Verification

## Decision

V0_6_0_RELEASE_EXECUTED_AND_VERIFIED

## Release summary

| Item | Result |
| --- | --- |
| Release commit | `dfcd512a6cf6c08d8fba42d7e3d1c4cb9a84030b` |
| Tag | `v0.6.0` |
| GitHub Release | https://github.com/Nam-Cheol/ni/releases/tag/v0.6.0 |
| Draft | `false` |
| Prerelease | `false` |
| Workflow | Release run `27126268646` succeeded |

The release was created by the repository's existing tag-triggered GitHub
Actions workflow and GoReleaser. No manual release workflow was invented.

## Assets

| Asset | SHA-256 verification |
| --- | --- |
| `namba-intent_0.6.0_checksums.txt` | downloaded |
| `namba-intent_0.6.0_darwin_amd64.tar.gz` | OK |
| `namba-intent_0.6.0_darwin_arm64.tar.gz` | OK |
| `namba-intent_0.6.0_linux_amd64.tar.gz` | OK |
| `namba-intent_0.6.0_linux_arm64.tar.gz` | OK |
| `namba-intent_0.6.0_windows_amd64.zip` | OK |

Assets were downloaded to `/tmp/namba-intent-v0.6.0-verify` and verified with
`shasum -a 256 -c namba-intent_0.6.0_checksums.txt`.

## Current-platform artifact proof

Host platform: `darwin/arm64`.

Archive tested:
`namba-intent_0.6.0_darwin_arm64.tar.gz`.

The archive contained:

- `LICENSE`
- `README.ko.md`
- `README.md`
- `namba-intent`
- `ni`

Verification:

| Command | Result |
| --- | --- |
| extracted `namba-intent --help` | passed; identifies Namba Intent and prompt-compiler-only boundary |
| extracted `namba-intent version` | `0.6.0` |
| extracted `namba-intent init . --yes` in temp project | created `docs/plan/**` and `.ni/**` planning skeleton |
| extracted `namba-intent status --proof --next-questions` | `NI Intent Readiness: BLOCKED` with first-run blockers and `Execution must not start.` |

## Public installer proof

Installer command:

```bash
BINDIR=/tmp/namba-intent-v0.6.0-install-bin.HTnwfm bash install.sh --version 0.6.0
```

Observed:

- selected repository `Nam-Cheol/ni`
- selected platform `darwin/arm64`
- selected asset `namba-intent_0.6.0_darwin_arm64.tar.gz`
- verified checksum for the selected asset
- installed `/tmp/namba-intent-v0.6.0-install-bin.HTnwfm/namba-intent`

Temporary PATH verification:

| Command | Result |
| --- | --- |
| `PATH=/tmp/namba-intent-v0.6.0-install-bin.HTnwfm:$PATH namba-intent --help` | passed |
| `PATH=/tmp/namba-intent-v0.6.0-install-bin.HTnwfm:$PATH namba-intent version` | `0.6.0` |

Uninstall removed the temporary installed binary. Follow-up checks confirmed
the temporary `BINDIR` no longer existed and `command -v namba-intent` did not
resolve from that temp path.

## Validation Before Tag

| Gate | Result |
| --- | --- |
| `gofmt -w .` | passed; no tracked diff remained |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | passed |
| `go run ./cmd/namba-intent --help` | passed |
| `go run ./cmd/namba-intent version` | `0.0.0-dev` source build |
| `go run ./cmd/ni --help` | passed; warning contained `ni is deprecated; use namba-intent.` |
| temp PATH `namba-intent --help` | passed |
| temp PATH `namba-intent version` | `0.6.0` with linker injection |
| temp PATH `namba-intent init . --yes` | passed |
| temp PATH `namba-intent status --proof --next-questions` | blocked as expected for blank first-run workspace |
| `python3 scripts/check-install-docs.py` | passed |
| `python3 scripts/check-install-ps1.py` | passed |
| `bash scripts/check-skill-packs.sh` | passed |
| `bash scripts/demo-check.sh` | passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | passed |
| `bash scripts/smoke.sh` | passed |
| `bash scripts/install-check.sh` | passed |
| `bash scripts/release-check.sh` | passed after clearing a stale local Go build cache and rerunning outside the sandbox so Go could recreate `/Users/namba/Library/Caches/go-build` |
| protected `.ni` diff | empty |

## Protected State

The protected root planning files were not edited:

```bash
git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
```

Result: empty before tag and after release verification.

## Known Deferrals

- Homebrew remains Planned / v0.5 candidate. Do not mark Homebrew Available.
- Windows real-host verification remains pending until a Windows transcript
  exists.
- Repository remains `Nam-Cheol/ni`.
- `.ni/` compatibility is preserved.
- `run` compiles a bounded handoff prompt and does not execute downstream work.
- Model workspace packs remain Experimental.
- No-terminal method remains Experimental / assisted.

## Next Task

Prepare the next maintenance task for post-v0.6.0 cleanup: update any remaining
pre-release roadmap prompts that should now point to docs/140, then decide
whether v0.6.1 needs a small docs/checker polish patch or whether the next
work should stay on external verification such as Windows real-host proof.
