# Readiness interview

`ni status --next-questions` turns deterministic readiness failures into
focused planning questions.

The questions are an interview aid for `ni-start`. They do not change readiness
state, do not resolve gaps, and do not use model judgment. The same readiness
rules that produce `issues[]` produce `next_questions[]` when requested.

`ni-start` may translate or summarize these deterministic questions into the
user's current language, but it must preserve the referenced IDs, locations,
answer shapes, command names, schema keys, target names, and status constants
exactly. Do not translate tokens such as `R014`, `OQ-001`, `SYNC-014`,
`ni status`, `.ni/contract.json`, `READY`, `BLOCKED`, or
`READY_WITH_DEFERRALS`. See
[`89_LANGUAGE_ADAPTIVE_AUTHORING.md`](89_LANGUAGE_ADAPTIVE_AUTHORING.md).

## Command

```bash
ni status --dir . --next-questions
ni status --dir . --proof --next-questions
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
  "group": "Evaluation evidence",
  "references": ["CAP-001"],
  "question": "CAP-001 has no evaluation. What evidence would prove this capability is complete?",
  "answer_shape": "test, review checklist, demo condition, user approval, protocol check, or manual inspection"
}
```

Questions are concise, specific, and tied to readiness rule failures. They
reference the relevant contract ID or planning path when the failure message
contains one. They are deterministic templates, not generated prose.

Each question should give the user a concrete answer shape, such as evidence,
an accepted decision, an explicit deferral, `not_applicable`, a mitigation, or
an explicit non-goal. The question must not imply implementation should begin
and must not pressure the user to accept uncertain planning state.

## Prioritization

`ni status --next-questions` returns a compact prioritized interview. It shows
at most three primary questions. When more deterministic questions exist, human
output reports how many lower-priority questions remain.

Question groups are ordered by readiness impact:

1. `First-run card`: when fresh `R014`, `R015`, and `R016` blockers appear
   together, ask exactly the purpose, actors/outcomes, and delivery-surface
   questions. Do not include unrelated lower-priority questions in that first
   card.
2. `Sync repairs`: when `SYNC-014`, `SYNC-015`, or `SYNC-016` exists, ask how
   to update the stale side of docs or contract. Do not ask the user to restate
   the whole project if one side already contains useful content.
3. `Risk decisions`: high-severity risks without mitigation ask for mitigation,
   owner/monitoring, explicit accepted-risk decision, or explicit deferral.
4. `Evaluation evidence`: accepted capabilities without evaluation ask what
   evidence proves completion.
5. `Scope boundaries`: missing non-goals ask what the project must explicitly
   avoid to reduce scope drift.
6. `Open blockers`: blocker open questions ask whether to resolve, defer with
   reason, or keep blocking with the missing information named.

If first-run blockers and first-run sync diagnostics appear together,
`ni status --next-questions` prioritizes sync repair because one side already
contains useful content. A first-run sync diagnostic shadows only its matching
generic first-run question: `SYNC-014` shadows `R014`, `SYNC-015` shadows
`R015`, and `SYNC-016` shadows `R016`. Other first-run gaps remain eligible and
are asked before unrelated lower-priority blockers. If neither side contains
useful content, the first-run card wins.

## Rule Examples

`R004` accepted capability without evaluation:

```text
CAP-001 has no evaluation. What evidence would prove this capability is complete?
Answer shape: test, review checklist, demo condition, user approval, protocol check, or manual inspection
```

`R006` high-severity risk without mitigation:

```text
RISK-001 is high severity and has no mitigation. What mitigation would reduce or monitor it, who owns it, or should this become an explicit accepted-risk decision?
```

`R010` missing non-goal:

```text
What must this project explicitly avoid so downstream work does not drift in scope?
Answer shape: one or more non-goals, or not_applicable with reason
```

`R009` blocker open question:

```text
OQ-001 is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?
Answer shape: accepted decision, deferral with reason, not_applicable, or keep blocking with reason
```

`R014` missing purpose:

```text
What should this project change, for whom, and why does it matter?
Answer shape: one or two sentences describing the desired reality change
```

`R015` missing actor or outcome:

```text
Who are the primary actors, and what outcome should each one get?
Answer shape: actor -> expected outcome
```

`R016` missing delivery surface:

```text
What is the likely delivery surface?
Answer shape: CLI, web app, conversation, document, workflow, research protocol, human service, or deferred with reason
```

`SYNC-014` purpose docs/contract drift:

```text
Project purpose is documented but missing from .ni/contract.json. Should .ni/contract.json be updated to match the docs, or is the docs text only a draft?
Answer shape: update contract / revise docs / revise both / keep blocker with reason
```

## First-run Card

When a fresh workspace reports `R014`, `R015`, and `R016`, `ni-start` should
group them into the opening planning card. The model should not ask broad
generic brainstorming questions and should not ask more than three questions at
once.

The card should say that `ni` is blocked only because the initial project
intent is not explicit enough to lock yet, and that implementation has not
started. Then it should ask:

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?

If the user's latest substantive message is Korean, ask the same three
human-facing questions in Korean while preserving tokens such as `ni`,
`BLOCKED`, `CLI`, `docs/plan/**`, and `.ni/contract.json`.

The user's answer should be recorded into both `docs/plan/**` and
`.ni/contract.json`: purpose in the project brief and `project.purpose`, actors
and outcomes in the actors doc and matching contract records, and delivery
surface in the delivery/operation doc plus `product_type` and
`delivery_surfaces` when clear. Uncertain answers remain assumptions or open
questions, clear exclusions become non-goals, and vague answers must not become
accepted decisions without confirmation.

After the authoring update, run or request:

```bash
ni status --dir . --proof --next-questions
```

The next readiness state still comes from the CLI.
`ni-start` should then show a planning proof block that summarizes what the
user said, which planning records and files changed, the before/after status
result, remaining blockers, and the next highest-priority question group.

`D001` deferred decision:

```text
DEC-001 is deferred. Does this deferred decision affect the next handoff, or should it remain visible without blocking?
```

`R012` docs/contract mismatch:

```text
DEC-001 differs between docs and contract. Which source is correct, and should the repair update docs, update the contract, defer the record, or mark it not_applicable?
```

## ni-start Usage

`ni-start` should run:

```bash
ni status --dir . --proof --next-questions
```

If JSON is easier to parse, use:

```bash
ni status --dir . --proof --json --next-questions
```

When questions are present, read the grouped output directly. Select the
highest-priority group in CLI order, preserve the group label, and ask at most
one group per turn unless the group itself is the compact `First-run card`.
Ask at most three primary questions at once. The skill may rephrase for
clarity, but it must preserve the referenced IDs, locations, readiness gap,
and answer shapes. It must not ask a question that implies implementation
work, pressures acceptance, invents broad brainstorming while deterministic
questions exist, or silently turns a blocker into a deferral.

If `Sync repairs` contains `SYNC-014`, `SYNC-015`, or `SYNC-016`, use those
repair questions instead of re-asking matching generic `R014`, `R015`, or
`R016` first-run questions. Ask whether to update contract, revise docs,
revise both, or keep the blocker with a reason.

After the user answers, `ni-start` updates `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json`, then runs or requests
`ni status --dir . --proof --next-questions` again. Readiness is blocked or
cleared by the deterministic CLI gate, not by model judgment.

The model's response after that update should include the planning proof block
defined in [`83_CONVERSATION_PROOF_CAPTURE.md`](83_CONVERSATION_PROOF_CAPTURE.md).

If no next questions are present and readiness is `READY` or
`READY_WITH_DEFERRALS`, `ni-start` may suggest `ni-end`. If readiness is
`BLOCKED` without questions, report the readiness issues directly.
