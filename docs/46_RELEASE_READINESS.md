# Release Readiness

This document is a pre-publish readiness record. It prepares `ni` for a future
release without announcing one, publishing one, or claiming package-manager or
hosted binary availability.

## Current Factual State

- License: `LICENSE` exists and contains the [MIT License](../LICENSE).
- CI: `.github/workflows/ci.yml` exists. It runs on `push` and
  `pull_request`, and it runs `go test ./...`, `bash scripts/quality.sh`, and
  `bash scripts/smoke.sh`.
- Security policy: `SECURITY.md` does not exist yet.
  TODO: define the vulnerability reporting policy before linking or advertising
  a security policy.
- Distribution: documented usage remains source, local build, and local install
  only. This document does not claim package distribution, Homebrew support,
  GoReleaser support, or a published binary release.
- Release automation: no release automation is part of release readiness.

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
- [ ] no runtime execution claims are present: `ni run` compiles a bounded
      prompt only and does not execute agents, shells, queues, or adapters.
- [ ] no false release/license/CI/security claims are present:
      - release docs do not claim hosted release or package availability;
      - license docs match the committed `LICENSE`;
      - CI docs match `.github/workflows/ci.yml`;
      - security docs do not claim a published policy before `SECURITY.md`
        exists.

## Local Verification

Run the local release-readiness check before any future release step:

```bash
bash scripts/release-check.sh
```

The script is a readiness gate only. It does not publish packages, create a
GitHub release, add release automation, or tag a commit.
