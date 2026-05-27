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
`ni status --dir . --next-questions` before describing the authoritative
readiness state.

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
4. Summarize the current purpose, actors, accepted capabilities, requirements,
   evaluations, risks, decisions, non-goals, artifacts, constraints, and open
   questions.
5. Run `ni status --dir . --next-questions` when command execution is
   available. If not, ask the user to run it and paste the exact output.

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
5. Refresh `.ni/session.json` only as bounded continuity state. It must not
   override the contract, docs, or lock.
6. Preserve trace links from capabilities to requirements, evaluations, risks,
   and artifacts.
7. Run or request `ni status --dir . --next-questions` after meaningful edits.
8. Report changed files, affected IDs, readiness blockers, and the next
   focused planning questions.

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
