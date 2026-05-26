# Contract model

The machine-readable project contract lives at `.ni/contract.json`.

The published JSON Schema for this file is
`schema/ni.contract.v0.json`. Other NI state files have versioned schemas in
`schema/`:

```text
schema/ni.project.v0.json
schema/ni.contract.v0.json
schema/ni.lock.v0.json
schema/ni.readiness-rules.v0.json
schema/ni.readiness-profiles.v0.json
schema/ni.feedback.v0.json
schema/ni.pressure.v0.json
schema/ni.amendment.v0.json
schema/ni.harness-candidate.v0.json
```

`schema/ni.harness-candidate.v0.json` validates the
`.ni/harness.candidates.json` ledger whose in-file schema value is
`ni.harness_candidates.v0`.

## Required top-level fields

```text
schema
readiness_profile
product_type
delivery_surfaces
interaction_mode
project
non_goals
capabilities
requirements
decisions
risks
evaluations
artifacts
open_questions
```

`readiness_profile` selects how strict `ni status` is when converting readiness issues into blockers or deferrals. The valid values are:

```text
concept
prototype
mvp
beta
production
```

The default profile created by `ni init` is `prototype`.

## Product shape

`product_type`, `delivery_surfaces`, and `interaction_mode` keep the contract usable for products that are not web or software systems.

`product_type` uses one of these values:

```text
software
conversation_product
research_protocol
operations_process
education_program
document_product
physical_product
mixed
```

`delivery_surfaces` is a non-empty array using one or more of these values:

```text
web
cli
api
conversation
document
workflow
human_service
physical
```

`interaction_mode` is a lowercase identifier such as:

```text
human_to_system
human_to_human
system_to_system
hybrid
```

If an older contract has no product shape fields, the CLI treats that as no value supplied and defaults to:

```text
product_type: software
delivery_surfaces: [cli]
interaction_mode: human_to_system
```

`ni init` can set these fields explicitly:

```bash
ni init --product-type conversation_product --surface conversation
```

Product shape affects scaffold text and `ni status` guidance only. Because these fields live in the contract, their values are included in normal lock hashing, but they must not add product-specific readiness authority or `ni run` execution behavior.

## ID prefixes

```text
CAP-001   capability
REQ-001   requirement
DEC-001   decision
RISK-001  risk
EVAL-001  evaluation
ART-001   artifact
OQ-001    open question
```

## Status values

Common status values:

```text
draft
accepted
deferred
rejected
not_applicable
```

## Traceability

A capability should be connected to:

```text
requirements[]
evaluations[]
risks[]
artifacts[]
```

A work packet generated later should trace back to IDs from the locked contract.

## Schema checks

Run the schema check with:

```bash
python3 scripts/check-schema.py
```

The check validates the published schema files and the current `.ni` state
files that are present in the repository.
