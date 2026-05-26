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
