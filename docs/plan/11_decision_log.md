# Decision log

## DEC-001: ni is a project intent compiler

Status: accepted

Rationale: The execution harness should be derived from the planning contract, not built first.

## DEC-002: ni run compiles a prompt in v0 and v1

Status: accepted

Rationale: Direct execution would move the project back into harness-first design.

## DEC-003: Codex skills are UX, not authority

Status: accepted

Rationale: Readiness and lock status must be deterministic CLI behavior.

## DEC-004: downstream targets are registry entries

Status: accepted

Rationale: Target-specific prompts and exports need deterministic names and boundaries without adding runtime adapters.

## DEC-005: feedback and pressure are inert until amended

Status: accepted

Rationale: Runtime observations are useful, but they must not silently alter locked acceptance criteria.

## DEC-006: harness candidates are derived proposals

Status: accepted

Rationale: A generated harness may help execution, but the kernel must not own execution state.

## DEC-007: relock requires an applied amendment

Status: accepted

Rationale: Locked planning docs should change only through an explicit user-visible amendment flow.

## DEC-008: collaboration checks are deterministic

Status: accepted

Rationale: Parallel planning changes need review without using model judgment as an authority.

## DEC-009: v0.2 primary authoring UX is model-user conversation, not contract editing commands

Status: accepted

Rationale: `ni init` should create the workspace, while `ni-start` keeps docs and contract synchronized from conversation. The CLI should validate readiness, lock or relock, and compile prompts rather than ask users to hand-author contract records through `add`, `list`, or `set` commands.

## DEC-010: v0.2 differentiation centers Intent Lock Protocol proof assets

Status: accepted

Rationale: The post-053 product direction is that `ni` is the Project Intent Compiler for AI Agents. The core message is "do not run the agent yet; compile the intent first," and the unique mechanism is the Intent Lock Protocol. v0.2 proof should come from ambiguous prompt blocking, non-software planning, benchmark protocol, status proof, downstream target story, README relaunch, README.ko companion sync, and release readiness, not from adding runtime execution.

## DEC-011: v0.3 packaging is README pamphlet, truthful distribution, and model workspace UX

Status: accepted

Rationale: The v0.3 public packaging direction should make `ni` easier to
understand and adopt without weakening the kernel boundary. README is a product
pamphlet with technical details in docs, the hero stays harness-neutral and
SVG-first, generated social imagery is optional, release binaries are required
before non-Go curl or package-manager availability claims, Codex/Claude packs
are planning UX rather than execution adapters, and no-terminal mode remains
assisted unless exact CLI proof is supplied.
