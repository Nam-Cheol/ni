# 두 번째 벤치마크 사례 선택

## Current benchmark state

첫 번째 v0.5 real benchmark evidence case는
`examples/benchmark-report/cases/internal-dashboard/`의 internal-dashboard
case다.

- Vague request: `BLOCKED`
- Answered artifact-readiness: `READY`
- Lock: isolated workspace only
- Prompt: 4000 characters
- Dashboard product readiness: `not_measured`

첫 번째 case는 hidden product assumption이 handoff 전에 visible blocker가 되는
모습을 보여준다는 점에서 가치가 있다. Claim boundary는 의도적으로 좁다.
`READY`는 benchmark planning-meeting artifact readiness만 의미하며 dashboard
product readiness를 의미하지 않는다.

## Candidate comparison

| Candidate | Differentiation value | Benchmark clarity | Lockability | Evidence quality | Cost | Boundary safety | Recommendation |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Research protocol | Very high: software가 아닌 human research, document, field-workflow surface를 분명히 보여준다. | High: 짧은 vague study request만으로 participant, method, consent, safety, artifact gap이 드러난다. | High: user answer가 있으면 fieldwork는 scope 밖에 둔 채 protocol artifact를 `BLOCKED`에서 `READY`로 옮길 수 있다. | Very high: meaningful blocker, high-severity risk mitigation, status proof, isolated lock, bounded human-team handoff를 만들기 좋다. | Low: 기존 fixture와 locked example이 vocabulary를 제공하지만 benchmark는 fresh isolated workspace를 써야 한다. | Very high: fieldwork, participant data collection, ethics approval, analysis, policy claim을 explicit non-goal로 둘 수 있다. | Select |
| Conversation product | Medium-high: web/CLI 밖 conversation surface를 보여주지만 기존 examples가 이미 conversation product를 다룬다. | Medium: "personal but not creepy"는 명확하지만 memory policy가 쉽게 커질 수 있다. | Medium: memory scope, consent, deletion, failure behavior가 답변되면 lockable하다. `BLOCKED` only case로도 가치가 있다. | High: memory, consent, trust, staleness, cross-project boundary blocker가 강하다. | Medium: 이해하기 쉽지만 safety/memory control scope를 신중히 잡아야 한다. | Medium-high: chatbot/runtime을 만들지 않으면 안전하지만 reader가 implementation quality를 추론하기 쉽다. | Defer |
| Operations process | High: workflow/process planning과 human-service boundary를 증명한다. | High: support handoff problem은 이해하고 scoring하기 쉽다. | High: issue scope, owners, escalation rules, accepted tools, evidence criteria가 있으면 lockable하다. | High: role, approval, severity, evidence blocker가 좋다. | Low-medium: 작은 workspace로 가능하지만 software support operations와 인접해 보일 수 있다. | High: tooling, SLA, staffing, automation claim을 피할 수 있다. | Strong later case |
| Education program | High: curriculum planning과 assessment boundary를 증명한다. | Medium-high: 요청은 단순하지만 readiness gap이 technical reader에게 blocker-grade로 덜 선명할 수 있다. | Medium-high: learner profile, outcomes, format, review, accessibility, assessment answer가 있으면 lockable하다. | Medium-high: outcome과 review evidence는 좋지만 privacy/safety 압력은 research보다 약하다. | Low: 설명하기 쉽고 작게 유지하기 쉽다. | High: training delivery, HR policy change, behavioral impact claim을 피할 수 있다. | Later case |

## Recommended second case

**Research protocol** case를 선택한다.

## Why this case

Research protocol은 software dashboard에서 가장 멀리 proof surface를 확장하면서도
감사하기 쉽기 때문에 두 번째 v0.5 benchmark case로 가장 적합하다. Request는 한
문장으로 vague하게 만들 수 있지만, readiness는 participant, consent, safety,
data handling, study artifact, review authority, non-goal에 대한 구체적 답변을
필요로 한다. 그래서 `ni`가 downstream work 전에 unclear intent를 block하고,
나중에 non-code project contract를 document/human-team handoff로 lock할 수
있음을 보여주기에 좋다.

Boundary safety도 강하다. 성공한 benchmark는 protocol artifact readiness, lock
validity, bounded prompt compilation만 주장할 수 있고 fieldwork, participant data
collection, ethics approval, analysis, cooling intervention outcome은 주장하지
않아도 된다.

기존 repository material은 context로만 사용해야 한다.

- `testdata/benchmark/vague-requests/community-heat-field-study/`는 적합한 vague
  fixture와 reviewer seed note를 제공한다.
- `examples/research-protocol/`은 research protocol이 lock될 수 있음을 이미
  보여주지만, v0.5 benchmark는 old lock/run output을 새 measurement로 재사용하지
  말고 fresh isolated evidence를 만들어야 한다.

## Draft vague request

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

## Expected first-run blockers

정확한 ID는 이후 benchmark workspace에서 배정되지만 expected blocker shape는
다음과 같다.

- `OQ-001`: Research question, decision, final study artifact가 accepted 상태가
  아니다.
- `OQ-002`: Participant 또는 observation scope, inclusion/exclusion criteria,
  location이 accepted 상태가 아니다.
- `OQ-003`: Consent, privacy, data handling, translation, accessibility,
  sensitive-data boundary가 accepted 상태가 아니다.
- `OQ-004`: Field-team safety, heat/weather stop rule, vulnerable-group
  safeguard가 accepted 상태가 아니다.
- `OQ-005`: Review owner, acceptance evidence, fieldwork 전 readiness criteria가
  accepted 상태가 아니다.
- Possible sync blocker: 이후 planning conversation이 docs 또는
  `.ni/contract.json` 중 하나만 업데이트하면 synchronization blocker가 떠야 한다.

## Expected user-answer packet

Selection 단계에서는 아래 항목을 채우지 않는다. 이후 answer packet은 다음을
수집해야 한다.

- Study decision: protocol이 도와야 할 decision과 required final artifact.
- Research questions: primary question, secondary question, explicit non-claim.
- Scope: neighborhood, observation window, participant/non-participant boundary,
  inclusion/exclusion criteria.
- Method: interview, observation, survey, sensor review, document review 중 무엇을
  쓸지와 그 method가 benchmark artifact에 충분한 이유.
- Consent and ethics: fieldwork 전에 ethics review가 필요한 항목, reviewer,
  approval 없이는 진행할 수 없는 항목.
- Data handling: allowed data, prohibited data, retention, de-identification,
  access limit.
- Safety: heat/weather stop rule, field-team safety owner, vulnerable-group
  safeguard, escalation path.
- Accessibility: language, translation, disability access, participant burden
  constraint.
- Non-goals: fieldwork 없음, participant data collection 없음, medical advice 없음,
  policy commitment 없음, intervention deployment 없음, `ni`의 statistical outcome
  claim 없음.
- Acceptance evidence: status proof, blocker disposition, risk mitigation review,
  isolated lock proof, bounded prompt count.

## Expected measurement table

이후 benchmark report는 existing protocol과 같은 rubric을 사용하되 result를
만들어내면 안 된다.

| Criterion | Direct-to-agent prompt | ni intent-lock path | Evidence to record |
| --- | --- | --- | --- |
| Missing acceptance criteria | `not_measured` until reviewed | `not_measured` until status proof exists | Research questions, method, consent, safety, data handling, artifact acceptance |
| Unmitigated high-risk items | `not_measured` until reviewed | `not_measured` until status proof exists | Participant safety, vulnerable groups, privacy, heat exposure, false policy claims |
| Unresolved blockers | `not_measured` until reviewed | `not_measured` until CLI status exists | `OQ-*` blockers and grouped next-question output |
| Hidden assumptions | `not_measured` until reviewed | `not_measured` until docs/contract review | Recruitment, ethics, method, output, timeline, data authority |
| Non-goal coverage | `not_measured` until reviewed | `not_measured` until docs/contract review | Fieldwork, data collection, intervention deployment, medical advice, policy promises |
| Stale plan detection | `not_measured` until lock/amendment scenario exists | `not_measured` unless intentionally tested | Lock hash or changed-plan proof, only if included |
| Target prompt boundedness | `not_measured` before lock | `not_measured` until prompt compilation after lock | Character count at or below 4000 |
| Readiness before execution | `not_measured` for direct path | `not_measured` until `ni status` runs | Authoritative status output |

## Non-execution boundary

이 selected case는 다음을 측정하지 않는다.

- implementation quality;
- downstream agent performance;
- real user impact;
- adoption;
- rework reduction;
- cost;
- latency;
- statistical effect size.

또한 fieldwork quality, ethics approval, participant outcome, intervention
effectiveness, policy readiness, analysis validity를 주장하면 안 된다. Downstream
agent, model API, generated prompt execution, shell adapter, queue, telemetry
path, runtime harness도 추가하지 않는다.

## Next task

Task 164: Create isolated research-protocol benchmark workspace and measure
initial BLOCKED readiness
