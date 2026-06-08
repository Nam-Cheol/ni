# Namba Intent Rename Implementation

## Current Status

State:
- v0.5.1 release: published and verified.
- rename decision: RENAME_TO_NAMBA_INTENT.
- product name: Namba Intent.
- primary command: `namba-intent`.
- repository: `Nam-Cheol/ni` retained.
- config directory: `.ni/` retained.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host verification: pending.
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

This implements the rename in the current tree without publishing v0.6.0.

## Decision Summary

| Area | Decision | Implemented? | Notes |
| --- | --- | --- | --- |
| product name | Namba Intent | Yes | README, CLI help, and current-tree docs use the new identity. |
| primary command | `namba-intent` | Yes | `cmd/namba-intent` calls the shared CLI implementation. |
| legacy `ni` shim | Keep for one transition release where safe. | Yes | `cmd/ni` delegates and warns: `ni is deprecated; use namba-intent.` |
| repository name | Keep `Nam-Cheol/ni`. | Yes | Repo rename is deferred. |
| config directory | Keep `.ni/`. | Yes | Existing contract, session, lock, and docs/plan paths remain unchanged. |
| release version | v0.6.0 future release. | Not released | No tag, publish, release workflow, or asset upload was performed. |
| Homebrew | Planned / v0.5 candidate. | Deferred | No Homebrew formula was created or marked Available. |
| Windows | Use `namba-intent.exe` as primary. | Implemented in installer | Real-host Windows verification remains pending. |

## Command Migration

| Old command | New command | Compatibility behavior | Notes |
| --- | --- | --- | --- |
| `ni init .` | `namba-intent init .` | `ni` shim delegates with warning. | `.ni/` remains the workspace directory. |
| `ni status --proof --next-questions` | `namba-intent status --proof --next-questions` | `ni` shim delegates with warning. | CLI remains readiness authority. |
| `ni end` | `namba-intent end` | `ni` shim delegates with warning. | Do not run on project root unless explicitly authorized. |
| `ni run --max-chars 4000` | `namba-intent run --max-chars 4000` | `ni` shim delegates with warning. | Prompt compilation only; no downstream execution. |
| `ni version` | `namba-intent version` | `ni` shim delegates with warning. | Source builds can still report `0.0.0-dev`. |

## Difference From namba-ai

| Surface | NambaAI | Namba Intent | Notes |
| --- | --- | --- | --- |
| command | `namba` | `namba-intent` | This project must not claim the `namba` command. |
| purpose | Codex workflow, SPEC execution, queue, sync, PR, land. | Pre-execution intent compile, readiness, lock, handoff prompt. | Names may share family branding; behavior does not. |
| execution behavior | May run implementation workflows. | Does not execute downstream work. | Namba Intent stops before runtime. |
| planning behavior | SPEC/workflow execution planning. | Intent Lock Protocol planning contract. | Kernel is authoritative. |
| queue / PR / release automation | In scope for namba-ai workflows. | Out of scope. | Not added here. |
| config directory | `.namba/` | `.ni/` | `.ni/` remains compatible. |
| target user | Workflow operators. | Users and agents compiling intent before handoff. | Boundary is before execution. |

## Implementation Surfaces

| Surface | Change made | Notes |
| --- | --- | --- |
| `cmd/namba-intent` | Added primary CLI entrypoint. | Calls `internal/cli.Run`. |
| `cmd/ni` | Converted to deprecated shim. | Prints warning to stderr, then delegates. |
| `internal/cli` | Moved shared CLI implementation out of `cmd/ni`. | Help and init copy now use Namba Intent / `namba-intent`. |
| `install.sh` | Installs `namba-intent` and verifies `namba-intent --help` / `version`. | Archive prefix changed to `namba-intent_`. |
| `install.ps1` | Installs `namba-intent.exe` under `%LOCALAPPDATA%\namba-intent\bin`. | No `ni` alias profile block required for primary path. |
| `README` | Current-tree product copy uses Namba Intent. | Warns latest published v0.5.1 may still use `ni`. |
| `README.ko` | Korean companion aligned with README. | Does not promise more than English. |
| install docs | Checker updated for Namba Intent markers. | Historical v0.5.x docs may retain `ni`. |
| `scripts/check-install-docs.py` | Current README expectations updated. | Keeps Homebrew and Windows proof boundaries. |
| `scripts/check-install-ps1.py` | Verifies `namba-intent.exe` primary install. | Forbids primary-path alias cleanup reliance. |
| release-check | Future GoReleaser markers use `namba-intent`. | Does not require hosted v0.6.0 assets. |
| skills | Rename pass still preserves `ni-*` skill IDs for transition. | Skills are UX; CLI is authority. |
| `.agents` | Existing local skill IDs may remain `ni-*`. | Renaming skill IDs is a separate compatibility task. |

## Compatibility Policy

`namba-intent` is primary. The `ni` command remains as a deprecated source-tree
shim for one transition release where safe and maintainable. The shim warning
is:

```text
ni is deprecated; use namba-intent.
```

Windows primary install does not rely on `ni`, because PowerShell defines
`ni -> New-Item`. Legacy v0.5.x alias cleanup remains historical guidance and
is not required for `namba-intent.exe`.

## Historical Docs Policy

Historical v0.5.x verification docs may keep `ni` when they describe real past
commands and public-release evidence. Current user-facing instructions should
use Namba Intent and `namba-intent`, with explicit notes when latest published
v0.5.1 still differs from current main.

## Claim-Boundary Audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| NambaAI distinction | Do not use `namba` for this project. | Primary command is `namba-intent`. | Yes | namba-ai keeps `namba`. |
| Namba Intent identity | Current tree may implement rename. | Help and README identify Namba Intent. | Yes | v0.6.0 is not published. |
| `namba-intent` command | Primary current-tree command works if validated. | Entrypoint exists and tests pass. | Yes | Hosted install still release-gated. |
| legacy `ni` shim | Warn and delegate only. | Shim warns on stderr and delegates. | Yes | Not primary Windows path. |
| `.ni/` compatibility | Do not rename state directory. | Protected `.ni` paths unchanged. | Yes | Root protected files were not edited. |
| Homebrew | Planned / v0.5 candidate only. | No Available claim added. | Yes | No formula publication. |
| Windows verification | Pending without transcript. | Static installer checks only. | Yes | No real-host claim. |
| run behavior | Prompt compilation only. | `run` remains bounded prompt compilation. | Yes | No agent/shell execution added. |
| runtime execution boundary | No task runner, SPEC runner, shell/Codex adapter, queue, PR/release automation. | No such behavior added. | Yes | Kernel boundary preserved. |

## Git Status / Inclusion Check

| Path or group | git status --short | Expected in v0.6.0? | Notes |
| --- | --- | --- | --- |
| `README.md` | modified | Yes | Current-tree Namba Intent surface. |
| `README.ko.md` | modified | Yes | Korean companion. |
| `cmd/ni/*` | modified/deleted old files | Yes | Shim only. |
| `cmd/namba-intent/*` | untracked new | Yes | Primary command. |
| `internal/*` | modified/new | Yes | Shared CLI package and prompt text. |
| `go.mod` | unchanged | Yes | No dependency change. |
| `go.sum` | unchanged | Yes | No dependency change. |
| `install.sh` | modified | Yes | Primary Unix installer. |
| `install.ps1` | modified | Yes | Primary Windows installer. |
| `scripts/*` | modified | Yes | Checkers and command-name smoke. |
| `docs/135*` | untracked | Yes | Decision source supplied by prior task. |
| `docs/136*` | untracked new | Yes | Implementation record. |
| `packages/*` | modified | Yes | Skill IDs remain `ni-*`; command examples use `namba-intent`. |
| `.agents/*` | modified | Yes | Local skill IDs remain `ni-*`; command examples use `namba-intent`. |
| `.ni/contract.json` | no diff expected | Yes | Protected. |
| `.ni/session.json` | no diff expected | Yes | Protected. |
| `.ni/plan.lock.json` | no diff expected | Yes | Protected. |
| unexpected files | existing `dist/*` / `.ni/generated/*` | No new action | Pre-existing generated artifacts were not executed. |

## Validation Results

Commands run:

| Command | Result | Notes |
| --- | --- | --- |
| `git status --short` | Passed | Dirty tree present before implementation. |
| `git log --oneline --decorate -20` | Passed | HEAD on `main`; v0.5.1 tag present behind HEAD. |
| `git tag --list v0.5.1` | Passed | `v0.5.1` exists. |
| `git tag --list v0.6.0` | Passed | No `v0.6.0` tag found. |
| `git rev-parse v0.5.1` | Passed | Resolved `b588f6b2e13111841081d186bd0e70d3c0bfbd6c`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed | No protected root `.ni` diff. |
| `gofmt -w cmd internal` | Passed | Go files formatted. |
| `go test ./...` | Environment failed | Default Go cache outside sandbox was not writable. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed | All Go packages passed after rename. |
| `python3 scripts/check-install-docs.py` | Passed | Current install doc markers aligned. |
| `python3 scripts/check-install-ps1.py` | Passed | Windows installer static safety check passed. |
| `bash scripts/check-skill-packs.sh` | Passed | Package and repo-local skill packs passed. |
| `bash scripts/test-install-sh.sh` | Passed | Unix installer dry-run, checksum, install, PATH, uninstall passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Passed | Source, build, temp install, command-name check passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Passed | Current-tree `namba-intent` CLI smoke passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/demo-check.sh` | Passed | Demo prompts/exports remained seed-only. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed | Broad quality wrapper passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed | Check-only release readiness gate passed; no release action. |
| Fresh-shell temp binary smoke | Passed | `namba-intent --help`, `version`, `init . --yes`, `status --proof --next-questions`, repeated init, lockfile safety, and `ni` shim delegation passed in `/private/tmp`. |

## Changes Made

- Added primary `namba-intent` command.
- Kept deprecated `ni` shim with stderr warning.
- Kept `.ni/` state paths.
- Updated primary install scripts for `namba-intent` / `namba-intent.exe`.
- Updated current README surfaces to Namba Intent while warning that public
  v0.5.1 may still use `ni`.
- Updated prompt and generated-doc wording for current-tree command examples.
- Updated initial install and release checkers for the new primary command.

## What This Implementation Proves

State only:
- current tree implements Namba Intent rename surfaces;
- `namba-intent` command works in current tree if validated;
- `.ni/` compatibility remains intact;
- no release action was performed.

## What This Implementation Does Not Prove

State:
- v0.6.0 has been published;
- public install retrieves `namba-intent`;
- Windows real-host execution works;
- Homebrew is Available;
- external users succeed;
- downstream execution succeeds.

## Recommended Next Task

A. Namba Intent first-user smoke.

## Next Task Prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: Namba Intent first-user smoke for current tree.

Use docs/136_NAMBA_INTENT_RENAME_IMPLEMENTATION.md as the implementation
record. Verify current-tree behavior only. Do not publish, tag, create a
GitHub release, upload assets, run release workflows, run GoReleaser publish,
create or publish Homebrew formula, mark Homebrew Available, run ni end on the
project root, relock the project root, execute generated prompts, or claim
Windows real-host verification without a transcript.

Required checks:
- build a temporary namba-intent binary outside the repo;
- verify command-name namba-intent --help from a fresh shell context;
- verify command-name namba-intent version from a fresh shell context;
- run namba-intent init . --yes in a temporary project;
- run namba-intent status --proof --next-questions in that temporary project;
- verify repeated namba-intent init . safety;
- verify lockfile safety in a temporary project;
- verify deprecated ni shim warning and delegation if built;
- verify git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json is empty.

Final response must separate current-tree proof from public-release proof.
```
