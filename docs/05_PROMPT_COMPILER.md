# Prompt compiler

`ni run` compiles a goal prompt from the locked project contract.

In v0 it does not run Codex, shell, or any adapter.

## Required constraints

```text
1. output must be 4000 characters or less
2. output must reference authoritative files
3. output must not paste full docs into the prompt
4. output must include source-of-truth rules
5. output must instruct the agent to stop on lock mismatch
6. output must require work items to trace to CAP/REQ/EVAL/RISK IDs
7. output must require validation before implementation
```

## Prompt shape

```text
You are executing a locked NI project plan.

Authoritative sources:
- .ni/plan.lock.json
- .ni/contract.json
- docs/plan/

Rules:
- Do not modify locked planning docs.
- Every task must map to CAP/REQ/EVAL/RISK IDs.
- Do not weaken acceptance criteria.
- If files do not match the lock hash, stop and report BLOCKED.

Goal:
Build the smallest valid next product increment from the locked plan.

Process:
1. Read lockfile and contract.
2. Produce a work graph, not a linear SPEC chain.
3. Select the smallest safe work packet.
4. Define validation before editing.
5. Implement only that packet.
6. Run validation.
7. Report evidence and next packets.
```
