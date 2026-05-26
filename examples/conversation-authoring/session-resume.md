# ni-start session resume transcript

This fixture is illustrative. It shows how long-running planning resumes from
persisted docs and session state. It is not a runnable harness and it does not
execute downstream tools.

## First planning session

**User**

Invoke `ni-start`. I want to plan an internal onboarding checklist compiler for
team leads.

**Model using ni-start**

Current planning state:

- Purpose: template still has TODO content.
- Delivery surface: document output for a human-to-human workflow.
- Accepted capabilities: none beyond the initialized template.
- Blocking gaps: purpose, actors, capabilities, evaluations, risks, artifacts,
  and blocker questions need concrete entries.

Focused questions:

1. Should the compiler produce a checklist only, or also assign owners?
2. What evidence should prove the checklist is complete enough to hand to a new
   team lead?
3. What must be out of scope for the first locked plan?

**User**

It should produce the checklist only. It must not assign owners or send
notifications. Evidence should be a review fixture with a complete onboarding
scenario and one scenario with missing source material.

**Model using ni-start**

I will update the planning docs and contract together:

- `docs/plan/00_project_brief.md` with the onboarding checklist purpose.
- `docs/plan/02_capabilities.md` with `CAP-001 Compile onboarding checklist`.
- `docs/plan/05_constraints.md` with non-goals for owner assignment and
  notifications.
- `docs/plan/07_evaluation_contract.md` with review fixture evaluation.
- `docs/plan/10_open_questions.md` with source ownership as a blocker.
- `.ni/contract.json` with matching capability, requirement, evaluation,
  artifact, non-goal, and open question records.
- `.ni/session.json` with the active focus, summary, pending question, and last
  updated docs.

Then I will run `ni status --dir . --next-questions`.

## Docs updated

Example persisted contract records after the first session:

```json
{
  "capabilities": [
    {
      "id": "CAP-001",
      "title": "Compile onboarding checklist",
      "status": "accepted",
      "requirements": ["REQ-001"],
      "evaluations": ["EVAL-001"],
      "artifacts": ["ART-001"]
    }
  ],
  "requirements": [
    {
      "id": "REQ-001",
      "title": "Checklist-only output",
      "status": "accepted",
      "acceptance": "The compiler produces an onboarding checklist and does not assign owners or send notifications."
    }
  ],
  "evaluations": [
    {
      "id": "EVAL-001",
      "title": "Onboarding checklist review fixtures",
      "method": "Review one complete onboarding scenario and one missing-source scenario for correct checklist output and visible gaps."
    }
  ],
  "non_goals": [
    {
      "id": "NG-001",
      "title": "Do not assign owners or send notifications."
    }
  ],
  "open_questions": [
    {
      "id": "OQ-001",
      "question": "Which source documents are authoritative for onboarding checklist content?",
      "status": "open",
      "blocker": true
    }
  ]
}
```

Example `.ni/session.json` carryover state:

```json
{
  "schema": "ni.session.v0",
  "active_planning_focus": "Resolve authoritative source documents for onboarding checklist content.",
  "last_planning_summary": "CAP-001 is accepted for checklist-only output. NG-001 excludes owner assignment and notifications. OQ-001 blocks readiness until source documents are named.",
  "pending_questions": [
    {
      "id": "OQ-001",
      "title": "Which source documents are authoritative for onboarding checklist content?",
      "blocker": true
    }
  ],
  "recent_decisions": [],
  "recent_risks": [],
  "last_readiness_status": "BLOCKED",
  "last_readiness_blockers": ["OQ-001"],
  "last_updated_docs": [
    "docs/plan/00_project_brief.md",
    "docs/plan/02_capabilities.md",
    "docs/plan/05_constraints.md",
    "docs/plan/07_evaluation_contract.md",
    "docs/plan/10_open_questions.md",
    ".ni/contract.json"
  ],
  "authority_order": [
    ".ni/plan.lock.json",
    ".ni/contract.json",
    "docs/plan/**",
    ".ni/session.json",
    "chat history"
  ],
  "notes": {
    "planning_aid_only": "Session state is a planning aid, not authority.",
    "locked_docs": "Session state must not override locked docs.",
    "completion": "Session state must not mark docs complete.",
    "raw_transcript": "Raw transcript is not the source of truth."
  }
}
```

Example readiness result:

```text
BLOCKED
- OQ-001: Which source documents are authoritative for onboarding checklist content?
```

## Session ends

The chat ends before the user answers `OQ-001`. No hidden chat memory is needed
for resume because the docs, contract, and session file carry the planning
state.

## Later session resumes

**User**

Invoke `ni-start` and resume planning.

**Model using ni-start**

I read `AGENTS.md`, `.ni/contract.json`, `docs/plan/**`, and
`.ni/session.json`. I also ran `ni status --dir . --next-questions`.

Resumed state:

- Active focus from session state: resolve authoritative source documents for
  onboarding checklist content.
- Verified against contract/docs: `CAP-001` is accepted, `REQ-001` requires
  checklist-only output, and `NG-001` excludes owner assignment and
  notifications.
- Pending blocker: `OQ-001` is still open in the contract and still reported by
  `ni status`.
- Readiness: `BLOCKED`.

Focused question:

1. For `OQ-001`, which source documents are authoritative for onboarding
   checklist content, or should this remain blocking because the owner is not
   known yet?

## Resume without session state

If `.ni/session.json` is missing in the later session, the model reconstructs
the same summary from `docs/plan/**`, `.ni/contract.json`, and
`ni status --dir . --next-questions`. It should state that no session file was
available and recreate bounded session state after the next meaningful planning
edit.

## Resume conflict example

If `.ni/session.json` says `OQ-001` was resolved but `.ni/contract.json` still
marks `OQ-001` open, the model reports:

```text
Session conflict: .ni/session.json says OQ-001 was resolved, but
.ni/contract.json marks OQ-001 open and blocking. The contract wins, so I will
treat OQ-001 as pending and ask it again.
```

