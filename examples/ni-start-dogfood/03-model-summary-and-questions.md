# 03. model summary and questions

## Model using ni-start

Current planning state:

- Current purpose: missing.
- Active readiness profile: `prototype`.
- Product type / delivery surfaces: missing.
- Accepted capabilities: none yet.
- Unresolved blocker questions: the initial template question is still open.
- Recent decisions: none accepted yet.
- Next recommended planning focus: answer the first-run intent blockers before
  any execution or lock.

I will run or request the readiness gate with proof:

```bash
ni status --dir ./refund-triage-plan --proof --next-questions
```

## CLI result

```text
NI Intent Readiness: BLOCKED

Blockers:
- R014 Project purpose is missing.
  Why it matters: ni cannot lock intent until it knows what reality the project is meant to change.
  Next: describe the project in one or two sentences: what should change, for whom, and why it matters.
- OQ-001 is marked as blocker.
  Why it matters: open blocker questions mean required intent is still unresolved.
  Next: answer or defer the blocker question, or keep it blocking with an explicit reason.
- R015 Actors or outcomes are missing.
  Why it matters: ni cannot judge readiness without knowing who uses or operates the product and what successful use looks like for them.
  Next: list the primary actors and the outcome each one expects.
- R016 Delivery surface is missing.
  Why it matters: downstream handoff depends on knowing whether the product is delivered as a CLI, web app, conversation, document, workflow, research protocol, human service, or another surface.
  Next: choose the likely delivery surface, or mark it deferred with an explicit reason.

Execution must not start.

Next questions:
First-run card:
1. R014: What should this project change, for whom, and why does it matter?
   Answer shape: one or two sentences describing the desired reality change
2. R015: Who are the primary actors, and what outcome should each one get?
   Answer shape: actor -> expected outcome
3. R016: What is the likely delivery surface?
   Answer shape: CLI, web app, conversation, document, workflow, research protocol, human service, or deferred with reason
```

## First-run conversation card

ni is blocked because the initial project intent is not explicit enough to lock
yet. I need three things before execution can safely start: what reality this
project should change, who it is for, and how it will be delivered.

Implementation has not started. This is still planning.

Focused questions:

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?

## Boundary shown

The model uses R014, R015, and R016 as the opening planning card. It keeps the
template open question visible through the CLI output, but it asks only the
three first-run questions and does not start implementation.
