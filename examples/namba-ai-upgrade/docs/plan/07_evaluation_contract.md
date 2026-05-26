# Evaluation contract

## EVAL-001: Limitation coverage review

Method: Inspect `docs/plan/00_project_brief.md` and `docs/plan/05_constraints.md` for the limitation list required by `REQ-001`.

## EVAL-002: SDD graph-boundary review

Method: Inspect `docs/plan/02_capabilities.md`, `docs/plan/04_domain_state.md`, and `docs/plan/09_execution_strategy.md` for dependency graph language and absence of mandatory execution tasks.

## EVAL-003: Pre-runtime boundary review

Method: Run `ni status` and inspect `docs/plan/03_interaction_contract.md` plus the generated prompt for no Codex execution, shell adapter, queue, or runtime state ownership.

## EVAL-004: Downstream target review

Method: Inspect `docs/plan/08_delivery_operation.md` for Codex, namba-ai, human-team, and generated harness targets with derived-only ownership.

## EVAL-005: Prompt artifact validation

Method: Run `ni end` and `ni run --target codex --out generated/codex.goal.prompt.txt`, then confirm the file is 4000 characters or less.
