# 006 - Lockfile end

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Implement `ni end`.

Goal:
Create `.ni/plan.lock.json` only when `ni status` is not `BLOCKED`.

Scope:
- internal/core/lock
- hash calculation for `.ni/contract.json` and docs/plan files
- `ni end --dir <path>`
- tests for successful lock and blocked refusal

Required behavior:
- Refuse lock on BLOCKED.
- Write schema `ni.lock.v0`.
- Include lock timestamp.
- Include source_of_truth.
- Include sha256 for contract and planning docs.

Non-goals:
- Do not compile prompt yet.
- Do not execute any adapter.

Validation:
- gofmt -w .
- go test ./...
- go run ./cmd/ni end --dir <ready-fixture-or-temp-copy>
- bash scripts/quality.sh
