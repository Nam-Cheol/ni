# 029 - Repository consistency audit after v0.1 RC

General rules for this task:
- Read AGENTS.md first.
- Keep the `ni-kernel` boundary intact.
- Do not add runtime execution, Codex adapters, shell adapters, queues, evidence
  runners, PR automation, release automation, plugin systems, or UI work.
- Before editing, summarize the intended changes and exact files you expect to
  modify.
- After editing, show changed files and validation results.

Task:
Audit repository documentation after the completed v0.1 release-candidate
state.

Goal:
Make README, roadmap, prompt archive, release notes, and docs agree that v0.1
RC is a pre-runtime project intent compiler.

Scope:
- Update `docs/08_ROADMAP.md` to show the v0.1 RC as complete and describe the
  next v0.1.0 and v0.2 phases.
- Update `prompts/README.md` so it does not imply the prompt sequence ends at
  `012`.
- Add this prompt as `prompts/029-repo-consistency-audit.md`.
- Add a short `docs/20_NEXT_STEPS.md` summarizing the next phase.
- Confirm README, release notes, roadmap, and AGENTS.md use the same product
  boundary.

Non-goals:
- Do not add runtime execution.
- Do not add Codex or shell adapters.
- Do not add queues, evidence runners, PR automation, release automation,
  plugin systems, TUI, or web UI.

Validation:
- `bash scripts/quality.sh`
- Search for stale roadmap references to `SPEC-011`, `SPEC-012`, and
  `SPEC-013` execution experiments. Update them or mark them historical.
