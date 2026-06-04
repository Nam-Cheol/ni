# Repo Cleanup, README Redesign, and Init TUI

## Current status

- v0.5.0 publication: verified.
- README two-path onboarding: revised in this task.
- ni init . Bubble Tea TUI: implemented in this task.
- Windows real-host execution: deferred on macOS-only development host.
- Homebrew: Planned / v0.5 candidate.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Task goal

This pass cleans accidental repo files, simplifies the README into a modern CLI
product surface, and upgrades `ni init .` from console prompts to a real Bubble
Tea v2 / Lip Gloss v2 TUI while preserving non-interactive fallback behavior.

## Repository cleanup

| File or path | Classification | Action | Reason | Notes |
| --- | --- | --- | --- | --- |
| `examples/namba-ai-upgrade/docs/plan/* 2.md` | delete | Deleted 12 files | Finder-style duplicate files with older TODO template content and no references. | Canonical files without ` 2` suffix remain. |
| `README 2.md`, `README.ko 2.md` | keep absent | No action | Existing validator already checks these root duplicate names. | No files present. |
| `install 2.sh`, `install 2.ps1` | keep absent | No action | Requested suspicious installer copies were not present. | No docs/scripts reference them. |
| `*.tmp`, `*.bak`, `Untitled*` | keep absent | No action | No matching local-only temp artifacts were present. | Search returned no files. |
| Release archives/assets | document as intentional | No action | Release and packaging files are part of existing documented distribution work. | No release publish/tag/upload performed. |

## README redesign

| Section | Before | After | Notes |
| --- | --- | --- | --- |
| Hero | Longer visual story with multiple card image groups. | One hero, one-line definition, short product explanation, one flow image. | Keeps existing safe assets only. |
| Install | Correct but heavier surrounding copy. | macOS and Windows primary paths only, plus detailed-doc link. | Homebrew remains Planned / v0.5 candidate. |
| First project | Present but embedded in longer narrative. | `mkdir`, `cd`, `ni init .`, `ni status`, `ni end`, `ni run` flow. | States `ni run` does not execute downstream work. |
| Feature list | Split across why/payoff sections. | `What ni does` command table. | CLI authority preserved. |
| Non-goals | Present near bottom. | Short `What ni does not do` section. | Runtime execution boundary preserved. |
| Status | Inline maturity note. | Concise status list. | No Homebrew or Windows overclaim. |

## README inspiration

The README structure was simplified in the style of modern CLI projects such as
uv, Deno, Bun, GitHub CLI, and Starship, without copying their content,
branding, wording, or layout wholesale.

## Init TUI architecture

| Layer | Responsibility | Package/path | Notes |
| --- | --- | --- | --- |
| command layer | Parse flags, choose TUI vs fallback, print post-TUI summary. | `cmd/ni/main.go` | Does not run `ni status`, `ni end`, or `ni run`. |
| domain init logic | Validate init options, build file plan, protect lockfile, write missing files. | `internal/core/docstore` | TUI does not own file writing. |
| TUI model | Hold wizard state, key handling, confirmation/cancel result. | `internal/tui/init` | Model, Init, Update, View are implemented. |
| TUI styling | Render the wizard with Lip Gloss v2 styles. | `internal/tui/init/model.go` | Visual styling is separate from writes. |
| file writer | Create only missing template files and skip existing files. | `internal/core/docstore` | `.ni/plan.lock.json` is never modified. |
| summary printer | Print plain text created/skipped/unchanged and next commands after TUI exits. | `cmd/ni/main.go` | Summary is not a readiness claim. |

## Bubble Tea v2 / Lip Gloss v2 usage

| Requirement | Implemented path | Pass? | Notes |
| --- | --- | --- | --- |
| Bubble Tea v2 import | `internal/tui/init/model.go` | Yes | Uses `charm.land/bubbletea/v2`. |
| Lip Gloss v2 import | `internal/tui/init/model.go` | Yes | Uses `charm.land/lipgloss/v2`. |
| model / Init / Update / View | `internal/tui/init/model.go` | Yes | `View()` returns `tea.View`. |
| AltScreen | `internal/tui/init/model.go` | Yes | Declared through `view.AltScreen = true`. |
| keyboard navigation | `internal/tui/init/model.go`, tests | Yes | Up, Down, Left, Right, Enter, Esc, and q are handled. |
| post-TUI plain text summary | `cmd/ni/main.go` | Yes | Printed after TUI exits. |
| non-interactive fallback | `cmd/ni/main.go`, tests | Yes | Existing CI/non-TTY behavior remains. |
| domain/render separation | `internal/core/docstore`, `internal/tui/init` | Yes | View renders only; domain writes. |

## TUI behavior

| Behavior | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `ni init .` | Launch guided TUI on an interactive TTY. | Command layer routes TTY positional init to TUI. | Yes | Non-TTY falls back. |
| current directory target | Use current directory when target is `.`. | Existing CLI tests cover `ni init . --yes`. | Yes | TUI uses same dir parsing. |
| interactive TUI | Step-based wizard with review before write. | Model has field stage and confirm stage. | Yes | No file write happens in `View`. |
| non-interactive fallback | Work in CI and pipes. | Existing fallback retained. | Yes | Plain output only. |
| existing file protection | Do not silently overwrite. | Domain plan marks existing files as skipped. | Yes | TUI offers missing-only, keep, abort. |
| lockfile protection | Do not modify `.ni/plan.lock.json`. | Domain and command layer return without writes. | Yes | Tested. |
| cancel | Write nothing before confirmation. | TUI returns canceled result. | Yes | Tested at model level. |
| confirm/write | Confirm returns intent; command writes through domain. | TUI produces intent and command calls docstore writer. | Yes | Tested at model/domain layers. |
| plain summary | Print target, created, skipped/unchanged, next commands. | Command layer prints summary after write or cancel. | Yes | Not styled as TUI. |

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
| Homebrew | Planned / v0.5 candidate only. | README and docs preserve Planned wording. | Yes | No Available claim added. |
| Windows real-host execution | Deferred unless transcript exists. | README preserves deferred wording. | Yes | No Windows verified claim added. |
| `ni init .` | Guided setup only. | TUI initializes planning artifacts only. | Yes | No readiness decision. |
| `ni run` | Bounded prompt compilation only. | README and code preserve non-execution boundary. | Yes | No runtime execution added. |
| READY | CLI status only. | README states `ni status` decides readiness. | Yes | TUI does not claim READY. |
| Model workspace packs | Experimental. | Status list preserves Experimental. | Yes | Skills are UX; CLI is authority. |
| No-terminal | Experimental / assisted. | Status list preserves boundary. | Yes | Not deterministic validation. |
| Runtime execution boundary | Kernel remains pre-runtime. | No shell/Codex adapter or queue added. | Yes | No downstream execution behavior. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| `README.md` | Modified | Yes | Redesigned public README. |
| `README.ko.md` | Modified | Yes | Companion update. |
| `cmd/ni/*` | Modified | Yes | Init routing and summary printer. |
| `internal/*` | Modified/added | Yes | Domain file plan and TUI package. |
| `go.mod` | Modified | Yes | Bubble Tea v2 and Lip Gloss v2 dependencies. |
| `go.sum` | Added | Yes | Module checksums. |
| `docs/22_INSTALL.md` | Unchanged | No | Existing markers remain accurate. |
| `docs/install-curl*` | Unchanged | No | Existing markers remain accurate. |
| `docs/124*` | Added | Yes | This implementation report and Korean companion. |
| install scripts | Unchanged | No | No release/publish work. |
| suspicious files removed | Deleted | Yes | `* 2.md` duplicates removed. |
| `.ni/contract.json` | Unchanged | No | Protected. |
| `.ni/session.json` | Unchanged | No | Protected. |
| `.ni/plan.lock.json` | Unchanged | No | Protected. |
| unexpected files | None expected | No | Recheck before commit. |

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
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build returned `0.0.0-dev`. |
| Non-interactive `ni init --dir /private/tmp/ni-init-smoke-203-a --yes` | Passed; planning files created and no lockfile/downstream artifact was created. |
| Repeated non-interactive init in the same temp directory | Passed; existing files were reported as unchanged. |
| Lockfile safety temp init | Passed; `.ni/plan.lock.json` was not modified and no contract was created. |
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

- Redesigned `README.md` and `README.ko.md`.
- Added Bubble Tea v2 / Lip Gloss v2 TUI package under `internal/tui/init`.
- Updated `cmd/ni/main.go` to route interactive init through the TUI and keep
  non-interactive fallback.
- Added domain file-plan and lockfile protection checks in `internal/core/docstore`.
- Added model/domain tests for init behavior.
- Deleted accidental `examples/namba-ai-upgrade/docs/plan/* 2.md` duplicates.
- Added this document and Korean companion.

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
