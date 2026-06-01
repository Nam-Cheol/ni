# Engineering skills applicability audit

This audit reviews selected public engineering skills from
<https://github.com/mattpocock/skills/tree/main/skills/engineering> for ideas
that may influence ni.

It does not copy, vendor, endorse, or claim compatibility with those skills.
It does not create new skills. It does not add runtime execution, task-runner
behavior, issue tracker publishing, PR automation, shell adapters, Codex exec,
or downstream agent execution.

## Boundary

Useful patterns may influence ni only when they preserve this repository's
architecture:

```text
ni-kernel
  docs contract
  readiness gate
  lockfile
  prompt compiler
  source-of-truth rule

ni-downstream-seeds
  project-specific work graph
  project-specific evaluation plan
  project-specific evidence rules
  project-specific adapter notes
```

The kernel owns deterministic planning contracts and gates. Downstream seed
material may be generated after lock, but it must remain derived and mutable.

## Applicability table

| Skill | Useful pattern for ni | Adapt now/later/not | Belongs in | Boundary risk | Recommended action |
| --- | --- | --- | --- | --- | --- |
| `diagnose` | Build a deterministic feedback loop before hypothesizing: reproduce, minimize, instrument, fix, and regression-test. For ni this maps to CLI status proof, fixture workspaces, golden tests, and lock-hash repros. | Later for validation/debug UX; now for internal development habits. | ni-kernel tests and docs; future skill pack guidance. | Medium | Use deterministic repro-loop language for `ni status`, docsync, lock mismatch, and prompt budget debugging. Do not add runtime instrumentation or production execution behavior. |
| `grill-with-docs` | Challenge fuzzy plans against existing domain language and documented decisions, asking one precise question at a time and updating docs as decisions crystallize. | Now | ni-start policy, readiness interview docs, model workspace packs. | Low | Adapt as a ni planning grill / docs-contract challenge pattern: preserve stable IDs, compare docs and contract, ask focused blocker questions, and update `docs/plan/**` plus `.ni/contract.json` together. |
| `improve-codebase-architecture` | Look for architectural friction using consistent domain vocabulary, locality, leverage, and testability. | Later | ni repo maintenance docs or separate contributor skill pack. | Medium | Use for ni maintainers when refactoring internal packages. Do not make it end-user ni-kernel behavior and do not generate visual reports as part of `ni status`, `ni end`, or `ni run`. |
| `prototype` | Treat prototypes as throwaway artifacts that answer one question and are deleted or absorbed after learning. | Later, cautiously | Separate downstream/sandboxed exploration seed, not ni-kernel. | High | Document as a downstream-only possibility after lock. Do not add prototype commands, runnable apps, persistence, or task-runner behavior to ni. |
| `setup-matt-pocock-skills` | Configure per-repo agent context before other skills operate: issue tracker, triage labels, and domain docs. | Later | Model-pack setup docs or separate package. | Medium | Use as inspiration for ni model-pack setup checks: locate project root, authority docs, language policy, and CLI path. Do not copy scaffolds or assume GitHub/Claude/Codex global installs. |
| `tdd` | Use behavior tests through public interfaces, one vertical slice at a time, with red-green-refactor discipline. | Now for ni internal development | ni-kernel test strategy and contributor docs. | Low | Prefer CLI-level and package-level behavior tests for readiness, docsync, lock, prompt budget, and skill-pack checks. Avoid speculative tests that encode imagined internals. |
| `to-issues` | Break a plan into independently grabbable vertical slices with dependencies and acceptance criteria. | Later | Downstream target seed or separate package. | Medium | May influence future issue-seed export after lock, but must not publish issues, mutate trackers, or become a task runner in core. |
| `to-prd` | Synthesize conversation and codebase context into a PRD with problem, solution, stories, decisions, testing, and out-of-scope sections. | Later | Downstream PRD/document seed, not ni-kernel. | Medium | Useful as a downstream export template after lock. Keep PRDs derived from the locked plan; do not publish to an issue tracker from ni core. |
| `triage` | Classify incoming work through a small state machine and separate need-info, ready-for-agent, ready-for-human, and won't-fix states. | Later | Feedback/pressure classification docs or separate model pack. | Medium | Could inform future pressure/feedback triage for planning conversations. Do not add issue comments, labels, closes, or tracker mutation to ni core. |
| `zoom-out` | Ask the agent to step up a layer and map relevant modules/callers using project vocabulary. | Now as a model skill-pack pattern | Model workspace packs and contributor docs. | Low | Use as a lightweight orientation prompt for ni maintainers and planning models. Keep it explanatory; no new kernel command required. |

## Expected ni adaptations

Immediate documentation and skill-pack influence:

- `grill-with-docs` becomes a ni planning grill pattern: challenge vague intent
  against docs and contract records, preserve terminology, ask one focused
  readiness question at a time, and update planning state visibly.
- `tdd` reinforces behavior-first tests through public CLI and package
  interfaces.
- `zoom-out` can shape model-pack orientation when a contributor or model needs
  the higher-level map before editing.

Later, outside core:

- `diagnose` can shape debugging guidance for deterministic repro loops around
  `ni status`, docsync, lock mismatches, and prompt budget failures.
- `to-prd` and `to-issues` may influence post-lock downstream seed formats, but
  only as derived artifacts. They must not publish issues from ni core.
- `triage` may influence future feedback/pressure classification.
- `setup-matt-pocock-skills` may inspire ni model-pack setup checks without
  copying external scaffolds.

Do not adapt into ni-kernel:

- `prototype` as runnable exploration, issue publishing, tracker mutation,
  shell execution, Codex execution, PR automation, queueing, or runtime agent
  orchestration.

## Non-goals

- No external skill files are copied into ni.
- No external skill package is vendored.
- No endorsement or compatibility claim is made.
- No new ni skills are created by this audit.
- No runtime execution, task runner, issue publishing, PR automation, shell
  adapter, Codex adapter, model API call, or Homebrew availability claim is
  added.
