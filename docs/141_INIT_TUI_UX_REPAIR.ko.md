# Init TUI UX Repair

## Current status

State:
- v0.6.0 release: published and verified
- primary command: namba-intent
- init TUI: Bubble Tea v2 / Lip Gloss v2
- reported UX issues: unclear help text, arrow-key editing conflict, unclear next steps
- Homebrew: Planned / v0.5 candidate
- Windows real-host verification: pending
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

## Repair goal

이번 수리는 `namba-intent init .` dogfooding 중 나온 첫 사용자 문제 3가지를
다룬다.

- setup guide가 각 질문의 의미와 거친 답변 허용 범위를 충분히 설명하지
  않았다.
- answer field 안에서 arrow key가 wizard step 이동으로 소비되어 텍스트
  편집이 어려웠다.
- init 뒤 plain text summary가 생성물, 현재 상태, status/end/run 경로,
  optional model help의 위치를 충분히 설명하지 않았다.

Init은 여전히 planning setup only이다. Skill 설치, plan lock, agent 실행,
prompt 실행, adapter 추가, runtime execution state 생성은 하지 않는다.

## Question/help copy

| Question / area | Old problem | New guidance | Notes |
| --- | --- | --- | --- |
| Language choice | 언어 선택은 가능하지만 readiness 맥락이 짧았다. | 선택한 언어가 labels와 review guidance에 쓰인다고 안내한다. | Menu는 choice list라 arrow navigation을 유지한다. |
| Korean guide tone | 첫 수리 뒤에도 번역투와 내부 문서 말투가 남아 있었다. | 한국어 copy를 구어체로 바꾸고 내부 용어를 줄였다. | `humanize-korean` fast-path 기준을 적용했다. |
| Korean guide layout | 질문 안내, 예시, 모를 때의 답이 한 덩어리처럼 보여 읽기 어려웠다. | `쓸 내용`, `예시`, `모르면`, `다음`, `READY`를 색이 들어간 짧은 라벨로 나누고 항목별로 줄을 분리한다. | 80x24에서도 입력 박스가 보이도록 TUI 안의 문장은 짧게 유지한다. |
| Project name | 단순 이름 입력처럼 보였다. | 나중에 다시 봤을 때 알아볼 수 있는 짧은 이름을 묻는다. | Folder-name default는 계속 허용한다. |
| Project goal | 완벽한 최종 답변이 필요한 것처럼 느껴질 수 있었다. | 누가 쓰는지, 무엇이 좋아져야 하는지, 한두 문장으로 자세히 써도 된다고 안내한다. | 모르면 "아직 모름"으로 남길 수 있다. |
| Target users / audience | Actor 같은 기술 용어가 앞에 나왔다. | 누가 보고, 쓰고, 검토하는지 묻는다. | Rough user list도 init에서는 허용한다. |
| Downstream agent task | Init이 agent를 실행하는 것처럼 들릴 수 있었다. | 계획을 잠근 뒤 맡길 일이며 지금 실행되는 일이 아니라고 말한다. | Non-execution boundary 보존. |
| Constraints / non-goals | Hard boundary를 놓치기 쉬웠다. | 이번에 절대 하지 말아야 할 일을 묻는다. | Default는 lock 전 downstream work 금지를 유지한다. |
| Success criteria | 답변 형태가 불명확했다. | 나중에 무엇을 확인하면 "이 계획이면 됐다"고 볼 수 있는지 묻는다. | 모르는 기준은 status에서 다시 볼 수 있다. |
| Known blockers / open questions | 불확실성을 숨기고 지나갈 수 있었다. | 아직 몰라서 결정하면 안 되는 점을 묻는다. | 불확실성은 lock 전까지 보이게 둔다. |
| Deferrals | Deferred scope와 accepted work가 섞일 수 있었다. | 중요하지만 이번에는 하지 않을 일을 묻는다. | Deferred work가 plan 안에 섞이지 않게 한다. |

Answer 단계 guide는 이제 다음을 색 라벨과 항목별 줄로 나누어 보여준다.

- 쓸 내용
- 예를 들면
- 아직 잘 모르겠다면
- `status`가 빠진 점을 알려준다는 짧은 안내
- READY는 "계획을 잠글 준비"라는 뜻이지 제품이 완성됐다는 뜻이 아니다

## Keyboard behavior

| Context | Key | Expected behavior | Implemented behavior | Notes |
| --- | --- | --- | --- | --- |
| Answer input | Left / Right | Answer text 안에서 cursor 이동. | Answer cursor가 좌우로 움직인다. | Wizard step을 바꾸지 않는다. |
| Answer input | Up / Down | Multiline text 안에서 이동하거나 single-line에서는 예측 가능하게 동작. | Paste된 multiline row 사이 이동; single-line에서는 step change 없음. | accidental step jump 방지. |
| Answer input | Backspace / Delete | 일반 텍스트 편집. | Backspace는 cursor 앞, Delete는 cursor 위치 문자를 지운다. | Cursor 보존. |
| Answer input | Enter | 적절할 때 계속 진행. | 다음 field로 이동하거나 마지막 field에서 confirmation으로 간다. | Newline 삽입은 하지 않는다. |
| Answer input | Tab / Shift+Tab | Next / previous field. | Tab은 다음, Shift+Tab은 이전 field. | Editing 중 details는 Ctrl+D. |
| Answer input | Ctrl+Right / Ctrl+Left | Next / previous step. | 다음/이전 field로 이동한다. | Arrow editing과 분리. |
| Answer input | Esc | 뒤로 가거나 첫 field에서는 cancel. | 이전 field로 이동, 첫 field에서는 cancel. | 기존 cancel semantics 유지. |
| Answer input | q | 텍스트 입력. | Answer field에서는 plain `q`가 글자로 입력된다. | 일반 문자를 빼앗지 않기 위해 Ctrl+Q가 quit이다. |
| Menus | Up / Down / Left / Right | Choice navigation. | Language, existing-file, confirmation choice가 arrow를 사용한다. | Context-aware arrows. |
| Menus | q | Quit/cancel. | Language, existing-file, confirmation stage에서 `q`로 cancel 가능. | Answer field에서는 q를 텍스트로 유지. |
| Confirmation screen | Enter | 선택된 동작 confirm. | 확인 선택 시 initial artifacts를 쓴다. | Lock은 하지 않는다. |
| Cancel/quit | Ctrl+C / Ctrl+Q | 안전하게 종료. | Pending TUI answers를 쓰지 않고 cancel/quit한다. | Global control 유지. |

## Post-init summary

Init이 끝나면 AltScreen 밖에서 plain text summary를 출력한다. 구조는 다음
질문에 답한다.

- What was created?
- What was skipped or unchanged?
- What state am I in now?
- What do I do next?
- How can I use an AI assistant?
- What does Namba Intent not do?

Next-step wording은 다음을 포함한다.

- Run `namba-intent status --proof --next-questions`
- If status is BLOCKED, answer the listed questions or refine docs/plan/**
- When status is READY and you agree with the plan, run `namba-intent end`
- Then run `namba-intent run --max-chars 4000` to compile a bounded handoff prompt
- `namba-intent run` does not execute the prompt or run an agent

## Model/skills guidance

Init은 model skill을 만들거나 설치하지 않는다. Skill과 model assistance는
선택적 UX layer다.

Model help가 필요하면 summary는 assistant에게 `docs/plan/**`,
`.ni/contract.json`, `.ni/session.json`을 읽게 하고, next questions 답변과
docs plus contract 동시 업데이트를 돕게 하라고 안내한다.

Authority boundary는 그대로다.

```text
Skills are UX; CLI is authority.
```

Model은 draft와 explanation을 도울 수 있다. Readiness를 claim하거나 project를
lock하거나 `namba-intent status`, `namba-intent end`, `namba-intent run`을
대체할 수 없다.

## Tests added

- `TestAnswerFieldArrowsEditTextInsteadOfChangingSteps`
- `TestAnswerFieldStepNavigationUsesTabEnterAndCtrlArrows`
- `TestMenuQQuitsWithoutStealingAnswerText`
- next steps, READY boundary, run boundary, CLI authority guidance를 확인하는 CLI init summary assertions
- visible `setup guide` / `guide` copy와 new help bars를 반영한 responsive render snapshots

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| init TUI | Guided planning setup only. | TUI는 질문을 더 명확히 하고 confirm 뒤 draft init answers만 쓴다. | Yes | Lock 또는 execution behavior 없음. |
| model assistance | Optional draft/help layer. | Post-init summary가 assistant help 사용법을 설명한다. | Yes | CLI remains authority. |
| skills | Skills are UX only. | Init이 skill을 만들거나 설치하지 않는다고 명시한다. | Yes | Skill installation path 없음. |
| status | CLI readiness gate. | Summary가 `namba-intent status --proof --next-questions`를 안내한다. | Yes | Model은 readiness를 claim할 수 없다. |
| end | READY와 user agreement 뒤 lock. | Summary가 READY와 사용자 동의 뒤 `end`를 실행하라고 말한다. | Yes | Project root `end`는 실행하지 않았다. |
| run | Bounded handoff prompt compilation only. | Summary가 `run`은 prompt 또는 agent를 실행하지 않는다고 말한다. | Yes | Generated prompt 실행 없음. |
| READY | Planning readiness, not product readiness. | TUI와 summary가 이 boundary를 말한다. | Yes | Product-readiness overclaim 방지. |
| Homebrew | Planned / v0.5 candidate. | 보존됨. | Yes | Available claim 없음. |
| Windows real-host | Transcript 전까지 pending. | 보존됨. | Yes | Windows verification claim 없음. |
| runtime execution | No task runner, SPEC runner, harness, adapter, queue, PR/release automation, or downstream execution layer. | Summary와 implementation에서 보존됨. | Yes | Runtime behavior 추가 없음. |
