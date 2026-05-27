# Claude Skill Pack

The Claude skill pack makes `ni` usable inside Claude Code or Claude
Skill-compatible workspaces as a model workspace workflow.

It is a UX distribution. It does not change `ni-kernel`, does not call Claude
APIs, and does not execute downstream work. The `ni` CLI remains the authority
for readiness, lock creation, lock hash verification, and prompt compilation.

```text
Claude skills -> conversation workflow
ni CLI -> readiness gate, lockfile, hash check, prompt compiler
```

## Pack Contents

```text
packages/claude-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

The downloadable archive is produced by:

```bash
bash scripts/package-claude-skills.sh
```

The output path is:

```text
dist/ni-claude-skills.zip
```

## Skill Behavior

| Skill | Role | Authority boundary |
| --- | --- | --- |
| `ni-start` | Conversation-driven authoring for `docs/plan/**`, `.ni/contract.json`, and `.ni/session.json`. | Must run or request `ni status --dir . --next-questions`; must not lock. |
| `ni-status-review` | Explain `ni status --proof` output and next questions. | Must preserve CLI status exactly; must not become a second readiness engine. |
| `ni-end` | Review readiness and ask for explicit lock confirmation. | Must run or request `ni status --dir .` before `ni end --dir .`; must not write `.ni/plan.lock.json` manually. |
| `ni-run` | Compile a bounded handoff prompt from a locked plan. | Must run or request `ni run --dir . --target <target> --max-chars 4000`; must not execute downstream work. |

## Installation

Claude-compatible environments may use different skill folder locations. This
repository does not assume or document a global default.

Use a directory that the user's environment documents and that the user has
verified:

```bash
bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
bash scripts/install-claude-skills.sh --target /path/to/skills
```

The installer has these safety properties:

- `--target` is required.
- `--dry-run` prints the copy operations without changing files.
- Existing skill directories are preserved unless `--force` is passed.
- The script copies skill files only; it does not install or invoke the `ni`
  binary.
- The script does not call Claude APIs.

## Boundaries

The Claude pack must not:

- execute implementation;
- add shell or Codex adapters;
- add evidence runners, queues, PR automation, release automation, or model
  orchestration;
- manually create, edit, or repair `.ni/plan.lock.json`;
- weaken readiness blockers to reach a ready state;
- generate downstream-owned execution state.

If the CLI reports a stale lock or hash mismatch, every skill must stop and
report `BLOCKED`.

## Validation

The repository quality check validates skill metadata for both repo-local Codex
skills and this Claude skill pack:

```bash
bash scripts/quality.sh
```

The packaging check verifies the archive can be produced:

```bash
bash scripts/package-claude-skills.sh
```

