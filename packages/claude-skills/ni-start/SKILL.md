---
name: ni-start
description: Continue NI planning in Claude-compatible workspaces while keeping the NI CLI as readiness authority.
---

# ni-start

Use this skill when the user asks Claude to start, continue, resume, or update
NI planning for a project.

`ni-start` is conversation-driven authoring. It helps the model and user turn
intent into synchronized planning docs and contract records before any
implementation begins.

It also supports resume mode for long-running planning. A later model session
must resume from persisted docs, `.ni/contract.json`, and bounded
`.ni/session.json` state instead of hidden chat memory.

## Authority Boundary

Skills are UX; the `ni` CLI is authority.

You may draft, review, and update:

```text
docs/plan/**
.ni/contract.json
.ni/session.json
```

You may read supporting project files such as `AGENTS.md`,
`.ni/project.json`, `.ni/readiness.rules.json`, `.ni/readiness.profiles.json`,
and `.ni/plan.lock.json`.

Do not declare readiness from model judgment. Run or request
`ni status --dir . --proof --next-questions` before describing the
authoritative readiness state.

Do not create `.ni/plan.lock.json`. Do not edit `.ni/plan.lock.json` manually.
If a lock hash mismatch exists, stop and report `BLOCKED`.

## Start Of Turn

1. Locate the project root from the user's workspace context. Do not assume a
   global install path or global `ni` path unless the user has documented and
   verified it.
2. Read `AGENTS.md`, `.ni/contract.json`, relevant `docs/plan/**`, and
   `.ni/session.json` if it exists.
3. If `.ni/plan.lock.json` exists, treat the plan as locked. Do not edit
   locked planning docs unless the user explicitly starts an amendment or
   relock flow.
4. If `.ni/session.json` exists, verify important session claims against
   `.ni/contract.json`, `docs/plan/**`, lock state, and the CLI result when
   available. If session state conflicts with contract records, the contract
   wins. If it conflicts with a valid lock, the lock and locked docs win.
5. Run or request `ni status --dir . --proof --next-questions` when command
   execution is available. If not, ask the user to run it and paste the exact
   output.
6. Summarize the current planning state before asking for more input. The start
   summary must include:
   - current purpose,
   - active readiness profile,
   - product type and delivery surfaces,
   - accepted capabilities,
   - unresolved blocker questions,
   - recent decisions,
   - next recommended planning focus.
7. Ask focused questions about the highest-impact readiness gaps. Ask at most
   one to three questions per turn. Prefer questions from the CLI
   `next_questions` result, preserving the referenced IDs, readiness gap, and
   allowed answer shapes such as evidence, decision, deferral, not_applicable,
   mitigation, or explicit non-goal. Avoid broad generic brainstorming unless
   the project is still empty.

## Authoring Loop

When the user provides new intent:

1. Extract planning state from the conversation: purpose, actors,
   capabilities, requirements, decisions, risks, evaluations, non-goals,
   constraints, artifacts, assumptions, and open questions.
2. Keep tentative, inferred, conflicting, or incomplete statements as draft
   records, assumptions, or open questions.
3. Preserve stable IDs where possible.
4. Update `docs/plan/**` and `.ni/contract.json` together when the turn changes
   planning state.
5. Refresh `.ni/session.json` only as bounded continuity state. Include active
   planning focus, last planning summary, active readiness profile, product
   type and delivery surfaces, pending questions, recent decisions, recent
   risks, last readiness status, last readiness blockers, last updated docs,
   and a note that raw transcript is not the source of truth. It must not
   override the contract, docs, or lock.
6. Preserve trace links from capabilities to requirements, evaluations, risks,
   and artifacts.
7. Run or request `ni status --dir . --proof --next-questions` after
   meaningful edits.
8. Report changed files, affected IDs, readiness blockers, and the next
   focused planning questions.
9. Suggest `ni-end` only when the CLI reports no blockers. Never declare
   completion by model judgment alone.

## Output Style

When responding during planning:

- Lead with the current planning summary or what changed.
- Name changed files, including `.ni/session.json` when refreshed.
- Name affected IDs and whether they are accepted, draft, assumption,
  rejected, deferred, resolved, or blockers.
- Ask only the next focused questions needed to unblock readiness.
- Ask at most one to three questions per turn.
- Prefer deterministic next questions from `ni status --proof --next-questions`.
- Preserve the CLI answer shapes: evidence, decision, deferral,
  not_applicable, mitigation, or explicit non-goal.
- If `ni status` reports `BLOCKED`, state the blockers plainly and keep
  planning open.

## Boundaries

Do not execute implementation.

Do not call Claude APIs.

Do not add shell adapters, Codex adapters, evidence runners, queues, PR
automation, release automation, plugin systems, model orchestration, TUI, or
web UI behavior.

Do not weaken risks, mitigations, requirements, evaluations, or non-goals to
make validation pass.

Do not silently delete planning records. Mark them rejected, deferred,
resolved, or not applicable when history matters.

Do not edit files outside `docs/plan/**`, `.ni/contract.json`, and
`.ni/session.json` during normal planning authoring unless the user explicitly
asks for a separate NI maintenance task.

Do not add user-facing contract `add`, `list`, or `set` commands.

Do not directly call downstream runtimes.
