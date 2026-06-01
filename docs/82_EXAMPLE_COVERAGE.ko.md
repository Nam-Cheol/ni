# Example Coverage

이 matrix는 각 public example이 무엇을 증명하는지, 어떻게 검증되는지, 그리고
non-execution boundary가 어디인지 기록한다. 예시들은 Project Intent Compiler
asset이다. Planning contract, grouped readiness proof, lock, prompt compilation
boundary를 검증하지만 예시 제품을 구현하거나 downstream agent를 실행하지 않는다.

## Verification command

Repository root에서 public demo check를 실행한다:

```bash
bash scripts/demo-check.sh
```

더 넓은 repository check:

```bash
bash scripts/quality.sh
```

## Coverage matrix

| Example | 증명하는 것 | Product type | Delivery surface | Expected status | demo-check coverage | Docs-only? | Korean companion |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `examples/ambiguous-prompt-blocked/` | vague execution이 handoff 전에 blocked 되며 grouped open-blocker question이 다음 turn을 안내한다. | `software` | `web` | `BLOCKED` | `ni status`와 grouped next-question rendering을 실행한다. | 아니오, blocked workspace fixture. | Yes |
| `examples/research-protocol/` | ni는 software-only가 아니며 research protocol은 fieldwork 전에 lock될 수 있다. | `research_protocol` | `document` | `READY` | status를 실행하고 existing lock에서 `human-team` prompt를 컴파일한다. | 아니오, locked workspace fixture. | Yes |
| `examples/conversation-product/` | conversation-surface planning은 chatbot runner가 되지 않고 lock될 수 있다. | `conversation_product` | `conversation` | `READY` | status, `human-team` prompt compile, seed-only export check를 실행한다. | 아니오, locked workspace fixture. | Yes |
| `examples/conversation-authoring/` | sustained model-user authoring은 docs, contract, session을 갱신하고 CLI proof는 stale sync를 잡아낸다. | `conversation_product` | `conversation`, `document` | `BLOCKED` | status와 `R012`를 확인하고 historical lock에서만 compile한다. | 아니오, historical lock material이 있는 blocked fixture. | Yes |
| `examples/namba-ai-upgrade/` | ni는 existing harness/workflow project의 upstream에서 계획하지만 그 harness가 되지 않는다. | `software` | `cli`, `document`, `workflow` | `BLOCKED` | status와 `R012`를 확인하고 historical lock에서만 Codex prompt를 컴파일한다. | 아니오, historical lock material이 있는 blocked fixture. | Yes |
| `examples/ni-start-dogfood/` | first-run card, grouped next questions, docs/contract/session update, re-status loop를 증명한다. | `conversation_product` | `conversation`, `document` | `READY_WITH_DEFERRALS` | status, grouped proof, existing lock의 `human-team` prompt compile을 실행한다. | 아니오, locked workspace fixture. | Yes |
| `examples/benchmark-report/` | `not_measured` boundary가 있는 benchmark/case-study reporting method와, historical blocked proof를 보존하고 resolved artifact-readiness transition을 status, lock, bounded prompt, before/after evidence, lessons로 package한 internal-dashboard case를 증명한다. | resolved dashboard benchmark artifact case는 `document_product` | isolated case workspace는 `document` | Dashboard case: benchmark artifact readiness에 한해 `READY` | required docs, historical blocked proof, resolved READY proof, blocker analysis, resolution path, answer packet, isolated lock summary, bounded prompt summary, before/after evidence, lessons, product/runtime claim의 남은 `not_measured` boundary를 확인한다. | 부분적: report template은 docs-only이고 dashboard case는 locked ni workspace를 가진다. | Yes |
| `examples/no-terminal-assisted/` | local CLI validation 전에도 docs와 contract draft는 만들 수 있지만 deterministic readiness claim은 하지 않는다. | draft `workflow` | draft `document` | claim하지 않음 | required file과 boundary wording만 확인한다. | 예, assisted draft. | Yes |

## Grouped next-question coverage

Grouped `ni status --proof --next-questions` UX는 다음 파일에서 직접 보여준다:

- `examples/ambiguous-prompt-blocked/05-next-questions.md`
- `examples/benchmark-report/cases/internal-dashboard/06-ni-status-proof.md`
- `examples/benchmark-report/cases/internal-dashboard/07-ni-next-questions.md`
- `examples/benchmark-report/cases/internal-dashboard/11-resolved-status-proof.md`
- `examples/benchmark-report/cases/internal-dashboard/12-resolved-next-questions.md`
- `examples/conversation-authoring/transcript.md`
- `examples/conversation-authoring/session-resume.md`
- `examples/ni-start-dogfood/03-model-summary-and-questions.md`
- `examples/ni-start-dogfood/06-status-proof.md`
- `examples/ni-start-dogfood/07-second-round-questions.md`
- `examples/ni-start-dogfood/README.md`

기대되는 model behavior는 group label을 보존하고, highest-priority group을 먼저
묻고, CLI answer shape를 사용하며, 사용자가 답한 뒤 `docs/plan/**`,
`.ni/contract.json`, `.ni/session.json`을 갱신하고
`ni status --dir . --proof --next-questions`를 다시 실행하거나 요청하는 것이다.

## Non-execution boundary

예시들은 다음을 하지 않는다:

- dashboard, assistant, research study, travel workflow, namba-ai change 구현;
- Codex, Claude APIs, model APIs, downstream agents, shell adapters 호출;
- queue, runtime execution, release automation, PR automation, evidence runner 생성;
- user-facing contract `add`, `list`, `set` authoring command 추가;
- benchmark result 조작 또는 statistical significance 주장.

Locked example은 CLI가 existing lock material을 먼저 검증하기 때문에 `ni run`으로
inert prompt를 컴파일할 수 있다. Blocked example은 historical lock 또는 generated
prompt material이 있어도 blocked 상태로 남는다.
