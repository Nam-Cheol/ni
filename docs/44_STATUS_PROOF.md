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

Human proof output is grouped so people and models can act on the result
without reinterpreting readiness rules:

- `Blockers` are readiness failures that prevent lock or execution.
- `Deferrals` are unresolved, non-blocking intent records that keep the status
  at `READY_WITH_DEFERRALS`.
- `Warnings` are non-blocking cautions, such as deferred decisions that
  downstream work must not depend on.
- `Passed checks` names deterministic checks that did not produce issues.

Every blocker includes fixed `Why it matters` and `Next` text. The CLI does
not auto-fix files or infer the answer; the next step text only names the
planning action needed from the user/model authoring loop.

```text
NI Intent Readiness: BLOCKED

Blockers:
- CAP-003 is accepted but has no linked EVAL.
  Why it matters: ni cannot prove this capability is verifiable.
  Next: define evidence and link an evaluation.
- RISK-002 is high severity but has no mitigation.
  Why it matters: high-severity risks can invalidate downstream work unless mitigation is explicit.
  Next: add mitigation, an owner, or an explicit accepted-risk decision.
- OQ-001 is marked as blocker.
  Why it matters: open blocker questions mean required intent is still unresolved.
  Next: answer or defer the blocker question, or keep it blocking with an explicit reason.
- DEC-004 conflicts with DEC-002.
  Why it matters: conflicting accepted decisions give downstream actors incompatible instructions.
  Next: revise, reject, or split one conflicting accepted decision.

Deferrals:
- None.

Warnings:
- None.

Passed checks:
- Required docs exist.
- Contract JSON is valid.
- Readiness profile definitions are valid.

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
but do not block the active profile; deferred decisions are displayed as
warnings because downstream work must avoid depending on them. `BLOCKED` proof
lists blocker evidence and may also list deferrals or warnings when they are
present.

JSON output remains a stable machine-readable readiness result. With
`--proof --json`, the `proof` field stays an array of deterministic proof
items with rule id, severity, references when present, and message. Human-only
group labels, why text, and next-step text are not added to the JSON schema.

## Determinism

The proof is deterministic because it is compiled from contract fields,
required planning docs, docs/contract sync findings, readiness profiles, and
accepted-decision contradiction checks. It does not call external APIs or an
LLM, it does not infer unstated intent, and it does not rewrite planning
records.
