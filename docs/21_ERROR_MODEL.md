# Error Model

`ni` uses deterministic exit codes so humans, Codex, and downstream tools can automate kernel checks without parsing prose.

## Exit Codes

| Code | Name | Meaning |
| --- | --- | --- |
| 0 | OK | The command completed successfully. |
| 1 | generic failure | The command failed outside a more specific category. |
| 2 | usage error | The invocation is malformed: unknown command, unknown option, or missing/invalid argument. |
| 3 | readiness blocked | The readiness gate is blocked and the requested operation cannot proceed. |
| 4 | stale lock / hash mismatch | `.ni/plan.lock.json` exists but locked file hashes no longer match the current workspace. |
| 5 | invalid contract | `.ni/contract.json` is missing required structure, malformed, uses an unsupported schema, or violates deterministic contract validation. |
| 6 | unsupported target | The requested prompt or export target is not supported. |
| 7 | semantic conflict | Planning inputs are syntactically valid but conflict semantically. |

## Structured Errors

Commands that support `--json` emit a structured error envelope when they cannot emit their normal JSON payload:

```json
{
  "error": {
    "code": "invalid_contract",
    "exit_code": 5,
    "message": "status failed: malformed contract JSON: unexpected end of JSON input",
    "details": {
      "command": "status"
    }
  }
}
```

The `error.code` field is stable automation surface. The human-readable `message` may include command context and lower-level validation details.

Data-bearing validation commands may still emit their normal structured result with a non-zero exit code when the result itself is the evidence, such as `ni conflicts --json` returning a conflict report with exit code 7.

## Automation Rules

- Treat code 0 as success.
- Treat code 2 as caller error; retry only after changing the invocation.
- Treat code 3 as a readiness gate failure; run `ni status` and resolve blocker issues before locking.
- Treat code 4 as `BLOCKED`; do not use locked planning docs until the lock is reconciled.
- Treat code 5 as a contract repair task; do not weaken acceptance criteria to make validation pass.
- Treat code 6 as a target selection error; inspect `ni targets`.
- Treat code 7 as a planning conflict; report conflicting sources or IDs instead of choosing silently.
