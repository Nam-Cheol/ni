# No-Terminal Assisted Draft

## 1. Purpose

This example shows how a model workspace can start an `ni`-shaped Project
Intent Compiler plan before the CLI is installed. It is intentionally a draft.
There is no lockfile, no hash proof, and no compiled handoff prompt.

## 2. What this proves

- Assisted planning can capture docs and a draft contract before local CLI
  access exists.
- Model judgment is not readiness, lock, or hash authority.
- The draft must graduate to full CLI validation before handoff or downstream
  seed generation.
- This example does not make deterministic validation claims.

## 3. Product type / surface

- `product_type`: draft `workflow`
- `delivery_surface`: draft `document`
- Expected `ni status`: not claimed until a teammate, CI job, or local CLI
  setup runs the command.
- Expected `ni run` target: not applicable.

## 4. Files

- `docs/plan/00_project_brief.md`: human-readable planning notes.
- `.ni/contract.json`: a draft contract aligned with the docs.
- Assumptions and blocker questions that stay visible instead of being treated
  as accepted decisions.

## 5. Commands

This example is intentionally docs-only. From the repository root:

```bash
test -f examples/no-terminal-assisted/README.md
test -f examples/no-terminal-assisted/README.ko.md
test -f examples/no-terminal-assisted/docs/plan/00_project_brief.md
test -f examples/no-terminal-assisted/.ni/contract.json
```

## 6. Expected output

The `test` commands should exit successfully. Do not describe the draft as
`READY`, `READY_WITH_DEFERRALS`, or `BLOCKED` unless a trusted CLI run is
available and quoted separately.

## 7. demo-check coverage

Covered by `bash scripts/demo-check.sh` as a docs-only example.

The demo check verifies required files and boundary wording. It does not run
`ni status`, `ni end`, or `ni run` for this intentionally assisted draft.

## 8. Korean companion

Korean companion docs exist: `README.ko.md`.

## 9. No-terminal checklist

- Start with a model pack or copied instructions.
- Create a `docs/plan` draft.
- Draft `.ni/contract.json` alongside the docs.
- Mark assumptions and open questions, especially blockers.
- Later validate with the CLI, a teammate, or a trusted runner.
- Do not treat model judgment as a lock.

## 10. Graduate before handoff

Use full `ni` before this draft guides implementation or downstream seed
generation. The CLI must produce readiness, lock creation, hash verification,
and prompt compilation with `ni status`, `ni end`, and `ni run`.

## 11. Non-execution boundary

This example does not add a web service, model API call, runtime execution,
shell adapter, queue, or model-authoritative skill.
