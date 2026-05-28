# Conversation Authoring UX Audit

Task 139 audits the v0.4 conversation-driven authoring experience. It is an
audit note only. It does not change CLI behavior, lock state, model packs, or
downstream execution behavior.

## Scope Reviewed

- `README.md` and `README.ko.md`
- `AGENTS.md`
- `docs/28_CONVERSATION_AUTHORING.md`
- `docs/29_AUTHORING_PROTOCOL.md`
- `docs/30_DOC_UPDATE_RULES.md`
- `docs/31_NI_START_BEHAVIOR.md`
- `docs/34_READINESS_INTERVIEW.md`
- `docs/35_NI_END_CONFIRMATION.md`
- `docs/36_NI_RUN_HANDOFF.md`
- `docs/37_MODEL_EDIT_SAFETY.md`
- `docs/no-terminal.md`
- `docs/55_MODEL_WORKSPACE_PACKS.md`
- `packages/claude-skills/**`
- `packages/codex-skills/**`
- `examples/ni-start-dogfood/**`
- `examples/conversation-authoring/**`

## Overall Finding

The kernel boundary is strong. The docs and skills consistently preserve the
rule that conversation is the authoring surface while `ni status`, `ni end`,
and `ni run` remain the authority.

The main v0.4 adoption risk is not conceptual drift. It is first-run
orientation: a new user can learn the protocol by reading several files, but
the immediate "what do I say after `ni init`?" path is still scattered across
README, protocol docs, skills, and examples.

## Top 10 UX Gaps

| Rank | Gap | User impact | Type |
| --- | --- | --- | --- |
| 1 | No single post-`ni init` starter script tells a new user what to say next. README says to use conversation, and examples show it, but the first usable prompt is not surfaced as a compact copyable path. | High: users may fall back to manual JSON editing or ask a model to implement too early. | Docs-only now; optional CLI hint later |
| 2 | The decision point between "continue planning" and "run `ni-end`" is clear in protocol docs, but not visible enough in first-run materials. `READY_WITH_DEFERRALS` especially needs a plain-language explanation at the moment users see it. | High: users can mistake "close enough" for lock readiness, or treat deferrals as hidden cleanup. | Docs-only |
| 3 | Readiness failures have `next_questions`, but the user-facing loop is not packaged as a reusable interview card. The examples show good questions, yet README and model-pack READMEs do not give users a small pattern for answering blocker questions. | High: blocked status can feel like a dead end instead of the next planning turn. | Docs-only |
| 4 | Docs/contract sync expectations are explicit, but diagnostics for drift are not yet framed as a first-class user experience. A user can learn that docs and `.ni/contract.json` must move together, but does not get a clear "these records disagree" proof surface from the docs reviewed. | High: sync drift is the most likely failure mode for model-authored planning state. | Requires CLI changes for best result |
| 5 | `ni-start` resume behavior is rigorous for models, but the user does not get a concise resume proof template. The docs say session claims must be verified, while examples show resume; the UX could show a predictable "I resumed from X, verified Y, conflicts Z" shape. | Medium-high: users may not trust resumed planning or may overtrust hidden chat memory. | Docs-only |
| 6 | Assumptions, decisions, risks, non-goals, and open questions are handled distinctly in protocol docs, but there is no short authoring cheat sheet for users and pack users. The distinctions are correct but require deeper reading. | Medium-high: ambiguous statements can be promoted too quickly by model hosts that only read the pack README. | Docs-only |
| 7 | Model workspace packs explain CLI authority clearly, but their setup flow is more installation-oriented than task-oriented. After installation, a user still needs a compact "open project, invoke `ni-start`, paste CLI proof if needed" workflow. | Medium: pack users may install correctly but not know the first working loop. | Docs-only |
| 8 | No-terminal mode avoids overstating deterministic validation, but it lacks a proof-capture template. Users are told to ask a trusted runner for `ni status`, `ni end`, and `ni run`, but not exactly what output to paste back and how the model should treat it. | Medium: no-terminal users may paste summaries instead of authoritative CLI output. | Docs-only |
| 9 | Korean companion coverage exists for public release and distribution docs, but the central conversation-authoring protocol docs are mostly English-only. `README.ko.md` points to the flow, yet Korean-first users lose depth after the entry page. | Medium: Korean adoption path is less complete for the v0.4 focus area. | Docs-only |
| 10 | Codex and Claude `ni-start` skills are aligned on authority, but the Codex skill contains a fuller resume and output contract than the Claude skill. That difference may create uneven behavior across model workspaces. | Medium: cross-host examples may not transfer cleanly. | Docs-only |

## Audit Question Answers

| Question | Answer |
| --- | --- |
| Can a new user understand what to say after `ni init`? | Partly. README and examples imply the path, but there is no compact first prompt or post-init conversation card. |
| Does `ni-start` clearly resume previous planning state? | Yes for model instructions, especially Codex and `docs/31_NI_START_BEHAVIOR.md`; less clearly as a user-visible resume proof. |
| Does the model know how to update docs and contract together? | Yes. `docs/29_AUTHORING_PROTOCOL.md`, `docs/30_DOC_UPDATE_RULES.md`, `docs/37_MODEL_EDIT_SAFETY.md`, and the skills repeat this strongly. |
| Does the user know when to continue planning vs run `ni-end`? | Partly. The rule exists, but first-run docs should make `BLOCKED`, `READY_WITH_DEFERRALS`, and `READY` transitions more obvious. |
| Are readiness failures translated into useful next questions? | Yes at the protocol level through `--next-questions`; the adoption gap is packaging that loop for users. |
| Are assumptions, decisions, risks, non-goals, and open questions handled distinctly? | Yes in deep docs; the distinction needs a compact user and pack quick reference. |
| Do model workspace packs explain CLI authority clearly? | Yes. Both packs consistently say skills are UX and CLI is authority. The missing piece is a task-first usage loop. |
| Does no-terminal mode avoid overstating deterministic validation? | Yes. It explicitly says deterministic readiness, locking, hash verification, and prompt compilation require the CLI. The gap is exact proof capture. |

## Docs-Only Gaps

These gaps can be addressed without CLI behavior changes:

- Gap 1 as a README/docs/model-pack quickstart, unless the project also wants
  `ni init` to print a next-step hint later.
- Gap 2 by adding a readiness transition card.
- Gap 3 by adding a reusable blocker interview card.
- Gap 5 by adding a user-visible resume proof template.
- Gap 6 by adding a record classification cheat sheet.
- Gap 7 by adding task-first pack usage instructions.
- Gap 8 by adding no-terminal proof-paste templates.
- Gap 9 by adding Korean companions for the v0.4 authoring docs.
- Gap 10 by aligning Claude and Codex skill wording.

## CLI-Change Gaps

These gaps would benefit from CLI work in a later task:

- Gap 1 if `ni init` should print a short "next conversation prompt" after
  workspace creation.
- Gap 4 if `ni status` should emit explicit docs/contract sync diagnostics,
  record-level mismatch categories, or JSON proof fields that model packs can
  preserve without reinterpretation.

No CLI work is recommended inside this audit task.

## Recommended Next 3 Tasks

1. Add a first-run conversation authoring card.

   Update README, model-pack READMEs, and `docs/no-terminal.md` with a compact
   path:

   ```text
   After ni init, say:
   "Invoke ni-start. Help me plan <project>. Ask only the next questions needed
   for ni status, and keep docs/plan/** plus .ni/contract.json synchronized."
   ```

   Include the transition rule: `BLOCKED` means continue planning,
   `READY_WITH_DEFERRALS` means review visible deferrals before `ni-end`, and
   `READY` means `ni-end` may begin confirmation.

2. Add readiness proof and blocker-question examples to the adoption path.

   Create a small reusable card for `ni status --proof --next-questions` that
   shows how a model should preserve CLI status, explain blockers, ask one to
   three next questions, and avoid treating model judgment as readiness.

3. Design docs/contract sync diagnostics for a later CLI task.

   Draft the expected `ni status` issue shape for docs/contract mismatch:
   affected record ID, doc path, contract field, severity, and next question.
   Keep the first step as a spec or doc proposal; implement CLI behavior only
   in a separate task.

## Suggested v0.4 Success Test

A fresh user should be able to:

1. Run `ni init`.
2. Copy one conversation prompt into a model workspace.
3. Answer focused blocker questions.
4. See exactly which files the model changed.
5. Understand why `ni status` is `BLOCKED`, `READY_WITH_DEFERRALS`, or `READY`.
6. Confirm `ni-end` only after a visible summary.
7. Run `ni-run` knowing it compiles a bounded prompt only.

If that path works without manual `.ni/contract.json` editing and without any
downstream execution claims, v0.4 conversation authoring hardening is pointed
at the right problem.
