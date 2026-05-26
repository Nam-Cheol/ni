# Initial setup manifest

This starter pack reframes `ni` as a project intent compiler rather than an execution harness.

## Main files

```text
README.md
AGENTS.md
MANIFEST.md

docs/00_START_HERE.md
docs/NI_BLUEPRINT.md
docs/01_PRODUCT_PRINCIPLES.md
docs/02_CONTRACT_MODEL.md
docs/03_READINESS_GATE.md
docs/04_LOCKFILE.md
docs/05_PROMPT_COMPILER.md
docs/06_CODEX_SKILLS.md
docs/07_GENERATED_HARNESS.md
docs/08_ROADMAP.md
docs/12_READINESS_PROFILES.md
docs/13_TARGET_HYPER_RUN.md
docs/14_TARGET_NAMBA_AI.md
docs/16_TARGET_SPEC_KIT.md
docs/17_TARGET_OUROBOROS.md
docs/18_EXAMPLES.md
docs/19_RELEASE_NOTES_v0.1.md

docs/plan/*.md
.ni/project.json
.ni/contract.json
.ni/pressure.json
.ni/readiness.rules.json
.ni/readiness.profiles.json

.agents/skills/ni-start/SKILL.md
.agents/skills/ni-end/SKILL.md
.agents/skills/ni-run/SKILL.md

prompts/*.md
scripts/*.sh
scripts/*.py
```

## How to use

1. Copy these files into a fresh `ni` repository.
2. Commit the starter state.
3. Open Codex in the repository root.
4. Run the prompts in `prompts/` in numeric order.
5. Commit after each successful prompt.
