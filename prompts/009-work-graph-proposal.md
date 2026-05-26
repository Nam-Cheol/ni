# 009 - Work graph proposal

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Implement a read-only work graph proposal from `.ni/contract.json`.

Goal:
Create a command that proposes a dependency graph without executing tasks.

Suggested command:
`ni graph --dir <path>`

Scope:
- internal/core/graph
- graph nodes from capabilities and artifacts
- dependency field support if present
- text output and optional JSON output

Non-goals:
- Do not run work packets.
- Do not create shell/Codex adapter.
- Do not mutate locked docs.

Validation:
- gofmt -w .
- go test ./...
- go run ./cmd/ni graph --dir .
- bash scripts/quality.sh
