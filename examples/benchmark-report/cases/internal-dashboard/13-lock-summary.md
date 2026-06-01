# Lock summary

Command:

```bash
go run ./cmd/ni end --dir examples/benchmark-report/cases/internal-dashboard/workspace
```

Output:

```text
locked plan at examples/benchmark-report/cases/internal-dashboard/workspace/.ni/plan.lock.json
status READY
```

Lockfile:

```text
examples/benchmark-report/cases/internal-dashboard/workspace/.ni/plan.lock.json
```

Locked at: `2026-05-29T13:49:39Z`

Readiness: `READY`

Scope: benchmark planning-meeting artifact readiness only. The lock does not
authorize dashboard implementation, downstream agent execution, model API
calls, shell adapters, queues, PR automation, release automation, or empirical
product claims.

Hashed lock sources:

- `.ni/contract.json`
- `docs/plan/00_project_brief.md`
- `docs/plan/01_actors_outcomes.md`
- `docs/plan/02_capabilities.md`
- `docs/plan/03_interaction_contract.md`
- `docs/plan/04_domain_state.md`
- `docs/plan/05_constraints.md`
- `docs/plan/06_risks_security.md`
- `docs/plan/07_evaluation_contract.md`
- `docs/plan/08_delivery_operation.md`
- `docs/plan/09_execution_strategy.md`
- `docs/plan/10_open_questions.md`
- `docs/plan/11_decision_log.md`
