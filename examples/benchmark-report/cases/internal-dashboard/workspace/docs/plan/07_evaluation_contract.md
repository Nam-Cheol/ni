# Evaluation contract

## EVAL-001: Benchmark artifact scope review

Method: Review `docs/plan/**` and `.ni/contract.json` to confirm `OQ-001`
through `OQ-004` are resolved only for benchmark planning-meeting artifact
readiness and not as dashboard product readiness.

Capability: CAP-001

## EVAL-002: Answer packet completeness review

Method: Confirm every required answer field is filled, the supported decision
is explicit, pass/fail criteria are testable, privacy boundaries are explicit,
and unresolved blockers are marked instead of answered implicitly.

Capability: CAP-002

## EVAL-003: Privacy and freshness boundary review

Method: Confirm the packet includes only allowed benchmark validation fields,
excludes sensitive data categories, states freshness expectations, and marks
stale or unknown source status as stale or TBD.

Capability: CAP-003

## EVAL-004: Isolated CLI readiness and prompt proof

Method: Run `go run ./cmd/ni status --dir
examples/benchmark-report/cases/internal-dashboard/workspace --proof
--next-questions`; if the CLI reports `READY` or `READY_WITH_DEFERRALS`, run
`go run ./cmd/ni end --dir examples/benchmark-report/cases/internal-dashboard/workspace`
and then `go run ./cmd/ni run --dir
examples/benchmark-report/cases/internal-dashboard/workspace --max-chars 4000`.
Record the status proof, next questions, lock summary, bounded prompt summary,
and prompt character count without executing downstream work.

Capability: CAP-004
