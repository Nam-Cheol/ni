# Post-Release Roadmap

This roadmap defines the next phases after the v0.4.0 release, release-asset
verification, and curl-installer verification work. It preserves
`ni` as `ni-kernel`: a Project Intent Compiler for AI Agents and a
deterministic pre-runtime control layer for accepted project intent.

The roadmap is directional, not a claim of implemented support. Future work
must stay inside the kernel boundary unless it is explicitly packaged as
downstream or separate-package integration work.

## Boundary

`ni-kernel` may keep improving:

- planning docs and contract synchronization;
- deterministic readiness validation;
- lockfile integrity and source-of-truth checks;
- bounded prompt compilation;
- inert downstream seed generation;
- target conformance explanations and examples.

`ni-kernel` must not become the owner of runtime execution, agent orchestration,
queues, adapters, evidence collection loops, or release workflows.

## Phases

### v0.4.x: post-release stabilization

Focus on small post-release fixes that make the current kernel easier to trust:

- fix adoption and documentation issues found by source, local binary, release
  binary, curl installer, model-workspace, or no-terminal-assisted users;
- keep release, curl installer, install, README, verification, and
  distribution docs accurate for the verified v0.4.0 state;
- improve examples and benchmark readability without overstating benchmark
  evidence;
- polish documentation around the Intent Lock Protocol, source-of-truth rules,
  and target boundaries;
- fix bugs in validation, locking, prompt compilation, target export, and
  command output.

This phase must not add runtime execution behavior. `ni run` remains prompt
compilation only.

### v0.4: conversation authoring UX hardening

Focus on making sustained model-user planning safer and more auditable:

- improve readiness rules for ambiguous, conflicting, tentative, or inferred
  planning records;
- improve `ni status` proof explanations so users can see why a plan is
  blocked, ready, or ready with deferrals;
- strengthen docs/contract sync checks and diagnostics;
- make non-goals, risks, mitigations, evaluations, and blocker questions easier
  to preserve across editing turns;
- improve Korean and English documentation parity checks where companion docs
  are maintained.

The CLI remains the authority. Skills and models remain UX.

### v0.5: evidence, authoring reliability, and adoption surfaces

Focus on making the current pre-runtime kernel more credible, easier to adopt,
and better supported by real planning evidence:

- publish real benchmark evidence and case studies that preserve
  `not_measured` boundaries and make no fake empirical, statistical
  significance, implementation-quality, or downstream-agent-performance claims;
- improve conversation-authoring reliability, especially docs/contract/session
  synchronization, grouped repair questions, proof capture, and preservation of
  assumptions, decisions, risks, evaluations, and non-goals;
- dogfood `ni-grill` against `ni` planning and improve planning challenge
  quality without replacing `ni status` as readiness authority;
- improve locked-plan change control, amendment, relock, and changed-intent UX
  while keeping lock and hash verification deterministic;
- consider Homebrew only as an optional distribution candidate after a tap or
  formula exists and `brew install`, `ni --help`, and `ni version` are tested;
- verify model workspace packs only where host-level install or discovery can
  be proved; otherwise keep them Experimental and CLI-authority bounded;
- expand product surfaces, especially non-software planning examples;
- keep downstream integrations as separate packages, target exports, seed
  formats, or downstream-owned notes rather than `ni-kernel` behavior.

Before a v0.5 task claims completion, use the acceptance evidence matrix in
[`95_V0_5_ACCEPTANCE_EVIDENCE.md`](95_V0_5_ACCEPTANCE_EVIDENCE.md). The matrix
defines the minimum files, commands, review proof, status vocabulary, and
`not_measured` boundaries for each lane.

After the first three v0.5 work packets, use
[`100_V0_5_WORK_PACKET_COMPLETION_AUDIT.md`](100_V0_5_WORK_PACKET_COMPLETION_AUDIT.md)
for the GRILL-003 through GRILL-005 closure record and the selected next
direction.

For the conversation proof capture reliability pass, use
[`101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md`](101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md)
to keep planning proof, CLI authority, no-terminal draft limits, benchmark
boundaries, and model workspace skill wording aligned.

For locked-plan change-control UX, use
[`102_CHANGE_CONTROL_UX_AUDIT.md`](102_CHANGE_CONTROL_UX_AUDIT.md) to preserve
the intended stale-lock, amended-intent, relock, and downstream handoff safety
model before changing diagnostics or lock behavior.

For the implemented stale-lock CLI diagnostic, use
[`103_STALE_LOCK_DIAGNOSTIC.md`](103_STALE_LOCK_DIAGNOSTIC.md) to preserve
`LOCK-STALE` wording, recovery flow, authority boundaries, and test coverage.

For practical amend/relock workflows after `LOCK-STALE`, use
[`104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md`](104_AMEND_RELOCK_WORKFLOW_EXAMPLES.md)
to preserve user examples, CLI recovery order, skill boundaries, no-terminal
limits, and non-execution claims.

For no-terminal stale-lock examples, use
[`106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md`](106_NO_TERMINAL_STALE_LOCK_EXAMPLES.md)
to distinguish model-only drafts, pasted CLI output, and trusted runner
transcripts without claiming deterministic no-terminal validation.

For no-terminal transcript quality, use
[`107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md`](107_NO_TERMINAL_TRANSCRIPT_QUALITY_CHECKLIST.md)
to distinguish unusable, partial, pasted, trusted runner, fixture-only, and
target-workspace transcripts without treating fixture evidence as project-root
state.

This phase may improve target seed quality and conformance as supporting work.
It must not make targets into executable adapters inside `ni-kernel`.

v0.5 is also the earliest scheduled point for Homebrew tap implementation as
distribution infrastructure. Homebrew remains Planned and deferred until the
external tap, formula, checksums, audit, local formula install, published tap
install, and `ni --help` / `ni version` validation have all passed.

### v0.6 or later: broader adoption evidence and ecosystem work

Focus on evidence and optional ecosystem work that should follow the v0.5
credibility baseline:

- publish broader benchmark data if v0.5 case studies show useful measurement
  patterns;
- add a landing page only if README and install docs are insufficient for
  adoption;
- explore a downstream package ecosystem only outside `ni-kernel`;
- add stronger adoption evidence from real users and maintained examples;
- add human-team handoff evaluation cases;
- document where readiness rules helped, where they were noisy, and where they
  need revision;
- continue to exclude `ni-kernel` runtime execution.

Benchmarks should evaluate intent quality and handoff readiness. They must not
become execution benchmarks or runtime performance claims.

### Later: optional downstream integrations

Later integrations may exist only as downstream packages, experiments, or
separate repositories. They must consume locked `ni` output rather than becoming
kernel-owned execution state.

Possible future packages may explore:

- tool-specific adapters outside `ni-kernel`;
- downstream harnesses that read locked seed packages;
- external evidence collection flows;
- optional automation around separate package release processes.

These are not committed kernel features. They must not change the rule that the
kernel stops at deterministic validation, locking, bounded prompt compilation,
and inert seed export.

## Still Forbidden In Core

The following remain forbidden as `ni-kernel` responsibilities:

- task runner;
- SPEC runner;
- Codex exec adapter;
- shell adapter;
- queue;
- multi-agent orchestration;
- PR automation;
- no release automation;
- execution evidence loop.

If any of these become useful, they belong downstream or in separate packages.
They must not become `ni run` behavior, lockfile state, source-of-truth state, or
kernel-owned lifecycle state.

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

Each research direction should preserve the Intent Lock Protocol: planning
conversation becomes explicit contract, deterministic gates decide readiness,
accepted intent is locked and hashed, and downstream handoff stops when intent
changes.
