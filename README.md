<p align="center">
  <img src="assets/hero.svg" alt="ni hero: Don't run the agent yet. Compile the intent first." width="100%">
</p>

<h1 align="center">Don't run the agent yet. Compile the intent first.</h1>

<p align="center"><strong>ni turns planning conversations into locked project contracts before AI agents or teams start work.</strong></p>

<p align="center">
  <a href="README.md"><kbd>English</kbd></a>
  <a href="README.ko.md"><kbd>한국어</kbd></a>
</p>

<p align="center">
  <a href=".github/workflows/ci.yml"><img alt="CI configured" src="https://img.shields.io/badge/CI-configured-25334a"></a>
  <a href="docs/22_INSTALL.md"><img alt="Install source-first" src="https://img.shields.io/badge/install-source--first-2d5a52"></a>
  <a href="LICENSE"><img alt="License MIT" src="https://img.shields.io/badge/license-MIT-f4b860"></a>
  <a href="docs/42_INTENT_LOCK_PROTOCOL.md"><img alt="Protocol Intent Lock" src="https://img.shields.io/badge/protocol-Intent%20Lock-7f92ff"></a>
</p>

<p align="center">
  <a href="#why-ni"><img src="assets/card-why.svg" alt="Why ni: prompts can hide users, risks, non-goals, acceptance, and blockers." width="32%"></a>
  <a href="#start-in-60-seconds"><img src="assets/card-start.svg" alt="Start path: initialize, check readiness, lock intent, and compile a prompt." width="32%"></a>
  <a href="#read-next"><img src="assets/card-docs.svg" alt="Docs map: protocol, commands, target boundaries, benchmark, and launch notes." width="32%"></a>
</p>

## Why ni

Agents fail less from lack of code ability and more from unclear intent.

`ni` is a Project Intent Compiler. It sits before execution, where vague goals
usually become hidden assumptions:

```text
planning conversation -> explicit contract -> readiness gate -> locked plan -> bounded prompt or seed
```

1. AI agents execute too early.
2. `ni` blocks ambiguous execution.
3. `ni` compiles intent into a locked project contract.
4. Then humans, models, or tools can work from that contract.

The payoff: `ni` makes unclear intent visible, blocks unsafe handoff, and
produces a bounded prompt or seed from a locked plan.

For the deeper product story, see [Why ni exists](docs/product-story.md).

## Start in 60 seconds

`ni` is currently source-first. From a checked-out repository:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
```

Now use conversation to fill `./my-plan/docs/plan/**` and
`./my-plan/.ni/contract.json`. The CLI, not the model, is the readiness
authority:

```bash
go run ./cmd/ni status --dir ./my-plan --next-questions
go run ./cmd/ni end --dir ./my-plan
go run ./cmd/ni run --dir ./my-plan --target generic --max-chars 4000
```

`ni run` compiles a prompt. It does not execute shell commands, queues, agents,
or downstream work.

## Choose your path

You do not need to be a Go user to start using the Intent Lock method. The full
deterministic `ni` gates still come from the CLI.

| Path | Status | Best for | Boundary |
| --- | --- | --- | --- |
| Source | Available | Contributors and early users who can run `go run ./cmd/ni ...`. | Full deterministic `status`, `end`, and `run`. |
| Release binary | Next | Terminal users who want `ni` without installing Go. | Wait for published GitHub Release assets and checksums. |
| Curl installer | Next | One-command install after release assets exist. | `install.sh` is present and tested locally, but public install waits on release assets. |
| Homebrew | Planned | macOS users who prefer package managers. | No tap or formula is published. |
| Claude skill pack | Available | Model-assisted planning with the packaged Claude skills. | Drafting UX only; use the CLI for readiness and lock authority. |
| Codex skill pack | Available | Model-assisted planning with the packaged Codex skills. | Drafting UX only; use the CLI for readiness and lock authority. |
| No-terminal model workflow | Available as assisted method | Teams that want to draft and review intent before anyone runs the CLI. | Not deterministic full `ni`; CLI output is required for authoritative validation. |

See [Install ni](docs/22_INSTALL.md) for source and local build paths,
[Curl Installer](docs/install-curl.md) for the release-gated installer, and
[No-Terminal Planning](docs/no-terminal.md) for the assisted workflow. See
[Distribution Strategy](docs/53_DISTRIBUTION_STRATEGY.md) and
[Homebrew Distribution Plan](docs/54_HOMEBREW_DISTRIBUTION.md) for planned
distribution tracks. Distribution automation is repository infrastructure, not
`ni` runtime execution.

This README does not claim package distribution or a published binary release.
For deterministic CLI use, use source, local build, or local install mode until a GitHub Release actually contains verified release assets; skill packs and
assisted no-terminal planning can help draft intent before that validation.

## What stays locked

The kernel owns the pre-runtime control layer:

- planning docs in `docs/plan/**`;
- `.ni/contract.json`;
- deterministic readiness through `ni status`;
- `.ni/plan.lock.json`;
- bounded prompt compilation through `ni run`.

After a lock exists, the lockfile is the source of truth. If the current plan no
longer matches the locked hashes, handoff stops as `BLOCKED`.

## What ni is not

`ni` is not a task runner, spec runner, multi-agent execution layer, queue,
shell adapter, PR automation system, release automation system, or runtime for
downstream work. Seed material is derived and mutable; the locked plan is the
authority.

## Read next

| Read | Why |
| --- | --- |
| [Why ni exists](docs/product-story.md) | The product story behind compile-before-run. |
| [Why ni](docs/why-ni.md) | The product argument, boundary, and benchmark framing. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | The rules for readiness, locking, hash trust, and blocked handoff. |
| [No-Terminal Planning](docs/no-terminal.md) | How to use the method before installing the CLI, without claiming validation. |
| [Command reference](docs/commands.md) | The implemented CLI surface. |
| [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/) | A small demo where vague intent correctly stops execution. |

## License

`ni` is licensed under the [MIT License](LICENSE).

Security policy and reporting scope are documented in [SECURITY.md](SECURITY.md).
