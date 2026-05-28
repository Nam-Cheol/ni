# Release Pipeline

이 문서는 `ni`의 publish하지 않는 release validation path를 설명한다. Release
pipeline은 repository distribution infrastructure이다. 이는 `ni-kernel`
<!-- ni-boundary-allow: explicit negative boundary statement. -->
runtime behavior가 아니며 task runner, SPEC runner, multi-agent execution
layer, adapter, queue, project execution harness가 되어서는 안 된다.

## Current Status

- Release binary availability: verified v0.3.0 GitHub Release assets와
  checksums에 대해 available이다.
- Curl installer availability: published release assets 기준으로 verified v0.3.0
  installer path에 대해 available이다.
- Homebrew availability: available하지 않다.
- Local release validation: `bash scripts/release-dry-run.sh`로 실행할 수
  있다.

## GitHub Release Workflow

Workflow는 `.github/workflows/release.yml`에 있다.

- Trigger: `v*`와 match되는 tag `push` events only.
- Permissions: tag push 이후 GoReleaser가 release assets를 만들 수 있도록
  `contents: write`를 사용한다.
- Required steps:
  - `actions/checkout@v4`로 repository를 check out한다;
  - `actions/setup-go@v5`와 `go.mod`로 Go를 set up한다;
  - `go test ./...`를 실행한다;
  - `bash scripts/quality.sh`를 실행한다;
  - `bash scripts/release-check.sh`를 실행한다;
  - GoReleaser를 `release --clean`으로 실행한다.

이 release path가 의도적으로 확장되기 전에는 branch, pull request, schedule,
manual dispatch에서 workflow를 실행하지 않는다.

## GoReleaser Archive Matrix

`.goreleaser.yaml`은 `./cmd/ni`를 `CGO_ENABLED=0`인 `ni` binary로 build한다.

| OS | Arch | Archive |
| --- | --- | --- |
| linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| darwin | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| darwin | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| windows | amd64 | `ni_<version>_windows_amd64.zip` |

Windows arm64는 의도적으로 제외한다. GoReleaser는
`ni_<version>_checksums.txt`도 작성한다.

## Local Dry Run

tag를 push하기 전에 dry run을 실행한다:

```bash
bash scripts/release-dry-run.sh
```

Dry run은 다음을 실행한다:

```text
go test ./...
bash scripts/quality.sh
bash scripts/smoke.sh
bash scripts/demo-check.sh
bash scripts/install-check.sh
bash scripts/release-check.sh
goreleaser check, if GoReleaser is installed
goreleaser release --snapshot --clean, if GoReleaser is installed
```

GoReleaser가 local에 설치되어 있지 않으면 script는 명시적인 설치 및 재실행
지침을 출력한다. 실행되지 않은 GoReleaser portion을 조용히 passed로 취급하면
안 된다.

## Stop Conditions

validation이 하나라도 실패하거나, workflow가 더 이상 `v*` tag에서만 실행되지
않거나, 명시적 release decision 없이 GoReleaser matrix가 바뀌거나, release
binary, curl installer, Homebrew availability를 검증 전에 claim하면 tag 전에
중단한다. Homebrew는 external package가 존재하고 verified되기 전까지 planned로
남는다.
