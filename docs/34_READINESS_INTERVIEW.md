# Readiness interview

`ni status --next-questions` turns deterministic readiness failures into
focused planning questions.

The questions are an interview aid for `ni-start`. They do not change readiness
state, do not resolve gaps, and do not use model judgment. The same readiness
rules that produce `issues[]` produce `next_questions[]` when requested.

## Command

```bash
ni status --dir . --next-questions
ni status --dir . --json --next-questions
```

Plain text output appends `question` lines after readiness issues. JSON output
adds `next_questions` only when `--next-questions` is present.

Without `--next-questions`, `ni status --json` keeps the existing status shape
and omits `next_questions`.

## Question Shape

Each question includes:

```json
{
  "rule_id": "R004",
  "severity": "blocker",
  "references": ["CAP-001"],
  "question": "CAP-001 has no evaluation. What evidence would prove this capability is complete: a test, review checklist, demo condition, user approval, or an explicit deferral?"
}
```

Questions are concise, specific, and tied to readiness rule failures. They
reference the relevant contract ID or planning path when the failure message
contains one. They are deterministic templates, not generated prose.

Each question should give the user a concrete answer shape, such as evidence,
an accepted decision, an explicit deferral, `not_applicable`, a mitigation, or
an explicit non-goal. The question must not imply implementation should begin
and must not pressure the user to accept uncertain planning state.

## Rule Examples

`R004` accepted capability without evaluation:

```text
CAP-001 has no evaluation. What evidence would prove this capability is complete: a test, review checklist, demo condition, user approval, or an explicit deferral?
```

`R006` high-severity risk without mitigation:

```text
RISK-001 is high severity and has no mitigation. What mitigation would reduce or monitor it, who owns it, or should this become an explicit accepted-risk decision?
```

`R010` missing non-goal:

```text
No non-goal is recorded. What explicit non-goal should bound the plan, or why is this boundary not_applicable?
```

`R009` blocker open question:

```text
OQ-001 is blocking readiness. What answer would resolve it: an accepted decision, a deferral with reason, not_applicable, or keeping it blocking with the missing information named?
```

`R014` missing purpose:

```text
project.purpose is missing a concrete purpose. What purpose should be recorded: a user outcome, evidence-backed assumption, accepted decision, or explicit deferral?
```

`R015` missing actor or outcome:

```text
docs/plan/01_actors_outcomes.md is missing an actor or outcome. Which actor needs what outcome, and should that record be accepted, kept as evidence, deferred, or marked not_applicable?
```

`R016` missing delivery surface:

```text
docs/plan/08_delivery_operation.md is missing a delivery surface. Which surface should the plan target: cli, web, api, conversation, document, workflow, human_service, physical, or a deferral with reason?
```

`D001` deferred decision:

```text
DEC-001 is deferred. Should it remain deferred with a reason, become an accepted or rejected decision, or be marked not_applicable?
```

`R012` docs/contract mismatch:

```text
DEC-001 differs between docs and contract. Which source is correct, and should the repair update docs, update the contract, defer the record, or mark it not_applicable?
```

## ni-start Usage

`ni-start` should run:

```bash
ni status --dir . --next-questions
```

If JSON is easier to parse, use:

```bash
ni status --dir . --json --next-questions
```

When questions are present, ask the highest-impact one to three questions from
the CLI result. The skill may rephrase for clarity, but it must preserve the
referenced IDs, the readiness gap, and the allowed outcomes. It must not ask a
question that implies implementation work, pressures acceptance, or silently
turns a blocker into a deferral.

If no next questions are present and readiness is `READY` or
`READY_WITH_DEFERRALS`, `ni-start` may suggest `ni-end`. If readiness is
`BLOCKED` without questions, report the readiness issues directly.
