# 터미널 없이 계획하기

CLI를 설치하기 전에도 `ni`의 planning method에서 이득을 얻을 수 있다. 핵심은
discipline이다. Intent를 명시하고, docs와 contract draft를 함께 유지하고,
blocker를 드러내며, plan이 검증되기 전에는 agent에게 work를 넘기지 않는다.

No-terminal planning은 validated `ni`와 같지 않다. Deterministic readiness,
locking, lock hash verification, prompt compilation에는 CLI가 필요하다.

## 세 가지 수준

| 수준 | 작동 방식 | 신뢰할 수 있는 것 |
| --- | --- | --- |
| Full `ni` | CLI가 설치되어 있고 `ni status`, `ni end`, `ni run`을 사용할 수 있다. | Authoritative readiness, lock creation, lock hash checks, bounded prompt compilation. |
| Model pack assisted | Claude/Codex-style skills가 planning docs authoring과 contract drafting을 안내한다. Lock 전에는 user, teammate, CI runner가 CLI validation을 실행해야 한다. | Model-assisted drafting에는 유용하지만 readiness와 lock claim은 CLI output 이후에만 authoritative하다. |
| Read-only method | Intent Lock checklist나 이 문서를 model session에 복사하고 plan을 검토하게 한다. | Learning과 self-review에는 유용하지만 authoritative하지 않고 validated `ni`와 equivalent하지 않다. |

## 수동 흐름

1. 이 repository에서 필요한 skill instructions를 download하거나 copy한다:
   `packages/claude-skills`, `packages/codex-skills`, `.agents/skills`.
2. Model에게 project의 `docs/plan` draft를 만들게 한다. Purpose, actors,
   capabilities, requirements, decisions, risks, evaluations, non-goals,
   constraints, artifacts, open questions를 다뤄야 한다.
3. Model에게 docs와 함께 `.ni/contract.json` draft를 유지하게 한다.
   Tentative하거나 conflicting한 statement는 assumptions, draft records, open
   questions로 남겨야 한다.
4. Blocker checklist를 visible하게 유지한다. Open blocker questions는 plan이
   거의 완성되어 보여도 lock을 막아야 한다.
5. 나중에 CLI를 사용할 수 있는 teammate, CI job, local setup에게 `ni status`를
   실행하게 한다. Result가 blocked라면 planning conversation을 계속한다.
6. `ni status`가 plan ready를 report하면 `ni end`로 lock을 만든다.
   `.ni/plan.lock.json`은 CLI만 만들어야 한다.
7. Final handoff prompt는 `ni run`으로 compile한다. `ni run`은 text를 compile할
   뿐 shell commands, agents, queues, downstream work를 실행하지 않는다.

## Intent Lock checklist

터미널 없이 작업할 때 이 checklist를 사용한다:

- Project purpose가 explicit한가?
- Actors와 outcomes가 named 상태인가?
- Every capability가 at least one requirement와 evaluation에 trace되는가?
- Non-goals가 visible한가?
- High-severity risks에 mitigations가 있는가?
- Open questions가 clear하게 marked되어 있는가, especially blockers?
- Accepted decisions가 assumptions와 rejected options에서 분리되어 있는가?
- Expected artifacts가 named 상태인가?
- Downstream handoff가 runtime execution이 아니라 planning output으로 bounded되어 있는가?

이 checklist는 learning과 drafting aid다. Model이 더 좋은 질문을 하게 도울 수는
있지만 `ni status`, `ni end`, `ni run`을 대체하지 않는다.

## Boundary

No-terminal planning은 hosted web app, model API calls, runtime execution, shell
adapters, queues, automation behavior를 추가해서는 안 된다. 이것은 kernel boundary를
보존하면서 Intent Lock method를 시작하는 docs-first 방식이다:

```text
model pack or copied checklist -> draft docs and contract
ni CLI -> deterministic readiness, lock, hash proof, prompt compile
```
