# Domain and state model

## Core entities

```text
project
contract
capability
requirement
decision
risk
evaluation
artifact
open question
lockfile
prompt
```

## State transition

```text
DRAFT -> READY_WITH_DEFERRALS -> READY -> LOCKED -> PROMPT_COMPILED
```

`BLOCKED` is a validation result, not a stable project state.
