# Execution strategy

## v0 execution strategy

Do not execute implementation automatically. Use `ni run` to compile a short prompt after the plan is locked.

## SDD relation

The plan borrows SDD clarity around boundaries and acceptance references, but rejects a mandatory sequential SPEC execution chain. Boundary candidates should form a graph:

```text
discover current namba-ai facts -> confirm upgrade boundaries
confirm upgrade boundaries -> propose implementation packets
confirm upgrade boundaries -> propose validation evidence
propose implementation packets -> downstream execution outside NI
```

## No execution tasks

This example does not create tasks, queues, tickets, PRs, shell commands, or adapter invocations. The only generated artifact required by the task is a prompt file saved under `generated/`.
