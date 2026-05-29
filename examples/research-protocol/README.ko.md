# Neighborhood Cooling Study Protocol

## 1. 목적

이 locked 예시는 software app이 아니라 research protocol을 계획한다. 제품은
fieldwork 전에 street-level cooling intervention을 비교하기 위한 문서화된
field-study method다.

## 2. 증명하는 것

- ni는 software가 아닌 제품 intent도 lock할 수 있다.
- `product_type=research_protocol`은 planning guidance를 바꾸지만 shared
  readiness gate는 바꾸지 않는다.
- delivery surface는 document일 수 있다.
- `ni run`은 data collection이나 analysis 없이 research team을 위한
  bounded handoff prompt를 컴파일할 수 있다.

## 3. 제품 유형 / 표면

- `product_type`: `research_protocol`
- `delivery_surface`: `document`
- 예상 `ni status`: `READY`
- 예상 `ni run` 대상: `human-team`

## 4. 파일

- `docs/plan/**`: protocol intent를 위한 locked planning docs.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, decisions.
- `.ni/plan.lock.json`: authoritative planning files의 hash를 담은 CLI lock.
- `generated/human-team.prompt.md`: 체크인된 human-team handoff.
- `generated/generic.prompt.txt`: 체크인된 generic downstream handoff.
- `contract-summary.md`: locked contract의 압축 요약.

## 5. 명령

Repository root에서:

```bash
go run ./cmd/ni status --dir examples/research-protocol
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/research-protocol --target human-team --max-chars 4000 --out "$tmpdir/human-team.prompt.md"
wc -m "$tmpdir/human-team.prompt.md"
rm -rf "$tmpdir"
```

## 6. 예상 출력

예상 상태: `READY`.

상태 명령은 다음으로 시작해야 한다.

```text
READY
profile: prototype
product type: research_protocol
delivery surfaces: document
```

`ni run`은 4000자 이하의 비어 있지 않은 prompt를 써야 한다.

## 7. 실행하지 않는 경계

이 예시는 study를 수행하거나 participant data를 수집하지 않는다. sensor
deployment, statistics, model API, ethics review 대체도 하지 않는다. ni는
locked planning contract를 검증하고 inert prompt material만 컴파일한다.
