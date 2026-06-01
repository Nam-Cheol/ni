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

## DEC-012: v0.3 visual identity is local deterministic SVG plus Markdown copy

Status: accepted

Rationale: README should work as a product pamphlet without making visual
rendering the source of product truth. The hero and core assets use local
deterministic SVG, important product copy remains Markdown text, SVG assets
avoid emoji, `foreignObject`, external fonts, external references, and long
text, README.ko stays within English canonical claims, remote capsule-style
renderers remain inspiration only, and README/assets checks guard regressions.

## DEC-013: v0.3 distribution release state is factual and evidence-gated

Status: accepted

Rationale: The public README and install docs must not claim more distribution
than exists. Source and local binary paths are available from this checkout;
release binaries are Available for the verified v0.3.0 GitHub Release archives
and checksums; curl install is Available for verified v0.3.0 release assets
after `install.sh` verification; Homebrew remains planned with no published tap
or formula; no-terminal mode is assisted unless exact CLI proof is supplied by
a trusted runner; model workspace packs are UX rather than authority; and
`ni-kernel` remains pre-runtime.

## DEC-014: v0.4 adoption roadmap is evidence-gated and pre-runtime

Status: accepted

Rationale: After v0.3.0 release binaries, curl installer verification, and
model workspace pack packaging, the next useful roadmap is adoption proof
rather than runtime expansion. v0.4 should verify a fresh-user install path,
advance Homebrew only while preserving Planned status until the tap/formula is
published and tested, verify model workspace packs, clarify README
use-anywhere paths, prepare announcement copy, publish one qualitative
benchmark case study, and optionally add a lightweight static landing page.
These surfaces help people try and trust `ni`; they do not make `ni-kernel` a
task runner, package publisher, Codex executor, shell adapter, hosted service,
or downstream runtime.

## DEC-015: v0.5 roadmap is evidence-quality, product-surface, and change-control focused

Status: accepted

Rationale: After v0.4 adoption hardening, the next useful roadmap should make
the locked-intent workflow more trustworthy and useful through real benchmark
data, broader product surfaces, stronger conversation authoring reliability,
and clearer lock/change-control UX. Optional Homebrew and landing-page work
remain evidence-gated and factual. Optional downstream integrations belong in
separate packages that consume locked output, not in `ni-kernel` as adapters or
runtime state.

## DEC-016: v0.4.0 post-release state is locked after verification

Status: accepted

Rationale: v0.4.0 release assets, checksums, current-platform binary behavior,
curl installer behavior, and install docs sync have been verified. The public
distribution state can therefore record release binary and curl installer paths
as Available for v0.4.0 while keeping Homebrew Planned / deferred, model
workspace packs as UX rather than authority, and `ni-kernel` pre-runtime.

## DEC-017: v0.5 roadmap lock follows evidence, reliability, and adoption hardening

Status: accepted

Rationale: v0.5 is not primarily a package-manager release. It is focused on
real benchmark evidence, conversation-authoring reliability, ni-grill planning
challenge quality, change-control UX, factual distribution/adoption surfaces,
non-software product-surface expansion, and downstream integration boundaries.
Homebrew remains a v0.5 candidate and can become Available only after tested
tap/formula evidence. Downstream execution remains outside `ni-kernel`.
GRILL-001 and GRILL-002 are addressed by roadmap and distribution alignment;
GRILL-003 remains a v0.5 acceptance-evidence improvement candidate; GRILL-004
remains a benchmark claim-boundary note; and GRILL-005 remains a model
workspace status preservation note.
