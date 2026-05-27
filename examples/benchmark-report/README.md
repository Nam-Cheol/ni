# Benchmark Report Example

This directory is a template for manually reporting the intent readiness
benchmark defined in
[`docs/43_BENCHMARK_PROTOCOL.md`](../../docs/43_BENCHMARK_PROTOCOL.md).

It contains no empirical results. Fill it only after running the protocol on a
specific request. Do not invent values for empty cells.

## Run Metadata

| Field | Value |
| --- | --- |
| Request fixture | `not_measured` |
| Scoring date | `not_measured` |
| Reviewer | `not_measured` |
| ni version or commit | `not_measured` |
| Direct prompt source | `not_measured` |
| ni workspace path | `not_measured` |

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
- `ni status` output,
- `ni end` output when readiness passed,
- `ni run` character count when a prompt was compiled,
- reviewer scoring notes.

## Boundary

This report must remain pre-runtime and intent-focused. It must not include
downstream execution traces, implementation results, telemetry, model API
outputs, or invented aggregate claims.
