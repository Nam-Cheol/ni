# AGENTS.md

## Critical instruction

Do not implement `ni` as a task runner first.

The initial product is `ni-kernel`: a project intent compiler that creates, validates, locks, and compiles planning contracts before any execution harness runs.

## Product architecture

Use this boundary throughout the repository:

```text
ni-kernel
  docs contract
  readiness gate
  lockfile
  prompt compiler
  source-of-truth rule

ni-generated-harness
  project-specific work graph
  project-specific evaluation plan
  project-specific evidence rules
  project-specific adapters
```

The kernel is authoritative. Generated harnesses are derived and mutable.

## Authority rules

1. Skills are UX; the CLI is authority.
2. A model may draft docs, detect gaps, propose work graphs, and propose a harness.
3. A model may not declare readiness without `ni status`.
4. A model may not lock a plan without `ni end`.
5. A model may not weaken acceptance criteria to pass validation.
6. A model may not silently edit locked planning docs.
7. If a lock hash mismatch exists, stop and report `BLOCKED`.

## Source-of-truth precedence

After `.ni/plan.lock.json` exists, use this order:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > chat transcript > model guess
```

## Initial command roadmap

Implement in this order:

```text
ni --help
ni version
ni init
ni status
ni end
ni run
```

`ni run` must initially compile a prompt only. Do not make it execute Codex or shell commands.

## Development rules

- Keep changes small.
- One prompt should map to one coherent commit.
- Prefer deterministic validation over model judgment.
- Every capability should map to at least one evaluation.
- High-severity risks require mitigation.
- Open blocker questions must prevent locking.
- Prompt output from `ni run` must be 4000 characters or less.

## Validation expectations

When Go code exists, run:

```bash
gofmt -w .
go test ./...
bash scripts/quality.sh
```

Before Go code exists, run:

```bash
bash scripts/quality.sh
```

## Forbidden early work

Do not add these before `ni status`, `ni end`, and `ni run` work:

- shell adapter,
- Codex adapter,
- evidence runner,
- queue,
- PR automation,
- release automation,
- plugin system,
- TUI or web UI.
