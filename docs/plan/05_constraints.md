# Constraints

## Hard constraints

- `ni run` prompt output must be 4000 characters or less when `--max-chars 4000` is used.
- Readiness must be rule-based, not model-feeling-based.
- Lockfile hash mismatch must block prompt compilation and target export.
- Codex is an adapter or UX target, not the kernel.
- Target exports must be seed material only.
- Feedback, pressure, and harness candidates must not silently change locked planning docs.
- Relock must require an explicit applied amendment when a prior lock exists.
- Collaboration checks must be deterministic and contract-local.
- After `ni init`, v0.2 authoring must flow through model-user conversation that updates docs and contract together.
- User-facing contract `add`, `list`, or `set` commands must not become the v0.2 primary authoring UX.
- Differentiation proof assets must remain pre-runtime evidence: demos, benchmark protocols, proof reports, target stories, README sync, and release checklists must not execute downstream agents or become kernel-owned runtime state.
- README must be a product pamphlet; technical details belong in docs and should be linked rather than expanded inline.
- README hero copy must avoid specific downstream harness product mentions.
- README and docs must not claim release binary, curl, Homebrew, or package-manager availability until those paths are implemented and verified.
- Release binary, curl installer, Homebrew, and package-manager status must
  match README.md, README.ko.md, docs/22_INSTALL.md,
  docs/53_DISTRIBUTION_STRATEGY.md, docs/54_HOMEBREW_DISTRIBUTION.md, and
  docs/69_MANUAL_RELEASE_STEPS.md; release-gated or planned paths stay labeled
  that way until verified evidence exists.
- The README hero uses SVG first; generated images and social cards are optional marketing assets, not kernel behavior.
- The README hero and core visual assets must be local deterministic SVG from
  the repository asset pipeline; remote capsule-style renderers are inspiration
  only and not a primary README dependency.
- Important product copy must remain Markdown text or accessible alt/link text;
  SVGs may include only short tested labels and must avoid emoji,
  `foreignObject`, external fonts, external references, and long text.
- The Korean companion README must stay within the canonical English README
  claims and must not add stronger install, package manager, runtime, or product
  availability claims.
- Visual regression checks must guard the README and assets surface.
- Distribution must support non-Go users through release binaries before curl or package-manager paths are presented as available.
- Model workspace packs may support Codex- and Claude-style planning workflows, but they must remain UX over docs and CLI proof.
- No-terminal mode is assisted planning only unless exact CLI output from a trusted runner supplies deterministic validation.

## Kernel boundary

`ni-kernel` owns docs contract, readiness gate, lockfile, prompt compiler, target registry, inert feedback and pressure ledgers, amendment/relock, and collaboration conflict checks.

`ni-generated-harness` owns project-specific work graphs, evaluation plans, evidence rules, adapters, and runtime decisions.

## Forbidden v0.2 behavior

- Do not add a shell adapter.
- Do not add a Codex execution adapter.
- Do not add an evidence runner.
- Do not add a queue.
- Do not add PR automation.
- Do not add release automation.
- Do not add a plugin system.
- Do not add a TUI or web UI.
- Do not add primary contract editing commands that make users hand-maintain `.ni/contract.json`.
- Do not turn model packs into execution adapters.
- Do not falsely claim package, curl, Homebrew, or binary availability.
- Do not replace the local README hero with a remote capsule-style renderer or
  make essential product copy image-only.
