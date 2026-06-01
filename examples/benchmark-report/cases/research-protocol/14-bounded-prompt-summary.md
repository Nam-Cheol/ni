# Bounded prompt summary

Command:

```bash
go run ./cmd/ni run --dir examples/benchmark-report/cases/research-protocol/workspace --max-chars 4000
```

Measurement command:

```bash
go run ./cmd/ni run --dir examples/benchmark-report/cases/research-protocol/workspace --max-chars 4000 | wc -m
```

Prompt target: `generic`

Prompt character count: `4000`

The prompt was bounded by `--max-chars 4000` and ended with:

```text
[truncated to max-chars]
```

Interpretation: `ni run` compiled prompt seed material only. It did not
execute Codex, shell commands, downstream agents, model API calls, adapters,
queues, PR automation, release automation, fieldwork, participant recruitment,
data collection, dashboard work, or research implementation.
