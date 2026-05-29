# Model edit safety

NI planning is conversation-authored, but model-authored edits must stay small,
visible, and conservative. A model may propose or apply a planning update from
conversation only when the update preserves the CLI as authority and keeps
uncertainty visible for human review.

## Core rules

Use these rules whenever a model updates `docs/plan/**`,
`.ni/contract.json`, or `.ni/session.json` from conversation.

1. Minimal diff rule: change only the files, records, and sections needed for
   the current planning turn. Do not rewrite unrelated prose, reorder stable
   records, or renumber IDs to make the plan look tidy.
2. Assumption vs decision rule: ambiguous, tentative, inferred, conflicting, or
   incomplete statements become assumptions, draft records, or open questions.
   They must not become accepted decisions.
3. User-confirmed decision rule: accepted decisions require explicit user
   confirmation or already-established accepted planning state. A model may
   suggest decision wording, but the suggestion is not accepted merely because
   the model wrote it.
4. No silent deletion rule: do not remove capabilities, requirements,
   evaluations, risks, non-goals, artifacts, decisions, or open questions
   without making the change visible. Prefer marking records as rejected,
   deferred, resolved, or not applicable when history matters.
5. Risk and evaluation integrity rule: do not weaken acceptance criteria,
   downgrade risks, remove mitigations, or loosen evaluations to make
   readiness pass. If validation fails, keep the gap visible.
6. Lock safety rule: after `.ni/plan.lock.json` exists, do not edit locked
   planning docs or matching contract records except through the amendment or
   relock flow. If a locked hash mismatch exists, stop and report `BLOCKED`.
7. Planning proof rule: after updating planning docs, show a short proof
   summary that names the user input captured, files changed, records affected,
   before/after `ni status --proof --next-questions` result, remaining
   assumptions or blockers, and next question group.

These rules apply to model-maintained authoring state. They do not add a task
runner, contract editing command surface, runtime adapter, evidence runner, or
automatic implementation behavior.

## Good planning updates

Good updates are narrow, traceable, and honest about uncertainty.

```text
User: We probably need a handoff prompt for Codex, but I am not sure whether it
should include evaluation details yet.

Good update:
- Add or update a draft capability for prompt handoff.
- Add an assumption that evaluation detail may be included.
- Add an open question asking whether evaluation detail is required for
  acceptance.
- Do not mark the inclusion decision as accepted.
```

```text
User: Yes, CAP-002 is accepted. It must produce a handoff prompt under 4000
characters and it should be checked by a golden fixture.

Good update:
- Mark CAP-002 accepted or keep it accepted if already accepted.
- Add or update the matching requirement and evaluation.
- Preserve existing IDs.
- Update both the relevant `docs/plan/**` file and `.ni/contract.json`.
- Summarize the changed records and run or report `ni status` when available.
```

```text
User: Remove the browser adapter from scope.

Good update:
- Mark the adapter capability, requirement, or decision as rejected or
  not applicable if it already exists.
- Add or update a non-goal that excludes browser adapter work.
- Keep any relevant risk or decision history visible.
```

## Bad planning updates

Bad updates hide uncertainty, create broad churn, or edit toward readiness
instead of truth.

```text
User: We probably need a handoff prompt.

Bad update:
- Mark "handoff prompt is accepted" as a decision.
- Add acceptance criteria the user did not confirm.
- Delete an open question because the model assumes the answer.
```

```text
User: Can we get this ready?

Bad update:
- Remove a high-severity risk to clear `ni status`.
- Downgrade an evaluation from fixture-based verification to "manual review"
  without user confirmation.
- Change accepted non-goals to make a requested implementation fit.
```

```text
Locked plan exists and hashes do not match.

Bad update:
- Edit `docs/plan/**` or `.ni/contract.json` directly.
- Treat the latest chat message as stronger authority than
  `.ni/plan.lock.json`.
- Continue as if the lock were still valid.
```

## Review checklist

Before finishing a model-authored planning turn, check:

- Did every accepted decision come from the user or existing accepted state?
- Are assumptions and unresolved questions still visible?
- Did the update preserve stable IDs and trace links?
- Were risks, mitigations, requirements, evaluations, and non-goals preserved
  unless explicitly changed?
- If the plan is locked, did the change use the amendment or relock flow?
- Did the response include a short planning proof summary?
