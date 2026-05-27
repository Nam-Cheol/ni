# 09. ni-run handoff

## User

Invoke `ni-run` for the human team.

## Model using ni-run

I will let the CLI verify the lock hashes and compile the target prompt.

```bash
ni run --dir ./refund-triage-plan \
  --target human-team \
  --max-chars 4000 \
  --out generated/human-team.prompt.txt
```

## CLI result

```text
compiled prompt at generated/human-team.prompt.txt
```

## Handoff boundary

The compiled prompt is downstream-compatible seed material only. It does not
start a support assistant, call a model API, open a shell, run Codex, approve
refunds, contact customers, or create a runtime.

The checked-in equivalent for this example is:

```bash
go run ./cmd/ni run --dir examples/ni-start-dogfood/workspace \
  --target human-team \
  --max-chars 4000 \
  --out /tmp/human-team.prompt.txt
```
