# 모호한 프롬프트 차단

## 1. 목적

이 예시는 Project Intent Compiler로서 ni의 핵심 효과를 보여준다. 모호한 실행
요청은 agent가 구현을 시작하기 전에 차단된다.

`build me a dashboard for my team` 같은 요청은 실행 가능해 보이지만,
implementation agent가 조용히 만들어낼 수 있는 제품 결정을 숨기고 있다.

## 2. 증명하는 것

- `ni status`가 불완전한 intent의 readiness 권한이다.
- blocker open question은 workspace를 `BLOCKED` 상태로 유지한다.
- 모델은 planning record를 작성할 수 있지만 readiness를 선언할 수 없다.
- Codex target prompt는 여기서 설명용일 뿐이며 blocked workspace에서
  컴파일되거나 실행되지 않는다.
- grouped `ni status --proof --next-questions` output은 모델이 임의의 순서를
  만들지 않고 다음 planning question을 제공한다.

## 3. 제품 유형 / 표면

- `product_type`: `software`
- `delivery_surface`: `web`
- 예상 `ni status`: `BLOCKED`
- 예상 `ni run` 대상: blocked 상태에서는 해당 없음

## 4. 파일

- `01-vague-request.md`: 모호한 사용자 요청.
- `02-direct-to-agent-risk.md`: 직접 구현 경로가 만들 가능성이 큰 가정.
- `03-ni-start-conversation.md`: `ni-start`가 intent를 planning record로
  캡처하는 방식.
- `04-ni-status-blocked.md`: 결정적인 blocked 상태.
- `05-next-questions.md`: lock 전에 필요한 집중 질문.
- `06-user-answers.md`: 모호함을 해소할 수 있는 예시 답변.
- `07-locked-contract-summary.md`: 답변 후 ready 상태 설명.
- `08-codex-target-prompt.md`: lock 이후의 설명용 bounded prompt.
- `workspace/docs/plan/**`: 캡처된 blocked planning docs.
- `workspace/.ni/contract.json`: 대응하는 machine-readable contract.

## 5. 명령

Repository root에서:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --proof --next-questions
```

## 6. 예상 출력

예상 상태: `BLOCKED`.

상태 명령에는 다음이 포함되어야 한다.

```text
BLOCKED
blocker R009: OQ-001 is a blocker open question
blocker R009: OQ-002 is a blocker open question
```

proof 명령은 grouped next question을 보여야 한다.

```text
Next questions:
Open blockers:
```

## 7. demo-check coverage

`bash scripts/demo-check.sh`가 이 예시를 검증한다.

demo check는 blocked workspace에 대해 `ni status`를 실행하고 open blocker의
grouped next question이 렌더링되는지 확인한다. workspace가 의도적으로 blocked
상태이므로 handoff를 컴파일하거나 실행하지 않는다.

## 8. Korean companion

Korean companion docs: `README.ko.md`.

## 9. 실행하지 않는 경계

이 예시는 Codex를 실행하거나 dashboard를 구현하지 않는다. shell adapter,
queue, model API, downstream tool도 사용하지 않는다. docs, contract capture,
readiness blocking, prompt handoff boundary를 검증하는 kernel proof asset이다.
