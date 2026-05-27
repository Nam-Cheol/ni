# ni-start Dogfood Transcript

## 1. Purpose

This public example shows the intended authoring loop after `ni init`: the
user plans through conversation, the model updates `docs/plan/**` and
`.ni/contract.json` together, and the CLI validates readiness before anything
locks or compiles.

## 2. What this proves

- The primary authoring UX is sustained planning conversation, not user-entered
  contract `add`, `list`, or `set` commands.
- The model may summarize, ask focused questions, and update docs plus contract
  records from confirmed answers.
- `ni status` controls readiness; the model never overrides a `BLOCKED` result.
- `ni end` writes the lock only after CLI readiness and explicit user
  confirmation.
- `ni run` compiles a bounded handoff prompt only; it does not execute
  downstream work.

## 3. Product type / surface

- `product_type`: `conversation_product`
- `delivery_surface`: `conversation`, `document`
- Expected `ni status`: `READY_WITH_DEFERRALS`
- Expected `ni run` target: `human-team`

## 4. Transcript outline

- `01-init.md`: `ni init` creates the planning workspace.
- `02-user-vague-idea.md`: the user gives a vague support-assistant idea.
- `03-model-summary-and-questions.md`: `ni-start` summarizes gaps and asks
  focused questions.
- `04-user-answers.md`: the user confirms scope, non-goals, and evidence.
- `05-docs-contract-delta.md`: the model updates docs and contract together.
- `06-status-proof.md`: `ni status` reports readiness state and blockers.
- `07-second-round-questions.md`: the model asks only the next blocker
  question from the gate.
- `08-ni-end-confirmation.md`: `ni-end` confirms CLI readiness before locking.
- `09-ni-run-handoff.md`: `ni-run` compiles a handoff prompt from the lock.

## 5. Files

- `workspace/docs/plan/**`: small completed planning workspace.
- `workspace/.ni/contract.json`: matching machine-readable contract.
- `workspace/.ni/session.json`: bounded continuity state below contract
  authority.
- `workspace/.ni/plan.lock.json`: CLI-written lockfile copied from the dogfood
  plan state.
- `workspace/generated/human-team.prompt.txt`: compiled prompt seed material.

## 6. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/ni-start-dogfood/workspace --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.txt"
wc -m "$tmpdir/human-team.prompt.txt"
rm -rf "$tmpdir"
```

## 7. Expected output

Expected status: `READY_WITH_DEFERRALS`.

The status command should start with:

```text
READY_WITH_DEFERRALS
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document
```

It should keep accepted deferrals visible:

```text
deferral D001: DEC-004 is deferred
deferral D002: OQ-002 remains open
```

## 8. Non-execution boundary

This example does not run a support assistant, call a model API, start Codex,
contact customers, approve refunds, create adapters, or manage a queue. It is a
kernel example for conversation-first authoring, readiness proof, lock
authority, and prompt compilation.
