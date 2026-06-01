# NI Grill

`ni-grill`은 `ni-end` 전에 draft NI contract를 압박 검토하는 model-facing
planning challenge skill이다.

이 문서는 [`90_ENGINEERING_SKILLS_APPLICABILITY.md`](90_ENGINEERING_SKILLS_APPLICABILITY.md)에서
확인한 `grill-with-docs` pattern을 NI에 맞게 적용한다. 외부 skill file을
복사하지 않는다. NI식 적용은 `docs/plan/**`, `.ni/contract.json`,
`ni status --proof --next-questions`를 기준으로 fuzzy project intent를
challenge하는 것이다.

## Boundary

`ni-grill`은 model workspace UX layer에 속한다.

Task runner, product implementation tool, readiness engine, lock writer,
prompt compiler, downstream adapter, execution harness가 아니다.

모든 NI skill과 같은 규칙을 따른다:

```text
Skills are UX; CLI is authority.
```

`ni-grill`은 weak assumptions, vague decisions, missing acceptance evidence,
docs/contract drift, risky handoff ambiguity, unsupported claims를 찾을 수
있다. 하지만 plan readiness를 선언하거나, lock을 approve하거나,
`.ni/plan.lock.json`을 edit하거나, CLI의 `BLOCKED` result를 override할 수
없다.

## 언제 쓰는가

`ni-grill`은 challenge할 planning content가 충분히 생긴 뒤, 사용자가
`ni-end`로 lock confirmation을 요청하기 전에 사용한다.

좋은 시점:

- `ni status`가 `BLOCKED`이고 grouped next questions를 더 날카롭게 물어야
  할 때;
- `ni status`가 `READY`에 가까워 보이지만 accepted records가 아직 vague할
  때;
- capabilities는 있지만 acceptance evidence가 약할 때;
- risks, non-goals, target handoff가 덜 구체적일 때;
- benchmark, proof, readiness claims에 `measured` 또는 `not_measured` label이
  필요할 때.

First-run brainstorming substitute로 쓰지 않는다. First-run blockers 또는 sync
diagnostics가 있으면 deterministic CLI questions를 먼저 묻는다.

## Required Process

Grill turn 시작 시 model은 다음을 읽는다:

- `AGENTS.md`;
- `docs/plan/**`;
- `.ni/contract.json`;
- `.ni/session.json` if present;
- `.ni/plan.lock.json` if present.

그 다음 run 또는 request한다:

```bash
ni status --dir . --proof --next-questions
```

Status output에 deterministic blockers가 있으면 `ni-grill`은 새로운 critique를
만들기 전에 그 blocker를 먼저 사용한다. Grouped next-question labels
`First-run card`, `Sync repairs`, `Risk decisions`, `Evaluation evidence`,
`Scope boundaries`, `Open blockers`를 preserve한다.

Deterministic blockers를 반영한 뒤에만 accepted 또는 nearly accepted planning
content에 extra pressure questions를 추가한다.

## Grill Categories

`ni-grill`은 다음을 pressure-test한다:

- purpose: specific reality change, single-problem focus, observable success;
- actors/outcomes: actor specificity, actor별 expected outcome, operators,
  reviewers, end users 분리;
- delivery surface: explicit surface와 docs/contract consistency;
- capabilities and requirements: accepted records의 trace links와 proof;
- evaluations: evidence가 test, review checklist, demo condition, user
  approval, protocol check, manual inspection 중 무엇인지;
- risks: high risks mitigation과 privacy/security/safety handling;
- non-goals: likely scope-drift traps의 explicit exclusion;
- decisions: accepted, deferred, rejected, not_applicable status의 의도적 사용;
- assumptions: uncertain statements가 assumptions 또는 open questions로 남아
  있는지;
- handoff: downstream actors가 할 일과 하지 말아야 할 일을 아는지;
- docs/contract sync: lock-critical records가 source 사이에서 일치하는지;
- claims: benchmark 또는 proof claims가 supported이고 measured 또는
  not_measured로 label되어 있는지.

## Finding Shape

각 grill finding은 concrete하고 answerable해야 한다:

```text
GRILL-001
- Affected planning ID or path: CAP-001 / docs/plan/02_capabilities.md
- Concern: Capability가 "usable report"라고 하지만 누가 accept하는지 정의하지 않는다.
- Why it matters: downstream work가 wrong reviewer에 맞춰질 수 있다.
- Question for the user: CAP-001은 누가 approve해야 하며 어떤 evidence가 completion proof인가?
- Expected answer shape: reviewer role plus test, review checklist, demo condition, user approval, protocol check, or manual inspection
- Blocks ni-end: yes
```

한 turn에서 grill questions는 최대 5개만 묻는다. Issue가 더 많으면 lock
blockers, high-risk ambiguity, acceptance evidence gaps,
privacy/security/safety risks, scope drift, target handoff ambiguity를 우선한다.

## Language Behavior

User-facing grill questions는 사용자의 latest substantive message 언어로 묻는다.
단, `R014`, `OQ-001`, `SYNC-014`, `GRILL-001`, `ni status`,
`.ni/contract.json`, `READY`, `BLOCKED` 같은 IDs, commands, paths, status
constants, target names, schema keys는 그대로 preserve한다.

CLI output은 English여도 된다. Model은 의미를 바꾸지 않는 선에서 사용자의
언어로 설명할 수 있다.

## Answer Handling

사용자가 grill questions에 답하면 model은 `docs/plan/**`,
`.ni/contract.json`, `.ni/session.json`을 함께 update한다.

Uncertainty는 assumptions 또는 open questions로 visible하게 남긴다. Clear
exclusions는 non-goals로 기록한다. Model은 uncertain answers를 accepted
decisions로 바꾸거나, risks/evaluations를 약화하거나, readiness를 통과하기 위해
edit해서는 안 된다.

Update 뒤에는 run 또는 request한다:

```bash
ni status --dir . --proof --next-questions
```

그리고 planning proof를 report한다:

- user input captured;
- interpreted planning records;
- updated artifacts;
- status before and after;
- remaining blockers;
- next question group.

## 다른 Skills와의 관계

`ni-start`는 main conversation 동안 planning state를 author and maintain한다.
`ni-grill`은 lock 전 draft plan을 challenge한다. `ni-end`는 CLI-ready planning
state를 summarize하고 explicit confirmation 뒤에만 `ni end`를 실행한다.

`ni-grill`은 답변이 planning edits를 요구하면 사용자를 `ni-start` flow로 돌려야
한다. `ni status`가 `READY` 또는 `READY_WITH_DEFERRALS`이고 사용자가 lock을
원할 때만 `ni-end`로 넘긴다.
