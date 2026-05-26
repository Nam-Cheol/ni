# 007 - Prompt compiler run

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Implement `ni run` as a prompt compiler.

Goal:
Read `.ni/plan.lock.json`, verify hashes, and print a 4000-character-or-less goal prompt.

Scope:
- internal/core/prompt
- lock hash verification
- `ni run --dir <path>`
- `ni run --out <path>`
- `ni run --max-chars 4000`
- tests for prompt budget and stale lock refusal

Required behavior:
- Refuse if lockfile is missing.
- Refuse if any locked file hash mismatches.
- Do not execute Codex or shell.
- Do not paste full docs into the prompt.

Validation:
- gofmt -w .
- go test ./...
- go run ./cmd/ni run --dir . --max-chars 4000
- bash scripts/quality.sh
