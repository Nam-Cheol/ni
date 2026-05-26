# 004 - Contract model

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Implement typed loading and validation primitives for `.ni/contract.json`.

Goal:
The CLI can load the contract and report malformed JSON or missing top-level fields.

Scope:
- internal/core/contract
- ID prefix helpers for CAP/REQ/DEC/RISK/EVAL/ART/OQ
- tests for valid and invalid contracts

Non-goals:
- Do not implement full readiness yet.
- Do not create lockfile.
- Do not compile prompts.
- Do not add JSON schema dependency unless clearly justified.

Validation:
- gofmt -w .
- go test ./...
- bash scripts/quality.sh
