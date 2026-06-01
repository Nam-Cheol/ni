# Blocker Analysis

This note explains why the research-protocol benchmark remained `BLOCKED`
after the initial status run. It does not answer the blocker questions, mark
them resolved or deferred, lock the workspace, compile a prompt, run
fieldwork, collect data, call downstream agents, or call model APIs.

`BLOCKED` is a valid benchmark result. In this case, ni prevented premature
research handoff by making readiness gaps explicit before any fieldwork,
participant recruitment, data collection, analysis, intervention placement, or
generated prompt execution. The case does not prove research quality,
intervention effectiveness, fieldwork safety, or outcome validity.

It does show that the readiness gaps are explicit, auditable, and tied to
planning records that must be updated only after user answers exist.

| Blocker | Unknown / why it blocks | Required answer | Unsafe assumption avoided | Expected planning update |
| --- | --- | --- | --- | --- |
| `OQ-001` | The research question, supported decision, decision owner, and final study artifact are not accepted. Without them, the protocol could optimize for an invented planning decision or an undefined final deliverable. | A user-confirmed research question, supported decision, decision owner, final study artifact, and what the study must not decide. | Assuming the study chooses intervention locations directly; assuming the artifact is a full statistical report; assuming the question covers all heat risk rather than a specific planning decision. | Update `docs/plan/00_project_brief.md`, `docs/plan/01_actors_outcomes.md`, `docs/plan/02_capabilities.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, and `.ni/session.json` if present. |
| `OQ-002` | Participant or observation units, inclusion and exclusion criteria, locations, sampling rules, and observation windows are not accepted. Without scope, the protocol could expose the wrong people or places to research activity. | User-confirmed observation units, participant groups if any, inclusion and exclusion criteria, locations or neighborhood types, observation window, sampling or selection rule, and out-of-scope populations or locations. | Recruiting vulnerable participants without criteria; treating all neighborhoods as interchangeable; assuming any hot day or any location is valid. | Update `docs/plan/01_actors_outcomes.md`, `docs/plan/02_capabilities.md`, `docs/plan/04_domain_state.md`, `docs/plan/05_constraints.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, and `.ni/session.json` if present. |
| `OQ-003` | Consent, privacy, allowed and prohibited data, storage/retention, translation, accessibility, and sensitive-data boundaries are not accepted. Without data and consent boundaries, lock would make unsafe collection or handling assumptions look trustworthy. | User-confirmed data types allowed and prohibited, consent approach, privacy constraints, storage/retention rule, translation needs, accessibility needs, sensitive-data boundaries, and data that must remain out of scope. | Collecting identifiable data unnecessarily; treating public observation as consent-free in all cases; ignoring translation or accessibility needs; storing sensitive notes without retention boundaries. | Update `docs/plan/03_interaction_contract.md`, `docs/plan/04_domain_state.md`, `docs/plan/05_constraints.md`, `docs/plan/06_risks_security.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, and `.ni/session.json` if present. |
| `OQ-004` | Field-team safety rules, heat/weather stop conditions, exposure limits, vulnerable-group safeguards, and escalation path are not accepted. Without safety rules, the plan could imply fieldwork may continue during dangerous conditions. | User-confirmed field team safety rules, heat/weather stop conditions, maximum exposure or shift limits, hydration/rest requirements, vulnerable-group safeguards, emergency escalation path, and conditions that cancel fieldwork. | Sending field teams out during dangerous heat; continuing observations despite unsafe conditions; treating vulnerable groups as ordinary observation targets without safeguards. | Update `docs/plan/05_constraints.md`, `docs/plan/06_risks_security.md`, `docs/plan/08_delivery_operation.md`, `docs/plan/09_execution_strategy.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, and `.ni/session.json` if present. |
| `OQ-005` | Reviewer, review audience, acceptance evidence, pass/fail criteria, pre-fieldwork readiness checklist, and approval owner are not accepted. Without review criteria, a draft protocol could be mistaken for fieldwork-ready evidence. | User-confirmed reviewer or approval owner, review audience, minimum protocol artifact, acceptance evidence, pass/fail criteria, pre-fieldwork readiness checklist, what is explicitly not required, and who can approve moving forward. | Treating a draft protocol as fieldwork-ready; treating informal agreement as approval; starting data collection before acceptance evidence exists. | Update `docs/plan/01_actors_outcomes.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/08_delivery_operation.md`, `docs/plan/10_open_questions.md`, `docs/plan/11_decision_log.md`, `.ni/contract.json`, and `.ni/session.json` if present. |

## Interpretation

The blocker set is useful because it stops execution at the correct boundary.
The direct prompt hides research, safety, ethics, and evidence assumptions; the
ni path records them as open questions tied to risks, evaluations, non-goals,
and synchronized planning state.

This evidence does not measure fieldwork, participant recruitment, data
collection, analysis, intervention placement, downstream agent performance,
research outcome validity, adoption, cost, latency, rework reduction, or
statistical effect size.

## Measured Result

- Readiness: `BLOCKED`
- Workspace locked: no
- Bounded prompt compiled: no
- Prompt count: `not_measured`
- Downstream execution: none
- Fieldwork: none
- Data collection: none
