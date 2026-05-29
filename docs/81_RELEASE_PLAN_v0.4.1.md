# v0.4.1 Stabilization Release Plan

Date: 2026-05-29

Status: Draft release plan only. This document does not publish a release,
create a tag, upload assets, mark Homebrew available, or add runtime
execution.

## Goal

Prepare a small stabilization release after v0.4 adoption hardening. The
release should make the existing pre-runtime authoring path easier to trust,
verify, and explain without expanding `ni-kernel` into a runner.

The release story is:

```text
ni init -> ni-start conversation -> ni status proof -> ni end -> ni run prompt
```

`ni run` remains a prompt compiler only. It must not execute Codex, shells,
model APIs, queues, adapters, SPEC workflows, or downstream work.

## Included

| Area | v0.4.1 stabilization scope | Evidence or source |
| --- | --- | --- |
| `ni-start` conversation UX hardening | Make first-run and resume behavior easier to follow: summarize current state, preserve stable IDs, ask one to three focused questions, and show changed files after planning edits. | [Conversation Authoring UX Audit](79_CONVERSATION_AUTHORING_UX_AUDIT.md), [ni-start behavior](31_NI_START_BEHAVIOR.md) |
| Readiness proof clarity | Keep `ni status --proof --next-questions` as the visible readiness explanation. Clarify blockers, deferrals, warnings, passed checks, and the rule that model judgment cannot declare readiness. | [Status proof](44_STATUS_PROOF.md), [Readiness interview](34_READINESS_INTERVIEW.md) |
| Docs/contract sync diagnostics | Stabilize the user-facing `R012` repair language for docs and `.ni/contract.json` drift: affected ID, location, problem, why it matters, suggested repair, and whether it blocks `ni-end`. | [Docs-contract sync](33_DOCS_CONTRACT_SYNC.md) |
| Next-question improvements | Package deterministic `next_questions` as the normal planning loop after `BLOCKED`, with ID-preserving questions and no pressure to accept incomplete decisions. | [Readiness interview](34_READINESS_INTERVIEW.md), [Status proof](44_STATUS_PROOF.md) |
| Model workspace pack UX | Keep Codex and Claude packs as UX layers. Improve task-first usage language around opening a project, invoking `ni-start`, preserving CLI proof, and stopping on lock mismatch. | [Model Workspace Packs](55_MODEL_WORKSPACE_PACKS.md), [Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.md) |
| Example coverage | Preserve checked examples that demonstrate blocked ambiguity, conversation authoring, research protocol planning, conversation products, and bounded handoff prompts without downstream execution. | [Examples](18_EXAMPLES.md), [Demo Verification](48_DEMO_VERIFICATION.md) |
| No-terminal assisted workflow | Improve the proof-capture story for users who cannot run a terminal: draft docs and contract with a model, then paste exact CLI output from a trusted runner before readiness, lock, or prompt claims. | [No-Terminal Planning](no-terminal.md) |
| Benchmark case study expansion | Keep benchmark evidence qualitative and transparent. Extend case-study coverage only as static, documented readiness comparisons; do not call external model APIs or execute downstream work. | [Benchmark Case Studies](77_BENCHMARK_CASE_STUDY.md), [Benchmark Protocol](43_BENCHMARK_PROTOCOL.md) |
| Homebrew decision | Keep Homebrew Planned for v0.4.1 unless the tap, formula, checksums, audit, install, and clean-environment verification are implemented before release. | [Homebrew Decision](80_HOMEBREW_DECISION.md) |

## Not Included

- Execution runtime.
- Codex exec adapter.
- Shell adapter.
- Task runner.
- SPEC runner.
- Queue.
- Multi-agent orchestration.
- Homebrew support if the tap and verification work are not implemented.

These exclusions are part of the release boundary, not missing launch chores.
They keep `ni-kernel` focused on the Intent Lock Protocol: docs contract,
readiness gate, lockfile, prompt compiler, and source-of-truth rule.

## Release Candidate Validation Checklist

Run these checks before tagging or publishing any `v0.4.1` release candidate:

- [ ] `go test ./...`
- [ ] quality: `bash scripts/quality.sh`
- [ ] smoke: `bash scripts/smoke.sh`
- [ ] demo-check: `bash scripts/demo-check.sh`
- [ ] install-check: `bash scripts/install-check.sh`
- [ ] release-check: `bash scripts/release-check.sh`
- [ ] fresh-install-check: `bash scripts/fresh-install-check.sh`
- [ ] skill-pack check: `bash scripts/check-skill-packs.sh`

This checklist is evidence collection only. It does not publish release assets,
push tags, update a Homebrew tap, or mark unverified distribution paths
Available.

## Availability Rules

- Only mark a path Available when the matching implementation and verification
  exist in the repository or published release assets.
- Keep Homebrew Planned unless the v0.4.1 release includes a verified tap
  formula and clean install proof.
- Keep model workspace packs as UX. They may guide planning, but `ni status`,
  `ni end`, and `ni run` remain the authority.
- Keep no-terminal assisted workflow Experimental unless exact CLI proof is
  supplied by a trusted runner.
- Keep benchmark results scoped to the measured case studies. Do not claim
  statistical significance or downstream implementation quality.

## Release Summary Draft

`v0.4.1` is a stabilization release for the adoption-hardening surface. It
polishes the path from `ni init` through `ni-start`, readiness proof review,
lock confirmation, and bounded prompt compilation. The focus is clarity:
readiness proofs should be easier to explain, docs/contract drift should be
easier to repair, next questions should feel like the next planning turn, and
model workspace packs should help users preserve CLI authority.

The release does not change the product category. `ni` remains a Project
Intent Compiler for AI Agents, not a task runner, execution harness, SPEC
runner, adapter layer, queue, or orchestration system.

