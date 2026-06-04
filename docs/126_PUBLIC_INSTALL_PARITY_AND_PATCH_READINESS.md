# Public Install Parity and Patch Readiness

## Current status

State:
- v0.5.0 publication: verified
- current-tree README two-path onboarding: implemented
- current-tree ni init . Bubble Tea TUI: implemented
- current-tree first-user smoke after TUI: FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES
- published v0.5.0 install parity: verified mismatch in this document
- Windows real-host execution: deferred on macOS-only development host
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Audit goal

This audit checks whether the public v0.5.0 install path matches the current
README onboarding path:

```bash
ni --help
ni version
ni init .
ni status --proof --next-questions
ni end
ni run --max-chars 4000
```

The lanes are separate. Current-tree behavior is not used as proof of published
v0.5.0 behavior.

## Decision

PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED

Justification: the published v0.5.0 binary installed through the curl installer
reports `0.5.0` and passes command-name `ni --help` / `ni version`, but the
current README-required `ni init .` step fails with `unknown init option: .`.
Current-tree `ni init .` works and the Bubble Tea v2 / Lip Gloss v2 TUI was
added after the v0.5.0 tag, so the right next action is a v0.5.1 patch release
plan rather than hiding the onboarding path.

## Lane comparison

| Lane | Binary source | Version output | ni init . behavior | ni status behavior | Result | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| current-tree | Local build from current checkout at `/private/tmp/ni-task205-current/bin/ni` | `0.0.0-dev` | Passed with `ni init . --yes`; created planning artifacts in temp project. | Passed and reported honest `BLOCKED` for empty first-user intent. | Pass with expected `BLOCKED` status | Source build, not release proof. |
| published v0.5.0 | Curl installer from `https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh`, installing release asset `ni_0.5.0_darwin_arm64.tar.gz` into temp BINDIR | `0.5.0` | Failed: `unknown init option: .` | Not run because init failed. | Mismatch | Help/version parity exists; README first-project parity does not. |
| Windows real-host, deferred unless run | Not run | Not observed | Not observed | Not observed | Deferred | No Windows host, VM, CI runner, or transcript was used. |

## Published v0.5.0 transcript

| Step | Command | Expected | Observed | Pass? | Notes |
| --- | --- | --- | --- | --- | --- |
| command-name ni --help | `HOME="$tmp/home" PATH="$tmp/bin:..." sh -lc "command -v ni; ni --help"` | `ni` resolves from temp BINDIR and help exits successfully. | `command -v ni` printed `/private/tmp/ni-task205-public.yvsGkZ/bin/ni`; help exited successfully and was redirected to `/tmp/ni-task205-public-help.txt`. | Yes | Fresh shell context with temp HOME and temp BINDIR. |
| command-name ni version | `ni version` | `0.5.0` | `0.5.0` | Yes | Confirms release linker version for the installed artifact. |
| ni init . | `"$tmp/bin/ni" init .` in temp project | README-required current-directory init succeeds. | `unknown init option: .` | No | This is the parity break. |
| ni status --proof --next-questions, if init succeeds | `ni status --proof --next-questions` | Run after init succeeds. | Not run. | n/a | Init failed, so no generated planning workspace existed for status. |
| generated artifact check, if init succeeds | `find "$tmp/project" -maxdepth 3 -type f` | `.ni/contract.json`, `.ni/session.json`, and `docs/plan/**` exist. | Not run. | n/a | Init failed before artifact creation. |

The temporary installer-created binary was removed with
`sh install.sh --uninstall`; the uninstall removed the temp binary and the
ni-managed PATH block from the temporary `.zshrc`.

## README parity audit

| README claim | Current-tree evidence | Published v0.5.0 evidence | Pass? | Required action |
| --- | --- | --- | --- | --- |
| Install v0.5.0 with curl installer, then open a new shell and run `ni --help`. | Current installer script still supports command-name verification. | Published v0.5.0 install lane resolved `ni` and help exited successfully. | Yes | Keep. |
| Run `ni version` after install. | Current-tree source build reports `0.0.0-dev`; release builds should report linker version. | Published v0.5.0 reported `0.5.0`. | Yes | Keep release/source distinction visible. |
| Start first project with `ni init .`. | Current-tree lane passed with temp project. | Published v0.5.0 failed with `unknown init option: .`. | No | Prepare v0.5.1 patch release plan; README now links this parity note. |
| `ni init .` opens guided project intent wizard. | Current-tree docs and code show Bubble Tea v2 / Lip Gloss v2 TUI added after v0.5.0. | Published v0.5.0 did not reach init. | No for v0.5.0 | Include TUI/positional init changes in v0.5.1 patch inputs. |
| Windows real-host global install is verified. | README says deferred. | No Windows run in this audit. | Yes | Keep deferred wording. |
| Homebrew is Available. | README says `Homebrew: Planned / v0.5 candidate`. | No Homebrew test or publication. | Yes | Keep Planned. |
| `ni run` executes downstream work. | README says bounded prompt compilation only. | Not tested here; boundary is a claim audit item. | Yes | Keep non-execution wording. |

## Patch-release assessment

| Surface | Current-tree behavior | Published v0.5.0 behavior | Patch release needed? | Notes |
| --- | --- | --- | --- | --- |
| ni init . | Works for current-directory init and supports guided/fallback behavior. | Fails with `unknown init option: .`. | Yes | README-required command is absent from v0.5.0. |
| Bubble Tea TUI | Implemented in `internal/tui/init` after v0.5.0. | Not present in v0.5.0 release artifact. | Yes | Include TUI and non-interactive fallback tests in patch validation. |
| macOS global install path handling | Current `install.sh` supports `--update-path` and reversible PATH block. | v0.5.0 tag `install.sh` lacks `--update-path`; current public raw `main` installer can install v0.5.0 with `--update-path`. | Yes | Installer script and binary are different public surfaces. |
| install.sh --update-path | Present in current tree and public raw `main` script used by README. | Absent at v0.5.0 tag. | Yes | Patch plan should explicitly validate installer/script state against v0.5.1. |
| Windows install.ps1 | Implemented in current tree. | Added after v0.5.0; no Windows real-host execution. | Yes, for release inclusion; real-host proof still deferred | Do not claim Windows execution without transcript. |
| README two-path onboarding | Implemented in current tree. | Published v0.5.0 binary cannot complete the current first-project path. | Yes | README note added; do not remove useful onboarding. |
| docs/120 acceptance | Added after v0.5.0 and describes global install acceptance. | Not part of v0.5.0 release tag. | Yes | Include in patch release docs set. |
| docs/125 first-user smoke | Current-tree smoke passed with notes. | Explicitly says it does not prove published v0.5.0 behavior. | Yes | This audit closes that proof gap as a mismatch. |

## Patch-release inputs

If patch is needed, list:
- expected version: v0.5.1
- changes to include:
  - current-tree `ni init .` positional target support
  - Bubble Tea v2 / Lip Gloss v2 guided init TUI
  - non-interactive fallback and existing-file/lockfile safety behavior
  - current README two-path onboarding note and docs/120 through docs/126 parity evidence
  - current `install.sh --update-path` behavior and Windows `install.ps1` packaging, with Windows execution still deferred unless a transcript exists
- validation required:
  - release artifact help/version for v0.5.1
  - temporary HOME/BINDIR curl installer install for v0.5.1
  - command-name `ni --help`
  - command-name `ni version`
  - `ni init .` in a temporary project
  - `ni status --proof --next-questions` if init succeeds
  - generated artifact check if init succeeds
  - protected project-root `.ni` diff remains empty
- release action explicitly not performed in this task.

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| published v0.5.0 | May claim help/version install proof, not current `ni init .` parity. | Help/version passed; `ni init .` failed. | No for full README parity | Decision is patch needed. |
| current-tree behavior | May claim current-tree onboarding works only from current-tree evidence. | Current lane passed with `0.0.0-dev` source build. | Yes | Kept separate from v0.5.0 lane. |
| macOS install | Must use temporary HOME/BINDIR for smoke and avoid user install pollution. | Temporary install used; uninstall confirmed. | Yes | No project-root `.ni` files changed. |
| Windows install | Must remain deferred without Windows transcript. | Deferred. | Yes | Static checks may pass later, but no real-host claim here. |
| Homebrew | Homebrew: Planned / v0.5 candidate. | Planned wording preserved. | Yes | No Homebrew Available claim. |
| ni init . | Current-tree feature; release parity must be verified before public claim. | v0.5.0 mismatch observed. | No for v0.5.0 | v0.5.1 patch release plan recommended. |
| ni run | Prompt compiler only. | Not run; docs preserve non-execution boundary. | Yes | Generated prompts were not executed. |
| READY | Must come from `ni status`, not model judgment. | Current temp project status was `BLOCKED`; no release status run because init failed. | Yes | No readiness overclaim. |
| Runtime execution boundary | No task runner, execution harness, shell/Codex adapter, queue, PR/release automation, or downstream execution layer. | No runtime behavior added or executed. | Yes | Audit only. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | `M` after this task | Yes | Narrow parity note linking docs/126. |
| README.ko.md | `M` after this task | Yes | Korean companion parity note. |
| docs/120* | clean | No | Added after v0.5.0; no edit in this task. |
| docs/124* | clean | No | Used as current-tree TUI evidence. |
| docs/125* | clean | No | Used as current-tree smoke evidence. |
| docs/126* | `??` after this task | Yes | Public install parity audit and Korean companion. |
| install.sh | clean | No | Current tree has `--update-path`; no edit. |
| install.ps1 | clean | No | No Windows real-host run. |
| scripts/* | clean | No | No script updates required. |
| temporary smoke directories | outside repo | No | `/private/tmp/ni-task205-current/**` and `/private/tmp/ni-task205-public.*`. |
| .ni/contract.json | clean and empty diff | No | Protected project-root planning state. |
| .ni/session.json | clean and empty diff | No | Protected project-root session state. |
| .ni/plan.lock.json | clean and empty diff | No | Protected project-root lockfile. |
| unexpected files | none observed before edits | No | Recheck before final. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Initially clean; after edits, expected README and docs/126 changes only. |
| `git log --oneline -5` | Top commits include `93dc241 Document first-user onboarding smoke after TUI` and `4ae4a85 Redesign README and add Bubble Tea init TUI`. |
| `git tag --list v0.5.0` | `v0.5.0`. |
| `git rev-parse v0.5.0` | `b8fec7fa9615a861acf4eba59733c800c70c6cca`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Empty before edits; rechecked after validation. |
| published v0.5.0 install smoke in temporary HOME/BINDIR | Partial pass, then mismatch: installer installed `0.5.0`; `ni init .` failed. |
| command-name `ni --help` from fresh shell context for published v0.5.0 | Passed; help exited successfully from temp BINDIR. |
| command-name `ni version` from fresh shell context for published v0.5.0 | Passed; output `0.5.0`. |
| `ni init .` with published v0.5.0 in temporary project | Failed: `unknown init option: .`. |
| `ni status --proof --next-questions` with published v0.5.0 if init succeeds | Not run because init failed. |
| current-tree first-user smoke for comparison | Passed with local source build; status reported expected `BLOCKED`. |
| `gofmt -w .` | Passed; no protected planning files changed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root reported `NI Intent Readiness: READY` with blockers, deferrals, and warnings none. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help includes `ni init [.]`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed after changing the next-task prompt line to an explicit `must not include` boundary phrase. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |

## Changes made

- Added this public install parity audit.
- Added the Korean companion document.
- Added a narrow README / README.ko parity note pointing to this audit.

## What this audit proves

- Published v0.5.0 does not match the current README first-user path because
  `ni init .` fails with `unknown init option: .`.
- A v0.5.1 patch release plan is the recommended next task.
- Project-root `.ni` files were untouched.

## What this audit does not prove

- Windows real-host execution works unless a Windows transcript exists.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.
- No-terminal is deterministic.

## Recommended next task

A. prepare v0.5.1 patch release plan

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: prepare a v0.5.1 patch release plan for public install parity.

Goal:
Create a release-plan document that packages the current-tree onboarding fixes
needed because published v0.5.0 passes help/version install proof but fails the
current README-required `ni init .` step with `unknown init option: .`.

Read:
- AGENTS.md
- README.md
- README.ko.md
- docs/22_INSTALL.md
- docs/120_GLOBAL_INSTALL_ACCEPTANCE.md
- docs/121_README_TWO_PATH_INSTALL_AND_INIT_TUI.md
- docs/124_REPO_CLEANUP_README_REDESIGN_INIT_TUI.md
- docs/125_FIRST_USER_ONBOARDING_SMOKE_AFTER_TUI.md
- docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md
- install.sh
- install.ps1
- cmd/ni/
- internal/
- scripts/release-check.sh

Rules:
- Do not publish, tag, create a GitHub release, upload assets, run release
  workflows, run goreleaser publish, create or publish a Homebrew formula, run
  `ni end` on the project root, relock the project root, edit
  `.ni/plan.lock.json`, execute generated prompts, or add runtime execution
  behavior.
- Do not mark Homebrew Available.
- Do not claim Windows real-host execution without a Windows transcript.
- Keep `Skills are UX; CLI is authority.`
- Keep `ni run` as bounded prompt compilation only.

Plan content required:
- decision: V0_5_1_PATCH_RELEASE_PLAN_READY or BLOCKED
- exact release rationale tied to docs/126 evidence
- included changes: `ni init .` positional support, Bubble Tea v2 / Lip Gloss v2
  TUI, non-interactive fallback, init safety behavior, current installer PATH
  handling, README/docs parity notes
- excluded changes (patch must not include): task runner, SPEC runner, shell/Codex adapter, queue, PR
  automation, release automation, Homebrew Available claim, Windows verified
  claim, downstream execution
- validation matrix for v0.5.1 local preflight and post-publication hosted
  artifact smoke
- protected `.ni` file boundary
- rollback / docs correction path if v0.5.1 publication is not approved
- complete next executable prompt for the authorized v0.5.1 release preflight

Validation:
- git status --short
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
Report changed files, release-plan decision, validation results, protected
`.ni` diff, and confirmation that no publish/tag/release/upload/project-root
relock/generated prompt execution occurred.
```
