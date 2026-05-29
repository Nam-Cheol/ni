# No-terminal assisted draft

## 1. 목적

이 example은 CLI를 설치하기 전에 model workspace에서 `ni` 형태의 Project Intent
Compiler plan을 시작하는 방식을 보여준다. 의도적으로 draft 상태다. Lockfile,
hash proof, compiled handoff prompt는 없다.

## 2. 증명하는 것

- local CLI access가 없어도 assisted planning은 docs와 draft contract를 캡처할
  수 있다.
- model judgment는 readiness, lock, hash authority가 아니다.
- handoff나 downstream seed generation 전에 full CLI validation으로 넘어가야
  한다.
- 이 예시는 deterministic validation claim을 하지 않는다.

## 3. 제품 유형 / 표면

- `product_type`: draft `workflow`
- `delivery_surface`: draft `document`
- 예상 `ni status`: teammate, CI job, local CLI setup이 command를 실행하기
  전까지 claim하지 않는다.
- 예상 `ni run` 대상: 해당 없음

## 4. 파일

- `docs/plan/00_project_brief.md`: 사람이 읽는 planning notes.
- `.ni/contract.json`: docs와 맞춰 둔 draft contract.
- Accepted decisions로 취급하지 않는 assumptions와 blocker questions.

## 5. 명령

이 예시는 의도적으로 docs-only다. Repository root에서:

```bash
test -f examples/no-terminal-assisted/README.md
test -f examples/no-terminal-assisted/README.ko.md
test -f examples/no-terminal-assisted/docs/plan/00_project_brief.md
test -f examples/no-terminal-assisted/.ni/contract.json
```

## 6. 예상 출력

`test` 명령은 성공해야 한다. Trusted CLI run을 따로 인용할 수 있기 전에는 이
draft를 `READY`, `READY_WITH_DEFERRALS`, `BLOCKED`로 설명하지 않는다.

## 7. demo-check coverage

`bash scripts/demo-check.sh`가 docs-only 예시로 검증한다.

demo check는 required file과 boundary wording을 확인한다. 의도적인 assisted
draft이므로 `ni status`, `ni end`, `ni run`을 실행하지 않는다.

## 8. Korean companion

Korean companion docs: `README.ko.md`.

## 9. No-terminal checklist

- Model pack 또는 copied instructions로 시작한다.
- `docs/plan` draft를 만든다.
- Docs와 함께 `.ni/contract.json`을 draft한다.
- Assumptions와 open questions를 표시한다, especially blockers.
- 나중에 CLI, teammate, trusted runner로 validate한다.
- Model judgment를 lock으로 취급하지 않는다.

## 10. Handoff 전에 full ni로 넘어가기

이 draft가 implementation 또는 downstream seed generation을 안내하기 전에는 full
`ni`를 사용한다. Readiness, lock creation, hash verification, prompt compilation은
CLI가 `ni status`, `ni end`, `ni run`으로 만들어야 한다.

## 11. 실행하지 않는 경계

이 example은 web service, model API call, runtime execution, shell adapter,
queue, model-authoritative skill을 추가하지 않는다.
