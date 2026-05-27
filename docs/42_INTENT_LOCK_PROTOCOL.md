# Intent Lock Protocol

The Intent Lock Protocol is a deterministic pre-runtime control layer that
defines:

1. how planning conversations become a project contract,
2. when the contract is ready to lock,
3. how the accepted plan is hashed,
4. what downstream actors may trust,
5. when execution must stop because intent changed.

It is the core mechanism of `ni-kernel`. `ni` does not merely write planning
docs; it compiles conversation into a locked, verifiable intent contract before
any downstream execution begins.

## Problem

Agents execute ambiguous intent too early.

Most agent systems start controlling behavior after a prompt, spec, work queue,
or runtime loop already exists. At that point the agent may be executing from
unclear goals, unresolved decisions, weak acceptance criteria, unmitigated
risks, or stale planning context.

The protocol moves control earlier. It asks whether the project intent is
explicit, accepted, validated, and unchanged before any downstream actor starts
work.

## Protocol inputs

The protocol reads planning state from:

- planning conversation,
- `docs/plan/**`,
- `.ni/contract.json`,
- `.ni/session.json`.

Conversation is source material. `docs/plan/**` and `.ni/contract.json` are the
model-maintained planning record. `.ni/session.json` is continuity state only;
it may help resume authoring, but it does not override the contract, docs, or
lock.

After `.ni/plan.lock.json` exists, source-of-truth precedence is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

## Protocol gates

The protocol is enforced through deterministic gates:

- `ni status` validates whether the contract is `BLOCKED`,
  `READY_WITH_DEFERRALS`, or `READY`.
- `ni end` may lock only a plan that passes the readiness gate.
- `.ni/plan.lock.json` records hashes for the accepted contract and required
  planning docs.
- stale lock refusal stops prompt compilation, seed export, feedback flows, and
  downstream handoff when current files no longer match the lock.

Models and skills may draft, explain, and guide. They do not declare readiness,
create locks, repair locks, or override stale-lock failures.

## Protocol outputs

When the gates pass, the protocol can produce:

- a locked project contract,
- a target prompt,
- seed export material.

The locked project contract is the authoritative intent snapshot. A target
prompt is a bounded downstream handoff compiled from the valid lock. Seed export
material is derived and mutable; it may help another tool start, but it does not
become kernel-owned execution state.

## Protocol rule

Downstream execution must not start from ambiguous or stale intent.

If readiness is `BLOCKED`, execution has not earned a trustworthy starting
point. If `.ni/plan.lock.json` is missing where a lock is required, execution
has no accepted intent snapshot. If locked hashes no longer match, intent has
changed and downstream execution must stop until the plan is amended and
relocked.

## Non-goals

The Intent Lock Protocol does not add:

- runtime execution,
- model API calls,
- shell or Codex adapters,
- contract authoring commands,
- queues or task-runner behavior.

The protocol ends at verified intent handoff. Runtime systems may consume the
locked contract, target prompt, or seed export, but they remain downstream of
`ni-kernel`.
