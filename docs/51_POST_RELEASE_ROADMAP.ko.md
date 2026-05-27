# Post-Release Roadmap

이 roadmap은 v0.2.0 launch work 이후의 다음 단계를 정의한다. `ni`는 계속
`ni-kernel`이어야 한다. 즉 AI Agents를 위한 Project Intent Compiler이며,
accepted project intent를 위한 deterministic pre-runtime control layer다.

이 roadmap은 방향성 문서이며, 구현된 support를 claim하지 않는다. Future work는
명시적으로 downstream 또는 separate-package integration work로 분리되지 않는 한
kernel boundary 안에 머물러야 한다.

## Boundary

`ni-kernel`은 다음을 계속 개선할 수 있다:

- planning docs와 contract synchronization;
- deterministic readiness validation;
- lockfile integrity와 source-of-truth checks;
- bounded prompt compilation;
- inert downstream seed generation;
- target conformance explanations와 examples.

`ni-kernel`은 runtime execution, agent orchestration, queues, adapters,
evidence collection loops, release workflows의 owner가 되면 안 된다.

## Phases

### v0.2.x: launch stabilization

현재 kernel을 더 신뢰하기 쉽게 만드는 작은 post-launch fixes에 집중한다:

- source-first users가 발견한 launch issues 수정;
- Intent Lock Protocol, source-of-truth rules, target boundaries 관련 docs polish;
- examples 확장, 특히 non-software examples와 locked handoff samples;
- validation, locking, prompt compilation, target export, command output의 bug
  fixes;
- release notes, README links, examples, public launch docs 정렬 유지.

이 phase는 runtime execution behavior를 추가하면 안 된다. `ni run`은 계속 prompt
compilation only다.

### v0.3: conversation authoring UX hardening

sustained model-user planning을 더 안전하고 auditable하게 만드는 데 집중한다:

- ambiguous, conflicting, tentative, inferred planning records를 위한 readiness
  rules 개선;
- plan이 blocked, ready, ready with deferrals인 이유를 사용자가 볼 수 있도록
  `ni status` proof explanations 개선;
- docs/contract sync checks와 diagnostics 강화;
- non-goals, risks, mitigations, evaluations, blocker questions가 editing turns
  사이에서 보존되도록 개선;
- companion docs가 유지되는 영역에서 Korean/English documentation parity checks
  개선.

CLI가 계속 authority다. Skills와 models는 계속 UX다.

### v0.4: target seed quality and conformance

locked-plan seed material을 더 유용하게 만들되 inert 상태를 유지하는 데 집중한다:

- built-in targets의 target seed formats 안정화;
- target conformance checks와 explanations 개선;
- human-team 및 tool-specific consumption shapes를 위한 더 명확한 handoff
  packets 추가;
- derived seed material을 보여주되 kernel-owned execution state로 바꾸지 않는
  examples 확장;
- generated work graphs, harness proposals, evaluation notes, adapter notes를
  mutable and downstream-owned 상태로 유지.

이 phase는 seed quality를 개선할 수 있다. 하지만 targets를 `ni-kernel` 내부의
executable adapters로 만들면 안 된다.

### v0.5: benchmark data and case studies

downstream agents를 실행하지 않고 planning quality에 관한 evidence를 쌓는 데
집중한다:

- existing benchmark protocol을 사용해 real benchmark reports 공개;
- direct-to-agent prompts와 locked `ni` intent를 ambiguity, traceability, risk
  coverage, handoff clarity 기준으로 비교;
- human-team handoff evaluation cases 추가;
- 더 많은 non-software product examples 추가;
- readiness rules가 도움이 된 곳, noisy했던 곳, revision이 필요한 곳 문서화.

Benchmarks는 intent quality와 handoff readiness를 평가해야 한다. execution
benchmarks나 runtime performance claims가 되면 안 된다.

### Later: optional downstream integrations

Later integrations는 downstream packages, experiments, 또는 separate repositories로만
존재할 수 있다. 이들은 locked `ni` output을 소비해야 하며 kernel-owned execution
state가 되면 안 된다.

Possible future packages는 다음을 탐색할 수 있다:

- `ni-kernel` 밖의 tool-specific adapters;
- locked seed packages를 읽는 downstream harnesses;
- external evidence collection flows;
- separate package release processes 주변의 optional automation.

이 항목들은 committed kernel features가 아니다. Kernel은 deterministic
validation, locking, bounded prompt compilation, inert seed export에서 멈춘다는
규칙을 바꾸면 안 된다.

## Still Forbidden In Core

다음은 `ni-kernel` responsibilities로 여전히 forbidden이다:

- task runner;
- SPEC runner;
- Codex exec adapter;
- shell adapter;
- queue;
- multi-agent orchestration;
- PR automation;
<!-- ni-boundary-allow: explicit negative boundary list item. -->
- no release automation;
<!-- ni-boundary-allow: explicit negative boundary list item. -->
- execution evidence loop.

이 항목들이 유용해지더라도 downstream 또는 separate packages에 속한다. `ni run`
behavior, lockfile state, source-of-truth state, kernel-owned lifecycle state가
되면 안 된다.

## Research Directions

Recommended next research directions:

- better readiness rules;
- better status proof explanations;
- stronger docs/contract sync;
- more non-software product examples;
- human-team handoff evaluation;
- real benchmark reports;
- target seed format stability;
- Korean/English doc parity checks.

각 research direction은 Intent Lock Protocol을 보존해야 한다. Planning
conversation은 explicit contract가 되고, deterministic gates가 readiness를
결정하며, accepted intent는 locked and hashed 상태가 되고, intent가 바뀌면
downstream handoff가 멈춰야 한다.
