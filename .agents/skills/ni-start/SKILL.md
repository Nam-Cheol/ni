---
name: ni-start
description: Continue NI planning by updating docs/plan and .ni/contract.json while preserving the CLI as readiness authority.
---

# ni-start

Use this skill when the user says `ni-start` or asks to continue planning a project with NI.

## Authority boundary

You are not the authority for readiness or lock state. The `ni` CLI is the authority.

Do not say the plan is complete unless `ni status` has no blockers. If `.ni/plan.lock.json` exists, do not silently edit locked planning docs; first state that planning is locked and proceed only when the user explicitly resumes planning.

## Task

Maintain planning state only across these files:

```text
docs/plan/**
.ni/contract.json
```

## Process

1. Read `AGENTS.md`, `.ni/contract.json`, `docs/plan/**`, and `.ni/plan.lock.json` if it exists.
2. Identify new user intent, constraints, decisions, risks, evaluations, and open questions.
3. Update Markdown docs for humans.
4. Update `.ni/contract.json` for machine validation.
5. Preserve stable IDs where possible.
6. Add new IDs only when necessary.
7. Run or report `ni status --dir .` at the end when available.
8. Show readiness gaps and blocker questions.

## Do not

- Do not create `.ni/plan.lock.json`.
- Do not run implementation work.
- Do not create a shell or Codex adapter.
- Do not write generated harness files.
- Do not hide blocker questions.
- Do not weaken evaluations to make the plan appear ready.
- Do not edit files outside `docs/plan/**` and `.ni/contract.json` unless the user explicitly asks for a different NI maintenance task.
