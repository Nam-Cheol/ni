# Domain and state model

## Core entities

```text
vague request
planning assumption
blocker open question
project contract
readiness result
lockfile
compiled prompt
downstream implementation agent
```

## State rule

The blocked workspace is pre-lock state. If the user answers the blocker
questions, the docs and contract may be updated together and checked again with
`ni status`.
