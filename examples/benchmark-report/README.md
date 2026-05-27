# Benchmark Report Example

## 1. Purpose

This directory is a deterministic template for manually reporting the intent
readiness benchmark defined in `docs/43_BENCHMARK_PROTOCOL.md`.

It contains no empirical results. Fill it only after running the protocol on a
specific request. Do not invent values for empty cells.

## 2. What this proves

- Benchmark reporting stays pre-runtime and evidence-based.
- Empty result cells remain `not_measured` until a real manual run exists.
- The benchmark compares direct-to-agent prompt readiness against the ni
  intent-lock path without executing either path.
- The report format makes status output, prompt boundedness, and reviewer
  notes auditable.

## 3. Product type / surface

- `product_type`: not applicable; this is a report template.
- `delivery_surface`: `document`
- Expected `ni status`: not applicable; this directory is not a ni workspace.
- Expected `ni run` target: not applicable.

## 4. Files

- `README.md`: the report template and boundary statement.
- `sample-report.md`: a fillable sample/template report with `not_measured`
  placeholders.
- `../../docs/43_BENCHMARK_PROTOCOL.md`: the benchmark protocol that defines
  the scoring method.

## 5. Commands

From the repository root:

```bash
test -f examples/benchmark-report/README.md
test -f examples/benchmark-report/sample-report.md
test -f docs/43_BENCHMARK_PROTOCOL.md
rg -n "not_measured|must not execute downstream agents|Target prompt boundedness" examples/benchmark-report/README.md examples/benchmark-report/sample-report.md docs/43_BENCHMARK_PROTOCOL.md
```

## 6. Expected output

The `test` commands should exit successfully.

The `rg` command should show `not_measured` markers in this template and
non-execution plus prompt-boundedness markers in the benchmark protocol.

## 7. Non-execution boundary

This report must remain intent-focused. It must not include downstream
execution traces, implementation results, telemetry, model API outputs, or
invented aggregate claims.

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
