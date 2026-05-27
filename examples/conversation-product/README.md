# Travel Concierge Triage

## 1. Purpose

This locked example plans a conversation product, not an app. The product is a
human-operated travel intake conversation that gathers trip intent,
constraints, risks, and open questions for a human concierge.

## 2. What this proves

- ni can lock intent for a conversation product before any runtime exists.
- `product_type=conversation_product` changes planning guidance without turning
  ni into a chatbot runner.
- The delivery surface can be conversation-first.
- `ni run` can compile a bounded handoff prompt for a human team.

## 3. Product type / surface

- `product_type`: `conversation_product`
- `delivery_surface`: `conversation`
- Expected `ni status`: `READY`
- Expected `ni run` target: `human-team`

## 4. Files

- `docs/plan/**`: locked planning docs for the conversation intent.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, and decisions.
- `.ni/plan.lock.json`: CLI-written lock with hashes for authoritative
  planning files.
- `generated/human-team.prompt.md`: checked-in compiled human-team handoff.
- `generated/codex.prompt.txt`: checked-in Codex target prompt.
- `contract-summary.md`: compact summary of the locked contract.

## 5. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir examples/conversation-product
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/conversation-product --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.md"
wc -m "$tmpdir/human-team.prompt.md"
rm -rf "$tmpdir"
```

## 6. Expected output

Expected status: `READY`.

The status command should start with:

```text
READY
profile: prototype
product type: conversation_product
delivery surfaces: conversation
```

The run command should write a non-empty prompt at or below 4000 characters.

## 7. Non-execution boundary

This example does not implement a chatbot, deploy a service, add a queue,
contact vendors, book travel, call a model API, or make regulated advice
claims. ni only validates the locked contract and compiles prompt material.
