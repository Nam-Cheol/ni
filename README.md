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

## Use ni anywhere

| Path | Status | What it means |
| --- | --- | --- |
| Source CLI | Available now | Run with `go run ./cmd/ni ...` while developing or trying the kernel. |
| Local binary | Available now | Build with `make build`, then run `./bin/ni ...`. |
| Local install | Available now | Install to a local bin path with `make install-local`. |
| Model workspace | Available now | Compile a locked prompt for a model workspace such as Codex or Claude; `ni` only produces the handoff text. |
| Package manager install | Planned | Not published yet. Use source or local install today. |
| Hosted service | Planned | Not available yet. |
| Model plugin | Planned | Not available yet. |

See [Install ni](docs/22_INSTALL.md) for the supported local paths.

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
| [Why ni](docs/why-ni.md) | The product argument and positioning. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | The rules for readiness, locking, hash trust, and blocked handoff. |
| [Command reference](docs/commands.md) | The implemented CLI surface. |
| [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/) | A small demo where vague intent correctly stops execution. |
