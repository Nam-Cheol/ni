# Conversation Authoring Fixture

## 1. Purpose

This fixture shows the sustained planning loop after `ni init`: a model and
user author planning docs through conversation while the CLI remains the
readiness, lock, and prompt compiler authority.

## 2. What this proves

- The user does not need contract authoring commands such as `contract add`,
  `contract set`, or `contract list`.
- A model may update `docs/plan/**` and `.ni/contract.json` together, then use
  `ni status` to verify readiness.
- `READY_WITH_DEFERRALS` can be locked when the deferrals are explicit and
  accepted.
- `ni run` compiles a human-team handoff prompt without executing downstream
  implementation.

## 3. Product type / surface

- `product_type`: `conversation_product`
- `delivery_surface`: `conversation`, `document`
- Expected `ni status`: `READY_WITH_DEFERRALS`
- Expected `ni run` target: `human-team`

## 4. Files

- `transcript.md`: model-user authoring loop and status checks.
- `ni-end-confirmation.md`: confirmation behavior before lock.
- `ni-run-handoff.md`: target selection, stale-lock refusal, and prompt
  compilation behavior.
- `session-resume.md`: bounded session resume below contract authority.
- `docs/plan/**`: completed human-facing plan docs.
- `.ni/contract.json`: matching machine-readable planning contract.
- `.ni/session.json`: bounded resume state below docs and contract.
- `.ni/plan.lock.json`: CLI-written lockfile for the completed plan.
- `generated/human-team.prompt.txt`: checked-in compiled handoff prompt.

## 5. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir examples/conversation-authoring
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/conversation-authoring --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.txt"
wc -m "$tmpdir/human-team.prompt.txt"
rm -rf "$tmpdir"
```

## 6. Expected output

Expected status: `READY_WITH_DEFERRALS`.

The status command should start with:

```text
READY_WITH_DEFERRALS
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document
```

It should also report the accepted deferrals:

```text
deferral D001: DEC-004 is deferred
deferral D002: OQ-002 remains open
```

The run command should write a non-empty prompt at or below 4000 characters.

## 7. Non-execution boundary

This fixture does not run a support assistant, contact customers, approve
refunds, call a model API, or invoke downstream tools. ni validates planning
state, writes the lock through `ni end`, and compiles a bounded prompt through
`ni run`.
