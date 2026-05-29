# Benchmark Report Example

## 1. Purpose

This directory is a deterministic template and small pre-runtime case library for
manually reporting the Project Intent Compiler benchmark method defined in
`docs/43_BENCHMARK_PROTOCOL.md`.

The template contains no empirical results. Fill it only after running the
protocol on a specific request. Do not invent values for empty cells. Case
directories may include manual qualitative readiness drills. Unavailable lock,
run, and prompt-count evidence must remain `not_measured`.

## 2. What this proves

- Benchmark reporting stays pre-runtime and evidence-based.
- Empty result cells remain `not_measured` until a real manual run exists.
- The benchmark compares direct-to-agent prompt readiness against the ni
  intent-lock path without executing either path.
- The report format makes status output, prompt boundedness, and reviewer
  notes auditable.
- It is a case-study method, not a claim of empirical results or statistical
  significance.
- The internal-dashboard case shows how a plausible dashboard request hides
  users, success criteria, data boundaries, risks, non-goals, and handoff
  evidence before any downstream work starts. It includes real `ni status`
  evidence and stops at `BLOCKED`.

## 3. Product type / surface

- `product_type`: not applicable; this is a report template.
- `delivery_surface`: `document`
- Expected `ni status`: not applicable; this directory is not a ni workspace.
- Expected `ni run` target: not applicable.

## 4. Files

- `README.md`: the report template and boundary statement.
- `README.ko.md`: Korean companion guide.
- `sample-report.md`: a fillable sample/template report with `not_measured`
  placeholders.
- `cases/internal-dashboard/`: manual qualitative readiness drill for a vague
  dashboard request, with an isolated ni workspace and checked-in blocked
  status proof.
- `../../docs/43_BENCHMARK_PROTOCOL.md`: the benchmark protocol that defines
  the scoring method.

## 5. Commands

From the repository root:

```bash
test -f examples/benchmark-report/README.md
test -f examples/benchmark-report/sample-report.md
test -f examples/benchmark-report/cases/internal-dashboard/README.md
test -f examples/benchmark-report/cases/internal-dashboard/04-measurement-table.md
test -f examples/benchmark-report/cases/internal-dashboard/06-ni-status-proof.md
test -f examples/benchmark-report/cases/internal-dashboard/07-ni-next-questions.md
test -f docs/43_BENCHMARK_PROTOCOL.md
go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
rg -n "not_measured|must not execute downstream agents|Target prompt boundedness|internal-dashboard|NI Intent Readiness: BLOCKED" examples/benchmark-report/README.md examples/benchmark-report/sample-report.md examples/benchmark-report/cases/internal-dashboard/*.md docs/43_BENCHMARK_PROTOCOL.md
```

## 6. Expected output

The `test` commands should exit successfully.

The `ni status` command should report `NI Intent Readiness: BLOCKED` for the
internal-dashboard workspace.

The `rg` command should show `not_measured` markers in this template and
dashboard case, the checked-in blocked status proof, plus non-execution and
prompt-boundedness markers in the benchmark protocol.

## 7. demo-check coverage

Covered by `bash scripts/demo-check.sh`.

The demo check verifies required files, runs `ni status` for the isolated
internal-dashboard workspace, and checks that lock/run evidence remains absent
or `not_measured`. It does not run `ni end`, `ni run`, dashboard code, or
downstream agents.

## 8. Korean companion

Korean companion docs exist: `README.ko.md`.

## 9. Non-execution boundary

This report must remain intent-focused. It must not include downstream
execution traces, implementation results, telemetry, model API outputs, or
invented aggregate claims. It must not claim statistical significance.

The internal-dashboard case is not a product demo. It must not become a
dashboard scaffold, runtime harness, queue, shell adapter, model call, or
downstream-agent run.

## Run Metadata

| Field | Value |
| --- | --- |
| Request fixture | `not_measured` |
| Scoring date | `not_measured` |
| Reviewer | `not_measured` |
| ni version or commit | `not_measured` |
| Direct prompt source | `not_measured` |
| ni workspace path | `not_measured` |

The checked-in internal-dashboard case fills these fields in its case files;
the generic report template above remains unfilled by design.

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
