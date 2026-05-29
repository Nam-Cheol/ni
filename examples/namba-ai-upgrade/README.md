# namba-ai Upgrade Planning

## 1. Purpose

This example uses ni as a Project Intent Compiler for the next namba-ai
upgrade before any runtime, SPEC-runner, Codex execution harness, or namba-ai
implementation work begins.

It exists to prove that ni keeps the kernel authoritative even for
software-oriented downstream seed material. The current fixture is verified
through `ni status`, which reports that a planning-doc body is missing and the
example is not freshly ready.

## 2. What this proves

- A software-oriented plan can still be pre-runtime and contract-first.
- `ni status` remains the readiness authority and catches stale planning docs.
- Grouped `ni status --proof --next-questions` output identifies the sync
  repair instead of letting a model restart broad planning.
- The historical `.ni/plan.lock.json` can still bound existing downstream
  prompt seed material.
- `ni run --target codex` produces seed material without invoking Codex when
  used against the existing lock.
- Downstream target notes stay derived and mutable; they do not become
  kernel-owned execution state.

## 3. Product type / surface

- `product_type`: `software`
- `delivery_surface`: `cli`, `document`, `workflow`
- Expected `ni status`: `BLOCKED`
- Expected `ni run` target: `codex` from the existing lock only

## 4. Files

- `docs/plan/**`: planning docs for the namba-ai upgrade intent.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, and decisions.
- `.ni/plan.lock.json`: historical CLI-written lock with hashes for
  authoritative planning files.
- `generated/codex.goal.prompt.txt`: checked-in compiled Codex target prompt
  from the existing lock.

Some duplicate `* 2.*` files are present as historical fixture material. They
are not the authoritative locked planning files.

## 5. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir examples/namba-ai-upgrade
go run ./cmd/ni status --dir examples/namba-ai-upgrade --proof --next-questions
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/namba-ai-upgrade --target codex --max-chars 4000 --out "$tmpdir/codex.goal.prompt.txt"
wc -m "$tmpdir/codex.goal.prompt.txt"
rm -rf "$tmpdir"
```

## 6. Expected output

Expected status: `BLOCKED`.

The status command should start with:

```text
BLOCKED
profile: prototype
product type: software
delivery surfaces: cli, document, workflow
```

It should also report the current fixture blocker:

```text
blocker R012: CAP-001
```

The run command may compile from the existing lockfile, but this example must
not be described as freshly ready until `ni status` passes again.

## 7. demo-check coverage

Covered by `bash scripts/demo-check.sh`.

The demo check verifies the current `BLOCKED` status and `R012` sync blocker,
then compiles the Codex prompt only from the existing lock.

## 8. Korean companion

Korean companion docs exist: `README.ko.md`.

## 9. Non-execution boundary

This example does not run Codex, modify namba-ai, create a shell adapter, start
a queue, call a model API, run a SPEC workflow, or execute downstream tools. Do
not call `ni end` for this fixture while `ni status` reports `BLOCKED`; ni only
validates planning state and compiles inert prompt seed material from an
existing lock.
