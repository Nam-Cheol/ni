# Measurement table

This is a qualitative before/after measurement for the vague research-protocol
request after synthetic benchmark fixture answers were applied to the isolated
workspace. It is not repeated benchmark data and does not claim research
protocol quality, fieldwork readiness, intervention decision readiness, or
downstream agent performance.

Synthetic answer provenance: the resolved answers are benchmark fixture data
created for ni evidence only. They are not real fieldwork approval, actual
research authorization, proof of research quality, or empirical evidence.

| Criterion | Direct-to-agent risk | ni-path evidence | Improved? | Evidence file or command reference |
| --- | --- | --- | --- | --- |
| Missing acceptance criteria | No pass/fail criteria for research question, method, consent, safety, data handling, artifact review, or pre-fieldwork readiness. | The initial workspace kept these gaps as `OQ-001` through `OQ-005`; the resolved workspace records a protocol memo, acceptance evidence, pass/fail criteria, and pre-fieldwork checklist from the synthetic fixture answers. | yes | `workspace/docs/plan/10_open_questions.md`; `11-resolved-status-proof.md` |
| Unmitigated high-risk items | Participant privacy, consent, vulnerable groups, heat/weather safety, biased recommendations, and false readiness claims could be ignored. | `RISK-001` through `RISK-005` are high severity and include mitigations; resolved status reports high-severity risks have mitigation. | yes | `workspace/docs/plan/06_risks_security.md`; `workspace/.ni/contract.json`; `11-resolved-status-proof.md` |
| Unresolved blockers | The direct request hides missing research question, locations, participants, data boundaries, safety rules, and review evidence. | The initial status proof reports `BLOCKED` because `OQ-001` through `OQ-005` were open; the resolved proof reports `READY` after all five were resolved with labeled fixture answers. | yes | `06-ni-status-proof.md`; `07-ni-next-questions.md`; `11-resolved-status-proof.md`; `12-resolved-next-questions.md` |
| Hidden assumptions | A downstream actor would need to guess research question, neighborhoods, participants, data sources, consent/privacy, safety rules, intervention criteria, and acceptance evidence. | The resolved workspace makes those choices explicit as synthetic fixture answers and keeps real fieldwork, real approval, and empirical claims outside scope. | yes | `02-direct-to-agent-risk.md`; `workspace/docs/plan/00_project_brief.md`; `workspace/docs/plan/10_open_questions.md` |
| Non-goal coverage | The direct request does not exclude fieldwork, participant recruitment, data collection, analysis, intervention placement, or model/runtime execution. | `NG-001` through `NG-004` exclude fieldwork, recruitment, data collection, analysis, intervention placement, downstream agents, model APIs, runtime work, and empirical claims. | yes | `workspace/docs/plan/05_constraints.md`; `workspace/.ni/contract.json`; `05-not-measured.md` |
| Delivery surface clarity | The direct request could be interpreted as a study plan, field guide, community process, intervention decision, analysis, or implementation task. | The workspace records `research_protocol` with `document`, `workflow`, and `human_service` surfaces, while keeping real fieldwork unauthorized. | yes | `workspace/docs/plan/00_project_brief.md`; `workspace/docs/plan/08_delivery_operation.md`; `11-resolved-status-proof.md` |
| Actor/outcome clarity | Research lead, field team, community stakeholders, reviewer, and supported decision are not accepted. | The resolved workspace names the City Resilience Program planning owner, privacy/safety reviewer, community engagement lead, field team, optional adult commenters, supported decision, and review outcome. | yes | `workspace/docs/plan/01_actors_outcomes.md`; `workspace/docs/plan/10_open_questions.md` |
| Evaluation evidence clarity | No reviewer, ethics checkpoint, fieldwork stop rule, acceptance evidence, or prompt boundary is defined. | `EVAL-001` through `EVAL-004` cover scope review, synthetic fixture completeness, non-execution boundary, status proof, lock proof, bounded prompt proof, and 4000-character prompt bound. | yes | `workspace/docs/plan/07_evaluation_contract.md`; `11-resolved-status-proof.md`; `13-lock-summary.md`; `14-bounded-prompt-summary.md` |
| Bounded handoff prompt availability | Unavailable; the direct prompt has no lock-verified bounded handoff. | The isolated workspace was locked with `ni end`, then `ni run --max-chars 4000` compiled a bounded generic prompt of 4000 characters. | yes | `13-lock-summary.md`; `14-bounded-prompt-summary.md` |

Measured readiness after synthetic answers: `READY`.

Lock: yes, inside the isolated benchmark workspace.

Bounded prompt: yes.

Prompt count: `4000`.
