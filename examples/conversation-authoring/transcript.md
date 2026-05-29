# ni-start conversation authoring transcript

This fixture is illustrative. It shows the shape of a sustained planning turn,
not exact output from a live model.

## Initial workspace

```bash
ni init --dir ./examples/refund-triage-plan \
  --product-type conversation_product \
  --surface conversation \
  --interaction-mode human_to_system
```

The new workspace contains template `docs/plan/**` files, `.ni/contract.json`,
and bounded `.ni/session.json` state. The user does not type contract
authoring commands; the model maintains those files from conversation.

The completed planning state produced by this transcript is checked in at
`examples/conversation-authoring/docs/plan/**` and
`examples/conversation-authoring/.ni/contract.json`.

## Turn 1

**User**

Invoke `ni-start`. I want to plan a refund triage assistant for support agents.

**Model using ni-start**

I will run the readiness gate with proof and grouped next questions before
asking anything. The CLI output is the interview driver.

```bash
ni status --dir . --proof --next-questions
```

Example CLI result:

```text
NI Intent Readiness: BLOCKED

Blockers:
- R014 Project purpose is missing.
- R015 Actors or outcomes are missing.
- R016 Delivery surface is missing.
- OQ-001 is marked as blocker.

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

Current planning state:

- Purpose: template still has TODO content.
- Delivery surface: conversation product for a human-to-system workflow.
- Accepted capabilities: none beyond the initialized template.
- Blocking gaps: purpose, actors, capabilities, evaluations, risks, artifacts,
  and blocker questions need concrete entries.

I will ask the highest-priority group from the CLI result first. The compact
`First-run card` is allowed to include all three primary questions, and I will
not add unrelated lower-priority questions.

Group: `First-run card`

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?

**User**

It changes refund triage from ad hoc judgment into a documented recommendation
workflow for support agents and supervisors. Support agents get a draft refund
recommendation; supervisors get escalations for ambiguous policy cases.
Deliver it as a conversation product with a document handoff.

It only drafts recommendations for support agents. It must never issue refunds
or contact customers. Readiness should be proven with three transcript
fixtures: eligible refund, ineligible refund, and policy ambiguity that
escalates to a supervisor.

**Model using ni-start**

I will update:

- `docs/plan/00_project_brief.md` with the refund triage purpose.
- `docs/plan/02_capabilities.md` with `CAP-001 Draft refund recommendations`.
- `docs/plan/05_constraints.md` with non-goals for issuing refunds and customer
  contact.
- `docs/plan/07_evaluation_contract.md` with transcript fixture evaluations.
- `docs/plan/10_open_questions.md` with remaining policy and owner questions.
- `.ni/contract.json` with matching capability, requirement, evaluation,
  risk, artifact, non-goal, and open question records.
- `.ni/session.json` with bounded continuity state and the latest CLI
  readiness blockers.

After the edit, I will run `ni status --dir . --proof --next-questions`.

Planning proof:

- User input captured:
  "Refund triage should become a documented recommendation workflow for
  support agents and supervisors, delivered through conversation plus a
  document handoff, with refund issuing and customer contact excluded."
- Interpreted planning records:
  - Purpose: documented refund recommendation workflow.
  - Actors/outcomes: support agents get draft recommendations; supervisors get
    ambiguity escalations.
  - Delivery surface: conversation product with document handoff.
  - Capabilities: `CAP-001` draft refund recommendations.
  - Requirements: `REQ-001` recommendations only.
  - Risks: `RISK-001` refund authority confusion.
  - Evaluations: `EVAL-001` transcript fixtures.
  - Decisions: none newly accepted beyond explicit scope exclusions.
  - Assumptions: none promoted to accepted decisions.
  - Non-goals: `NG-001` no refunds issued or customer contact.
  - Open questions: `OQ-001` authoritative refund policy source remains
    blocking.
- Updated planning artifacts:
  - `docs/plan/00_project_brief.md`: purpose recorded.
  - `docs/plan/02_capabilities.md`: `CAP-001` recorded.
  - `docs/plan/05_constraints.md`: refund and customer-contact non-goals
    recorded.
  - `docs/plan/07_evaluation_contract.md`: transcript fixture evaluation
    recorded.
  - `docs/plan/10_open_questions.md`: `OQ-001` recorded as blocking.
  - `.ni/contract.json`: matching capability, requirement, evaluation, risk,
    artifact, non-goal, and open question records updated.
  - `.ni/session.json`: bounded continuity state and latest readiness blockers
    refreshed.
- Status result:
  - before: `BLOCKED` because `R014`, `R015`, `R016`, and `OQ-001` were
    unresolved.
  - after: pending the next `ni status --dir . --proof --next-questions` run.
- Remaining blockers:
  - `OQ-001`: authoritative refund policy source.
- Next question group:
  - pending CLI output.

## Persisted planning records

Example contract records after the turn:

```json
{
  "capabilities": [
    {
      "id": "CAP-001",
      "title": "Draft refund recommendations",
      "status": "accepted",
      "requirements": ["REQ-001"],
      "evaluations": ["EVAL-001"],
      "risks": ["RISK-001"],
      "artifacts": ["ART-001"]
    }
  ],
  "requirements": [
    {
      "id": "REQ-001",
      "title": "Recommendations only",
      "status": "accepted"
    }
  ],
  "evaluations": [
    {
      "id": "EVAL-001",
      "title": "Refund triage transcript fixtures",
      "method": "Review eligible refund, ineligible refund, and ambiguous policy transcripts for correct recommendation and escalation behavior."
    }
  ],
  "non_goals": [
    {
      "id": "NG-001",
      "title": "Do not issue refunds or contact customers."
    }
  ],
  "open_questions": [
    {
      "id": "OQ-001",
      "title": "Which policy source is authoritative for refund eligibility?",
      "status": "open",
      "blocker": true
    }
  ]
}
```

Example human-facing doc update:

```text
CAP-001 Draft refund recommendations

The assistant helps support agents evaluate refund requests and draft a
recommendation. It does not issue refunds, send customer messages, or bypass
supervisor review when the policy is ambiguous.

Evaluation: EVAL-001 checks three transcript fixtures: eligible refund,
ineligible refund, and ambiguous policy escalation.
```

## Gate result

```bash
ni status --dir . --proof --next-questions
```

Example result:

```text
NI Intent Readiness: BLOCKED

Blockers:
- OQ-001 is marked as blocker.

Execution must not start.

Next questions:
Open blockers:
1. OQ-001: OQ-001 is blocking readiness. Should it be resolved, deferred with reason, or kept blocking with the missing information named?
   Answer shape: accepted decision, deferral with reason, not_applicable, or keep blocking with reason
```

**Model using ni-start**

`ni status` reports `BLOCKED`, so I will not suggest `ni-end` yet. The next
focused group is `Open blockers`: which refund policy source should `REQ-001`
and `EVAL-001` treat as authoritative, should it be deferred with reason, or
should it remain blocking?

## Turn 2

**User**

Use the internal refund policy page as the source of truth. If it conflicts
with the ticket text, escalate.

**Model using ni-start**

I will resolve `OQ-001`, add `DEC-001` that the internal refund policy page is
authoritative, add `RISK-003` for stale policy references, update the transcript
evaluation to include conflict escalation, and refresh `.ni/session.json` as
bounded continuity state. Then I will run
`ni status --dir . --proof --next-questions` again.

Changed files:

- `docs/plan/00_project_brief.md`
- `docs/plan/02_capabilities.md`
- `docs/plan/03_interaction_contract.md`
- `docs/plan/04_domain_state.md`
- `docs/plan/06_risks_security.md`
- `docs/plan/07_evaluation_contract.md`
- `docs/plan/10_open_questions.md`
- `docs/plan/11_decision_log.md`
- `.ni/contract.json`
- `.ni/session.json`

Example result:

```text
READY_WITH_DEFERRALS
- DEC-004: Final production owner for maintaining the refund policy page.
- OQ-002: Which support dashboard will eventually display the recommendation draft?
```

`ni status` reports `READY_WITH_DEFERRALS`, so the planning gate can move to
`ni-end` only after the deferrals are summarized and the user explicitly
confirms the lock. The deferrals remain visible; the model does not silently
resolve them to make the plan look cleaner.

At no point did the user type contract `add`, `list`, or `set` commands. The
model maintained the docs and contract from conversation, while the CLI reported
gaps and readiness.

## Turn 3

**User**

Invoke `ni-end`.

**Model using ni-end**

I will check the gate through the CLI before asking for lock confirmation.

```bash
ni status --dir .
```

CLI result:

```text
READY_WITH_DEFERRALS
- DEC-004 is deferred.
- OQ-002 remains open and non-blocking.
```

The CLI says the prototype plan can lock with visible deferrals. The lock will
hash `.ni/contract.json` and `docs/plan/**`.

Confirm that I should run `ni end --dir .` and let the CLI write
`.ni/plan.lock.json`?

**User**

Confirmed. Run ni end.

**Model using ni-end**

```bash
ni end --dir .
```

CLI result:

```text
locked plan at .ni/plan.lock.json
status READY_WITH_DEFERRALS
```

The CLI wrote `.ni/plan.lock.json`. The model did not create or edit the lock
by hand.

## Turn 4

**User**

Invoke `ni-run` for the human team.

**Model using ni-run**

I will let the CLI verify the lock hashes and compile the target prompt.

```bash
ni run --dir . --target human-team --max-chars 4000 --out generated/human-team.prompt.txt
```

CLI result:

```text
compiled prompt at generated/human-team.prompt.txt
```

The prompt is 4000 characters or less. `ni` compiled this prompt only. It did
not execute implementation, Codex, shell commands, adapters, queues, PR
automation, or downstream runtime work.
