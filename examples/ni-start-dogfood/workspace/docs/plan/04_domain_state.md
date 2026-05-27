# Domain state

## Inputs

- Support ticket facts supplied to the support agent.
- Internal refund policy page, selected in conversation as the authoritative
  eligibility source.
- Transcript fixtures for eligible, ineligible, and ambiguous refund cases.

## State categories

- Eligible recommendation: ticket facts match policy and no exception is
  required.
- Ineligible recommendation: ticket facts do not satisfy the policy.
- Escalation: policy is ambiguous, stale, missing, or conflicts with ticket
  facts.

## Source of truth

The policy page is authoritative for eligibility. The NI source-of-truth order
for the authored plan remains `.ni/plan.lock.json`, `.ni/contract.json`,
`docs/plan/**`, `.ni/session.json`, then chat history.
