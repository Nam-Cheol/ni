# v0.5 Acceptance Evidence

## Purpose

This document defines what counts as completion evidence for v0.5 roadmap
lanes. It turns the broad v0.5 direction into inspectable evidence expectations
before later work claims that a lane is done.

This is planning acceptance evidence. It is not execution evidence, does not
prove downstream implementation quality, and does not authorize runtime
execution. CLI readiness and lock state remain authoritative for NI plans:
`ni status` decides readiness, `ni end` writes locks, and `ni run` compiles
bounded prompts from valid locks.

## Evidence principles

1. Evidence must be inspectable.
2. Evidence must reference files, commands, or documented decisions.
3. Evidence must preserve `not_measured` boundaries.
4. Evidence must distinguish `Available`, `Experimental`, and `Planned`.
5. Evidence must not turn model judgment into authority.
6. Evidence must not weaken readiness gates.
7. Evidence must not imply downstream execution.

## Evidence matrix

| Lane | Completion claim | Required evidence | Verification command or file | Status vocabulary | Not measured / not claimed |
| --- | --- | --- | --- | --- | --- |
| Real benchmark evidence | v0.5 benchmark cases show how NI exposes readiness gaps before handoff. | Case workspace or docs; status proof; next questions; before/after evidence if resolved; `not_measured` table; no downstream execution statement; visible claim-boundary markers. | `docs/77_BENCHMARK_CASE_STUDY.md`; `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`; `examples/benchmark-report/**`; `bash scripts/demo-check.sh` | `measured`; `not_measured`; `BLOCKED`; `READY` for artifact-readiness only | Implementation quality; downstream agent performance; adoption; cost; latency; statistical effect size |
| Conversation-authoring reliability | `ni-start` / `ni-grill` / `ni-end` flow is easier to follow and safer. | First-run card docs; grouped next questions; planning proof capture; examples; skill pack wording. | `docs/31_NI_START_BEHAVIOR.md`; `docs/83_CONVERSATION_PROOF_CAPTURE.md`; `examples/ni-start-dogfood/**`; `bash scripts/check-skill-packs.sh` | `documented`; `verified by examples`; `Experimental` | Real user success rate; model quality across providers |
| ni-grill quality | `ni-grill` challenges planning quality without becoming a readiness gate. | `docs/91_NI_GRILL.md`; `docs/92_NI_GRILL_OUTPUT_BUDGET.md`; ni-grill examples; severity and budget rules; CLI authority boundary. | `examples/ni-grill/**`; `bash scripts/demo-check.sh`; `bash scripts/check-skill-packs.sh` | `advisory`; `severity-labeled`; `CLI decides` | Real reduction in planning defects; user satisfaction |
| Change-control UX | Locked plan changes are easier to understand and audit. | Amended docs; amendment record; status proof; relock behavior; stale lock or hash mismatch examples. | Future docs, tests, or examples that show amendment/relock and stale-lock proof | `Planned`; `audited`; `verified` | Team-scale merge success; production workflow adoption |
| Homebrew | Homebrew is available only after tested tap/formula evidence. | Tap repository; formula; sha256; `brew install` output; `ni --help`; `ni version`. | `docs/80_HOMEBREW_DECISION.md`; future Homebrew verification doc | Current status: `Planned` / v0.5 candidate; later `Experimental` or `Available` only with evidence | Package-manager adoption; long-term formula maintenance |
| Model workspace packs | Skill packs are easier to install or use in supported model workspaces. | Package scripts; zip contents; `SKILL.md` files; status-preservation doc; host-level install verification if claiming `Available`. | `bash scripts/check-skill-packs.sh`; `bash scripts/package-claude-skills.sh`; `bash scripts/package-codex-skills.sh`; `docs/75_MODEL_PACK_INSTALL_VERIFICATION.md`; `docs/99_MODEL_WORKSPACE_STATUS.md` | `Experimental`; `Available` only for a specific host path after host-level install and usage verification | Model provider behavior; global host compatibility unless tested |
| No-terminal assisted workflow | Users can start drafting without CLI, but trusted validation still requires CLI. | `docs/no-terminal.md`; `examples/no-terminal-assisted/**`; clear not-deterministic warning. | `docs/no-terminal.md`; `bash scripts/demo-check.sh` | `Experimental`; `assisted`; `not deterministic` | Trusted readiness; lock; hash verification; prompt compilation without CLI |
| Product surface expansion | NI supports more than software planning. | Examples for research protocol, operations process, education program, document product, physical product planning, or similar; status proof or docs-only boundary. | `examples/**`; `docs/82_EXAMPLE_COVERAGE.md` | `example-backed`; `docs-only`; `measured` | Real-world adoption; downstream implementation quality |
| Downstream integrations | Downstream integrations remain separate packages, seed exports, or target formats. | Separate package or docs; no `ni-kernel` runtime execution; target/export conformance. | `docs/45_TARGET_STORY.md`; target conformance docs/tests when present; `bash scripts/smoke.sh` | `seed`; `export`; `separate package`; `not ni-kernel runtime` | Downstream execution success |

## Claim boundary rules

- Do not say "improves implementation quality" unless implementation was measured.
- Do not say "reduces rework" unless repeated trials measured it.
- Do not say `Available` unless install or verification evidence exists.
- Do not say `READY` means the product is ready unless the readiness scope says that.
- Do not say model workspace packs are global unless host-level install was verified.
- Do not say no-terminal is deterministic.

## GRILL-003 closure

GRILL-003 is addressed by this acceptance evidence matrix. That does not mean
every v0.5 lane is complete. It means each lane now has clearer completion
evidence requirements that future tasks must satisfy before claiming completion.

GRILL-004 is addressed by `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`, which keeps
benchmark claim limits visible next to the case-study and example surfaces.
GRILL-005 is addressed by `docs/99_MODEL_WORKSPACE_STATUS.md`, which preserves
the Experimental broad product path, names verified repository evidence, and
keeps host-level/global install and provider behavior as unverified unless
documented.

## How to use this document

Before starting a v0.5 task:

1. Identify the task's lane in the evidence matrix.
2. Check the required evidence before editing.
3. Keep `not_measured` boundaries visible next to any claim.
4. Update examples, docs, or tests before claiming completion.
5. Use `ni-grill` to challenge whether the evidence is sufficient.
6. Use `ni status` for deterministic readiness; do not substitute model judgment.
