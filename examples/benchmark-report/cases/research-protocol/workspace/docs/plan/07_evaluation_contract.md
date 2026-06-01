# Evaluation contract

## EVAL-001: Initial workspace scope review.

Method: Review docs/plan/** and .ni/contract.json to confirm the case is a
research_protocol benchmark with document, workflow, and human_service surfaces
and that it records the vague neighborhood cooling study request without
inventing blocker answers. Linked capabilities: CAP-001.

## EVAL-002: Blocker visibility review.

Method: Run go run ./cmd/ni status --dir
examples/benchmark-report/cases/research-protocol/workspace --proof
--next-questions and confirm OQ-001 through OQ-005 remain open blockers before
any lock or prompt compilation. Linked capabilities: CAP-002.

## EVAL-003: Non-execution boundary review.

Method: Confirm the benchmark case records no lockfile, no bounded prompt, no
fieldwork, no participant recruitment, no data collection, no analysis, no
intervention placement decision, no downstream agent, no model API call, and
no empirical outcome claim. Linked capabilities: CAP-003.
