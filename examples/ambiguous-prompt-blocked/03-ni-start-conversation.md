# 03. ni-start conversation

`ni-start` does not execute the dashboard request. It extracts planning state
from the conversation and writes draft records to `docs/plan/**` plus
`.ni/contract.json`.

Captured facts:

- Product type: software.
- Surface: web.
- Initial purpose: plan a customer-issue dashboard before implementation.
- Non-goal: do not connect live customer systems or build downstream automation.

Draft assumptions:

- The primary user is probably a support lead.
- The first version may use static or imported issue data.
- The initial workflow may be triage visibility, not ticket mutation.

Blocker questions:

- OQ-001: Which team and first workflow must the dashboard serve?
- OQ-002: What data source and data sensitivity boundary may the first version use?

The example workspace at `workspace/` represents this captured-but-blocked
state.
