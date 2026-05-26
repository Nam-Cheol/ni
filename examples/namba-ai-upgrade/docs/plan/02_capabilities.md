# Capabilities

## CAP-001: Capture current namba-ai upgrade limitations as planning evidence

The plan records limitations that affect an upgrade before implementation starts: runtime coupling, implicit acceptance criteria, sequence pressure, unclear evidence ownership, and Codex-source-of-truth drift.

## CAP-002: Model SDD sequencing and collaboration issues as graph-oriented planning constraints

The upgrade plan must not require a single total-order SPEC chain. It should describe boundaries as graph nodes with:

```text
boundary candidate
depends_on
acceptance_refs
validation_refs
risk_refs
artifact_refs
```

Independent boundaries may move in parallel when they do not change the same decisions, requirements, risks, or artifacts. Conflicting edits must be resolved in planning before downstream implementation.

## CAP-003: Preserve a Codex-limited, pre-runtime planning compiler boundary

NI may compile a Codex target prompt after the plan is locked. NI must not invoke Codex, retain Codex execution state, or treat a Codex transcript as more authoritative than the lockfile and contract.

## CAP-004: Define downstream targets and validation expectations without execution tasks

Downstream targets are handoff destinations, not NI-owned queues:

- Codex prompt handoff,
- namba-ai seed guidance,
- human-team review,
- generated harness proposal.
