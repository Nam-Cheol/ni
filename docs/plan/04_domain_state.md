# Domain and state model

## Core entities

```text
project
contract
readiness profile
product type
delivery surface
interaction mode
session state
capability
requirement
decision
risk
evaluation
artifact
open question
target
feedback
pressure item
harness candidate
amendment
collaboration diff
collaboration conflict
lockfile
prompt
export seed
```

## State transition

```text
DRAFT -> READY_WITH_DEFERRALS -> READY -> LOCKED -> PROMPT_COMPILED
LOCKED -> AMENDMENT_DRAFT -> AMENDMENT_APPLIED -> RELOCKED
```

`BLOCKED` is a validation result, not a stable project state.

In v0.2, draft planning state is normally produced by model-user conversation after `ni init`, then persisted into `docs/plan/**` and `.ni/contract.json` by `ni-start`. The CLI does not become an interactive contract editor.

## Source-of-truth state

After `.ni/plan.lock.json` exists, source-of-truth precedence is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat transcript
```

## Derived state

Feedback, pressure, harness candidates, prompts, and target exports are derived or inert. They may point to evidence and pressure, but they do not become accepted planning state until an explicit amendment is applied and the plan is relocked.
