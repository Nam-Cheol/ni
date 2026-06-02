# Stale Lock Diagnostic

## Current status

ni는 이제 CLI diagnostic surfaces에서 stale existing locks를 surface한다.
Diagnostic label은 `LOCK-STALE`이다.

`ni status --proof --next-questions`는 current planning contract에 대해 여전히
`READY`를 report할 수 있다. 하지만 existing `.ni/plan.lock.json`이 current
lockable planning inputs와 더 이상 match하지 않으면 warning을 표시한다.

## What stale means

Stale lock은 다음을 뜻한다:

- `.ni/plan.lock.json`이 존재한다; 그리고
- current lockable planning inputs가 `.ni/plan.lock.json`에 기록된 hashes와 다른
  hashes를 만든다.

이 비교는 `ni run`이 사용하는 것과 같은 lock verification path를 사용한다.
Lockable inputs는 `.ni/contract.json`과 lock에 기록된 required `docs/plan/**`
files다. `.ni/session.json`은 locked docs 아래의 mutable continuity state로
남으며, current lock semantics에서는 hash되지 않는다.

## What stale does not mean

Stale은 다음을 증명하지 않는다:

- implementation failure;
- product unreadiness;
- benchmark failure;
- downstream execution failure;
- adoption, cost, latency, quality, downstream agent performance.

이는 current planning inputs와 existing locked plan이 더 이상 match하지 않는다는
뜻일 뿐이다. 따라서 downstream handoff는 stale lock에 의존하면 안 된다.

## User-facing behavior

`ni status --proof --next-questions`는 readiness semantics와 lock freshness를
분리해서 유지한다. Current planning contract가 ready이면 readiness는 `READY`로
남을 수 있지만, stale existing lock은 `Warnings` 아래에 `LOCK-STALE`로 표시된다.

Warning text는 다음과 같다:

```text
WARNING: LOCK-STALE existing lock is stale. Current planning inputs differ from .ni/plan.lock.json.
```

가능하면 diagnostic은 첫 번째 mismatched lockable input path와 recovery action도
함께 표시한다.

`ni end`는 lock이 없을 때 normal first-lock flow를 유지한다. Existing lock이
current이면 lock이 current이고 `ni end`가 CLI readiness flow를 통해 refresh한다고
말한다. Existing lock이 stale이면 changed intent가 review된 뒤 `ni end`가
CLI-authoritative relock step이라고 말한다.

`ni run`은 stale handoff를 계속 refuse한다. 자동으로 relock하지 않고,
`.ni/plan.lock.json`을 mutate하지 않으며, downstream work를 execute하지 않는다.
Stale-lock refusal에는 recovery guidance가 포함된다.

## Recovery flow

```text
review changed intent
-> run ni status --proof --next-questions
-> run ni end
-> run ni run --max-chars 4000
```

Review step은 human 및 planning-state work다. Intent가 바뀌면 `docs/plan/**`와
`.ni/contract.json`을 함께 update하고, 그 다음 CLI gates에 의존한다.
Practical user workflows는
[`104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md`](104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md)를
참고한다.
No-terminal assisted examples는
[`106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md`](106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md)를
참고한다.

## Authority boundary

- CLI is authority.
- Skills are UX.
- Skills do not determine readiness.
- Skills do not lock or relock.
- No-terminal assisted workflow does not provide deterministic validation.
- `ni run` compiles a bounded handoff prompt and does not execute downstream
  work.

## Test coverage

이 task는 stale-lock에 대해 다음 focused coverage를 추가한다:

- no lock exists: `ni status --proof --next-questions`가 `LOCK-STALE`를 emit하지
  않는다;
- lock exists and is current: `ni status --proof --next-questions`가
  `LOCK-STALE`를 emit하지 않는다;
- lock exists and planning inputs change: `ni status --proof --next-questions`가
  `LOCK-STALE`를 emit한다;
- lock exists and planning inputs change: `ni run`이 stale handoff를 refuse하고
  recovery guidance를 포함한다;
- temporary fixture 안에서 relock 후: stale warning이 사라지고 `ni run`이 bounded
  handoff를 다시 compile한다.

## Remaining follow-up candidates

- Broader changed-intent fixtures 추가.
- No-terminal stale-lock explanation 개선.
- Model workspace stale-lock wording verification 추가.
