# Risks and security

## RISK-001: model self-approval

Severity: high

Mitigation: readiness and lock state are CLI-enforced. The model cannot declare completion alone.

## RISK-002: criteria weakening

Severity: high

Mitigation: accepted evaluations, risks, non-goals, and lock hashes must be preserved unless an explicit amendment is applied and relocked.

## RISK-003: host or runtime lock-in

Severity: medium

Mitigation: Codex and downstream tools are targets or UX surfaces. Kernel concepts remain provider-neutral.

## RISK-004: readiness profiles look like execution stages

Severity: high

Mitigation: profiles are documented and validated as planning readiness profiles only, with no runtime packets or agent execution semantics.

## RISK-005: target exports become execution state

Severity: high

Mitigation: exports require a valid lock, verify hashes, avoid external binaries, and write only seed or handoff files.

## RISK-006: feedback mutates accepted truth

Severity: high

Mitigation: feedback ingest must not modify `.ni/contract.json` or `.ni/plan.lock.json`; it can only create inert records for review.

## RISK-007: pressure ledger becomes acceptance shortcut

Severity: high

Mitigation: pressure items require explicit promotion and still need an amendment before accepted planning criteria change.

## RISK-008: harness candidates become kernel runtime

Severity: high

Mitigation: harness candidates must require user acceptance, validation evidence, and a non-execution flag.

## RISK-009: relock bypass

Severity: high

Mitigation: relock refuses to replace an existing lock unless an amendment tied to the current source lock has been applied.

## RISK-010: collaboration checks rely on model judgment

Severity: medium

Mitigation: diff and conflict checks are deterministic and contract-local.

## RISK-011: CLI contract editing displaces conversation authoring

Severity: high

Mitigation: keep v0.2 authoring in `ni-start` conversation updates to `docs/plan/**` and `.ni/contract.json`; do not add contract `add`, `list`, or `set` commands as primary UX.

## RISK-012: differentiation proof drifts into runtime claims

Severity: high

Mitigation: keep positioning, demos, benchmark protocol, status proof, target story, README relaunch, README.ko sync, and release readiness as pre-runtime proof assets. They may validate intent readiness and compile bounded prompts or seed notes, but they must not claim agent execution, shell execution, queues, adapters, or release automation inside `ni-kernel`.

## RISK-013: packaging claims outrun implemented distribution

Severity: high

Mitigation: README, distribution docs, visual assets, and model workspace pack
docs must use factual availability language. Release binaries are required for
non-Go users before curl or package-manager paths are claimed as available;
model packs remain planning UX rather than execution adapters; no-terminal mode
must show CLI-produced proof before claiming deterministic validation. The v0.3
distribution release-state lock keeps README, install, distribution, Homebrew,
manual release, and post-release roadmap docs aligned with source/local
availability, verified v0.3.0 release binary and curl installer availability,
planned Homebrew status, and the pre-runtime kernel boundary.

## RISK-014: visual identity drifts into fragile or overclaiming README surface

Severity: high

Mitigation: Keep primary README visuals local, deterministic, and generated
from checked SVG sources; keep important product copy in Markdown; prohibit
emoji, `foreignObject`, external fonts, external references, and long text in
core SVGs; keep README.ko within English canonical claims; treat remote
capsule-style renderers as inspiration only; and run asset plus README surface
regression checks before relock.

## RISK-015: adoption roadmap turns distribution or promotion into kernel runtime

Severity: high

Mitigation: Keep v0.4 adoption work as documentation, packaging verification,
announcement copy, qualitative benchmark evidence, and optional static landing
page work outside `ni-kernel` runtime behavior. Homebrew stays Planned until the
tap and formula are published and `brew install Nam-Cheol/tap/ni`, `ni --help`,
and `ni version` pass. `ni run` remains a prompt compiler only and must not call
Codex, shell adapters, package publishing, queues, PR automation, or downstream
execution.
