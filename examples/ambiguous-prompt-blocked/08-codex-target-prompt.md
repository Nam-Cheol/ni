# 08. Codex target prompt

This is the kind of bounded prompt `ni run --target codex` can compile after
the plan is locked. It is seed material, not execution.

```text
You are implementing a locked NI plan.

Authority:
- Trust .ni/plan.lock.json first.
- If the locked plan hash does not match the current contract/docs, stop and
  report BLOCKED.
- Do not reinterpret open product intent from chat memory.

Goal:
Build a read-only web dashboard for a support lead's daily customer-issue
triage workflow.

Accepted requirements:
- Use a CSV export from the support tool as input.
- Show open issues grouped by priority, owner, and age bucket.
- Flag issues older than 48 hours.
- Provide a clear empty and malformed CSV state.
- Keep the dashboard read-only.

Non-goals:
- Do not connect live customer systems.
- Do not assign tickets.
- Do not send notifications.
- Do not create queue, shell adapter, or task-runner behavior.

Evaluation:
- Use a fixed CSV fixture to verify counts by priority, owner, and age bucket.
- Verify overdue flags around the 48-hour boundary.
- Verify no mutation controls are present.
```

The prompt tells Codex what to build only after ni has blocked and resolved the
ambiguous execution request.
