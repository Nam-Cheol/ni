---
name: Bug report
about: Report a problem with ni-kernel behavior, validation, locking, or prompt compilation
title: "Bug: "
labels: bug
assignees: ""
---

## ni Version

<!-- Paste `ni version`, `go run ./cmd/ni version`, or the commit SHA. -->

## Command

<!-- Include the exact command and flags. -->

```bash

```

## Expected Result

<!-- What should have happened? -->

## Actual Result

<!-- What happened instead? Include output or error text when useful. -->

## Workspace Shape

<!-- Describe the relevant workspace state. Share only non-sensitive details. -->

- `docs/plan/**`:
- `.ni/contract.json`:
- `.ni/plan.lock.json`:
- Was the plan locked?
- Were any files edited after locking?

## Boundary Check

<!-- ni-kernel validates, locks, and compiles intent. It does not execute downstream work. -->

- Does this bug affect readiness validation, docs/contract sync, lock/hash checks, prompt compilation, or inert seed generation?
- Does the report involve a downstream runtime, task runner, queue, Codex exec, PR automation, or release automation?
