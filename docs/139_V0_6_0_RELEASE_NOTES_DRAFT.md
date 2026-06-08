# v0.6.0 Release Notes Draft

## Current status

State after publication:
- v0.5.1 release: published and verified
- v0.6.0 release: published and verified on macOS darwin/arm64
- Namba Intent rename: released
- primary command: namba-intent
- deprecated ni shim: transition-only
- .ni/ compatibility: preserved
- public install retrieval of namba-intent: verified for `install.sh --version 0.6.0` on darwin/arm64
- Homebrew: Planned / v0.5 candidate
- Windows real-host verification: pending
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

Post-release evidence is recorded in
[docs/140 post-release verification](140_V0_6_0_POST_RELEASE_VERIFICATION.md).

## Historical draft goal

This document originally drafted conservative v0.6.0 release notes for the
Namba Intent rename before publication. It did not itself create a tag, upload
assets, create Homebrew material, run `namba-intent end` on the project root,
relock the project root, execute generated prompts, or add downstream
execution behavior.

## Decision

V0_6_0_RELEASE_NOTES_SUPERSEDED_BY_POST_RELEASE_VERIFICATION

Justification: the draft accurately described the rename before publication.
It is now superseded by docs/140 for release metadata, hosted artifacts,
checksums, public install retrieval, and current-platform artifact proof.
Windows real-host verification is still pending, Homebrew remains Planned /
v0.5 candidate, repository rename is deferred, model workspace packs remain
Experimental, the no-terminal method remains Experimental / assisted, and
external user validation has not been performed.

## Release notes draft

### Summary

v0.6.0 introduces the Namba Intent product name and makes `namba-intent` the
primary CLI command.

Namba Intent is a Project Intent Compiler for AI Agents:

```text
Don't run the agent yet.
Compile the intent first.
```

It turns planning conversations into docs contracts, checks readiness, locks
accepted plans, and compiles bounded downstream handoff prompts before any
execution harness runs.

### Why this release exists

The short `ni` command is being retired as the primary public command for two
reasons:

- PowerShell defines `ni -> New-Item`, so `ni` is unsafe and confusing as a
  Windows-first command.
- `namba` is already reserved for the existing namba-ai CLI, which is a
  Codex/SPEC workflow tool.

The new name keeps Namba Intent distinct from NambaAI: Namba Intent is
pre-execution intent compile / readiness / lock / bounded handoff prompt.
NambaAI is Codex workflow / SPEC execution / queue / sync / PR / land oriented.

### Highlights

- Product name becomes Namba Intent.
- Primary command becomes `namba-intent`.
- `ni` remains a deprecated transition shim when present.
- `.ni/` remains the state and config directory for compatibility.
- Repository remains `Nam-Cheol/ni` for now.
- Release archive names move toward `namba-intent_<version>_<os>_<arch>`.
- `namba-intent run --max-chars 4000` compiles a bounded handoff prompt only.
- Namba Intent remains distinct from NambaAI and does not use the `namba`
  command.

### Breaking / migration notes

New scripts, docs, and user workflows should call `namba-intent`.

The transition shim may keep `ni` usable for one release where safe, but it is
not the primary command. When invoked from the current tree, it warns:

```text
ni is deprecated; use namba-intent.
```

Existing project state remains compatible because `.ni/contract.json`,
`.ni/session.json`, `.ni/plan.lock.json`, and `docs/plan/**` paths do not
change.

### Added

- Primary `namba-intent` CLI entrypoint.
- Release configuration for `namba-intent` archives and checksums.
- Current-tree installer paths for `namba-intent` and `namba-intent.exe`.
- Conservative rename, first-user smoke, readiness, and release-note evidence
  docs.

### Changed

- README and current install docs use Namba Intent and `namba-intent` for
  current/future surfaces.
- Historical v0.5.1 public release evidence stays tied to the older `ni`
  command.
- Windows install guidance moves away from PowerShell `ni` alias cleanup for
  the new primary command.
- Skill-pack examples use `namba-intent` as the CLI authority while preserving
  transition skill IDs.

### Deprecated

- `ni` as the primary command.

The `ni` shim is transition-only. It must not become the documented primary
path and must not be used to claim Windows compatibility without real Windows
PowerShell evidence.

### Compatibility

- `.ni/` remains unchanged.
- Existing planning docs and lock paths remain unchanged.
- `namba-intent status`, `namba-intent end`, and `namba-intent run` remain the
  readiness, lock, and prompt compiler authorities.
- Skills are UX; CLI is authority.

### Install notes

v0.6.0 assets are published and public install retrieval for
`namba-intent` is verified on darwin/arm64 in docs/140. The verified commands
were:

```bash
namba-intent --help
namba-intent version
namba-intent init .
```

Historical v0.5.1 evidence remains tied to the older `ni` command. Homebrew
remains Planned / v0.5 candidate.

### Known deferrals

- Windows real-host execution transcript.
- Homebrew Available proof.
- External user validation.
- Repository rename.
- Model workspace host behavior.
- No-terminal deterministic validation.

### What this release does not do

- It does not make Namba Intent a task runner, SPEC runner, execution harness,
  shell adapter, Codex exec adapter, queue, PR automation system, release
  automation system, or downstream execution layer.
- It does not make `run` execute downstream work.
- It does not rename `.ni/`.
- It does not rename the repository.
- It does not make Homebrew Available.
- It does not verify Windows real-host execution without a transcript.
- It does not claim Windows real-host execution without a transcript.

### Validation evidence

Evidence is recorded in:

- [docs/135 rename plan](135_NAMBA_INTENT_RENAME_PLAN.md)
- [docs/136 rename implementation](136_NAMBA_INTENT_RENAME_IMPLEMENTATION.md)
- [docs/137 first-user smoke](137_NAMBA_INTENT_FIRST_USER_SMOKE.md)
- [docs/138 readiness audit](138_V0_6_0_PUBLIC_INSTALL_PARITY_AND_RELEASE_READINESS.md)

This draft also records current validation commands below.

### Maintainer checklist before publication

- Confirm intended release commit and clean working tree.
- Confirm `v0.6.0` tag absence before authorized tag creation.
- Run full Go, quality, install, smoke, and release-check gates.
- Run a release-like artifact dry-run and verify version injection.
- Verify archive names and checksum names use `namba-intent`.
- Verify the deprecated `ni` shim warning and delegation.
- Verify hosted release metadata, assets, and checksums after publication.
- Verify public `install.sh` retrieval of `namba-intent` after publication.
- Keep Windows real-host, Homebrew, external-user, model-workspace, and
  no-terminal deferrals explicit unless separate evidence exists.

## Rename rationale

| Issue | Evidence | v0.6.0 response | Notes |
| --- | --- | --- | --- |
| PowerShell ni alias conflict | `docs/134_WINDOWS_POWERSHELL_ALIAS_FIX.md` records `ni -> New-Item`, and docs/135 chooses a safer command. | Make `namba-intent` primary. | PowerShell alias cleanup remains legacy v0.5.x guidance for `ni`. |
| namba-ai command namespace | docs/135 records that `namba` is already reserved for the existing namba-ai CLI. | Do not use `namba`; use `namba-intent`. | Keeps Namba Intent distinct from NambaAI. |
| Namba Intent identity | README, docs/135, and docs/136 use Project Intent Compiler language. | Product name becomes Namba Intent. | Pre-runtime intent compile boundary stays intact. |
| .ni compatibility | docs/135 and docs/136 preserve `.ni/`. | Keep `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`, and `docs/plan/**`. | No state directory rename in v0.6.0. |
| public install parity | docs/138 says public install retrieval of `namba-intent` is release-gated. | Draft notes require post-publication verification. | Current tree is not public release proof. |

## Migration notes

| Old | New | Compatibility | Notes |
| --- | --- | --- | --- |
| `ni init .` | `namba-intent init .` | `ni` shim may delegate with warning where present. | Creates planning workspace only. |
| `ni status --proof --next-questions` | `namba-intent status --proof --next-questions` | `ni` shim may delegate with warning where present. | CLI remains readiness authority. |
| `ni end` | `namba-intent end` | `ni` shim may delegate with warning where present. | Do not run on project root in this task. |
| `ni run --max-chars 4000` | `namba-intent run --max-chars 4000` | `ni` shim may delegate with warning where present. | Bounded prompt compilation only. |
| `ni version` | `namba-intent version` | `ni` shim may delegate with warning where present. | Source builds can report `0.0.0-dev`. |
| `.ni/` | `.ni/` | Preserved. | Existing planning state remains compatible. |

## NambaAI distinction

| Surface | NambaAI | Namba Intent | Notes |
| --- | --- | --- | --- |
| command | `namba` | `namba-intent` | This project must not claim the `namba` command. |
| purpose | Codex workflow, SPEC execution, queue, sync, PR, land. | Pre-execution intent compile, readiness, lock, bounded handoff prompt. | Same product family, different job. |
| execution behavior | May run implementation workflows. | Does not execute downstream work. | Namba Intent stops before runtime. |
| planning behavior | SPEC/workflow execution planning. | Intent Lock Protocol planning contract. | Kernel remains authoritative. |
| queue / PR / release automation | In scope for namba-ai workflows. | Out of scope for Namba Intent. | Namba Intent must not include queue, PR automation, or release automation. |
| config directory | `.namba/` in namba-ai. | `.ni/` remains compatible. | Do not rename `.ni/` in v0.6.0. |

## Release notes claim audit

| Claim area | Expected boundary | Observed wording | Pass? | Notes |
| --- | --- | --- | --- | --- |
| v0.6.0 publication status | Must have post-release proof after publication. | Superseded by docs/140. | Yes | Tag, workflow, release, assets, and checksums verified. |
| public install retrieval | Must be verified before claiming public `namba-intent` retrieval. | Superseded by docs/140. | Yes | `install.sh --version 0.6.0` verified on darwin/arm64. |
| namba-intent command | May describe released primary command. | Draft says primary command becomes `namba-intent`. | Yes | Public install proof is in docs/140. |
| deprecated ni shim | Transition-only. | Draft says `ni` is deprecated and not primary. | Yes | Warning text preserved. |
| .ni compatibility | Must preserve `.ni/`. | Draft says `.ni/` remains state/config directory. | Yes | No `.ni` rename. |
| NambaAI distinction | Must not confuse Namba Intent with NambaAI. | Draft includes distinction and reserves `namba` for namba-ai. | Yes | Command namespace clear. |
| Homebrew | Must remain Planned / v0.5 candidate. | Draft preserves this exact status. | Yes | No Homebrew Available claim. |
| Windows real-host verification | Must remain pending without transcript. | Draft says pending. | Yes | Static docs are not real-host proof. |
| run behavior | Prompt compilation only. | Draft says `run` compiles a bounded handoff prompt only. | Yes | No execution claim. |
| runtime execution boundary | Must not include task runner, SPEC runner, shell/Codex adapter, queue, PR/release automation, or downstream execution layer. | Draft lists these as excluded. | Yes | Kernel boundary preserved. |

## Validation evidence

| Evidence | Result | Notes |
| --- | --- | --- |
| docs/135 rename plan | RENAME_TO_NAMBA_INTENT. | Establishes name, command, `.ni/`, repo, and NambaAI boundary. |
| docs/136 rename implementation | Current-tree rename implemented. | Records primary command, transition shim, installers, and checks. |
| docs/137 first-user smoke | NAMBA_INTENT_FIRST_USER_SMOKE_PASS_WITH_NOTES. | Pre-release command-name, init/status, repeated init, lockfile safety, and shim checks passed. |
| docs/138 readiness audit | V0_6_0_RELEASE_READINESS_READY_WITH_NOTES. | Pre-release rename/readiness ready with explicit release-time deferrals. |
| docs/140 post-release verification | V0_6_0_RELEASE_EXECUTED_AND_VERIFIED. | Release metadata, assets, checksums, darwin/arm64 artifact, and installer path verified. |
| current validation commands | See Validation results. | Run in this task after docs/139 creation. |
| protected .ni diff | Empty after validation. | `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` unchanged. |

## Known deferrals

| Deferral | Reason | Required future evidence | Blocks release notes? |
| --- | --- | --- | --- |
| v0.6.0 publication | Superseded by docs/140. | Human-approved tag, release workflow, release metadata. | Closed |
| public install retrieval of namba-intent | Superseded by docs/140. | Isolated install from published v0.6.0 assets plus help/version proof. | Closed |
| hosted v0.6.0 artifacts | Superseded by docs/140. | Asset inventory and checksum verification. | Closed |
| Windows real-host verification | macOS/current environment cannot prove it. | Windows PowerShell install, new-session help/version, init, uninstall transcript. | No |
| Homebrew Available | No tap/formula/install proof. | Tap/formula, checksums, audit, install, `namba-intent --help`, `namba-intent version`. | No |
| external user validation | No external tester transcript. | Tester transcript and comprehension review. | No |
| repository rename | Deferred to avoid URL/install churn. | Separate repo-rename plan and verified redirects/install docs. | No |
| model workspace host behavior | Host/provider behavior varies and remains Experimental. | Host-specific install/discovery/provider proof. | No |
| no-terminal deterministic validation not claimed | No-terminal method remains Experimental / assisted. | Trusted CLI proof for target workspaces. | No |

## Git status / inclusion check

| Path or group | git status --short | Expected in v0.6.0? | Notes |
| --- | --- | --- | --- |
| README.md | clean at task start | Yes | Current-tree Namba Intent surface already tracked. |
| README.ko.md | clean at task start | Yes | Korean companion already tracked. |
| docs/135* | clean at task start | Yes | Rename plan pair tracked. |
| docs/136* | clean at task start | Yes | Rename implementation pair tracked. |
| docs/137* | clean at task start | Yes | First-user smoke pair tracked. |
| docs/138* | clean at task start | Yes | v0.6.0 readiness audit pair tracked. |
| docs/139* | added in this task | Yes | Release notes draft pair. |
| CHANGELOG.md | absent | No | Not added to avoid release-history confusion. |
| RELEASE.md | absent | No | Not added to avoid implying publication. |
| .ni/contract.json | no diff | No direct edit | Protected. |
| .ni/session.json | no diff | No direct edit | Protected. |
| .ni/plan.lock.json | no diff | No direct edit | Protected. |
| unexpected files | none at task start | No | No generated prompt execution. |

## Validation results

| Command | Result |
| --- | --- |
| `git status --short` | Passed; shows only `docs/51_POST_RELEASE_ROADMAP*` modifications and new `docs/139*` files. |
| `git log --oneline --decorate -20` | Checked before editing; HEAD was `dfcbf7a Clarify Namba Intent install and release boundaries`. |
| `git tag --list v0.5.1` | Passed; `v0.5.1` exists. |
| `git tag --list v0.6.0` | Passed; empty. |
| `git rev-parse v0.5.1` | Passed; `b588f6b2e13111841081d186bd0e70d3c0bfbd6c`. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty before and after edits. |
| Required ripgrep scans | Reviewed rename, release, install, Homebrew, Windows, `.ni`, and runtime boundary surfaces. |
| `gofmt -w .` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed. |
| `git diff --check` | Passed. |

## Changes made

| File | Why |
| --- | --- |
| docs/139_V0_6_0_RELEASE_NOTES_DRAFT.md | Added conservative v0.6.0 release notes draft, migration notes, claim audit, deferrals, and next task prompt. |
| docs/139_V0_6_0_RELEASE_NOTES_DRAFT.ko.md | Added Korean companion without widening English claims. |
| docs/51_POST_RELEASE_ROADMAP.md | Added a narrow pointer to the v0.6.0 release notes draft. |
| docs/51_POST_RELEASE_ROADMAP.ko.md | Added the matching Korean roadmap pointer. |

## What this historical draft proved

State only:
- v0.6.0 release notes are ready under audited boundaries with known notes.
- Release notes described the Namba Intent rename before publication.
- Known deferrals remain explicit.
- Release action evidence is now recorded in docs/140.

## What docs/140 now proves

State:
- v0.6.0 has been published.
- Hosted v0.6.0 artifacts exist.
- Public install retrieves namba-intent.
- Windows real-host execution works.
- Homebrew is Available.
- External users succeed.
- Downstream execution succeeds.

## Recommended next task

A. v0.6.0 artifact dry-run

Selection rationale: release notes are ready with notes, and docs/138 already
records release-candidate readiness. The next useful proof is local
release-like artifact generation and installer validation without publication.

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: v0.6.0 artifact dry-run for Namba Intent without publishing.

Use docs/135_NAMBA_INTENT_RENAME_PLAN.md,
docs/136_NAMBA_INTENT_RENAME_IMPLEMENTATION.md,
docs/137_NAMBA_INTENT_FIRST_USER_SMOKE.md,
docs/138_V0_6_0_PUBLIC_INSTALL_PARITY_AND_RELEASE_READINESS.md, and
docs/139_V0_6_0_RELEASE_NOTES_DRAFT.md as source evidence.

Goal:
Verify local release-like v0.6.0 artifacts for the Namba Intent rename without
publishing, tagging, uploading assets, creating a GitHub Release, running a
release workflow, or running GoReleaser publish.

Required boundaries:
- Do not publish.
- Do not tag.
- Do not create a GitHub release.
- Do not upload assets.
- Do not run release workflows.
- Do not run GoReleaser publish.
- Do not create or publish a Homebrew formula.
- Do not mark Homebrew Available.
- Do not claim public install retrieves namba-intent before v0.6.0 release.
- Do not claim hosted v0.6.0 artifacts exist.
- Do not claim Windows real-host verification without transcript.
- Do not run ni end or namba-intent end on the project root.
- Do not relock the project root.
- Do not edit .ni/plan.lock.json.
- Do not execute generated prompts.
- Do not add task runner, SPEC runner, shell adapter, Codex exec adapter,
  queue, PR automation, release automation, or downstream execution behavior.

Checks:
- Record git status, v0.5.1 tag, v0.6.0 tag absence, v0.5.1 rev, and protected
  .ni diff.
- Build local temporary artifacts with version injection for v0.6.0.
- Verify archive names use namba-intent_<version>_<os>_<arch>.
- Verify checksum file uses namba-intent_<version>_checksums.txt.
- Verify extracted namba-intent --help and namba-intent version.
- Verify deprecated ni shim warning and version if included in the archive.
- Run install.sh --dry-run --version 0.6.0 against local/fake release assets if
  supported by the existing test harness.
- Keep current-tree proof separate from public-release proof.

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
Report changed files, artifact dry-run decision, archive/checksum evidence,
installer dry-run result, validation results, protected .ni diff result,
confirmation that no publication/tag/release/upload/root relock/generated
prompt execution/Homebrew Available/downstream execution occurred, and the
selected next task.
```
