# Manual Release Steps

이 절차는 public GitHub Release를 수동으로 준비하고 publish하기 위한 것이다. 이는
`ni-kernel` behavior가 아니며 product 내부 release automation으로 바뀌면 안 된다.

Planned version: `v0.3.0`

## Before Tagging

1. git tree가 clean한지 확인한다:

   ```bash
   git status --short
   ```

2. publish하지 않는 release dry run을 실행한다:

   ```bash
   bash scripts/release-dry-run.sh
   ```

3. normal quality gate를 실행한다:

   ```bash
   bash scripts/quality.sh
   go test ./...
   ```

4. [docs/68_RELEASE_NOTES_v0.3.0.md](68_RELEASE_NOTES_v0.3.0.md)를 다시 읽고 must
   not claim:
   hosted binaries, curl installer availability, package manager distribution,
   runtime execution, adapters, queues, PR/release automation inside
   `ni-kernel`을 claim하지 않는지 확인한다.

## Tag And Publish Assets

1. local에서 annotated tag를 만든다:

   ```bash
   git tag -a v0.3.0 -m "ni v0.3.0"
   ```

2. tag만 push한다:

   ```bash
   git push origin v0.3.0
   ```

3. GitHub Actions release workflow가 끝날 때까지 기다린다.

4. GitHub Release가 존재하고 expected archives를 모두 포함하는지 확인한다:

   - `ni_<version>_linux_amd64.tar.gz`
   - `ni_<version>_linux_arm64.tar.gz`
   - `ni_<version>_darwin_amd64.tar.gz`
   - `ni_<version>_darwin_arm64.tar.gz`
   - `ni_<version>_windows_amd64.zip`
   - `ni_<version>_checksums.txt`

5. Checksums가 uploaded archives와 일치하는지 확인한다.

6. Assets와 checksums가 존재한 뒤에만 README install status를 release-gated 또는
   planned wording에서 release binary path available wording으로 변경한다. Published
   release assets에 대해 installer를 검증하기 전에는 curl installer availability를
   available로 표시하지 않는다.

## Stop Conditions

Release notes must not imply excluded runtime or automation boundaries are
included.

Check가 실패하거나, tag가 wrong commit을 가리키거나, assets가 missing이거나,
checksums가 일치하지 않거나, release notes가 `ni`를 task runner, SPEC runner,
multi-agent execution layer, Codex exec adapter, shell adapter, queue, PR
automation system, release automation system, downstream execution runtime으로
암시하면 release를 중단한다.
