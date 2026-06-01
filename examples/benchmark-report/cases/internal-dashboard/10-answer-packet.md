# Internal dashboard benchmark answer packet

This packet was created to collect user answers to the internal-dashboard
benchmark blockers. It did not resolve the blockers by itself. At packet
creation time, the benchmark remained `BLOCKED` until answers were provided,
the isolated benchmark workspace was updated, and `ni status` reported
readiness that allowed lock.

Use these answers only inside:

```text
examples/benchmark-report/cases/internal-dashboard/workspace/
```

Do not apply them to the repository-root planning lock or root `.ni/` state.

## Original Status Before Answers

- Readiness: `BLOCKED`
- Lock created: no
- `ni-run` prompt compiled: no
- Required answers: `OQ-001` through `OQ-004`

## Resolved Status After Answers

Task 161 applied user-provided answers for `OQ-001` through `OQ-004` to the
isolated workspace as benchmark planning-meeting artifact readiness. The
measured resolved status is recorded in `11-resolved-status-proof.md`; lock and
prompt evidence are recorded in `13-lock-summary.md` and
`14-bounded-prompt-summary.md`.

## How to use this packet

1. Fill in the answers below.
2. Use the answers to update only the benchmark workspace.
3. Run `ni status --proof --next-questions` against the benchmark workspace.
4. Lock only if readiness becomes `READY` or `READY_WITH_DEFERRALS` and user
   confirmation exists.
5. Run `ni run` only after a valid benchmark workspace lock.

## OQ-001 - Primary user and supported decision

Prompt:

Who is the primary dashboard user, and what decision should the dashboard help
them make?

Required answer fields:

- Primary user:
- Secondary users, if any:
- Decision supported:
- Decision timing:
- What should not be supported:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- Guessing which "customer team" role matters.
- Assuming all customer-facing teams need the same dashboard.

## OQ-002 - Attention signals and ranking criteria

Prompt:

What makes an account "need attention," and how should accounts be ranked?

Required answer fields:

- Attention signals:
- Thresholds or review rules:
- Ranking logic:
- Freshness expectations:
- Signals explicitly excluded:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- Inventing account-health metrics.
- Treating any negative signal as equally urgent.

## OQ-003 - Source systems, privacy, and access boundaries

Prompt:

What data may be used, where does it come from, how fresh must it be, and who
may see it?

Required answer fields:

- Source systems:
- Allowed fields:
- Prohibited fields:
- Freshness requirement:
- Access roles:
- Privacy/security constraints:
- Data that must remain out of scope:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- Exposing sensitive customer data.
- Assuming stale data is acceptable.
- Assuming all internal users can see the same fields.

## OQ-004 - Planning-meeting acceptance evidence

Prompt:

What evidence is enough for the planning meeting to accept the result?

Required answer fields:

- Meeting audience:
- Meeting date or timing:
- Minimum artifact:
- Pass/fail criteria:
- What is explicitly not required:
- Who approves acceptance:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- Treating any prototype as sufficient.
- Treating a live dashboard as required if a planning memo or mock is enough.
- Treating unclear approval as acceptance.

## After answers are provided

Expected next steps:

1. Update benchmark workspace `docs/plan`.
2. Update benchmark workspace `.ni/contract.json`.
3. Update benchmark workspace `.ni/session.json`.
4. Run `ni status --proof --next-questions`.
5. If still `BLOCKED`, document remaining blockers.
6. If `READY` or `READY_WITH_DEFERRALS`, lock only inside the benchmark
   workspace.
7. If locked, compile a bounded prompt from the benchmark workspace.
8. Update the benchmark measurement table honestly.

Rules:

- Do not edit root `.ni/plan.lock.json`.
- Do not run repository-root `ni end` or `ni relock`.
- Do not execute downstream agents.
- Do not implement the dashboard.
- Do not fake prompt or lock evidence.
