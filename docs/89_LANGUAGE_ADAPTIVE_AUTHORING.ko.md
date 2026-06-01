# 언어 적응형 authoring

`ni-start`는 사용자의 최근 실질 메시지 언어에 맞춰 사용자-facing planning
question을 물어야 한다.

이 규칙은 UX 규칙일 뿐이다. CLI 권한 모델, readiness rule, schema key, command
name, ID, path, lock hash, status value를 바꾸지 않는다.

## 정책

- 사용자가 한국어로 말하면 planning question은 한국어로 묻는다.
- 사용자가 영어로 말하면 planning question은 영어로 묻는다.
- 사용자가 언어를 명시적으로 요청하면 그 언어를 사용한다.
- 대화가 섞여 있으면 사용자의 최신 명시적 언어 선호를 우선한다.
- Project documentation language가 선언되지 않았다면 human-facing prose는
  사용자의 현재 언어를 사용한다.
- Planning docs는 project가 선언한 documentation language를 따를 수 있다.
- ni repository 자체는 English canonical doc과 Korean companion doc 정책을
  유지한다.
- CLI output은 영어로 남아도 된다.
- 모델은 CLI proof와 next-question output을 사용자의 언어로 요약하거나 번역할 수
  있지만, 의미를 바꾸면 안 되고 authoritative ID, path, command, schema key,
  status constant는 그대로 보존해야 한다.

## 그대로 보존할 것

Machine-readable 또는 command-facing token은 번역하거나 다시 쓰지 않는다.

```text
R014
OQ-001
SYNC-014
ni status
ni end
ni run
.ni/contract.json
READY
BLOCKED
READY_WITH_DEFERRALS
```

모든 contract ID, file path, schema key, target name, status constant,
command name에도 같은 규칙을 적용한다.

## 예시

한국어 사용자가 `ni init` 이후 질문한 경우:

```text
ni는 초기 프로젝트 의도가 아직 lock할 만큼 명확하지 않아서 BLOCKED 상태입니다.
구현은 아직 시작되지 않았고, 지금은 planning 단계입니다.

1. 이 프로젝트는 누구를 위해 어떤 현실을 바꿔야 하나요?
2. 주요 actor는 누구이고, 각 actor가 얻어야 하는 outcome은 무엇인가요?
3. 예상 delivery surface는 CLI, web app, conversation, document, workflow,
   research protocol, human service 중 무엇인가요?
```

영어 사용자가 `ni init` 이후 질문한 경우:

```text
ni is blocked because the initial project intent is not explicit enough to lock
yet. Implementation has not started. This is still planning.

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?
```

영어 CLI output을 한국어로 설명하는 경우:

```text
CLI는 `R014 Project purpose is missing`을 보고했습니다. 즉,
`docs/plan/00_project_brief.md`와 `.ni/contract.json`에 lock 가능한 목적이
아직 부족합니다. `R014`, `docs/plan/00_project_brief.md`,
`.ni/contract.json`은 그대로 유지합니다.
```

## 경계

- Localization으로 readiness gate를 약화하지 않는다.
- Model-translated text를 CLI output보다 authoritative하게 만들지 않는다.
- Command name, file path, schema key, ID, status constant, target name을
  번역하지 않는다.
- 친절한 번역을 lock, readiness proof, prompt compiler output으로 취급하지
  않는다.
