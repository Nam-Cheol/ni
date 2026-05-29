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
- `ni-start` must use grouped `ni status --proof --next-questions` output as
  its primary planning interview when present.
- `ni-start` must show a concise planning proof block after meaningful
  authoring updates, naming changed files, affected IDs, before/after CLI
  status, remaining blockers, and the next question group.
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

## Copy This Folder

Claude-compatible environments may use different skill folder locations. This
repository does not assume a global path.

Use a target directory that your environment documents and that you have
verified. Copy only the skill folders into that target:

```bash
TARGET=/path/to/verified/claude-skills
mkdir -p "$TARGET"
cp -R packages/claude-skills/ni-start "$TARGET/"
cp -R packages/claude-skills/ni-status-review "$TARGET/"
cp -R packages/claude-skills/ni-end "$TARGET/"
cp -R packages/claude-skills/ni-run "$TARGET/"
```

From the zip archive, unpack first, then copy the same skill folders:

```bash
unzip -q dist/ni-claude-skills.zip -d /tmp/ni-claude-skills-unpacked
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-start "$TARGET/"
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-status-review "$TARGET/"
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-end "$TARGET/"
cp -R /tmp/ni-claude-skills-unpacked/ni-claude-skills/ni-run "$TARGET/"
```

This is a file-copy workflow only. Do not describe the target as a global
Claude install path unless that specific host path and loading behavior have
been verified.

## Guarded Install Script

The Claude pack also includes a guarded copy script. Use a target directory
that your environment documents and that you have verified:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

The installer copies skill directories only after the target is supplied. It
preserves existing skill directories unless `--force` is passed.

Manual copy from this source tree or from the unpacked zip archive is also
available when the user has verified the target folder for the current Claude
compatible host. Do not describe that target as a global Claude install path.

## Verify The Pack

List the skills:

```bash
find packages/claude-skills -mindepth 1 -maxdepth 1 -type d -name 'ni-*' -print | sort
```

Check the `SKILL.md` files:

```bash
find packages/claude-skills -path '*/SKILL.md' -print | sort
bash scripts/check-skill-packs.sh
```

Package the zip:

```bash
bash scripts/package-claude-skills.sh
```

Inspect the archive:

```bash
unzip -l dist/ni-claude-skills.zip
```

See `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md` for the full installation and
verification status.

## What This Does Not Do

- Does not run Claude APIs.
- Does not execute implementation or downstream work.
- Does not replace `ni` CLI validation for readiness, locking, hash checks, or
  prompt compilation.
