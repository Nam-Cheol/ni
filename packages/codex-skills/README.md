# NI Codex Skill Pack

This package contains Codex-style NI workflow skills for repo-local use.

The skills help a Codex workspace author, review, lock, and compile NI planning
contracts. They are UX only. The `ni` CLI remains the authority for readiness,
lock creation, lock hash verification, and prompt compilation.

## Status

Status: Experimental.

Verified: source files, package zip, and metadata checks.

Not verified: global host install, provider behavior, cross-machine install,
or global Codex skill discovery.

Boundary: Skills are UX; CLI is authority.

## Skills

| Skill | Purpose |
| --- | --- |
| `ni-start` | Continue conversation-driven planning and keep `docs/plan/**`, `.ni/contract.json`, and `.ni/session.json` synchronized. |
| `ni-grill` | Challenge accepted or nearly accepted planning content before lock; it does not execute work or approve readiness by model judgment. |
| `ni-status-review` | Explain `ni status --proof` output and identify the next planning question without becoming a second readiness engine. |
| `ni-end` | Review CLI readiness, ask for explicit confirmation, and lock only through `ni end`. |
| `ni-run` | Compile a bounded handoff prompt from a valid lock without executing downstream work. |

## Authority Rules

- Skills are UX; the CLI is authority.
- Run or request `ni status` before any readiness claim.
- `ni-start` must use grouped `ni status --proof --next-questions` output as
  its primary planning interview when present.
- `ni-start` must ask user-facing planning questions in the language of the
  user's latest substantive message while preserving IDs, commands, paths,
  schema keys, target names, and status constants exactly.
- `ni-start` must show a concise planning proof block after meaningful
  authoring updates, naming changed files, affected IDs, before/after CLI
  status, remaining blockers, and the next question group.
- Skills may help draft or explain proof-related planning text.
- Skills do not determine readiness.
- Skills do not lock plans.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- `ni-grill` challenges planning quality before lock. It does not execute work.
- If `ni status` is `BLOCKED`, `ni-grill` should use deterministic blockers
  before inventing new critique.
- `ni-grill` must use `Critical`, `High`, `Medium`, `Low`, and `Note`
  severity labels, show at most 5 findings by default, and show at most 3
  `Critical`/`High` findings first when they exist.
- `ni-grill` must ask user-facing questions in the user's latest substantive
  language while preserving IDs, commands, paths, status constants, and
  severity labels.
- `ni-grill` never approves lock by model judgment.
- Run or request `ni end` before any lock claim.
- Run or request `ni run` before any compiled handoff prompt claim.
- Never edit `.ni/plan.lock.json` manually.
- Stop and report `BLOCKED` on stale locks or hash mismatches.
- Do not call `codex exec`.
- Do not execute downstream work.
- Do not add shell/Codex adapters, evidence runners, queues, model
  orchestration, PR automation, release automation, TUI, or web UI behavior.

## Packaging

From the repository root:

```bash
bash scripts/package-codex-skills.sh
```

The archive is written to:

```text
dist/ni-codex-skills.zip
```

## Copy This Folder

This repository verifies repo-local skill usage only. Copy the skill directories
into a workspace-local `.agents/skills/` directory:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-grill .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

From the zip archive, unpack first, then copy the same skill folders:

```bash
unzip -q dist/ni-codex-skills.zip -d /tmp/ni-codex-skills-unpacked
mkdir -p .agents/skills
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-start .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-grill .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-status-review .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-end .agents/skills/
cp -R /tmp/ni-codex-skills-unpacked/ni-codex-skills/ni-run .agents/skills/
```

Then run the relevant `ni` CLI commands from the project workspace when a skill
asks for authority.

Global Codex skill installation and discovery paths are not claimed by this
package. Treat global usage as experimental or planned until verified in a
specific Codex environment.

## Manual Copy And Zip Usage

Manual copy is available from this source tree or from the zip archive after it
is unpacked. Copy only the skill directories into a target folder that the user
has verified for the current model host. Do not describe that target as a
global Codex install path unless that host-specific path and loading behavior
have been verified.

Codex dry-run install support is planned.

## Verify The Pack

List the skills:

```bash
find packages/codex-skills -mindepth 1 -maxdepth 1 -type d -name 'ni-*' -print | sort
```

Check the `SKILL.md` files:

```bash
find packages/codex-skills -path '*/SKILL.md' -print | sort
bash scripts/check-skill-packs.sh
```

Package the zip:

```bash
bash scripts/package-codex-skills.sh
```

Inspect the archive:

```bash
unzip -l dist/ni-codex-skills.zip
```

See `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md` for the full installation and
verification status.

See `docs/99_MODEL_WORKSPACE_STATUS.md` for the Experimental broad product
status and not_verified host/provider boundaries.

See `docs/89_LANGUAGE_ADAPTIVE_AUTHORING.md` for the language-adaptive
authoring rule used by `ni-start`.

## What This Does Not Do

- Does not run Codex APIs or `codex exec`.
- Does not execute implementation or downstream work.
- Does not replace `ni` CLI validation for readiness, locking, hash checks, or
  prompt compilation.
