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
  "question": "For CAP-001, what evidence proves this capability works, or should that evidence be deferred?"
}
```

Questions are concise, specific, and tied to readiness rule failures. They
reference the relevant contract ID or planning path when the failure message
contains one. Some questions explicitly allow `deferred` or `not_applicable`
answers where that is a valid planning status or a useful way to record that the
gap should remain visible.

## Rule Examples

`R004` accepted capability without evaluation:

```text
For CAP-001, what evidence proves this capability works, or should that evidence be deferred?
```

`R006` high-severity risk without mitigation:

```text
For RISK-001, what mitigation, owner, or explicit accepted-risk decision is required?
```

`R010` missing non-goal:

```text
For R010, what must this project explicitly avoid?
```

`R009` blocker open question:

```text
For OQ-001, what decision resolves this blocker, should it be deferred, or why must it remain blocking?
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
