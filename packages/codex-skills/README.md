# NI Codex Skill Pack

This package contains Codex-style NI workflow skills for repo-local use.

The skills help a Codex workspace author, review, lock, and compile NI planning
contracts. They are UX only. The `ni` CLI remains the authority for readiness,
lock creation, lock hash verification, and prompt compilation.

## Skills

| Skill | Purpose |
| --- | --- |
| `ni-start` | Continue conversation-driven planning and keep `docs/plan/**`, `.ni/contract.json`, and `.ni/session.json` synchronized. |
| `ni-status-review` | Explain `ni status --proof` output and identify the next planning question without becoming a second readiness engine. |
| `ni-end` | Review CLI readiness, ask for explicit confirmation, and lock only through `ni end`. |
| `ni-run` | Compile a bounded handoff prompt from a valid lock without executing downstream work. |

## Authority Rules

- Skills are UX; the CLI is authority.
- Run or request `ni status` before any readiness claim.
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

## Repo-Local Usage

This repository verifies repo-local skill usage only. Copy the skill directories
into a workspace-local `.agents/skills/` directory:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

Then run the relevant `ni` CLI commands from the project workspace when a skill
asks for authority.

Global Codex skill installation and discovery paths are not claimed by this
package. Treat global usage as experimental or planned until verified in a
specific Codex environment.
