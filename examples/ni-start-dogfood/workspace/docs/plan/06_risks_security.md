# Risks and security

## RISK-001: Refund authority may be implied

Severity: high
Status: accepted
Mitigation: Require recommendation language, agent review, and transcript checks for approval or customer-contact wording.

The assistant must not sound like it can approve or issue refunds.

## RISK-002: Ticket context may contain private customer data

Severity: medium
Status: accepted
Mitigation: Use only ticket facts needed for triage and keep generated handoffs inside the support-agent workflow.

Transcript fixtures should avoid unnecessary personal data.

## RISK-003: Policy source may be stale or unclear

Severity: high
Status: accepted
Mitigation: Cite policy, escalate conflicts or unclear sections, and keep policy ownership visibly deferred.

The plan treats unresolved ownership as a deferral, not as a hidden assumption.

## RISK-004: Docs and contract records may drift

Severity: medium
Status: accepted
Mitigation: Name changed docs and affected IDs each authoring turn, then run ni status.

This risk is specific to conversation-driven authoring.
