---
name: ni-start
description: Continue NI planning by updating docs/plan and .ni/contract.json while preserving the CLI as readiness authority.
---

# ni-start

Use this skill when the user says `ni-start` or asks to continue planning a project with NI.

`ni-start` is the primary authoring UX after `ni init`. It turns a sustained
model-user planning conversation into synchronized updates to `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json`.

It also supports resume mode for long-running planning. A later model session
must resume from persisted docs, `.ni/contract.json`, and bounded
`.ni/session.json` state rather than hidden chat memory.

## Authority boundary

Skills are UX; the CLI is authority.

Skills are UX; CLI is authority.

You are not the authority for readiness or lock state. The `namba-intent` CLI is the authority.

Do not say the plan is complete unless `namba-intent status` has no blockers. If `.ni/plan.lock.json` exists, do not silently edit locked planning docs; first state that planning is locked and proceed only through the amendment or relock flow when the user explicitly resumes planning.

If your interpretation conflicts with `namba-intent status`, report the CLI result and stop. Do not override, reinterpret, or soften a CLI `BLOCKED` result.

`LOCK-STALE` means the existing lock no longer matches current planning inputs.
- Skills may help draft amended planning text.
- Skills may help explain `LOCK-STALE`.
- Skills do not determine readiness.
- Skills do not lock or relock.
- Skills do not replace `namba-intent status`, `namba-intent end`, or `namba-intent run`.
- Skills do not update `.ni/plan.lock.json`.

Recovery order: `review changed intent -> namba-intent status --proof --next-questions -> namba-intent end -> namba-intent run --max-chars 4000`.

## Language-adaptive questions

Ask user-facing planning questions in the language of the user's latest
substantive message. If the user explicitly requests a language, use that
language. If the conversation is mixed, prefer the latest explicit language
preference.

Preserve IDs, command names, file paths, schema keys, target names, and status
constants exactly. Do not translate tokens such as `R014`, `OQ-001`,
`SYNC-014`, `namba-intent status`, `namba-intent end`, `namba-intent run`, `.ni/contract.json`, `READY`,
`BLOCKED`, or `READY_WITH_DEFERRALS`.

CLI output may remain English. You may explain or summarize CLI proof and
next-question output in the user's language, but do not alter meaning and do
not make model-translated text authoritative over the CLI. Do not use
localization to weaken readiness gates.

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

## Resume mode

Use resume mode whenever the user asks to continue, resume, pick up, or
re-enter planning after an earlier model session.

Resume mode starts from persisted files:

- If `.ni/session.json` exists, read it as a planning aid and verify every
  important claim against `.ni/contract.json`, `docs/plan/**`, lock state, and
  `namba-intent status --dir . --proof --next-questions` when available.
- If `.ni/session.json` is missing, invalid, empty, or too stale to trust,
  reconstruct the planning summary from `.ni/contract.json`, `docs/plan/**`,
  lock state, and CLI readiness output.
- Do not depend on hidden chat memory or a raw transcript.
- Do not store full raw transcript content by default.

When session state conflicts with contract records, the contract wins. Report
the conflict using the relevant IDs, ignore or correct the stale session entry
in the next planning update, and continue from the contract value. When session
state conflicts with locked docs or a valid lock, the lock and locked docs win.
If a lock hash mismatch exists, stop and report `BLOCKED`.

The resumed summary should name:

- the active focus from `.ni/session.json`, or that focus was reconstructed,
- which session claims were confirmed against docs and contract,
- which session claims were stale or conflicted,
- the active readiness profile,
- product type and delivery surfaces,
- accepted and draft capabilities,
- decisions, non-goals, risks, and assumptions that affect readiness,
- open blocker questions and the CLI readiness status.

## Start-of-turn process

1. Read `AGENTS.md`, `.ni/contract.json`, `docs/plan/**`, `.ni/session.json`
   if it exists, and `.ni/plan.lock.json` if it exists.
2. If this is a resumed session, compare `.ni/session.json` to the contract,
   docs, and lock state. Report conflicts, use the contract or locked docs as
   authority, and reconstruct from docs and contract if session state is
   missing or unusable.
3. Run or request `namba-intent status --dir . --proof --next-questions` when available.
   The CLI output is the source for readiness status, active profile, proof,
   blockers, and deterministic next questions. If command execution is not
   available, ask the user to run it and paste the exact output.
4. Summarize current planning state in a few concrete bullets before asking for
   more input. The start summary must include:
   - current purpose,
   - active readiness profile,
   - product type and delivery surfaces,
   - accepted capabilities,
   - unresolved blocker questions,
   - recent decisions,
   - next recommended planning focus.
   You may also mention whether focus came from `.ni/session.json` or was
   reconstructed, and any session conflicts discovered during resume.
5. Identify missing required planning areas from the current docs, contract,
   and readiness profile. Check for:
   - missing or TODO purpose, actors, capabilities, requirements, risks,
     evaluations, artifacts, constraints, delivery expectations, non-goals, or
     open questions,
   - accepted capabilities without linked requirements or evaluations,
   - high-severity risks without mitigation,
   - blocker questions that still affect acceptance criteria or scope.
6. Ask focused questions from the grouped CLI `next_questions` result when it
   is present. Read the groups in CLI order, select the highest-priority group,
   preserve the group label, and ask at most one group per turn unless the
   group is the compact `First-run card`. Ask at most three primary questions
   at once. You may lightly rephrase for clarity, but preserve the referenced
   IDs, locations, readiness gap, and allowed outcomes. Preserve concrete
   answer shapes such as evidence, decision, deferral, not_applicable,
   mitigation, or explicit non-goal. Do not invent broad generic brainstorming
   questions while deterministic next questions exist.

## Grouped next-question handling

`namba-intent status --dir . --proof --next-questions` is the conversation driver when it
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
`namba-intent status --dir . --proof --next-questions` again. Explain that readiness is
blocked or cleared by deterministic CLI gates, not model judgment.

## Planning proof capture

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
- Do not claim lock unless `namba-intent end` actually created `.ni/plan.lock.json`.
- Skills may help draft or explain proof-related planning text.
- Skills do not determine readiness.
- Skills do not lock plans.
- Skills do not replace `namba-intent status`, `namba-intent end`, or `namba-intent run`.
- Keep uncertain user statements as assumptions or open questions.
- Record clear exclusions as non-goals.
- If docs and contract disagree, say so and run or request status again.
- In no-terminal mode, call this a draft audit trail only; it becomes trusted
  only after CLI validation.

## First-run opening card

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

Run or request `namba-intent status --dir . --proof --next-questions` again after the
update, and use the CLI result as the next authority.

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
   - active readiness profile,
   - product type and delivery surfaces,
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
11. Run or report `namba-intent status --dir . --proof --next-questions` at the end when available.
12. Show the planning proof block with changed files and affected IDs.
13. Show readiness gaps and next questions from the CLI result.
14. Reflect the CLI readiness status and blockers back into `.ni/session.json`
    without treating the session file as readiness authority.

Continue this loop until `namba-intent status` reports no blocking issues. Suggest
`ni-end` only when the readiness gate reports `READY` or
`READY_WITH_DEFERRALS`. Never declare completion by model judgment alone.

## Output style

When responding during planning:

- Lead with the current planning summary or what changed.
- Name the files changed when an answer updates planning state, including
  `.ni/session.json` when refreshed.
- Include the planning proof block after meaningful authoring updates.
- Name affected IDs and whether they are accepted, draft, assumption, rejected,
  deferred, resolved, or blockers.
- Ask only the next focused questions needed to unblock readiness.
- Ask at most one grouped next-question section per turn, except for the
  compact `First-run card`.
- Ask at most three primary questions at once.
- Avoid broad generic brainstorming while deterministic next questions exist.
- Use deterministic next questions from `namba-intent status --proof --next-questions`
  directly.
- Preserve the CLI group labels, including `First-run card`, `Sync repairs`,
  `Risk decisions`, `Evaluation evidence`, `Scope boundaries`, and
  `Open blockers`.
- Preserve the CLI answer shapes: evidence, decision, deferral,
  not_applicable, mitigation, or explicit non-goal.
- If `namba-intent status` reports `BLOCKED`, state the blockers plainly and keep
  planning open.
- If a lock hash mismatch exists, stop and report `BLOCKED`.

## Example turn shape

```text
Current state:
- Purpose: ...
- Accepted capabilities: CAP-001, CAP-002
- Blocking gaps: OQ-003 needs an owner; CAP-002 has no evaluation

Focused group from namba-intent status --proof --next-questions:
Evaluation evidence

1. CAP-002 has no evaluation. What evidence would prove this capability is
   complete?
   Answer shape: test, review checklist, demo condition, user approval,
   protocol check, or manual inspection

After you answer, I will update docs/plan/**, .ni/contract.json, and
.ni/session.json, then run namba-intent status --dir . --proof --next-questions
```

## Do not

- Do not create `.ni/plan.lock.json`.
- Do not edit `.ni/plan.lock.json` manually.
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
