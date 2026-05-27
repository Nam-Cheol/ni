<p align="center">
  <img src="assets/hero.svg" alt="ni hero banner: Project Intent Compiler for AI Agents" width="100%">
</p>

<p align="center">
  <a href="README.md" aria-label="Read in English"><img alt="English" src="assets/badge-english.svg" width="84" height="28"></a>
  <a href="README.ko.md" aria-label="Read in Korean"><img alt="한국어" src="assets/badge-korean.svg" width="84" height="28"></a>
</p>

<p align="center">
  <a href="LICENSE"><img alt="License MIT" src="https://img.shields.io/badge/license-MIT-f4b860"></a>
  <a href=".github/workflows/ci.yml"><img alt="CI workflow exists" src="https://img.shields.io/badge/CI-workflow%20exists-25334a"></a>
  <a href="SECURITY.md"><img alt="Security policy exists" src="https://img.shields.io/badge/security-policy%20exists-2d5a52"></a>
  <a href="docs/00_START_HERE.md"><img alt="Docs index exists" src="https://img.shields.io/badge/docs-index%20exists-5b8def"></a>
</p>

<h1 align="center">Don't run the agent yet. Compile the intent first.</h1>

<p align="center"><strong>ni turns planning conversations into locked project contracts before implementation work starts.</strong></p>

`ni` is a Project Intent Compiler for AI Agents. It makes intent explicit,
checks whether the plan is ready, locks the accepted contract, and compiles a
bounded handoff prompt or derived seed material.

## Why ni

<p align="center">
  <img src="assets/card-pain-vague-intent.svg" alt="Vague intent: a plausible prompt can still hide missing users, acceptance criteria, risks, non-goals, or blockers." width="32%">
  <img src="assets/card-pain-early-execution.svg" alt="Early execution: work should not begin just because a request sounds plausible." width="32%">
  <img src="assets/card-pain-rework.svg" alt="Rework: hidden assumptions become expensive after people, models, or tools start from the wrong plan." width="32%">
</p>

### Vague intent

A prompt can sound actionable while users, acceptance criteria, risks,
non-goals, or blocker questions are still missing.

### Early execution

Work should not begin just because a request sounds plausible.

### Rework

Hidden assumptions become expensive after people, models, or tools start from
the wrong plan.

## What ni gives you

<p align="center">
  <img src="assets/card-payoff-capture-intent.svg" alt="Capture intent: planning conversation becomes explicit docs and a contract draft." width="32%">
  <img src="assets/card-payoff-lock-contract.svg" alt="Lock contract: readiness and lock commands gate the accepted plan, hashes, and source of truth." width="32%">
  <img src="assets/card-payoff-handoff-safely.svg" alt="Handoff safely: a valid locked plan compiles into a bounded prompt or derived seed material." width="32%">
</p>

### Capture intent

Planning conversation becomes explicit docs and a contract draft.

### Lock contract

`ni status` and `ni end` gate readiness, hashes, and lock creation.

### Handoff safely

`ni run` compiles a bounded prompt or seed from a valid locked plan. It does not
execute shell commands, queues, agents, or downstream work.

## Start in 60 seconds

Start from source when you are evaluating or developing from this checkout:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
```

Use conversation to fill `./my-plan/docs/plan/**` and
`./my-plan/.ni/contract.json`, then let the CLI make the authoritative call:

```bash
go run ./cmd/ni status --dir ./my-plan --next-questions
go run ./cmd/ni end --dir ./my-plan
go run ./cmd/ni run --dir ./my-plan --target generic --max-chars 4000
```

## Choose your path

| Path | Status | Use it when | Boundary |
| --- | --- | --- | --- |
| Source | Available | You can run `go run ./cmd/ni ...`. | Full deterministic `status`, `end`, and `run`. |
| Local binary | Available after local build | You want `./bin/ni` or a local install from this checkout. | Build or install locally; no hosted binary is claimed. |
| Model workspaces | Available as assisted planning | You want Codex or Claude to help draft docs and contract records. | Skills are UX; the CLI remains readiness and lock authority. |
| No-terminal method | Available as assisted method | You want to learn or draft the Intent Lock method before a CLI run. | Useful drafting, not deterministic validation. |
| Release binary | Available | You want `ni` without Go from a published release. | Download an OS/arch archive from GitHub Releases and verify it with the checksum file. |
| Curl installer | Available | You want a small shell installer for release assets. | Download and inspect `install.sh`; it installs the release archive and verifies the checksum when available. |
| Homebrew | Planned | You prefer a package manager. | No tap or formula is published. |

Release status: GitHub Release binaries and checksums are available. Package
manager distribution is not available. The curl installer has been verified
against the published `v0.3.0` release assets.

Curl installer, with inspection first:

```bash
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version 0.3.0
BINDIR="$HOME/.local/bin" sh install.sh --version 0.3.0
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

Manual release download:

1. Open <https://github.com/Nam-Cheol/ni/releases> and choose the archive for
   your OS and architecture.
2. Download the matching archive and `ni_<version>_checksums.txt` from the same
   release.
3. Verify the archive with the checksum file.
4. Unpack the archive, place `ni` on your `PATH` if desired, then run
   `ni --help` and `ni version`.

License: `ni` is licensed under the [MIT License](LICENSE).

See [Install ni](docs/22_INSTALL.md), [No-Terminal Planning](docs/no-terminal.md),
and [Model Workspace Packs](docs/55_MODEL_WORKSPACE_PACKS.md) for details.

## Demo

The best first demo is a blocked one:

```bash
go run ./cmd/ni status --dir examples/ambiguous-prompt-blocked/workspace
```

```text
BLOCKED
```

That result is the point. A vague request should stop before handoff. See the
[Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/) walkthrough.

## What ni is not

`ni` is not a task runner, spec runner, multi-agent execution layer, queue,
shell adapter, PR automation system, release automation system, or runtime for
downstream work. The kernel owns planning contracts, readiness, lockfiles, hash
checks, and prompt compilation.

## Read next

| Read | Why |
| --- | --- |
| [Why ni exists](docs/product-story.md) | The short product story behind compile-before-run. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | The deeper rules for readiness, locking, hash trust, and blocked handoff. |
| [Install ni](docs/22_INSTALL.md) | Source, local build, release binary, and curl installer details. |
| [Command reference](docs/commands.md) | The implemented CLI surface. |
| [README Visual Wireframe](docs/63_README_VISUAL_WIREFRAME.md) | The visual layout contract for this README. |
