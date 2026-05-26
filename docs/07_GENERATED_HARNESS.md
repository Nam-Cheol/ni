# Generated harness

The generated harness is derived from a locked project contract.

The harness is not the kernel. It is a mutable proposal that selects part of the locked plan and describes how another tool or agent could execute it later.

It may contain:

```text
work graph
evaluation plan
evidence plan
implementation packets
review checklist
```

It must not change the locked planning contract unless a new planning cycle begins.
It must not execute work by itself.
It must not add adapters as part of the kernel.

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

## Machine-readable shape

The canonical proposal shape is JSON:

```text
schema
source_lock_hash
selected_capabilities
work_packets
validations
evidence_locations
known_risks
non_goals
```

`source_lock_hash` is the sha256 of `.ni/plan.lock.json`, prefixed with `sha256:`.

`selected_capabilities` is a list of `CAP-*` IDs from the locked contract.

`work_packets` are proposed units of work. They may include capability IDs, dependency packet IDs, artifact IDs, validation IDs, evidence IDs, and risk IDs. A work packet is descriptive only; it is not a queue item and must not be run by `ni`.

`validations` describe checks that should produce evidence. They are derived from evaluations in the planning contract.

`evidence_locations` describe where proof should be collected, such as artifact paths, command output, changed files, or generated reports.

`known_risks` carries risk IDs and mitigations from the contract.

`non_goals` preserves constraints that generated work must not violate.

## CLI behavior

`ni harness plan --dir <path>` may print a read-only proposal derived from the locked plan.

The command must:

- verify `.ni/plan.lock.json` exists,
- verify locked file hashes,
- refuse stale locks with `BLOCKED`,
- avoid modifying files,
- avoid executing work packets,
- avoid creating shell, Codex, or remote adapters.
