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

Release는 repository의 기존 tag-triggered GitHub Actions workflow와 GoReleaser로
생성했다. 별도 manual release workflow는 만들지 않았다.

## Assets

| Asset | SHA-256 verification |
| --- | --- |
| `namba-intent_0.6.0_checksums.txt` | downloaded |
| `namba-intent_0.6.0_darwin_amd64.tar.gz` | OK |
| `namba-intent_0.6.0_darwin_arm64.tar.gz` | OK |
| `namba-intent_0.6.0_linux_amd64.tar.gz` | OK |
| `namba-intent_0.6.0_linux_arm64.tar.gz` | OK |
| `namba-intent_0.6.0_windows_amd64.zip` | OK |

Assets는 `/tmp/namba-intent-v0.6.0-verify`에 download했고
`shasum -a 256 -c namba-intent_0.6.0_checksums.txt`로 검증했다.

## Current-platform artifact proof

Host platform: `darwin/arm64`.

Tested archive:
`namba-intent_0.6.0_darwin_arm64.tar.gz`.

Archive contents:

- `LICENSE`
- `README.ko.md`
- `README.md`
- `namba-intent`
- `ni`

Verification:

| Command | Result |
| --- | --- |
| extracted `namba-intent --help` | passed; Namba Intent와 prompt-compiler-only boundary 확인 |
| extracted `namba-intent version` | `0.6.0` |
| extracted `namba-intent init . --yes` in temp project | `docs/plan/**` and `.ni/**` planning skeleton 생성 |
| extracted `namba-intent status --proof --next-questions` | `NI Intent Readiness: BLOCKED`, first-run blockers, `Execution must not start.` |

## Public installer proof

Installer command:

```bash
BINDIR=/tmp/namba-intent-v0.6.0-install-bin.HTnwfm bash install.sh --version 0.6.0
```

Observed:

- selected repository `Nam-Cheol/ni`
- selected platform `darwin/arm64`
- selected asset `namba-intent_0.6.0_darwin_arm64.tar.gz`
- selected asset checksum verified
- installed `/tmp/namba-intent-v0.6.0-install-bin.HTnwfm/namba-intent`

Temporary PATH verification:

| Command | Result |
| --- | --- |
| `PATH=/tmp/namba-intent-v0.6.0-install-bin.HTnwfm:$PATH namba-intent --help` | passed |
| `PATH=/tmp/namba-intent-v0.6.0-install-bin.HTnwfm:$PATH namba-intent version` | `0.6.0` |

Uninstall은 temporary installed binary를 제거했다. Follow-up checks에서 temporary
`BINDIR`이 더 이상 존재하지 않고 `command -v namba-intent`가 temp path에서
resolve되지 않음을 확인했다.

## Validation Before Tag

| Gate | Result |
| --- | --- |
| `gofmt -w .` | passed; tracked diff 없음 |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | passed |
| `go run ./cmd/namba-intent --help` | passed |
| `go run ./cmd/namba-intent version` | `0.0.0-dev` source build |
| `go run ./cmd/ni --help` | passed; warning contained `ni is deprecated; use namba-intent.` |
| temp PATH `namba-intent --help` | passed |
| temp PATH `namba-intent version` | `0.6.0` with linker injection |
| temp PATH `namba-intent init . --yes` | passed |
| temp PATH `namba-intent status --proof --next-questions` | blank first-run workspace에서 expected BLOCKED |
| `python3 scripts/check-install-docs.py` | passed |
| `python3 scripts/check-install-ps1.py` | passed |
| `bash scripts/check-skill-packs.sh` | passed |
| `bash scripts/demo-check.sh` | passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | passed |
| `bash scripts/smoke.sh` | passed |
| `bash scripts/install-check.sh` | passed |
| `bash scripts/release-check.sh` | stale local Go build cache clear 후, Go가 `/Users/namba/Library/Caches/go-build`를 recreate할 수 있도록 sandbox 밖에서 rerun하여 passed |
| protected `.ni` diff | empty |

## Protected State

Protected root planning files는 edit하지 않았다.

```bash
git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
```

Result: tag 전과 release verification 후 모두 empty.

## Known Deferrals

- Homebrew remains Planned / v0.5 candidate. Homebrew Available로 표시하지 않는다.
- Windows real-host verification은 Windows transcript가 생길 때까지 pending.
- Repository remains `Nam-Cheol/ni`.
- `.ni/` compatibility is preserved.
- `run` compiles a bounded handoff prompt and does not execute downstream work.
- Model workspace packs remain Experimental.
- No-terminal method remains Experimental / assisted.

## Next Task

Post-v0.6.0 cleanup으로 남은 pre-release roadmap prompts가 docs/140을 가리키도록
정리한 뒤, v0.6.1 docs/checker polish patch가 필요한지 또는 Windows real-host
proof 같은 external verification을 다음 작업으로 둘지 결정한다.
