# v0.6.0 Public Install Parity and Release Readiness

## Current status

State:
- v0.5.1 release: published and verified
- v0.6.0 release: not published
- Namba Intent rename: implemented in current tree
- primary command: namba-intent
- deprecated ni shim: transition-only
- .ni/ compatibility: preserved
- public install retrieval of namba-intent: not verified until v0.6.0 release
- Homebrew: Planned / v0.5 candidate
- Windows real-host verification: pending
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

## Audit goal

This audit checks whether the current-tree rename is ready for v0.6.0 release
preparation and whether public install docs are safely bounded. It does not
publish, tag, upload assets, run release workflows, run GoReleaser publish, run
`namba-intent end` on the project root, relock the project root, execute
generated prompts, create Homebrew material, or add downstream execution
behavior.

## Decision

V0_6_0_RELEASE_READINESS_READY_WITH_NOTES

Rationale: current-tree command rename, installer configuration, release asset
naming, `.ni/` compatibility, deprecated `ni` shim behavior, and non-execution
boundaries are ready for v0.6.0 release preparation. Public v0.6.0 install
retrieval, hosted artifacts, Windows real-host execution, Homebrew availability,
and external user validation remain explicit future evidence gates.

## Current-tree rename readiness

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| product name | Namba Intent | README, CLI help, and docs use Namba Intent for current-tree surfaces. | Yes | v0.6.0 is not published. |
| primary command | `namba-intent` | `cmd/namba-intent` exists and delegates to shared CLI logic. | Yes | `go run ./cmd/namba-intent --help` passed. |
| cmd/namba-intent | primary entrypoint | `cmd/namba-intent/main.go` calls `internal/cli.Run`. | Yes | No duplicated command logic. |
| deprecated ni shim | warn and delegate | `cmd/ni/main.go` prints `ni is deprecated; use namba-intent.` and delegates. | Yes | Transition-only. |
| internal CLI sharing | shared implementation | `internal/cli` owns command dispatch and command-name option. | Yes | Shim passes `CommandName: "ni"` only for compatibility output. |
| .ni compatibility | preserve `.ni/` | Init/status/lock/session paths still use `.ni/`. | Yes | No `.namba-intent` directory introduced. |
| namba-ai distinction | do not claim `namba` | docs/135 and docs/136 distinguish NambaAI `namba` from Namba Intent `namba-intent`. | Yes | Repo remains `Nam-Cheol/ni`. |
| no downstream execution | run compiles prompt only | README, skills, and CLI prompt wording preserve prompt-compiler-only boundary. | Yes | No shell/Codex adapter added. |

## Public install boundary

| Surface | Expected boundary | Observed state | Pass? | Required action |
| --- | --- | --- | --- | --- |
| README.md | Separate current main from latest public v0.5.1. | States v0.6.0 is upcoming and v0.5.1 may still use `ni`. | Yes | Keep note until v0.6.0 is verified. |
| README.ko.md | Korean companion must not widen claims. | Matches README boundary. | Yes | Keep in sync. |
| install.sh | Future primary binary is `namba-intent`. | Selects `namba-intent_<version>` assets and installs `namba-intent`. | Yes | Do not claim public retrieval before v0.6.0. |
| install.ps1 | Future primary binary is `namba-intent.exe`. | Selects `namba-intent_<version>_windows_amd64.zip` and installs under `%LOCALAPPDATA%\namba-intent\bin`. | Yes | Real Windows transcript still pending. |
| docs/22_INSTALL.md | Public v0.5.1 and current-main v0.6.0 paths must be separated. | Updated to mark public `namba-intent` installer retrieval as Release-gated. | Yes | Reverify after v0.6.0 publication. |
| docs/install-curl.md | Curl docs must not claim public `namba-intent` retrieval. | Updated to separate historical v0.5.1 `ni` evidence from future `namba-intent` assets. | Yes | Reverify after hosted assets exist. |
| public latest release | Latest public release may still be v0.5.1. | `git tag --list v0.5.1` returned `v0.5.1`; `v0.6.0` tag absent. | Yes | Do not treat current tree as published release. |
| v0.6.0 release status | Not published. | No local `v0.6.0` tag found. | Yes | Publication is a future human-approved task. |

## Installer readiness

| Installer | Expected future v0.6.0 behavior | Current-tree state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| install.sh | Install `namba-intent` from `namba-intent_<version>_<os>_<arch>` archive. | Implemented. | Yes | Validation uses local fake release assets. |
| install.ps1 | Install `namba-intent.exe` from `namba-intent_<version>_windows_amd64.zip`. | Implemented. | Yes | Static check only on macOS host. |
| uninstall | Remove installer-managed primary binary and PATH entry. | Unix and PowerShell uninstall target `namba-intent` names. | Yes | No public uninstall was run. |
| verification commands | Print `namba-intent --help` and `namba-intent version`. | Implemented in both installers. | Yes | Matches README current-tree command. |
| legacy ni handling | Keep shim transition-only. | Source-tree shim exists; installer primary path does not rely on `ni`. | Yes | Windows `ni` alias cleanup is historical only. |
| Windows alias handling | Do not rely on PowerShell `ni`. | `install.ps1` says alias cleanup is not required for `namba-intent.exe`. | Yes | Real-host Windows verification pending. |

## Release tooling readiness

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| version injection | Inject `ni/internal/version.Version`. | `.goreleaser.yaml` uses `-X ni/internal/version.Version={{ .Version }}`. | Yes | Source `go run` still reports `0.0.0-dev`. |
| asset names | Future artifacts use `namba-intent_...`. | `.goreleaser.yaml`, install scripts, and release pipeline docs use `namba-intent_<version>...`. | Yes | Hosted assets do not exist yet. |
| checksum names | Future checksum file is `namba-intent_<version>_checksums.txt`. | Config and installers match. | Yes | Public v0.6.0 checksum not published. |
| release-check | Check-only release gate. | `scripts/release-check.sh` checks tests, quality, smoke, install, docs, and release boundaries. | Yes | Does not publish. |
| install-check | Source/build/temp install gate. | `scripts/install-check.sh` verifies `namba-intent` command-name path and installer tests. | Yes | Uses temp paths. |
| GoReleaser config | Primary command plus shim. | `.goreleaser.yaml` builds `namba-intent` and `ni` shim in one archive. | Yes | GoReleaser publish not run. |
| GitHub workflow | Tag-only release workflow. | `.github/workflows/release.yml` runs on `v*` tags and calls GoReleaser. | Yes | Workflow not run in this task. |

## Docs and skills audit

| Surface | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| README | Current-tree Namba Intent, public v0.5.1 distinction. | Bounded. | Yes | No v0.6.0 publication claim. |
| README.ko | Korean companion. | Bounded. | Yes | Does not promise more than English. |
| current docs | Use Namba Intent and `namba-intent` for current/future surfaces. | docs/22, install-curl, docs/67, docs/120 updated for current primary command. | Yes | Historical v0.5.x docs preserve `ni`. |
| historical docs | Preserve actual past commands. | v0.5.1 release evidence remains `ni`. | Yes | This is intentional. |
| Claude skills | Skills are UX; CLI is authority. | Package README and skill docs use `namba-intent` authority while preserving `ni-*` skill IDs. | Yes | Broad status remains Experimental. |
| Codex skills | Skills are UX; CLI is authority. | Package README and `.agents/skills` use `namba-intent` command examples. | Yes | No skill replaces CLI gates. |
| .agents skills | Repo-local UX only. | `ni-start`, `ni-grill`, `ni-end`, and `ni-run` preserve CLI authority. | Yes | No downstream execution behavior. |

## Current-tree evidence

| Evidence | Result | Notes |
| --- | --- | --- |
| docs/136 implementation | Pass | Records rename implementation and validation. |
| docs/137 smoke | Pass with notes | Records command-name help/version, init/status, repeated init, lockfile safety, and shim delegation. |
| go test | Pass | `GOCACHE=/private/tmp/ni-go-cache go test ./...`. |
| install docs check | Pass | `python3 scripts/check-install-docs.py`. |
| install ps1 check | Pass | `python3 scripts/check-install-ps1.py`. |
| quality | Pass | `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh`. |
| release-check | Pass | `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh`. |
| protected .ni diff | Pass | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` empty. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks v0.6.0 readiness? |
| --- | --- | --- | --- |
| v0.6.0 publication | This task is non-publishing. | Human-approved tag, release workflow, release metadata. | No, blocks publication claim only. |
| public install retrieval of namba-intent | Hosted v0.6.0 assets do not exist yet. | Isolated install from published v0.6.0 assets plus help/version proof. | No, blocks public install claim only. |
| hosted artifacts | No release action performed. | Asset inventory and checksum verification. | No, blocks hosted artifact claim only. |
| Windows real-host execution | macOS host cannot prove it. | Windows PowerShell install, new-session help/version, uninstall transcript. | No, unless Windows verified claim is required. |
| Homebrew Available | No tap/formula/install proof. | Tap/formula, checksums, audit, install, `namba-intent --help`, `namba-intent version`. | No, because Homebrew remains Planned. |
| external user validation | No external tester transcript. | Tester transcript and comprehension review. | No, known note. |
| model workspace host behavior | Global host/provider behavior not verified. | Host-specific install/discovery/runtime proof. | No, because Experimental is preserved. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.6.0 publication status | Must say not published. | Preserved. | Yes | No tag or release action. |
| public install | Must not claim public `namba-intent` retrieval. | docs/22 and curl docs mark it Release-gated. | Yes | v0.5.1 evidence remains `ni`. |
| Namba Intent identity | Current tree may use it. | Used in README/help/docs. | Yes | Release claim remains future-gated. |
| namba-intent command | Primary current-tree command. | Implemented and validated. | Yes | Public install still not verified. |
| deprecated ni shim | Transition-only. | Warns and delegates. | Yes | Not primary path. |
| namba-ai distinction | Do not use `namba`. | Preserved in docs/135 and docs/136. | Yes | Repo name still `Nam-Cheol/ni`. |
| Homebrew | Planned / v0.5 candidate. | Preserved. | Yes | No Available claim. |
| Windows real-host verification | Pending. | Preserved. | Yes | Static checks only. |
| run behavior | Prompt compilation only. | Preserved. | Yes | No prompt execution. |
| runtime execution boundary | No task runner, SPEC runner, shell/Codex adapter, queue, PR/release automation, downstream execution layer. | Preserved. | Yes | No runtime behavior added. |

## Git status / inclusion check

| Path or group | git status --short | Expected in v0.6.0? | Notes |
| --- | --- | --- | --- |
| README.md | clean | Yes | Bounded current-tree surface already tracked. |
| README.ko.md | clean | Yes | Korean companion already tracked. |
| cmd/namba-intent/* | clean | Yes | Primary command tracked. |
| cmd/ni/* | clean | Yes | Deprecated shim tracked. |
| internal/cli/* | clean | Yes | Shared implementation tracked. |
| install.sh | clean | Yes | Future `namba-intent` installer tracked. |
| install.ps1 | clean | Yes | Future Windows installer tracked. |
| docs/135* | clean | Yes | Rename plan tracked. |
| docs/136* | clean | Yes | Rename implementation tracked. |
| docs/137* | untracked before this task | Yes | First-user smoke evidence from prior current-tree smoke. |
| docs/138* | added in this task | Yes | This audit pair. |
| scripts/* | modified | Yes | Install docs checker updated for Release-gated current-main installer. |
| packages/* | modified | Yes | Claude package authority wording aligned. |
| .agents/* | clean | Yes | Existing repo-local skills remain. |
| .ni/contract.json | no diff | Yes | Protected. |
| .ni/session.json | no diff | Yes | Protected. |
| .ni/plan.lock.json | no diff | Yes | Protected. |
| unexpected files | none observed beyond docs/137 untracked before this task | No | No generated prompts executed. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; docs/137 existed as untracked before this task; docs/138 and bounded docs/checker edits added in this task. |
| `git log --oneline --decorate -20` | Passed; HEAD `80acd80 Implement Namba Intent rename`. |
| `git tag --list v0.5.1` | Passed; `v0.5.1` exists. |
| `git tag --list v0.6.0` | Passed; empty. |
| `git rev-parse v0.5.1` | Passed; `b588f6b2e13111841081d186bd0e70d3c0bfbd6c`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/namba-intent --help` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/namba-intent version` | Passed; source output `0.0.0-dev`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; deprecation warning verified. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff --check` | Passed. |

## Changes made

- Updated `docs/22_INSTALL.md` to separate verified v0.5.1 `ni` evidence from
  release-gated current-main `namba-intent` installer behavior.
- Updated `docs/install-curl.md` and `docs/install-curl.ko.md` with the same
  public install boundary.
- Updated `docs/120_GLOBAL_INSTALL_ACCEPTANCE.md` and Korean companion to use
  `namba-intent` as the primary global install acceptance command.
- Updated `docs/67_RELEASE_PIPELINE.md` and Korean companion to match the
  actual GoReleaser primary binary plus shim configuration.
- Updated `packages/claude-skills/README.md` to name `namba-intent` as CLI
  authority.
- Updated `scripts/check-install-docs.py` to enforce the new Release-gated
  current-main installer boundary.
- Added this audit and Korean companion.

## What this audit proves

State only:
- current-tree rename readiness is ready with notes;
- public install docs are bounded after this audit;
- next release preparation can proceed;
- no release action was performed.

## What this audit does not prove

State:
- v0.6.0 has been published;
- public install retrieves namba-intent;
- hosted v0.6.0 artifacts exist;
- Windows real-host execution works;
- Homebrew is Available;
- external users succeed;
- downstream execution succeeds.

## Recommended next task

A. v0.6.0 release notes draft

Selection rationale: readiness passes with notes, and the remaining proof gaps
are expected release-time or external verification gates. Release notes are the
next safest preparation artifact before artifact dry-run or publication.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: draft v0.6.0 Namba Intent release notes without publishing.

Use docs/135_NAMBA_INTENT_RENAME_PLAN.md,
docs/136_NAMBA_INTENT_RENAME_IMPLEMENTATION.md,
docs/137_NAMBA_INTENT_FIRST_USER_SMOKE.md, and
docs/138_V0_6_0_PUBLIC_INSTALL_PARITY_AND_RELEASE_READINESS.md as source
evidence.

Goal:
Create v0.6.0 release notes that explain the Namba Intent rename, primary
`namba-intent` command, transition-only `ni` shim, preserved `.ni/`
compatibility, installer and release asset rename, and prompt-compiler-only
runtime boundary.

Required boundaries:
- Do not publish, tag, create a GitHub release, upload assets, run release
  workflows, run GoReleaser publish, create or publish Homebrew formula, mark
  Homebrew Available, claim Windows real-host verification, run ni end or
  namba-intent end on the project root, relock the project root, edit
  .ni/plan.lock.json, execute generated prompts, or add downstream execution
  behavior.
- State that v0.6.0 is not published unless this task explicitly verifies a
  new tag/release created outside the task.
- State that public install retrieval of `namba-intent` is not verified until
  v0.6.0 artifacts are published and checked.
- Preserve Homebrew: Planned / v0.5 candidate, Model workspace packs:
  Experimental, No-terminal method: Experimental / assisted, and Skills are
  UX; CLI is authority.

Deliverables:
- Add docs/139_V0_6_0_RELEASE_NOTES_DRAFT.md.
- Add docs/139_V0_6_0_RELEASE_NOTES_DRAFT.ko.md.
- Include sections for summary, included changes, migration notes, install
  boundary, validation evidence, known deferrals, what this does not prove,
  and recommended next task.

Validation:
- git status --short
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
- git diff --check

Final response:
Report changed files, release-note decision, public install boundary, validation
results, protected .ni diff result, confirmation that no publication/tag/release
action occurred, and the selected next task.
```
