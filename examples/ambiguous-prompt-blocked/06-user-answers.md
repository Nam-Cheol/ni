# 06. User answers

User:

> This is for the support lead. v0 should show which open issues need attention
> before the daily standup. Use a CSV export from our support tool, not live API
> access. The dashboard should be read-only. It should group issues by priority,
> owner, and age bucket, and flag anything older than 48 hours. Do not assign
> tickets or send notifications yet.

Planning update:

- OQ-001 can resolve to an accepted decision: support-lead daily triage is the
  first workflow.
- OQ-002 can resolve to an accepted decision: v0 uses a CSV export and must not
  connect to live systems.
- Requirements can be tightened around read-only grouping, age buckets, and
  overdue flags.
- Evaluation can check a fixed CSV fixture against expected dashboard counts.
- Non-goals remain explicit: no live API, ticket assignment, notifications, or
  downstream execution.

The model may now update docs and contract, then ask the CLI again. The model
still does not declare readiness by itself.
