# 벤치마크 사례 연구: 환불 분류 어시스턴트

이 문서는 수동 정성 사례 연구다. 하나의 모호한 요청과 저장소에 포함된
ni 계획 예제를 `docs/43_BENCHMARK_PROTOCOL.md`의 루브릭으로 비교한다.
반복 실험이 아니며, 외부 모델 API를 호출하지 않고, 다운스트림 작업을
실행하지 않으며, 통계적 유의성을 주장하지 않는다.

## 사례

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

이 사례 연구는 실행 전 의도 준비도를 측정한다.

- 누락된 인수 기준,
- 완화되지 않은 high 위험,
- 해결되지 않은 blocker,
- 숨은 가정,
- non-goal 적용 범위,
- 경계가 정해진 target prompt의 존재 여부.

측정하지 않는 것은 구현 품질, 모델 품질, 다운스트림 에이전트 성능,
사용자 만족도, 런타임 안전성, 비용, 지연 시간, 통계적 효과 크기,
프로덕션 결과다.

## 산출물: A. 직접 에이전트 프롬프트

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

## 산출물: B. ni Intent-Lock 경로

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

## 수동 준비도 비교

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

## 관찰

직접 프롬프트는 그럴듯하지만 구현 액터에게 넘길 준비가 되지 않았다.
다운스트림 액터는 시작 전에 환불 권한 경계, 정책 출처, escalation
동작, 평가 증거, non-goal을 스스로 발명해야 한다.

ni 경로가 프로덕션 작업을 완성한 것은 아니다. 실행 전에 의도를 감사
가능하게 만들었다. 첫 readiness check는 authoritative policy source에서
멈췄고, accepted plan은 명시적인 requirement, high-risk mitigation,
evaluation, non-goal, visible deferral을 포함했으며, target handoff
prompt는 설정된 4,000자 한도 안에 있었다.

## 한계

이 사례 연구는 의도적으로 좁다. 모든 ni plan이 개선된다는 것,
다운스트림 구현이 성공한다는 것, 또는 어떤 프로세스가 통계적으로 더
낫다는 것을 증명하지 않는다. 이 문서는 Intent Lock Protocol이 하나의
모호한 prompt를 다운스트림 실행 없이 bounded, lock-verified handoff로
바꾸는 투명한 before/after 사례 하나를 보여준다.
