# NI Claude Skill Pack

This package contains Claude Skill-compatible NI workflow instructions.

The skills help a Claude Code or Claude Skill-compatible workspace author,
review, lock, and compile NI planning contracts. They are UX only. The `ni`
CLI remains the authority for readiness, lock creation, lock hash verification,
and prompt compilation.

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
- Do not call Claude APIs.
- Do not execute downstream work.
- Do not add shell/Codex adapters, evidence runners, queues, model
  orchestration, PR automation, release automation, TUI, or web UI behavior.

## Packaging

From the repository root:

```bash
bash scripts/package-claude-skills.sh
```

The archive is written to:

```text
dist/ni-claude-skills.zip
```

## Installation

Claude-compatible environments may use different skill folder locations. This
repository does not assume a global path.

Use a target directory that your environment documents and that you have
verified:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

The installer copies skill directories only after the target is supplied. It
preserves existing skill directories unless `--force` is passed.

