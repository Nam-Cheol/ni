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
| User-global skill pack | Install the same `ni` UX into a user's global model workspace skill folder. | Planned; Codex global installation is not claimed because it has not been verified here. | Must not imply the CLI is bundled, installed, or bypassed. |
| Downloadable zip skill pack | Provide a portable archive with skills, prompts, and README instructions. | Planned. | The archive packages UX only; users still need a trusted way to invoke the CLI gates. |
| Manual copy-paste workflow | Let no-terminal users copy planning text into a model workspace and paste CLI-produced proofs back into the conversation. | Experimental as a workflow pattern; not a supported no-terminal product path. | A model may not declare readiness, lock, or handoff from copied text unless the copied result came from the CLI. |
| Future package installer | Install or update model workspace packs through a future package mechanism. | Planned. | Installer work is distribution infrastructure, not `ni-kernel` execution behavior. |

## Environment Matrix

| Environment | Repo-local skills | User-global pack | Downloadable zip pack | Manual copy-paste | Future package installer |
| --- | --- | --- | --- | --- | --- |
| Codex-style skill folders | Available: `.agents/skills/ni-start`, `.agents/skills/ni-end`, and `.agents/skills/ni-run` exist in this repo. | Unverified: no global Codex install path is verified or documented as available. | Planned: can reuse the repo-local skill shape after packaging is defined. | Experimental: users can paste docs and CLI outputs, but repo-local skills are the preferred current path. | Planned. |
| Claude Skills / slash commands | Unverified: this repo has no Claude-specific skill or slash-command package. | Unverified: no Claude distribution path is verified. | Planned only after files match Claude's required structure and are tested there. | Experimental: generic instructions can be copied, but Claude-specific distribution is not claimed. | Planned. |
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
| `ni-status-review` | Explain `ni status` or `ni status --proof` output, identify blockers, and suggest the next planning question. | No standalone skill file exists yet. | Planned. |
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

### User-global skill pack

A user-global pack should mirror the repo-local behavior but live outside a
specific project checkout. It remains planned until the global install location,
update behavior, and host-specific loading rules are verified.

It must include:

- the same behavioral instructions as the repo-local skills;
- a visible requirement to locate the project root before editing;
- a visible requirement to run or request `ni status` before readiness claims;
- a visible requirement to stop on lock hash mismatches;
- an update and uninstall story.

### Downloadable zip skill pack

A zip pack should be a portable archive, not an installer that changes kernel
behavior. The expected archive shape is:

```text
ni-model-workspace-pack/
  README.md
  codex/
    skills/
      ni-start/SKILL.md
      ni-end/SKILL.md
      ni-run/SKILL.md
  generic/
    instructions.md
    status-review.md
    manual-copy-paste.md
```

Claude-specific files should be added only after the repository verifies the
required Claude Skills or slash-command structure. Until then, the zip may
include generic instructions for Claude users, but it must not claim Claude
Skills distribution.

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
- Claude Skills or slash-command distribution must not be described as
  available until files match Claude's required structure and are tested there.
- Downloadable zip packs must not be described as available until the archive is
  produced, inspected, and documented.
- Manual copy-paste may be described as an experimental workflow, not as a
  complete no-terminal product.
- Future package installer work must stay planned until there is a real
  implementation and validation path.
