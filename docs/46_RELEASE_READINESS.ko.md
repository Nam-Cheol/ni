# Release Readiness

이 문서는 `ni`의 release readiness facts를 기록한다. Published release-binary와
curl-installer paths는 문서화하지만 package manager나 hosted service
availability는 claim하지 않는다.

## Current Factual State

- License: `LICENSE`가 존재하며 [MIT License](../LICENSE)를 포함한다.
- CI: `.github/workflows/ci.yml`이 존재한다. 이 workflow는 `push`와
  `pull_request`에서 실행되며 `go test ./...`, `bash scripts/quality.sh`,
  `bash scripts/smoke.sh`를 실행한다.
- Security policy: `SECURITY.md`가 존재하며 early source-first scope,
  reporting limits, secret-handling guidance, runtime boundary를 문서화한다.
- Distribution: documented available usage는 source, local build, local install,
  verified v0.3.0 GitHub Release binaries, verified v0.3.0 curl installer
  path다. Package manager distribution, Homebrew support, Scoop support, hosted
  service availability는 available하지 않다.
- Repository release workflow: `.github/workflows/release.yml`은 tagged GitHub
  Releases에 대해서만 release binaries를 build한다. 이는 distribution
  infrastructure이지 `ni` runtime behavior가 아니다.
- Release notes: `docs/68_RELEASE_NOTES_v0.3.0.ko.md`는 first public GitHub
  Release notes이다. Package-manager availability를 claim하지 않는다.

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
- [ ] no runtime execution claims are present: `ni run`은 bounded prompt만
      compile하며 agents, shells, queues, adapters를 실행하지 않는다.
- [ ] no false release/license/CI/security claims are present:
      - release docs는 published GitHub Release에 존재하는 hosted release
        assets만 claim한다;
      - license docs는 committed `LICENSE`와 일치한다;
      - CI docs는 `.github/workflows/ci.yml`과 일치한다;
      - security docs는 private reporting channels, enterprise support, runtime
        security features를 claim하지 않는다.

## Local Verification

future release changes 전에 local release-readiness check를 실행한다:

```bash
bash scripts/release-check.sh
bash scripts/release-dry-run.sh
```

This script is a readiness gate only. It does not publish packages, create a
GitHub release, upload binaries, or tag a commit. Source, local build, temporary
local install paths는 `bash scripts/install-check.sh`로 검증한다.
