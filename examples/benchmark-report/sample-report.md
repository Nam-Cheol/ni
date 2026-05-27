# Sample Benchmark Report Template

This is a sample/template for manual reporting only. It contains no empirical
results, no aggregate claims, and no downstream execution traces. Keep cells as
`not_measured` until a real reviewer runs the protocol.

## Run Metadata

| Field | Value |
| --- | --- |
| Request fixture | `not_measured` |
| Scoring date | `not_measured` |
| Reviewer role | `not_measured` |
| ni version or commit | `not_measured` |
| Direct prompt source | `not_measured` |
| ni workspace path | `not_measured` |
| Status command | `not_measured` |
| Target prompt command | `not_measured` |

## Fixture Review Aids

| Fixture File | Used? | Notes |
| --- | --- | --- |
| `request.md` | `not_measured` | Direct-to-agent prompt source. |
| `expected-hidden-assumptions.md` | `not_measured` | Reviewer aid, not measured output. |
| `expected-readiness-gaps.md` | `not_measured` | Reviewer aid, not measured output. |
| `suggested-ni-questions.md` | `not_measured` | Interview aid, not an answer key. |

## Metric Table

| Metric | Direct-to-agent Prompt | ni Intent-Lock Path | Notes |
| --- | --- | --- | --- |
| Missing acceptance criteria count | `not_measured` | `not_measured` | Count criteria gaps visible before execution. |
| Unmitigated high-risk count | `not_measured` | `not_measured` | Count high risks without mitigation or accepted rationale. |
| Unresolved blocker count | `not_measured` | `not_measured` | Count blockers that should prevent trustworthy execution. |
| Hidden assumption count | `not_measured` | `not_measured` | Count material assumptions a downstream actor would need to invent. |
| Non-goal coverage | `not_measured` | `not_measured` | Use `none`, `partial`, or `explicit`. |
| Stale plan detection | `not_measured` | `not_measured` | Use `not_applicable`, `passes`, or `blocked`. |
| Target prompt boundedness | `not_measured` | `not_measured` | Record character count and pass/fail against the configured maximum. |
| Readiness status before execution | `not_measured` | `not_measured` | For the ni path, use authoritative `ni status` output. |

## Evidence

Paste or link only evidence produced by a real manual run:

- request text,
- direct prompt,
- reviewer notes for direct prompt scoring,
- ni-start conversation summary,
- `ni status` output,
- `ni end` output when readiness passed,
- `ni run` character count when a prompt was compiled,
- reviewer scoring notes for the ni path.

## Observations

`not_measured`

## Conclusions

`not_measured`

Do not claim that ni reduces rework, improves implementation quality, or
outperforms direct prompting unless real benchmark evidence has been collected
and reported.
