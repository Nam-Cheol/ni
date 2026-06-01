# Execution strategy

## v0 execution strategy

Do not execute implementation, fieldwork, participant recruitment, data
collection, analysis, intervention placement, downstream agents, generated
prompts, model APIs, shell adapters, queues, telemetry paths, release
automation, or runtime harnesses.

For this initial benchmark measurement, run only:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions
```

If status is `BLOCKED`, stop. Do not run `ni end` or `ni run`.
