# No-Terminal Planning

No-terminal planning is an assisted drafting path for people who want to start
using the Intent Lock Protocol before they install or run the `ni` CLI. It can
help a model conversation create initial `docs/plan/**` notes and a draft
`.ni/contract.json`, but it cannot finish a trusted lock on its own.

User-facing no-terminal questions should follow the language of the user's
latest substantive message. If the user speaks Korean, ask Korean planning
questions; if the user speaks English, ask English planning questions; if the
user requests a language, use it. Preserve IDs, commands, file paths, schema
keys, target names, and status constants exactly. A model may explain English
CLI proof in the user's language, but the proof remains authoritative only as
CLI output.

The rule is simple: no-terminal users can start planning, learn the protocol,
and prepare a draft for review. Trusted readiness, lock creation, hash
verification, and prompt compilation require CLI-validated `ni`.

## What No-Terminal Mode Is

No-terminal mode is:

- a way to start planning through model conversation;
- useful for drafting `docs/plan/**` and a contract draft;
- useful for learning `ni`'s Intent Lock Protocol before installation;
- useful for preparing questions, assumptions, non-goals, and handoff notes
  before a teammate or local setup runs the CLI;
- an assisted authoring path, not the source of authority.

## What No-Terminal Mode Is Not

No-terminal mode is not:

- deterministic readiness validation;
- a real `ni` lock;
- lock hash verification;
- a replacement for `ni status`, `ni end`, or `ni run`;
- execution;
- a hosted service;
- model authority.

Do not describe a no-terminal draft as `READY`, `READY_WITH_DEFERRALS`, or
locked unless a trusted CLI run produced that result and the exact proof is
available.

## What Users Can Do Without Terminal

Without direct terminal access, a user can:

- copy or load the `ni-start` guidance into a model workspace;
- describe the project idea, actors, delivery surface, constraints, and risks;
- ask the model to draft planning docs and `.ni/contract.json`;
- ask the model to show a draft planning proof block after meaningful updates;
- keep assumptions, draft decisions, explicit non-goals, and open blocker
  questions visible;
- hand the draft to a teammate, CI job, or later local setup for CLI
  validation.

Without CLI proof, a user cannot trust:

- readiness status;
- lock creation;
- lock hash checks;
- prompt compilation;
- downstream seed generation from a locked plan.

In no-terminal mode, planning proof capture is a draft audit trail only. It can
show what the model believes it changed, but it is not deterministic validation
and becomes trusted only after a CLI run validates the drafted docs and
contract.

## No-Terminal Checklist

Use this checklist when you are starting without a local CLI:

- Copy or load `ni-start` guidance from `.agents/skills/ni-start/SKILL.md`,
  `packages/codex-skills/ni-start/SKILL.md`, or
  `packages/claude-skills/ni-start/SKILL.md`.
- Describe the project idea in conversation: what should change, for whom, and
  how it will be delivered.
- Ask the model to create a `docs/plan/**` draft.
- Ask the model to draft `.ni/contract.json` alongside the docs.
- Mark uncertain statements as assumptions or open questions.
- Mark explicit exclusions as non-goals.
- Ask for a concise planning proof block after meaningful updates.
- Do not treat the draft as locked.
- Later validate with the CLI or a teammate who can run the CLI.

## Intent Lock Drafting Checklist

Use this checklist while drafting:

- Is the project purpose explicit?
- Are actors and outcomes named?
- Does every draft capability trace to at least one requirement and evaluation
  when those records exist?
- Are non-goals visible?
- Are high-severity risks paired with mitigations?
- Are open questions marked clearly, especially blockers?
- Are accepted decisions separated from assumptions and rejected options?
- Are expected artifacts named?
- Is the downstream handoff bounded to planning output, not runtime execution?

This checklist helps a model ask better questions. It does not replace
`ni status`, `ni end`, or `ni run`.

## Graduation Path To Full ni

Move from assisted drafting to full `ni` as soon as the plan might guide
implementation, budget, review, or downstream seed generation.

1. Install `ni` through the verified release binary path or the curl installer.
2. Run `ni status --proof --next-questions` in the drafted workspace.
3. Fix blockers by continuing the planning conversation and updating
   `docs/plan/**` and `.ni/contract.json` together.
4. Run `ni end` only after `ni status` reports readiness and the user confirms
   the accepted plan.
5. Run `ni run` only after `.ni/plan.lock.json` exists and the lock hashes are
   valid.
6. Keep model workspace skills as UX. They guide the conversation, but they do
   not override CLI readiness, locking, hash verification, or prompt
   compilation.

## Team Handoff Path

A no-terminal workflow can still be useful on a team:

1. A user without terminal access drafts the plan with a model.
2. A teammate with the CLI runs `ni status --proof --next-questions`.
3. The teammate returns the blockers, proof, and next questions exactly as CLI
   output or as a faithful summary with IDs.
4. The user continues planning with the model and keeps assumptions, non-goals,
   and open blockers visible.
5. The teammate validates again.
6. Locking happens only after CLI validation clears blockers and the user
   confirms the accepted plan.

If the teammate runs `ni end`, the generated `.ni/plan.lock.json` becomes the
source of truth. After that, do not silently edit locked planning docs.

## Model Workspace Pack Relation

Codex and Claude skill packs can guide the conversation. They can help a model
ask the first-run card, draft docs, draft contract records, preserve stable IDs,
and explain CLI blockers.

They do not execute downstream work. They do not call model APIs as part of
`ni`. They do not override CLI readiness. They do not create a real lock.
Global install may remain Experimental unless host-specific loading and
discovery have been verified.

## Example

See `examples/no-terminal-assisted/` for a docs-only example. It shows an
assisted draft with `docs/plan/00_project_brief.md` and `.ni/contract.json`,
keeps the draft marked "not locked", and requires later CLI validation.

The example intentionally does not run `ni status`, `ni end`, or `ni run`,
because that would imply a trusted CLI workspace instead of no-terminal
assisted drafting.

## Boundary

No-terminal planning must not add a hosted web app, model API calls, runtime
execution, shell adapters, queues, or automation behavior. It is a docs-first
way to start the Intent Lock Protocol while preserving the kernel boundary:

```text
model pack or copied guidance -> draft docs and contract
ni CLI -> deterministic readiness, lock, hash proof, prompt compile
```
