# Release Notes: v0.1 Release Candidate

`ni` v0.1 is the first release-candidate shape of the Project Intent Compiler
for AI Agents.
It is usable by someone who understands the pre-runtime contract compiler model:
plan first, validate deterministically, lock the accepted plan, then derive
prompts or seed material.

## Included

- CLI bootstrap with `--help` and `version`.
- `ni init` for planning docs, `.ni/contract.json`, readiness profiles, pressure
  ledger, and harness candidate skeletons.
- `ni status` deterministic readiness evaluation.
- `ni end` lockfile creation for ready plans.
- `ni run` prompt compilation with `--target`, `--out`, and `--max-chars`.
- Built-in targets: `generic`, `codex`, `human-team`, `hyper-run`, `namba-ai`,
  `ouroboros`, and `spec-kit`.
- `ni export` seed packages for `hyper-run`, `namba-ai`, `ouroboros`, and
  `spec-kit`.
- Product-shape fields for non-software planning: `product_type`,
  `delivery_surfaces`, and `interaction_mode`.
- Readiness profiles: `concept`, `prototype`, `mvp`, `beta`, and `production`.
- Inert feedback ingest and pressure tracking.
- Amendment and relock flow for explicit post-lock planning changes.
- Collaboration diff and conflict checks.
- Read-only work graph and downstream harness seed proposal support.

## Kernel Boundary

The v0.1 kernel owns planning authority only:

```text
docs contract
readiness gate
lockfile
prompt compiler
source-of-truth rule
```

Harness seed proposals and exports are downstream seed material:

```text
project-specific work graph
project-specific evaluation plan
project-specific evidence rules
project-specific adapter notes
```

The kernel does not own runtime execution state.

## Not Included

v0.1 does not include:

- shell adapters,
- Codex adapters,
- evidence runners,
- queues,
- PR automation,
- release automation,
- package publishing,
- plugin systems,
- TUI or web UI.

`ni run` compiles a prompt only. It does not execute that prompt.

## Upgrade and Use Notes

- Existing planning work should be represented in `docs/plan/**` and
  `.ni/contract.json`.
- Run `ni status` before trusting readiness.
- Run `ni end` only when readiness is `READY` or `READY_WITH_DEFERRALS`.
- Treat `.ni/plan.lock.json` as authoritative after lock.
- If locked docs or the contract change, use an explicit amendment and `ni
  relock`.
- Use `ni export` only for derived seed packages after lock.

## Validation

The release candidate is expected to pass:

```bash
go test ./...
bash scripts/quality.sh
```

The README quickstart uses `go run ./cmd/ni ...`, which matches the current
unpackaged repository state.
