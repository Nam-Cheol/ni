# Release Readiness

이 문서는 publish 전 release readiness 기록이다. `ni`의 future release를
준비하지만 release를 announce, publish, 또는 package manager나 hosted binary
availability가 있다고 claim하지 않는다.

## Current Factual State

- License: `LICENSE`가 존재하며 [MIT License](../LICENSE)를 포함한다.
- CI: `.github/workflows/ci.yml`이 존재한다. 이 workflow는 `push`와
  `pull_request`에서 실행되며 `go test ./...`, `bash scripts/quality.sh`,
  `bash scripts/smoke.sh`를 실행한다.
- Security policy: `SECURITY.md`가 존재하며 early source-first scope,
  reporting limits, secret-handling guidance, runtime boundary를 문서화한다.
- Distribution: documented available usage는 source, local build, local
  install이다. Future release assets를 위한 GoReleaser configuration과
  tag-triggered GitHub release workflow가 있지만, 이 문서는 hosted binary
  availability, package distribution, Homebrew support, Scoop support를
  claim하지 않는다.
- Repository release workflow: `.github/workflows/release.yml`은 tagged GitHub
  Releases에 대해서만 release binaries를 build한다. 이는 distribution
  infrastructure이지 `ni` runtime behavior가 아니다.
- Release draft: `docs/47_RELEASE_DRAFT_v0.2.0.ko.md`는 factual GitHub release
  draft일 뿐이다. Tag `v0.2.0`을 suggest하지만 tag 생성, release publish,
  binary upload, package-manager availability claim을 하지 않는다.

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
- [ ] v0.2.0 release draft exists and clearly separates included capability
      claims from not-included runtime, adapter, automation, and binary-package
      scope.
- [ ] no runtime execution claims are present: `ni run`은 bounded prompt만
      compile하며 agents, shells, queues, adapters를 실행하지 않는다.
- [ ] no false release/license/CI/security claims are present:
      - release docs는 published release가 assets를 포함하기 전에 hosted release
        assets availability를 claim하지 않는다;
      - license docs는 committed `LICENSE`와 일치한다;
      - CI docs는 `.github/workflows/ci.yml`과 일치한다;
      - security docs는 private reporting channels, enterprise support, runtime
        security features를 claim하지 않는다.

## Local Verification

future release step 전에 local release-readiness check를 실행한다:

```bash
bash scripts/release-check.sh
```

This script is a readiness gate only. It does not publish packages, create a
GitHub release, upload binaries, or tag a commit. Source, local build, temporary
local install paths는 `bash scripts/install-check.sh`로 검증한다.
