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

Do not say the plan is complete unless `ni status` has no blockers. If `.ni/plan.lock.json` exists, do not silently edit locked planning docs; first state that planning is locked and proceed only through the amendment or relock flow when the user explicitly resumes planning.

If your interpretation conflicts with `ni status`, report the CLI result and stop. Do not override, reinterpret, or soften a CLI `BLOCKED` result.

## Edit discipline

Keep planning edits narrow, visible, and grounded in the user's confirmed
intent.

- Minimal diff rule: change only the files, records, and sections needed for
  the current planning turn. Do not rewrite unrelated prose, reorder stable
  records, or renumber IDs for style.
- Assumption vs decision rule: tentative, inferred, conflicting, or incomplete
  statements stay visible as assumptions, draft records, or open questions.
  They do not become accepted decisions.
- User-confirmed decision rule: accepted decisions require explicit user
  confirmation or already-established accepted planning state. You may propose
  decision wording, but your proposed wording is not acceptance by itself.
- No silent deletion rule: do not remove planning records without making the
  change visible. Prefer marking records rejected, deferred, resolved, or not
  applicable when history matters.
- Lock safety rule: after `.ni/plan.lock.json` exists, do not edit locked
  `docs/plan/**` content or matching `.ni/contract.json` records except through
  the amendment or relock flow. If a lock hash mismatch exists, stop and report
  `BLOCKED`.
- Risk and evaluation integrity rule: do not weaken risks, mitigations,
  requirements, evaluations, or non-goals to reach readiness.
- Change summary rule: after updating docs, show a short summary naming changed
  files, affected IDs, and any remaining assumptions, blockers, or readiness
  gaps.

## Task

Maintain planning state across these files:

```text
docs/plan/**
.ni/contract.json
.ni/session.json
```

The skill may read supporting files such as `AGENTS.md`, `.ni/project.json`,
`.ni/readiness.rules.json`, `.ni/readiness.profiles.json`, and
`.ni/plan.lock.json` to understand authority, profile, and lock state. It must
not make downstream harness state the source of truth.

`.ni/session.json` is persistent carryover context only. It is below
`.ni/contract.json` and `docs/plan/**` in the source-of-truth order, must not
override locked docs, must not mark docs complete, and must not store full raw
chat logs by default.

## Start-of-turn process

1. Read `AGENTS.md`, `.ni/contract.json`, `docs/plan/**`, `.ni/session.json` if it exists, and `.ni/plan.lock.json` if it exists.
2. Summarize current planning state in a few concrete bullets:
   - active planning focus from `.ni/session.json`, verified against docs and contract,
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
4. Run or request `ni status --dir . --next-questions` when available.
5. Ask focused questions about the highest-impact gaps. Prefer one to three
   questions from the CLI `next_questions` result. You may lightly rephrase for
   clarity, but preserve the referenced IDs, readiness gap, and allowed outcomes.
   Do not ask broad generic questions such as "What else should we add?"

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
5. Update `.ni/session.json` as bounded continuity state:
   - active planning focus,
   - last planning summary,
   - pending questions,
   - recent decisions,
   - recent risks,
   - last readiness status,
   - last readiness blockers,
   - last updated docs,
   - the explicit note that raw transcript is not the source of truth.
6. Preserve stable IDs where possible.
7. Add new IDs only when necessary.
8. Keep trace links intact: capabilities should point to their requirements,
   evaluations, risks, and artifacts when those records exist.
9. Record decisions, assumptions, risks, non-goals, and open questions in both
   the relevant plan docs and contract fields when they affect readiness.
10. Preserve existing risks, mitigations, requirements, evaluations, and
    non-goals unless the user explicitly changes them.
11. Run or report `ni status --dir . --next-questions` at the end when available.
12. Show a short change summary with changed files and affected IDs.
13. Show readiness gaps and next questions from the CLI result.
14. Reflect the CLI readiness status and blockers back into `.ni/session.json`
    without treating the session file as readiness authority.

Continue this loop until `ni status` reports no blocking issues. Suggest
`ni-end` only when the readiness gate reports `READY` or
`READY_WITH_DEFERRALS`. Never declare completion by model judgment alone.

## Output style

When responding during planning:

- Lead with the current planning summary or what changed.
- Name the files changed when an answer updates planning state, including
  `.ni/session.json` when refreshed.
- Name affected IDs and whether they are accepted, draft, assumption, rejected,
  deferred, resolved, or blockers.
- Ask only the next focused questions needed to unblock readiness.
- Prefer deterministic next questions from `ni status --next-questions`.
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
2. For OQ-003, what decision resolves this blocker, should it be deferred, or
   why must it remain blocking?

After you answer, I will update docs/plan/** and .ni/contract.json, then run
ni status --dir . --next-questions
```

## Do not

- Do not create `.ni/plan.lock.json`.
- Do not run implementation work.
- Do not create a shell or Codex adapter.
- Do not write generated harness files.
- Do not hide blocker questions.
- Do not weaken evaluations to make the plan appear ready.
- Do not convert ambiguous user statements into accepted decisions.
- Do not silently delete planning records.
- Do not edit locked planning docs except through the amendment or relock flow.
- Do not edit files outside `docs/plan/**`, `.ni/contract.json`, and
  `.ni/session.json` unless the user explicitly asks for a different NI
  maintenance task.
- Do not add user-facing contract `add`, `list`, or `set` commands.
- Do not create SPEC runner behavior.
- Do not directly call downstream runtimes.
