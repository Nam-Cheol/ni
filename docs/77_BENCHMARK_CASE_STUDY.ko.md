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

사례 3은 이제 before/after artifact-readiness package를 포함한다. Original
`BLOCKED` dashboard request proof를 보존하고, 답변이 적용된 benchmark-artifact
variant가 `READY`에 도달해 isolated workspace를 lock하고 bounded prompt를
compile한 증거를 기록한다. 이 `READY` claim은 benchmark planning-meeting
artifact readiness에만 적용된다. Dashboard product readiness와 product outcome은
`not_measured`로 남는다.

Research-protocol benchmark package도 resolved synthetic-fixture variant를
포함한다. Original `BLOCKED` proof를 보존하고, isolated workspace 안에서만
`OQ-001`부터 `OQ-005`까지 clearly labeled synthetic answer를 적용해 `READY`에
도달한 뒤 isolated workspace를 lock하고 4000자 bounded prompt를 compile한다. 이
`READY` claim은 benchmark fixture readiness에만 적용되며 real fieldwork approval
또는 research quality가 아니다.

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
  - B. `BLOCKED`에서 answered artifact readiness, isolated lock, bounded
    prompt로 이어지는 before/after ni readiness 경로

세 번째 사례는 software dashboard request에 대한 수동 정성 before/after
readiness package다. Internal dashboard는 downstream actor가 너무 일찍 만들기
쉬운 예시다. web surface가 뻔해 보여도 users, decisions, account signals,
privacy constraints, data freshness, success criteria, risks, non-goals가 빠져
있을 수 있다.

이 case는 original `BLOCKED` proof와 이후 answered artifact-readiness proof를
함께 보존한다. 이후 proof는 isolated benchmark workspace만 lock했고 bounded
prompt만 compile했다. Dashboard를 구현하지 않았고, generated prompt를 실행하지
않았으며, downstream agent를 호출하거나 dashboard product readiness를 주장하지
않았다.

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
- resolved status proof:
  `examples/benchmark-report/cases/internal-dashboard/11-resolved-status-proof.md`
- resolved next-question proof:
  `examples/benchmark-report/cases/internal-dashboard/12-resolved-next-questions.md`
- lock summary:
  `examples/benchmark-report/cases/internal-dashboard/13-lock-summary.md`
- bounded prompt summary:
  `examples/benchmark-report/cases/internal-dashboard/14-bounded-prompt-summary.md`
- before/after evidence:
  `examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.md`
- lessons:
  `examples/benchmark-report/cases/internal-dashboard/16-lessons.md`
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

답변 적용 전 측정된 readiness blocker:

- `OQ-001`: primary dashboard user와 supported decision이 accepted 상태가
  아니었다;
- `OQ-002`: "needs attention"이 observable account signal로 정의되지 않았다;
- `OQ-003`: source systems, account fields, freshness rules, privacy
  constraints, access controls가 accepted 상태가 아니었다;
- `OQ-004`: planning-meeting acceptance evidence가 accepted 상태가 아니었다.

사용자 답변 적용 후 측정된 readiness:

- `OQ-001`부터 `OQ-004`는 benchmark planning-meeting artifact readiness에
  대해서만 resolved 상태다.
- accepted delivery surface는 dashboard delivery가 아니라 `document`다.
- `ni status --proof --next-questions`는 blocker와 deferral 없이 `READY`를
  보고한다.
- `ni end`는 isolated workspace lock
  `examples/benchmark-report/cases/internal-dashboard/workspace/.ni/plan.lock.json`을
  만들었다.
- `ni run --max-chars 4000`은 4000자 generic prompt를 compile했다.
- 이 case는 여전히 dashboard implementation quality, production readiness,
  downstream agent performance, rework reduction, adoption, cost, latency,
  statistical effect size를 claim하지 않는다.

체크인된 docs/contract records:

- `docs/plan/01_actors_outcomes.md`는 planning owner, product lead, internal
  operations lead를 primary benchmark artifact user로 기록한다.
- `docs/plan/02_capabilities.md`와 `.ni/contract.json`은 artifact readiness,
  answer-packet review, non-execution boundary, lock 후 bounded prompt compile
  capability를 기록한다.
- `docs/plan/06_risks_security.md`와 `.ni/contract.json`은 false
  product-readiness claim을 포함한 high-severity risk 4개와 mitigation을
  기록한다.
- `docs/plan/07_evaluation_contract.md`는 scope review, answer-packet
  completeness, privacy/freshness review, isolated CLI readiness/prompt proof를
  기록한다.
- `docs/plan/08_delivery_operation.md`는 document surface를 기록하고 이
  benchmark case에서 dashboard delivery가 허가되지 않았다고 명시한다.
- `docs/plan/10_open_questions.md`는 `OQ-001`부터 `OQ-004`를 benchmark artifact
  readiness에 대해 resolved/non-blocking으로 표시한다.

2026-05-29에 확인한 original `ni status` output excerpt:

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

2026-05-29에 확인한 resolved `ni status` output excerpt:

```text
$ go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
NI Intent Readiness: READY

Blockers:
- None.

Deferrals:
- None.

Warnings:
- None.

Passed checks:
- Required docs exist.
- Contract JSON is valid.
- Readiness profile definitions are valid.
- Capability and evaluation traceability rules passed.
- High-severity risks have mitigation.
- Decision statuses are valid and accepted decisions do not conflict.
- No blocker open questions are present.
- At least one non-goal is recorded.
- Docs and contract are synchronized.

Execution may proceed only after lock.
```

Resolved proof, next-question disposition, lock summary, bounded prompt
summary는 `11-resolved-status-proof.md`, `12-resolved-next-questions.md`,
`13-lock-summary.md`, `14-bounded-prompt-summary.md`에 체크인되어 있다.

Before/after evidence package와 lessons는 `15-before-after-evidence.md`와
`16-lessons.md`에 체크인되어 있다.

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

이 표는 user answer를 적용한 뒤 dashboard request에 대한 한 리뷰어의 수동
정성 평가다. 반복 benchmark data가 아니다. ni path evidence는 benchmark
artifact readiness에 대한 실제 status, lock, prompt proof이며 dashboard product
readiness가 아니다.

| Criterion | Direct-to-agent risk | ni-path evidence | Improved? | Evidence |
| --- | --- | --- | --- | --- |
| Missing acceptance criteria | account health, priority ranking, freshness, performance, usability, meeting acceptance의 pass/fail check가 빠져 있다. | resolved workspace는 benchmark planning-meeting artifact readiness에 한해 required OQ fields complete, supported decision clear, testable criteria, explicit privacy boundary, unresolved blockers visible을 pass/fail criteria로 정의한다. | yes | `workspace/docs/plan/04_domain_state.md`; `workspace/docs/plan/07_evaluation_contract.md`; `11-resolved-status-proof.md` |
| Unmitigated high-risk items | customer data exposure, incorrect prioritization, stale signal, false product-readiness claim risk가 보이지만 mitigation이 없다. | `RISK-001`부터 `RISK-004`는 high severity이고 mitigation을 갖는다. `ni status`도 high-severity risks have mitigation을 보고한다. | yes | `workspace/docs/plan/06_risks_security.md`; `workspace/.ni/contract.json`; `11-resolved-status-proof.md` |
| Unresolved blockers | primary users, source systems, required fields, meeting date, launch surface가 unknown이다. | `OQ-001`부터 `OQ-004`는 benchmark artifact readiness에 대해 resolved이고 `ni status --proof --next-questions`는 blocker/deferral 없이 `READY`를 보고한다. | yes | `workspace/docs/plan/10_open_questions.md`; `11-resolved-status-proof.md`; `12-resolved-next-questions.md` |
| Hidden assumptions | users, metrics, source systems, privacy review, deadline, visualization format을 downstream actor가 발명해야 한다. | workspace는 scope shift를 `DEC-003`으로 기록하고 artifact answer를 product answer로 취급하지 않는다. Dashboard product readiness, implementation quality, empirical impact는 out of scope다. | yes | `workspace/docs/plan/11_decision_log.md`; `workspace/docs/plan/05_constraints.md`; `workspace/.ni/contract.json` |
| Non-goal coverage | CRM replacement, workflow automation, forecasting, write-back behavior, downstream agents, live integration 제외가 없다. | `NG-001`부터 `NG-004`는 dashboard implementation, live integration, runtime state, downstream agents, model APIs, automation, release work, dashboard product readiness, empirical impact claim을 제외한다. | yes | `workspace/.ni/contract.json`; `workspace/docs/plan/05_constraints.md`; `workspace/docs/plan/08_delivery_operation.md` |
| Delivery surface clarity | web dashboard라고 가정하지만 prototype, report, embedded CRM view, planning document가 구분되지 않는다. | resolved workspace는 accepted delivery surface를 isolated benchmark planning workspace/report evidence인 `document`로 바꾸고 web dashboard delivery는 unauthorized로 유지한다. | yes | `workspace/docs/plan/00_project_brief.md`; `workspace/docs/plan/08_delivery_operation.md`; `11-resolved-status-proof.md` |
| Actor/outcome clarity | "Customer team"과 "who needs attention"은 implementation을 이끌기엔 너무 넓다. | Actor/outcome clarity는 benchmark artifact audience와 supported decision에 대해 accepted이고 customer-dashboard product actor/outcome은 readiness claim 밖에 남는다. | yes | `workspace/docs/plan/01_actors_outcomes.md`; `workspace/docs/plan/10_open_questions.md`; `12-resolved-next-questions.md` |
| Evaluation evidence clarity | correctness, freshness, access, meeting readiness evidence가 없다. | `EVAL-001`부터 `EVAL-004`는 scope review, answer packet completeness, privacy/freshness boundary review, status proof, lock proof, bounded prompt proof를 다룬다. | yes | `workspace/docs/plan/07_evaluation_contract.md`; `11-resolved-status-proof.md`; `13-lock-summary.md`; `14-bounded-prompt-summary.md` |
| Bounded handoff prompt availability | 없음; 직접 프롬프트에는 lock 검증된 compiled target prompt가 없다. | isolated workspace는 `ni end`로 locked 되었고 `ni run --max-chars 4000`이 4000자 generic prompt를 compile했다. | yes | `13-lock-summary.md`; `14-bounded-prompt-summary.md` |

## 사례 3에서 개선된 것

개선은 dashboard가 설계되거나 구현되었다는 뜻이 아니다. 개선은 benchmark가
먼저 왜 실행을 기다려야 하는지 드러낸 뒤, user answer를 artifact-readiness
boundary에서만 accepted했다는 뜻이다. 직접 요청은 users, outcomes, data
boundaries, risks, evaluation evidence, non-goals를 숨긴다. ni path는 그
항목들을 synchronized docs/contract records로 만들고 scope shift를 기록했으며,
CLI가 `READY`를 보고한 뒤에만 lock하고 bounded prompt seed material을
compile했다.

## 사례 3에서 측정하지 않은 것

이 사례는 CLI가 `READY`를 보고한 뒤 isolated lockfile creation, compiled
prompt availability, prompt character count를 측정했다. Agent behavior,
dashboard quality, development time, user adoption, reduced rework, cost,
latency, downstream agent performance, statistical effect는 측정하지 않았다.
Model API, dashboard implementation, downstream agent는 실행하지 않았다.

## 사례 3 실행하지 않는 경계

Dashboard case는 pre-runtime benchmark evidence로만 남는다. Runtime demo,
shell adapter, dashboard scaffold, queue, telemetry collector, downstream agent
harness가 되면 안 된다. 이 사례의 역할은 누군가 dashboard를 만들기 전에
intent를 compile해야 하는 이유를 보여주는 것이다.

## Research-Protocol Benchmark Package

선택된 두 번째 v0.5 benchmark case는
`examples/benchmark-report/cases/research-protocol/`에 체크인되어 있다. 이 case는
다음 vague request의 initial readiness 측정을 보존한다.

```text
Help us plan a summer neighborhood cooling study so we can decide where to
place shade and cooling interventions.
```

Initial measured status는 `BLOCKED`였다. Isolated workspace는
`product_type=research_protocol`과 delivery surfaces `document`, `workflow`,
`human_service`를 기록한다. `ni status --proof --next-questions`는 research
question, participant 또는 observation scope, consent/privacy/data/accessibility
boundary, heat/weather field safety, vulnerable-group safeguard, review owner,
acceptance evidence를 다루는 `OQ-001`부터 `OQ-005`를 open blocker로 보고한다.

Resolved variant는 이 다섯 blocker에 synthetic benchmark fixture answer를
적용한다. 이 answer는 real fieldwork approval, actual research authorization,
proof of research quality, empirical evidence가 아니다. 이 answer를
`docs/plan/**`, `.ni/contract.json`, `.ni/session.json`에 반영한 뒤
`ni status --proof --next-questions`는 `READY`를 보고했고, `ni end`는
`workspace/.ni/plan.lock.json`을 만들었으며, `ni run --max-chars 4000`은 4000자
bounded prompt를 compile했다.

이 case에는 blocker analysis, resolution path, synthetic answer packet, resolved
status proof, lock summary, bounded prompt summary가 포함된다.

- `examples/benchmark-report/cases/research-protocol/08-blocker-analysis.md`
- `examples/benchmark-report/cases/research-protocol/09-resolution-path.md`
- `examples/benchmark-report/cases/research-protocol/10-answer-packet.md`
- `examples/benchmark-report/cases/research-protocol/11-resolved-status-proof.md`
- `examples/benchmark-report/cases/research-protocol/12-resolved-next-questions.md`
- `examples/benchmark-report/cases/research-protocol/13-lock-summary.md`
- `examples/benchmark-report/cases/research-protocol/14-bounded-prompt-summary.md`

이 문서들은 각 blocker가 왜 중요한지, later user answer가 어떤 방식으로
resolve할 수 있는지, 어떤 unsafe assumption을 피하는지, resolved synthetic
fixture가 isolated workspace만 update한 뒤 허용된 순서로 `ni status`, `ni end`,
`ni run`을 실행했음을 설명한다.

이 research-protocol case는 research protocol quality, fieldwork readiness,
intervention decision readiness, participant outcome, adoption, rework
reduction, cost, latency, downstream agent performance, statistical effect를
주장하지 않는다.

## 관찰

세 직접 프롬프트는 그럴듯하지만 다운스트림 액터에게 넘길 준비가 되지
않았다. 환불 사례에서는 다운스트림 액터가 시작 전에 환불 권한 경계,
정책 출처, escalation 동작, 평가 증거, non-goal을 스스로 발명해야 한다.
research-protocol 사례에서는 sampling, consent, field safety, ethics,
evidence, review boundary를 스스로 발명해야 한다. dashboard 사례에서는
users, account signals, data boundaries, prioritization criteria, delivery
surface, meeting acceptance, non-goals를 스스로 발명해야 한다.

ni 경로가 프로덕션 작업을 완성한 것은 아니다. 실행 전에 의도를 감사 가능하게
만들었다. 환불 사례의 readiness check는 authoritative policy source에서 멈춘 뒤
`READY_WITH_DEFERRALS`에 도달했고, original research-protocol package는 다섯
research blocker가 답변될 때까지 `BLOCKED`에서 멈췄다. Resolved
research-protocol fixture는 `READY`에 도달했고, accepted plan은 명시적인
requirement, high-risk mitigation, evaluation, non-goal을 포함했다. Target
handoff prompt는 설정된 4,000자 한도 안에 있었다. Dashboard와 research
benchmark package는 모두 `BLOCKED`에서 isolated workspace `READY` lock 및
4000자 prompt로 가는 transition proof를 포함하지만, product readiness,
implementation quality, downstream agent performance, user impact, adoption,
rework reduction, cost, latency, statistical effect, real research approval,
fieldwork readiness는 `not_measured`로 남긴다.

## 한계

이 사례 연구 보고서는 의도적으로 좁다. 모든 ni plan이 개선된다는 것,
다운스트림 구현이 성공한다는 것, 또는 어떤 프로세스가 통계적으로 더
낫다는 것을 증명하지 않는다. 이 문서는 Intent Lock Protocol이 unclear intent를
드러내고, 충분한 답변이 있을 때 downstream work를 실행하지 않은 채 bounded,
lock-verified handoff seed material을 만들 수 있음을 보여주는 세 개의 투명한
readiness case를 기록한다. Dashboard case는 artifact-readiness evidence일 뿐
dashboard product가 ready임을 증명하지 않는다. Research-protocol resolved case는
synthetic fixture readiness evidence일 뿐 real fieldwork나 research approval이
ready임을 증명하지 않는다.
