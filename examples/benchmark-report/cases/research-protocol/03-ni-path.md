# ni path

This benchmark case uses an isolated ni workspace:

```text
examples/benchmark-report/cases/research-protocol/workspace/
```

The workspace records the request as a `research_protocol` with `document`,
`workflow`, and `human_service` delivery surfaces. It intentionally keeps the
research blockers unresolved.

Expected path for this initial task:

1. Initialize the isolated workspace.
2. Record the vague research-protocol request in docs and `.ni/contract.json`.
3. Keep `OQ-001` through `OQ-005` open and blocking.
4. Run `ni status --proof --next-questions`.
5. If status is `BLOCKED`, stop.

Measured result:

- Readiness: `BLOCKED`
- Lock: no
- Bounded prompt: no
- Prompt count: `not_measured`

Stop rule: because status is `BLOCKED`, this task does not run `ni end` or
`ni run`.
