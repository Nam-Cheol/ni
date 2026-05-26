# Capabilities

## CAP-001: initialize planning workspace

Create planning docs, `.ni/contract.json`, readiness rules, readiness profiles, and inert ledger files for a new project.

## CAP-002: validate readiness

Detect missing docs, invalid contract JSON, missing evaluations, unmitigated high risks, blocker questions, missing non-goals, and profile-specific deferrals.

## CAP-003: lock accepted plan

Create `.ni/plan.lock.json` with hashes for `.ni/contract.json` and required `docs/plan/**` files.

## CAP-004: compile target prompt

Generate a 4000-character-or-less prompt from the locked plan for `generic`, `codex`, `human-team`, or seed-oriented downstream targets.

## CAP-005: provide Codex planning skills

Keep `ni-start`, `ni-end`, and `ni-run` as UX helpers while the CLI remains readiness and lock authority.

## CAP-006: select readiness profile

Support planning readiness profiles for `concept`, `prototype`, `mvp`, `beta`, and `production`.

## CAP-007: preserve related-work boundary

Position `ni` against host enhancers, SPEC-first toolkits, Hyper Run-style evidence loops, heavy harness templates, and design-time compilers.

## CAP-008: model product types and surfaces

Apply product-type, delivery-surface, and interaction-mode guidance without changing readiness authority.

## CAP-009: register downstream targets

List supported prompt and export targets through a deterministic target registry.

## CAP-010: export target seed packages

Write locked seed material for downstream targets after hash verification, without creating runtime packets or calling external tools.

## CAP-011: ingest downstream feedback

Record inert downstream feedback without changing `.ni/contract.json` or `.ni/plan.lock.json`.

## CAP-012: maintain pressure ledger

Turn feedback into visible planning pressure that requires explicit promotion before it can influence an amendment.

## CAP-013: manage harness candidate lifecycle

Propose, validate, accept, or retire generated harness candidates as derived and non-executing material.

## CAP-014: amend and relock locked plans

Require an explicit applied amendment before replacing a valid existing lock.

## CAP-015: detect collaboration conflicts

Compare planning states and report deterministic contract conflicts, including lock mismatches and weakened accepted requirements.
