# 010 - Generated harness contract

General rules for this task:
- Read AGENTS.md first.
- Keep the ni-kernel boundary intact.
- Do not add execution adapters unless this prompt explicitly asks for it.
- Before editing, summarize the intended changes and exact files you expect to modify.
- After editing, show changed files and validation results.

Task:
Define the generated harness contract.

Goal:
Create a machine-readable proposal format for project-specific harnesses derived from a locked plan.

Scope:
- docs/07_GENERATED_HARNESS.md
- .ni/harness.schema.example.json or docs/examples/generated-harness.json
- optional `ni harness plan` command if small and read-only

Required fields:
- source lock hash
- selected capabilities
- work packets
- validations
- evidence locations
- known risks
- non-goals

Non-goals:
- Do not execute the harness.
- Do not create adapters.

Validation:
- gofmt -w . if Go changed
- go test ./... if Go exists
- bash scripts/quality.sh
