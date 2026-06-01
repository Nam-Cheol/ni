# NI Grill Output Budget

`ni-grill` is useful only when its critique is compact enough for the user to
answer. This document defines the severity model, output budget, and finding
format for grill output.

The boundary remains unchanged:

```text
Skills are UX; CLI is authority.
```

`ni-grill` may apply advisory planning pressure. It must not become a second
readiness gate, approve lock, refuse lock by model judgment, execute generated
prompts, or weaken acceptance criteria to make a plan look ready.

## Why The Budget Exists

Planning critique can expand faster than intent can be resolved. A long list
of equal-weight concerns makes it harder for users to see what should change
before `ni-end`.

The output budget forces `ni-grill` to:

- preserve deterministic `ni status` blockers first;
- show the highest-risk planning concerns before editorial issues;
- keep questions answerable in one turn;
- separate CLI blockers from advisory findings;
- summarize lower-priority issues instead of flooding the user.

## Severity Model

Severity is `ni-grill`'s advisory planning pressure. It is not the same as CLI
readiness. Only `ni status`, user confirmation, and `ni end` govern lock.

Use these labels exactly:

| Severity | Meaning | Examples |
| --- | --- | --- |
| `Critical` | Likely must be addressed before `ni-end` because it points to a serious planning-quality or claim-boundary problem the CLI may not fully capture. | Public claim not supported by evidence; privacy/safety boundary ambiguous; accepted capability has weak evidence wording even if an evaluation exists; benchmark says `READY` but `not_measured` boundary is hidden; downstream handoff could imply implementation authorization. |
| `High` | Important before lock, but may be resolved through user confirmation, deferral, or explicit non-goal. | Actor or outcome too broad; delivery surface ambiguity in prose; risk mitigation vague; target handoff unclear. |
| `Medium` | Improves clarity and reduces rework, but may not block the next lock if the user accepts the tradeoff. | Wording could be more precise; evidence could name a reviewer more clearly; assumptions should be promoted to open questions. |
| `Low` | Editorial or maintainability improvement. | Duplicate wording; docs organization; minor example clarity. |
| `Note` | Observation only. No question is required unless the user wants to refine. | Well-handled non-goal; clear claim boundary; useful pattern to repeat. |

## Output Budget

By default, `ni-grill` should show at most 5 findings in one turn.

Rules:

- Show deterministic CLI blockers from `ni status` before advisory findings.
- If `Critical` or `High` findings exist, show at most 3 of them first.
- Ask no more than 5 user-facing questions in one turn.
- Avoid mixing many categories unless required by the top risks.
- Do not repeat every deterministic blocker unless explaining it helps the
  user answer.
- If more issues exist, write: `N additional lower-priority findings were not shown.`

When `ni status` is `BLOCKED`, prioritize deterministic blockers and keep any
secondary grill critique short. When `ni status` is `READY` or
`READY_WITH_DEFERRALS`, focus on claim quality, public handoff, risk clarity,
and overclaim prevention.

## Prioritization

Order findings by:

1. deterministic CLI blockers from `ni status`;
2. `Critical` ni-grill findings;
3. `High` ni-grill findings;
4. acceptance evidence gaps;
5. privacy/security/safety risks;
6. claim-boundary risks;
7. non-goal or scope-drift risks;
8. handoff ambiguity;
9. `Medium` or `Low` editorial issues;
10. `Note` observations.

## Finding Format

Use this shape:

```text
Grill findings:
1. GRILL-001 — <severity> — <category>
   Affected: <file path or planning ID>
   Concern: <specific concern>
   Why it matters: <why downstream handoff or lock quality could suffer>
   Question: <user-facing question>
   Answer shape: <expected answer form>
   Suggested action: <resolve / defer / mark non-goal / clarify / keep as note>
   Blocks ni-end: <CLI decides / likely yes / maybe / no>
```

Use `Blocks ni-end: CLI decides` when the finding corresponds to deterministic
readiness. Use `Blocks ni-end: likely yes` only for severe planning-quality
issues that should be resolved before lock. Use `Blocks ni-end: maybe` for
user-confirmable tradeoffs. Use `Blocks ni-end: no` for clarity, editorial, or
observation findings.

## Language Behavior

Ask user-facing questions in the user's latest substantive language. Preserve
IDs, commands, paths, schema keys, target names, status constants, and severity
labels exactly unless the surrounding prose also gives a translation. Examples:
`GRILL-001`, `R014`, `OQ-001`, `SYNC-014`, `ni status`,
`.ni/contract.json`, `READY`, `READY_WITH_DEFERRALS`, `BLOCKED`, `Critical`,
`High`, `Medium`, `Low`, and `Note`.

## Good And Bad Output

Bad:

- Lists 12 findings with equal weight.
- Asks generic questions such as "Can you clarify this?"
- Repeats all `ni status` blockers and then adds many unrelated critiques.
- Says "do not lock" by model judgment alone.
- Mixes execution advice into a planning challenge.

Good:

- Shows the top 3 findings when those are enough.
- Labels severity and category.
- Asks a specific user-facing question with an answer shape.
- Preserves CLI authority.
- Summarizes omitted lower-priority findings.

## Before And After

Bad:

```text
GRILL-001: Evidence is vague. Can you clarify?
GRILL-002: Risk is unclear. Can you clarify?
GRILL-003: Handoff is unclear. Can you clarify?
GRILL-004: Non-goal is unclear. Can you clarify?
GRILL-005: Claim is unclear. Can you clarify?
GRILL-006: The docs could be better.
```

Good:

```text
Grill findings:
1. GRILL-001 — Critical — claim boundary
   Affected: examples/benchmark-report/cases/research-protocol/15-before-after-evidence.md
   Concern: The `READY` transition can be quoted without the nearby
   `not_measured` research-approval boundary.
   Why it matters: a reader could mistake benchmark artifact readiness for
   real fieldwork authorization.
   Question: Should the transition row say that real research approval,
   fieldwork authorization, research quality, and intervention effectiveness
   remain `not_measured`?
   Answer shape: yes/no plus exact row wording, or rationale that the existing
   scope note is sufficient.
   Suggested action: clarify
   Blocks ni-end: maybe

2 additional lower-priority findings were not shown.
```

## CLI Authority Boundary

`ni-grill` may say a finding likely should be addressed before lock, but it
must still route readiness and locking through the CLI:

```text
ni status --dir . --proof --next-questions
ni end --dir .
```

If `ni status` reports `BLOCKED`, report the deterministic blockers and stop
before `ni-end`. If a lock hash mismatch exists, stop and report `BLOCKED`.
