# Lockfile

`ni end` creates `.ni/plan.lock.json`.

The lockfile records a hash snapshot of the accepted planning contract and docs. After this point, `ni run` must verify that current files still match the lock.

## Required lockfile contents

```json
{
  "schema": "ni.lock.v0",
  "locked_at": "2026-05-26T00:00:00Z",
  "status": "LOCKED",
  "contract_hash": "sha256:...",
  "docs": [
    {"path": ".ni/contract.json", "sha256": "sha256:..."},
    {"path": "docs/plan/00_project_brief.md", "sha256": "sha256:..."}
  ],
  "source_of_truth": [
    ".ni/plan.lock.json",
    ".ni/contract.json",
    "docs/plan/**",
    ".ni/session.json",
    "chat history"
  ]
}
```

## Lock behavior

```text
ni end
  runs ni status
  refuses BLOCKED
  writes plan.lock.json if ready

ni run
  reads plan.lock.json
  verifies hashes
  refuses stale docs
```

The lock hashes `.ni/contract.json` and required `docs/plan/**` files. It does
not hash `.ni/session.json`, because session state is mutable carryover context
below locked docs in the source-of-truth order.
