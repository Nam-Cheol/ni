# Project brief

## Purpose

`ni` is a project intent compiler. In v0.2, it turns model-user planning conversations into synchronized human planning docs, `.ni/contract.json`, a deterministic readiness result, a lockfile, and a bounded handoff prompt for downstream tools.

## v0.2 focus

`ni init` creates the initial planning structure. After initialization, the primary authoring interface is sustained model-user conversation through `ni-start`, which updates `docs/plan/**` and `.ni/contract.json` together. The CLI remains authoritative for deterministic gaps (`ni status`), explicit lock or relock (`ni end` or `ni relock`), and prompt compilation (`ni run`).

User-facing contract `add`, `list`, or `set` commands are not part of the v0.2 primary authoring UX.

## Later direction

The later roadmap keeps the kernel narrow while making the planning contract useful across more project types and downstream targets. The kernel owns:

- related-work boundaries,
- readiness profiles,
- product type and delivery-surface guidance,
- downstream target registry,
- locked target exports,
- feedback ingest,
- pressure ledger,
- harness candidate lifecycle,
- amendment and relock flow,
- collaboration diff and conflict checks.

## Problem

Agent and SPEC systems often mix planning, execution, evidence collection, and project growth into one runtime. `ni` separates those concerns. Planning can become explicit and locked before any generated harness or downstream runtime begins work.

## Success definition

A user can start with `ni init`, plan through model-user conversation, let `ni-start` update planning docs and `.ni/contract.json`, run `ni status`, lock or relock through the CLI, and compile a 4000-character-or-less target prompt from the valid lock.

## Boundary

Generated harnesses, seed packages, prompts, and downstream feedback are derived material. They may inform future amendments, but they do not replace `.ni/plan.lock.json`, `.ni/contract.json`, or `docs/plan/**` as source of truth.
