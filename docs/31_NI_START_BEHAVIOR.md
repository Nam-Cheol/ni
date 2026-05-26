# ni-start behavior

`ni-start` is the sustained conversation mode for NI planning. After
`ni init`, it is the intended authoring interface for users who want a model to
maintain planning state without manually editing `.ni/contract.json`.

The skill does not replace the CLI gates. It writes planning state; the CLI
validates, locks, and compiles.

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
summary names the active planning focus, purpose, accepted capabilities,
delivery surface, decisions, non-goals, open questions, and known readiness
blockers. Claims from `.ni/session.json` should be checked against the contract,
docs, and `ni status`.

## Gap detection

`ni-start` should identify missing required planning areas from the current
contract and docs, with `ni status --next-questions` as the first source for
readiness-blocking interview prompts. Common gaps include:

- purpose, actors, outcomes, or delivery surface still marked TODO,
- capabilities without requirements or evaluations,
- missing artifacts for accepted capabilities,
- high-severity risks without mitigation,
- blocker questions that affect scope or acceptance criteria,
- constraints or non-goals that conflict with requested behavior,
- docs and `.ni/contract.json` disagreeing about the same record.

Questions should be focused on the gaps that block readiness. Prefer questions
returned by the CLI. Instead of asking "What else should the plan include?",
ask a concrete question such as:

```text
For CAP-002, what evidence proves this capability works, or should that evidence be deferred?
```

Questions must preserve the relevant IDs, avoid implying implementation work,
avoid pressuring acceptance, and allow `deferred` or `not_applicable` where
those are appropriate planning outcomes.

## Persistence Rules

After the user answers, `ni-start` updates both planning forms in the same
authoring pass and refreshes session state:

- `docs/plan/**` explains the plan for humans,
- `.ni/contract.json` stores stable IDs, statuses, and trace links for the CLI.
- `.ni/session.json` stores the latest focus, short summary, pending questions,
  recent decisions, recent risks, readiness status, readiness blockers, and docs
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
ni status --dir . --next-questions
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
