# ni-start behavior

`ni-start` is the sustained conversation mode for NI planning. After
`ni init`, it is the intended authoring interface for users who want a model to
maintain planning state without manually editing `.ni/contract.json`.

The skill does not replace the CLI gates. It writes planning state; the CLI
validates, locks, and compiles.

It also supports resume mode for long-running planning. Resume mode reconstructs
or verifies planning continuity from persisted project files, not hidden chat
memory.

```text
ni init -> ni-start conversation -> docs/plan + .ni/contract.json -> ni status
```

## Responsibilities

At the start of a planning turn, `ni-start` reads the current planning source:

- `AGENTS.md` for repository authority rules,
- `docs/plan/**` for human-readable planning state,
- `.ni/contract.json` for machine-readable planning state,
- `.ni/session.json` for non-authoritative carryover context,
- `.ni/plan.lock.json` if present, to detect locked-plan authority.

It then summarizes the current state before asking for more input. A useful
start summary must name:

1. current purpose,
2. active readiness profile,
3. product type and delivery surfaces,
4. accepted capabilities,
5. unresolved blocker questions,
6. recent decisions,
7. next recommended planning focus.

Claims from `.ni/session.json` should be checked against the contract, docs,
and `ni status --dir . --proof --next-questions`.

## First-run conversation card

On a fresh workspace after `ni init`, `ni status --proof --next-questions`
commonly reports the first intent blockers:

- `R014 Project purpose is missing`
- `R015 Actors or outcomes are missing`
- `R016 Delivery surface is missing`

`ni-start` should turn those blockers into a concise opening planning card
instead of making the user feel stuck. The card should explain that `ni` is
blocked only because the initial intent is not explicit enough to lock, and
that implementation has not started.

Recommended wording:

```text
ni is blocked because the initial project intent is not explicit enough to lock
yet. I need three things before execution can safely start: what reality this
project should change, who it is for, and how it will be delivered.

Implementation has not started. This is still planning.
```

Then ask at most three focused questions, grouped around the first-run
blockers:

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?

The model may mention that the CLI also surfaced a template open question, but
it should not let generic brainstorming displace the three first-run questions.
When the user answers, `ni-start` records purpose in
`docs/plan/00_project_brief.md` and `project.purpose`, actors and outcomes in
`docs/plan/01_actors_outcomes.md` and matching contract records, and delivery
surface in `docs/plan/08_delivery_operation.md`, `product_type`, and
`delivery_surfaces` as appropriate.

Clear exclusions should be recorded as non-goals. Uncertain, tentative, or
vague answers should stay visible as assumptions or blocker open questions
until the user confirms them. The model must not convert vague answers into
accepted decisions merely to pass readiness. After recording the answer,
`ni-start` runs or requests `ni status --dir . --proof --next-questions`
again and uses the CLI result as the next authority.

If that status result reports `SYNC-014`, `SYNC-015`, or `SYNC-016`, the
first-run answer was recorded inconsistently. `ni-start` should repair the
stale side of the docs/contract pair, or keep uncertain intent as an explicit
assumption, open blocker question, or deferral. It must not proceed to
`ni-end` while these sync diagnostics block readiness.

`ni status --next-questions` groups these prompts under headings such as
`First-run card`, `Sync repairs`, `Risk decisions`, `Evaluation evidence`,
`Scope boundaries`, and `Open blockers`. `ni-start` must preserve the group,
ID, location, and answer-shape fields when asking the user, and ask only the
top questions returned by the CLI.

## Grouped next-question handling

`ni-start` should run or request this command at the start of a planning turn
and after every meaningful authoring update:

```bash
ni status --dir . --proof --next-questions
```

The grouped `Next questions` section is the primary conversation driver when
it is present. `ni-start` should read the groups in CLI order, select the
highest-priority group, and ask at most one group per turn unless the group is
the compact `First-run card`. It should ask at most three primary questions at
once. It must not invent broad brainstorming questions while deterministic
next questions exist.

Group labels should be shown or preserved in the model response. If the CLI
prints `Location` or `Answer shape`, keep those fields visible enough that the
user can answer in the expected form. Readiness remains blocked by
deterministic CLI gates, not by model judgment.

Group-specific rules:

- `First-run card`: use the compact three-question card for purpose,
  actors/outcomes, and delivery surface. Do not add unrelated lower-priority
  questions.
- `Sync repairs`: ask the repair questions and ask whether to update contract,
  revise docs, revise both, or keep the blocker with reason. Do not restart
  broad planning if one side already has useful content. Do not re-ask
  `R014`, `R015`, or `R016` when matching `SYNC-014`, `SYNC-015`, or
  `SYNC-016` repair questions are present.
- `Risk decisions`: ask for mitigation, owner, monitoring plan, accepted-risk
  decision, or explicit deferral. Do not suggest lowering risk severity to
  pass readiness.
- `Evaluation evidence`: ask what evidence proves a capability is complete.
  Offer answer shapes such as test, review checklist, demo condition, user
  approval, protocol check, or manual inspection.
- `Scope boundaries`: ask for explicit non-goals and explain that non-goals
  prevent scope drift.
- `Open blockers`: ask whether to resolve, defer with reason, or keep
  blocking. Do not silently convert open questions into accepted decisions.

After the user answers the selected group, `ni-start` updates `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json` together, then runs or requests
`ni status --dir . --proof --next-questions` again.

## Planning proof capture

After every meaningful authoring update, `ni-start` should report a concise
planning proof block. The block makes the edit auditable for the user without
turning the model into a readiness authority.

Required shape:

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

- Keep it concise.
- Do not expose hidden chain-of-thought.
- Do not invent file changes or contract IDs that were not updated.
- If no files changed, say `No planning artifacts were updated.`
- Do not claim readiness unless the after-status came from the CLI.
- Do not claim lock unless `ni end` actually created `.ni/plan.lock.json`.
- Keep uncertain user statements as assumptions or open questions.
- Record clear exclusions as non-goals.
- If docs and contract disagree, say so and use the next status result or
  request one before proceeding.

In no-terminal mode, this block is only a draft audit trail. It becomes trusted
only after the drafted docs and contract pass CLI validation.

## Resume mode

When a later model session resumes planning, `ni-start` first uses the same
authoritative inputs as a normal turn. If `.ni/session.json` exists, it may use
the file for active focus, last summary, pending questions, recent decisions and
risks, last readiness status, and recently updated docs.

Every session claim must be verified against `.ni/contract.json`,
`docs/plan/**`, lock state, and `ni status --dir . --proof --next-questions`
when available. If the session file conflicts with a contract record, the
contract wins and the conflict is reported. If it conflicts with a locked plan,
the lock and locked docs win. A lock hash mismatch stops the turn with
`BLOCKED`.

If `.ni/session.json` is missing, invalid, empty, or stale, `ni-start`
reconstructs a resumed summary from `docs/plan/**`, `.ni/contract.json`, and
CLI readiness output. The next meaningful planning edit should recreate or
refresh `.ni/session.json` with bounded continuity state. It should not store a
full raw transcript by default.

## Gap detection

`ni-start` should identify missing required planning areas from the current
contract and docs, with `ni status --proof --next-questions` as the first
source for readiness-blocking interview prompts. Common gaps include:

- purpose, actors, outcomes, or delivery surface still marked TODO,
- capabilities without requirements or evaluations,
- missing artifacts for accepted capabilities,
- high-severity risks without mitigation,
- blocker questions that affect scope or acceptance criteria,
- constraints or non-goals that conflict with requested behavior,
- docs and `.ni/contract.json` disagreeing about the same record.

Questions should be focused on the gaps that block readiness. When the CLI
returns grouped next questions, ask the highest-priority group first, ask at
most one group per turn, and ask at most three primary questions at once.
Avoid broad generic brainstorming while deterministic next questions exist.
Instead of asking "What else should the plan include?", ask a concrete
question such as:

```text
CAP-002 has no evaluation. What evidence would prove this capability is complete: a test, review checklist, demo condition, user approval, or an explicit deferral?
```

Questions must preserve the relevant IDs, avoid implying implementation work,
avoid pressuring acceptance, and preserve the CLI answer shapes. Valid answer
shapes may include evidence, an accepted decision, an explicit deferral,
`not_applicable`, a mitigation, or an explicit non-goal.

## Persistence Rules

After the user answers, `ni-start` updates both planning forms in the same
authoring pass and refreshes session state:

- `docs/plan/**` explains the plan for humans,
- `.ni/contract.json` stores stable IDs, statuses, and trace links for the CLI.
- `.ni/session.json` stores the latest focus, short summary, active readiness
  profile, product type and delivery surfaces, pending questions, recent
  decisions, recent risks, readiness status, readiness blockers, and docs
  changed in the turn.

The skill records purpose, actors, capabilities, requirements, decisions,
risks, evaluations, non-goals, constraints, artifacts, assumptions, and open
questions when the conversation changes them. Tentative or inferred statements
stay visible as assumptions or open questions until the user confirms them.
Session state is only a planning aid. It must not override locked docs, must not
mark docs complete, and must not store full raw chat logs by default.

Existing IDs remain stable. New IDs are appended only when a distinct record is
needed. Obsolete records should be marked rejected, deferred, or not applicable
when preserving history matters.

## Readiness Handoff

After meaningful updates, `ni-start` runs or requests:

```bash
ni status --dir . --proof --next-questions
```

The status result is authoritative. If it reports `BLOCKED`, `ni-start` keeps
planning open and asks the highest-priority group from `next_questions`. If
the CLI returns no next questions, show the readiness issues directly. If it
reports `READY` or `READY_WITH_DEFERRALS`, `ni-start` may suggest moving to
`ni-end`.

`ni-start` must never declare completion by model judgment alone.

## Boundaries

`ni-start` must not:

- add user-facing contract `add`, `list`, or `set` commands,
- run implementation work,
- create SPEC runner behavior,
- create shell, Codex, queue, evidence-runner, or downstream runtime adapters,
- directly call downstream runtimes,
- create or edit `.ni/plan.lock.json`,
- weaken accepted requirements or evaluations to make validation pass.

The skill may propose downstream seed ideas only as planning content. Downstream
harness material remains derived and mutable, not kernel-owned execution state.
