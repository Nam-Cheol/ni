# Research protocol benchmark answer packet

이 packet은 research-protocol benchmark blocker에 대한 human-fillable form이다.
Answer를 채우지 않고, blocker를 resolve하지 않으며, readiness를 강제로 만들지
않고, lock을 만들거나 prompt를 compile하지 않는다. Research work, data
collection, fieldwork도 허가하지 않는다.

답변은 다음 위치 안에서만 사용한다.

```text
examples/benchmark-report/cases/research-protocol/workspace/
```

repository-root planning lock이나 root `.ni/` state에는 적용하지 않는다.

## Status

- Current readiness: `BLOCKED`
- Lock created: no
- `ni-run` prompt compiled: no
- Required answers: `OQ-001` through `OQ-005`

## How to use this packet

1. 아래 답변란을 채운다.
2. 답변은 isolated research-protocol benchmark workspace 갱신에만 사용한다.
3. 해당 workspace를 대상으로 `ni status --proof --next-questions`를 실행한다.
4. readiness가 `READY` 또는 `READY_WITH_DEFERRALS`가 되고 user confirmation이
   있을 때만 lock한다.
5. valid benchmark workspace lock이 생긴 뒤에만 `ni run`을 실행한다.

## OQ-001 - Research question, supported decision, and final study artifact

Prompt:

이 study는 어떤 research question에 답해야 하며, 어떤 decision을 지원해야 하고,
planning에 충분한 final artifact는 무엇인가?

Required answer fields:

- Research question:
- Supported decision:
- Decision owner:
- Final study artifact:
- What should not be decided by this study:

Unsafe assumptions to avoid:

- study가 intervention location을 직접 선택한다고 가정하기.
- final artifact가 full statistical report라고 가정하기.
- research question이 specific planning decision이 아니라 모든 heat risk에 관한
  것이라고 가정하기.

## OQ-002 - Participant/observation scope, criteria, locations, and windows

Prompt:

누구 또는 무엇을, 어디에서, 어떤 time window 동안, 어떤 inclusion/exclusion
criteria에 따라 observe하는가?

Required answer fields:

- Observation units:
- Participant groups, if any:
- Inclusion criteria:
- Exclusion criteria:
- Locations or neighborhood types:
- Observation window:
- Sampling or selection rule:
- Out-of-scope populations or locations:

Unsafe assumptions to avoid:

- Criteria 없이 vulnerable participant를 모집하기.
- 모든 neighborhood를 interchangeable하게 다루기.
- 아무 hot day 또는 아무 location이 valid하다고 가정하기.

## OQ-003 - Consent, privacy, data handling, translation, accessibility, and sensitive-data boundaries

Prompt:

어떤 data를 collect 또는 use할 수 있고, consent는 어떻게 처리하며, 어떤 data가
sensitive하고, 어떤 accessibility 또는 translation support가 필요한가?

Required answer fields:

- Data types allowed:
- Data types prohibited:
- Consent approach:
- Privacy constraints:
- Storage/retention rule:
- Translation needs:
- Accessibility needs:
- Sensitive-data boundaries:
- Data that must remain out of scope:

Unsafe assumptions to avoid:

- identifiable data를 불필요하게 collect하기.
- public observation은 모든 경우 consent-free라고 보기.
- translation 또는 accessibility need를 무시하기.
- retention boundary 없이 sensitive note를 저장하기.

## OQ-004 - Field safety, heat/weather stop rules, vulnerable-group safeguards, and escalation path

Prompt:

Hot weather 또는 unsafe condition 동안 field team과 participant를 보호하는 safety
rule은 무엇인가?

Required answer fields:

- Field team safety rules:
- Heat/weather stop conditions:
- Maximum exposure or shift limits:
- Hydration/rest requirements:
- Vulnerable-group safeguards:
- Emergency escalation path:
- Conditions that cancel fieldwork:

Unsafe assumptions to avoid:

- dangerous heat에 field team을 내보내기.
- unsafe condition에도 observation을 계속하기.
- vulnerable group을 safeguard 없이 ordinary observation target으로 다루기.

## OQ-005 - Reviewer, acceptance evidence, and pre-fieldwork readiness criteria

Prompt:

누가 이 protocol을 review하며, fieldwork가 시작되기 전에 어떤 evidence면 충분한가?

Required answer fields:

- Reviewer or approval owner:
- Review audience:
- Minimum protocol artifact:
- Acceptance evidence:
- Pass/fail criteria:
- Pre-fieldwork readiness checklist:
- What is explicitly not required:
- Who can approve moving forward:

Unsafe assumptions to avoid:

- draft protocol을 fieldwork-ready로 보기.
- informal agreement를 approval로 보기.
- acceptance evidence가 생기기 전에 data collection을 시작하기.

## After answers are provided

Expected next steps:

1. isolated benchmark workspace만 update한다.
2. `docs/plan`을 update한다.
3. `.ni/contract.json`을 update한다.
4. 있으면 `.ni/session.json`을 update한다.
5. `ni status --proof --next-questions`를 실행한다.
6. `BLOCKED`이면 remaining blocker를 문서화한다.
7. `READY` 또는 `READY_WITH_DEFERRALS`이면 benchmark workspace 안에서만
   lock한다.
8. lock되면 benchmark workspace에서 bounded prompt를 compile한다.
9. measurement table을 정직하게 update한다.

Rules:

- root `.ni/plan.lock.json`을 편집하지 않는다.
- repository-root `ni end` 또는 `ni relock`을 실행하지 않는다.
- research-protocol benchmark workspace가 `BLOCKED`로 남아 있는 동안 그
  workspace에서 `ni end` 또는 `ni run`을 실행하지 않는다.
- downstream agent를 실행하지 않는다.
- fieldwork를 실행하지 않는다.
- participant를 recruit하지 않는다.
- data를 collect하지 않는다.
- model API를 호출하지 않는다.
- prompt 또는 lock evidence를 조작하지 않는다.
