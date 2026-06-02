---
name: ni-grill
description: Challenge NI planning quality before ni-end by reading docs, contract, and ni status proof without replacing the CLI readiness gate.
---

# ni-grill

Use this skill when the user says `ni-grill`, asks to pressure-test a draft NI
plan, or wants focused critique before `ni-end`.

`ni-grill` challenges planning quality before lock. It does not execute work.
It is a planning challenge UX layer for `ni-kernel`, not a product
implementation tool and not downstream execution.

## Authority

Skills are UX; CLI is authority.

`ni status` is the authority for `BLOCKED`, `READY_WITH_DEFERRALS`, and
`READY`. `ni-grill` never approves lock by model judgment. If `ni status` is
`BLOCKED`, `ni-grill` should use deterministic blockers before inventing new
critique.

If ni status is BLOCKED, ni-grill should use deterministic blockers before
inventing new critique.

If `.ni/plan.lock.json` exists, respect the source-of-truth order and do not
silently edit locked planning docs. If a lock hash mismatch exists, stop and
report `BLOCKED`.

`LOCK-STALE` means the existing lock no longer matches current planning inputs.
- Skills may help draft amended planning text.
- Skills may help explain `LOCK-STALE`.
- Skills do not determine readiness.
- Skills do not lock or relock.
- Skills do not replace `ni status`, `ni end`, or `ni run`.
- Skills do not update `.ni/plan.lock.json`.

Recovery order: `review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000`.

## Start Process

1. Read `AGENTS.md`, `docs/plan/**`, `.ni/contract.json`,
   `.ni/session.json` if present, and `.ni/plan.lock.json` if present.
2. Run or request:

```bash
ni status --dir . --proof --next-questions
```

3. Preserve grouped next questions first: `First-run card`, `Sync repairs`,
   `Risk decisions`, `Evaluation evidence`, `Scope boundaries`, and
   `Open blockers`.
4. Do not ask generic brainstorming questions when deterministic blockers or
   grouped next questions exist.
5. Apply grill questions only to accepted or nearly accepted planning content.

## Grill Scope

Pressure-test these areas:

- purpose: specific reality change, single problem focus, observable success;
- actors/outcomes: specific actors, actor outcomes, operator/reviewer/end-user
  separation;
- delivery surface: explicit surface and docs/contract consistency;
- capabilities and requirements: accepted records have proof and trace links;
- evaluations: evidence is test, review checklist, demo condition, user
  approval, protocol check, or manual inspection;
- risks: high risks have mitigation; privacy, security, and safety are named;
- non-goals: scope-drift traps are excluded;
- decisions: statuses are accepted, deferred, rejected, or not_applicable;
- assumptions: uncertain statements remain assumptions or open questions;
- handoff: downstream actors know what to do and what not to do;
- docs/contract sync: `docs/plan/**` and `.ni/contract.json` agree;
- claims: benchmark or proof claims are supported and labeled measured or
  not_measured.

## Finding Format

Use severity labels: `Critical`, `High`, `Medium`, `Low`, and `Note`.
Severity is advisory planning pressure, not CLI readiness. Skills are UX; CLI
is authority.

Show at most 5 findings by default and ask at most 5 user-facing questions in
one turn. If `Critical` or `High` findings exist, show at most 3 of them first.
Prioritize deterministic CLI blockers, then `Critical`, `High`, acceptance
evidence gaps, privacy/security/safety risks, claim-boundary risks,
non-goal/scope drift risks, handoff ambiguity, `Medium`/`Low` editorial issues,
and `Note` observations. If more findings exist, summarize:

```text
N additional lower-priority findings were not shown.
```

When `ni status` is `BLOCKED`, prioritize deterministic blockers and avoid a
long secondary critique. When `ni status` is `READY` or
`READY_WITH_DEFERRALS`, focus on claim quality, public handoff, risk clarity,
and overclaim prevention.

For each finding, use this shape:

```text
Grill findings:
1. GRILL-001 — <severity> — <category>
   Affected: <file path or planning ID>
   Concern: <specific concern>
   Why it matters: <why downstream handoff or lock quality could suffer>
   Question: <user-facing question>
   Answer shape: <expected answer form>
   Suggested action: <resolve / defer / mark non-goal / clarify / keep as note>
   Blocks ni-end: <CLI decides / likely yes / maybe / no>
```

Use `Blocks ni-end: CLI decides` for deterministic readiness findings, `likely
yes` only for severe planning-quality issues that should be resolved before
lock, `maybe` for user-confirmable tradeoffs, and `no` for clarity/editorial
findings.

Ask user-facing grill questions in the user's latest substantive language.
Preserve IDs, commands, paths, status constants, target names, schema keys, and
severity labels: `R014`, `OQ-001`, `SYNC-014`, `GRILL-001`, `ni status`,
`.ni/contract.json`, `READY`, `READY_WITH_DEFERRALS`, `BLOCKED`, `Critical`,
`High`, `Medium`, `Low`, and `Note`.

## Updating After Answers

When the user answers grill questions, update `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json` together. Do not create accepted
decisions from uncertainty; record uncertainty as assumptions or open
questions.

After updates, run or request `ni status --dir . --proof --next-questions`
again and show planning proof:

- user input captured;
- interpreted planning records;
- updated artifacts;
- status before and after;
- remaining blockers;
- next question group.

- Skills may help draft or explain proof-related planning text.
- Skills do not determine readiness.
- Skills do not lock plans.
- Skills do not replace `ni status`, `ni end`, or `ni run`.

## Do not

- Do not execute downstream work.
- Do not implement the product.
- Do not replace CLI readiness.
- Do not declare docs complete by model judgment.
- Do not weaken acceptance criteria, risks, mitigations, requirements,
  evaluations, or non-goals to pass validation.
- Do not manually edit `.ni/plan.lock.json`.
