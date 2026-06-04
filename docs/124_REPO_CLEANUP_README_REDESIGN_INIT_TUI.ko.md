# Repo Cleanup, README Redesign, and Init TUI

## Current status

- v0.5.0 publication: verified.
- README two-path onboarding: revised in this task.
- ni init . Bubble Tea TUI: implemented in this task.
- Windows real-host execution: macOS-only development host에서는 deferred.
- Homebrew: Planned / v0.5 candidate.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Task goal

이번 pass는 accidental repo files를 정리하고, README를 modern CLI product README로
단순화하며, `ni init .`을 Bubble Tea v2 / Lip Gloss v2 TUI로 upgrade한다.
Non-interactive fallback은 유지한다.

## Repository cleanup

| File or path | Classification | Action | Reason | Notes |
| --- | --- | --- | --- | --- |
| `examples/namba-ai-upgrade/docs/plan/* 2.md` | delete | 12 files 삭제 | Finder-style duplicate이며 더 오래된 TODO template content이고 참조가 없다. | ` 2` suffix가 없는 canonical files 유지. |
| `README 2.md`, `README.ko 2.md` | keep absent | No action | Existing validator가 root duplicate names를 이미 검사한다. | No files present. |
| `install 2.sh`, `install 2.ps1` | keep absent | No action | 요청된 suspicious installer copies는 존재하지 않았다. | Docs/scripts reference 없음. |
| `*.tmp`, `*.bak`, `Untitled*` | keep absent | No action | Matching local-only temp artifacts가 없었다. | Search returned no files. |
| Release archives/assets | document as intentional | No action | Release and packaging files는 기존 documented distribution work의 일부다. | Publish/tag/upload 없음. |

## README redesign

| Section | Before | After | Notes |
| --- | --- | --- | --- |
| Hero | 여러 card image group이 있는 긴 visual story. | Hero, one-line definition, 짧은 product 설명, 하나의 flow image. | Existing safe assets만 사용. |
| Install | 정확하지만 주변 copy가 길었다. | macOS/Windows primary paths와 detailed-doc link만 노출. | Homebrew는 Planned / v0.5 candidate. |
| First project | 긴 narrative 안에 있었다. | `mkdir`, `cd`, `ni init .`, `ni status`, `ni end`, `ni run` flow. | `ni run`은 downstream work를 실행하지 않는다고 명시. |
| Feature list | why/payoff sections에 분산. | `What ni does` command table. | CLI authority 유지. |
| Non-goals | 하단에 존재. | 짧은 `What ni does not do` section. | Runtime execution boundary 유지. |
| Status | Inline maturity note. | Concise status list. | Homebrew/Windows overclaim 없음. |

## README inspiration

README structure는 uv, Deno, Bun, GitHub CLI, Starship 같은 modern CLI projects의
간결한 구조에서 inspiration을 받았다. Content, branding, wording, layout은 복사하지
않았다.

## Init TUI architecture

| Layer | Responsibility | Package/path | Notes |
| --- | --- | --- | --- |
| command layer | Flags parse, TUI/fallback 선택, post-TUI summary 출력. | `cmd/ni/main.go` | `ni status`, `ni end`, `ni run`을 실행하지 않는다. |
| domain init logic | Init options validate, file plan build, lockfile protect, missing files write. | `internal/core/docstore` | TUI가 file writing을 소유하지 않는다. |
| TUI model | Wizard state, key handling, confirmation/cancel result. | `internal/tui/init` | Model, Init, Update, View 구현. |
| TUI styling | Lip Gloss v2 styles로 wizard render. | `internal/tui/init/model.go` | Visual styling은 writes와 분리. |
| file writer | Missing template files만 create하고 existing files는 skip. | `internal/core/docstore` | `.ni/plan.lock.json`은 수정하지 않는다. |
| summary printer | TUI 종료 후 plain text created/skipped/unchanged/next commands 출력. | `cmd/ni/main.go` | Summary는 readiness claim이 아니다. |

## Bubble Tea v2 / Lip Gloss v2 usage

| Requirement | Implemented path | Pass? | Notes |
| --- | --- | --- | --- |
| Bubble Tea v2 import | `internal/tui/init/model.go` | Yes | `charm.land/bubbletea/v2` 사용. |
| Lip Gloss v2 import | `internal/tui/init/model.go` | Yes | `charm.land/lipgloss/v2` 사용. |
| model / Init / Update / View | `internal/tui/init/model.go` | Yes | `View()`는 `tea.View` 반환. |
| AltScreen | `internal/tui/init/model.go` | Yes | `view.AltScreen = true`로 선언. |
| keyboard navigation | `internal/tui/init/model.go`, tests | Yes | Up, Down, Left, Right, Enter, Esc, q 처리. |
| post-TUI plain text summary | `cmd/ni/main.go` | Yes | TUI 종료 후 출력. |
| non-interactive fallback | `cmd/ni/main.go`, tests | Yes | Existing CI/non-TTY behavior 유지. |
| domain/render separation | `internal/core/docstore`, `internal/tui/init` | Yes | View는 render만 하고 domain이 write. |

## TUI behavior

| Behavior | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `ni init .` | Interactive TTY에서 guided TUI launch. | Command layer가 TTY positional init을 TUI로 route. | Yes | Non-TTY는 fallback. |
| current directory target | Target `.`이면 current directory 사용. | Existing CLI tests가 `ni init . --yes` cover. | Yes | TUI도 같은 dir parsing 사용. |
| interactive TUI | Step-based wizard와 write 전 review. | Field stage와 confirm stage 존재. | Yes | `View`에서 file write 없음. |
| non-interactive fallback | CI와 pipes에서 작동. | Existing fallback 유지. | Yes | Plain output only. |
| existing file protection | Silent overwrite 금지. | Domain plan이 existing files를 skipped로 mark. | Yes | TUI는 missing-only, keep, abort 제공. |
| lockfile protection | `.ni/plan.lock.json` 수정 금지. | Domain과 command layer가 write 없이 return. | Yes | Tested. |
| cancel | Confirmation 전 write 없음. | TUI가 canceled result 반환. | Yes | Model-level test. |
| confirm/write | Confirm result를 domain writer가 사용. | TUI는 intent만 반환하고 command가 docstore writer 호출. | Yes | Model/domain layers에서 test. |
| plain summary | target, created, skipped/unchanged, next commands 출력. | Command layer가 write/cancel 뒤 summary 출력. | Yes | TUI styling 아님. |

## Tests added

- `TestBuildFilePlanClassifiesCreateAndExisting`
- `TestInitWithOptionsProtectsLockfile`
- `TestModelInitialStateUsesAltScreen`
- `TestUpdateHandlesUpDownAndLeftRight`
- `TestUpdateHandlesEnterEscAndQ`
- `TestConfirmPathReturnsIntent`
- `TestCancelPathWritesNothingSignal`
- `TestExistingFileChoices`

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Homebrew | Planned / v0.5 candidate only. | README와 docs가 Planned wording을 보존. | Yes | Available claim 없음. |
| Windows real-host execution | Transcript 전까지 deferred. | README가 deferred wording 보존. | Yes | Windows verified claim 없음. |
| `ni init .` | Guided setup only. | TUI는 planning artifacts만 initialize. | Yes | Readiness 결정 없음. |
| `ni run` | Bounded prompt compilation only. | README와 code가 non-execution boundary 보존. | Yes | Runtime execution 없음. |
| READY | CLI status only. | README가 `ni status` decides readiness라고 명시. | Yes | TUI READY claim 없음. |
| Model workspace packs | Experimental. | Status list가 Experimental 보존. | Yes | Skills are UX; CLI is authority. |
| No-terminal | Experimental / assisted. | Status list가 boundary 보존. | Yes | Deterministic validation 아님. |
| Runtime execution boundary | Kernel은 pre-runtime. | Shell/Codex adapter 또는 queue 추가 없음. | Yes | Downstream execution behavior 없음. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| `README.md` | Modified | Yes | Public README redesign. |
| `README.ko.md` | Modified | Yes | Companion update. |
| `cmd/ni/*` | Modified | Yes | Init routing과 summary printer. |
| `internal/*` | Modified/added | Yes | Domain file plan과 TUI package. |
| `go.mod` | Modified | Yes | Bubble Tea v2와 Lip Gloss v2 dependencies. |
| `go.sum` | Added | Yes | Module checksums. |
| `docs/22_INSTALL.md` | Unchanged | No | Existing markers accurate. |
| `docs/install-curl*` | Unchanged | No | Existing markers accurate. |
| `docs/124*` | Added | Yes | Implementation report와 Korean companion. |
| install scripts | Unchanged | No | Release/publish work 없음. |
| suspicious files removed | Deleted | Yes | `* 2.md` duplicates removed. |
| `.ni/contract.json` | Unchanged | No | Protected. |
| `.ni/session.json` | Unchanged | No | Protected. |
| `.ni/plan.lock.json` | Unchanged | No | Protected. |
| unexpected files | None expected | No | Commit 전 recheck. |

## Validation results

| Command | Result |
| --- | --- |
| `go get charm.land/bubbletea/v2 charm.land/lipgloss/v2` | Passed after network approval. |
| `GOCACHE=/private/tmp/ni-go-cache go mod tidy` | Passed after network approval. |
| `gofmt -w ...` | Passed. |
| `gofmt -w .` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; `NI Intent Readiness: READY`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build는 `0.0.0-dev` 반환. |
| Non-interactive `ni init --dir /private/tmp/ni-init-smoke-203-a --yes` | Passed; planning files 생성, lockfile/downstream artifact 생성 없음. |
| 같은 temp directory에서 repeated non-interactive init | Passed; existing files는 unchanged로 report. |
| Lockfile safety temp init | Passed; `.ni/plan.lock.json` 수정 없음, contract 생성 없음. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `python3 scripts/check-readme-surface.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

## Changes made

- `README.md`와 `README.ko.md`를 redesign했다.
- `internal/tui/init`에 Bubble Tea v2 / Lip Gloss v2 TUI package를 추가했다.
- `cmd/ni/main.go`가 interactive init을 TUI로 route하고 non-interactive fallback을 유지하도록 update했다.
- `internal/core/docstore`에 domain file-plan과 lockfile protection checks를 추가했다.
- Init behavior에 대한 model/domain tests를 추가했다.
- Accidental `examples/namba-ai-upgrade/docs/plan/* 2.md` duplicates를 삭제했다.
- 이 문서와 English companion을 추가했다.

## What this task proves

- Accidental repo files were audited and cleaned where safe.
- README now has a simpler modern CLI structure.
- `ni init .` uses Bubble Tea v2 / Lip Gloss v2 for interactive onboarding.
- Non-interactive fallback remains available.
- Domain init logic and TUI rendering are separated.
- `ni run` remains bounded prompt compilation only.

## What this task does not prove

- Windows real-host execution works unless a Windows transcript exists.
- Homebrew is Available.
- Downstream execution succeeds.
- No-terminal is deterministic.
- External users succeed.
- README design is final.

## Recommended next task

A. first-user onboarding smoke after TUI

## Next task prompt

Proceed in `/Users/namba/Documents/project/ni`.

Goal: run a first-user onboarding smoke after the Bubble Tea init TUI pass.

Rules:
- Do not publish, tag, release, upload assets, run root `ni end`, relock the
  project root, or execute generated prompts.
- Do not mark Homebrew Available.
- Do not claim Windows real-host execution verified unless a Windows transcript
  exists.
- Preserve `Skills are UX; CLI is authority.`

Scope:
- Build a temporary `ni` binary.
- In a temporary project directory, run non-interactive `ni init . --yes`,
  repeated `ni init . --yes`, lockfile protection, `ni status --proof
  --next-questions`, and help/version checks.
- If practical on an interactive TTY, manually smoke `ni init .` TUI and record
  only what was observed; otherwise state that interactive TTY smoke remains
  deferred.
- Update the smallest appropriate docs with smoke results and boundaries.

Validation:
- `gofmt -w .`
- `GOCACHE=/private/tmp/ni-go-cache go test ./...`
- `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help`
- `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version`
- `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh`
- `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json`

Final response: report changed files, smoke commands, results, remaining
interactive/host limitations, and protected `.ni` diff status.
