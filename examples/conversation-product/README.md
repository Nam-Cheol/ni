# Travel Concierge Triage

ni is not a software spec generator.
ni compiles project intent for any product surface.

This locked example plans a conversation product, not an app. The product is a
human-operated travel intake conversation that gathers enough trip intent,
constraints, risks, and open questions for a human concierge to continue.

## What this proves

- `product_type` is `conversation_product`.
- The delivery surface is `conversation`, not web or CLI.
- Evaluation is transcript and checklist based: intake coverage, unsupported
  request escalation, and human concierge handoff quality.
- `ni run` compiles a handoff prompt. It does not deploy an assistant, contact
  vendors, make bookings, or execute implementation.

## Included files

- `docs/plan/**`: locked planning docs for the conversation intent.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, and decisions.
- `.ni/plan.lock.json`: CLI-written lock with hashes for the authoritative
  planning files.
- `generated/human-team.prompt.md`: compiled human-team handoff prompt.
- `contract-summary.md`: compact summary of the locked contract.

## Try it

From the repository root:

```bash
go run ./cmd/ni status --dir examples/conversation-product
go run ./cmd/ni run --dir examples/conversation-product --target human-team --out examples/conversation-product/generated/human-team.prompt.md
```

Expected status: `READY`.

## Boundary

This example stops at intent lock and handoff. It does not implement a chatbot,
add a queue, integrate live support, book travel, or make regulated advice
claims.
