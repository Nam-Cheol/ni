# v0.5.1 Release Candidate Validation

## Current status

State:
- v0.5.0 publication: verified.
- Public install parity decision: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- v0.5.1 patch plan decision: V0_5_1_PATCH_PLAN_READY_WITH_NOTES.
- current-tree first-user smoke after TUI: FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES.
- v0.5.1 release: not published.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host execution: deferred on the macOS-only development host.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

This validates current-tree readiness for a v0.5.1 release candidate without
publishing it. It does not tag, publish, upload assets, create a GitHub Release,
run release workflows, create or publish a Homebrew formula, run `ni end` on the
project root, relock the project root, execute generated prompts, or add runtime
execution behavior.

## Decision

V0_5_1_RC_VALIDATION_PASS_WITH_NOTES

Justification: the audited current tree covers the v0.5.1 install-parity patch
scope and the local validation gates passed. Notes remain because v0.5.1 has not
been published, release artifacts do not exist yet, `install.sh` has not
retrieved v0.5.1, Windows real-host execution is still deferred, Homebrew remains
Planned / v0.5 candidate, and external user validation has not been performed.

## Patch scope validation

| Scope item | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- |
| `ni init .` positional target | Current-directory init works. | `cmd/ni/main.go` help and temp binary smoke accept `ni init . --yes`. | Yes | Published v0.5.0 lacks this path; current tree has it. |
| Bubble Tea v2 TUI | Interactive guided init uses Bubble Tea v2. | `internal/tui/init/model.go` imports `charm.land/bubbletea/v2`. | Yes | Interactive launch is routed only where safe. |
| Lip Gloss v2 styling | TUI view uses Lip Gloss v2 styling. | `internal/tui/init/model.go` imports `charm.land/lipgloss/v2` and renders styled views. | Yes | Styling remains render-only. |
| non-interactive fallback | CI/non-TTY still works. | `ni init . --yes` passed in a temporary project. | Yes | No TUI requirement for automation. |
| domain/render separation | File writing lives outside TUI rendering. | `internal/core/docstore` owns init file planning and writes; `internal/tui/init` owns model/view. | Yes | Supports deterministic tests. |
| existing file protection | Init must not silently overwrite planning files. | Repeated `ni init . --yes` reported existing/unchanged files and created no files. | Yes | This is additive-only behavior. |
| lockfile protection | Init must not modify `.ni/plan.lock.json`. | Sentinel lockfile SHA-256 stayed `1ede15b645ba3dab7c56783753a6420477a7ec58847c3e80acb2d0fa792d5960`. | Yes | The command stopped with a lockfile warning. |
| post-TUI plain summary | Init should report what happened after guided setup. | `cmd/ni/main.go` prints plain init summaries; tests assert summary text. | Yes | `ni init . --yes` output includes created/unchanged files and next commands. |
| README onboarding | macOS and Windows first paths are prominent and bounded. | README and README.ko show two primary install paths plus `ni init .`. | Yes | They preserve the public v0.5.0 parity note. |
| install docs/checkers | Docs and validators align with current behavior. | `python3 scripts/check-install-docs.py` and `python3 scripts/check-install-ps1.py` passed. | Yes | Windows checker is static safety only. |
| no downstream execution | Patch must not make `ni run` execute downstream work. | README/docs and help keep `ni run` as bounded prompt compilation. | Yes | No runtime execution behavior was added in this validation. |

## Repository state

| Surface | Observed state | Notes |
| --- | --- | --- |
| git status | `README.md` and `README.ko.md` modified; docs/126, docs/127, and docs/128 untracked after this task. | No staging or commit performed. |
| v0.5.0 tag | Present; `git rev-parse v0.5.0` returned `b8fec7fa9615a861acf4eba59733c800c70c6cca`. | Public mismatch baseline. |
| v0.5.1 tag | Absent. | Correct for RC validation; no tag was created. |
| diff `v0.5.0..HEAD` | 60 files changed, 7302 insertions, 476 deletions before docs/128. | Includes current-tree init, installer, docs, and checker work. |
| protected `.ni` diff | Empty. | `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` unchanged in the project root. |
| generated artifacts | No unexpected repo artifact was created by this validation. | Temporary binary/project lived under `/private/tmp`. |

docs/126 and docs/127 are currently untracked release-planning evidence. They
are expected in the v0.5.1 RC planning bundle but were not staged here.

## Current-tree first-user validation

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| command-name `ni --help` | `PATH=/private/tmp/ni-v051-rc-bin.3tTS9H:... zsh -f -c 'ni --help'` | PATH-resolved command works and lists `ni init [.]`. | Passed. | Yes | `command -v ni` resolved to the temporary binary. |
| command-name `ni version` | `PATH=/private/tmp/ni-v051-rc-bin.3tTS9H:... zsh -f -c 'ni version'` | Source/current-tree binary prints a version. | `0.0.0-dev`. | Yes | Release builds must inject `0.5.1`. |
| `ni init . --yes` | Temp project `ni init . --yes`. | Creates planning workspace. | Created `.ni/contract.json`, `.ni/session.json`, and `docs/plan/**`; did not create `.ni/plan.lock.json`. | Yes | No downstream prompt was executed. |
| `ni status --proof --next-questions` | Temp project status. | Reports actual first-run readiness. | `NI Intent Readiness: BLOCKED` with R014, OQ-001, R015, and R016. | Yes | BLOCKED is expected for an empty first-run scaffold. |
| repeated init | Temp project `ni init . --yes` again. | Does not silently overwrite. | Reported existing/unchanged files and `created files: none`. | Yes | Existing planning docs preserved. |
| lockfile safety | Create sentinel `.ni/plan.lock.json`, then run `ni init . --yes`. | Init must not modify lockfile. | Warning printed; no files changed; sentinel hash unchanged. | Yes | This used a temporary project, not the repo root. |

## TUI model coverage

| Behavior | Evidence | Pass? | Notes |
| --- | --- | --- | --- |
| model | `internal/tui/init/model.go`; `internal/tui/init/model_test.go`. | Yes | Model state is directly unit-tested. |
| Init | `TestModelInitialStateUsesAltScreen`. | Yes | Initial state exposes alt-screen view. |
| Update | `updateForTest` drives key messages through the model. | Yes | Covered through navigation and confirmation tests. |
| View | `m.View()` is checked for rendered content. | Yes | View output is not treated as readiness authority. |
| AltScreen | `view.AltScreen` asserted in tests. | Yes | TUI behavior only. |
| Lip Gloss v2 | `charm.land/lipgloss/v2` import and styled view code. | Yes | Evidence is code-level plus package tests. |
| Up/Down | Key tests cover `tea.KeyUp` and `tea.KeyDown`. | Yes | Cursor movement covered. |
| Left/Right | Key tests cover `tea.KeyLeft` and `tea.KeyRight`. | Yes | Horizontal movement covered. |
| Enter | Key tests cover `tea.KeyEnter`. | Yes | Select and confirm paths covered. |
| Esc | Key tests cover `tea.KeyEsc`. | Yes | Back/cancel path covered. |
| q | Key tests cover `tea.KeyRunes` with `q`. | Yes | Cancel path covered. |
| cancel | `TestModelCanCancelAtConfirm` and `q` path. | Yes | Canceled result asserted. |
| confirm | `TestModelCanConfirmIntent`. | Yes | Confirmed result asserted. |
| post-TUI summary | `cmd/ni/main_test.go` summary assertions and `confirmGuidedIntent` summary. | Yes | Plain text summary exists outside TUI rendering. |

## README/install docs audit

| Surface | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| README.md | macOS/Windows primary paths, Homebrew not Available, `ni run` compile-only, READY not product readiness, TUI not authority. | Matches; includes v0.5.0 parity note and status list. | Yes | `ni init .` is guided project intent setup. |
| README.ko.md | Korean companion must not promise more than English. | Matches English boundaries. | Yes | Preserves required control words. |
| docs/22_INSTALL.md | Install behavior and uninstall guidance aligned. | Check passed. | Yes | Windows real-host execution remains deferred. |
| docs/install-curl.md | Curl installer guidance bounded to available release assets and command-name verification. | Check passed. | Yes | Does not claim v0.5.1 retrieval before release. |
| docs/install-curl.ko.md | Korean companion bounded like English. | Check passed. | Yes | Homebrew remains outside curl path. |

## Version/release gate

| Gate | Required evidence before release | Current state | Blocks RC? | Notes |
| --- | --- | --- | --- | --- |
| release version injection | Release binary reports `0.5.1`. | `.goreleaser.yaml` injects `ni/internal/version.Version={{ .Version }}`. | No | Must be verified in artifact dry-run/release build. |
| v0.5.1 tag | Tag exists at intended release commit. | Absent. | No | Correct for pre-publication RC validation. |
| artifact build | Archives and checksums generated for v0.5.1. | Not run. | No | Next uncertainty. |
| artifact version output | Extracted artifact `ni version` reports `0.5.1`. | Not available yet. | No | Required before publication. |
| `install.sh` v0.5.1 retrieval | Installer retrieves v0.5.1 after release. | Not possible before publication. | No | Must not be claimed now. |
| curl installer isolated install | Temporary install runs `ni --help` and `ni version`. | Current-tree binary smoke only; no hosted v0.5.1 install. | No | Required after artifact exists. |
| checksum verification | `ni_0.5.1_checksums.txt` matches artifacts. | Not available yet. | No | Artifact dry-run should verify. |

## v0.5.1 release notes delta

Draft delta:
- Fix: public install parity mismatch from v0.5.0 where installed `ni --help`
  and `ni version` passed but `ni init .` failed.
- Add: `ni init .` positional target support.
- Add: Bubble Tea v2 / Lip Gloss v2 guided init TUI.
- Add: non-interactive fallback preservation for CI and non-TTY contexts.
- Add: existing file and `.ni/plan.lock.json` protection during init.
- Improve: README macOS / Windows two-path onboarding.
- Improve: install docs and checker alignment for command-name verification.
- Does not add downstream execution.
- Does not make Homebrew Available.
- Does not verify Windows real-host execution.
- Does not make no-terminal deterministic.

This delta is draft-only. It must not be copied into published release notes
until the actual v0.5.1 artifacts and installer path are verified.

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks v0.5.1 RC? |
| --- | --- | --- | --- |
| Windows real-host execution | Current development host is macOS-only. | Windows install/new-session/help/version/init/uninstall transcript. | No |
| Homebrew Available | No published tap/formula/install proof. | Tap, formula, checksum, audit, install, `ni --help`, `ni version`, uninstall proof. | No |
| external user validation | No external user transcript in this task. | User or separate-host install/init transcript. | No |
| model workspace host behavior | Host-level/global install and provider behavior remain unverified. | Host-specific discovery/install/provider transcript. | No |
| no-terminal deterministic validation not claimed | No-terminal remains Experimental / assisted. | Trusted CLI transcript for target workspace. | No |
| README raster images | Raster image generation was out of scope. | Separate asset task if still useful. | No |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.5.1 publication status | Must say not published. | Preserved. | Yes | No release action occurred. |
| published v0.5.0 behavior | Must distinguish installed help/version pass from `ni init .` failure. | Preserved in README/docs/126/docs/127. | Yes | Public mismatch is the reason for v0.5.1. |
| current-tree behavior | May claim local current-tree support only. | Current-tree binary smoke passed. | Yes | Not a hosted artifact claim. |
| Homebrew | Must remain Planned / v0.5 candidate. | Preserved. | Yes | No Homebrew Available claim. |
| Windows real-host execution | Must remain deferred without transcript. | Preserved. | Yes | Static installer checks only. |
| `ni init .` | Guided setup only. | Creates planning docs and `.ni` skeleton. | Yes | Does not run agents. |
| `ni run` | Bounded prompt compilation only. | README/help/docs preserve compile-only wording. | Yes | Must not execute downstream work. |
| READY | CLI readiness only, not product readiness. | Root `ni status` reports READY; docs keep CLI authority. | Yes | Model did not declare readiness independently. |
| TUI readiness boundary | TUI may collect intent, not decide readiness. | TUI model writes no readiness result. | Yes | `ni status` remains the gate. |
| runtime execution boundary | `ni-kernel` must not include task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, or downstream execution layer. | Preserved. | Yes | No runtime behavior was added. |

## Git status / inclusion check

| Path or group | `git status --short` | Expected in v0.5.1? | Notes |
| --- | --- | --- | --- |
| README.md | `M README.md` | Yes | Public parity note and onboarding. |
| README.ko.md | `M README.ko.md` | Yes | Korean companion. |
| cmd/ni/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Contains positional init and tests. |
| internal/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Contains docstore/TUI support. |
| go.mod | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Bubble Tea v2 / Lip Gloss v2 dependencies. |
| go.sum | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Dependency checksums. |
| install.sh | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Current install path handling. |
| install.ps1 | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Windows User PATH installer. |
| docs/124* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Init TUI evidence. |
| docs/125* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | First-user smoke evidence. |
| docs/126* | `?? docs/126_*` | Yes | Public v0.5.0 mismatch evidence. |
| docs/127* | `?? docs/127_*` | Yes | v0.5.1 patch plan. |
| docs/128* | `?? docs/128_*` | Yes | This RC validation and Korean companion. |
| scripts/* | clean relative to HEAD; changed in `v0.5.0..HEAD` | Yes | Install/release/checker alignment. |
| `.ni/contract.json` | no diff | No direct edit | Protected. |
| `.ni/session.json` | no diff | No direct edit | Protected. |
| `.ni/plan.lock.json` | no diff | No direct edit | Protected. |
| unexpected files | none observed in repo status | No | Temporary files are outside repo. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; expected modified/untracked docs only. |
| `git log --oneline --decorate -20` | Passed; post-v0.5.0 onboarding/install commits visible. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git tag --list v0.5.1` | Empty; no v0.5.1 tag exists. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff --name-only v0.5.0..HEAD` | Passed; candidate files listed. |
| `git diff --stat v0.5.0..HEAD` | Passed; 60 files, 7302 insertions, 476 deletions before docs/128. |
| required ripgrep scans | Passed; no scope-expanding runtime/release claim found that blocks RC validation. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root reported `NI Intent Readiness: READY` with no blockers, deferrals, or warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help includes `ni init [.]`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build reports `0.0.0-dev`. |
| temporary current-tree binary build | Passed; `/private/tmp/ni-v051-rc-bin.3tTS9H/ni`. |
| command-name temp `ni --help` | Passed. |
| command-name temp `ni version` | Passed; `0.0.0-dev`. |
| temp `ni init . --yes` | Passed. |
| temp `ni status --proof --next-questions` | Passed; command ran and correctly reported BLOCKED for blank first-run intent. |
| repeated temp `ni init . --yes` | Passed; no overwrites. |
| temp lockfile safety | Passed; sentinel hash unchanged. |
| TUI model behavior tests | Passed through `go test ./...`, including `ni/internal/tui/init`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed after docs/128 addition. |
| `bash scripts/demo-check.sh` | Passed after docs/128 addition. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed after docs/128 addition. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed after docs/128 addition. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed after docs/128 addition. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed after docs/128 addition. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

No Go files were touched by this RC validation doc task, so `gofmt -w .` was
not required for the final edited set.

## Changes made

| File | Why |
| --- | --- |
| `docs/128_V0_5_1_RELEASE_CANDIDATE_VALIDATION.md` | Added English RC validation, decision, evidence, release notes delta, and next task prompt. |
| `docs/128_V0_5_1_RELEASE_CANDIDATE_VALIDATION.ko.md` | Added Korean companion with the same boundaries. |
| `docs/51_POST_RELEASE_ROADMAP.md` | Added a narrow pointer to docs/128. |
| `docs/51_POST_RELEASE_ROADMAP.ko.md` | Added the matching Korean pointer. |

## What this validation proves

- The current tree appears ready for a v0.5.1 RC under the audited criteria.
- The first-user current-tree path works with notes: a blank scaffold correctly
  remains BLOCKED until real intent is supplied.
- The planned v0.5.1 patch scope is covered in the current tree.
- Release gates are explicit and still pending for actual artifacts.
- No release action was performed.

## What this validation does not prove

- v0.5.1 has been published.
- v0.5.1 artifacts exist.
- `install.sh` retrieves v0.5.1.
- Windows real-host execution works.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.

## Recommended next task

Selected next task: A. v0.5.1 artifact dry-run.

Selection rationale: RC validation passes with notes, and the next uncertainty
is whether the release artifact path injects `0.5.1`, produces consistent
checksums, and preserves install parity before any publication action.

## Next executable Codex prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Goal:
Run a v0.5.1 artifact dry-run for the current release-candidate tree. This is
pre-publication validation only.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md
- docs/127_V0_5_1_PATCH_RELEASE_PLAN.md
- docs/128_V0_5_1_RELEASE_CANDIDATE_VALIDATION.md
- docs/22_INSTALL.md
- docs/install-curl.md
- install.sh
- install.ps1
- .goreleaser.yaml
- .github/workflows/release.yml
- cmd/ni/
- internal/
- scripts/release-check.sh

Required checks:
- git status --short
- git tag --list v0.5.1
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- Check whether goreleaser is installed.
- If goreleaser is installed, run a local snapshot/dry-run artifact build only,
  inspect generated archive names, checksums, and extracted current-platform
  `ni --help` / `ni version` output.
- If goreleaser is not installed, do not install it automatically; document the
  unavailable command and run the closest local `go build` linker-flag check
  that proves `ni version` can report `0.5.1`.
- Build a temporary current-platform binary outside the repo with
  `-ldflags "-X ni/internal/version.Version=0.5.1"` and verify command-name
  `ni --help`, `ni version`, `ni init . --yes`, repeated init, and lockfile
  protection in a temporary project.

Decision:
Use exactly one:
- V0_5_1_ARTIFACT_DRY_RUN_PASS
- V0_5_1_ARTIFACT_DRY_RUN_PASS_WITH_NOTES
- V0_5_1_ARTIFACT_DRY_RUN_BLOCKED

Rules:
- Do not publish, tag, create a GitHub Release, upload assets, run release
  workflows, run goreleaser publish, create or publish a Homebrew formula, run
  ni end on the project root, relock the project root, execute generated
  prompts, add runtime execution behavior, or mark v0.5.1 as released.
- Do not mark Homebrew Available.
- Do not claim Windows real-host execution verified unless a Windows transcript
  exists.
- Do not claim `install.sh` retrieves v0.5.1 unless a real hosted v0.5.1 asset
  exists and the installer retrieves it.
- Keep Skills are UX; CLI is authority.
- Keep `ni run` as bounded prompt compilation only.

Final response:
Report changed files, artifact dry-run decision, version-injection result,
checksum/archive result if available, current-platform binary smoke, known
deferrals, validation results, protected .ni diff, and confirmation that no
publish/tag/release/upload/project-root relock/generated prompt execution
occurred.
```
