# First-User Onboarding Smoke After TUI

## Current status

State:
- v0.5.0 publication: verified
- README redesigned: completed
- ni init . Bubble Tea TUI: implemented
- current-tree first-user smoke: 이 문서에서 verified
- Windows real-host execution: macOS-only development host에서는 deferred
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Smoke goal

이 smoke는 current repository tree 기준으로 Bubble Tea TUI 이후 first-user
onboarding을 검증한다. Command-name resolution, README path coherence,
temporary project init, init overwrite safety, lockfile safety, TUI model
coverage, local interactive TTY run을 확인한다.

Release publication, Homebrew availability, Windows real-host execution,
external-user success, downstream execution은 검증하지 않는다.

## Smoke decision

FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES

Notes:
- Current-tree command-name onboarding은 tested macOS shell context에서 passed.
- `ni status --proof --next-questions`는 init 뒤 실행됐고, empty first-user
  template에 project purpose, actors, delivery surface, open blocker question이
  남아 있어 정직하게 `BLOCKED`를 report했다.
- Tested binary는 current checkout에서 build했고 `0.0.0-dev`를 반환했다. 이것은
  published v0.5.0 artifact가 current TUI behavior를 포함한다는 proof가 아니다.
- Windows host, VM, CI runner, tester transcript를 사용하지 않았으므로 Windows
  real-host execution은 deferred 상태다.

## Smoke lanes

| Lane | Purpose | Binary source | Result | Notes |
| --- | --- | --- | --- | --- |
| current-tree command-name smoke | `ni`가 command name으로 resolve되고 basic commands가 동작하는지 확인. | `/private/tmp/ni-first-user-smoke/bin/ni` built from current checkout | Pass | Temp binary directory를 PATH 앞에 두었다. |
| non-interactive first-user project smoke | `ni init . --yes`가 planning artifacts를 만들고 status가 실행되는지 확인. | current-tree temp binary | Pass with expected `BLOCKED` status | `BLOCKED`는 incomplete first-user intent 때문이며 command failure가 아니다. |
| repeated init safety smoke | Existing files를 silently overwrite하지 않는지 확인. | current-tree temp binary | Pass | Existing files는 `unchanged`로 report됐고 representative hashes unchanged. |
| lockfile safety smoke | Existing `.ni/plan.lock.json`을 `ni init . --yes`가 수정하지 않는지 확인. | current-tree temp binary | Pass | Sentinel lockfile hash unchanged; relock 없음. |
| TUI model test smoke | Go tests로 Bubble Tea model behavior 확인. | source tests | Pass | `internal/tui/init` tests가 cover. |
| interactive TTY smoke | Real TTY에서 wizard가 열리고 plain summary로 종료되는지 확인. | current-tree temp binary | Pass | AltScreen open, arrow navigation, confirmation write, summary after exit observed. |
| Windows real-host smoke | Real Windows에서 install/command-name flow 확인. | not run | Deferred | Windows host 또는 transcript 없음. |

## Command-name smoke

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| resolve command | `command -v ni` | Absolute invocation proof가 아니라 command lookup path. | `/private/tmp/ni-first-user-smoke/bin/ni` | Yes | Fresh `env -i` shell과 temp PATH 사용. |
| help | `ni --help` | Usage가 성공적으로 출력된다. | `init`, `status`, `end`, `run` 포함 command list 출력. | Yes | Project files untouched. |
| version | `ni version` | Version command succeeds. | `0.0.0-dev` | Yes | Current-tree source build이며 release artifact proof 아님. |

## Project artifact check

| Artifact | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `.ni/contract.json` | `ni init . --yes`가 생성. | Present. | Yes | Temporary project only. |
| `.ni/session.json` | `ni init . --yes`가 생성. | Present. | Yes | Temporary project only. |
| `docs/plan/**` | `ni init . --yes`가 생성. | 12 planning docs present. | Yes | `00_project_brief.md` through `11_decision_log.md`. |
| `.ni/plan.lock.json` | Init이 만들면 안 됨. | First-user/repeated-init smokes에서 absent. | Yes | `ni end` 실행 없음. |
| downstream generated prompt | Init/status가 만들거나 실행하면 안 됨. | `.ni/generated`와 `.ni/run.log` absent. | Yes | `ni run` 실행 없음. |

## TUI behavior check

| Behavior | Evidence | Pass? | Notes |
| --- | --- | --- | --- |
| Bubble Tea v2 model | `internal/tui/init/model.go` imports `charm.land/bubbletea/v2`; tests passed. | Yes | Source-level proof. |
| Init | `Model.Init()` exists and returns no command; `go test ./internal/tui/init` passed. | Yes | Model-level smoke. |
| Update | `TestUpdateHandlesUpDownAndLeftRight`, `TestUpdateHandlesEnterEscAndQ`. | Yes | Keyboard behavior tested. |
| View | `TestModelInitialStateUsesAltScreen` calls `View()`. | Yes | `View()` returns `tea.View`. |
| AltScreen | `view.AltScreen = true`; real TTY에서 AltScreen enter/exit control flow observed. | Yes | Interactive smoke에서 관찰. |
| Lip Gloss v2 styling | `internal/tui/init/model.go` imports `charm.land/lipgloss/v2`; tests passed. | Yes | Styling is render-only. |
| Up/Down | Model test and real TTY cursor movement. | Yes | Down and Up/Down navigation observed. |
| Left/Right | `TestUpdateHandlesUpDownAndLeftRight`. | Yes | Model-level coverage. |
| Enter | `TestUpdateHandlesEnterEscAndQ`; real TTY confirmation wrote artifacts. | Yes | Confirmation flow passed. |
| Esc | `TestUpdateHandlesEnterEscAndQ`. | Yes | Model-level coverage. |
| q | `TestUpdateHandlesEnterEscAndQ`. | Yes | Model-level coverage. |
| plain text summary | Real TTY printed `initialized ni planning workspace at .` after AltScreen exit. | Yes | Summary는 readiness claim이 아니다. |

## README alignment

| README claim | Smoke evidence | Pass? | Notes |
| --- | --- | --- | --- |
| macOS install path leads to command-name verification. | Current-tree command-name smoke passed with temp PATH. | Yes | Curl installer는 이번 smoke에서 실행하지 않았다. |
| Windows install path is documented but real-host execution is deferred. | README는 verified claim 전에 Windows transcript가 필요하다고 말한다. | Yes | Windows transcript 없음. |
| First project uses `ni init .`. | Temporary first-user project ran `ni init . --yes`; TTY ran `ni init .`. | Yes | Non-interactive and interactive paths both worked locally. |
| First project uses `ni status --proof --next-questions`. | Command ran and reported actual `BLOCKED` state with next questions. | Yes | CLI authority preserved. |
| First project later uses `ni end`. | README includes the step. | Yes | 이번 smoke에서는 실행하지 않음. |
| First project later uses `ni run --max-chars 4000`. | README states prompt compilation only. | Yes | 이번 smoke에서는 실행하지 않음. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| macOS current-tree command-name smoke | Tested shell context만 verify 가능. | Temp PATH와 current-tree binary로 verified. | Yes | 모든 shell config proof 아님. |
| Windows real-host execution | Windows transcript 없으면 deferred. | Deferred. | Yes | Windows verified claim 없음. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | Planned wording preserved. | Yes | `brew install` claim 추가 없음. |
| ni init . | Planning artifacts only. | Docs와 `.ni` skeleton 생성; lockfile 없음. | Yes | TUI는 readiness authority가 아니다. |
| ni status | CLI-authoritative readiness gate. | Proof와 next questions를 포함해 `BLOCKED` report. | Yes | Model judgment가 status를 override하지 않음. |
| ni end | CLI gate가 허용한 뒤 lock authority. | Not run. | Yes | Project-root 또는 fixture lock을 `ni end`로 만들지 않음. |
| ni run | Bounded prompt compiler only. | Not run. | Yes | Downstream prompt execution 없음. |
| READY | `ni status`에서만 나와야 함. | Empty temp project에 READY claim 없음. | Yes | Actual state는 `BLOCKED`. |
| TUI readiness boundary | TUI가 readiness를 결정하거나 lock하면 안 됨. | TUI는 init artifacts만 작성. | Yes | Status는 separate. |
| Runtime execution boundary | Task runner, execution harness, shell/Codex adapter, queue, PR/release automation, downstream execution layer 없음. | Runtime behavior 추가/실행 없음. | Yes | Help/version/init/status/tests만 사용. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | clean before docs/125 | No | Read-only audit; README change 불필요. |
| README.ko.md | clean before docs/125 | No | Read-only audit; companion README change 불필요. |
| cmd/ni/* | clean before docs/125 | No | Code blocker 없음. |
| internal/core/docstore/* | clean before docs/125 | No | Narrow fix 불필요. |
| internal/tui/init/* | clean before docs/125 | No | Existing tests covered required behavior. |
| go.mod | clean before docs/125 | No | Dependency change 없음. |
| go.sum | clean before docs/125 | No | Dependency change 없음. |
| docs/124* | clean before docs/125 | No | Context로만 사용. |
| docs/125* | added | Yes | This smoke result and Korean companion. |
| temporary smoke directories | outside repo | No | `/private/tmp/ni-first-user-smoke/**`. |
| `.ni/contract.json` | no diff | No | Protected project-root planning state. |
| `.ni/session.json` | no diff | No | Protected project-root session state. |
| `.ni/plan.lock.json` | no diff | No | Protected project-root lockfile. |
| unexpected files | none observed before docs/125 | No | `git status --short`로 rechecked. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Smoke 전 passed; unstaged changes 없음. |
| `git log -1 --oneline` | `4ae4a85 Redesign README and add Bubble Tea init TUI`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Smoke 전/중 passed; no output. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /private/tmp/ni-first-user-smoke/bin/ni ./cmd/ni` | Passed. |
| Fresh temp PATH shell의 `command -v ni` | Passed; `/private/tmp/ni-first-user-smoke/bin/ni`로 resolved. |
| Fresh temp PATH shell의 `ni --help` | Passed. |
| Fresh temp PATH shell의 `ni version` | Passed; output `0.0.0-dev`. |
| Temporary first-user project의 `ni init . --yes` | Passed; planning files created. |
| Temporary first-user project의 `ni status --proof --next-questions` | Passed; actual `BLOCKED` state와 next questions report. |
| Temporary project의 repeated `ni init . --yes` | Passed; existing files reported unchanged. |
| Repeated-init representative `shasum -a 256` check | Passed; `.ni/contract.json`, `.ni/session.json`, `docs/plan/00_project_brief.md` hashes unchanged. |
| Lockfile fixture `ni init . --yes` | Passed; warning printed and `.ni/plan.lock.json` hash unchanged. |
| Temp `GOCACHE`의 `go test ./internal/tui/init` | Passed. |
| Temp `GOCACHE`의 `go test ./cmd/ni` | Passed. |
| Temporary project의 interactive TTY `ni init .` | Passed; AltScreen, navigation, confirmation, artifact creation, plain summary observed. |

## Changes made

- Added `docs/125_FIRST_USER_ONBOARDING_SMOKE_AFTER_TUI.md`.
- Added `docs/125_FIRST_USER_ONBOARDING_SMOKE_AFTER_TUI.ko.md`.
- Go code changes were not needed.
- README claim changes were not needed.

## What this smoke proves

- Current-tree command-name onboarding works for the tested macOS shell context.
- Non-interactive `ni init . --yes` creates expected planning artifacts.
- `ni status --proof --next-questions` runs after init and reports the actual
  readiness state.
- TUI model behavior is covered by tests.
- Real local TTY onboarding opens the Bubble Tea TUI, accepts navigation and
  confirmation, and exits to a plain text summary.
- Project-root `.ni` files were not modified.

## What this smoke does not prove

- Published v0.5.0 artifact includes current-tree TUI behavior.
- Windows real-host execution works unless a Windows transcript exists.
- Homebrew is Available.
- Every shell configuration works.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.

## Recommended next task

F. public install parity / patch-release readiness audit

## Next task prompt

Proceed in `/Users/namba/Documents/project/ni`.

Goal: run a public install parity / patch-release readiness audit after the
current-tree first-user onboarding smoke passed with notes.

Use the current repository tree as source of truth, but do not claim that the
published v0.5.0 artifact includes the Bubble Tea TUI or latest onboarding
behavior unless the hosted artifact is explicitly tested and proves it.

Scope:
- Read `AGENTS.md`, `README.md`, `README.ko.md`, `docs/22_INSTALL.md`,
  `docs/install-curl.md`, `docs/install-curl.ko.md`,
  `docs/124_REPO_CLEANUP_README_REDESIGN_INIT_TUI.md`,
  `docs/124_REPO_CLEANUP_README_REDESIGN_INIT_TUI.ko.md`, and
  `docs/125_FIRST_USER_ONBOARDING_SMOKE_AFTER_TUI.md`.
- Compare public install claims with current-tree smoke evidence.
- Identify whether a patch release or additional hosted-artifact verification
  is needed before public docs imply the TUI/onboarding behavior is available
  through the released install paths.
- Keep Homebrew: Planned / v0.5 candidate.
- Keep Windows real-host execution deferred unless a Windows transcript exists.
- Keep Model workspace packs: Experimental.
- Keep No-terminal method: Experimental / assisted.
- Preserve `Skills are UX; CLI is authority.`

Rules:
- Do not publish, tag, release, upload assets, run release workflows, create or
  publish a Homebrew formula, run `ni end` on the project root, relock the
  project root, execute generated prompts, add task-runner behavior, add shell
  or Codex adapters, or make `ni run` execute downstream work.
- Do not edit `.ni/contract.json`, `.ni/session.json`, or `.ni/plan.lock.json`.

Expected output:
- Add or update the narrowest docs needed to record public install parity,
  patch-release readiness, and remaining proof gaps.
- Include a claim-boundary table for release binary, curl installer, Homebrew,
  Windows, TUI behavior, and `ni run`.
- Run `gofmt -w .`, `GOCACHE=/private/tmp/ni-go-cache go test ./...`,
  `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh`, and
  `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json`.
- Final report must list changed files, validation results, and confirmation
  that no project-root lockfile edit, root relock, release action, Homebrew
  Available claim, Windows verified claim, or downstream execution was added.
