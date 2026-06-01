# Internal dashboard benchmark answer packet

이 packet은 internal-dashboard benchmark blocker에 대한 사용자 답변을 수집하기
위해 만들어진 문서이다. 이 문서만으로 blocker가 해결되지는 않았다. Packet
작성 시점에는 답변이 제공되고, isolated benchmark workspace가 갱신되며,
`ni status`가 lock 가능한 readiness를 보고하기 전까지 benchmark가 `BLOCKED`
상태였다.

이 답변은 다음 위치 안에서만 사용한다.

```text
examples/benchmark-report/cases/internal-dashboard/workspace/
```

repository-root planning lock이나 root `.ni/` state에는 적용하지 않는다.

## Original Status Before Answers

- Readiness: `BLOCKED`
- Lock created: no
- `ni-run` prompt compiled: no
- Required answers: `OQ-001` through `OQ-004`

## Resolved Status After Answers

Task 161은 `OQ-001`부터 `OQ-004`에 대한 사용자 답변을 isolated workspace에
benchmark planning-meeting artifact readiness로 적용했다. 측정된 resolved
status는 `11-resolved-status-proof.md`에, lock과 prompt evidence는
`13-lock-summary.md`와 `14-bounded-prompt-summary.md`에 기록되어 있다.

## How to use this packet

1. 아래 답변란을 채운다.
2. 답변은 benchmark workspace 갱신에만 사용한다.
3. benchmark workspace를 대상으로 `ni status --proof --next-questions`를
   실행한다.
4. readiness가 `READY` 또는 `READY_WITH_DEFERRALS`가 되고 사용자 확인이
   있을 때만 lock한다.
5. 유효한 benchmark workspace lock이 생긴 뒤에만 `ni run`을 실행한다.

## OQ-001 - Primary user and supported decision

Prompt:

주 dashboard 사용자는 누구이며, dashboard는 어떤 결정을 돕기 위한 것인가?

Required answer fields:

- Primary user:
- Secondary users, if any:
- Decision supported:
- Decision timing:
- What should not be supported:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- 어떤 "customer team" 역할이 중요한지 추측하기.
- 모든 customer-facing team이 같은 dashboard를 필요로 한다고 가정하기.

## OQ-002 - Attention signals and ranking criteria

Prompt:

어떤 account가 "need attention" 상태이며, account는 어떤 기준으로 정렬되어야
하는가?

Required answer fields:

- Attention signals:
- Thresholds or review rules:
- Ranking logic:
- Freshness expectations:
- Signals explicitly excluded:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- account-health metric을 지어내기.
- 모든 negative signal을 똑같이 긴급한 것으로 다루기.

## OQ-003 - Source systems, privacy, and access boundaries

Prompt:

어떤 data를 사용할 수 있고, 그 data는 어디에서 오며, 얼마나 최신이어야 하고,
누가 볼 수 있는가?

Required answer fields:

- Source systems:
- Allowed fields:
- Prohibited fields:
- Freshness requirement:
- Access roles:
- Privacy/security constraints:
- Data that must remain out of scope:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- sensitive customer data를 노출하기.
- stale data가 허용된다고 가정하기.
- 모든 internal user가 같은 field를 볼 수 있다고 가정하기.

## OQ-004 - Planning-meeting acceptance evidence

Prompt:

planning meeting에서 결과를 acceptance하기에 충분한 evidence는 무엇인가?

Required answer fields:

- Meeting audience:
- Meeting date or timing:
- Minimum artifact:
- Pass/fail criteria:
- What is explicitly not required:
- Who approves acceptance:

Optional notes:

- Context or rationale:
- Follow-up questions:

Unsafe assumptions to avoid:

- 아무 prototype이나 sufficient하다고 보기.
- planning memo나 mock만으로 충분한 경우에도 live dashboard가 필요하다고 보기.
- approval이 불명확한 상태를 acceptance로 간주하기.

## After answers are provided

Expected next steps:

1. benchmark workspace `docs/plan`을 갱신한다.
2. benchmark workspace `.ni/contract.json`을 갱신한다.
3. benchmark workspace `.ni/session.json`을 갱신한다.
4. `ni status --proof --next-questions`를 실행한다.
5. 여전히 `BLOCKED`이면 남은 blocker를 문서화한다.
6. `READY` 또는 `READY_WITH_DEFERRALS`이면 benchmark workspace 안에서만
   lock한다.
7. lock되면 benchmark workspace에서 bounded prompt를 compile한다.
8. benchmark measurement table을 정직하게 갱신한다.

Rules:

- root `.ni/plan.lock.json`을 편집하지 않는다.
- repository-root `ni end` 또는 `ni relock`을 실행하지 않는다.
- downstream agent를 실행하지 않는다.
- dashboard를 구현하지 않는다.
- prompt 또는 lock evidence를 조작하지 않는다.
