# Checklists

## Before implementation starts

```text
[ ] docs/plan exists
[ ] .ni/contract.json exists
[ ] ni status returns READY or READY_WITH_DEFERRALS
[ ] .ni/plan.lock.json exists
[ ] lock hashes match current files
[ ] ni run output is 4000 characters or less
```

## Before accepting a model-generated plan

```text
[ ] capabilities have evaluations
[ ] high risks have mitigations
[ ] blocker questions are resolved
[ ] non-goals are explicit
[ ] source-of-truth order is preserved
[ ] no acceptance criterion was weakened to pass
```

## Before adding execution adapters

```text
[ ] ni status is implemented
[ ] ni end is implemented
[ ] ni run prompt compiler is implemented
[ ] stale lock detection is implemented
[ ] test fixtures cover blocked cases
```
