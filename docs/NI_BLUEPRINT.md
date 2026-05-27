# NI blueprint

## Definition

`ni` is the Project Intent Compiler for AI Agents.

Don't run the agent yet. Compile the intent first.

It compiles planning conversations and planning documents into a locked,
versioned, verifiable project contract before Codex, Claude, Spec Kit, Hyper
Run, namba-ai, a generated harness, or a human team starts execution.

Lock intent before any harness runs.

## Why this exists

Traditional harnesses try to stabilize model behavior by controlling the
external environment, retrying failures, and enforcing execution loops. That is
useful, but it still begins too late. It assumes the plan is already correct
enough to execute.

`ni` moves the control boundary earlier. It asks whether the project intent, constraints, risks, evaluations, and delivery model are explicit enough to start execution at all.

## Kernel responsibilities

The kernel owns:

```text
1. document templates
2. contract schema
3. readiness rules
4. lockfile creation
5. hash validation
6. prompt compilation
7. source-of-truth precedence
```

The kernel does not own:

```text
1. provider-specific agent execution
2. shell command execution
3. PR automation
4. multi-agent orchestration
5. project-specific implementation strategy
6. Codex adapter behavior
7. task, SPEC, or evidence-runner lifecycle state
```

## Conceptual pipeline

```text
ni-start
  Continue planning conversation.
  Update docs/plan and .ni/contract.json.
  Show readiness gaps.

ni-status
  Deterministically validate the contract.
  Output BLOCKED, READY_WITH_DEFERRALS, or READY.

ni-end
  If not blocked, lock the current plan snapshot.
  Write .ni/plan.lock.json.

ni-run
  Compile a short goal prompt from the locked contract.
  Do not execute in v0.
```

## Core entities

```text
CAP     capability
REQ     requirement
DEC     decision
RISK    risk
EVAL    evaluation
ART     artifact
OQ      open question
```

Every accepted capability must connect to at least one evaluation. High-severity risks must have mitigation. Blocker open questions must prevent locking.

## Design stance

The model may produce plans. The CLI decides whether the plan is structurally ready. The user decides whether the ready plan is accepted.
