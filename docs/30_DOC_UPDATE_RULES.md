# Document update rules

Planning docs are the reviewable narrative form of `.ni/contract.json`. The
model should keep both forms synchronized during authoring conversation.

## File responsibilities

| File | Primary responsibility |
| --- | --- |
| `docs/plan/00_project_brief.md` | purpose, problem, success definition, kernel boundary |
| `docs/plan/01_actors_outcomes.md` | actors, users, reviewers, downstream consumers, desired outcomes |
| `docs/plan/02_capabilities.md` | capabilities and capability-level scope |
| `docs/plan/03_interaction_contract.md` | interaction mode, user/system responsibilities, handoff behavior |
| `docs/plan/04_domain_state.md` | domain entities, state, lifecycle, terminology |
| `docs/plan/05_constraints.md` | hard constraints, non-goals, forbidden behavior |
| `docs/plan/06_risks_security.md` | risks, severity, mitigations, trust boundaries |
| `docs/plan/07_evaluation_contract.md` | evaluations, methods, readiness evidence |
| `docs/plan/08_delivery_operation.md` | delivery expectations and operating model |
| `docs/plan/09_execution_strategy.md` | downstream strategy after lock, without kernel-owned execution state |
| `docs/plan/10_open_questions.md` | open questions, assumptions, blockers, resolutions |
| `docs/plan/11_decision_log.md` | accepted, rejected, deferred, and not-applicable decisions |

If a turn changes a contract field, update the matching plan file in the same
authoring pass. If a turn changes a plan file in a way that affects readiness,
update `.ni/contract.json` in the same pass.

## Contract update rules

The model should update `.ni/contract.json` from conversation instead of asking
the user to edit JSON by hand. Contract changes must preserve schema shape:

- `project.purpose` captures why the project exists.
- `product_type`, `delivery_surfaces`, and `interaction_mode` capture product
  shape without adding runtime behavior.
- `non_goals[]` captures explicit exclusions.
- `capabilities[]` captures scoped powers and links to requirements,
  evaluations, risks, and artifacts.
- `requirements[]` captures behavior and acceptance criteria.
- `decisions[]` captures explicit planning choices.
- `risks[]` captures severity, status, and mitigation.
- `evaluations[]` captures verification method.
- `artifacts[]` captures planned outputs and paths.
- `open_questions[]` captures unresolved assumptions and blockers.

The model must not weaken accepted requirements, evaluations, risks, or
non-goals to make validation pass. If readiness fails, fix the underlying
planning gap or keep the plan blocked.

## Assumptions and questions

Assumptions belong in the relevant plan doc and should be mirrored as open
questions when they affect readiness. An assumption can become an accepted
decision only after the user confirms it or the conversation provides explicit
evidence.

Resolved questions should keep their ID and record the resolution. Blocker
questions should remain blockers until the answer is reflected in the relevant
contract fields and plan docs.

## Validation rhythm

After meaningful authoring changes, run `ni status` when possible. Before
declaring lock readiness, rely on `ni status`, not model judgment. Before
compiling downstream prompt material, rely on `ni end` and `ni run`.

After the status check, include a concise planning proof summary. It should
only name files and contract fields or IDs that actually changed, quote the
before/after status result from `ni status --proof --next-questions` when
available, and keep unresolved uncertainty visible as assumptions or open
questions. If no files changed, say that no planning artifacts were updated.

For repository changes, also run the project validation commands required by
AGENTS.md.
