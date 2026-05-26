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

The new workspace contains template `docs/plan/**` files and
`.ni/contract.json` with TODO planning records.

## Turn 1

**User**

Invoke `ni-start`. I want to plan a refund triage assistant for support agents.

**Model using ni-start**

Current planning state:

- Purpose: template still has TODO content.
- Delivery surface: conversation product for a human-to-system workflow.
- Accepted capabilities: none beyond the initialized template.
- Blocking gaps: purpose, actors, capabilities, evaluations, risks, artifacts,
  and blocker questions need concrete entries.

Focused questions:

1. Should the assistant only draft refund recommendations, or may it approve
   refunds automatically?
2. What transcript evidence should prove the assistant handled a refund request
   correctly?
3. What must be out of scope for the first locked plan?

**User**

It only drafts recommendations for support agents. It must never issue refunds
or contact customers. Readiness should be proven with three transcript fixtures:
eligible refund, ineligible refund, and policy ambiguity that escalates to a
supervisor.

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

After the edit, I will run `ni status --dir .`.

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
      "status": "accepted",
      "acceptance": "The assistant drafts refund recommendations for support agents and never issues refunds or contacts customers."
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
      "question": "Which policy source is authoritative for refund eligibility?",
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
ni status --dir .
```

Example result:

```text
BLOCKED
- OQ-001: Which policy source is authoritative for refund eligibility?
```

**Model using ni-start**

`ni status` reports `BLOCKED`, so I will not suggest `ni-end` yet. The next
focused question is: which refund policy source should `REQ-001` and
`EVAL-001` treat as authoritative?

## Turn 2

**User**

Use the internal refund policy page as the source of truth. If it conflicts
with the ticket text, escalate.

**Model using ni-start**

I will resolve `OQ-001`, add a decision that the internal refund policy page is
authoritative, add a risk for stale policy references, and update the transcript
evaluation to include conflict escalation. Then I will run `ni status --dir .`
again.

Example result:

```text
READY
```

`ni status` reports `READY`, so the planning gate can move to `ni-end`.
