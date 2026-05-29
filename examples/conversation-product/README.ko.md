# Travel Concierge Triage

## 1. 목적

이 locked 예시는 app이 아니라 conversation product를 위한 Project Intent
Compiler 사용을 보여준다. 제품은 human concierge를 위해 trip intent,
constraints, risks, open questions를 모으는 human-operated travel intake
conversation이다.

## 2. 증명하는 것

- ni는 runtime이 없어도 conversation product intent를 lock할 수 있다.
- `product_type=conversation_product`는 ni를 chatbot runner로 바꾸지 않고
  planning guidance만 바꾼다.
- delivery surface는 conversation-first일 수 있다.
- conversation-driven authoring은 conversation surface를 lock 가능한 계획으로
  만들 수 있다.
- `ni run`은 human team을 위한 bounded handoff prompt를 컴파일할 수 있다.

## 3. 제품 유형 / 표면

- `product_type`: `conversation_product`
- `delivery_surface`: `conversation`
- 예상 `ni status`: `READY`
- 예상 `ni run` 대상: `human-team`

## 4. 파일

- `docs/plan/**`: conversation intent를 위한 locked planning docs.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, decisions.
- `.ni/plan.lock.json`: authoritative planning files의 hash를 담은 CLI lock.
- `generated/human-team.prompt.md`: 체크인된 human-team handoff.
- `generated/codex.prompt.txt`: 체크인된 Codex target prompt.
- `contract-summary.md`: locked contract의 압축 요약.

## 5. 명령

Repository root에서:

```bash
go run ./cmd/ni status --dir examples/conversation-product
go run ./cmd/ni status --dir examples/conversation-product --proof --next-questions
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/conversation-product --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.md"
wc -m "$tmpdir/human-team.prompt.md"
rm -rf "$tmpdir"
```

## 6. 예상 출력

예상 상태: `READY`.

상태 명령은 다음으로 시작해야 한다.

```text
READY
profile: prototype
product type: conversation_product
delivery surfaces: conversation
```

`ni run`은 4000자 이하의 비어 있지 않은 prompt를 써야 한다.

## 7. demo-check coverage

`bash scripts/demo-check.sh`가 이 예시를 검증한다.

demo check는 `READY` status를 확인하고, existing lock에서 `human-team` prompt를
컴파일하며, downstream export package가 seed-only 상태인지 검증한다.

## 8. Korean companion

Korean companion docs: `README.ko.md`.

## 9. 실행하지 않는 경계

이 예시는 chatbot을 구현하거나 service를 배포하지 않는다. queue를 추가하지
않고, vendor contact, travel booking, model API 호출, regulated advice claim도
하지 않는다. ni는 locked contract를 검증하고 prompt material만 컴파일한다.
