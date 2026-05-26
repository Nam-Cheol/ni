# Start here

This repository starts from a strict premise: `ni` should not begin as a generic agent harness.

The first implementation target is a planning kernel:

```text
conversation -> docs contract -> readiness gate -> lockfile -> 4000-char goal prompt
```

Generated harnesses and downstream execution systems may consume locked output
later, but they must not become kernel-owned runtime behavior.

## Prompt archive

Use available prompts in numeric order. The original bootstrap sequence is:

```text
000-readonly-review.md
001-reframe-docs-to-ni-kernel.md
002-bootstrap-cli.md
003-ni-init-docs-template.md
004-contract-model.md
005-readiness-status.md
006-lockfile-end.md
007-prompt-compiler-run.md
008-codex-skills.md
009-work-graph-proposal.md
010-generated-harness-contract.md
011-dogfood-ni-project.md
012-codex-exec-experiment-later.md
```

Later follow-up prompts may use higher task numbers, such as
`029-repo-consistency-audit.md`. Stop after each prompt, review the diff, run
validation, and commit.

## The key invariant

If the model says planning is complete but `ni status` says blocked, planning is blocked.

If the model says implementation can start but `.ni/plan.lock.json` is missing or stale, implementation is blocked.
