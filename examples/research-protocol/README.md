# Neighborhood Cooling Study Protocol

## 1. Purpose

This locked example plans a research protocol, not a software app. The product
is a documented field-study method for comparing street-level cooling
interventions before any fieldwork begins.

## 2. What this proves

- ni can lock intent for a non-software product.
- `product_type=research_protocol` changes planning guidance without changing
  the shared readiness gate.
- The delivery surface can be a document.
- `ni run` can compile a bounded handoff prompt for a research team without
  collecting data or running analysis.

## 3. Product type / surface

- `product_type`: `research_protocol`
- `delivery_surface`: `document`
- Expected `ni status`: `READY`
- Expected `ni run` target: `human-team`

## 4. Files

- `docs/plan/**`: locked planning docs for the protocol intent.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, and decisions.
- `.ni/plan.lock.json`: CLI-written lock with hashes for authoritative
  planning files.
- `generated/human-team.prompt.md`: checked-in compiled human-team handoff.
- `generated/generic.prompt.txt`: checked-in generic downstream handoff.
- `contract-summary.md`: compact summary of the locked contract.

## 5. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir examples/research-protocol
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/research-protocol --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.md"
wc -m "$tmpdir/human-team.prompt.md"
rm -rf "$tmpdir"
```

## 6. Expected output

Expected status: `READY`.

The status command should start with:

```text
READY
profile: prototype
product type: research_protocol
delivery surfaces: document
```

The run command should write a non-empty prompt at or below 4000 characters.

## 7. Non-execution boundary

This example does not implement the study, collect participant data, deploy
sensors, run statistics, call a model API, or replace ethics review. ni only
validates the locked planning contract and compiles inert prompt material.
