# Release Readiness

This document is a pre-publish readiness record. It prepares `ni` for a future
release without announcing one, publishing one, or claiming package-manager or
hosted binary availability.

## Current Factual State

- License: `LICENSE` exists and contains the [MIT License](../LICENSE).
- CI: `.github/workflows/ci.yml` exists. It runs on `push` and
  `pull_request`, and it runs `go test ./...`, `bash scripts/quality.sh`, and
  `bash scripts/smoke.sh`.
- Security policy: `SECURITY.md` exists and documents the early source-first
  scope, reporting limits, secret-handling guidance, and runtime boundary.
- Distribution: documented usage remains source, local build, and local install
  only. This document does not claim package distribution, Homebrew support,
  GoReleaser support, or a published binary release.
- Release automation: no release automation is part of release readiness.
- Release draft: `docs/47_RELEASE_DRAFT_v0.2.0.md` is a factual GitHub release
  draft only. It suggests tag `v0.2.0` but does not create a tag, publish a
  release, upload binaries, or claim package-manager availability.

## Readiness Checklist

- [ ] quality passes through `bash scripts/quality.sh`.
- [ ] tests pass through `go test ./...`.
- [ ] README and README.ko are in sync for release, license, CI, security,
      install, and runtime-boundary claims.
- [ ] examples exist under `examples/` and include runnable planning
      workspaces or report templates.
- [ ] status proof works with
      `go run ./cmd/ni status --dir examples/conversation-product --proof`.
- [ ] benchmark protocol exists at
      [docs/43_BENCHMARK_PROTOCOL.md](43_BENCHMARK_PROTOCOL.md).
- [ ] v0.2.0 release draft exists and clearly separates included capability
      claims from not-included runtime, adapter, automation, and binary-package
      scope.
- [ ] no runtime execution claims are present: `ni run` compiles a bounded
      prompt only and does not execute agents, shells, queues, or adapters.
- [ ] no false release/license/CI/security claims are present:
      - release docs do not claim hosted release or package availability;
      - license docs match the committed `LICENSE`;
      - CI docs match `.github/workflows/ci.yml`;
      - security docs do not claim private reporting channels, enterprise
        support, or runtime security features.

## Local Verification

Run the local release-readiness check before any future release step:

```bash
bash scripts/release-check.sh
```

The script is a readiness gate only. It does not publish packages, create a
GitHub release, add release automation, or tag a commit.
