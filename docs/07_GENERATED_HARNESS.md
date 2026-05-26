# Generated harness

The generated harness is derived from a locked project contract.

It may contain:

```text
work graph
evaluation plan
evidence plan
adapter choice
implementation packets
review checklist
```

It must not change the locked planning contract unless a new planning cycle begins.

## Work graph instead of sequence

`ni` should avoid forcing all work into a total order.

```text
CAP-001 -> EVAL-001
CAP-002 -> EVAL-002
CAP-003 depends_on CAP-001
```

Independent nodes may be executed in parallel by different people or agents. Nodes that touch the same artifact or decision should be serialized.

## Generated harness contract

A generated harness must declare:

```text
source lock hash
selected capabilities
work packets
validation commands
evidence locations
known risks
non-goals
```
