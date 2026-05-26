# Contract model

The machine-readable project contract lives at `.ni/contract.json`.

## Required top-level fields

```text
schema
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
