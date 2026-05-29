# Conversation Authoring Fixture

## 1. 목적

이 fixture는 `ni init` 이후의 지속적인 planning loop를 보여준다. 모델과
사용자는 대화로 planning docs를 작성하지만, readiness, lock, prompt compiler
권한은 CLI에 남아 있다.

또한 historical locked fixture material만으로 fresh ready state를 주장할 수
없다는 점을 보여준다. 현재 planning docs와 contract는 여전히 `ni status`를
통과해야 한다.

## 2. 증명하는 것

- 사용자는 `contract add`, `contract set`, `contract list` 같은 authoring
  command를 입력할 필요가 없다.
- 모델은 `docs/plan/**`와 `.ni/contract.json`을 함께 갱신한 뒤 `ni status`로
  readiness를 검증해야 한다.
- `ni status`는 historical lockfile과 generated prompt가 있어도 stale
  docs/contract synchronization을 잡아낼 수 있다.
- 체크인된 `ni run` material은 inert handoff seed material이며 downstream
  execution이 아니다.

## 3. 제품 유형 / 표면

- `product_type`: `conversation_product`
- `delivery_surface`: `conversation`, `document`
- 예상 `ni status`: `BLOCKED`
- 예상 `ni run` 대상: existing lock에서만 `human-team`

## 4. 파일

- `transcript.md`: model-user authoring loop와 status checks.
- `ni-end-confirmation.md`: lock 전 confirmation behavior.
- `ni-run-handoff.md`: target selection, stale-lock refusal, prompt
  compilation behavior.
- `session-resume.md`: contract authority 아래의 bounded session resume.
- `docs/plan/**`: human-facing plan docs.
- `.ni/contract.json`: machine-readable planning contract.
- `.ni/session.json`: docs와 contract 아래의 bounded resume state.
- `.ni/plan.lock.json`: historical CLI-written lockfile.
- `generated/human-team.prompt.txt`: existing lock에서 컴파일된 handoff prompt.

## 5. 명령

Repository root에서:

```bash
go run ./cmd/ni status --dir examples/conversation-authoring
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/conversation-authoring --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.txt"
wc -m "$tmpdir/human-team.prompt.txt"
rm -rf "$tmpdir"
```

## 6. 예상 출력

예상 상태: `BLOCKED`.

상태 명령은 다음으로 시작해야 한다.

```text
BLOCKED
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document
```

현재 docs/contract synchronization blocker도 보여야 한다.

```text
blocker R012
```

`ni run`은 existing lockfile에서 prompt를 컴파일할 수 있지만, `ni status`가
다시 통과하기 전에는 이 fixture를 fresh ready 상태로 설명하면 안 된다.

## 7. 실행하지 않는 경계

이 fixture는 support assistant를 실행하지 않고, customer contact, refund
approval, model API 호출, downstream tool invocation도 하지 않는다. `ni status`가
`BLOCKED`를 보고하는 동안 이 fixture에 `ni end`를 호출하지 않는다. ni는
planning state를 검증하고 existing lock에서 bounded prompt material만
컴파일한다.
