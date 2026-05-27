# Capabilities

## CAP-001: initialize conversation planning workspace

`ni init` creates the initial `docs/plan/**` structure, `.ni/contract.json`, readiness config, and bounded session state for conversation-driven authoring.

## CAP-002: report deterministic readiness gaps

Detect missing docs, invalid contract JSON, missing evaluations, unmitigated high risks, blocker questions, missing non-goals, and profile-specific deferrals.

## CAP-003: confirm and lock accepted plan

`ni-end` confirms the accepted plan with the user, then the CLI creates `.ni/plan.lock.json` with hashes for `.ni/contract.json` and required `docs/plan/**` files.

## CAP-004: compile bounded handoff prompt

`ni-run` compiles a 4000-character-or-less handoff prompt from the locked plan for `generic`, `codex`, `human-team`, or seed-oriented downstream targets. It does not execute Codex, shell commands, adapters, or downstream runtimes.

## CAP-005: author planning state through model-user conversation

After `ni init`, use sustained model-user conversation as the primary docs authoring interface. `ni-start` maintains `docs/plan/**`, `.ni/contract.json`, and bounded `.ni/session.json` continuity state together. `ni status` supplies deterministic readiness gaps, `ni-end` confirms and locks through CLI authority, and `ni-run` compiles a handoff prompt. User-facing contract `add`, `list`, or `set` commands are not part of the v0.2 primary UX.

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

## CAP-016: prove Project Intent Compiler differentiation

Maintain the v0.2 positioning and proof assets that show `ni` as a pre-runtime Project Intent Compiler for AI Agents. The plan must tie the message "do not run the agent yet; compile the intent first" to the Intent Lock Protocol, the ambiguous prompt blocking demo, the non-software planning demo, the intent readiness benchmark protocol, the status proof report, the downstream target story, README relaunch, README.ko companion sync, and release readiness checklist.

## CAP-017: package v0.3 public product surface

Lock the v0.3 product packaging direction across README pamphlet strategy,
visual identity, distribution, and model workspace packs. README is a product
pamphlet with technical depth moved to docs, no specific harness product claims
in the hero, SVG-first visual assets, optional generated social images, factual
distribution language for release binaries, curl, and package managers, Codex-
and Claude-style model workspace packs as planning UX, and no-terminal mode as
assisted rather than deterministic validation.

The visual identity portion of this capability is locked as local and
deterministic: README is a product pamphlet, the hero uses local SVG generated
from the checked asset pipeline, important product copy remains Markdown text,
SVG assets avoid emoji, `foreignObject`, external fonts, external references,
and long text, README.ko does not exceed English canonical claims, remote
capsule-style renderers are inspiration only, and visual regression checks guard
the README and assets surface.
