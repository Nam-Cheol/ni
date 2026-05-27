# Evaluation contract

## EVAL-001: ni init creates conversation-authoring workspace files

Method: CLI integration test for `docs/plan/**`, `.ni/contract.json`, readiness config, and bounded session state.

## EVAL-002: capability without evaluation is blocked

Method: readiness fixture test.

## EVAL-003: high risk without mitigation is blocked

Method: readiness fixture test.

## EVAL-004: stale lock blocks run

Method: lock and prompt fixture test.

## EVAL-005: handoff prompt budget is enforced

Method: unit and CLI integration test for `ni run --target codex --max-chars 4000`.

## EVAL-006: conversation authoring preserves CLI authority and excludes contract add/list/set primary UX

Method: static quality check for skills, conversation-authoring docs, and forbidden contract editing commands.

## EVAL-007: readiness profiles initialize correctly

Method: CLI integration tests for `ni init --profile`.

## EVAL-008: readiness profile appears in status

Method: CLI integration test for `ni status`.

## EVAL-009: related-work boundary is enforced in docs

Method: static core-boundary quality check.

## EVAL-010: product type guidance is deterministic

Method: contract and readiness fixture tests for supported product types, delivery surfaces, and status guidance.

## EVAL-011: downstream target registry is deterministic

Method: target unit test and CLI target listing test.

## EVAL-012: target exports are locked seed material only

Method: CLI integration tests that verify lock preconditions and forbidden runtime paths.

## EVAL-013: feedback ingest is inert

Method: CLI integration tests that confirm feedback does not modify contract or lock files.

## EVAL-014: pressure promotion is explicit

Method: pressure ledger unit and CLI tests.

## EVAL-015: harness candidate lifecycle is non-executing

Method: harness unit and CLI lifecycle tests.

## EVAL-016: amendment and relock require explicit approval

Method: CLI integration test for amend, stale lock rejection, applied amendment, relock, and previous lock archive.

## EVAL-017: collaboration conflicts are deterministic

Method: collab unit tests for changed IDs, contradictory decisions, weakened requirements, and lock mismatch handling.

## EVAL-018: differentiation proof assets stay pre-runtime and synchronized

Method: static quality and release-readiness checks verify README.md, README.ko.md, docs/40_POSITIONING.md through docs/46_RELEASE_READINESS.md, examples, status proof output, and prompt budget claims preserve Project Intent Compiler positioning without runtime execution claims.
