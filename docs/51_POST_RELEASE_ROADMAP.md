# Post-Release Roadmap

This roadmap defines the next phase after the v0.2.0 launch work. It preserves
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

### v0.2.x: launch stabilization

Focus on small post-launch fixes that make the current kernel easier to trust:

- fix launch issues found by source-first users;
- polish documentation around the Intent Lock Protocol, source-of-truth rules,
  and target boundaries;
- expand examples, especially non-software examples and locked handoff samples;
- fix bugs in validation, locking, prompt compilation, target export, and
  command output;
- keep release notes, README links, examples, and public launch docs aligned.

This phase must not add runtime execution behavior. `ni run` remains prompt
compilation only.

### v0.3: conversation authoring UX hardening

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

### v0.4: target seed quality and conformance

Focus on making locked-plan seed material more useful while keeping it inert:

- stabilize target seed formats for built-in targets;
- improve target conformance checks and explanations;
- add clearer handoff packets for human-team and tool-specific consumption
  shapes;
- expand examples that show derived seed material without turning it into
  kernel-owned execution state;
- keep generated work graphs, harness proposals, evaluation notes, and adapter
  notes mutable and downstream-owned.

This phase may improve seed quality. It must not make targets into executable
adapters inside `ni-kernel`.

### v0.5: benchmark data and case studies

Focus on evidence about planning quality without running downstream agents:

- publish real benchmark reports using the existing benchmark protocol;
- compare direct-to-agent prompts against locked `ni` intent for ambiguity,
  traceability, risk coverage, and handoff clarity;
- add human-team handoff evaluation cases;
- add more non-software product examples;
- document where readiness rules helped, where they were noisy, and where they
  need revision.

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
