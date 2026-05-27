# ni에 기여하기

`ni` 개선에 관심을 가져주셔서 고맙습니다.

`ni`는 AI agent를 위한 Project Intent Compiler다. 현재 제품은
`ni-kernel`이며, downstream execution이 시작되기 전에 planning contract를
생성하고, 검증하고, 잠그고, 컴파일하는 deterministic pre-runtime layer다.

`ni`는 execution runtime이 아니다. `ni`는 일을 실행하는 도구로 커지는 것이
아니라 Intent Lock Protocol을 보호해야 한다.

## Product Boundary

kernel이 소유하는 것은 다음이다:

- planning docs와 `.ni/contract.json` sync,
- deterministic readiness validation,
- `.ni/plan.lock.json` lock/hash correctness,
- bounded prompt compilation,
- locked intent에서 파생된 inert downstream seed material.

kernel이 소유하지 않아야 하는 것은 다음이다:

- task runners,
- SPEC runners,
- Codex execution,
- shell adapters,
- queues,
- agent teams,
- PR automation,
- release automation,
- downstream runtime state.

task runners, SPEC runners, Codex exec, queues, agent teams, PR/release
automation을 core에 추가하는 contribution은 `ni-kernel` scope 밖이다.

## Preferred Contributions

좋은 kernel contribution은 보통 다음 영역을 개선한다:

- readiness rules,
- docs/contract sync validation,
- examples,
- target seed formats,
- benchmark fixtures,
- documentation clarity,
- lock/hash correctness,
- conversation authoring UX.

Target seed material은 derived, inert, locked-plan dependent 상태를 유지할 때
환영한다. 어떤 change가 downstream work를 시작하거나, schedule하거나,
track하거나, complete한다면 그것은 `ni-kernel` 밖에 속한다.

## Issue 또는 PR을 열기 전에

proposal이 `ni-kernel`에 속하는지, 아니면 downstream seed material에 속하는지
먼저 확인해달라. 유용한 질문은 다음이다:

```text
이 change는 locked intent를 validate/compile하는가, 아니면 work를 execute하는가?
```

work를 execute한다면 현재 core boundary 밖이다.

bug report에는 다음을 포함해달라:

- `ni` version 또는 commit,
- command와 flags,
- expected result,
- actual result,
- workspace shape, 공유해도 안전한 범위의 `docs/plan/**`,
  `.ni/contract.json`, `.ni/plan.lock.json` 상태.

public issues에 secrets, credentials, proprietary planning contracts, 민감한
prompts를 공유하지 말아달라.

## Pull Requests

changes는 작게 유지하고 하나의 coherent intent에 묶어달라. code change가
있다면 관련 checks의 validation evidence를 포함해달라.

repository validation:

```bash
bash scripts/quality.sh
```

Go files를 touched했다면 다음도 실행한다:

```bash
go test ./...
```

validation을 통과시키기 위해 acceptance criteria, risks, mitigations,
evaluations, non-goals를 약화하지 말아달라.

## Contributor License Agreement 없음

현재 contribution flow에는 contributor license agreement를 추가하지 않는다.
