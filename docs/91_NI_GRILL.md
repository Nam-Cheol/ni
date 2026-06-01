# NI Grill

`ni-grill` is a model-facing planning challenge skill for draft NI contracts
before `ni-end`.

It adapts the useful `grill-with-docs` pattern identified in
[`90_ENGINEERING_SKILLS_APPLICABILITY.md`](90_ENGINEERING_SKILLS_APPLICABILITY.md)
without copying external skill files. The adaptation is specific to NI:
challenge fuzzy project intent against `docs/plan/**`, `.ni/contract.json`,
and `ni status --proof --next-questions`.

## Boundary

`ni-grill` belongs to the model workspace UX layer.

It is not a task runner, product implementation tool, readiness engine, lock
writer, prompt compiler, downstream adapter, or execution harness.

The rule is the same as every NI skill:

```text
Skills are UX; CLI is authority.
```

`ni-grill` can find weak assumptions, vague decisions, missing acceptance
evidence, docs/contract drift, risky handoff ambiguity, and unsupported claims.
It cannot declare a plan ready, approve a lock, edit `.ni/plan.lock.json`, or
override a `BLOCKED` result from the CLI.

## When To Use It

Use `ni-grill` after enough planning content exists to challenge, and before
the user asks `ni-end` to confirm a lock.

Good moments:

- `ni status` is `BLOCKED` and grouped next questions need sharper framing;
- `ni status` is near `READY` but accepted records still feel vague;
- capabilities exist but acceptance evidence is thin;
- risks, non-goals, or target handoff are under-specified;
- benchmark, proof, or readiness claims need labels such as `measured` or
  `not_measured`.

Avoid using it as a first-run brainstorming substitute. If first-run blockers
or sync diagnostics exist, ask the deterministic CLI questions first.

## Required Process

At the start of a grill turn, the model reads:

- `AGENTS.md`;
- `docs/plan/**`;
- `.ni/contract.json`;
- `.ni/session.json` when present;
- `.ni/plan.lock.json` when present.

Then it runs or requests:

```bash
ni status --dir . --proof --next-questions
```

If the status output contains deterministic blockers, `ni-grill` uses those
before inventing new critique. It preserves grouped next-question labels:
`First-run card`, `Sync repairs`, `Risk decisions`, `Evaluation evidence`,
`Scope boundaries`, and `Open blockers`.

Only after deterministic blockers are accounted for should `ni-grill` add
extra pressure questions against accepted or nearly accepted planning content.

## Grill Categories

`ni-grill` pressure-tests:

- purpose: specific reality change, single-problem focus, observable success;
- actors/outcomes: actor specificity, expected outcome per actor, separation of
  operators, reviewers, and end users;
- delivery surface: explicit surface and docs/contract consistency;
- capabilities and requirements: accepted records have trace links and proof;
- evaluations: evidence is test, review checklist, demo condition, user
  approval, protocol check, or manual inspection;
- risks: high risks have mitigation and privacy/security/safety handling;
- non-goals: likely scope-drift traps are explicitly excluded;
- decisions: accepted, deferred, rejected, and not_applicable statuses are used
  intentionally;
- assumptions: uncertain statements remain assumptions or open questions;
- handoff: downstream actors know what to do and what not to do;
- docs/contract sync: lock-critical records agree across sources;
- claims: benchmark or proof claims are supported and labeled measured or
  not_measured.

## Finding Shape

Each grill finding must be concrete and answerable:

```text
GRILL-001
- Affected planning ID or path: CAP-001 / docs/plan/02_capabilities.md
- Concern: The capability says "usable report" but does not define who accepts it.
- Why it matters: downstream work may optimize for the wrong reviewer.
- Question for the user: Who must approve CAP-001, and what evidence counts?
- Expected answer shape: reviewer role plus test, review checklist, demo condition, user approval, protocol check, or manual inspection
- Blocks ni-end: yes
```

Ask at most five grill questions in one turn. If more issues exist, prioritize
lock blockers, high-risk ambiguity, acceptance evidence gaps,
privacy/security/safety risks, scope drift, and target handoff ambiguity.

## Language Behavior

Ask user-facing grill questions in the language of the user's latest
substantive message. Preserve IDs, commands, paths, status constants, target
names, and schema keys exactly, including `R014`, `OQ-001`, `SYNC-014`,
`GRILL-001`, `ni status`, `.ni/contract.json`, `READY`, and `BLOCKED`.

CLI output may remain English. A model may explain it in the user's language
without changing its meaning.

## Answer Handling

When the user answers grill questions, the model updates `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json` together.

Uncertainty stays visible as assumptions or open questions. Clear exclusions
become non-goals. The model must not convert uncertain answers into accepted
decisions, weaken risks or evaluations, or edit toward readiness.

After updates, run or request:

```bash
ni status --dir . --proof --next-questions
```

Then report planning proof:

- user input captured;
- interpreted planning records;
- updated artifacts;
- status before and after;
- remaining blockers;
- next question group.

## Relationship To Other Skills

`ni-start` authors and maintains planning state during the main conversation.
`ni-grill` challenges a draft plan before lock. `ni-end` summarizes CLI-ready
planning state and asks for explicit confirmation before `ni end`.

`ni-grill` can hand the user back to `ni-start` when answers require planning
edits. It should hand the user to `ni-end` only when `ni status` reports
`READY` or `READY_WITH_DEFERRALS` and the user wants to lock.
