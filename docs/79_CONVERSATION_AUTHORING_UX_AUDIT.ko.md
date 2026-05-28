# Conversation Authoring UX Audit

Task 139는 v0.4 conversation-driven authoring experience를 audit한다. 이
문서는 audit note일 뿐이다. CLI behavior, lock state, model packs, downstream
execution behavior를 변경하지 않는다.

## Scope Reviewed

- `README.md` and `README.ko.md`
- `AGENTS.md`
- `docs/28_CONVERSATION_AUTHORING.md`
- `docs/29_AUTHORING_PROTOCOL.md`
- `docs/30_DOC_UPDATE_RULES.md`
- `docs/31_NI_START_BEHAVIOR.md`
- `docs/34_READINESS_INTERVIEW.md`
- `docs/35_NI_END_CONFIRMATION.md`
- `docs/36_NI_RUN_HANDOFF.md`
- `docs/37_MODEL_EDIT_SAFETY.md`
- `docs/no-terminal.md`
- `docs/55_MODEL_WORKSPACE_PACKS.md`
- `packages/claude-skills/**`
- `packages/codex-skills/**`
- `examples/ni-start-dogfood/**`
- `examples/conversation-authoring/**`

## Overall Finding

Kernel boundary는 강하다. Docs와 skills는 conversation이 authoring surface이고
`ni status`, `ni end`, `ni run`이 authority라는 규칙을 일관되게 보존한다.

v0.4 adoption의 주요 risk는 concept drift가 아니다. 첫 실행 orientation이다.
새 사용자는 여러 파일을 읽으면 protocol을 이해할 수 있지만, `ni init` 직후
"무엇을 말해야 하는가?" 경로는 README, protocol docs, skills, examples에 나뉘어
있다.

## Top 10 UX Gaps

| Rank | Gap | User impact | Type |
| --- | --- | --- | --- |
| 1 | `ni init` 이후 새 사용자가 바로 말할 수 있는 starter script가 한 곳에 없다. README는 conversation을 쓰라고 말하고 examples는 보여주지만, compact copyable path는 전면에 드러나지 않는다. | High: 사용자가 manual JSON editing으로 돌아가거나 model에게 너무 일찍 implementation을 요청할 수 있다. | 지금은 docs-only; 나중에 optional CLI hint 가능 |
| 2 | "continue planning"과 "`ni-end` 실행"의 결정 지점은 protocol docs에는 명확하지만 first-run materials에는 충분히 보이지 않는다. 특히 `READY_WITH_DEFERRALS`는 사용자가 보는 순간 plain-language 설명이 필요하다. | High: 사용자가 "거의 됨"을 lock readiness로 오해하거나 deferrals를 숨은 cleanup으로 볼 수 있다. | Docs-only |
| 3 | Readiness failures에는 `next_questions`가 있지만, user-facing loop가 reusable interview card로 포장되어 있지 않다. Examples는 좋은 질문을 보여주지만 README와 model-pack READMEs는 blocker questions에 답하는 작은 pattern을 제공하지 않는다. | High: blocked status가 다음 planning turn이 아니라 dead end처럼 느껴질 수 있다. | Docs-only |
| 4 | Docs/contract sync expectations는 명시적이지만 drift diagnostics가 아직 first-class user experience로 보이지 않는다. 사용자는 docs와 `.ni/contract.json`이 함께 움직여야 함을 배울 수 있지만, reviewed docs 안에서는 "이 records가 불일치한다"는 proof surface가 명확하지 않다. | High: sync drift는 model-authored planning state에서 가장 가능성 높은 failure mode다. | Best result에는 CLI changes 필요 |
| 5 | `ni-start` resume behavior는 model에게는 엄격하지만, user에게 concise resume proof template로 보이지 않는다. Docs는 session claims를 verify해야 한다고 말하고 examples는 resume을 보여주지만, 예측 가능한 "X에서 resume했고 Y를 verify했으며 conflict는 Z" shape가 필요하다. | Medium-high: 사용자가 resumed planning을 신뢰하지 못하거나 hidden chat memory를 과신할 수 있다. | Docs-only |
| 6 | Assumptions, decisions, risks, non-goals, open questions는 protocol docs에서 구분되지만 users와 pack users를 위한 짧은 authoring cheat sheet가 없다. 구분은 맞지만 깊은 reading이 필요하다. | Medium-high: pack README만 읽는 model host에서는 ambiguous statement가 너무 빨리 accepted로 승격될 수 있다. | Docs-only |
| 7 | Model workspace packs는 CLI authority를 명확히 설명하지만 setup flow가 task-oriented라기보다 installation-oriented다. 설치 후에도 사용자는 "project 열기, `ni-start` invoke, 필요하면 CLI proof paste"라는 compact working loop가 필요하다. | Medium: pack users가 설치는 제대로 해도 첫 working loop를 모를 수 있다. | Docs-only |
| 8 | No-terminal mode는 deterministic validation을 과장하지 않지만 proof-capture template가 없다. 사용자는 trusted runner에게 `ni status`, `ni end`, `ni run`을 요청하라는 안내를 받지만, 정확히 어떤 output을 다시 paste해야 하고 model이 그것을 어떻게 다뤄야 하는지는 부족하다. | Medium: no-terminal users가 authoritative CLI output 대신 summary를 paste할 수 있다. | Docs-only |
| 9 | Korean companion coverage는 public release와 distribution docs에는 있지만 central conversation-authoring protocol docs는 대부분 English-only다. `README.ko.md`는 flow를 가리키지만 Korean-first users는 entry page 이후 depth를 잃는다. | Medium: v0.4 focus area에서 Korean adoption path가 덜 완성되어 있다. | Docs-only |
| 10 | Codex와 Claude `ni-start` skills는 authority 면에서는 aligned지만 Codex skill이 Claude skill보다 resume과 output contract가 더 풍부하다. 이 차이는 model workspaces 간 behavior를 uneven하게 만들 수 있다. | Medium: cross-host examples가 깔끔하게 이전되지 않을 수 있다. | Docs-only |

## Audit Question Answers

| Question | Answer |
| --- | --- |
| Can a new user understand what to say after `ni init`? | 부분적으로 그렇다. README와 examples가 path를 암시하지만 compact first prompt 또는 post-init conversation card는 없다. |
| Does `ni-start` clearly resume previous planning state? | Model instructions 기준으로는 그렇다. 특히 Codex와 `docs/31_NI_START_BEHAVIOR.md`가 강하다. User-visible resume proof로는 덜 명확하다. |
| Does the model know how to update docs and contract together? | 그렇다. `docs/29_AUTHORING_PROTOCOL.md`, `docs/30_DOC_UPDATE_RULES.md`, `docs/37_MODEL_EDIT_SAFETY.md`, skills가 이 규칙을 반복한다. |
| Does the user know when to continue planning vs run `ni-end`? | 부분적으로 그렇다. Rule은 있지만 first-run docs에서 `BLOCKED`, `READY_WITH_DEFERRALS`, `READY` transitions를 더 명확히 보여야 한다. |
| Are readiness failures translated into useful next questions? | Protocol level에서는 `--next-questions`를 통해 그렇다. Adoption gap은 그 loop를 users에게 포장하는 것이다. |
| Are assumptions, decisions, risks, non-goals, and open questions handled distinctly? | Deep docs에서는 그렇다. 이 distinction에는 compact user and pack quick reference가 필요하다. |
| Do model workspace packs explain CLI authority clearly? | 그렇다. 두 packs 모두 skills는 UX이고 CLI가 authority라고 일관되게 말한다. 부족한 부분은 task-first usage loop다. |
| Does no-terminal mode avoid overstating deterministic validation? | 그렇다. deterministic readiness, locking, hash verification, prompt compilation에는 CLI가 필요하다고 명확히 말한다. Gap은 exact proof capture다. |

## Docs-Only Gaps

이 gaps는 CLI behavior changes 없이 해결할 수 있다:

- Gap 1은 README/docs/model-pack quickstart로 해결한다. 나중에 project가
  `ni init` next-step hint를 원할 때만 CLI hint를 추가한다.
- Gap 2는 readiness transition card로 해결한다.
- Gap 3은 reusable blocker interview card로 해결한다.
- Gap 5는 user-visible resume proof template로 해결한다.
- Gap 6은 record classification cheat sheet로 해결한다.
- Gap 7은 task-first pack usage instructions로 해결한다.
- Gap 8은 no-terminal proof-paste templates로 해결한다.
- Gap 9는 v0.4 authoring docs의 Korean companions를 추가해서 해결한다.
- Gap 10은 Claude와 Codex skill wording을 align해서 해결한다.

## CLI-Change Gaps

다음 gaps는 later task에서 CLI work가 있으면 효과가 크다:

- Gap 1: `ni init`이 workspace creation 이후 짧은 "next conversation prompt"를
  출력해야 한다면 CLI change가 필요하다.
- Gap 4: `ni status`가 explicit docs/contract sync diagnostics,
  record-level mismatch categories, 또는 model packs가 reinterpretation 없이
  보존할 JSON proof fields를 emit해야 한다면 CLI change가 필요하다.

이 audit task 안에서는 CLI work를 추천하지 않는다.

## Recommended Next 3 Tasks

1. Add a first-run conversation authoring card.

   README, model-pack READMEs, `docs/no-terminal.md`에 compact path를 추가한다:

   ```text
   After ni init, say:
   "Invoke ni-start. Help me plan <project>. Ask only the next questions needed
   for ni status, and keep docs/plan/** plus .ni/contract.json synchronized."
   ```

   Transition rule도 포함한다: `BLOCKED`는 continue planning,
   `READY_WITH_DEFERRALS`는 visible deferrals review 후 `ni-end`,
   `READY`는 `ni-end` confirmation을 시작할 수 있음을 뜻한다.

2. Add readiness proof and blocker-question examples to the adoption path.

   `ni status --proof --next-questions`를 위한 작은 reusable card를 만든다.
   Model이 CLI status를 보존하고, blockers를 설명하고, one to three next
   questions를 묻고, model judgment를 readiness처럼 다루지 않는 방법을 보여준다.

3. Design docs/contract sync diagnostics for a later CLI task.

   Docs/contract mismatch에 대한 expected `ni status` issue shape를 draft한다:
   affected record ID, doc path, contract field, severity, next question.
   첫 단계는 spec 또는 doc proposal로 유지하고, CLI behavior 구현은 separate task에서
   진행한다.

## Suggested v0.4 Success Test

Fresh user는 다음을 할 수 있어야 한다:

1. `ni init` 실행.
2. Model workspace에 conversation prompt 하나 copy.
3. Focused blocker questions에 답하기.
4. Model이 바꾼 files를 정확히 보기.
5. `ni status`가 왜 `BLOCKED`, `READY_WITH_DEFERRALS`, 또는 `READY`인지 이해하기.
6. Visible summary 이후에만 `ni-end` confirm.
7. `ni-run`이 bounded prompt만 compile한다는 점을 알고 실행.

이 path가 manual `.ni/contract.json` editing 없이, downstream execution claims 없이
작동하면 v0.4 conversation authoring hardening은 올바른 문제를 겨냥하고 있다.
