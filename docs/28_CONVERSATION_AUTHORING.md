# Conversation authoring

`ni` authoring starts after `ni init` creates a planning workspace. From that
point, the primary authoring interface is sustained model-user conversation,
not manual contract commands.

The user explains intent, corrects drafts, answers questions, rejects scope,
and approves decisions. The model turns that conversation into:

- human-readable planning docs under `docs/plan/**`,
- the machine-readable `.ni/contract.json`,
- bounded continuity state in `.ni/session.json`,
- visible open questions and assumptions,
- readiness gaps that can be checked by `ni status`.

The CLI remains the authority for validation, lock, and prompt compilation.
Conversation is the authoring surface; `ni status`, `ni end`, and `ni run` are
the gates and compilers.

## Extraction model

During planning conversation, the model extracts these elements:

| Conversation signal | Planning record |
| --- | --- |
| The project reason, target outcome, or problem to solve | `project.purpose`, `docs/plan/00_project_brief.md` |
| People, systems, reviewers, operators, and downstream consumers | `docs/plan/01_actors_outcomes.md` |
| Product actions or user-visible powers | `capabilities[]`, `docs/plan/02_capabilities.md` |
| Required behavior, acceptance criteria, or constraints on a capability | `requirements[]`, linked from `capabilities[].requirements` |
| Explicit choices the user accepts or rejects | `decisions[]`, `docs/plan/11_decision_log.md` |
| Possible failures, hazards, or trust boundaries | `risks[]`, `docs/plan/06_risks_security.md` |
| Ways to verify the plan or later work | `evaluations[]`, `docs/plan/07_evaluation_contract.md` |
| Work products, docs, prompts, or outputs the plan must produce | `artifacts[]`, linked from `capabilities[].artifacts` |
| Out-of-scope behavior | `non_goals[]`, relevant plan docs |
| Product boundaries, policies, budgets, profiles, and delivery surfaces | `.ni/contract.json`, `docs/plan/05_constraints.md` |
| Unknowns, conflicts, or unstated acceptance criteria | `open_questions[]`, `docs/plan/10_open_questions.md` |

The model should preserve user language where it affects meaning, but it should
normalize records into stable IDs and statuses that the CLI can validate.
`.ni/session.json` may carry a short planning summary between model sessions,
but it is below contract and docs authority and must be corrected when it
conflicts with them.

## Uncertainty rule

Uncertain statements do not become accepted decisions.

If the user says something tentative, inferred, conflicting, or incomplete, the
model must record it as one of:

- a draft requirement or draft capability,
- an assumption visible in the relevant planning doc,
- an open question in `.ni/contract.json` and `docs/plan/10_open_questions.md`.

Examples of uncertain signals include "probably", "maybe", "for now",
"I think", "unless", "not sure", missing owners, missing evaluation methods,
and contradictions with accepted decisions or non-goals.

An open question should be marked as a blocker when it prevents a deterministic
readiness decision, affects acceptance criteria, or could change high-severity
risk mitigation.

## Authority boundary

The model may draft, reconcile, and explain planning state. It may not declare
the plan complete by judgment alone.

Authoritative checks remain:

- `ni status` determines readiness.
- `ni end` locks a ready plan and writes `.ni/plan.lock.json`.
- `ni run` compiles a bounded prompt from a valid locked plan.

If the model believes the docs are complete but `ni status` reports `BLOCKED`,
the plan is blocked. If a lock hash mismatch exists, the model must stop and
report `BLOCKED` instead of editing around the mismatch.

After lock, the source-of-truth order is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

## Non-authoring commands

The authoring protocol must not add user-facing contract `add`, `list`, or
`set` commands. Users should not have to manually edit `.ni/contract.json`.
The model and skills maintain it from conversation, and the CLI validates the
result.
