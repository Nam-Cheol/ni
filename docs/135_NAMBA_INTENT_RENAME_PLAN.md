# Namba Intent Rename Plan

## Current status

State:
- v0.5.1 release: published and verified.
- current product/CLI name: ni.
- observed Windows conflict: PowerShell `ni -> New-Item`.
- existing namba-ai command: namba.
- Homebrew: Planned / v0.5 candidate.
- Windows real-host verification: pending.
- Skills are UX; CLI is authority.
- ni is currently a pre-runtime Project Intent Compiler for AI Agents.

## Rename goal

The rename should avoid the PowerShell built-in `ni` alias conflict and make
this project clearly distinct from the existing namba-ai project.

The future product identity is:

- Product name: Namba Intent.
- CLI command: `namba-intent`.
- Legacy/internal name: `ni`, transitional only.
- Config directory: keep `.ni/` for compatibility.
- Repository: keep `Nam-Cheol/ni` for now; repo rename is deferred.

The tagline remains:

```text
Don't run the agent yet.
Compile the intent first.
```

## Decision

RENAME_TO_NAMBA_INTENT

Justification: the current `ni` command has a real PowerShell alias conflict,
and the `namba` command already belongs to the existing namba-ai project. The
next product-facing name should be Namba Intent, and the next primary CLI
command should be `namba-intent`.

## Name decision

| Area | Decision | Reason | Notes |
| --- | --- | --- | --- |
| product name | Namba Intent | Keeps the Namba product family while describing the pre-execution intent compiler. | Product-facing copy should move from `ni` to Namba Intent in v0.6.0. |
| CLI command | `namba-intent` | Avoids PowerShell `ni -> New-Item` and avoids the existing namba-ai `namba` command. | v0.6.0 should make this the primary command. |
| repo name | Keep `Nam-Cheol/ni` for now. | Repo rename adds release, install, and URL churn beyond the command rename. | Deferred until command migration is stable. |
| config directory | Keep `.ni/`. | Existing locks, contract paths, docs, tests, and user workspaces rely on it. | `.ni/` rename is explicitly deferred. |
| old command compatibility | Prefer a compatibility shim for one transition release. | Reduces breakage for current macOS/Linux users while letting Windows avoid the unsafe short command. | Shim must warn: `ni is deprecated; use namba-intent`. |
| release version | v0.6.0. | Changing the primary command is user-facing and should not be a tiny patch. | Release notes must call out the migration clearly. |

## Difference from namba-ai

| Surface | NambaAI | Namba Intent | Notes |
| --- | --- | --- | --- |
| command | `namba` | `namba-intent` | Do not use `namba` for this project. |
| purpose | Codex workflow, SPEC execution, queue, sync, PR, and land flows. | Pre-execution project intent compile, readiness, lock, and handoff prompt. | Product names may share family branding, but command scope must stay separate. |
| execution behavior | May run implementation workflows. | Does not run implementation. | Namba Intent remains a pre-runtime control layer. |
| planning behavior | SPEC/workflow-oriented project execution planning. | Intent Lock Protocol: docs contract, readiness gate, lockfile, prompt compiler, source-of-truth rule. | Kernel remains authoritative. |
| queue / PR / release automation | In scope for namba-ai workflows. | Out of scope. | Namba Intent must not become queue, PR automation, release automation, or downstream execution layer. |
| config directory | `.namba/` in namba-ai. | `.ni/` remains compatible. | Do not rename `.ni/` in v0.6.0. |
| target user | Users running Codex/SPEC workflow operations. | Users and agents that need to compile and lock intent before downstream work. | The boundary is before execution, not during execution. |

## Name usage audit

The required scan was run:

```bash
rg -n "\bni\b|ni.exe|ni init|ni status|ni end|ni run|Nam-Cheol/ni|NambaAI|namba-ai|\bnamba\b|\.ni" README.md README.ko.md docs install.sh install.ps1 scripts cmd internal packages .agents
```

Observed surfaces:

| Surface | Current usage | Migration implication |
| --- | --- | --- |
| README and README.ko | Product name, installer URLs, Windows notes, command examples, and flow examples use `ni`. | Update product copy to Namba Intent and command examples to `namba-intent`, while preserving `.ni/` compatibility notes. |
| docs | Install, release, Homebrew, Windows alias, model workspace, and readiness docs use `ni` heavily. | Update current/future docs carefully; historical release docs may keep old names with clear context. |
| install.sh | Installs `ni`, uses `Nam-Cheol/ni`, and prints `ni --help` / `ni version`. | Change primary installed binary to `namba-intent`; decide whether to install a deprecated `ni` shim on safe platforms. |
| install.ps1 | Installs `%LOCALAPPDATA%\ni\bin\ni.exe` and manages the PowerShell alias workaround. | Prefer installing `namba-intent.exe`; avoid relying on a Windows `ni` shim because of alias behavior. |
| scripts | Validation, release, install, and demo scripts assume `ni` archive and command names. | Update checkers and release scripts in one migration pass. |
| cmd/internal tests | Help text, init text, protected `.ni` paths, stale-lock messages, and run prompts use `ni`. | Update command/product text without changing kernel semantics or `.ni` paths. |
| packages and .agents | Skill names, examples, zip names, and CLI-authority wording use `ni`. | Update skill pack user-facing examples; keep skill authority text tied to the renamed CLI. |

This audit shows the rename should not be a blind text replacement. Historical
v0.5.x evidence should keep its original command where it documents actual past
behavior, while v0.6.0 user-facing instructions should promote Namba Intent and
`namba-intent`.

## Migration scope

| Surface | Change | Include in v0.6.0? | Notes |
| --- | --- | --- | --- |
| binary name | Build and distribute `namba-intent`. | Yes | Primary command for all new docs and installs. |
| install.sh | Install `namba-intent`; optionally install a deprecated `ni` shim on safe Unix-like hosts. | Yes | Must keep uninstall reversible and scoped to installer-managed files. |
| install.ps1 | Install `namba-intent.exe`. | Yes | Avoid Windows `ni` shim unless a real-host transcript proves safe behavior. |
| README | Rename product-facing identity and command examples. | Yes | Keep release truth and Homebrew boundaries honest. |
| README.ko | Korean companion update. | Yes | Preserve exact boundary phrases such as `Skills are UX; CLI is authority.` |
| docs | Update forward-looking docs and migration docs. | Yes | Historical evidence docs may retain `ni` as past-state proof. |
| package skills | Update skill pack examples and authority language. | Yes | Skill names may need a separate compatibility decision. |
| release assets | Rename archives from `ni_...` to `namba-intent_...`. | Yes | Checksums must match the new asset names. |
| checksums | Generate checksums for new asset names. | Yes | Do not reuse old checksum files. |
| CI / release scripts | Update artifact, command, and checker expectations. | Yes | Ensure `go test`, install checks, and release checks use the new primary command. |
| `.ni/` | Keep unchanged. | No | Compatibility path; explicit deferred rename. |
| Homebrew | Do not publish formula in this rename task. | No | Formula naming should be revisited after command migration. |
| repo rename | Keep `Nam-Cheol/ni`. | No | Deferred to avoid URL/install churn in the command migration. |

## Compatibility policy

Recommendation: B. compatibility shim.

`namba-intent` should be the primary command in v0.6.0. A deprecated `ni` shim
may remain for one transition release on platforms where it is safe and
maintainable. The shim should print a warning before delegating:

```text
ni is deprecated; use namba-intent.
```

Windows should avoid relying on `ni` because PowerShell already defines
`ni -> New-Item`. If any Windows compatibility shim is considered, it must be
validated on a real Windows PowerShell host before being documented.

Compatibility rules:

- New docs should use `namba-intent`.
- Existing `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json`, and
  `docs/plan/**` behavior must remain compatible.
- `ni run` behavior after rename becomes `namba-intent run`: bounded handoff
  prompt compilation only, not downstream execution.
- No deprecated alias may become kernel-owned execution state.

## Risks

| Risk | Impact | Mitigation |
| --- | --- | --- |
| breaking existing users | Current users may have scripts or habits around `ni`. | Provide one transition release with a warning shim where safe, plus clear release notes. |
| confusion with existing namba-ai | Users may assume `namba-intent` is the same as `namba`. | Keep the distinction table in README, docs, and release notes. |
| docs drift | Some docs describe historical v0.5.x behavior while others describe v0.6.0. | Separate historical evidence from current instructions and use explicit version context. |
| install script migration | Install/uninstall may leave old binaries, PATH blocks, or profile blocks. | Add migration tests for fresh install, update install, uninstall, and old-binary cleanup. |
| release asset naming | Installers and checksums may point to old `ni_...` assets. | Update GoReleaser, install scripts, and release checks together. |
| Windows PowerShell behavior | `ni` invokes `New-Item` unless the alias is removed. | Make `namba-intent.exe` primary on Windows and keep Windows `ni` compatibility deferred unless real-host proof exists. |
| Homebrew formula naming | Homebrew core already has an unrelated `ni` formula. | Defer Homebrew and evaluate `namba-intent` formula naming separately. |

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| namba-ai distinction | Must not use `namba` for Namba Intent. | This plan reserves `namba` for namba-ai and chooses `namba-intent`. | Yes | Avoids CLI collision. |
| Namba Intent identity | Product rename plan only; implementation deferred. | This document chooses the future identity but does not rename code. | Yes | v0.6.0 implementation is a later task. |
| `run` behavior | Prompt compilation only. | Preserved as bounded handoff prompt compilation. | Yes | No downstream execution added. |
| Homebrew | Planned / v0.5 candidate only. | Preserved. | Yes | No formula publication or Available claim. |
| Windows verification | Pending until transcript exists. | Preserved. | Yes | Static installer checks are not real-host proof. |
| runtime execution boundary | No task runner, SPEC runner, shell/Codex adapter, queue, PR automation, release automation, or downstream execution layer. | Preserved. | Yes | Rename does not change kernel scope. |

## Recommended next task

A. implement Namba Intent rename

## Next task prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Task: Implement the v0.6.0 Namba Intent rename.

Use docs/135_NAMBA_INTENT_RENAME_PLAN.md and
docs/135_NAMBA_INTENT_RENAME_PLAN.ko.md as the authoritative rename plan.

Decision:
- Product name: Namba Intent.
- Primary CLI command: namba-intent.
- Do not use namba as this project's CLI command.
- Keep .ni/ for compatibility.
- Keep repository Nam-Cheol/ni for now.
- Prefer a compatibility shim for one transition release only where safe.
- Windows should use namba-intent.exe as the primary command and must not rely
  on ni unless real-host PowerShell behavior is proven safe.

Scope:
- Rename product-facing current instructions from ni to Namba Intent where they
  describe the current v0.6.0+ product.
- Rename primary command examples from ni to namba-intent.
- Update Go command metadata, help text, tests, installers, release asset names,
  checksums expectations, docs, README, README.ko, scripts, package skills, and
  .agents examples needed for the new primary command.
- Preserve historical v0.5.x evidence docs where they describe past verified
  ni behavior, but add version context if needed.
- Preserve .ni/contract.json, .ni/session.json, .ni/plan.lock.json, schema
  paths, lockfile path, and docs/plan/** behavior.

Compatibility policy:
- namba-intent is primary.
- A deprecated ni shim may be kept for a transition release on platforms where
  it is safe and maintainable.
- The shim must warn: ni is deprecated; use namba-intent.
- Do not make Windows ni compatibility a claim without a real Windows
  PowerShell transcript.

Forbidden:
- Do not rename .ni/.
- Do not rename the repository.
- Do not publish, tag, create a GitHub release, upload assets, run release
  workflows, create or publish a Homebrew formula, or mark Homebrew Available.
- Do not run ni end on the project root.
- Do not relock the project root.
- Do not execute generated prompts.
- Do not add runtime execution behavior.
- Do not make run execute downstream work.
- Do not add task runner, SPEC runner, queue, PR automation, release automation,
  shell adapter, or Codex exec adapter behavior.

Validation:
- git status --short
- gofmt -w .
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Final response:
- changed files
- primary rename behavior
- compatibility shim behavior
- validation results
- protected .ni diff result
- confirmation that no publication, tag, release, upload, root relock, prompt
  execution, Homebrew publication, or downstream execution behavior occurred
```
