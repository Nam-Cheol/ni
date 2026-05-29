# No-terminal assisted draft

## 1. 목적

이 example은 `ni`의 no-terminal assisted drafting flow를 보여준다. User는 CLI를
설치하기 전에 model workspace에서 시작하고, model에게 `docs/plan/**`과
`.ni/contract.json`을 draft하게 하며, later CLI validation 전까지 결과를 not
locked 상태로 유지한다.

의도적으로 docs-only다. Lockfile, hash proof, compiled handoff prompt는 없다.

## 2. 짧은 흐름

1. User가 model workspace에서 시작하고 `ni-start` guidance를 load하거나 복사한다.
2. Model이 first-run card를 묻는다:
   - 이 project는 무엇을, 누구를 위해, 왜 바꿔야 하는가?
   - Primary actors는 누구이고 각 actor가 얻어야 할 outcome은 무엇인가?
   - Likely delivery surface는 무엇인가?
3. User가 initial project idea와 clear exclusions를 답한다.
4. Model이 `docs/plan/00_project_brief.md`와 `.ni/contract.json`을 draft한다.
5. Model은 무엇이 바뀌었다고 해석했는지 draft planning proof block으로 보여줄 수
   있지만, 이것은 deterministic validation이 아니다.
6. Draft는 "not locked"로 표시되고 assumptions, non-goals, blocker questions를
   visible하게 유지한다.
7. 나중에 readiness, locking, hash trust, prompt compilation, downstream seed
   generation 전에 CLI validation이 필요하다.

## 3. 증명하는 것

- Local CLI access가 없어도 assisted planning은 docs와 draft contract를 캡처할 수
  있다.
- No-terminal 사용자는 planning을 시작할 수 있지만 CLI validation 없이는 trusted
  lock을 끝낼 수 없다.
- Model judgment는 readiness, lock, hash, prompt compilation authority가 아니다.
- Model workspace packs는 drafting을 guide할 수 있지만 downstream work를 실행하지
  않고 CLI readiness를 override하지 않는다.
- 이 예시는 deterministic validation claim을 하지 않는다.
- No-terminal planning proof는 trusted CLI run이 docs와 contract를 validate하기
  전까지 draft-only다.

## 4. 제품 유형 / 표면

- `product_type`: draft `workflow`
- `delivery_surface`: draft `document`
- 예상 `ni status`: teammate, CI job, local CLI setup이 command를 실행하기 전까지
  claim하지 않는다.
- 예상 `ni run` 대상: 해당 없음

## 5. 파일

- `docs/plan/00_project_brief.md`: 사람이 읽는 planning notes.
- `.ni/contract.json`: docs와 맞춰 둔 draft contract.
- Accepted decisions로 취급하지 않는 assumptions와 blocker questions.

## 6. No-terminal checklist

- `ni-start` guidance를 복사하거나 load한다.
- Project idea를 설명한다.
- Model에게 `docs/plan/**` draft를 만들게 한다.
- Model에게 `.ni/contract.json`을 draft하게 한다.
- Model에게 concise draft planning proof block을 요청한다.
- Uncertain statements는 assumptions 또는 open questions로 표시한다.
- Explicit exclusions는 non-goals로 표시한다.
- Draft를 locked 상태로 취급하지 않는다.
- 나중에 CLI 또는 teammate로 validate한다.

## 7. Team handoff path

1. Terminal access가 없는 user가 model과 draft한다.
2. CLI가 있는 teammate가 `ni status --proof --next-questions`를 실행한다.
3. Teammate가 blockers, proof, questions를 돌려준다.
4. User가 model과 planning을 계속한다.
5. Lock은 deterministic CLI validation이 blockers를 clear하고 user가 accepted plan을
   confirm한 뒤에만 일어난다.

## 8. 명령

이 예시는 의도적으로 docs-only다. Repository root에서:

```bash
test -f examples/no-terminal-assisted/README.md
test -f examples/no-terminal-assisted/README.ko.md
test -f examples/no-terminal-assisted/docs/plan/00_project_brief.md
test -f examples/no-terminal-assisted/.ni/contract.json
```

## 9. 예상 출력

`test` 명령은 성공해야 한다. Trusted CLI run을 따로 인용할 수 있기 전에는 이
draft를 `READY`, `READY_WITH_DEFERRALS`, `BLOCKED`로 설명하지 않는다.

## 10. demo-check coverage

`bash scripts/demo-check.sh`가 docs-only example로 검증한다.

Demo check는 required files와 boundary wording을 확인한다. 이 의도적인 assisted
draft에 대해 `ni status`, `ni end`, `ni run`을 실행하지 않는다.

## 11. Korean companion

Korean companion docs: `README.ko.md`.

## 12. Handoff 전에 full ni로 넘어가기

이 draft가 implementation 또는 downstream seed generation을 안내하기 전에는 full
`ni`를 사용한다:

1. Release binary path 또는 curl installer로 `ni`를 설치한다.
2. `ni status --proof --next-questions`를 실행한다.
3. Docs와 contract draft의 blockers를 해결한다.
4. Readiness와 user confirmation 이후에만 `ni end`를 실행한다.
5. Valid lock이 존재한 뒤에만 `ni run`을 실행한다.

## 13. 실행하지 않는 경계

이 example은 web service, model API call, runtime execution, shell adapter,
queue, model-authoritative skill을 추가하지 않는다.
