# Init TUI UX Repair

## Current status

State:
- v0.6.0 release: published and verified
- primary command: namba-intent
- init TUI: Bubble Tea v2 / Lip Gloss v2
- reported UX issues: unclear help text, arrow-key editing conflict, unclear next steps
- Homebrew: Planned / v0.5 candidate
- Windows real-host verification: pending
- Skills are UX; CLI is authority.
- Namba Intent is a pre-runtime Project Intent Compiler for AI Agents.

## Repair goal

This repair addresses three first-user problems reported during
`namba-intent init .` dogfooding:

- the setup guide did not explain why each question mattered or how rough an
  answer could be;
- answer fields used arrow keys for wizard movement, which conflicted with
  ordinary text editing;
- the post-init text did not clearly say what was created, what state the user
  was in, how to continue with status/end/run, or how optional model help fits
  under CLI authority.

The repair keeps init as planning setup only. It does not install skills, lock
the plan, run agents, execute prompts, add adapters, or create runtime
execution state.

## Question/help copy

| Question / area | Old problem | New guidance | Notes |
| --- | --- | --- | --- |
| Language choice | User could choose a language but did not get readiness context yet. | The guide keeps this short and says labels/review guidance use the selected language. | Menus still use arrows because they are choice lists. |
| Korean guide tone | The first repair still sounded like translated internal documentation. | Korean copy now uses plain conversational labels and fewer internal terms. | Based on the `humanize-korean` fast-path rewrite rules. |
| Korean guide layout | The question guide, example, uncertainty fallback, and READY note looked like one text blob. | The Korean guide now separates `쓸 내용`, `예시`, `모르면`, `다음`, and `READY` with colored short labels and per-item lines. | The TUI copy stays short enough to keep the answer box visible at 80x24. |
| Project name | The field looked like a simple label request. | Asks for a short name the user can recognize later. | Folder-name default remains acceptable. |
| Project goal | User could think a perfect final answer was required. | Asks who uses it, what should improve, and says one or two detailed sentences are okay. | Uncertainty can be written as "still unknown". |
| Target users / audience | Actor expectations were implicit and too technical. | Asks who will see, use, or review the result. | Rough user lists are acceptable during init. |
| Downstream agent task | Could sound like init would run the agent. | Says this is work for after the plan is locked, not something init runs now. | Preserves non-execution boundary. |
| Constraints / non-goals | User might omit hard boundaries. | Asks what must not happen in this round. | Default still blocks downstream work before lock. |
| Success criteria | User might not know the answer shape. | Asks what the user would check later to say the plan is good enough. | Unknown criteria stay visible for later status. |
| Known blockers / open questions | User might hide uncertainty to get through the wizard. | Asks what is still unknown or needs a decision. | Uncertainty remains visible before lock. |
| Deferrals | Deferred scope could be confused with accepted work. | Asks what matters but should wait until later. | Default remains acceptable when nothing is deferred. |

Every answer-stage guide now separates:

- What to write
- Example
- If you are not sure
- a short note that `status` shows what is missing
- READY means the plan is ready to lock, not that the product is finished

## Keyboard behavior

| Context | Key | Expected behavior | Implemented behavior | Notes |
| --- | --- | --- | --- | --- |
| Answer input | Left / Right | Move within the answer text. | Moves the answer cursor left/right. | Does not change wizard step. |
| Answer input | Up / Down | Move within multiline text, or stay predictable for single-line text. | Moves between pasted multiline rows when possible; otherwise no step change. | Prevents accidental step jumps. |
| Answer input | Backspace / Delete | Edit text normally. | Backspace removes before cursor; Delete removes at cursor. | Cursor is preserved. |
| Answer input | Enter | Continue when appropriate. | Moves to the next field, or to confirmation from the last field. | Does not insert a newline. |
| Answer input | Tab / Shift+Tab | Next / previous field. | Tab advances; Shift+Tab goes back. | Details drawer uses Ctrl+D while editing. |
| Answer input | Ctrl+Right / Ctrl+Left | Next / previous step. | Moves to next/previous field. | Separate from arrow editing. |
| Answer input | Esc | Go back or cancel from first answer. | Previous field, or cancel if already at the first field. | Preserves existing cancel semantics. |
| Answer input | q | Type text. | Plain `q` remains editable text in answer fields. | Ctrl+Q still quits; this avoids stealing a normal letter. |
| Menus | Up / Down / Left / Right | Navigate choices. | Language, existing-file, and confirmation choices still use arrows. | Context-aware arrows. |
| Menus | q | Quit/cancel. | Language, existing-file, and confirmation stages accept `q`. | Answer fields keep `q` as text. |
| Confirmation screen | Enter | Confirm selected choice. | Writes initial artifacts only when selected confirmation is accepted. | Still does not lock. |
| Cancel/quit | Ctrl+C / Ctrl+Q | Quit safely. | Cancels and exits without writing pending TUI answers. | Existing global controls preserved. |

## Post-init summary

After init exits, successful init prints a plain text summary outside AltScreen.
The structure answers:

- What was created?
- What was skipped or unchanged?
- What state am I in now?
- What do I do next?
- How can I use an AI assistant?
- What does Namba Intent not do?

Next-step wording now includes:

- Run `namba-intent status --proof --next-questions`
- If status is BLOCKED, answer the listed questions or refine docs/plan/**
- When status is READY and you agree with the plan, run `namba-intent end`
- Then run `namba-intent run --max-chars 4000` to compile a bounded handoff prompt
- `namba-intent run` does not execute the prompt or run an agent

## Model/skills guidance

Init does not create or install model skills. Skills and model assistance are
optional UX layers.

If a user wants model help, the summary tells them to ask the assistant to read
`docs/plan/**`, `.ni/contract.json`, and `.ni/session.json`, then help answer
next questions and update docs plus contract together.

The authority boundary remains:

```text
Skills are UX; CLI is authority.
```

Models may draft and explain. They may not claim readiness, lock the project,
or replace `namba-intent status`, `namba-intent end`, and `namba-intent run`.

## Tests added

- `TestAnswerFieldArrowsEditTextInsteadOfChangingSteps`
- `TestAnswerFieldStepNavigationUsesTabEnterAndCtrlArrows`
- `TestMenuQQuitsWithoutStealingAnswerText`
- CLI init summary assertions for next steps, READY boundary, run boundary, and CLI authority guidance
- Updated responsive render snapshots for visible `setup guide` / `guide` copy and new help bars

## Claim-boundary audit

| Claim area | Expected boundary | Observed state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| init TUI | Guided planning setup only. | TUI asks clearer questions and writes draft init answers only after confirmation. | Yes | No lock or execution behavior added. |
| model assistance | Optional draft/help layer. | Post-init summary explains how to ask an assistant for help. | Yes | CLI remains authority. |
| skills | Skills are UX only. | Init explicitly does not create or install skills. | Yes | No skill installation path added. |
| status | CLI readiness gate. | Summary points to `namba-intent status --proof --next-questions`. | Yes | Model cannot claim readiness. |
| end | Lock only after READY and user agreement. | Summary says to run `end` only when status is READY and the user agrees. | Yes | Project root `end` was not run by this repair. |
| run | Bounded handoff prompt compilation only. | Summary says `run` does not execute the prompt or run an agent. | Yes | No generated prompt executed. |
| READY | Planning readiness, not product readiness. | TUI and summary state this boundary. | Yes | Prevents product-readiness overclaim. |
| Homebrew | Planned / v0.5 candidate. | Preserved. | Yes | No Available claim added. |
| Windows real-host | Pending until transcript exists. | Preserved. | Yes | No Windows verification claim added. |
| runtime execution | No task runner, SPEC runner, harness, adapter, queue, PR/release automation, or downstream execution layer. | Preserved in summary and implementation. | Yes | No runtime behavior added. |
