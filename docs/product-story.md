# Why ni exists

AI agents are getting better at doing work. The harder problem is still knowing
whether the work should start.

Many agent failures begin before the first file changes. A user gives a goal
that sounds clear enough to act on, the agent finds a plausible path, and the
system starts turning ambiguity into implementation. Missing users, unstated
acceptance criteria, hidden risks, non-goals, unresolved questions, and stale
planning context all become part of the work anyway.

That is the moment `ni` is built for.

## The execution problem

Agents often execute ambiguous intent because the interface invites it. A
prompt asks for momentum. The model is rewarded for being helpful. The user may
not know which details are structurally important until the work is already in
motion.

Better prompts help, but prompt refinement alone is not enough. A polished
prompt can still hide the same unresolved intent. It can sound decisive while
leaving the project without accepted constraints, evaluated capabilities,
explicit non-goals, mitigated high-severity risks, or a clear answer to whether
the plan changed after agreement.

The failure is not only wording. It is control.

## Move the boundary

`ni` moves the control boundary before execution.

Instead of asking an agent to infer the plan while it works, `ni` asks the
planning conversation to become a project contract first. The contract is not a
longer prompt. It is the shared record of what the project is for, what must be
true, what is out of scope, what risks matter, what evidence will count, and
which open questions still block a handoff.

When the contract is not ready, execution should not begin. That refusal is a
feature: it keeps uncertainty visible while it is still cheap to resolve.

## From conversation to handoff

The product shape is intentionally small:

```text
conversation -> contract -> lock -> bounded handoff
```

Conversation becomes contract when the planning record is explicit enough to
validate.

Contract becomes lock when the accepted plan passes the readiness gate and is
hashed into a source of truth.

Lock becomes bounded handoff when `ni` compiles a short prompt or inert seed
material from the locked plan.

Downstream work starts only after intent is trustworthy. If the plan changes
after lock, or the current files no longer match the locked hashes, the handoff
stops instead of pretending the old agreement still holds.

## Why this matters

AI work is fast enough that the cost of unclear intent shows up quickly:
unwanted scope, confident wrong turns, brittle acceptance, and arguments about
what was "really" requested.

`ni` does not try to make downstream agents smarter by adding another execution
loop. It makes their starting line safer. The goal is not more ceremony; it is
a better moment to say yes, no, or not yet.

When intent is still ambiguous, `ni` should make that obvious. When intent is
accepted and locked, downstream actors get a bounded artifact they can trust.
That is the product promise: compile the intent first, then let the work begin.

## Related work

For comparisons with adjacent agent, specification, and execution-harness
projects, see [Related work](11_RELATED_WORK.md) and the
[differentiation map](41_DIFFERENTIATION.md). For the technical protocol behind
this story, see the [Intent Lock Protocol](42_INTENT_LOCK_PROTOCOL.md).
