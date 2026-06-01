# 02. Grill Questions

The model first runs or requests:

```bash
ni status --dir . --proof --next-questions
```

If the CLI returns grouped deterministic questions, `ni-grill` asks those
first. In this example, assume no higher-priority `Sync repairs` or
`Open blockers` were returned, so grill pressure is applied to accepted draft
content.

```text
ni-grill result:

Grill findings:
1. GRILL-001 — Critical — acceptance evidence
   Affected: REQ-001 / CAP-001
   Concern: "Useful for support work" does not name who accepts the
   recommendation or what quality bar they apply.
   Why it matters: downstream work could optimize for generic plausibility
   instead of supervisor-approved triage.
   Question: Who must approve CAP-001 output, and what concrete evidence proves
   it is acceptable?
   Answer shape: reviewer role plus test, review checklist, demo condition,
   user approval, protocol check, or manual inspection
   Suggested action: clarify
   Blocks ni-end: likely yes

2. GRILL-002 — High — evaluation evidence
   Affected: EVAL-001 / docs/plan/07_evaluation_contract.md
   Concern: The evaluation says manual inspection but does not define the
   inspection checklist.
   Why it matters: an accepted capability has evidence only in name, not in an
   auditable condition.
   Question: What checklist items must EVAL-001 inspect before CAP-001 counts
   as complete?
   Answer shape: 3-5 checklist items tied to CAP-001 and REQ-001
   Suggested action: resolve
   Blocks ni-end: maybe

3. GRILL-003 — High — privacy/security risk
   Affected: RISK-001 / docs/plan/06_risks_security.md
   Concern: Refund authority is mitigated, but policy-data handling is not
   explicitly covered.
   Why it matters: privacy/security risk can invalidate the handoff even when
   refund authority is constrained.
   Question: Should policy and ticket data handling be a separate risk, a
   mitigation on RISK-001, or out of scope with reason?
   Answer shape: new risk, mitigation update, or not_applicable with reason
   Suggested action: resolve
   Blocks ni-end: maybe

4. GRILL-004 — Medium — non-goal boundary
   Affected: NG-001 / docs/plan/05_constraints.md
   Concern: The non-goal excludes refund issuing and customer contact but does
   not exclude queue, adapter, or runtime integration work.
   Why it matters: ni must remain pre-runtime and downstream work must not
   start from a hidden execution surface.
   Question: Should queues, shell/Codex adapters, and live support integrations
   be explicit non-goals for this plan?
   Answer shape: one or more non-goals, or not_applicable with reason
   Suggested action: mark non-goal
   Blocks ni-end: maybe
```

Only four questions are asked, below the five-question limit.
