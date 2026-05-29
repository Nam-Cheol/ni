# 벤치마크 사례 연구: 의도 준비도

이 문서는 수동 정성 사례 연구 보고서다. 저장소에 포함된 모호한 요청과
ni 계획 예제를 `docs/43_BENCHMARK_PROTOCOL.md`의 루브릭으로 비교한다.
반복 실험이 아니며, 외부 모델 API를 호출하지 않고, 다운스트림 작업을
실행하지 않으며, 통계적 유의성을 주장하지 않는다.

## 사례 1: 환불 분류 어시스턴트

- 픽스처:
  `testdata/benchmark/vague-requests/refund-triage-assistant/`
- ni 경로 산출물:
  `examples/ni-start-dogfood/`
- 리뷰어 역할: 저장소 관리자가
  `docs/43_BENCHMARK_PROTOCOL.md` 루브릭을 적용
- 채점일: 2026-05-28
- 비교 대상:
  - A. 직접 에이전트에 전달하는 프롬프트 준비도
  - B. `ni-start -> ni status -> ni-end -> ni-run` 준비도

이 사례는 체크인된 ni-start dogfood 전사와 같은 출발 의도를 사용한다.
즉, 지원 상담원을 위한 환불 분류 어시스턴트를 만들고 싶다는 모호한
요청이다. 직접 경로는 정적 프롬프트 텍스트만 검토했다. ni 경로는
체크인된 계획 산출물과 검증된 CLI 출력으로 검토했다.

## 측정한 것

이 사례 연구들은 실행 전 의도 준비도를 측정한다.

- 누락된 인수 기준,
- 완화되지 않은 high 위험,
- 해결되지 않은 blocker,
- 숨은 가정,
- non-goal 적용 범위,
- 경계가 정해진 target prompt의 존재 여부.

측정하지 않는 것은 구현 품질, 모델 품질, 다운스트림 에이전트 성능,
사용자 만족도, 런타임 안전성, 비용, 지연 시간, 통계적 효과 크기,
프로덕션 결과다.

사례 3은 사례 1과 2보다 의도적으로 가볍다. 완료된 lock/run 측정이
아니라 docs-only 수동 readiness drill이다. 사용할 수 없는 cell은
`not_measured`로 남긴다.

## 사례 1 산출물: A. 직접 에이전트 프롬프트

출처:
`testdata/benchmark/vague-requests/refund-triage-assistant/request.md`

```text
I want a refund triage assistant for support agents.
```

픽스처의 수동 리뷰 노트:

```text
Expected hidden assumptions include:
- The assistant only drafts recommendations and does not approve, issue, or
  initiate refunds.
- Support agents, not customers, are the primary users.
- A specific internal refund policy source exists and is authoritative.
- Ambiguous policy cases should escalate to a supervisor instead of inventing a
  policy interpretation.

Expected readiness gaps include:
- Missing acceptance criteria for recommendation scope, policy citation,
  escalation behavior, privacy handling, transcript review, and runtime
  boundary.
- High-risk gaps around refund authority and stale or unclear policy
  interpretation.
- No explicit non-goals for refund approval, customer contact, supervisor
  replacement, or live integration work.
```

## 사례 1 산출물: B. ni Intent-Lock 경로

체크인된 ni-start dogfood 전사는 다음 계획 경로를 기록한다.

- 모호한 사용자 아이디어: `examples/ni-start-dogfood/02-user-vague-idea.md`
- 확인된 답변: `examples/ni-start-dogfood/04-user-answers.md`
- docs와 contract 변경:
  `examples/ni-start-dogfood/05-docs-contract-delta.md`
- status 증거: `examples/ni-start-dogfood/06-status-proof.md`
- end 확인: `examples/ni-start-dogfood/08-ni-end-confirmation.md`
- run handoff: `examples/ni-start-dogfood/09-ni-run-handoff.md`
- 계획 workspace: `examples/ni-start-dogfood/workspace/`
- 컴파일된 prompt:
  `examples/ni-start-dogfood/workspace/generated/human-team.prompt.txt`

마지막 blocker가 답변되기 전의 `ni status --next-questions` 결과:

```text
BLOCKED
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document

blocker R009: OQ-001 is a blocker open question

question R009 OQ-001: Which refund policy source is authoritative for refund eligibility?
```

blocker가 답변되고 계획이 업데이트된 뒤, 2026-05-28에 확인한 명령 출력:

```text
$ go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace
READY_WITH_DEFERRALS
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document
interaction mode: human_to_system
guidance: product_type=conversation_product changes planning guidance only; readiness authority remains the shared gate.
guidance: Cover conversation turns, failure handling, transcript evaluation, and human handoff.
guidance: conversation surface: specify turn boundaries, memory expectations, refusals, and escalation.
guidance: document surface: specify audience, structure, review workflow, and publication format.
deferral D001: DEC-004 is deferred
deferral D002: OQ-002 remains open
```

체크인된 lock 데이터를 덮어쓰지 않기 위해 임시 복사본에서 `ni end`를
확인했다.

```text
$ go run ./cmd/ni end --dir /private/tmp/ni-benchmark-case-study.HlL2Rc/workspace
locked plan at /private/tmp/ni-benchmark-case-study.HlL2Rc/workspace/.ni/plan.lock.json
status READY_WITH_DEFERRALS
```

`ni run`도 임시 복사본에서 확인했다.

```text
$ go run ./cmd/ni run --dir /private/tmp/ni-benchmark-case-study.HlL2Rc/workspace --target human-team --max-chars 4000 --out /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt
compiled prompt at /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt

$ wc -m /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt
    3805 /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt
```

실제 컴파일된 prompt 산출물:

```text
Human-team handoff

Goal: Hand this locked NI plan to a PM/dev/design/research team for coordinated implementation planning.

Project: Refund Triage Assistant Plan - Plan a support-agent assistant that drafts refund recommendations from tickets and the internal policy page, escalates ambiguity, and excludes refund approval, customer contact, and runtimes.
Readiness: READY_WITH_DEFERRALS
Target: human-team (handoff)
Locked at: 2026-05-26T13:48:04Z

Authoritative sources:
- .ni/plan.lock.json is authoritative for lock state, hashes, and source-of-truth order.
- .ni/contract.json carries accepted CAP/REQ/EVAL/RISK IDs and acceptance criteria.
- docs/plan/ contains locked planning context; use only when hashes match.
- .ni/session.json is a planning aid below locked docs; it must not override contract or docs.
Source of truth: .ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history

Rules:
- Treat .ni/plan.lock.json as authoritative over .ni/contract.json, docs/plan, session state, and chat.
- Verify locked file hashes before using planning docs; on lock mismatch report BLOCKED and stop.
- If there are conflicting requirements, report BLOCKED with conflicting sources or IDs and stop.
- Do not weaken acceptance criteria, risk mitigations, or blocker handling to proceed.
- Every work item must trace to CAP/REQ/EVAL/RISK IDs.
- Keep ni at the pre-runtime boundary; this prompt is seed material, not kernel-owned execution state.
- Do not make ni run execute shell, Codex, adapters, queues, PR automation, or agent teams.

Target instructions:
- PM: maintain scope against accepted capabilities, non-goals, and open blocker questions.
- Dev: identify implementation packets and validation mapped to CAP/REQ/EVAL/RISK IDs.
- Design/research: confirm user-facing assumptions, evidence needs, and unresolved risks without changing acceptance criteria.
- Execution responsibility stays outside ni; this is a handoff artifact, not orchestration.

Process:
1. Review .ni/plan.lock.json first, then .ni/contract.json and docs/plan.
2. Assign next ownership only after confirming no lock mismatch or conflicting requirements.
3. Record validation evidence and blockers for the team before implementation proceeds.

Accepted capabilities:
- CAP-001: Draft refund recommendations for support agents.
- CAP-002: Escalate ambiguous or conflicting refund policy cases.
- CAP-003: Maintain synchronized human docs and machine contract records from conversation.

Accepted requirements:
- REQ-001: Draft recommendations only; never approve, issue, or initiate refunds.
- REQ-002: Cite ticket facts and the internal refund policy section used.
- REQ-003: Escalate policy ambiguity or ticket-policy conflict to a supervisor.
- REQ-004: Escalation language must not invent policy or customer commitments.
- REQ-005: Model-authored planning turns must update docs/plan/** and .ni/contract.json together.
- REQ-006: Readiness, lock, and prompt compilation must be delegated to ni status, ni end, and ni run rather than model judgment.

Risks:
- RISK-001 (high): Require recommendation language, agent review, and transcript checks for approval or customer-contact wording.
- RISK-002 (medium): Use only ticket facts needed for triage and keep generated handoffs inside the support-agent workflow.
- RISK-003 (high): Cite policy, escalate conflicts or unclear sections, and keep policy ownership visibly deferred.
- RISK-004 (medium): Name changed docs and affected IDs each authoring turn, then run ni status.

Open questions:
- OQ-002 blocker=false: Which support dashboard will eventually display the recommendation draft?

Expected output: team handoff with owners, next packets, validation evidence, risks, decisions needed, and BLOCKED if the lock or requirements conflict.
```

## 사례 1 수동 준비도 비교

아래 수치는 이 단일 사례에 대한 한 리뷰어의 수동 루브릭 적용이다. 반복
수치 벤치마크 데이터가 아니다.

| Metric | A. 직접 프롬프트 | B. ni intent-lock 경로 |
| --- | --- | --- |
| Missing acceptance criteria | 6개 범주 누락: 추천 권한, 정책 인용, escalation 동작, privacy 처리, transcript 평가, runtime boundary | locked contract에 blocker급 gap 0개; accepted requirement 6개와 evaluation 3개가 존재 |
| Unmitigated high risks | high-risk 영역 2개가 보이지만 완화 없음: 환불 권한 암시, 오래되었거나 불명확한 정책 해석 | 완화되지 않은 high risk 0개; `RISK-001`과 `RISK-003`은 high severity이며 mitigation을 포함 |
| Unresolved blockers | blocker급 미지수 5개: 권한, 정책 출처, escalation 기준, 평가 증거, handoff target | blocker 0개; status는 `READY_WITH_DEFERRALS`이고 non-blocking deferral 2개가 보임 |
| Hidden assumptions | fixture note에 material assumption 7개 | 측정 범위에서는 hidden 0개로 계산; 남은 불확실성은 `DEC-004`와 `OQ-002` deferral로 보임 |
| Non-goal coverage | 없음 | 명시적: 환불 발행 또는 승인 없음, 고객 연락 없음, supervisor 대체 없음, runtime 또는 live integration 없음 |
| Bounded target prompt availability | 없음; 직접 프롬프트에는 lock 검증된 compiled target prompt가 없음 | 통과; `ni run --max-chars 4000`이 3,805자 `human-team` prompt를 생성 |

## 사례 2: 동네 냉방 연구 프로토콜

- 픽스처:
  `testdata/benchmark/vague-requests/community-heat-field-study/`
- ni 경로 산출물:
  `examples/research-protocol/`
- 리뷰어 역할: 저장소 관리자가
  `docs/43_BENCHMARK_PROTOCOL.md` 루브릭을 적용
- 채점일: 2026-05-29
- 비교 대상:
  - A. 직접 에이전트에 전달하는 프롬프트 준비도
  - B. `ni-start -> ni status -> ni-end -> ni-run` 준비도

두 번째 사례는 소프트웨어가 아닌 research protocol 요청이다. 벤치마크
증거가 software assistant나 dashboard에만 머물지 않는지 확인하기 위해
포함했다. 직접 경로는 정적 프롬프트 텍스트만 검토했다. ni 경로는
체크인된 locked planning artifact와 검증된 CLI 출력으로 검토했다.

## 사례 2 산출물: A. 직접 에이전트 프롬프트

출처:
`testdata/benchmark/vague-requests/community-heat-field-study/request.md`

```text
Category: research protocol

Plan a field study to learn how residents deal with extreme heat and what
support would help them. Make it practical for the team to run soon.
```

픽스처의 수동 리뷰 노트:

```text
Expected hidden assumptions include:
- Residents can be recruited quickly and safely.
- Extreme heat exposure can be discussed without creating participant risk.
- The team has authority to collect, store, and analyze resident responses.
- "Support would help" means a program, physical resource, communication plan,
  or policy recommendation.
- The study method can be selected without a formal ethics or consent review.
- "Soon" gives enough time for protocol design, translation, accessibility,
  and field-team training.

Expected readiness gaps include:
- Missing acceptance criteria for research questions, sampling, consent,
  safety, data handling, field readiness, and final artifacts.
- High-risk gaps around participant safety, vulnerable populations, privacy,
  and environmental exposure.
- No explicit non-goals for intervention delivery, medical advice, policy
  commitments, or quantitative claims.
```

## 사례 2 산출물: B. ni Intent-Lock 경로

체크인된 research-protocol 예제는 locked planning path를 기록한다.

- 계획 workspace: `examples/research-protocol/`
- contract summary: `examples/research-protocol/contract-summary.md`
- planning docs: `examples/research-protocol/docs/plan/`
- locked contract: `examples/research-protocol/.ni/contract.json`
- lockfile: `examples/research-protocol/.ni/plan.lock.json`
- 컴파일된 prompt:
  `examples/research-protocol/generated/human-team.prompt.md`

픽스처의 suggested ni questions:

```text
- What are the exact research questions and expected study artifacts?
- Who may participate, and which groups require extra safeguards?
- What consent, privacy, translation, and accessibility requirements apply?
- What weather, location, or field-team safety conditions must stop the study?
- Which recommendations or interventions are out of scope for this protocol?
- Who reviews and accepts the protocol before fieldwork begins?
```

2026-05-29에 확인한 `ni status` 출력:

```text
$ go run ./cmd/ni status --dir examples/research-protocol
READY
profile: prototype
product type: research_protocol
delivery surfaces: document
interaction mode: human_to_human
guidance: product_type=research_protocol changes planning guidance only; readiness authority remains the shared gate.
guidance: Cover hypothesis, data handling, method, analysis, ethics, and reproducibility evidence.
guidance: document surface: specify audience, structure, review workflow, and publication format.
```

체크인된 lock 데이터를 덮어쓰지 않기 위해 임시 복사본에서 `ni end`를
확인했다.

```text
$ go run ./cmd/ni end --dir /private/tmp/ni-research-case-study.dYqnEa/research-protocol
locked plan at /private/tmp/ni-research-case-study.dYqnEa/research-protocol/.ni/plan.lock.json
status READY
```

컴파일된 prompt는 실행하지 않고 `ni run`만 확인했다.

```text
$ go run ./cmd/ni run --dir examples/research-protocol --target human-team --max-chars 4000 --out /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md
compiled prompt at /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md

$ wc -m /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md
    3621 /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md
```

Locked contract summary:

```text
Product type: research_protocol
Delivery surface: document
Readiness: READY
Accepted capabilities:
- CAP-001: Define the study question, hypothesis, sampling frame, and inclusion criteria.
- CAP-002: Specify data handling, ethics, and participant-contact boundaries.
- CAP-003: Define reproducible analysis and evidence requirements for the protocol.
Accepted requirements:
- REQ-001: The protocol must state the research question and hypothesis in falsifiable terms.
- REQ-002: The sampling frame must define eligible street segments, observation windows, and exclusion criteria.
- REQ-003: The protocol must separate public environmental observations from any human-subject interaction.
- REQ-004: The protocol must require ethics review before any participant contact or personally identifying data collection.
- REQ-005: The analysis plan must name variables, comparison method, reproducibility evidence, and reviewer acceptance criteria.
High risks:
- RISK-001: Ambiguous sampling criteria could bias the study; mitigated by explicit inclusion and exclusion criteria plus independent reviewer reproduction.
- RISK-002: Participant privacy or ethics boundaries may be crossed; mitigated by separating environmental observation from participant contact and requiring ethics review before personal data collection.
Non-goals:
- No fieldwork, participant data collection, sensor deployment, analysis runtime, or ethics approval from ni.
```

3,621자로 제한 안에 들어간 target prompt excerpt:

```text
Project: Neighborhood Cooling Study Protocol - Plan a documented field research protocol for comparing street-level cooling interventions before any study execution begins.
Readiness: READY
Target: human-team (handoff)

Accepted capabilities:
- CAP-001: Define the study question, hypothesis, sampling frame, and inclusion criteria.
- CAP-002: Specify data handling, ethics, and participant-contact boundaries.
- CAP-003: Define reproducible analysis and evidence requirements for the protocol.

Accepted requirements:
- REQ-001: The protocol must state the research question and hypothesis in falsifiable terms.
- REQ-002: The sampling frame must define eligible street segments, observation windows, and exclusion criteria.
- REQ-003: The protocol must separate public environmental observations from any human-subject interaction.
- REQ-004: The protocol must require ethics review before any participant contact or personally identifying data collection.
- REQ-005: The analysis plan must name variables, comparison method, reproducibility evidence, and reviewer acceptance criteria.
```

## 사례 2 수동 준비도 비교

아래 수치는 이 단일 사례에 대한 한 리뷰어의 수동 루브릭 적용이다. 반복
수치 벤치마크 데이터가 아니다.

| Metric | A. 직접 프롬프트 | B. ni intent-lock 경로 |
| --- | --- | --- |
| Missing acceptance criteria | 7개 범주 누락: research question, sampling, consent, safety, data handling, field readiness, final artifact | locked contract에 blocker급 gap 0개; accepted requirement 5개와 evaluation 3개가 존재 |
| Unmitigated high risks | high-risk 영역 4개가 보이지만 완화 없음: participant safety, vulnerable populations, privacy, environmental exposure | 완화되지 않은 high risk 0개; `RISK-001`과 `RISK-002`는 high severity이며 mitigation을 포함 |
| Unresolved blockers | blocker급 미지수 5개: participant criteria, location, study method, consent process, review owner | blocker 0개; status는 `READY` |
| Hidden assumptions | fixture note에 material assumption 6개 | 측정 범위에서는 hidden 0개로 계산; 가정은 requirement, risk, evaluation, non-goal로 표현됨 |
| Non-goal coverage | 없음 | 명시적: ni가 fieldwork, participant data collection, sensor deployment, analysis runtime, ethics approval을 수행하지 않음 |
| Bounded target prompt availability | 없음; 직접 프롬프트에는 lock 검증된 compiled target prompt가 없음 | 통과; `ni run --max-chars 4000`이 3,621자 `human-team` prompt를 생성 |

## 사례 3: 모호한 내부 대시보드 요청

- 픽스처:
  `testdata/benchmark/vague-requests/customer-dashboard/`
- 사례 산출물:
  `examples/benchmark-report/cases/internal-dashboard/`
- 리뷰어 역할: 저장소 관리자가
  `docs/43_BENCHMARK_PROTOCOL.md` 루브릭을 적용
- 채점일: 2026-05-29
- 비교 대상:
  - A. 직접 에이전트에 전달하는 프롬프트 준비도
  - B. `ni-start -> ni status` 준비도 경로. `BLOCKED`에서 멈춘다.

세 번째 사례는 software dashboard request에 대한 수동 정성 readiness
drill이다. Internal dashboard는 downstream actor가 너무 일찍 만들기 쉬운
예시다. web surface가 뻔해 보여도 users, decisions, account signals,
privacy constraints, data freshness, success criteria, risks, non-goals가
빠져 있을 수 있다.

사례 1과 2와 달리 이 사례는 locked workspace가 아니며 handoff prompt를
compile하지 않는다. 대신 checked-in ni planning workspace와 실제
`ni status --proof --next-questions` 출력을 포함한다. authoritative status가
`BLOCKED`이므로 `ni end`나 `ni run`은 실행하지 않았고 prompt character count는
`not_measured`로 남긴다.

## 사례 3 산출물: A. 직접 에이전트 프롬프트

출처:
`testdata/benchmark/vague-requests/customer-dashboard/request.md`

```text
Category: software dashboard

Build a dashboard for the customer team so they can see what is going on with
accounts and know who needs attention. It should be easy to use and ready for
the next planning meeting.
```

픽스처의 수동 리뷰 노트:

```text
Expected hidden assumptions include:
- The dashboard is for customer success managers, not executives, support, or
  sales.
- "Needs attention" means renewal risk, support escalation, product usage drop,
  or unpaid invoice status.
- Required account data already exists in a trusted system and can be accessed
  safely.
- The dashboard can expose customer health signals without additional privacy
  review.
- The next planning meeting defines the delivery deadline and minimum useful
  scope.
- A simple table or chart view is enough to satisfy "easy to use."
- Historical trends are either required or out of scope.
```

## 사례 3 산출물: B. ni Intent-Lock 경로

체크인된 case artifact는 측정된 pre-runtime ni path를 기록한다.

- original request:
  `examples/benchmark-report/cases/internal-dashboard/01-original-request.md`
- direct risk analysis:
  `examples/benchmark-report/cases/internal-dashboard/02-direct-to-agent-risk.md`
- ni path expectations:
  `examples/benchmark-report/cases/internal-dashboard/03-ni-path.md`
- measurement table:
  `examples/benchmark-report/cases/internal-dashboard/04-measurement-table.md`
- not-measured boundary:
  `examples/benchmark-report/cases/internal-dashboard/05-not-measured.md`
- status proof:
  `examples/benchmark-report/cases/internal-dashboard/06-ni-status-proof.md`
- next-question proof:
  `examples/benchmark-report/cases/internal-dashboard/07-ni-next-questions.md`
- blocker analysis:
  `examples/benchmark-report/cases/internal-dashboard/08-blocker-analysis.md`
- resolution path:
  `examples/benchmark-report/cases/internal-dashboard/09-resolution-path.md`
- answer packet:
  `examples/benchmark-report/cases/internal-dashboard/10-answer-packet.md`
- planning workspace:
  `examples/benchmark-report/cases/internal-dashboard/workspace/`

픽스처의 suggested ni-start questions:

```text
- Who are the primary users, and what decision should the dashboard help them
  make?
- Which account signals and source systems are allowed for the first version?
- What does "needs attention" mean in observable terms?
- What acceptance checks must pass before the planning meeting?
- Which dashboard behaviors are explicitly out of scope for this iteration?
- What privacy, access-control, or data-freshness constraints apply?
```

lock 전 측정된 readiness blocker:

- `OQ-001`: primary dashboard user와 supported decision이 아직 accepted 상태가
  아니다;
- `OQ-002`: "needs attention"이 observable account signal로 정의되지 않았다;
- `OQ-003`: source systems, account fields, freshness rules, privacy
  constraints, access controls가 아직 accepted 상태가 아니다;
- `OQ-004`: planning-meeting acceptance evidence가 아직 accepted 상태가 아니다.

체크인된 docs/contract records:

- `docs/plan/01_actors_outcomes.md`는 requested customer-team actor를 기록하고
  exact role/outcome은 blocker intent로 남긴다.
- `docs/plan/02_capabilities.md`와 `.ni/contract.json`은 request capture와
  readiness blocking을 위한 accepted planning capability를 기록한다.
- `docs/plan/06_risks_security.md`와 `.ni/contract.json`은 blocker를 보존하는
  mitigation을 가진 high-severity risk 3개를 기록한다.
- `docs/plan/07_evaluation_contract.md`는 planning-capture review와
  blocked-readiness proof를 기록한다.
- `docs/plan/08_delivery_operation.md`는 web surface를 기록하고 이 benchmark
  case에서 dashboard delivery가 허가되지 않았다고 명시한다.
- `docs/plan/10_open_questions.md`는 blocker open question 4개를 유지한다.

2026-05-29에 확인한 `ni status` output excerpt:

```text
$ go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
NI Intent Readiness: BLOCKED

Blockers:
- OQ-001 is marked as blocker.
- OQ-002 is marked as blocker.
- OQ-003 is marked as blocker.
- OQ-004 is marked as blocker.

Passed checks:
- Required docs exist.
- Contract JSON is valid.
- Readiness profile definitions are valid.
- Capability and evaluation traceability rules passed.
- High-severity risks have mitigation.
- Decision statuses are valid and accepted decisions do not conflict.
- At least one non-goal is recorded.
- Docs and contract are synchronized.

Execution must not start.
```

Full proof와 next-question output은 `06-ni-status-proof.md`와
`07-ni-next-questions.md`에 체크인되어 있다.

Blocker analysis와 resolution path는 `08-blocker-analysis.md`와
`09-resolution-path.md`에 체크인되어 있다. 이 문서들은 각 blocker가 왜 lock을
막는지, later에 어떤 user answer가 필요한지, 어떤 unsafe assumption을 피하는지,
그리고 future resolved variant가 gate를 약화하지 않고 `ni status`, `ni end`,
`ni run`으로 진행할 수 있는 방법을 설명한다.

| Blocker | Required answer | Expected planning update | Unsafe assumption avoided |
| --- | --- | --- | --- |
| `OQ-001` | Primary dashboard user와 supported decision을 확인한다. | Actor/outcome, capability, open-question, contract record를 update한다. | Customer-team role 또는 decision을 추측하지 않는다. |
| `OQ-002` | Observable attention signal, threshold, ordering criteria, review rule을 정의한다. | Capability, domain-state, evaluation, open-question, contract record를 update한다. | Account-health metric이나 ranking formula를 발명하지 않는다. |
| `OQ-003` | Source system, allowed field, freshness, privacy constraint, access control을 확인한다. | Domain-state, constraint, risk, open-question, contract record를 update한다. | Sensitive data exposure나 stale data 허용을 가정하지 않는다. |
| `OQ-004` | Meeting timing, audience, minimum artifact, pass/fail evidence를 확인한다. | Evaluation, delivery-operation, open-question, contract record를 update한다. | Prototype, memo, live dashboard 중 아무거나 충분하다고 판단하지 않는다. |

| Step | Action | Expected result |
| --- | --- | --- |
| 1 | User answers `OQ-001`. | Primary user is explicit. |
| 2 | User answers `OQ-002`. | Attention signals and ranking criteria are explicit. |
| 3 | User answers `OQ-003`. | Source systems, privacy, and access constraints are explicit. |
| 4 | User answers `OQ-004`. | Meeting acceptance evidence is explicit. |
| 5 | Run `ni status`. | May become `READY` or `READY_WITH_DEFERRALS` if no new blockers appear. |
| 6 | Run `ni end` only after confirmation. | Lock may be created. |
| 7 | Run `ni run` only after lock. | Bounded prompt may be compiled. |

## 사례 3 수동 측정표

이 표는 dashboard request에 대한 한 리뷰어의 수동 정성 평가다. 반복
benchmark data가 아니다. ni path evidence는 blocked pre-runtime workspace에
대한 실제 status proof이며, 완료된 lock/run 측정이 아니다.

| Criterion | Direct-to-agent risk | ni-path evidence | Improved? | Evidence |
| --- | --- | --- | --- | --- |
| Missing acceptance criteria | account health, priority ranking, freshness, performance, usability, meeting acceptance의 pass/fail check가 빠져 있다. | `OQ-002`, `OQ-003`, `OQ-004`가 signal definition, data freshness, meeting evidence를 lock 전 blocker로 유지한다. | yes | `workspace/docs/plan/10_open_questions.md`; `06-ni-status-proof.md` |
| Unmitigated high-risk items | customer data exposure, incorrect prioritization, stale account signal risk가 보이지만 mitigation이 없다. | `RISK-001`부터 `RISK-003`은 high severity이고 mitigation을 갖는다. `ni status`도 high-severity risks have mitigation을 보고한다. | yes | `workspace/docs/plan/06_risks_security.md`; `workspace/.ni/contract.json`; `06-ni-status-proof.md` |
| Unresolved blockers | primary users, source systems, required fields, meeting date, launch surface가 unknown이다. | `ni status --proof --next-questions`가 four blocker open questions와 `Execution must not start`를 보고한다. | yes | `06-ni-status-proof.md`; `07-ni-next-questions.md` |
| Hidden assumptions | users, metrics, source systems, privacy review, deadline, visualization format을 downstream actor가 발명해야 한다. | workspace는 그 assumption들을 accepted dashboard scope가 아니라 blocker questions 또는 non-goals로 기록한다. | yes | `workspace/docs/plan/01_actors_outcomes.md`; `workspace/docs/plan/10_open_questions.md`; `workspace/.ni/contract.json` |
| Non-goal coverage | CRM replacement, workflow automation, forecasting, write-back behavior, downstream agents, live integration 제외가 없다. | `NG-001`부터 `NG-003`은 dashboard implementation, live customer-system integration, CRM write-back, runtime state, downstream agents, model APIs, queues, PR/release automation을 제외한다. | yes | `workspace/.ni/contract.json`; `workspace/docs/plan/08_delivery_operation.md` |
| Delivery surface clarity | web dashboard라고 가정하지만 prototype, report, embedded CRM view, planning document가 구분되지 않는다. | workspace는 requested surface를 `web`으로 기록하고 `OQ-004`로 실제 meeting evidence와 deliverable readiness를 계속 block한다. | yes | `workspace/docs/plan/00_project_brief.md`; `workspace/docs/plan/08_delivery_operation.md`; `06-ni-status-proof.md` |
| Actor/outcome clarity | "Customer team"과 "who needs attention"은 implementation을 이끌기엔 너무 넓다. | `OQ-001`과 `OQ-002`가 actor, decision, attention signals가 accepted 되기 전 readiness를 block한다. | yes | `workspace/docs/plan/01_actors_outcomes.md`; `workspace/docs/plan/10_open_questions.md`; `07-ni-next-questions.md` |
| Evaluation evidence clarity | correctness, freshness, access, meeting readiness evidence가 없다. | `EVAL-001`과 `EVAL-002`는 planning capture와 blocked-readiness proof를 다루며 product evidence는 `OQ-002`부터 `OQ-004`로 계속 block된다. | yes | `workspace/docs/plan/07_evaluation_contract.md`; `06-ni-status-proof.md` |
| Bounded handoff prompt availability | 없음; 직접 프롬프트에는 lock 검증된 compiled target prompt가 없다. | `ni status`가 `BLOCKED`이므로 bounded prompt는 compile하지 않았다. prompt character count는 `not_measured`로 남는다. | no | `06-ni-status-proof.md`; no `08-ni-lock-summary.md`; no `09-ni-run-prompt-summary.md` |

## 사례 3에서 개선된 것

개선은 dashboard가 설계되거나 구현되었다는 뜻이 아니다. 개선은 benchmark가
왜 실행을 기다려야 하는지 드러낸다는 뜻이다. 직접 요청은 users, outcomes,
data boundaries, risks, evaluation evidence, non-goals를 숨긴다. ni path는
그 항목들을 synchronized docs/contract records로 만들고, high risks를 blocker
preservation으로 mitigate하며, plan을 blocked 상태로 남겼다. 따라서
`BLOCKED`는 유용한 result다. ni는 premature handoff를 막고 readiness gap을
명시적으로 만들었지만 implementation quality를 증명하지는 않았다.

## 사례 3에서 측정하지 않은 것

이 사례는 lockfile creation, compiled prompt availability, prompt character
count, agent behavior, dashboard quality, development time, user adoption,
reduced rework, statistical effect를 측정하지 않았다. Cost, latency,
downstream agent performance도 측정하지 않았다. `ni end`, `ni run`, model API,
dashboard implementation, downstream agent를 실행하지 않았다.

## 사례 3 실행하지 않는 경계

Dashboard case는 pre-runtime benchmark evidence로만 남는다. Runtime demo,
shell adapter, dashboard scaffold, queue, telemetry collector, downstream agent
harness가 되면 안 된다. 이 사례의 역할은 누군가 dashboard를 만들기 전에
intent를 compile해야 하는 이유를 보여주는 것이다.

## 관찰

세 직접 프롬프트는 그럴듯하지만 다운스트림 액터에게 넘길 준비가 되지
않았다. 환불 사례에서는 다운스트림 액터가 시작 전에 환불 권한 경계,
정책 출처, escalation 동작, 평가 증거, non-goal을 스스로 발명해야 한다.
research-protocol 사례에서는 sampling, consent, field safety, ethics,
evidence, review boundary를 스스로 발명해야 한다. dashboard 사례에서는
users, account signals, data boundaries, prioritization criteria, delivery
surface, meeting acceptance, non-goals를 스스로 발명해야 한다.

ni 경로가 프로덕션 작업을 완성한 것은 아니다. 실행 전에 의도를 감사
가능하게 만들었다. 환불 사례의 readiness check는 authoritative policy
source에서 멈춘 뒤 `READY_WITH_DEFERRALS`에 도달했고, research-protocol
사례는 `READY`에 도달했다. 두 accepted plan은 명시적인 requirement,
high-risk mitigation, evaluation, non-goal을 포함했으며, 두 target
handoff prompt 모두 설정된 4,000자 한도 안에 있었다. dashboard case는 실제
blocked status proof를 추가한다. docs와 contract는 synchronized이고, high
risks는 mitigation을 가지며, non-goals는 explicit하지만 blocker question 4개가
open 상태라 execution은 멈춘다. lock, run, prompt character count는
`not_measured`로 남긴다.

## 한계

이 사례 연구 보고서는 의도적으로 좁다. 모든 ni plan이 개선된다는 것,
다운스트림 구현이 성공한다는 것, 또는 어떤 프로세스가 통계적으로 더
낫다는 것을 증명하지 않는다. 이 문서는 Intent Lock Protocol이 두 개의
모호한 prompt를 다운스트림 실행 없이 bounded, lock-verified handoff로
바꾸는 투명한 before/after 사례와, fake lock/run evidence를 만들지 않고
`BLOCKED`에서 멈추는 measured dashboard readiness case를 보여준다.
