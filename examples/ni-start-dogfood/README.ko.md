# ni-start 도그푸드 대화 기록

## 1. 목적

이 공개 예시는 `ni init` 이후 Project Intent Compiler의 의도한 작성 흐름을
보여준다. 사용자는 계획을 대화로 설명하고, 모델은 확인된 답변을 바탕으로
`docs/plan/**`, `.ni/contract.json`, `.ni/session.json`을 함께 갱신하며, CLI가
잠금과 프롬프트 컴파일 전에 준비 상태를 검증한다.

## 2. 증명하는 것

- 기본 작성 UX는 사용자가 contract `add`, `list`, `set` 명령을 입력하는
  방식이 아니라 지속적인 계획 대화다.
- 모델은 요약, 집중 질문, 문서, 계약, session 기록의 동기화 갱신을 할 수 있다.
- `ni status`가 준비 상태의 권한이며, 모델은 `BLOCKED` 결과를 덮어쓸 수
  없다.
- `ni-start`는 grouped `ni status --proof --next-questions` output을 사용하고,
  group label과 answer shape를 보존하며 highest-priority group을 먼저 묻는다.
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
- `02-user-vague-idea.md`: 사용자가 fresh workspace에서 `ni-start`를
  호출한다.
- `03-model-summary-and-questions.md`: `ni-start`가
  `ni status --proof --next-questions`를 실행하고 first-run blocker
  `R014`/`R015`/`R016`을 확인한 뒤 opening card 질문을 묻는다.
- `04-user-answers.md`: 사용자가 purpose, actor/outcome, delivery surface,
  초기 scope, non-goal, evidence를 답한다.
- `05-docs-contract-delta.md`: 모델이 docs와 contract를 함께 갱신한다.
- `06-status-proof.md`: first-run 답변을 기록한 뒤 `ni status`를 다시
  실행하고 다음 blocker를 보고한다.
- `07-second-round-questions.md`: 모델이 gate가 요구하는 다음 blocker
  질문만 묻는다.
- `08-ni-end-confirmation.md`: `ni-end`가 lock 전에 CLI 준비 상태를 확인한다.
- `09-ni-run-handoff.md`: `ni-run`이 lock에서 handoff 프롬프트를 컴파일한다.

## 5. 명령

Repository root에서:

```bash
go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace
go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace --proof --next-questions
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/ni-start-dogfood/workspace --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.txt"
wc -m "$tmpdir/human-team.prompt.txt"
rm -rf "$tmpdir"
```

## 6. 예상 출력

예상 상태: `READY_WITH_DEFERRALS`.

상태 명령은 다음으로 시작해야 한다.

```text
READY_WITH_DEFERRALS
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document
```

수락된 deferral도 계속 보여야 한다.

```text
deferral D001: DEC-004 is deferred
deferral D002: OQ-002 remains open
```

proof 명령은 grouped handoff deferral question도 보여야 한다.

```text
Next questions:
Handoff deferrals:
```

`ni run`은 4000자 이하의 비어 있지 않은 handoff 프롬프트를 써야 한다.

## 7. demo-check coverage

`bash scripts/demo-check.sh`가 이 예시를 검증한다.

demo check는 `READY_WITH_DEFERRALS`를 확인하고 grouped proof command를 실행하며,
workspace에 이미 CLI-written lock이 있기 때문에 `human-team` prompt만 컴파일한다.

## 8. Korean companion

Korean companion docs: `README.ko.md`.

## 9. 실행하지 않는 경계

이 예시는 support assistant를 실행하지 않고, model API를 호출하지 않으며,
Codex 실행, 고객 연락, 환불 승인, adapter, queue를 만들지 않는다. 커널의
conversation-first authoring, readiness proof, lock authority, prompt
compilation을 보여주는 예시다.
