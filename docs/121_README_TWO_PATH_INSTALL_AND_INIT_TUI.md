# README Two-Path Install and Init TUI

## Current status

- v0.5.0 publication: verified
- macOS/Linux global install path handling: implemented
- Windows User PATH installer: implemented
- Windows real-host execution: deferred on macOS-only development host
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- ni is a pre-runtime Project Intent Compiler for AI Agents.

## Pass goal

This pass simplifies README onboarding and improves `ni init .`. A new user
should see two primary install paths, verify `ni` by global command name, then
start guided project intent setup in the current directory without confusing
that flow with downstream execution.

## README install simplification

| Area | Before | After | Notes |
| --- | --- | --- | --- |
| macOS path | README mixed source, local binary, release binary, curl installer, model workspace, no-terminal, and Homebrew rows before the first success path. | README shows one macOS primary path: inspect `install.sh`, dry-run, install, open a new shell, run `ni --help`, `ni version`, and `ni init .`. | Detailed source/release/local build alternatives moved to `docs/22_INSTALL.md`. |
| Windows path | README contained a longer installer section after a broad path matrix. | README shows one Windows primary path: inspect `install.ps1`, dry-run, install, open new PowerShell, run `ni --help`, `ni version`, and `ni init .`. | Windows real-host execution remains deferred until a Windows transcript exists. |
| Homebrew | Homebrew appeared in the README path matrix as Planned. | README mentions only `Homebrew: Planned / v0.5 candidate` and sends package-manager details to docs. | No Homebrew Available claim. |
| advanced install docs | README carried much of the install manual. | README links to `docs/22_INSTALL.md` for source, local build, Linux, release archive, and advanced uninstall details. | README remains a product pamphlet. |
| uninstall | macOS and Windows uninstall steps were present but embedded in long install sections. | Each primary path includes its uninstall command. | Uninstall remains bounded to installer-managed binary/PATH entries. |
| verification | Verification was described across several sections. | Both primary paths require `ni --help`, `ni version`, and `ni init .` after opening a new shell/session. | Global command verification remains the install success criterion. |

## First-use tutorial

README now centers the first project around:

```bash
mkdir my-project
cd my-project
ni init .
ni status --proof --next-questions
ni end
ni run --max-chars 4000
```

This preserves non-execution boundaries: `ni init .` creates initial intent
artifacts, `ni status --proof --next-questions` checks readiness, `ni end`
locks the accepted plan, and `ni run --max-chars 4000` compiles a bounded
downstream handoff prompt. `ni run` does not execute downstream work.

## ni init . TUI behavior

| Behavior | Expected | Implemented | Notes |
| --- | --- | --- | --- |
| current-directory target | `ni init .` works from an existing project directory. | Implemented. | Positional target support was added while preserving `--dir`. |
| interactive prompts | Interactive stdin can guide a minimal project-intent setup. | Implemented. | The flow asks for project name, goal, audience, eventual downstream task, constraints/non-goals, success criteria, blockers, and deferrals. |
| non-interactive behavior | Existing scripts and non-interactive runs stay stable. | Implemented. | `--dir` remains non-interactive by default; `--interactive` opts in during tests or piped stdin. |
| generated files | Init writes `.ni/contract.json`, `.ni/session.json`, and `docs/plan/**`. | Implemented through existing `docstore.InitWithOptions`. | Guided fields are passed through the existing writer, not a parallel contract writer. |
| existing file handling | Existing planning files are not silently overwritten. | Implemented. | Existing paths are reported; non-interactive init adds missing files only. Interactive mode can keep existing files, add missing files only, or abort. |
| lockfile safety | Existing `.ni/plan.lock.json` is not modified. | Implemented. | If a lockfile exists, init warns, points to status/amend/relock flow, and exits without writing files. |
| abort behavior | User can abort before writes. | Implemented. | Interactive confirmation happens before writing guided artifacts. |

## Tests added

- `TestInitCurrentDirectoryTarget`
- `TestInitNonInteractiveBehaviorIsStable`
- `TestInitInteractiveFlowFromStdin`
- `TestInitDoesNotSilentlyOverwriteExistingFiles`
- `TestInitDoesNotModifyExistingLockfile`
- `TestInitDoesNotCreateDownstreamExecutionArtifacts`

## What this pass proves

- README now presents two primary install paths
- `ni init .` onboarding was implemented or improved for tested behavior
- global command verification remains the install success criterion
- `ni run` remains bounded prompt compilation only

## What this pass does not prove

- Windows real-host execution works unless a Windows transcript exists
- Homebrew is Available
- downstream execution succeeds
- no-terminal is deterministic
- external users succeed

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| macOS install | Available path must verify global command name. | README and install docs require new shell plus `ni --help` and `ni version`. | Yes | Local validation uses macOS host and temporary paths. |
| Windows install | Installer code may be documented; real-host execution must not be claimed. | README says real-host execution is deferred on this macOS-only host. | Yes | `scripts/check-install-ps1.py` remains static safety proof only. |
| Homebrew | Planned / v0.5 candidate only. | README and docs keep Homebrew Planned. | Yes | No `brew install` path is presented. |
| ni init . | Guided intent setup only. | CLI writes planning artifacts and does not run agents. | Yes | Existing writer is reused. |
| ni run | Bounded prompt compilation only. | README says it does not execute downstream work. | Yes | No runtime behavior was added. |
| READY | CLI readiness only, not product readiness. | README keeps readiness tied to `ni status`. | Yes | Model judgment is not readiness. |
| Model workspace packs | Experimental. | README links to model workspace docs and does not mark packs Available. | Yes | Skills are UX; CLI is authority. |
| No-terminal | Experimental / assisted. | README links to no-terminal docs without deterministic validation claims. | Yes | Trusted CLI proof remains required. |
| Runtime execution boundary | No task runner, shell adapter, Codex exec adapter, queue, or downstream execution. | README and code preserve non-execution boundary. | Yes | No execution adapter was added. |

## Git status / inclusion check

| Path or group | git status --short | Expected in next commit? | Notes |
| --- | --- | --- | --- |
| README.md | M | Yes | Two-path install and first project tutorial. |
| README.ko.md | M | Yes | Korean companion mirrors README.md. |
| cmd/ni/* | M | Yes | `ni init .` target, guided prompts, tests. |
| internal/* | M | Yes | Guided init fields in existing docstore writer. |
| docs/22_INSTALL.md | M | Yes | Detailed install matrix and primary path handoff. |
| docs/install-curl* | M | Yes | Post-install `ni init .` step. |
| docs/120* | unchanged | No | Existing global install acceptance remains current. |
| docs/121* | A | Yes | This audit and Korean companion. |
| install.sh | unchanged | No | No installer shell behavior changed. |
| install.ps1 | unchanged | No | No PowerShell installer behavior changed. |
| scripts/* | M | Yes | README/install/release checkers updated for two-path README; `demo-check.sh` gets a writable default Go cache. |
| .ni/contract.json | unchanged | No | Protected project-root planning state. |
| .ni/session.json | unchanged | No | Protected project-root planning state. |
| .ni/plan.lock.json | unchanged | No | Protected project-root lockfile. |
| unexpected files | none observed | No | Recheck before commit. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; expected README, docs, Go, and scripts changes only. |
| `gofmt -w .` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; `NI Intent Readiness: READY`, blockers/deferrals/warnings none. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build printed `0.0.0-dev`. |
| Temp cwd `ni init . --yes` via a built temporary binary | Passed; `.ni/contract.json`, `.ni/session.json`, and `docs/plan/00_project_brief.md` created. |
| `python3 scripts/check-install-docs.py` | Passed |
| `python3 scripts/check-install-ps1.py` | Passed |
| `bash scripts/check-skill-packs.sh` | Passed |
| `bash scripts/demo-check.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; no output. |

## Changes made

- Simplified README and README.ko.md around macOS and Windows primary install paths.
- Added `ni init .` positional target handling and a small guided interactive init flow.
- Kept non-interactive `--dir` initialization stable for scripts.
- Added existing-file and lockfile safety handling for init.
- Added install docs handoff for first-use `ni init .`.
- Updated install docs checker so README no longer needs the detailed path matrix.

## Recommended next task

Selected: A. generate README visual assets from docs/116 prompts

This is selected because README/TUI onboarding is stable enough for the next
product-facing pass, while Windows real-host verification remains unavailable
on this macOS-only development host.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: Generate and verify README visual assets from docs/116 prompts.

Read AGENTS.md, README.md, README.ko.md, docs/116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.md, docs/121_README_TWO_PATH_INSTALL_AND_INIT_TUI.md, and the current assets/ files.

Goal:
Improve README visuals without changing install availability claims or execution boundaries.

Rules:
- Do not publish, tag, release, upload, or push.
- Do not edit .ni/contract.json, .ni/session.json, or .ni/plan.lock.json.
- Do not run ni end or relock the project root.
- Do not imply Homebrew is Available.
- Do not imply Windows real-host execution is verified.
- Do not imply ni run executes downstream work.
- Keep exact product/status text in Markdown or deterministic SVG, not AI-generated raster text.

Work:
1. Audit existing README image references and confirm which files exist.
2. Generate or update only visual assets that can be verified locally.
3. Keep README.md and README.ko.md alt text accurate and claim-safe.
4. Add a short docs note with assets changed, verification, and claim-boundary audit.
5. Run README/rendering and quality checks available in the repo.

Validation:
- git status --short
- python3 scripts/check-install-docs.py
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
Report changed files, visual assets changed, README/README.ko status, claim-boundary audit, validation results, protected .ni diff, and confirmation that no publish/tag/release/upload/project-root relock/generated prompt execution occurred.
```
