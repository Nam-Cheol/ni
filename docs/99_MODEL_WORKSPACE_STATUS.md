# Model Workspace Status

## Current status

Model workspace packs are **Experimental** as a broad product path.

Repo-local skill files, package source folders, zip packaging scripts, and
metadata checks are verified in this repository. Host-level/global install,
provider runtime behavior, and cross-machine installation are not verified by
this repository unless a later host-specific verification document says so.

## What is verified

| Evidence | Status | Verification |
| --- | --- | --- |
| Repo-local skill files | Verified | `.agents/skills/**/SKILL.md` |
| Claude skill package source | Verified | `packages/claude-skills/**` |
| Codex skill package source | Verified | `packages/codex-skills/**` |
| Skill metadata | Verified | `bash scripts/check-skill-packs.sh` |
| Claude zip package | Verified | `bash scripts/package-claude-skills.sh` |
| Codex zip package | Verified | `bash scripts/package-codex-skills.sh` |
| CLI authority wording | Verified | skill docs and README files |

## What is not verified

| Claim | Status | Reason |
| --- | --- | --- |
| Global Claude install | not_verified | host-level install was not tested |
| Global Codex install | not_verified | host-level install was not tested |
| Provider runtime behavior | not_verified | no provider API or host behavior was tested |
| Cross-machine install | not_verified | no multi-machine install matrix |
| Skills replace CLI validation | false | CLI remains authority |

## Status vocabulary

- **Experimental:** pack source and packaging are available, but host-level
  install or provider behavior is not fully verified.
- **Available:** may only be used for a specific host path after install and
  usage verification for that path.
- **not_verified:** no claim should be made.
- **UX layer:** model instructions that help authoring but do not decide
  readiness.
- **CLI authority:** `ni status`, `ni end`, `ni run`, lock hashes, and prompt
  compiler remain authoritative.

## Rules for README/docs

- Do not say "Available" for model workspace packs as a broad product path
  unless host-level install is verified.
- Do not say global install works unless tested.
- Do not imply skills can lock, validate, or compile without CLI.
- Do not imply provider behavior is guaranteed.
- Do not imply no-terminal mode is deterministic.
- Do not hide Experimental status behind marketing copy.

## Rules for skills

Each NI skill should state or preserve:

- Skills are UX; CLI is authority.
- Do not execute downstream work.
- Do not modify `.ni/plan.lock.json` manually.
- Do not approve readiness by model judgment.
- Run or request `ni status --proof --next-questions` when relevant.
- Use the user's current language for questions, while preserving IDs,
  commands, file paths, and status constants.

## How status may become Available later

1. Choose one host environment.
2. Install skill pack using documented method.
3. Confirm the host discovers the skills.
4. Invoke `ni-start`, `ni-grill`, `ni-end`, and `ni-run`.
5. Confirm skills preserve CLI authority.
6. Confirm no downstream execution.
7. Record verification doc.
8. Update only that host path to Available.

## Boundary

Model workspace packs do not add runtime execution, model APIs, adapters,
downstream agents, queues, or task running.
