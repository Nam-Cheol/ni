# ni v0.2.0 — Project Intent Compiler for AI Agents

Tag suggestion: `v0.2.0`

Summary: Don't run the agent yet. Compile the intent first.

This is a draft GitHub release note for a future source-first `ni` release. It
does not publish a release, create a tag, upload binaries, or claim package
manager availability.

## Category

`ni` is a Project Intent Compiler for AI Agents. It turns planning
conversation into a checked project contract, locks the accepted intent, and
compiles bounded handoff material before downstream execution begins.

The core mechanism is the Intent Lock Protocol: a deterministic pre-runtime
control layer for deciding when intent is ready, what downstream actors may
trust, and when execution must stop because intent changed.

## Included

- Intent Lock Protocol.
- Conversation-driven authoring model.
- `ni init`, `ni status`, `ni end`, and `ni run`.
- Status proof report through `ni status --proof`.
- Locked plan hash validation.
- Target prompt and export seed outputs.
- Ambiguous prompt blocked demo.
- Non-software planning demos.
- Benchmark protocol.
- Korean companion README.

## Not Included

- Execution runtime.
- Codex adapter.
- Shell adapter.
- SPEC runner.
- Task queue.
- Multi-agent orchestration.
- PR/release automation.
- Binary package distribution.

## Source-First Distribution

This draft assumes repository-source use only: `go run ./cmd/ni ...`, local
builds, or local installs from the checked-out source tree. It does not claim
hosted release assets, Homebrew support, GoReleaser support, or published
binary packages.

## Validation Commands

Run these checks before any manual tag or GitHub release step:

```bash
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
```

If a demo-check script is added later, run it as an optional extra gate before
publishing release notes.

## Release Boundary

`ni run` compiles a prompt only. It does not execute Codex, shells, model APIs,
queues, adapters, SPEC workflows, or multi-agent systems.

This draft does not add release automation. Any tag or GitHub release should be
created manually after the local validation gate and CI pass for the exact
commit being released.
