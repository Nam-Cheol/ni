# ni v0.4.0 Release Preflight

Date: 2026-05-29

Release: `ni v0.4.0 - Conversation Authoring Hardening`

Status: preflight record only. 이 문서는 tag를 만들거나 push하지 않고, GitHub
Release를 publish하지 않고, assets를 upload하지 않고, Homebrew를 Available로
표시하지 않고, runtime execution behavior를 추가하지 않는다.

## Scope Confirmation

`v0.4.0`은 `v0.3.0` 이후 adoption hardening을 다룬다:

- `R014`, `R015`, `R016` first-run proof text;
- first-run `ni-start` conversation card;
- `SYNC-014`, `SYNC-015`, `SYNC-016` docs/contract sync diagnostics;
- grouped and prioritized next questions;
- `ni-start`와 model workspace packs의 grouped-question usage;
- example coverage와 유지되는 Korean companion docs;
- no-terminal assisted workflow hardening;
- benchmark case study expansion;
- Homebrew deferred to `v0.5`;
- conversation proof capture;
- v0.4.0 release planning.

포함하지 않는 것:

- task runner;
- SPEC runner;
- execution harness;
- Codex exec adapter;
- shell adapter;
- downstream agents;
- queue;
- PR automation;
- `ni-kernel` 내부 release automation;
- Homebrew availability.

runtime execution behavior 없음. `ni run`은 bounded prompt compiler로 남아야
하며 shells, agents, queues, adapters, downstream work를 실행하면 안 된다.

## Git State

Preflight에서 확인한 상태:

| Check | Result |
| --- | --- |
| Branch | `main` |
| Working tree | preflight edits 전 clean |
| `HEAD == origin/main` | preflight edits 전 yes: `6f80074eb09fbcc5e81b64abd98f172ea69c7e66` |
| Recent tags | `v0.3.0`, `v0.1.0` |
| `v0.4.0` tag absent before release | yes |

이 preflight task는 이 문서, Korean companion, release-check update를 추가할 수
있다. Tag를 만들기 전 최종 working tree를 다시 확인해야 한다.

## Version State

Expected release version: `0.4.0`.

Checked-in development version source는 `internal/version/version.go`이며,
`Version` default는 `0.0.0-dev`다. `cmd/ni/main.go`는 `ni version`에서 이 값을
출력한다.

GoReleaser는 `.goreleaser.yaml`의 다음 ldflags로 release value를 inject한다:

```text
-s -w -X ni/internal/version.Version={{ .Version }}
```

따라서 repository는 `internal/version/version.go`에 `0.4.0`을 hard-code하면 안
된다. `v0.4.0` tag에서 GoReleaser가 만든 binary release는 `ni version`에서
binary release는 `0.4.0`을 report해야 한다.

## Release Workflow

`.github/workflows/release.yml`은 다음 tag push에서 실행된다:

```yaml
push:
  tags:
    - "v*"
```

Workflow는 full history로 checkout하고, `go.mod` 기준으로 Go를 설치한 뒤 다음을
실행한다:

- `go test ./...`;
- `bash scripts/quality.sh`;
- `bash scripts/release-check.sh`;
- `goreleaser/goreleaser-action@v6` with `release --clean`.

따라서 `v0.4.0` tag push에서 release workflow가 실행된다. Manual release
operator는 생성된 assets와 checksums를 별도로 검증해야 한다.

## Validation Commands

Tagging 전에 실행한다:

- `go test ./...`
- `bash scripts/quality.sh`
- `bash scripts/smoke.sh`
- `bash scripts/demo-check.sh`
- `bash scripts/install-check.sh`
- `bash scripts/release-check.sh`
- `bash scripts/fresh-install-check.sh` if present
- `bash scripts/check-skill-packs.sh`
- `bash scripts/package-claude-skills.sh`
- `bash scripts/package-codex-skills.sh`
- `bash scripts/release-dry-run.sh`

GoReleaser local dry run:

- GoReleaser가 설치되어 있으면 `goreleaser check`와
  `goreleaser release --snapshot --clean`을 실행한다.
- GoReleaser가 없으면 local GoReleaser portion은 skipped로 기록하고 GitHub
  Actions 또는 GoReleaser가 설치된 다른 machine에서 보완해야 한다.

## Validation Results

Current-task results:

| Command | Result |
| --- | --- |
| `go test ./...` | passed |
| `bash scripts/quality.sh` | passed |
| `bash scripts/smoke.sh` | passed |
| `bash scripts/demo-check.sh` | passed |
| `bash scripts/install-check.sh` | passed |
| `bash scripts/release-check.sh` | passed |
| `bash scripts/fresh-install-check.sh` | network access로 passed; `v0.4.0` assets는 release 전 존재하지 않으므로 현재 published `v0.3.0` release assets를 검증 |
| `bash scripts/check-skill-packs.sh` | passed |
| `bash scripts/package-claude-skills.sh` | passed; `dist/ni-claude-skills.zip` 생성 |
| `bash scripts/package-codex-skills.sh` | passed; `dist/ni-codex-skills.zip` 생성 |
| `bash scripts/release-dry-run.sh` | tags 생성, push, publishing 없이 passed |
| `goreleaser check` | skipped; local machine에 `goreleaser`가 설치되어 있지 않음 |
| `goreleaser release --snapshot --clean` | skipped; local machine에 `goreleaser`가 설치되어 있지 않음 |

Local GoReleaser config validation은 `goreleaser` binary가 이 machine에 없기
때문에 `bash scripts/release-check.sh`의 repository checks로 제한된다.
Release assets를 publish하기 전에 GitHub Actions 또는 GoReleaser가 설치된 다른
machine에서 GoReleaser check와 snapshot archive build를 보완해야 한다.

## Manual Release Steps

Preflight가 통과한 뒤에만 실행한다:

```bash
git status --short --branch
git fetch origin --tags
git tag --list v0.4.0
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/demo-check.sh
bash scripts/install-check.sh
bash scripts/release-check.sh
bash scripts/fresh-install-check.sh
bash scripts/check-skill-packs.sh
bash scripts/package-claude-skills.sh
bash scripts/package-codex-skills.sh
bash scripts/release-dry-run.sh
git tag -a v0.4.0 -m "ni v0.4.0"
git push origin v0.4.0
```

Tag push 이후:

1. GitHub Actions release workflow가 끝날 때까지 기다린다.
2. Release assets와 `ni_0.4.0_checksums.txt`를 verify한다.
3. Current-platform binary를 `ni --help`와 `ni version`으로 verify한다.
4. Curl installer를 `--version 0.4.0`으로 verify한다.
5. Availability status가 verification 후 바뀌는 경우에만 install docs를 update한다.

## Guardrails

- Preflight 중 tag creation 없음.
- Preflight 중 tag push 없음.
- Preflight 중 GitHub Release publication 없음.
- Homebrew availability claim 없음.
- runtime execution behavior 없음.
- shell adapter 없음.
- Codex exec adapter 없음.
- downstream agents 없음.
- queues 없음.
- user-facing contract `add`, `list`, `set` commands 없음.
- `ni end` 실행 없음.
- `ni relock` 실행 없음.
- `.ni/plan.lock.json` manual edits 없음.
