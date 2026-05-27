# 07. Locked contract summary

After the user answers, `ni status` can return `READY`, and `ni end` can lock
the accepted plan.

Locked intent:

- Build a read-only web dashboard for a support lead's daily triage workflow.
- Input is a CSV export from the support tool.
- Show open issues grouped by priority, owner, and age bucket.
- Flag issues older than 48 hours.
- Do not connect live APIs, assign tickets, send notifications, or create any
  automation runner.

Accepted evidence:

- A fixed CSV fixture produces expected counts by priority, owner, and age
  bucket.
- A fixture containing issues older and younger than 48 hours proves overdue
  flag behavior.
- Review confirms the UI is read-only and contains no mutation affordances.

Lock boundary:

- `.ni/plan.lock.json` is the authority after lock.
- Downstream prompts may trust the locked hashes.
- If intent changes, execution must stop and the plan must be amended or
  relocked.
