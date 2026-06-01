# Measurement table

This is an initial-state qualitative measurement for the vague
research-protocol request. It records what became visible before execution. It
does not report repeated benchmark data and does not claim research protocol
quality, fieldwork readiness, intervention decision readiness, or downstream
agent performance.

| Criterion | Direct-to-agent risk | ni-path evidence | Improved? | Evidence file or command reference |
| --- | --- | --- | --- | --- |
| Missing acceptance criteria | No pass/fail criteria for research question, method, consent, safety, data handling, artifact review, or pre-fieldwork readiness. | The workspace keeps these gaps as OQ-001 through OQ-005 and status reports `BLOCKED`; no acceptance criteria are invented. | yes | `workspace/docs/plan/10_open_questions.md`; `06-ni-status-proof.md` |
| Unmitigated high-risk items | Participant privacy, consent, vulnerable groups, heat/weather safety, biased recommendations, and false readiness claims could be ignored. | RISK-001 through RISK-005 are high severity and include mitigations; status reports high-severity risks have mitigation. | yes | `workspace/docs/plan/06_risks_security.md`; `workspace/.ni/contract.json`; `06-ni-status-proof.md` |
| Unresolved blockers | The direct request hides missing research question, locations, participants, data boundaries, safety rules, and review evidence. | `ni status --proof --next-questions` reports `BLOCKED` because OQ-001 through OQ-005 are open blockers. | yes | `06-ni-status-proof.md`; `07-ni-next-questions.md` |
| Hidden assumptions | A downstream actor would need to guess research question, neighborhoods, participants, data sources, consent/privacy, safety rules, intervention criteria, and acceptance evidence. | The ni workspace names those assumptions as explicit blocker questions instead of treating them as accepted decisions. | yes | `02-direct-to-agent-risk.md`; `workspace/docs/plan/10_open_questions.md` |
| Non-goal coverage | The direct request does not exclude fieldwork, participant recruitment, data collection, analysis, intervention placement, or model/runtime execution. | NG-001 through NG-004 exclude fieldwork, recruitment, data collection, analysis, intervention placement, downstream agents, model APIs, runtime work, and empirical claims. | yes | `workspace/docs/plan/05_constraints.md`; `workspace/.ni/contract.json`; `05-not-measured.md` |
| Delivery surface clarity | The direct request could be interpreted as a study plan, field guide, community process, intervention decision, analysis, or implementation task. | The workspace records `research_protocol` with `document`, `workflow`, and `human_service` surfaces, while keeping fieldwork unauthorized. | yes | `workspace/docs/plan/00_project_brief.md`; `workspace/docs/plan/08_delivery_operation.md`; `06-ni-status-proof.md` |
| Actor/outcome clarity | Research lead, field team, community stakeholders, reviewer, and supported decision are not accepted. | The workspace names likely actors while leaving the accepted study decision, reviewer, and fieldwork readiness gate open as blockers. | partial | `workspace/docs/plan/01_actors_outcomes.md`; `workspace/docs/plan/10_open_questions.md` |
| Evaluation evidence clarity | No reviewer, ethics checkpoint, fieldwork stop rule, acceptance evidence, or prompt boundary is defined. | EVAL-001 through EVAL-003 define initial workspace, blocker visibility, and non-execution boundary reviews; protocol quality evidence remains unresolved. | partial | `workspace/docs/plan/07_evaluation_contract.md`; `05-not-measured.md` |
| Bounded handoff prompt availability | Unavailable; the direct prompt has no lock-verified bounded handoff. | No prompt was compiled because status is `BLOCKED`; prompt count is `not_measured`. | no | `06-ni-status-proof.md`; `05-not-measured.md` |

Measured readiness: `BLOCKED`.

Lock: no.

Bounded prompt: no.

Prompt count: `not_measured`.
