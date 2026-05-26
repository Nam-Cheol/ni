---
name: ni-start
description: Continue NI planning by updating docs/plan and .ni/contract.json while preserving the CLI as readiness authority.
---

# ni-start

Use this skill when the user says `ni-start` or asks to continue planning a project with NI.

`ni-start` is the primary authoring UX after `ni init`. It turns a sustained
model-user planning conversation into synchronized updates to `docs/plan/**`
and `.ni/contract.json`.

## Authority boundary

You are not the authority for readiness or lock state. The `ni` CLI is the authority.

Do not say the plan is complete unless `ni status` has no blockers. If `.ni/plan.lock.json` exists, do not silently edit locked planning docs; first state that planning is locked and proceed only when the user explicitly resumes planning.

If your interpretation conflicts with `ni status`, report the CLI result and stop. Do not override, reinterpret, or soften a CLI `BLOCKED` result.

## Task

Maintain planning state only across these files:

```text
docs/plan/**
.ni/contract.json
```

The skill may read supporting files such as `AGENTS.md`, `.ni/project.json`,
`.ni/readiness.rules.json`, `.ni/readiness.profiles.json`, and
`.ni/plan.lock.json` to understand authority, profile, and lock state. It must
not make downstream harness state the source of truth.

## Start-of-turn process

1. Read `AGENTS.md`, `.ni/contract.json`, `docs/plan/**`, and `.ni/plan.lock.json` if it exists.
2. Summarize current planning state in a few concrete bullets:
   - purpose and delivery surface,
   - accepted capabilities,
   - known decisions and non-goals,
   - open blocker questions,
   - readiness state if `ni status --dir .` has already been run.
3. Identify missing required planning areas from the current docs, contract,
   and readiness profile. Check for:
   - missing or TODO purpose, actors, capabilities, requirements, risks,
     evaluations, artifacts, constraints, delivery expectations, non-goals, or
     open questions,
   - accepted capabilities without linked requirements or evaluations,
   - high-severity risks without mitigation,
   - blocker questions that still affect acceptance criteria or scope.
4. Ask focused questions about the highest-impact gaps. Do not ask broad
   generic questions such as "What else should we add?" Prefer one to three
   specific questions tied to IDs, blockers, or missing plan areas.

## Authoring loop

After the user answers:

1. Extract new or corrected purpose, actors, capabilities, requirements,
   decisions, risks, evaluations, non-goals, constraints, artifacts,
   assumptions, and open questions.
2. Classify each record as accepted, draft, assumption, deferred, rejected, or
   blocker. Tentative, inferred, conflicting, or incomplete statements stay
   visible as assumptions or open questions.
3. Update Markdown docs for humans.
4. Update `.ni/contract.json` for machine validation.
5. Preserve stable IDs where possible.
6. Add new IDs only when necessary.
7. Keep trace links intact: capabilities should point to their requirements,
   evaluations, risks, and artifacts when those records exist.
8. Record decisions, assumptions, risks, non-goals, and open questions in both
   the relevant plan docs and contract fields when they affect readiness.
9. Run or report `ni status --dir .` at the end when available.
10. Show readiness gaps and blocker questions from the CLI result.

Continue this loop until `ni status` reports no blocking issues. Suggest
`ni-end` only when the readiness gate reports `READY` or
`READY_WITH_DEFERRALS`. Never declare completion by model judgment alone.

## Output style

When responding during planning:

- Lead with the current planning summary or what changed.
- Name the files changed when an answer updates planning state.
- Ask only the next focused questions needed to unblock readiness.
- If `ni status` reports `BLOCKED`, state the blockers plainly and keep
  planning open.
- If a lock hash mismatch exists, stop and report `BLOCKED`.

## Example turn shape

```text
Current state:
- Purpose: ...
- Accepted capabilities: CAP-001, CAP-002
- Blocking gaps: OQ-003 needs an owner; CAP-002 has no evaluation

Focused questions:
1. For CAP-002, should acceptance be measured by a transcript fixture or a CLI
   smoke test?
2. Who owns the blocker decision in OQ-003?

After you answer, I will update docs/plan/** and .ni/contract.json, then run
ni status --dir .
```

## Do not

- Do not create `.ni/plan.lock.json`.
- Do not run implementation work.
- Do not create a shell or Codex adapter.
- Do not write generated harness files.
- Do not hide blocker questions.
- Do not weaken evaluations to make the plan appear ready.
- Do not edit files outside `docs/plan/**` and `.ni/contract.json` unless the user explicitly asks for a different NI maintenance task.
- Do not add user-facing contract `add`, `list`, or `set` commands.
- Do not create SPEC runner behavior.
- Do not directly call downstream runtimes.
