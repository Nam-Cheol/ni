# Neighborhood Cooling Study Protocol

ni is not a software spec generator.
ni compiles project intent for any product surface.

This locked example plans a research protocol, not a software app. The product
is a documented field-study method for comparing street-level cooling
interventions before any fieldwork begins.

## What this proves

- `product_type` is `research_protocol`.
- The delivery surface is `document`, not web or CLI.
- Evaluation is protocol review: sampling reproducibility, ethics boundary
  review, and analysis reproducibility review.
- `ni run` compiles a handoff prompt for a research team. It does not collect
  data, run analysis, deploy sensors, or execute implementation.

## Included files

- `docs/plan/**`: locked planning docs for the protocol intent.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, and decisions.
- `.ni/plan.lock.json`: CLI-written lock with hashes for the authoritative
  planning files.
- `generated/human-team.prompt.md`: compiled human-team handoff prompt.
- `contract-summary.md`: compact summary of the locked contract.

## Try it

From the repository root:

```bash
go run ./cmd/ni status --dir examples/research-protocol
go run ./cmd/ni run --dir examples/research-protocol --target human-team --out examples/research-protocol/generated/human-team.prompt.md
```

Expected status: `READY`.

## Boundary

This example stops at intent lock and handoff. It does not implement the study,
collect participant data, create a dashboard, or replace ethics review.
