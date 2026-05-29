# Evaluation contract

## EVAL-001: Planning capture review

Method: Review `docs/plan/**` and `.ni/contract.json` to confirm the vague
dashboard request is represented as planning state, not as executable dashboard
implementation scope.

## EVAL-002: Blocked readiness proof

Method: Run `go run ./cmd/ni status --dir
examples/benchmark-report/cases/internal-dashboard/workspace --proof
--next-questions` from the repository root and verify it reports `BLOCKED`
with blocker open questions before downstream work starts.
