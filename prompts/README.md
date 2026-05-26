# Prompt sequence

Run available prompts in numeric order. Treat each prompt as one focused task
and one commit.

This archive does not imply the project stopped at `012`. The original
bootstrap sequence is preserved, and later task prompts may use higher task
numbers when they are added for follow-up audits or release work.

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
029-repo-consistency-audit.md
```

Do not skip directly to execution adapters. The kernel must validate, lock, and
compile prompts first. Later prompts must preserve the same product boundary:
`ni-kernel` owns planning authority, and generated harnesses remain downstream
seed material.
