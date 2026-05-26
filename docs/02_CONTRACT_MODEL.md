# Contract model

The machine-readable project contract lives at `.ni/contract.json`.

## Required top-level fields

```text
schema
readiness_profile
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
