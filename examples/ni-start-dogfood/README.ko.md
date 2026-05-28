# ni-start 도그푸드 대화 기록

## 1. 목적

이 공개 예시는 `ni init` 이후의 의도한 작성 흐름을 보여준다. 사용자는
계획을 대화로 설명하고, 모델은 확인된 답변을 바탕으로 `docs/plan/**`와
`.ni/contract.json`을 함께 갱신하며, CLI가 잠금과 프롬프트 컴파일 전에
준비 상태를 검증한다.

## 2. 증명하는 것

- 기본 작성 UX는 사용자가 contract `add`, `list`, `set` 명령을 입력하는
  방식이 아니라 지속적인 계획 대화다.
- 모델은 요약, 집중 질문, 문서와 계약 기록의 동기화 갱신을 할 수 있다.
- `ni status`가 준비 상태의 권한이며, 모델은 `BLOCKED` 결과를 덮어쓸 수
  없다.
- `ni end`는 CLI 준비 상태와 사용자의 명시적 확인 뒤에만 lockfile을 쓴다.
- `ni run`은 제한된 handoff 프롬프트만 컴파일하며 downstream 실행을 하지
  않는다.
- `ni-start`는 시작 시 purpose, active readiness profile, product type /
  delivery surface, accepted capability, unresolved blocker question, recent
  decision, next planning focus를 요약하고 1개에서 3개의 집중 질문만 묻는다.

## 3. 제품 유형 / 표면

- `product_type`: `conversation_product`
- `delivery_surface`: `conversation`, `document`
- 예상 `ni status`: `READY_WITH_DEFERRALS`
- 예상 `ni run` 대상: `human-team`

## 4. 대화 기록 개요

- `01-init.md`: `ni init`이 계획 workspace를 만든다.
- `02-user-vague-idea.md`: 사용자가 모호한 support assistant 아이디어를
  제시한다.
- `03-model-summary-and-questions.md`: `ni-start`가 gap을 요약하고 집중
  질문을 한다.
- `04-user-answers.md`: 사용자가 scope, non-goal, evidence를 확인한다.
- `05-docs-contract-delta.md`: 모델이 docs와 contract를 함께 갱신한다.
- `06-status-proof.md`: `ni status`가 준비 상태와 blocker를 보고한다.
- `07-second-round-questions.md`: 모델이 gate가 요구하는 다음 blocker
  질문만 묻는다.
- `08-ni-end-confirmation.md`: `ni-end`가 lock 전에 CLI 준비 상태를 확인한다.
- `09-ni-run-handoff.md`: `ni-run`이 lock에서 handoff 프롬프트를 컴파일한다.

## 5. 명령

Repository root에서:

```bash
go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace --proof --next-questions
```

## 6. 실행하지 않는 경계

이 예시는 support assistant를 실행하지 않고, model API를 호출하지 않으며,
Codex 실행, 고객 연락, 환불 승인, adapter, queue를 만들지 않는다. 커널의
conversation-first authoring, readiness proof, lock authority, prompt
compilation을 보여주는 예시다.
