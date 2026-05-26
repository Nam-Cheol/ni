# Codex Skill Dogfood

This manual test validates the repository-local Codex skills as planning UX.
The CLI remains authoritative for readiness, lock state, hash checks, and prompt
budget enforcement.

## Scope

Validate these skills:

```text
.agents/skills/ni-start/SKILL.md
.agents/skills/ni-end/SKILL.md
.agents/skills/ni-run/SKILL.md
```

The test must not add a Codex adapter, shell adapter, queue, evidence runner,
or execution harness. It must not call `codex exec`. `ni run` compiles a prompt
only.

## Fixture

Use [`docs/examples/codex-skill-dogfood.fixture.md`](examples/codex-skill-dogfood.fixture.md)
as the transcript fixture. It starts from a vague user idea, shows how
`ni-start` should update planning docs plus `.ni/contract.json`, shows
`ni-end` refusing a blocked state or locking only through the CLI, and shows
`ni-run` compiling a prompt without executing it.

## Manual Transcript Plan

### 1. Start From Vague Intent

User:

```text
ni-start
I want a local tool that turns fuzzy project ideas into something an
implementation agent can safely use.
```

Expected skill behavior:

- Read `AGENTS.md`, `.ni/contract.json`, `docs/plan/**`, and the lockfile if it
  exists.
- Extract intent, constraints, risks, evaluations, and blocker questions.
- Update only `docs/plan/**` and `.ni/contract.json` unless the user explicitly
  asks for a different NI maintenance task.
- Preserve stable IDs where possible.
- Run or report `ni status --dir .`.
- Report readiness gaps without declaring readiness by model judgment.

Pass condition:

- Human planning docs and machine contract agree on the same capability,
  requirement, evaluation, risk, and open question IDs.
- Any unresolved blocker remains visible in `ni status` output.

### 2. Refuse Blocked Lock

User:

```text
ni-end
```

Expected skill behavior when the CLI reports `BLOCKED`:

```text
ni status --dir .
```

The skill reports the CLI blockers and stops. It does not run `ni end`, does
not create `.ni/plan.lock.json`, and does not say the plan is ready.

Pass condition:

- The transcript contains the CLI `BLOCKED` result.
- No manual lockfile edit occurs.

### 3. Lock Ready Plan Through CLI

After the blocker question is resolved and `ni-start` updates the planning
state, run:

```text
ni-end
```

Expected skill behavior when the CLI reports `READY` or
`READY_WITH_DEFERRALS`:

```text
ni status --dir .
ni end --dir .
```

The skill confirms that `.ni/plan.lock.json` was created by the CLI and reports
the readiness status plus lockfile path.

Pass condition:

- Lock creation is attributable to `ni end --dir .`.
- The skill does not hand-author, repair, or reinterpret the lockfile.

### 4. Compile Prompt Only

User:

```text
ni-run
```

Expected skill behavior:

```text
ni run --dir . --max-chars 4000
```

The skill reports the command and confirms the prompt is 4000 characters or
less. It does not call `codex exec`, shells, adapters, queues, or downstream
runtimes.

Pass condition:

- The output is a prompt derived from the locked plan.
- Missing or stale lock errors are reported as `BLOCKED`.
- No execution command appears in the transcript.

## Acceptance Checklist

[ ] `ni-start` treats skill output as draft planning UX only.
[ ] `ni-start` preserves CLI authority and exposes blocker questions.
[ ] `ni-end` stops on CLI `BLOCKED`.
[ ] `ni-end` locks only by running `ni end --dir .` after CLI readiness.
[ ] `ni-run` compiles a prompt only and keeps output at or below 4000
characters.
[ ] The transcript contains no Codex adapter, shell adapter, or execution
harness behavior.
[ ] `bash scripts/quality.sh` passes after the documentation and skill wording
changes.

## Current Repository Check

The repository-local skills include explicit guard text:

- `ni-start`: `ni status` wins over model interpretation.
- `ni-end`: `ni status` and `ni end` are the only readiness and lock
  authorities.
- `ni-run`: stale or missing locks stop prompt compilation; the skill must not
  synthesize a replacement prompt.
