# Open questions

## OQ-001: exact Go package layout

Blocker: false

Status: resolved

Resolution: use `cmd/ni` for the CLI and `internal/core/*` for deterministic kernel packages.

## OQ-002: JSON schema library

Blocker: false

Status: resolved

Resolution: use minimal Go validation first; add schema tooling only if deterministic validation outgrows the current model.
