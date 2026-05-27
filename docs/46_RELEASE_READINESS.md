# Release Readiness

This document records release readiness facts for `ni`. It documents the
published release-binary and curl-installer paths without claiming package
manager or hosted service availability.

## Current Factual State

- License: `LICENSE` exists and contains the [MIT License](../LICENSE).
- CI: `.github/workflows/ci.yml` exists. It runs on `push` and
  `pull_request`, and it runs `go test ./...`, `bash scripts/quality.sh`, and
  `bash scripts/smoke.sh`.
- Security policy: `SECURITY.md` exists and documents the early source-first
  scope, reporting limits, secret-handling guidance, and runtime boundary.
- Distribution: documented available usage includes source, local build, local
  install, GitHub Release binaries, and the curl installer. Package manager
  distribution, Homebrew support, Scoop support, and hosted service
  availability are not available.
- Repository release workflow: `.github/workflows/release.yml` builds release
  binaries only for tagged GitHub Releases. It is distribution infrastructure,
  not `ni` runtime behavior.
- Release notes: `docs/68_RELEASE_NOTES_v0.3.0.md` is a factual GitHub Release
  notes for the first public GitHub Release. It does not claim package-manager
  availability.

## Readiness Checklist

- [ ] quality passes through `bash scripts/quality.sh`.
- [ ] tests pass through `go test ./...`.
- [ ] install-check passes through `bash scripts/install-check.sh`.
- [ ] README and README.ko are in sync for release, license, CI, security,
      install, and runtime-boundary claims.
- [ ] examples exist under `examples/` and include runnable planning
      workspaces or report templates.
- [ ] status proof works with
      `go run ./cmd/ni status --dir examples/conversation-product --proof`.
- [ ] benchmark protocol exists at
      [docs/43_BENCHMARK_PROTOCOL.md](43_BENCHMARK_PROTOCOL.md).
- [ ] v0.3.0 release notes exist and clearly separate included capability
      claims from not-included runtime, adapter, automation, and binary-package
      scope.
- [ ] no runtime execution claims are present: `ni run` compiles a bounded
      prompt only and does not execute agents, shells, queues, or adapters.
- [ ] no false release/license/CI/security claims are present:
      - release docs only claim hosted release assets that exist in the
        published GitHub Release;
      - license docs match the committed `LICENSE`;
      - CI docs match `.github/workflows/ci.yml`;
      - security docs do not claim private reporting channels, enterprise
        support, or runtime security features.

## Local Verification

Run the local release-readiness check before future release changes:

```bash
bash scripts/release-check.sh
bash scripts/release-dry-run.sh
```

The script is a readiness gate only. It does not publish packages, create a
GitHub release, upload binaries, or tag a commit. It includes
`bash scripts/install-check.sh` to verify source, local build, and temporary
local install paths.
