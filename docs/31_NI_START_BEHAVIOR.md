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

Questions should be focused on the gaps that block readiness. Ask at most one
to three focused questions per turn and prefer questions returned by the CLI.
Avoid broad generic brainstorming unless the project is still empty. Instead
of asking "What else should the plan include?", ask a concrete question such
as:

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
planning open and asks the highest-impact one to three questions from
`next_questions`. If the CLI returns no next questions, show the readiness
issues directly. If it reports `READY` or `READY_WITH_DEFERRALS`, `ni-start` may
suggest moving to `ni-end`.

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
