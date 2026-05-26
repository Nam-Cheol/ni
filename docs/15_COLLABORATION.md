# Collaboration

`ni diff` and `ni conflicts` compare two planning states without resolving or mutating them.

These commands support parallel planning review inside the kernel boundary. They do not talk to Git hosts, create PRs, auto-resolve changes, or create execution queues.

## Inputs

Both commands accept a project directory, a `.ni/contract.json` file, or a `.ni/plan.lock.json` file:

```bash
ni diff --base <path-or-lock> --head <path-or-lock>
ni conflicts --base <path-or-lock> --head <path-or-lock>
```

Use `--json` for machine-readable output.

When a lockfile is available, the command verifies the lock hashes for that planning state. A mismatch is reported as a blocking collaboration conflict because the state has changed without a relock.

## Diff

`ni diff` reports contract-level additions, removals, and modifications across:

- non-goals,
- capabilities,
- requirements,
- decisions,
- risks,
- evaluations,
- artifacts,
- open questions.

`ni diff` is informational. It exits nonzero only when inputs cannot be loaded.

## Conflicts

`ni conflicts` reports semantic conflicts and exits nonzero when any blocking conflict is found.

Detected conflicts include:

- the same `CAP`, `REQ`, `DEC`, `RISK`, or `EVAL` ID changed between base and head,
- a new accepted decision contradicts an existing accepted decision,
- a capability was removed while its evaluation or artifact still exists,
- risk severity was lowered without a new mitigation or amendment reason,
- accepted requirement wording was weakened,
- a lock hash mismatch exists without relock.

The checks are deterministic and contract-local. They intentionally do not use model judgment.

## Review Flow

1. Run `ni diff --base <base> --head <head>` to inspect all contract changes.
2. Run `ni conflicts --base <base> --head <head>` before accepting parallel planning changes.
3. If conflicts are present, resolve them in planning docs and `.ni/contract.json`.
4. Run `ni status`.
5. Use `ni end` or `ni relock` only after readiness and amendment rules allow it.

After `.ni/plan.lock.json` exists, the lock remains the highest planning source of truth. Collaboration commands may point out stale or contradictory states, but they do not replace `ni status`, `ni end`, or lock verification.
