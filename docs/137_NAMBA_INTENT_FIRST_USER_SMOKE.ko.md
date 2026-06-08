# Namba Intent First-User Smoke

## Current status

State:
- v0.5.1 release: published and verified
- v0.6.0 release: not published
- Namba Intent rename: implemented in current tree
- primary command: namba-intent
- deprecated ni shim: transition-only
- .ni/ compatibility: preserved
- Homebrew: Planned / v0.5 candidate
- Windows real-host verification: pending
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

## Smoke goal

이 smoke는 temporary PATH에서 resolve되는 current-tree `namba-intent` binary로
Namba Intent onboarding을 검증합니다. Public v0.6.0 install 또는 hosted release
artifact 검증이 아닙니다.

## Decision

NAMBA_INTENT_FIRST_USER_SMOKE_PASS_WITH_NOTES

Notes:
- Current-tree `namba-intent` command-name, first-user init, first-user status,
  repeated init, lockfile safety, deprecated `ni` shim checks가 통과했습니다.
- Blank first-user project는 purpose, actors/outcomes, delivery surface, first
  open blocker question이 unresolved라서 예상대로 `BLOCKED`를 보고했습니다.
- Public v0.6.0 install, Windows real-host execution, Homebrew availability,
  external user success, downstream execution은 증명하지 않았습니다.

## Command-name smoke

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| command-name resolution | `command -v namba-intent` | Temporary PATH에서 resolve되고, typed absolute path가 아니다. | `/tmp/namba-intent-first-user-smoke/bin/namba-intent` | Yes | Temporary current-tree binary only. |
| help | `namba-intent --help` | Help가 Namba Intent와 `namba-intent`를 primary로 사용한다. | Help는 "Namba Intent is a Project Intent Compiler for AI Agents."로 시작하고 usage line은 `namba-intent`를 사용한다. | Yes | Help는 run이 bounded prompt를 compile할 뿐 agents, shell commands, queues, PR automation, release automation을 실행하지 않는다고 말한다. |
| version | `namba-intent version` | Version command가 실행된다. | `0.0.0-dev` | Yes | Local development build에 맞는 값이며 release proof가 아니다. |

## First-user project smoke

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| initialize blank project | `namba-intent init . --yes` | Lockfile 또는 downstream execution 없이 planning workspace artifacts를 만든다. | `docs/plan/**`, `.ni/contract.json`, `.ni/session.json`, `.ni/project.json`, `.ni/pressure.json`, `.ni/harness.candidates.json`, `.ni/readiness.rules.json`, `.ni/readiness.profiles.json` 생성. | Yes | Output은 next command를 제안만 했고, status 외에는 실행하지 않았다. |
| status proof and questions | `namba-intent status --proof --next-questions` | 실제 planning state를 보고한다. Blank workspace는 `BLOCKED`일 수 있다. | `NI Intent Readiness: BLOCKED`, blockers는 R014, OQ-001, R015, R016. | Yes | Output은 "Execution must not start."를 포함했다. |

## Generated artifact check

| Artifact | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `.ni/contract.json` | Init 후 존재. | 존재. | Yes | Init fixture가 생성. |
| `.ni/session.json` | Init 후 존재. | 존재. | Yes | Init fixture가 생성. |
| `docs/plan/**` | Init 후 존재. | `00_project_brief.md`부터 `11_decision_log.md`까지 12개 planning docs 존재. | Yes | Planning skeleton only. |
| `.ni/plan.lock.json` | First-user init/status 후 없어야 한다. | 없음. | Yes | Lock 생성 없음. |
| downstream generated prompt | Init/status가 생성하거나 실행하지 않는다. | Fixture에 `.ni/generated/**` 없음. | Yes | `run` 실행 없음. |

## Repeated init safety

`/tmp/namba-intent-first-user-smoke/repeat.60uuYh`에서 첫
`namba-intent init . --yes`는 planning skeleton을 만들었습니다. 두 번째
`namba-intent init . --yes`는 다음을 보고했습니다:

- `existing Namba Intent planning files found; namba-intent init will not overwrite them.`
- `adding missing files only.`
- 기존 planning docs와 `.ni` files를 `exists`와 `unchanged`로 보고
- `created files: none`

`.ni/plan.lock.json`은 생성되지 않았고 downstream prompt도 실행되지 않았습니다.

## Lockfile safety

`/tmp/namba-intent-first-user-smoke/lock.0JtvKg`에서 init으로 blank planning
workspace를 만든 뒤, temporary project 안에만 fixture `.ni/plan.lock.json`을
배치했습니다. 두 번째 init 전 SHA-256은 다음과 같습니다:

```text
61d7265d58137893995f2a693b83d213874ab5a2028ecb693857350d0eb7c16c
```

두 번째 `namba-intent init . --yes`는 다음을 보고했습니다:

```text
warning: .ni/plan.lock.json already exists; this project is already locked.
Use namba-intent status --proof --next-questions, then the amend/relock flow for locked planning changes.
No files changed by namba-intent init; the lockfile was not modified.
```

Command 후 SHA-256은 그대로였습니다:

```text
61d7265d58137893995f2a693b83d213874ab5a2028ecb693857350d0eb7c16c
```

Project root에서 `namba-intent end`를 실행하지 않았고 root relock도 하지
않았습니다.

## Deprecated ni shim

| Behavior | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| deprecation warning | `ni is deprecated; use namba-intent.`가 stderr로 출력된다. | stderr-only capture에 정확히 `ni is deprecated; use namba-intent.`가 있었다. | Yes | Warning은 stdout help output과 분리됐다. |
| delegation | Shim이 behavior를 delegate한다. | `ni --help`, `ni version`, `ni init . --yes`가 shared CLI behavior로 실행됐다. | Yes | `ni version`은 `0.0.0-dev`; temp init은 planning artifacts를 생성. |
| primary command boundary | `ni`는 transition-only이며 primary command가 아니다. | README/docs와 primary binary smoke는 `namba-intent`를 사용한다. `ni` help는 shim 호출 때문에 legacy command name을 사용한다. | Yes | `ni`를 primary path로 만들지 않는다. |
| Windows support boundary | Windows `ni` shim support를 claim하지 않는다. | Windows `ni` shim support를 claim하거나 verify하지 않았다. | Yes | Current Windows primary installer는 `namba-intent.exe` 사용. |

## TUI model coverage

| Behavior | Evidence | Pass? | Notes |
| --- | --- | --- | --- |
| Bubble Tea v2 model | `internal/tui/init/model_test.go` imports `charm.land/bubbletea/v2`. | Yes | Existing focused tests cover the model. |
| Init | `TestModelInitialStateUsesAltScreen`가 `NewModel`을 생성한다. | Yes | Default project name도 확인. |
| Update | `TestUpdateHandlesUpDownAndLeftRight`, `TestUpdateHandlesEnterEscAndQ`, `TestConfirmPathReturnsIntent`, `TestCancelPathWritesNothingSignal`, `TestExistingFileChoices`. | Yes | Direct model transition tests. |
| View | `TestModelInitialStateUsesAltScreen`가 `m.View()` 호출. | Yes | Brittle terminal screenshot 없음. |
| AltScreen | `TestModelInitialStateUsesAltScreen`가 `view.AltScreen` assert. | Yes | AltScreen behavior cover. |
| Lip Gloss v2 | `internal/tui/init/model.go` uses `charm.land/lipgloss/v2`. | Yes | Model/view compilation과 tests로 cover; screenshot assertion 불필요. |
| keyboard navigation | `TestUpdateHandlesUpDownAndLeftRight`, `TestUpdateHandlesEnterEscAndQ`. | Yes | Up/Down, Left/Right, Enter, Esc, q cover. |
| cancel | `TestCancelPathWritesNothingSignal` 및 `TestUpdateHandlesEnterEscAndQ`의 q path. | Yes | Canceled result 검증. |
| confirm | `TestConfirmPathReturnsIntent`. | Yes | Confirmed intent result 검증. |
| plain summary | `internal/cli/cli_test.go` init tests assert init summaries. | Yes | Existing CLI tests cover post-init textual summary paths. |

## README/docs audit

| Surface | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `README.md` | Namba Intent, `namba-intent` primary, v0.6.0 publication overclaim 없음. | Namba Intent와 `namba-intent`를 사용하고, v0.6.0 publish 전 latest published v0.5.1은 여전히 `ni`일 수 있다고 말한다. | Yes | Homebrew는 Planned / v0.5 candidate. |
| `README.ko.md` | Korean companion이 claim을 넓히지 않는다. | README와 aligned; v0.6.0 not-published, Windows/Homebrew boundaries 유지. | Yes | `Skills are UX; CLI is authority.` 유지. |
| install docs | Current-tree installer behavior와 v0.5.1 public-release evidence를 분리한다. | `docs/22_INSTALL.md`와 curl install docs는 verified v0.5.1 `ni` release evidence를 보존하고, README와 실제 installers는 upcoming/current-tree `namba-intent`를 문서화한다. | Yes with notes | Public v0.5.1 proof가 아직 `namba-intent`가 아니라 `ni`인 점은 의도된 boundary. |
| `docs/135` | Rename plan이 product/execution boundaries를 보존한다. | Tracked docs/135 pair가 있고 `namba-intent`, `.ni/`, future v0.6.0 release, non-execution boundaries를 유지한다. | Yes | Root relock 없음. |
| `docs/136` | Implementation record가 rename/claim boundaries를 보존한다. | Tracked docs/136 pair가 current-tree rename, transition-only `ni`, no public v0.6.0 proof를 기록한다. | Yes | 이 smoke는 다음 record로 docs/137을 추가한다. |
| skill packs | Skills are UX and must not replace CLI authority. | `.agents/skills`와 package skill docs는 `namba-intent` commands를 사용하고 documented transition skill IDs는 유지한다. | Yes | Skill pack execution behavior 추가 없음. |

## Installer surface audit

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `install.sh` | `namba-intent`를 primary로 install한다. | `BIN_NAME="namba-intent"`와 asset prefix `namba-intent_`를 사용하고 next steps는 `namba-intent --help`, `namba-intent version`을 verify한다. | Yes | Current-tree script only; public v0.6.0 install proof 아님. |
| `install.ps1` | `namba-intent.exe`를 primary로 install한다. | `%LOCALAPPDATA%\namba-intent\bin` 기본, target은 `namba-intent.exe`, verification steps도 `namba-intent`. | Yes | Static/current-tree audit only. |
| uninstall | Primary install을 제거한다. | Unix uninstall은 installed `namba-intent` binary와 managed PATH block을 제거; PowerShell uninstall은 `namba-intent.exe`와 User PATH entry를 제거. | Yes | Public uninstall은 실행하지 않았다. |
| Windows alias handling | Windows primary path는 `ni` alias cleanup에 의존하지 않는다. | PowerShell installer는 `PowerShell ni alias cleanup is not required for namba-intent.exe.`라고 말한다. | Yes | Windows real-host transcript 없음. |
| public release boundary | Public v0.6.0 install works를 claim하지 않는다. | README/docs는 v0.6.0 not published와 v0.5.1 public proof가 historical `ni`임을 말한다. | Yes | Release, tag, asset upload, GoReleaser publish 없음. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Namba Intent identity | Project Intent Compiler for AI Agents. | README, help, docs가 이 identity를 사용한다. | Yes | Pre-runtime boundary preserved. |
| namba-intent command | Primary current-tree command. | PATH-resolved temp binary smoke 통과. | Yes | Current-tree only. |
| ni shim | Deprecated transition path only. | Stderr warning과 delegation 확인. | Yes | Primary 아님. |
| .ni compatibility | `.ni/` unchanged. | Init은 `.ni/*`를 만들고 root protected `.ni` diff는 edit 전 empty였다. | Yes | `.ni` rename 없음. |
| v0.6.0 publication status | Not published. | `git tag --list v0.6.0`는 empty. | Yes | Release action 없음. |
| public install | Public install이 `namba-intent`를 가져온다고 claim하지 않는다. | Claim 없음. | Yes | Public v0.5.1 proof는 `ni`. |
| Homebrew | Planned / v0.5 candidate. | Available claim 관찰/추가 없음. | Yes | Homebrew formula work 없음. |
| Windows real-host verification | Transcript 전까지 pending. | Pending; installer surface만 audit. | Yes | Windows host 사용 안 함. |
| run behavior | Bounded prompt compile only. | Help와 prompt code가 non-execution wording 유지. | Yes | 이 smoke에서 `run` 실행 없음. |
| runtime execution boundary | Task runner, SPEC runner, execution harness, shell/Codex adapter, queue, PR/release automation, downstream execution layer 없음. | 해당 behavior를 추가하거나 실행하지 않았다. | Yes | Smoke는 help/version/init/status만 사용. |

## Git status / inclusion check

| Path or group | git status --short | Expected in v0.6.0? | Notes |
| --- | --- | --- | --- |
| `README.md` | clean at smoke start | Yes | Tracked; current-tree Namba Intent surface. |
| `README.ko.md` | clean at smoke start | Yes | Tracked Korean companion. |
| `cmd/namba-intent/*` | clean at smoke start | Yes | Tracked primary entrypoint. |
| `cmd/ni/*` | clean at smoke start | Yes | Tracked deprecated shim. |
| `internal/cli/*` | clean at smoke start | Yes | Tracked shared CLI. |
| `install.sh` | clean at smoke start | Yes | Tracked current-tree Unix installer. |
| `install.ps1` | clean at smoke start | Yes | Tracked current-tree PowerShell installer. |
| `docs/135*` | clean at smoke start | Yes | Tracked rename plan pair. |
| `docs/136*` | clean at smoke start | Yes | Tracked rename implementation pair. |
| `docs/137*` | new in this task | Yes | This smoke record pair. |
| `scripts/*` | clean at smoke start | Yes | Existing checkers and smoke scripts. |
| `packages/*` | clean at smoke start | Yes | Existing package surfaces. |
| `.agents/*` | clean at smoke start | Yes | Existing repo-local skill surfaces. |
| `.ni/contract.json` | no diff | Yes | Protected root file; not edited. |
| `.ni/session.json` | no diff | Yes | Protected root file; not edited. |
| `.ni/plan.lock.json` | no diff | Yes | Protected root file; not edited. |
| unexpected files | none in repo at smoke start | No | Temporary binaries/projects는 repo 밖 `/tmp/namba-intent-first-user-smoke` 아래 생성. |

## Validation results

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Passed | Smoke start에서 empty. |
| `git log --oneline --decorate -20` | Passed | `HEAD -> main` at `80acd80 Implement Namba Intent rename`; `v0.5.1` tag는 HEAD 뒤에 있음. |
| `git tag --list v0.5.1` | Passed | `v0.5.1` 반환. |
| `git tag --list v0.6.0` | Passed | Empty. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed | Documentation edit 전 empty. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /tmp/namba-intent-first-user-smoke/bin/namba-intent ./cmd/namba-intent` | Passed | Temporary current-tree binary를 repo 밖에 build. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /tmp/namba-intent-first-user-smoke/bin/ni ./cmd/ni` | Passed | Temporary deprecated shim binary를 repo 밖에 build. |
| `command -v namba-intent` | Passed | Temporary PATH에서 resolve. |
| `namba-intent --help` | Passed | Namba Intent와 `namba-intent` primary 사용. |
| `namba-intent version` | Passed | `0.0.0-dev` 반환. |
| `namba-intent init . --yes` | Passed | Temp project에 first-user planning workspace 생성. |
| `namba-intent status --proof --next-questions` | Passed | Truthful `BLOCKED` state와 next questions 보고. |
| repeated `namba-intent init . --yes` | Passed | Existing files unchanged; lockfile 생성 없음. |
| lockfile fixture `namba-intent init . --yes` | Passed | Existing lockfile warning; lockfile SHA-256 unchanged. |
| `command -v ni` | Passed | Temporary PATH에서 transition shim resolve. |
| `ni --help` | Passed | Stderr deprecation warning과 delegated help. |
| `ni version` | Passed | `0.0.0-dev`와 warning 반환. |
| `ni init . --yes` | Passed | Delegated temp init; lockfile 생성 없음. |
| `gofmt -w .` | Passed | Output 없음. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed | `internal/tui/init` 포함 모든 packages 통과. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed | JSON, schema, formatting, README surface, install docs, skill packs, prompt budget, core boundary, assets, Go tests, smoke checks 통과. Script 내부 fixture `end`와 `relock` path는 project root와 분리됨. |

## Changes made

- `docs/137_NAMBA_INTENT_FIRST_USER_SMOKE.md`: current-tree first-user smoke 기록.
- `docs/137_NAMBA_INTENT_FIRST_USER_SMOKE.ko.md`: Korean companion.

## What this smoke proves

State only:
- current-tree namba-intent command-name path works;
- current-tree Namba Intent first-user init path works;
- deprecated ni shim behavior is transition-safe;
- .ni/ compatibility remains intact;
- no release action was performed.

## What this smoke does not prove

State:
- v0.6.0 has been published;
- public install retrieves namba-intent;
- Windows real-host execution works;
- Homebrew is Available;
- external users succeed;
- downstream execution succeeds.

## Recommended next task

A. v0.6.0 public install parity / release readiness audit
