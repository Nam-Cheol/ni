# Language-adaptive authoring

`ni-start` should ask user-facing planning questions in the language of the
user's latest substantive message.

This is a UX rule only. It does not change the CLI authority model, readiness
rules, schema keys, command names, IDs, paths, lock hashes, or status values.

## Policy

- If the user speaks Korean, ask planning questions in Korean.
- If the user speaks English, ask planning questions in English.
- If the user explicitly requests a language, use that language.
- If the conversation is mixed, prefer the user's latest explicit language
  preference.
- If no project documentation language is declared, use the user's current
  language for human-facing prose.
- Planning docs may follow the project's declared documentation language.
- The ni repository itself keeps English canonical docs and Korean companion
  docs where companions exist.
- CLI output may remain English.
- A model may summarize or translate CLI proof and next-question output into
  the user's language, but it must preserve meaning and keep authoritative
  IDs, paths, commands, schema keys, and status constants unchanged.

## Preserve exactly

Do not translate or rewrite machine-readable or command-facing tokens:

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

The same rule applies to all contract IDs, file paths, schema keys, target
names, status constants, and command names.

## Examples

Korean user after `ni init`:

```text
ni는 초기 프로젝트 의도가 아직 lock할 만큼 명확하지 않아서 BLOCKED 상태입니다.
구현은 아직 시작되지 않았고, 지금은 planning 단계입니다.

1. 이 프로젝트는 누구를 위해 어떤 현실을 바꿔야 하나요?
2. 주요 actor는 누구이고, 각 actor가 얻어야 하는 outcome은 무엇인가요?
3. 예상 delivery surface는 CLI, web app, conversation, document, workflow,
   research protocol, human service 중 무엇인가요?
```

English user after `ni init`:

```text
ni is blocked because the initial project intent is not explicit enough to lock
yet. Implementation has not started. This is still planning.

1. What should this project change, for whom, and why does it matter?
2. Who are the primary actors, and what outcome should each one get?
3. What is the likely delivery surface: CLI, web app, conversation, document,
   workflow, research protocol, human service, or something else?
```

English CLI output explained in Korean:

```text
CLI는 `R014 Project purpose is missing`을 보고했습니다. 즉,
`docs/plan/00_project_brief.md`와 `.ni/contract.json`에 lock 가능한 목적이
아직 부족합니다. `R014`, `docs/plan/00_project_brief.md`,
`.ni/contract.json`은 그대로 유지합니다.
```

## Boundaries

- Do not use localization to weaken readiness gates.
- Do not make model-translated text authoritative over CLI output.
- Do not translate command names, file paths, schema keys, IDs, status
  constants, or target names.
- Do not treat a friendly translation as a lock, readiness proof, or prompt
  compiler output.
