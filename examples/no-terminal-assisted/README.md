# No-Terminal Assisted Draft

This example shows how a model workspace can start an `ni`-shaped plan before
the CLI is installed. It is intentionally a draft. There is no lockfile, no hash
proof, and no compiled handoff prompt.

Expected `ni status`: not claimed until a teammate, CI job, or local CLI setup
runs the command.

## What the model may draft

- `docs/plan/00_project_brief.md`: human-readable planning notes.
- `.ni/contract.json`: a draft contract aligned with the docs.
- Assumptions and blocker questions that stay visible instead of being treated
  as accepted decisions.

## No-terminal checklist

- Start with a model pack or copied instructions.
- Create a `docs/plan` draft.
- Draft `.ni/contract.json` alongside the docs.
- Mark assumptions and open questions, especially blockers.
- Later validate with the CLI, a teammate, or a trusted runner.
- Do not treat model judgment as a lock.

## Graduate before handoff

Use full `ni` before this draft guides implementation or downstream seed
generation. The CLI must produce readiness, lock creation, hash verification,
and prompt compilation with `ni status`, `ni end`, and `ni run`.

This example does not add a web service, model API call, runtime execution,
shell adapter, queue, or model-authoritative skill.
