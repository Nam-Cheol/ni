# Model Workspace Packs

This document defines how `ni` can be used inside model workspaces without
treating it as a terminal-only Go CLI.

The strategy is intentionally conservative. Model packs are UX distribution:
they help a model and user author, review, lock, and compile planning state.
They do not become the authority. The CLI remains the deterministic authority
for readiness, lock creation, lock hash verification, and prompt compilation.

```text
skills and instructions -> user/model planning UX
ni CLI -> readiness gate, lockfile, hash check, prompt compiler
```

No pack may add model API calls, execution runtime behavior, downstream
adapters, or hidden implementation steps. If the pack cannot get a valid CLI
result, it must report `BLOCKED` instead of substituting model judgment.

## Status Legend

| Status | Meaning |
| --- | --- |
| Available | Matching files exist in this repository and the workflow can be used in that environment today. |
| Experimental | The workflow can be attempted from existing docs or copied instructions, but it is not a packaged or fully verified distribution path. |
| Planned | The strategy is accepted, but matching files, packaging, or installer support have not been built. |
| Unverified | The environment may support a similar shape, but this repository has not verified the required host-specific file structure or install path. |

## Pack Types

| Pack type | Purpose | Current status | Boundary |
| --- | --- | --- | --- |
| Repo-local `ni` skills | Use skill files checked into the current project workspace. | Available for Codex-style repo-local skills only. | The model may edit planning docs and `.ni/contract.json`; `ni status`, `ni end`, and `ni run` remain authority. |
| User-global skill pack | Install the same `ni` UX into a user's global model workspace skill folder. | Available only through a user-provided, verified target directory for the Claude pack; Codex global installation is not claimed because it has not been verified here. | Must not imply the CLI is bundled, installed, or bypassed. |
| Downloadable zip skill pack | Provide a portable archive with skills, prompts, and README instructions. | Available for the Codex and Claude skill packs after `scripts/check-skill-packs.sh` and the package scripts pass. | The archive packages UX only; users still need a trusted way to invoke the CLI gates. |
| Manual copy-paste workflow | Let no-terminal users copy planning text into a model workspace and paste CLI-produced proofs back into the conversation. | Experimental as a workflow pattern; not a supported no-terminal product path. | A model may not declare readiness, lock, or handoff from copied text unless the copied result came from the CLI. |
| Future package installer | Install or update model workspace packs through a future package mechanism. | Planned. | Installer work is distribution infrastructure, not `ni-kernel` execution behavior. |

## Environment Matrix

| Environment | Repo-local skills | User-global pack | Downloadable zip pack | Manual copy-paste | Future package installer |
| --- | --- | --- | --- | --- | --- |
| Codex-style skill folders | Available: `.agents/skills/ni-start`, `.agents/skills/ni-end`, and `.agents/skills/ni-run` exist in this repo; `packages/codex-skills` also includes `ni-status-review`. | Unverified: no global Codex install path is verified or documented as available. | Available: `scripts/package-codex-skills.sh` creates `dist/ni-codex-skills.zip`. | Experimental: users can paste docs and CLI outputs, but repo-local skills are the preferred current path. | Planned. |
| Claude Skills / slash commands | Available as repository files under `packages/claude-skills`; slash-command behavior is not claimed. | Available only through `scripts/install-claude-skills.sh --target <verified-dir>`; no global Claude path is assumed. | Available: `scripts/package-claude-skills.sh` creates `dist/ni-claude-skills.zip`. | Experimental: generic instructions can still be copied when the host does not load skills. | Planned. |
| Generic model instruction packs | Experimental: a model can read repository docs, but `.agents/skills` is a Codex-style convention. | Planned. | Planned. | Experimental: use visible instructions plus pasted CLI proofs. | Planned. |

The matrix deliberately does not mark global Codex installation or Claude Skills
distribution as available. Those claims require host-specific verification that
has not happened in this repository.

## Skill Surface

The model workspace pack should cover these UX actions.

| UX action | Role | Current repository state | Status |
| --- | --- | --- | --- |
| `ni-start` | Continue planning conversation and update `docs/plan/**`, `.ni/contract.json`, and session continuity from user intent. | `.agents/skills/ni-start/SKILL.md` exists. | Available for repo-local Codex-style skills. |
| `ni-end` | Review CLI readiness, ask for explicit user confirmation, then let the CLI write `.ni/plan.lock.json`. | `.agents/skills/ni-end/SKILL.md` exists. | Available for repo-local Codex-style skills. |
| `ni-run` | Compile a 4000-character-or-less handoff prompt from a valid lock. | `.agents/skills/ni-run/SKILL.md` exists. | Available for repo-local Codex-style skills. |
| `ni-status-review` | Explain `ni status` or `ni status --proof` output, identify blockers, and suggest the next planning question. | `packages/claude-skills/ni-status-review/SKILL.md` and `packages/codex-skills/ni-status-review/SKILL.md` exist. | Available for the Claude and Codex skill packs. |
| `ni-readme-pamphlet-review` | Review README changes against the pamphlet strategy in `docs/52_README_PAMPHLET_STRATEGY.md`. | No standalone skill file exists yet. | Planned and optional. |

`ni-status-review` is useful because status output is the main proof that a
model must preserve rather than reinterpret. It should be a review skill, not a
second readiness engine.

`ni-readme-pamphlet-review` is useful for release and documentation polish, but
it is not part of the Intent Lock Protocol. It should stay optional and should
not be required for locking or prompt compilation.

## Required Pack Contents

### Repo-local Codex-style pack

The repository-local pack is the current available shape:

```text
.agents/skills/ni-start/SKILL.md
.agents/skills/ni-end/SKILL.md
.agents/skills/ni-run/SKILL.md
docs/06_CODEX_SKILLS.md
docs/31_NI_START_BEHAVIOR.md
docs/35_NI_END_CONFIRMATION.md
docs/36_NI_RUN_HANDOFF.md
```

The pack assumes the model workspace can read repository files and can either
run the CLI or ask the user to provide exact CLI output.

The packaged Codex skill source is:

```text
packages/codex-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

### User-global skill pack

A user-global pack should mirror the repo-local behavior but live outside a
specific project checkout. For Claude-compatible environments, this repository
provides a safe copy script that requires a user-provided, verified target
directory. It does not claim or assume a global Claude path.

It must include:

- the same behavioral instructions as the repo-local skills;
- a visible requirement to locate the project root before editing;
- a visible requirement to run or request `ni status` before readiness claims;
- a visible requirement to stop on lock hash mismatches;
- an update and uninstall story.

### Downloadable zip skill pack

A zip pack is a portable archive, not an installer that changes kernel
behavior. The Codex archive shape is:

```text
ni-codex-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

It is produced by:

```bash
bash scripts/package-codex-skills.sh
```

The output path is:

```text
dist/ni-codex-skills.zip
```

The Claude archive shape is:

```text
ni-claude-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

It is produced by:

```bash
bash scripts/package-claude-skills.sh
```

The output path is:

```text
dist/ni-claude-skills.zip
```

### Manual copy-paste workflow

Manual copy-paste is for users who can participate in model-guided planning but
cannot run terminal commands themselves.

The workflow is:

1. The user copies the project goal and current planning docs into the model
   workspace.
2. The model proposes edits to `docs/plan/**` and `.ni/contract.json` as visible
   patches or replacement snippets.
3. A trusted local actor applies the edits and runs `ni status`.
4. The user pastes the exact status proof back into the model workspace.
5. If the status is `BLOCKED`, the model asks the next planning question.
6. If the status is `READY` or `READY_WITH_DEFERRALS`, the model may guide the
   user through `ni-end`, but only the CLI may create the lock.
7. After lock creation, the user pastes the lock result or `ni run` output back
   into the workspace.

The manual workflow is not a claim that `ni` works without the CLI. It is a
way to keep no-terminal users in the loop while a trusted runner produces the
authoritative proof.

### Future package installer

A future package installer may install or update model workspace packs. It
should be treated as distribution infrastructure and should not be required for
the kernel to function.

The installer may:

- copy skill or instruction files into verified host locations;
- report what it installed and where;
- preserve existing user edits by default;
- verify pack version metadata.

The installer must not:

- run downstream implementation;
- call model APIs;
- weaken readiness checks;
- generate or repair lockfiles from model judgment;
- hide CLI failures behind friendly model summaries.

## Authority Rules For Every Pack

- Skills are UX; the CLI is authority.
- The model may draft docs, detect gaps, and propose edits.
- The model may not declare readiness without `ni status`.
- The model may not lock without `ni end`.
- The model may not compile a replacement prompt if `ni run` reports a missing
  or stale lock.
- Open blocker questions prevent locking.
- High-severity risks still require mitigation.
- Prompt output from `ni run` must stay at or below 4000 characters.
- After `.ni/plan.lock.json` exists, source-of-truth precedence is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

If a lock hash mismatch exists, every pack must stop and report `BLOCKED`.

## Availability Rules

- Repo-local Codex-style skills may be described as available in this
  repository.
- Global Codex installation must not be described as available until a real
  install location and loading behavior are verified.
- Claude skill installation must require a user-provided, verified target
  directory; the repository must not assume a global Claude path.
- Claude slash-command behavior must not be described as available because this
  pack provides skills, not slash-command integration.
- Downloadable zip packs may be described as available only when the archive is
  produced, inspected, and documented.
- Manual copy-paste may be described as an experimental workflow, not as a
  complete no-terminal product.
- Future package installer work must stay planned until there is a real
  implementation and validation path.

See [Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.md)
for the current verification command, install paths, and status language.
