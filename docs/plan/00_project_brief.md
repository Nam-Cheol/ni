# Project brief

## Purpose

`ni` is the Project Intent Compiler for AI Agents. In v0.4, it turns the verified v0.3.0 release, release binary, curl installer, and model workspace pack work into an adoption roadmap while preserving the same deterministic readiness, lockfile, and bounded prompt compiler boundary before any execution harness runs.

## v0.4 adoption focus

The v0.4 adoption roadmap makes the existing product easier to try and share
without turning `ni-kernel` into runtime execution:

- verify a fresh-user install path from the public v0.3.0 release and curl
  installer evidence;
- advance the Homebrew route from Planned only after tap, formula, checksum,
  audit, install, `ni --help`, and `ni version` evidence exists;
- verify model workspace packs as portable planning UX, including source,
  manual copy, zip packaging, and host-specific dry-run or install evidence
  where available;
- keep README and README.ko clear that `ni` can be used from source, local
  binary, verified release binaries, verified curl installer, and model
  workspaces, while avoiding false Homebrew or package-manager availability
  claims;
- prepare the announcement kit around the Project Intent Compiler message and
  verified install paths;
- publish a first benchmark case study as qualitative intent-readiness evidence,
  not a runtime or statistical performance claim;
- treat a lightweight static landing page as optional public doorway work, with
  README remaining the canonical quick entry;
- preserve the pre-runtime boundary: `ni run` compiles prompts only and does
  not execute Codex, shell commands, package publishing, adapters, queues, PR
  automation, or downstream work.

## v0.3 focus

The v0.3 packaging plan makes the public product surface easier to adopt without
turning `ni-kernel` into a runtime:

- `README.md` is a product pamphlet, with detailed protocol, command,
  distribution, and target material moved to docs.
- The README hero uses repository-local deterministic SVG first; generated
  image or social card assets are optional marketing outputs.
- Important product copy remains Markdown text near the visual surface; SVGs
  may carry only tested short labels and must avoid emoji, `foreignObject`,
  external fonts, external references, and long text.
- The Korean companion README stays synchronized with the English canonical
  claims and must not exceed them.
- Remote capsule-style renderers are inspiration only; primary README visuals
  come from the local asset pipeline and are guarded by README and asset
  regression checks.
- README copy avoids specific harness product claims in the hero and avoids
  claiming curl, package-manager, or binary availability before those paths are
  implemented and verified.
- The v0.3.0 distribution release state is factual: release binaries are
  Available for the verified v0.3.0 GitHub Release archives and checksums, the
  curl installer is Available for verified v0.3.0 release assets after
  `install.sh` verification, Homebrew is Planned with no published tap or
  formula, model workspace packs are UX rather than authority, and no-terminal
  mode is assisted unless a trusted runner supplies exact CLI proof.
- Distribution must keep release binaries as the trust base for curl installer
  availability; package-manager paths stay planned until published and tested.
- Model workspace packs support Codex- and Claude-style planning workflows as
  UX over docs and CLI proof, not execution adapters.
- No-terminal mode is an assisted planning workflow; it is not full
  deterministic validation unless a trusted runner supplies exact `ni` CLI
  output.

## v0.2 foundation

`ni init` creates the initial planning structure. After initialization, the primary authoring interface is sustained model-user conversation through `ni-start`, which updates `docs/plan/**` and `.ni/contract.json` together. The CLI remains authoritative for deterministic gaps (`ni status`), explicit lock or relock (`ni end` or `ni relock`), and prompt compilation (`ni run`).

User-facing contract `add`, `list`, or `set` commands are not part of the v0.2 primary authoring UX.

The public v0.2 message is: do not run the agent yet; compile the intent first. The unique mechanism is the Intent Lock Protocol, which defines how planning conversations become a contract, when the contract is ready to lock, how the accepted plan is hashed, what downstream actors may trust, and when execution must stop because intent changed.

## Differentiation and packaging proof assets

The v0.2 and v0.3 direction is supported by:

- ambiguous prompt blocking demo,
- non-software planning demo,
- intent readiness benchmark protocol,
- status proof report,
- downstream target story,
- README relaunch,
- README.ko companion sync,
- release readiness checklist,
- README pamphlet strategy,
- distribution strategy and Homebrew plan,
- model workspace pack docs for Codex and Claude,
- visual asset rules.

## Later direction

The later roadmap keeps the kernel narrow while making the planning contract useful across more project types and downstream targets. The kernel owns:

- related-work boundaries,
- readiness profiles,
- product type and delivery-surface guidance,
- downstream target registry,
- locked target exports,
- feedback ingest,
- pressure ledger,
- harness candidate lifecycle,
- amendment and relock flow,
- collaboration diff and conflict checks.

## Problem

Agent and SPEC systems often mix planning, execution, evidence collection, and project growth into one runtime. `ni` separates those concerns. Planning can become explicit and locked before any generated harness or downstream runtime begins work.

## Success definition

A user can start with `ni init`, plan through model-user conversation, let `ni-start` update planning docs and `.ni/contract.json`, run `ni status --proof`, lock or relock through the CLI, and compile a 4000-character-or-less target prompt from the valid lock. The resulting contract presents `ni-kernel` as a pre-runtime intent compiler and packages `ni` through truthful README, distribution, visual, and model workspace surfaces, not as a task runner, SPEC runner, package availability claim, model execution adapter, or multi-agent execution layer.

## Boundary

Generated harnesses, seed packages, prompts, and downstream feedback are derived material. They may inform future amendments, but they do not replace `.ni/plan.lock.json`, `.ni/contract.json`, or `docs/plan/**` as source of truth.
