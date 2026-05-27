# Ambiguous prompt blocked

This example shows ni's core payoff:

> This is ni's core payoff: it blocks ambiguous execution before the agent starts.

A vague request like "build me a dashboard for my team" sounds actionable, but
it hides product decisions that an implementation agent would otherwise make
silently: who the dashboard serves, what data it may use, which workflow matters
first, what must not be built, and how success will be evaluated.

The example walks the same prompt through two paths:

1. Direct-to-agent execution begins with hidden assumptions.
2. `ni-start` captures the intent as planning records.
3. `ni status` returns `BLOCKED` because blocker questions remain open.
4. The user answers focused questions.
5. `ni end` can lock the clarified plan.
6. `ni run` compiles a bounded Codex target prompt.

The checked-in workspace under
[`workspace/`](workspace/) is intentionally the blocked state from step 4. It
contains synchronized `docs/plan/**` and `.ni/contract.json` records, with two
blocker open questions.

Verify the documented blocked state:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

Expected status:

```text
BLOCKED
blocker R009: OQ-001 is a blocker open question
blocker R009: OQ-002 is a blocker open question
```

## Files

- [`01-vague-request.md`](01-vague-request.md): the ambiguous user prompt.
- [`02-direct-to-agent-risk.md`](02-direct-to-agent-risk.md): what a direct
  implementation path would likely assume.
- [`03-ni-start-conversation.md`](03-ni-start-conversation.md): how `ni-start`
  records intent without pretending the plan is ready.
- [`04-ni-status-blocked.md`](04-ni-status-blocked.md): the deterministic
  blocked status.
- [`05-next-questions.md`](05-next-questions.md): focused questions that remove
  ambiguity.
- [`06-user-answers.md`](06-user-answers.md): user answers that turn assumptions
  into accepted planning records.
- [`07-locked-contract-summary.md`](07-locked-contract-summary.md): the ready
  plan summary after `ni end`.
- [`08-codex-target-prompt.md`](08-codex-target-prompt.md): a bounded prompt
  that `ni run` could hand to Codex after lock.

## Boundary

This demo does not execute Codex, implement the dashboard, start shell adapters,
create a queue, or add task-runner behavior. It is a kernel example: docs,
contract, readiness, lock narrative, and prompt compilation only.
