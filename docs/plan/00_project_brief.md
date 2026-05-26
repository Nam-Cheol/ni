# Project brief

## Purpose

`ni` is a project intent compiler. It turns planning conversations into human planning docs, a machine-readable contract, a readiness result, a lockfile, and bounded prompt or seed material for downstream tools.

## Post-v0 direction

The v1 roadmap keeps the kernel narrow while making the planning contract useful across more project types and downstream targets. The kernel owns:

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

A user can update planning docs and `.ni/contract.json`, run `ni status`, lock or relock through the CLI, and compile a 4000-character-or-less target prompt from the valid lock.

## Boundary

Generated harnesses, seed packages, prompts, and downstream feedback are derived material. They may inform future amendments, but they do not replace `.ni/plan.lock.json`, `.ni/contract.json`, or `docs/plan/**` as source of truth.
