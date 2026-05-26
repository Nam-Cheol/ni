# Evaluation contract

## EVAL-001: Refund recommendation transcript fixtures

Method: Review eligible and ineligible refund transcript fixtures for correct recommendation, cited facts, cited policy source, and absence of refund issuance or customer-contact language.

The eligible fixture should produce a recommendation draft. The ineligible
fixture should explain the policy reason without contacting the customer.

## EVAL-002: Ambiguous policy escalation transcript fixture

Method: Review a policy ambiguity or ticket-policy conflict fixture and verify supervisor escalation with no invented policy and a clear next action for the support agent.

The ambiguity fixture proves that unresolved policy cases remain human-owned.

## EVAL-003: Conversation authoring audit fixture

Method: Check the transcript, docs/plan sections, and .ni/contract.json for matching CAP, REQ, EVAL, RISK, ART, NG, DEC, and OQ records after the model-authored turn.

This evaluation proves the authoring loop, not the support assistant runtime.
