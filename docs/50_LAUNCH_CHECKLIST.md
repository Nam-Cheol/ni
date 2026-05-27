# Public Launch Checklist

This checklist is a deterministic pre-public-sharing gate for `ni`. It prepares
the repository for a public link, announcement draft, or maintainer review. It
does not publish packages, create GitHub releases, tag commits, post to social
channels, or run downstream agents.

Use it after the release-readiness gate and before sharing the repository
publicly:

```bash
bash scripts/launch-check.sh
```

## Positioning

- [ ] README tagline still says `Project Intent Compiler for AI Agents`.
- [ ] README explains `ni-kernel` as a deterministic pre-runtime control layer.
- [ ] README says `ni` is not an execution harness, task runner, SPEC runner,
      multi-agent execution layer, queue, shell adapter, Codex adapter, PR
      automation, or release automation.
- [ ] Positioning links point to
      [docs/40_POSITIONING.md](40_POSITIONING.md) and
      [docs/42_INTENT_LOCK_PROTOCOL.md](42_INTENT_LOCK_PROTOCOL.md).

## README Rendering

- [ ] `README.md` renders with balanced Markdown fences.
- [ ] Public demo commands in README use `go run`, `bash scripts/demo-check.sh`,
      `bash scripts/quality.sh`, or `bash scripts/install-check.sh`.
- [ ] README links to install, security, release-readiness, benchmark, target
      story, and launch checklist docs.
- [ ] README does not claim hosted binaries, package-manager distribution, or
      package publishing.

## README.ko Parity

- [ ] `README.ko.md` exists while Korean companion docs are maintained.
- [ ] `README.ko.md` links to the Korean release-readiness and launch checklist
      docs.
- [ ] `README.ko.md` keeps the same launch-sensitive claims as `README.md`:
      source-first distribution, MIT license, CI, security policy, demo
      verification, install/build verification, benchmark protocol, and no
      execution runtime.

## License

- [ ] `LICENSE` exists.
- [ ] README and README.ko state the repository is licensed under the
      [MIT License](../LICENSE).
- [ ] Launch material does not imply any additional commercial, hosted, or
      enterprise license.

## Security Policy

- [ ] `SECURITY.md` exists.
- [ ] `SECURITY.ko.md` exists while Korean companion docs are maintained.
- [ ] README links to `SECURITY.md`.
- [ ] Security docs do not claim private reporting channels, runtime sandboxing,
      downstream agent security, or enterprise support.

## CI

- [ ] `.github/workflows/ci.yml` exists.
- [ ] CI runs on `push` and `pull_request`.
- [ ] CI runs `go test ./...`, `bash scripts/quality.sh`, and
      `bash scripts/smoke.sh`.
- [ ] Public docs do not claim release, publish, package, or downstream-runtime
      automation from CI.

## Demo Verification

- [ ] [docs/48_DEMO_VERIFICATION.md](48_DEMO_VERIFICATION.md) exists.
- [ ] `bash scripts/demo-check.sh` passes.
- [ ] Public demos prove `BLOCKED`, `READY`, or `READY_WITH_DEFERRALS` states
      without running downstream agents.
- [ ] Demo outputs remain source-tree reproducible and do not require hosted
      services.

## Benchmark Protocol

- [ ] [docs/43_BENCHMARK_PROTOCOL.md](43_BENCHMARK_PROTOCOL.md) exists.
- [ ] Benchmark docs state the benchmark is not an execution benchmark.
- [ ] Benchmark docs state downstream agents must not be executed.
- [ ] Benchmark docs keep direct-to-agent comparisons separate from runtime
      performance claims.

## Release Notes

- [ ] [docs/68_RELEASE_NOTES_v0.3.0.md](68_RELEASE_NOTES_v0.3.0.md) exists.
- [ ] The release notes draft is factual and source-first.
- [ ] The draft says it does not publish a release, create a tag, upload
      binaries, or claim package-manager availability.
- [ ] The draft separates included kernel capability claims from not-included
      runtime, adapter, automation, and binary-package scope.
- [ ] [docs/69_MANUAL_RELEASE_STEPS.md](69_MANUAL_RELEASE_STEPS.md) exists and
      keeps tag push, asset confirmation, checksum confirmation, and README
      install-status updates manual.

## Install/Build Verification

- [ ] [docs/22_INSTALL.md](22_INSTALL.md) exists.
- [ ] `bash scripts/install-check.sh` passes.
- [ ] Source, local build, and temporary local install commands work.
- [ ] Install docs do not claim Homebrew, GoReleaser, package-manager, or hosted
      binary availability.

## Issue Templates

- [ ] `.github/ISSUE_TEMPLATE/bug_report.md` exists.
- [ ] `.github/ISSUE_TEMPLATE/feature_request.md` exists.
- [ ] `.github/ISSUE_TEMPLATE/boundary_question.md` exists.
- [ ] Issue templates ask reporters to preserve the ni-kernel boundary and call
      out downstream-runtime, task-runner, queue, PR automation, or release
      automation requests.

## No False Claims

- [ ] Public docs do not claim a published binary release.
- [ ] Public docs do not claim package publishing, package-manager support, or
      Homebrew support.
- [ ] Public docs do not claim hosted service availability.
- [ ] Public docs do not claim release automation.
- [ ] Public docs do not claim downstream execution, model API execution, shell
      execution, queue execution, or multi-agent orchestration.

## No Execution-Runtime Drift

- [ ] `ni run` is still documented as prompt compilation only.
- [ ] Target exports are still documented as inert downstream seed material.
- [ ] No launch item requires running Codex, Claude, model APIs, shells,
      downstream agents, queues, adapters, or generated harnesses.
- [ ] Any future public launch blocker involving execution belongs outside
      `ni-kernel` unless it is reduced to deterministic validation, locking,
      bounded prompt compilation, or inert seed generation.

## Local Gate

Run the deterministic launch gate:

```bash
bash scripts/launch-check.sh
```

The gate may call local validation scripts that already exist in the repository:
`quality.sh`, `smoke.sh`, `demo-check.sh`, `release-check.sh`, and
`install-check.sh`. It must not publish anything, create a GitHub release, tag a
commit, post announcements, upload packages, or run downstream agents.
