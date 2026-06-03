# Post-Release Roadmap

이 roadmap은 verified v0.5.0 release, release-asset verification,
curl-installer verification work 이후의 다음 단계들을 정의한다. `ni`는 계속 `ni-kernel`이어야 한다. 즉 AI Agents를 위한 Project
Intent Compiler이며, accepted project intent를 위한 deterministic pre-runtime
control layer다.

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

### v0.4.x: post-release stabilization

현재 kernel을 더 신뢰하기 쉽게 만드는 작은 post-release fixes에 집중한다:

- source, local binary, release binary, curl installer, model-workspace,
  no-terminal-assisted users가 발견한 adoption/documentation issues 수정;
- verified v0.4.0 state에 맞게 release, curl installer, install, README,
  verification, distribution docs를 정확하게 유지;
- benchmark evidence를 과장하지 않으면서 examples와 benchmark readability 개선;
- Intent Lock Protocol, source-of-truth rules, target boundaries 관련 docs polish;
- validation, locking, prompt compilation, target export, command output의 bug
  fixes.

이 phase는 runtime execution behavior를 추가하면 안 된다. `ni run`은 계속 prompt
compilation only다.

### v0.4: conversation authoring UX hardening

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

### v0.5: evidence, authoring reliability, and adoption surfaces

현재 pre-runtime kernel을 더 credible하고 adopt하기 쉽게 만들며, real planning
evidence로 뒷받침하는 데 집중한다:

- `not_measured` boundaries를 보존하고 fake empirical claims, statistical
  significance claims, implementation-quality claims, downstream-agent-performance
  claims를 하지 않는 real benchmark evidence와 case studies 공개;
- docs/contract/session synchronization, grouped repair questions, proof capture,
  assumptions/decisions/risks/evaluations/non-goals 보존을 포함한
  conversation-authoring reliability 개선;
- `ni-grill`을 `ni` planning에 dogfood하고, `ni status` readiness authority를
  대체하지 않으면서 planning challenge quality 개선;
- locked-plan change control, amendment, relock, changed-intent UX를 개선하되
  lock/hash verification은 deterministic하게 유지;
- tap 또는 formula가 존재하고 `brew install`, `ni --help`, `ni version`이
  tested된 뒤에만 Homebrew를 optional distribution candidate로 검토;
- host-level install 또는 discovery가 proof될 때만 model workspace packs를
  verify하고, 그렇지 않으면 Experimental과 CLI-authority boundary를 유지;
- 특히 non-software planning examples 같은 product surfaces 확장;
- downstream integrations는 separate packages, target exports, seed formats,
  downstream-owned notes로만 유지하고 `ni-kernel` behavior로 만들지 않기.

v0.5 task가 completion을 claim하기 전에는
[`95_V0_5_ACCEPTANCE_EVIDENCE.md`](95_V0_5_ACCEPTANCE_EVIDENCE.md)의 acceptance
evidence matrix를 사용한다. 이 matrix는 lane별 minimum files, commands, review
proof, status vocabulary, `not_measured` boundaries를 정의한다.

첫 세 개 v0.5 work packet 이후에는
[`100_V0_5_WORK_PACKET_COMPLETION_AUDIT.ko.md`](100_V0_5_WORK_PACKET_COMPLETION_AUDIT.ko.md)를
GRILL-003부터 GRILL-005까지의 closure record와 selected next direction으로
사용한다.

Conversation proof capture reliability pass에는
[`101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md`](101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md)를
사용해 planning proof, CLI authority, no-terminal draft limits, benchmark
boundaries, model workspace skill wording을 aligned 상태로 유지한다.

Locked-plan change-control UX에는
[`102_CHANGE_CONTROL_UX_AUDIT.ko.md`](102_CHANGE_CONTROL_UX_AUDIT.ko.md)를 사용해
diagnostics 또는 lock behavior를 바꾸기 전에 intended stale-lock, amended-intent,
relock, downstream handoff safety model을 보존한다.

구현된 stale-lock CLI diagnostic에는
[`103_STALE_LOCK_DIAGNOSTIC.ko.md`](103_STALE_LOCK_DIAGNOSTIC.ko.md)를 사용해
`LOCK-STALE` wording, recovery flow, authority boundaries, test coverage를
보존한다.

`LOCK-STALE` 이후 practical amend/relock workflows에는
[`104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md`](104_AMEND_RELOCK_WORKFLOW_EXAMPLES.ko.md)를
사용해 user examples, CLI recovery order, skill boundaries, no-terminal limits,
non-execution claims를 보존한다.

Model workspace stale-lock wording에는
[`105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md`](105_MODEL_WORKSPACE_STALE_LOCK_WORDING_VERIFICATION.ko.md)를
사용해 Experimental status, host-verification limits, skill drafting-only
boundaries, CLI-authoritative recovery order를 보존한다.

No-terminal stale-lock examples는
[`106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md`](106_NO_TERMINAL_STALE_LOCK_EXAMPLES.ko.md)를
사용해 model-only drafts, pasted CLI output, trusted runner transcripts를 구분하고
deterministic no-terminal validation claim을 피한다.

No-terminal transcript quality에는
[`107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.ko.md`](107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.ko.md)를
사용해 unusable, partial, pasted, trusted runner, fixture-only, target-workspace
transcripts를 구분하고 fixture evidence를 project-root state로 취급하지 않는다.

Broader changed-intent fixture coverage에는
[`108_CHANGED_INTENT_FIXTURE_COVERAGE.ko.md`](108_CHANGED_INTENT_FIXTURE_COVERAGE.ko.md)를
사용해 representative stale-lock cases, non-lockable false positives, fixture
relock recovery, project-root safety를 lock semantics 변경 없이 verify한다.

Reliability docs release-readiness sweep에는
[`109_V0_5_RELIABILITY_DOCS_RELEASE_READINESS_SWEEP.ko.md`](109_V0_5_RELIABILITY_DOCS_RELEASE_READINESS_SWEEP.ko.md)를
사용해 v0.5 reliability set 전반의 consistency, navigation, Korean companion
parity, validation surface, overclaim boundaries를 verify한다.

v0.5 release-candidate readiness decision에는
[`110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md`](110_V0_5_RELEASE_CANDIDATE_READINESS_AUDIT.ko.md)를
사용해 release-note polish 전에 `RC_READY_WITH_DEFERRALS` decision, validation
evidence, explicit deferrals, claim boundaries를 보존한다.

v0.5 RC polish와 release-note wording에는
[`111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md`](111_V0_5_RC_POLISH_RELEASE_NOTES_DRAFT.ko.md)를
사용해 release action 전에 draft-only release-note language, known deferrals,
validation commands, release-note claim audit을 보존한다.

Final v0.5 release-note preflight에는
[`112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md`](112_V0_5_RELEASE_NOTES_FINAL_PREFLIGHT.ko.md)를
사용해 release-note wording, adjacent RC docs, Korean companion parity, git
visibility, validation commands, forbidden release or availability claims를
verify한다.

v0.5 artifact dry-run audit에는
[`113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md`](113_V0_5_ARTIFACT_DRY_RUN_AUDIT.ko.md)를
사용해 release action 전에 dry-run/check-only evidence로 release artifact,
install, build, version/help, script, hosted availability boundaries를 verify한다.

v0.5 release publication checklist에는
[`114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md`](114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md)를
사용해 future human-controlled release actions를 gated 상태로 유지하고,
checklist task에서 explicitly not run으로 표시하며, Homebrew, hosted-artifact,
cross-platform verification deferrals와 분리한다.

v0.5 publication human approval packet에는
[`115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.ko.md`](115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.ko.md)를
사용해 final human decision을 explicit하고 Codex가 unselected로 둔 상태로 유지하며,
tag, GitHub release creation, asset upload, checksum publication, Homebrew work,
generated prompt execution, availability-claim upgrade와 분리한다.

README onboarding repair and visual prompt pack에는
[`116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.ko.md`](116_README_ONBOARDING_AND_VISUAL_PROMPT_PASS.ko.md)를
사용해 DO_NOT_APPROVE_FIX_FIRST fix-first context, bounded macOS and Windows
install/uninstall wording, README image prompt reproducibility,
non-execution claim boundaries를 publication approval 재검토 전에 보존한다.

v0.5.0 post-release verification에는
[`117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md`](117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md)를
사용해 actual release evidence, asset and checksum verification,
current-platform binary output, curl installer proof, README/install sync,
remaining deferrals를 보존한다. Homebrew, model workspace, no-terminal,
benchmark claims는 이 문서만으로 upgrade하지 않는다.

Homebrew implementation audit에는
[`118_HOMEBREW_IMPLEMENTATION_AUDIT.ko.md`](118_HOMEBREW_IMPLEMENTATION_AUDIT.ko.md)를
사용해 local draft formula, source archive sha256 evidence, tap absence proof,
local Homebrew validation boundaries, future availability gate를 보존한다. Tap
publication 또는 README `brew install` wording 전까지 이 gate를 유지한다.

이 phase는 supporting work로 target seed quality와 conformance를 개선할 수 있다.
하지만 targets를 `ni-kernel` 내부의 executable adapters로 만들면 안 된다.

v0.5는 Homebrew tap implementation을 distribution infrastructure로 다룰 수 있는
가장 이른 scheduled point이기도 하다. External tap, formula, checksums, audit,
local formula install, published tap install, `ni --help` / `ni version`
validation이 모두 통과하기 전까지 Homebrew는 Planned and deferred로 남는다.

### v0.6 or later: broader adoption evidence and ecosystem work

v0.5 credibility baseline 이후의 evidence와 optional ecosystem work에 집중한다:

- v0.5 case studies가 유용한 measurement patterns를 보여준 경우 broader
  benchmark data 공개;
- README와 install docs만으로 adoption이 부족한 경우에만 landing page 검토;
- downstream package ecosystem은 `ni-kernel` 밖에서만 탐색;
- real users와 maintained examples에서 stronger adoption evidence 추가;
- human-team handoff evaluation cases 추가;
- readiness rules가 도움이 된 곳, noisy했던 곳, revision이 필요한 곳 문서화;
- `ni-kernel` runtime execution은 계속 제외.

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
