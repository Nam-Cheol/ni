# Codex Skill Pack

The Codex skill pack makes `ni` usable inside a repo-local Codex workspace as a
model workspace workflow.

It is a UX distribution. It does not change `ni-kernel`, does not call
`codex exec`, and does not execute downstream work. The `ni` CLI remains the
authority for readiness, lock creation, lock hash verification, and prompt
compilation.

```text
Codex skills -> conversation workflow
ni CLI -> readiness gate, lockfile, hash check, prompt compiler
```

## Pack Contents

```text
packages/codex-skills/
  README.md
  README.ko.md
  ni-start/SKILL.md
  ni-status-review/SKILL.md
  ni-end/SKILL.md
  ni-run/SKILL.md
```

The downloadable archive is produced by:

```bash
bash scripts/package-codex-skills.sh
```

The output path is:

```text
dist/ni-codex-skills.zip
```

## Skill Behavior

| Skill | Role | Authority boundary |
| --- | --- | --- |
| `ni-start` | Conversation-driven authoring for `docs/plan/**`, `.ni/contract.json`, and `.ni/session.json`. | Must run or request `ni status --dir . --next-questions`; must not lock. |
| `ni-status-review` | Explain `ni status --proof` output and next questions. | Must preserve CLI status exactly; must not become a second readiness engine. |
| `ni-end` | Review readiness and ask for explicit lock confirmation. | Must run or request `ni status --dir .` before `ni end --dir .`; must not write `.ni/plan.lock.json` manually. |
| `ni-run` | Compile a bounded handoff prompt from a locked plan. | Must run or request `ni run --dir . --target <target> --max-chars 4000`; must not execute downstream work. |

## Repo-Local Installation

Repo-local usage is the verified path. Copy the skill directories into the
workspace-local `.agents/skills/` directory:

```bash
mkdir -p .agents/skills
cp -R packages/codex-skills/ni-start .agents/skills/
cp -R packages/codex-skills/ni-status-review .agents/skills/
cp -R packages/codex-skills/ni-end .agents/skills/
cp -R packages/codex-skills/ni-run .agents/skills/
```

The package does not install or invoke the `ni` binary. It only provides skill
instructions. The skills must run or request the `ni` CLI commands that supply
authority.

Global Codex skill installation and discovery paths are not claimed here. Treat
global usage as experimental or planned until verified in a specific Codex
environment.

## Boundaries

The Codex pack must not:

- execute implementation;
- call `codex exec`;
- add shell or Codex adapters;
- add evidence runners, queues, PR automation, release automation, or model
  orchestration;
- manually create, edit, or repair `.ni/plan.lock.json`;
- weaken readiness blockers to reach a ready state;
- generate downstream-owned execution state.

If the CLI reports a stale lock or hash mismatch, every skill must stop and
report `BLOCKED`.

## Validation

The repository quality check validates skill metadata for repo-local skills and
packaged skill distributions:

```bash
bash scripts/quality.sh
```

The packaging check verifies the archive can be produced:

```bash
bash scripts/package-codex-skills.sh
```
