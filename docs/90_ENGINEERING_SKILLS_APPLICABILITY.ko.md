# Engineering skills 적용 가능성 감사

이 문서는
<https://github.com/mattpocock/skills/tree/main/skills/engineering>의 공개
engineering skill 중 일부가 ni에 어떤 개념적 영향을 줄 수 있는지 검토한다.

<!-- ni-boundary-allow: explicit negative boundary statement. -->
이 문서는 외부 skill을 복사하거나 vendor하지 않는다. endorsement 또는
compatibility를 주장하지 않는다. 새 skill을 만들지 않는다. Runtime execution,
<!-- ni-boundary-allow: explicit negative boundary statement. -->
task-runner behavior, issue tracker publishing, PR automation, shell adapter,
Codex exec, downstream agent execution을 추가하지 않는다.

## 경계

유용한 pattern은 다음 architecture를 보존할 때만 ni에 영향을 줄 수 있다.

```text
ni-kernel
  docs contract
  readiness gate
  lockfile
  prompt compiler
  source-of-truth rule

ni-downstream-seeds
  project-specific work graph
  project-specific evaluation plan
  project-specific evidence rules
  project-specific adapter notes
```

Kernel은 deterministic planning contract와 gate를 소유한다. Downstream seed
material은 lock 이후 생성될 수 있지만 derived and mutable이어야 한다.

## 적용 가능성 표

| Skill | ni에 유용한 pattern | 지금/나중/비적용 | 속하는 위치 | Boundary risk | 권장 action |
| --- | --- | --- | --- | --- | --- |
| `diagnose` | 가설보다 먼저 deterministic feedback loop를 만든다: reproduce, minimize, instrument, fix, regression-test. ni에서는 CLI status proof, fixture workspace, golden test, lock-hash repro에 대응된다. | 나중 validation/debug UX; 지금은 내부 개발 습관 | ni-kernel test와 docs; future skill pack guidance | Medium | `ni status`, docsync, lock mismatch, prompt budget debugging에 repro-loop language를 사용한다. Runtime instrumentation이나 production execution behavior는 추가하지 않는다. |
| `grill-with-docs` | 모호한 계획을 기존 domain language와 documented decision에 비춰 challenge하고, 한 번에 하나의 정확한 질문을 묻고, 결정이 굳어질 때 docs를 갱신한다. | 지금 | ni-start policy, readiness interview docs, model workspace packs | Low | ni planning grill / docs-contract challenge pattern으로 적용한다. Stable ID를 보존하고, docs와 contract를 비교하고, focused blocker question을 묻고, `docs/plan/**`과 `.ni/contract.json`을 함께 갱신한다. |
| `improve-codebase-architecture` | 일관된 domain vocabulary, locality, leverage, testability로 architecture friction을 찾는다. | 나중 | ni repo maintenance docs 또는 separate contributor skill pack | Medium | ni maintainer가 internal package를 refactor할 때 사용한다. End-user ni-kernel behavior로 만들지 않고, `ni status`, `ni end`, `ni run`에 visual report generation을 넣지 않는다. |
| `prototype` | Prototype을 하나의 질문에 답하는 throwaway artifact로 취급하고, 학습 뒤 삭제하거나 흡수한다. | 나중, 주의 | Separate downstream/sandboxed exploration seed, ni-kernel 아님 | High | Lock 이후 downstream-only 가능성으로만 문서화한다. Prototype command, runnable app, persistence, task-runner behavior를 ni에 추가하지 않는다. |
| `setup-matt-pocock-skills` | 다른 skill이 동작하기 전에 repo별 agent context를 설정한다: issue tracker, triage label, domain docs. | 나중 | Model-pack setup docs 또는 separate package | Medium | ni model-pack setup check의 inspiration으로만 사용한다: project root, authority docs, language policy, CLI path 확인. External scaffold를 복사하거나 GitHub/Claude/Codex global install을 가정하지 않는다. |
| `tdd` | Public interface를 통한 behavior test, 한 번에 하나의 vertical slice, red-green-refactor discipline. | 지금 ni 내부 개발 | ni-kernel test strategy와 contributor docs | Low | Readiness, docsync, lock, prompt budget, skill-pack check에 CLI-level 및 package-level behavior test를 선호한다. 상상한 internal shape를 고정하는 speculative test는 피한다. |
| `to-issues` | Plan을 dependency와 acceptance criteria가 있는 independently grabbable vertical slice로 나눈다. | 나중 | Downstream target seed 또는 separate package | Medium | Lock 이후 future issue-seed export에 영향을 줄 수 있다. Core에서 issue를 publish하거나 tracker를 mutate하거나 task runner가 되면 안 된다. |
| `to-prd` | Conversation과 codebase context를 problem, solution, stories, decisions, testing, out-of-scope가 있는 PRD로 종합한다. | 나중 | Downstream PRD/document seed, ni-kernel 아님 | Medium | Lock 이후 downstream export template으로 유용하다. PRD는 locked plan에서 파생되어야 하며, ni core가 issue tracker에 publish하면 안 된다. |
| `triage` | Incoming work를 state machine으로 분류하고 need-info, ready-for-agent, ready-for-human, won't-fix를 구분한다. | 나중 | Feedback/pressure classification docs 또는 separate model pack | Medium | Future pressure/feedback triage에 영향을 줄 수 있다. Issue comment, label, close, tracker mutation을 ni core에 추가하지 않는다. |
| `zoom-out` | Agent에게 한 단계 위로 올라가 project vocabulary로 relevant module/caller map을 설명하게 한다. | 지금 model skill-pack pattern | Model workspace packs와 contributor docs | Low | ni maintainer와 planning model의 orientation prompt로 사용한다. 설명용으로 유지하고 새 kernel command는 필요 없다. |

## 예상되는 ni 적용

즉시 문서와 skill pack에 영향을 줄 수 있는 것:

- `grill-with-docs`는 ni planning grill pattern이 된다. Vague intent를 docs와
  contract record에 비춰 challenge하고, terminology를 보존하고, 한 번에 하나의
  focused readiness question을 묻고, planning state를 visible하게 갱신한다.
- `tdd`는 public CLI와 package interface를 통한 behavior-first test를 강화한다.
- `zoom-out`은 contributor 또는 model이 edit 전에 higher-level map이 필요할 때
  model-pack orientation을 돕는다.

나중, core 밖에서:

- `diagnose`는 `ni status`, docsync, lock mismatch, prompt budget failure에 대한
  deterministic repro loop guidance를 형성할 수 있다.
- `to-prd`와 `to-issues`는 post-lock downstream seed format에 영향을 줄 수 있지만
  derived artifact로만 남아야 한다. ni core가 issue를 publish하면 안 된다.
- `triage`는 future feedback/pressure classification에 영향을 줄 수 있다.
- `setup-matt-pocock-skills`는 external scaffold를 복사하지 않는 범위에서 ni
  model-pack setup check에 inspiration을 줄 수 있다.

ni-kernel에 적용하지 말아야 하는 것:

- Runnable exploration으로서의 `prototype`, issue publishing, tracker mutation,
  shell execution, Codex execution, PR automation, queueing, runtime agent
  orchestration.

## Non-goals

- 외부 skill file을 ni에 복사하지 않는다.
- 외부 skill package를 vendor하지 않는다.
- Endorsement 또는 compatibility claim을 하지 않는다.
- 이 audit으로 새 ni skill을 만들지 않는다.
- Runtime execution, task runner, issue publishing, PR automation, shell
  adapter, Codex adapter, model API call, Homebrew availability claim을
  추가하지 않는다.
