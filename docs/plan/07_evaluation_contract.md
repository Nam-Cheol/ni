# Evaluation contract

## EVAL-001: incomplete docs are blocked

Method: fixture test. A project missing required docs must return `BLOCKED`.

## EVAL-002: capability without evaluation is blocked

Method: fixture test. Accepted capabilities must link to at least one evaluation.

## EVAL-003: high risk without mitigation is blocked

Method: fixture test.

## EVAL-004: stale lock blocks run

Method: fixture test. Modify a locked doc and verify `ni run` refuses.

## EVAL-005: prompt budget is enforced

Method: unit test. `ni run --max-chars 4000` must not exceed the limit.
