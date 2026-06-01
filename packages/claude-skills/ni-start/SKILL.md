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

Skills are UX; CLI is authority.

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

## Language-Adaptive Questions

Ask user-facing planning questions in the language of the user's latest
substantive message. If the user explicitly requests a language, use that
language. If the conversation is mixed, prefer the latest explicit language
preference.

Preserve IDs, command names, file paths, schema keys, target names, and status
constants exactly. Do not translate tokens such as `R014`, `OQ-001`,
`SYNC-014`, `ni status`, `ni end`, `ni run`, `.ni/contract.json`, `READY`,
`BLOCKED`, or `READY_WITH_DEFERRALS`.

CLI output may remain English. You may explain or summarize CLI proof and
next-question output in the user's language, but do not alter meaning and do
not make model-translated text authoritative over the CLI. Do not use
localization to weaken readiness gates.

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
7. Ask focused questions from the grouped CLI `next_questions` result when it
   is present. Read the groups in CLI order, select the highest-priority group,
   preserve the group label, and ask at most one group per turn unless the
   group is the compact `First-run card`. Ask at most three primary questions
   at once. Preserve the referenced IDs, locations, readiness gap, and allowed
   answer shapes such as evidence, decision, deferral, not_applicable,
   mitigation, or explicit non-goal. Do not invent broad generic brainstorming
   questions while deterministic next questions exist.

## Grouped Next-Question Handling

`ni status --dir . --proof --next-questions` is the conversation driver when it
returns grouped next questions. Preserve group labels such as:

- `First-run card`
- `Sync repairs`
- `Risk decisions`
- `Evaluation evidence`
- `Scope boundaries`
- `Open blockers`

Use the groups this way:

- `First-run card`: use the compact three-question card. Do not add unrelated
  lower-priority questions.
- `Sync repairs`: ask repair questions and ask whether to update contract,
  revise docs, revise both, or keep the blocker with reason. Do not restart
  broad planning if one side already has useful content. Do not re-ask
  `R014`, `R015`, or `R016` when `SYNC-014`, `SYNC-015`, or `SYNC-016`
  already provide matching repair questions.
- `Risk decisions`: ask for mitigation, owner, monitoring plan, accepted-risk
  decision, or explicit deferral. Do not suggest lowering risk severity to
  pass readiness.
- `Evaluation evidence`: ask what evidence proves a capability is complete.
  Offer the CLI answer shapes: test, review checklist, demo condition, user
  approval, protocol check, or manual inspection.
- `Scope boundaries`: ask for explicit non-goals and explain that non-goals
  prevent scope drift.
- `Open blockers`: ask whether to resolve, defer with reason, or keep
  blocking. Do not silently convert open questions into accepted decisions.

After the user answers the selected group, update `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json`, then run or request
`ni status --dir . --proof --next-questions` again. Explain that readiness is
blocked or cleared by deterministic CLI gates, not model judgment.

## Planning Proof Capture

After every meaningful authoring update, show a concise planning proof block.
The block must describe what changed in planning state without exposing hidden
chain-of-thought and without making you the readiness authority.

Use this shape:

```text
Planning proof:
- User input captured:
  "<short paraphrase of user answer>"
- Interpreted planning records:
  - Purpose: ...
  - Actors/outcomes: ...
  - Delivery surface: ...
  - Capabilities: CAP-001 ...
  - Requirements: REQ-001 ...
  - Risks: RISK-001 ...
  - Evaluations: EVAL-001 ...
  - Decisions: DEC-001 accepted/deferred/rejected if applicable
  - Assumptions: ASM-001 or open question if applicable
  - Non-goals: NG-001 if applicable
  - Open questions: OQ-001 ...
- Updated planning artifacts:
  - docs/plan/00_project_brief.md: purpose clarified
  - docs/plan/01_actors_outcomes.md: primary actors added
  - docs/plan/03_interaction_contract.md: delivery surface recorded
  - .ni/contract.json: project.purpose, actors/outcomes, delivery_surfaces updated
  - .ni/session.json: active focus and pending questions updated
- Status result:
  - before: BLOCKED because R014/R015/R016
  - after: BLOCKED/READY_WITH_DEFERRALS/READY because ...
- Remaining blockers:
  - OQ-001 ...
- Next question group:
  - Sync repairs / Risk decisions / Evaluation evidence / Open blockers / none
```

Rules:

- Keep it short.
- Do not invent file changes, contract fields, or IDs.
- If no files were changed, say no planning artifacts were updated.
- Do not claim readiness unless CLI status proves it.
- Do not claim lock unless `ni end` actually created `.ni/plan.lock.json`.
- Keep uncertain user statements as assumptions or open questions.
- Record clear exclusions as non-goals.
- If docs and contract disagree, say so and run or request status again.
- In no-terminal mode, call this a draft audit trail only; it becomes trusted
  only after CLI validation.

## First-run Opening Card

When a fresh project reports the first-run blockers `R014`, `R015`, and
`R016`, use them as the opening planning card. Do not ask broad generic
brainstorming questions and do not ask more than three questions at once. If
the CLI also reports a template blocker such as `OQ-001`, keep it visible but
do not let it displace the three foundational questions.

Use this framing:

```text
ni is blocked because the initial project intent is not explicit enough to lock
yet. I need three things before execution can safely start: what reality this
project should change, who it is for, and how it will be delivered.

Implementation has not started. This is still planning.
```

If the user's latest substantive message is Korean, use Korean human-facing
wording for the framing and questions while preserving `ni`, `BLOCKED`, IDs,
commands, file paths, schema keys, target names, and status constants.

Then ask:

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?

After the user answers, record clear purpose in `docs/plan/00_project_brief.md`
and `project.purpose`, actors and outcomes in
`docs/plan/01_actors_outcomes.md` and matching contract records, and delivery
surface in `docs/plan/08_delivery_operation.md`, `product_type`, and
`delivery_surfaces` when clear. Record uncertain answers as assumptions or
open questions. Record clear exclusions as non-goals. Do not convert vague
answers into accepted decisions without confirmation.

Run or request `ni status --dir . --proof --next-questions` again after the
update, and use the CLI result as the next authority.

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
8. Report the planning proof block, including changed files, affected IDs,
   readiness blockers, and the next focused planning questions.
9. Suggest `ni-end` only when the CLI reports no blockers. Never declare
   completion by model judgment alone.

## Output Style

When responding during planning:

- Lead with the current planning summary or what changed.
- Name changed files, including `.ni/session.json` when refreshed.
- Include the planning proof block after meaningful authoring updates.
- Name affected IDs and whether they are accepted, draft, assumption,
  rejected, deferred, resolved, or blockers.
- Ask only the next focused questions needed to unblock readiness.
- Ask at most one grouped next-question section per turn, except for the
  compact `First-run card`.
- Ask at most three primary questions at once.
- Use deterministic next questions from `ni status --proof --next-questions`
  directly.
- Preserve the CLI group labels, including `First-run card`, `Sync repairs`,
  `Risk decisions`, `Evaluation evidence`, `Scope boundaries`, and
  `Open blockers`.
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
