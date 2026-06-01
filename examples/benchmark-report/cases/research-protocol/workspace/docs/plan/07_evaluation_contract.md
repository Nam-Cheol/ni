# Evaluation contract

## EVAL-001: Initial workspace scope review.

Method: Review docs/plan/** and .ni/contract.json to confirm the case is a
research_protocol benchmark with document, workflow, and human_service surfaces
and that it labels the resolved answers as synthetic benchmark fixture data,
not real research approval. Linked capabilities: CAP-001.

## EVAL-002: Blocker visibility review.

Method: Run go run ./cmd/ni status --dir
examples/benchmark-report/cases/research-protocol/workspace --proof
--next-questions and confirm OQ-001 through OQ-005 are resolved only after the
synthetic fixture answers are reflected in docs/plan/** and .ni/contract.json.
Linked capabilities: CAP-002.

## EVAL-003: Non-execution boundary review.

Method: Confirm the benchmark case records no lockfile, no bounded prompt, no
fieldwork, no participant recruitment, no data collection, no analysis, no
intervention placement decision, no downstream agent, no model API call, and
no empirical outcome claim unless the lockfile and bounded prompt were created
only by the allowed isolated CLI commands after readiness cleared. Linked
capabilities: CAP-003.

## EVAL-004: Isolated CLI readiness, lock, and prompt proof.

Method: Run go run ./cmd/ni status --dir
examples/benchmark-report/cases/research-protocol/workspace --proof
--next-questions; if the CLI reports READY or READY_WITH_DEFERRALS, run go run
./cmd/ni end --dir examples/benchmark-report/cases/research-protocol/workspace
and then go run ./cmd/ni run --dir
examples/benchmark-report/cases/research-protocol/workspace --max-chars 4000.
Record status proof, next questions, lock summary, bounded prompt summary, and
prompt character count without executing downstream work. Linked capabilities:
CAP-004.
