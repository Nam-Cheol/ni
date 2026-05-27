# How ni Works

`ni-kernel` is the authoritative layer for turning planning conversation into
a locked project contract. Downstream seed material is derived and mutable.

```text
conversation -> docs/plan/** + .ni/contract.json -> ni status -> ni end -> ni run
```

## Kernel Ownership

The kernel is authoritative for:

- `docs/plan/**`;
- `.ni/contract.json`;
- deterministic readiness validation;
- `.ni/plan.lock.json`;
- lock hash verification;
- bounded prompt compilation;
- inert downstream seed exports and proposals.

Downstream seed material may include prompts, work-graph proposals,
evaluation-plan proposals, evidence-rule notes, harness seed proposals, or
handoff packets. It must not become kernel-owned execution state.

## Source Of Truth

Before a lock exists, `docs/plan/**` and `.ni/contract.json` are maintained
together from planning conversation. Skills and models are UX; the CLI is
authority.

After `.ni/plan.lock.json` exists, source-of-truth precedence is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

If locked hashes no longer match current planning files, target handoff stops
with `BLOCKED`.

## The Core Gates

`ni status` is the readiness gate. It validates whether the contract is
`BLOCKED`, `READY_WITH_DEFERRALS`, or `READY`. Model judgment cannot override
this result.

`ni end` is the lock gate. It may write `.ni/plan.lock.json` only after the
readiness gate passes. The lock hashes `.ni/contract.json` and required
`docs/plan/**` files.

`ni run` is the prompt compiler. It requires a valid lock, verifies locked
hashes, and prints or writes a bounded target prompt. It does not execute that
prompt.

## Commands

The initial core path is:

```text
ni init -> ni status -> ni end -> ni run
```

Other implemented kernel commands inspect targets, export locked seed material,
record inert feedback and pressure, manage explicit amendments, compare
planning states, and propose inert graph or downstream tool material.

See the full [Command Reference](commands.md).

## Target Boundaries

Targets are consumption shapes for a locked plan. They are not integrations
that `ni` executes, adapters that `ni` owns, or lifecycle state that becomes
part of `ni-kernel`.

Built-in targets include prompt and handoff targets for coding models, model
workspaces, and human teams, plus inert seed targets for downstream tools and
execution environments.

See [Target Story](45_TARGET_STORY.md) for target-by-target boundaries and
named downstream examples.

## Change After Lock

Locked planning docs must not be silently edited. If intent changes after a
lock, the change needs an explicit amendment or relock flow. A stale lock is a
hard stop for handoff, because downstream actors can no longer trust that the
compiled prompt or seed material matches accepted intent.

See [Intent Lock Protocol](intent-lock-protocol.md) for the public protocol
overview and [Protocol Specification](42_INTENT_LOCK_PROTOCOL.md) for the
deeper mechanism.
