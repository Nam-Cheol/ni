# No-terminal assisted draft

이 example은 CLI를 설치하기 전에 model workspace에서 `ni` 형태의 plan을 시작하는
방식을 보여준다. 의도적으로 draft 상태다. Lockfile, hash proof, compiled handoff
prompt는 없다.

Expected `ni status`: teammate, CI job, local CLI setup이 command를 실행하기
전까지 claim하지 않는다.

## Model이 draft할 수 있는 것

- `docs/plan/00_project_brief.md`: 사람이 읽는 planning notes.
- `.ni/contract.json`: docs와 맞춰 둔 draft contract.
- Accepted decisions로 취급하지 않는 assumptions와 blocker questions.

## No-terminal checklist

- Model pack 또는 copied instructions로 시작한다.
- `docs/plan` draft를 만든다.
- Docs와 함께 `.ni/contract.json`을 draft한다.
- Assumptions와 open questions를 표시한다, especially blockers.
- 나중에 CLI, teammate, trusted runner로 validate한다.
- Model judgment를 lock으로 취급하지 않는다.

## Handoff 전에 full ni로 넘어가기

이 draft가 implementation 또는 downstream seed generation을 안내하기 전에는 full
`ni`를 사용한다. Readiness, lock creation, hash verification, prompt compilation은
CLI가 `ni status`, `ni end`, `ni run`으로 만들어야 한다.

이 example은 web service, model API call, runtime execution, shell adapter,
queue, model-authoritative skill을 추가하지 않는다.
