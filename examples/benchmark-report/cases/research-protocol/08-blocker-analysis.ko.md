# Blocker 분석

이 문서는 initial status run 뒤 research-protocol benchmark가 왜 `BLOCKED`로
남아 있는지 설명한다. Blocker question에 답하지 않고, resolved 또는 deferred로
표시하지 않으며, workspace를 lock하거나 prompt를 compile하지 않는다.
Fieldwork, data collection, downstream agent, model API도 실행하지 않는다.

`BLOCKED`는 유효한 benchmark result다. 이 case에서 ni는 fieldwork,
participant recruitment, data collection, analysis, intervention placement,
generated prompt execution 전에 readiness gap을 명시적으로 드러내어 premature
research handoff를 막았다. 이 case는 research quality, intervention
effectiveness, fieldwork safety, outcome validity를 증명하지 않는다.

이 case가 보여주는 것은 readiness gap이 명시적이고 audit 가능하며, user answer가
존재한 뒤에만 update해야 하는 planning record와 연결되어 있다는 점이다.

| Blocker | Unknown / why it blocks | Required answer | Unsafe assumption avoided | Expected planning update |
| --- | --- | --- | --- | --- |
| `OQ-001` | Research question, supported decision, decision owner, final study artifact가 accepted 상태가 아니다. 이것들이 없으면 protocol은 발명된 planning decision이나 정의되지 않은 final deliverable에 맞춰질 수 있다. | User가 확인한 research question, supported decision, decision owner, final study artifact, 그리고 study가 결정하지 말아야 할 것. | study가 intervention location을 직접 선택한다고 가정하는 것, artifact가 full statistical report라고 가정하는 것, research question이 특정 planning decision이 아니라 모든 heat risk를 다룬다고 가정하는 것. | `docs/plan/00_project_brief.md`, `docs/plan/01_actors_outcomes.md`, `docs/plan/02_capabilities.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. |
| `OQ-002` | Participant 또는 observation unit, inclusion/exclusion criteria, location, sampling rule, observation window가 accepted 상태가 아니다. Scope가 없으면 protocol이 잘못된 사람이나 장소를 research activity에 노출할 수 있다. | User가 확인한 observation unit, 필요한 경우 participant group, inclusion criteria, exclusion criteria, location 또는 neighborhood type, observation window, sampling 또는 selection rule, out-of-scope population 또는 location. | Criteria 없이 vulnerable participant를 모집하는 것, 모든 neighborhood가 interchangeable하다고 보는 것, 아무 hot day나 아무 location이 valid하다고 가정하는 것. | `docs/plan/01_actors_outcomes.md`, `docs/plan/02_capabilities.md`, `docs/plan/04_domain_state.md`, `docs/plan/05_constraints.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. |
| `OQ-003` | Consent, privacy, allowed/prohibited data, storage/retention, translation, accessibility, sensitive-data boundary가 accepted 상태가 아니다. Data와 consent boundary 없이는 unsafe collection 또는 handling assumption이 trustworthy한 것처럼 보일 수 있다. | User가 확인한 data types allowed/prohibited, consent approach, privacy constraints, storage/retention rule, translation needs, accessibility needs, sensitive-data boundaries, out-of-scope data. | Identifiable data를 불필요하게 수집하는 것, public observation은 모든 경우 consent-free라고 가정하는 것, translation/accessibility need를 무시하는 것, retention boundary 없이 sensitive note를 저장하는 것. | `docs/plan/03_interaction_contract.md`, `docs/plan/04_domain_state.md`, `docs/plan/05_constraints.md`, `docs/plan/06_risks_security.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. |
| `OQ-004` | Field-team safety rule, heat/weather stop condition, exposure limit, vulnerable-group safeguard, escalation path가 accepted 상태가 아니다. Safety rule이 없으면 위험한 조건에서도 fieldwork를 계속할 수 있다는 암시가 생긴다. | User가 확인한 field team safety rules, heat/weather stop conditions, maximum exposure 또는 shift limits, hydration/rest requirements, vulnerable-group safeguards, emergency escalation path, fieldwork cancel condition. | Dangerous heat에 field team을 내보내는 것, unsafe condition에서도 observation을 계속하는 것, vulnerable group을 ordinary observation target으로 safeguard 없이 다루는 것. | `docs/plan/05_constraints.md`, `docs/plan/06_risks_security.md`, `docs/plan/08_delivery_operation.md`, `docs/plan/09_execution_strategy.md`, `docs/plan/10_open_questions.md`, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. |
| `OQ-005` | Reviewer, review audience, acceptance evidence, pass/fail criteria, pre-fieldwork readiness checklist, approval owner가 accepted 상태가 아니다. Review criteria가 없으면 draft protocol이 fieldwork-ready evidence로 오해될 수 있다. | User가 확인한 reviewer 또는 approval owner, review audience, minimum protocol artifact, acceptance evidence, pass/fail criteria, pre-fieldwork readiness checklist, explicitly not required 항목, moving forward를 approve할 수 있는 사람. | Draft protocol을 fieldwork-ready로 보는 것, informal agreement를 approval로 간주하는 것, acceptance evidence가 생기기 전에 data collection을 시작하는 것. | `docs/plan/01_actors_outcomes.md`, `docs/plan/07_evaluation_contract.md`, `docs/plan/08_delivery_operation.md`, `docs/plan/10_open_questions.md`, `docs/plan/11_decision_log.md`, `.ni/contract.json`, 그리고 있으면 `.ni/session.json`을 update한다. |

## 해석

이 blocker set은 올바른 boundary에서 execution을 멈추기 때문에 유용하다.
Direct prompt는 research, safety, ethics, evidence assumption을 숨긴다. ni path는
그 assumption을 risk, evaluation, non-goal, synchronized planning state와 연결된
open question으로 기록한다.

이 evidence는 fieldwork, participant recruitment, data collection, analysis,
intervention placement, downstream agent performance, research outcome validity,
adoption, cost, latency, rework reduction, statistical effect size를 측정하지
않는다.

## 측정된 결과

- Readiness: `BLOCKED`
- Workspace locked: no
- Bounded prompt compiled: no
- Prompt count: `not_measured`
- Downstream execution: none
- Fieldwork: none
- Data collection: none
