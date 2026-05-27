# namba-ai Upgrade Planning

## 1. Purpose

This locked example plans the next namba-ai upgrade before any runtime,
SPEC-runner, Codex execution harness, or namba-ai implementation work begins.

It exists to prove that ni can compile a downstream Codex prompt from a locked
project contract while keeping the kernel authoritative.

## 2. What this proves

- A software-oriented plan can still be pre-runtime and contract-first.
- `.ni/plan.lock.json` remains the source of truth for downstream prompt use.
- `ni run --target codex` produces seed material without invoking Codex.
- Downstream target notes stay derived and mutable; they do not become
  kernel-owned execution state.

## 3. Product type / surface

- `product_type`: `software`
- `delivery_surface`: `cli`, `document`, `workflow`
- Expected `ni status`: `READY`
- Expected `ni run` target: `codex`

## 4. Files

- `docs/plan/**`: locked planning docs for the namba-ai upgrade intent.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, and decisions.
- `.ni/plan.lock.json`: CLI-written lock with hashes for authoritative
  planning files.
- `generated/codex.goal.prompt.txt`: checked-in compiled Codex target prompt.

Some duplicate `* 2.*` files are present as historical fixture material. They
are not the authoritative locked planning files.

## 5. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir examples/namba-ai-upgrade
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/namba-ai-upgrade --target codex --max-chars 4000 --out "$tmpdir/codex.goal.prompt.txt"
wc -m "$tmpdir/codex.goal.prompt.txt"
rm -rf "$tmpdir"
```

## 6. Expected output

Expected status: `READY`.

The status command should start with:

```text
READY
profile: prototype
product type: software
delivery surfaces: cli, document, workflow
```

The run command should write a non-empty prompt at or below 4000 characters.

## 7. Non-execution boundary

This example does not run Codex, modify namba-ai, create a shell adapter, start
a queue, call a model API, run a SPEC workflow, or execute downstream tools. ni
only validates the locked contract and compiles an inert Codex target prompt.
