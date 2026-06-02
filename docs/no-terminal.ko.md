# 터미널 없이 계획하기

No-terminal planning은 사용자가 `ni` CLI를 설치하거나 실행하기 전에 Intent Lock
Protocol을 시작할 수 있게 하는 assisted drafting path다. Model conversation으로
초기 `docs/plan/**` notes와 draft `.ni/contract.json`을 만들 수 있지만, 이것만으로
trusted lock을 끝낼 수는 없다.

No-terminal에서 사용자-facing question은 사용자의 최근 실질 메시지 언어를 따라야
한다. 사용자가 한국어로 말하면 한국어 planning question을 묻고, 영어로 말하면
영어로 묻고, 언어를 명시적으로 요청하면 그 언어를 사용한다. ID, command, file
path, schema key, target name, status constant는 정확히 보존한다. 모델은 영어 CLI
proof를 사용자의 언어로 설명할 수 있지만, proof는 CLI output으로만 authoritative하다.

규칙은 단순하다. No-terminal 사용자는 planning을 시작하고 protocol을 배우고 review
대상 draft를 준비할 수 있다. Trusted readiness, lock creation, hash verification,
prompt compilation은 CLI-validated `ni`가 필요하다.

## No-terminal mode가 하는 일

No-terminal mode는:

- model conversation으로 planning을 시작하는 방법이다;
- `docs/plan/**`과 contract draft 작성에 유용하다;
- 설치 전에 `ni`의 Intent Lock Protocol을 배우는 데 유용하다;
- teammate 또는 local setup이 CLI를 실행하기 전에 questions, assumptions,
  non-goals, handoff notes를 준비하는 데 유용하다;
- assisted authoring path이지 source of authority가 아니다.

## No-terminal mode가 아닌 것

No-terminal mode는 다음이 아니다:

- deterministic readiness validation;
- 실제 `ni` lock;
- lock hash verification;
- `ni status`, `ni end`, `ni run`의 replacement;
- execution;
- hosted service;
- model authority.

Trusted CLI run이 결과를 만들고 exact proof를 확인할 수 있기 전에는 no-terminal
draft를 `READY`, `READY_WITH_DEFERRALS`, locked 상태로 설명하지 않는다.

## 터미널 없이 할 수 있는 일

터미널 접근 없이도 사용자는:

- model workspace에 `ni-start` guidance를 복사하거나 load할 수 있다;
- project idea, actors, delivery surface, constraints, risks를 설명할 수 있다;
- model에게 planning docs와 `.ni/contract.json`을 draft하게 할 수 있다;
- 의미 있는 update 뒤 model에게 draft planning proof block을 보여 달라고 할 수 있다;
- assumptions, draft decisions, explicit non-goals, open blocker questions를
  visible하게 유지할 수 있다;
- draft를 CLI validation이 가능한 teammate, CI job, later local setup에 넘길 수
  있다.

CLI proof 없이 신뢰할 수 없는 것은:

- readiness status;
- lock creation;
- lock hash checks;
- prompt compilation;
- locked plan에서 나온 downstream seed generation.

No-terminal mode에서 planning proof capture는 draft audit trail일 뿐이다. Model이
무엇을 바꾸었다고 해석했는지는 보여줄 수 있지만 deterministic validation은 아니며,
drafted docs와 contract를 CLI run이 validate한 뒤에만 trusted 상태가 된다.

Existing lock이 stale일 수 있을 때 no-terminal assistance는 changed intent review를
draft하거나 pasted `LOCK-STALE` output을 explain할 수 있다. 그래도 trusted
runner의 exact CLI output 없이는 readiness를 determine하거나, lock freshness를
prove하거나, hashes를 verify하거나, relock하거나, authoritative bounded handoff를
compile할 수 없다. Examples는
[`106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md`](106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md)를
참고한다.
Unusable, partial, pasted, fixture-only, target-workspace transcript를 구분하는
checklist는
[`107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.ko.md`](107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.ko.md)를
참고한다.

## No-terminal checklist

Local CLI 없이 시작할 때 이 checklist를 사용한다:

- `.agents/skills/ni-start/SKILL.md`,
  `packages/codex-skills/ni-start/SKILL.md`, 또는
  `packages/claude-skills/ni-start/SKILL.md`의 `ni-start` guidance를 복사하거나
  load한다.
- Conversation에서 project idea를 설명한다: 무엇이, 누구를 위해, 어떤 delivery로
  바뀌어야 하는가.
- Model에게 `docs/plan/**` draft를 만들게 한다.
- Model에게 docs와 함께 `.ni/contract.json`을 draft하게 한다.
- Uncertain statements는 assumptions 또는 open questions로 표시한다.
- Explicit exclusions는 non-goals로 표시한다.
- 의미 있는 update 뒤 concise planning proof block을 요청한다.
- Draft를 locked 상태로 취급하지 않는다.
- 나중에 CLI 또는 CLI를 실행할 수 있는 teammate로 validate한다.

## Intent Lock drafting checklist

Drafting 중에는 이 checklist를 사용한다:

- Project purpose가 explicit한가?
- Actors와 outcomes가 named 상태인가?
- Draft capability가 관련 record가 존재할 때 at least one requirement와
  evaluation에 trace되는가?
- Non-goals가 visible한가?
- High-severity risks에 mitigations가 있는가?
- Open questions가 clear하게 marked되어 있는가, especially blockers?
- Accepted decisions가 assumptions와 rejected options에서 분리되어 있는가?
- Expected artifacts가 named 상태인가?
- Downstream handoff가 runtime execution이 아니라 planning output으로 bounded되어
  있는가?

이 checklist는 model이 더 좋은 질문을 하게 돕는다. `ni status`, `ni end`,
`ni run`을 대체하지 않는다.

## Full ni로 넘어가는 경로

Plan이 implementation, budget, review, downstream seed generation을 안내할 수 있는
순간 assisted drafting에서 full `ni`로 넘어간다.

1. Verified release binary path 또는 curl installer로 `ni`를 설치한다.
2. Drafted workspace에서 `ni status --proof --next-questions`를 실행한다.
3. Planning conversation을 계속하면서 `docs/plan/**`과 `.ni/contract.json`을 함께
   update해 blockers를 해결한다.
4. `ni status`가 readiness를 report하고 user가 accepted plan을 confirm한 뒤에만
   `ni end`를 실행한다.
5. `.ni/plan.lock.json`이 존재하고 lock hashes가 valid한 뒤에만 `ni run`을
   실행한다.
6. Model workspace skills는 UX로 유지한다. Conversation을 guide하지만 CLI
   readiness, locking, hash verification, prompt compilation을 override하지 않는다.

## Team handoff path

No-terminal workflow는 team에서도 유용하다:

1. Terminal access가 없는 user가 model과 plan을 draft한다.
2. CLI가 있는 teammate가 `ni status --proof --next-questions`를 실행한다.
3. Teammate가 blockers, proof, next questions를 CLI output 그대로 또는 IDs가
   보존된 faithful summary로 돌려준다.
4. User는 model과 planning을 계속하고 assumptions, non-goals, open blockers를
   visible하게 유지한다.
5. Teammate가 다시 validate한다.
6. Locking은 CLI validation이 blockers를 clear하고 user가 accepted plan을
   confirm한 뒤에만 일어난다.

Teammate가 `ni end`를 실행하면 생성된 `.ni/plan.lock.json`이 source of truth가
된다. 그 뒤에는 locked planning docs를 silently edit하지 않는다.

## Model workspace pack과의 관계

Codex와 Claude skill packs는 conversation을 guide할 수 있다. First-run card를
묻고, docs를 draft하고, contract records를 draft하고, stable IDs를 보존하고, CLI
blockers를 설명하는 데 도움이 된다.

그 pack들은 downstream work를 실행하지 않는다. `ni`의 일부로 model APIs를 호출하지
않는다. CLI readiness를 override하지 않는다. 실제 lock을 만들지 않는다. Global
install은 host-specific loading과 discovery가 verified되기 전까지 Experimental로
남을 수 있다. 현재 Experimental status와 not_verified host/provider boundary는
[Model Workspace Status](99_MODEL_WORKSPACE_STATUS.ko.md)를 참고한다.

## Example

Docs-only example은 `examples/no-terminal-assisted/`를 참고한다. 이 example은
`docs/plan/00_project_brief.md`와 `.ni/contract.json`이 있는 assisted draft를
보여주고, draft를 "not locked"로 표시하며, later CLI validation을 요구한다.

이 example은 의도적으로 `ni status`, `ni end`, `ni run`을 실행하지 않는다. 그렇게
하면 no-terminal assisted drafting이 아니라 trusted CLI workspace를 의미하게 되기
때문이다.

## Boundary

No-terminal planning은 hosted web app, model API calls, runtime execution, shell
adapters, queues, automation behavior를 추가해서는 안 된다. 이것은 kernel boundary를
보존하면서 Intent Lock Protocol을 시작하는 docs-first 방식이다:

```text
model pack or copied guidance -> draft docs and contract
ni CLI -> deterministic readiness, lock, hash proof, prompt compile
```
