# Status proof

`ni status --proof` prints a deterministic readiness proof from the same
pre-runtime rules that decide `BLOCKED`, `READY_WITH_DEFERRALS`, or `READY`.
It is an explanation of the gate result, not a model-generated review.

The proof does not edit planning docs, does not resolve issues, and does not
start execution. If the proof reports `BLOCKED`, execution must not start.

## Commands

```bash
ni status --dir . --proof
ni status --dir . --proof --next-questions
ni status --dir . --proof --json
ni status --dir . --proof --json --next-questions
```

`--next-questions` uses the same readiness issues as the proof and appends a
deterministic interview section. JSON output includes `proof` only when
`--proof` is present and includes `next_questions` only when
`--next-questions` is present.

## Human output

```text
NI Intent Readiness: BLOCKED

Proof:
- CAP-003 is accepted but has no linked EVAL.
- RISK-002 is high severity but has no mitigation.
- OQ-001 is marked as blocker.
- DEC-004 conflicts with DEC-002.

Execution must not start.

Next questions:
1. For CAP-003, what evidence proves this capability works, or should that evidence be deferred?
2. For RISK-002, what mitigation, owner, or explicit accepted-risk decision is required?
3. For OQ-001, what decision resolves this blocker, should it be deferred, or why must it remain blocking?
```

## Rule sources

Proof items are derived from readiness, sync, and conflict rules:

```text
R001 required planning docs exist
R002 .ni/contract.json is valid JSON
R003 at least one capability exists
R004 every accepted capability has at least one linked evaluation
R005 every evaluation has a method
R006 every high-severity risk has mitigation
R007 every accepted capability has at least one artifact or requirement
R008 decisions use accepted, deferred, rejected, or not_applicable
R009 blocker open questions prevent lock
R010 at least one non-goal exists
R011 readiness profile definitions are valid
R012 planning docs and contract are synchronized
R013 accepted decisions do not contradict each other
D001 deferred decisions remain explicit
D002 non-blocking open questions remain explicit
```

`READY` proof means no blocker or deferral issue was produced for the active
profile. `READY_WITH_DEFERRALS` proof lists deferral items that remain explicit
but do not block the active profile. `BLOCKED` proof lists blocker evidence and
may also list deferrals when both are present.

## Determinism

The proof is deterministic because it is compiled from contract fields,
required planning docs, docs/contract sync findings, readiness profiles, and
accepted-decision contradiction checks. It does not call external APIs or an
LLM, and it does not infer unstated intent.
