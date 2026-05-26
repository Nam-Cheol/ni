# ni-end confirmation transcript

This fixture is illustrative. It shows the confirmation shape for `ni-end`; it
is not exact output from a live model.

## Ready with deferrals

**User**

Invoke `ni-end` for this plan.

**Model using ni-end**

I will check readiness through the CLI before discussing lock confirmation.

```bash
ni status --dir .
```

CLI result:

```text
READY_WITH_DEFERRALS
profile: prototype
product type: conversation_product
```

The CLI reports `READY_WITH_DEFERRALS`, so the plan can be locked only after
you confirm the contract summary below.

Pre-lock summary:

- Project: Refund Triage Assistant Plan.
- Purpose: draft refund recommendations for support agents without issuing
  refunds or contacting customers.
- Readiness profile: prototype.
- Product type: conversation_product.
- Delivery surfaces: conversation, document.
- Interaction mode: human_to_system.
- Accepted capabilities: `CAP-001 Draft refund recommendations`, `CAP-002
  Escalate ambiguous or conflicting cases`, `CAP-003 Maintain docs and contract
  from conversation`.
- Linked requirement: `REQ-001 Recommendations only`.
- Linked evaluation: `EVAL-001 Refund triage transcript fixtures`.
- Linked artifact: `ART-001 docs/plan/`.
- High-severity risks: `RISK-001 Refund authority may be implied`, `RISK-003
  Policy source may be stale or unclear`.
- Deferred decisions: `DEC-004 Final production policy owner`.
- Open non-blocking questions: `OQ-002 Which dashboard will display the
  recommendation draft?`
- Source files to lock: `.ni/contract.json` and `docs/plan/**`.

Confirm that I should run `ni end --dir .` and let the CLI write
`.ni/plan.lock.json`?

**User**

What happens to the deferred production policy owner decision?

**Model using ni-end**

It remains visible as `DEC-004` with `status=deferred` in the planning contract
and is included in the lock context. The CLI has classified it as non-blocking
under the current readiness profile, so it does not prevent locking, but it is
not resolved by the lock.

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

`ni end` wrote `.ni/plan.lock.json`. The locked readiness status is
`READY_WITH_DEFERRALS`.

## Blocked refusal

**User**

Invoke `ni-end`.

**Model using ni-end**

```bash
ni status --dir .
```

CLI result:

```text
BLOCKED
- OQ-001: Which policy source is authoritative for refund eligibility?
```

`BLOCKED`. I will not run `ni end --dir .`. Resolve `OQ-001` through planning
authoring, then run `ni status --dir .` again before attempting to lock.
