# First-User Onboarding Smoke After TUI

## Current status

State:
- v0.5.0 publication: verified
- README redesigned: completed
- ni init . Bubble Tea TUI: implemented
- current-tree first-user smoke: verified in this document
- Windows real-host execution: deferred on macOS-only development host
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Smoke goal

This smoke verifies first-user onboarding after the Bubble Tea TUI using the
current repository tree. It checks command-name resolution, README path
coherence, temporary-project initialization, init overwrite safety, lockfile
safety, TUI model coverage, and a local interactive TTY run.

It does not verify release publication, Homebrew availability, Windows
real-host execution, external-user success, or downstream execution.

## Smoke decision

FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES

Notes:
- Current-tree command-name onboarding passed in a tested macOS shell context.
- `ni status --proof --next-questions` ran after init and honestly reported
  `BLOCKED` for the empty first-user template because project purpose, actors,
  delivery surface, and an open blocker question remain unresolved.
- The tested binary was built from the current checkout and returned
  `0.0.0-dev`; this is not proof that the published v0.5.0 artifact includes
  the current TUI behavior.
- Windows real-host execution remains deferred because no Windows host, VM, CI
  runner, or tester transcript was used.

## Smoke lanes

| Lane | Purpose | Binary source | Result | Notes |
| --- | --- | --- | --- | --- |
| current-tree command-name smoke | Verify `ni` resolves by command name and basic commands work. | `/private/tmp/ni-first-user-smoke/bin/ni` built from current checkout | Pass | PATH placed the temp binary directory first. |
| non-interactive first-user project smoke | Verify `ni init . --yes` creates planning artifacts and status runs. | current-tree temp binary | Pass with expected `BLOCKED` status | `BLOCKED` reflects incomplete first-user intent, not a command failure. |
| repeated init safety smoke | Verify existing files are not silently overwritten. | current-tree temp binary | Pass | Existing files were reported as `unchanged`; representative hashes were unchanged. |
| lockfile safety smoke | Verify `ni init . --yes` does not modify an existing `.ni/plan.lock.json`. | current-tree temp binary | Pass | Sentinel lockfile hash stayed unchanged; no relock occurred. |
| TUI model test smoke | Verify Bubble Tea model behavior through Go tests. | source tests | Pass | Covered by `internal/tui/init` tests. |
| interactive TTY smoke | Verify a real TTY opens the wizard and exits to plain summary. | current-tree temp binary | Pass | AltScreen opened, arrows moved the cursor, confirmation wrote artifacts, summary printed after exit. |
| Windows real-host smoke | Verify Windows install and command-name flow on real Windows. | not run | Deferred | No Windows host or transcript was available. |

## Command-name smoke

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| resolve command | `command -v ni` | Path from command lookup, not absolute invocation proof. | `/private/tmp/ni-first-user-smoke/bin/ni` | Yes | Run through a fresh `env -i` shell with temp PATH. |
| help | `ni --help` | Usage prints successfully. | Help printed command list including `init`, `status`, `end`, and `run`. | Yes | No project files touched. |
| version | `ni version` | Version command succeeds. | `0.0.0-dev` | Yes | Current-tree source build, not release artifact proof. |

## Project artifact check

| Artifact | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `.ni/contract.json` | Created by `ni init . --yes`. | Present. | Yes | Temporary project only. |
| `.ni/session.json` | Created by `ni init . --yes`. | Present. | Yes | Temporary project only. |
| `docs/plan/**` | Created by `ni init . --yes`. | 12 planning docs present. | Yes | `00_project_brief.md` through `11_decision_log.md`. |
| `.ni/plan.lock.json` | Not created by init. | Absent in first-user and repeated-init smokes. | Yes | No `ni end` was run. |
| downstream generated prompt | Not created or executed by init/status. | `.ni/generated` and `.ni/run.log` absent. | Yes | No `ni run` was run. |

## TUI behavior check

| Behavior | Evidence | Pass? | Notes |
| --- | --- | --- | --- |
| Bubble Tea v2 model | `internal/tui/init/model.go` imports `charm.land/bubbletea/v2`; tests passed. | Yes | Source-level proof. |
| Init | `Model.Init()` exists and returns no command; `go test ./internal/tui/init` passed. | Yes | Model-level smoke. |
| Update | `TestUpdateHandlesUpDownAndLeftRight`, `TestUpdateHandlesEnterEscAndQ`. | Yes | Keyboard behavior tested. |
| View | `TestModelInitialStateUsesAltScreen` calls `View()`. | Yes | `View()` returns `tea.View`. |
| AltScreen | `view.AltScreen = true`; real TTY emitted AltScreen enter/exit control flow. | Yes | Observed in interactive smoke. |
| Lip Gloss v2 styling | `internal/tui/init/model.go` imports `charm.land/lipgloss/v2`; tests passed. | Yes | Styling is render-only. |
| Up/Down | Model test and real TTY cursor movement. | Yes | Down and Up/Down navigation observed. |
| Left/Right | `TestUpdateHandlesUpDownAndLeftRight`. | Yes | Model-level coverage. |
| Enter | `TestUpdateHandlesEnterEscAndQ`; real TTY confirmation wrote artifacts. | Yes | Confirmation flow passed. |
| Esc | `TestUpdateHandlesEnterEscAndQ`. | Yes | Model-level coverage. |
| q | `TestUpdateHandlesEnterEscAndQ`. | Yes | Model-level coverage. |
| plain text summary | Real TTY printed `initialized ni planning workspace at .` after AltScreen exit. | Yes | Summary is not a readiness claim. |

## README alignment

| README claim | Smoke evidence | Pass? | Notes |
| --- | --- | --- | --- |
| macOS install path leads to command-name verification. | Current-tree command-name smoke passed with temp PATH. | Yes | This smoke did not run the curl installer. |
| Windows install path is documented but real-host execution is deferred. | README says Windows transcript is required before verified claim. | Yes | No Windows transcript was produced. |
| First project uses `ni init .`. | Temporary first-user project ran `ni init . --yes`; TTY ran `ni init .`. | Yes | Non-interactive and interactive paths both worked locally. |
| First project uses `ni status --proof --next-questions`. | Command ran and reported actual `BLOCKED` state with next questions. | Yes | CLI authority preserved. |
| First project later uses `ni end`. | README includes the step. | Yes | Not run in this smoke. |
| First project later uses `ni run --max-chars 4000`. | README states prompt compilation only. | Yes | Not run in this smoke. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| macOS current-tree command-name smoke | May verify only the tested shell context. | Verified with temp PATH and current-tree binary. | Yes | Does not prove every shell config. |
| Windows real-host execution | Deferred without Windows transcript. | Deferred. | Yes | No Windows verified claim. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | Planned wording preserved. | Yes | No `brew install` claim added. |
| ni init . | Initializes planning artifacts only. | Created docs and `.ni` skeleton; no lockfile. | Yes | TUI does not own readiness. |
| ni status | CLI-authoritative readiness gate. | Reported `BLOCKED` with proof and next questions. | Yes | Model judgment did not override status. |
| ni end | Lock authority only after CLI gate permits it. | Not run. | Yes | No project-root or fixture lock was created with `ni end`. |
| ni run | Bounded prompt compiler only. | Not run. | Yes | No downstream prompt executed. |
| READY | Must come from `ni status`. | No READY claim made for empty temp project. | Yes | Actual state was `BLOCKED`. |
| TUI readiness boundary | TUI must not determine readiness or lock plans. | TUI wrote init artifacts only. | Yes | Status remained separate. |
| Runtime execution boundary | No task runner, execution harness, shell/Codex adapter, queue, PR/release automation, or downstream execution layer. | No runtime behavior added or executed. | Yes | Smoke used only help/version/init/status/tests. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | clean before docs/125 | No | Read-only audit; no README change needed. |
| README.ko.md | clean before docs/125 | No | Read-only audit; no companion README change needed. |
| cmd/ni/* | clean before docs/125 | No | No code blocker found. |
| internal/core/docstore/* | clean before docs/125 | No | No narrow fix needed. |
| internal/tui/init/* | clean before docs/125 | No | Existing tests covered required behavior. |
| go.mod | clean before docs/125 | No | No dependency change. |
| go.sum | clean before docs/125 | No | No dependency change. |
| docs/124* | clean before docs/125 | No | Used as context only. |
| docs/125* | added | Yes | This smoke result and Korean companion. |
| temporary smoke directories | outside repo | No | `/private/tmp/ni-first-user-smoke/**`. |
| `.ni/contract.json` | no diff | No | Protected project-root planning state. |
| `.ni/session.json` | no diff | No | Protected project-root session state. |
| `.ni/plan.lock.json` | no diff | No | Protected project-root lockfile. |
| unexpected files | none observed before docs/125 | No | Rechecked with `git status --short`. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed before smoke; no unstaged changes. |
| `git log -1 --oneline` | `4ae4a85 Redesign README and add Bubble Tea init TUI`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed before and during smoke; no output. |
| `GOCACHE=/private/tmp/ni-go-cache go build -o /private/tmp/ni-first-user-smoke/bin/ni ./cmd/ni` | Passed. |
| `command -v ni` in fresh temp PATH shell | Passed; resolved `/private/tmp/ni-first-user-smoke/bin/ni`. |
| `ni --help` in fresh temp PATH shell | Passed. |
| `ni version` in fresh temp PATH shell | Passed; output `0.0.0-dev`. |
| `ni init . --yes` in temporary first-user project | Passed; planning files created. |
| `ni status --proof --next-questions` in temporary first-user project | Passed; reported actual `BLOCKED` state and next questions. |
| repeated `ni init . --yes` in temporary project | Passed; existing files reported unchanged. |
| repeated-init representative `shasum -a 256` check | Passed; `.ni/contract.json`, `.ni/session.json`, and `docs/plan/00_project_brief.md` hashes unchanged. |
| lockfile fixture `ni init . --yes` | Passed; warning printed and `.ni/plan.lock.json` hash unchanged. |
| `go test ./internal/tui/init` with temp `GOCACHE` | Passed. |
| `go test ./cmd/ni` with temp `GOCACHE` | Passed. |
| interactive TTY `ni init .` in temporary project | Passed; AltScreen, navigation, confirmation, artifact creation, and plain summary observed. |

## Changes made

- Added `docs/125_FIRST_USER_ONBOARDING_SMOKE_AFTER_TUI.md`.
- Added `docs/125_FIRST_USER_ONBOARDING_SMOKE_AFTER_TUI.ko.md`.
- No Go code changes were needed.
- No README claim changes were needed.

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
