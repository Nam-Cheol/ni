# Conversation Authoring Fixture

## 1. Purpose

This fixture shows ni as a Project Intent Compiler during the sustained
planning loop after `ni init`: a model and user author planning docs through
conversation while the CLI remains the readiness, lock, and prompt compiler
authority.

It also shows that historical locked fixture material is not enough to claim a
fresh ready state. The current planning docs and contract must still pass
`ni status`.

## 2. What this proves

- The user does not need contract authoring commands such as `contract add`,
  `contract set`, or `contract list`.
- A model may update `docs/plan/**` and `.ni/contract.json` together, then use
  `ni status` to verify readiness.
- The model should consume grouped `ni status --proof --next-questions` output
  as the next planning interview instead of inventing broad questions.
- `ni status` can catch stale docs/contract synchronization even when a
  historical lockfile and generated prompt are present.
- The checked-in `ni run` material remains inert handoff seed material; it is
  not downstream execution.

## 3. Product type / surface

- `product_type`: `conversation_product`
- `delivery_surface`: `conversation`, `document`
- Expected `ni status`: `BLOCKED`
- Expected `ni run` target: `human-team` from the existing lock only

## 4. Files

- `transcript.md`: model-user authoring loop and status checks.
- `ni-end-confirmation.md`: confirmation behavior before lock.
- `ni-run-handoff.md`: target selection, stale-lock refusal, and prompt
  compilation behavior.
- `session-resume.md`: bounded session resume below contract authority.
- `docs/plan/**`: completed human-facing plan docs.
- `.ni/contract.json`: matching machine-readable planning contract.
- `.ni/session.json`: bounded resume state below docs and contract.
- `.ni/plan.lock.json`: historical CLI-written lockfile for the fixture.
- `generated/human-team.prompt.txt`: checked-in compiled handoff prompt from
  the existing lock.

## 5. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir examples/conversation-authoring
go run ./cmd/ni status --dir examples/conversation-authoring --proof --next-questions
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/conversation-authoring --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.txt"
wc -m "$tmpdir/human-team.prompt.txt"
rm -rf "$tmpdir"
```

## 6. Expected output

Expected status: `BLOCKED`.

The status command should start with:

```text
BLOCKED
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document
```

It should also keep the docs/contract synchronization blockers visible:

```text
blocker R012
```

The run command may compile from the existing lockfile, but the fixture must
not be described as freshly ready until `ni status` passes again.

## 7. demo-check coverage

Covered by `bash scripts/demo-check.sh`.

The demo check verifies the current `BLOCKED` status and sync blocker, and it
compiles the `human-team` prompt only from the existing lock.

## 8. Korean companion

Korean companion docs exist: `README.ko.md`.

## 9. Non-execution boundary

This fixture does not run a support assistant, contact customers, approve
refunds, call a model API, or invoke downstream tools. Do not call `ni end` for
this fixture while `ni status` reports `BLOCKED`; ni only validates planning
state and compiles bounded prompt material from an existing lock.
