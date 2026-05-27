<p align="center">
  <img src="assets/hero.svg" alt="ni hero: Don't run the agent yet. Compile the intent first." width="100%">
</p>

<h1 align="center">ni</h1>

<p align="center"><strong>Project Intent Compiler for AI Agents.</strong></p>

<p align="center">
  <a href="README.md"><kbd>English</kbd></a>
  <a href="README.ko.md"><kbd>한국어</kbd></a>
</p>

<p align="center">
  <a href=".github/workflows/ci.yml"><kbd>CI</kbd></a>
  <a href="SECURITY.md"><kbd>Security</kbd></a>
  <a href="LICENSE"><kbd>MIT License</kbd></a>
  <a href="docs/00_START_HERE.md"><kbd>Docs</kbd></a>
</p>

## Don't Run The Agent Yet

`ni` turns planning conversations into locked project contracts before AI
agents or teams start work.

The current product is `ni-kernel`: a deterministic pre-runtime control layer
for intent. It validates whether a plan is explicit enough to trust, locks the
accepted version, and compiles bounded handoff prompts or inert seed material.
It does not run the downstream work.

```text
conversation -> planning contract -> readiness gate -> intent lock -> bounded handoff
```

## Why ni

Prompts can sound actionable while still hiding the details that make execution
trustworthy: users, acceptance criteria, risks, non-goals, blocker questions,
and whether the accepted plan has changed.

`ni` moves control before execution. It makes intent explicit, checks it
deterministically, locks the accepted snapshot, and stops handoff when the
current files no longer match the lock.

Read the deeper product argument in [Why ni](docs/why-ni.md).

## Start In 3 Steps

```bash
go run ./cmd/ni init --dir <path> --profile prototype
go run ./cmd/ni status --dir <path>
go run ./cmd/ni end --dir <path>
go run ./cmd/ni run --dir <path> --target codex --max-chars 4000
```

1. Create a planning workspace.
2. Use model-user conversation to maintain `docs/plan/**` and
   `.ni/contract.json`; `ni status` is the readiness authority.
3. Lock with `ni end`, then compile a bounded prompt with `ni run`.

`ni run` prints or writes a prompt. It does not execute Codex, shell commands,
agents, queues, or adapters.

## Visual Flow

```text
vague request
  -> docs/plan/** + .ni/contract.json
  -> ni status
  -> ni end
  -> .ni/plan.lock.json
  -> ni run / ni export
  -> downstream prompt or inert seed material
```

## One Demo

The fastest demo is [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/).
It starts from:

```text
Build me a dashboard for my team.
```

Run:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

Expected result:

```text
BLOCKED
```

That is success: execution has not earned a trustworthy starting point yet.

## Install And Use

`ni` is currently source-first. It does not claim package-manager distribution,
hosted service availability, or a published binary release.
Release status: does not claim package distribution or a published binary release.

```bash
go run ./cmd/ni --help
make build
make install-local
```

See [Install](docs/22_INSTALL.md) for source, local build, or local install
details.

## Trust Signals

- deterministic readiness through `ni status`;
- lockfile hash verification through `.ni/plan.lock.json`;
- prompt compilation budget, initially 4000 characters;
- no downstream execution in `ni run`;
- repository validation through `bash scripts/quality.sh`;
- Public demo verification through `bash scripts/demo-check.sh`;
- Source/build/install verification through `bash scripts/install-check.sh`;
- licensed under the [MIT License](LICENSE);
- CI validation in `.github/workflows/ci.yml`.

## What ni Is Not

`ni` is not a task runner, spec runner, multi-agent execution layer, queue,
shell adapter, PR automation system, release automation system, or
downstream-specific runtime. Downstream prompts and seed packages are derived
and mutable; they do not become kernel-owned execution state.

## Read Next

- [Why ni](docs/why-ni.md)
- [How ni works](docs/how-ni-works.md)
- [Intent Lock Protocol](docs/intent-lock-protocol.md)
- [Protocol Specification](docs/42_INTENT_LOCK_PROTOCOL.md)
- [Command Reference](docs/commands.md)
- [Positioning](docs/40_POSITIONING.md)
- [Target Story](docs/45_TARGET_STORY.md)
- [Benchmark Protocol](docs/43_BENCHMARK_PROTOCOL.md)
- [Release Readiness](docs/46_RELEASE_READINESS.md)
- [Launch Checklist](docs/50_LAUNCH_CHECKLIST.md)
