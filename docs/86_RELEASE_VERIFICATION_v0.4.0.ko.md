# v0.4.0 릴리스 검증

날짜: 2026-05-29

범위: 수동 태그 푸시와 GitHub Actions 릴리스 빌드 이후 게시된 GitHub Release
`v0.4.0` 자산을 검증한다. 이 release-asset verification은 새 릴리스를
게시하지 않고, 태그를 푸시하지 않고, 릴리스 버전을 바꾸지 않고, Homebrew를
Available로 표시하지 않고, 패키지 매니저 availability claim을 추가하지 않고,
runtime execution behavior를 추가하지 않는다.

저장소 릴리스:
`https://github.com/Nam-Cheol/ni/releases/tag/v0.4.0`

릴리스 메타데이터:

- 태그: `v0.4.0`
- 이름: `v0.4.0`
- 초안: `false`
- 프리릴리스: `false`
- 게시 시각: `2026-05-29T06:58:39Z`
- 대상: `main`

## 자산 목록

게시된 릴리스에는 예상 OS/아키텍처 자산이 모두 포함되어 있다.

| 자산 | 크기 | 다이제스트 |
| --- | ---: | --- |
| `ni_0.4.0_checksums.txt` | 471 | `sha256:b24746a824084b01ebcc5c706afe9e5e43ca8bdb2b8606a766edcaa1bb4a70ca` |
| `ni_0.4.0_darwin_amd64.tar.gz` | 1,209,092 | `sha256:b7b503cf998a963fc21174967617710bdb0d8efecd9d4202e6899f2e0e36a9a1` |
| `ni_0.4.0_darwin_arm64.tar.gz` | 1,149,041 | `sha256:da5e3b715a79e2a284b3095c2e2f5e813ecccc7b2ea89657e363af48d2603813` |
| `ni_0.4.0_linux_amd64.tar.gz` | 1,190,005 | `sha256:ca4733be1cc67417fbd51b2d6ef3866b6d3eb3021feaa1f049e39f5e6266dd6f` |
| `ni_0.4.0_linux_arm64.tar.gz` | 1,107,207 | `sha256:4e2bffb3505a11b6a0bd6ff2939994f7629aa715768a77a96357b2f062b22733` |
| `ni_0.4.0_windows_amd64.zip` | 1,249,396 | `sha256:95f9c09b815327e0040bd0263fe4db35949472e99bbc22d70f916789bb101f2f` |

예상 플랫폼 범위가 모두 존재한다.

- 체크섬 파일
- macOS amd64
- macOS arm64
- Linux amd64
- Linux arm64
- Windows amd64

## 다운로드

자산은 임시 로컬 디렉터리에 다운로드했다.

```text
/private/tmp/ni-v0.4.0-assets.leSaCK
```

다운로드한 파일:

```text
ni_0.4.0_checksums.txt
ni_0.4.0_darwin_amd64.tar.gz
ni_0.4.0_darwin_arm64.tar.gz
ni_0.4.0_linux_amd64.tar.gz
ni_0.4.0_linux_arm64.tar.gz
ni_0.4.0_windows_amd64.zip
```

## 체크섬 검증

명령:

```bash
shasum -a 256 -c ni_0.4.0_checksums.txt
```

출력:

```text
ni_0.4.0_darwin_amd64.tar.gz: OK
ni_0.4.0_darwin_arm64.tar.gz: OK
ni_0.4.0_linux_amd64.tar.gz: OK
ni_0.4.0_linux_arm64.tar.gz: OK
ni_0.4.0_windows_amd64.zip: OK
```

## 아카이브 검증

각 아카이브는 정상적으로 압축 해제되었다.

아카이브 내용:

```text
darwin amd64:  LICENSE, README.md, README.ko.md, ni
darwin arm64:  LICENSE, README.md, README.ko.md, ni
linux amd64:   LICENSE, README.md, README.ko.md, ni
linux arm64:   LICENSE, README.md, README.ko.md, ni
windows amd64: LICENSE, README.md, README.ko.md, ni.exe
```

바이너리 형식 확인:

```text
extract/darwin_amd64/ni:      Mach-O 64-bit executable x86_64
extract/darwin_arm64/ni:      Mach-O 64-bit executable arm64
extract/linux_amd64/ni:       ELF 64-bit LSB executable, x86-64, statically linked
extract/linux_arm64/ni:       ELF 64-bit LSB executable, ARM aarch64, statically linked
extract/windows_amd64/ni.exe: PE32+ executable (console) x86-64, for MS Windows
```

## 현재 플랫폼 바이너리

현재 검증 플랫폼:

```text
Darwin arm64
```

현재 플랫폼 자산:

```text
ni_0.4.0_darwin_arm64.tar.gz
```

`ni --help` 출력:

```text
ni is a project intent compiler.

Usage:
  ni --help
  ni amend create --title <title> [--dir <path>]
  ni amend list [--dir <path>]
  ni amend show <id> [--dir <path>]
  ni amend apply <id> [--dir <path>]
  ni conflicts --base <path-or-lock> --head <path-or-lock> [--json]
  ni diff --base <path-or-lock> --head <path-or-lock> [--json]
  ni end --dir <path>
  ni export --target hyper-run|namba-ai|ouroboros|spec-kit --out <dir> [--dir <path>]
  ni feedback add --file <path> [--dir <path>]
  ni feedback list [--dir <path>] [--json]
  ni graph --dir <path> [--json]
  ni harness plan --dir <path> [--json]
  ni harness candidates [--dir <path>] [--json]
  ni harness propose --from-pressure <id> [--dir <path>]
  ni harness validate <candidate-id> --evidence <path> [--dir <path>]
  ni harness accept <candidate-id> [--dir <path>]
  ni harness retire <candidate-id> [--dir <path>]
  ni init --dir <path> [--profile concept|prototype|mvp|beta|production] [--product-type <type>] [--surface <surface>] [--interaction-mode <mode>]
  ni pressure status [--dir <path>] [--json]
  ni pressure promote <id> [--dir <path>]
  ni pressure retire <id> [--dir <path>]
  ni relock --dir <path>
  ni run --dir <path> [--target <target>] [--out <path>] [--max-chars N]
  ni status --dir <path> [--json] [--proof] [--next-questions]
  ni targets [--json]
  ni version

Commands:
  amend   Create, inspect, and apply explicit contract amendments.
  conflicts Detect semantic planning conflicts between two contracts or locked plans.
  diff     Show contract-level changes between two contracts or locked plans.
  end      Lock the accepted planning contract.
  export   Write locked-plan seed artifacts for a downstream target.
  feedback Record and list inert downstream feedback.
  graph    Propose a read-only work graph.
  harness  Manage inert generated harness proposals.
  init     Create planning docs and .ni skeleton.
  pressure Track inert planning pressure without changing readiness rules.
  relock   Create a new lock from an explicitly amended plan.
  run      Compile a goal prompt from the locked plan.
  status   Validate planning readiness.
  targets  List downstream prompt/export targets.
  version  Print the ni version.
```

`ni version` 출력:

```text
0.4.0
```

## 결과

게시된 `v0.4.0` 릴리스 자산은 verified 상태다.

체크섬이 통과했고, 모든 아카이브가 정상적으로 압축 해제되었으며, 현재
플랫폼이 아닌 아카이브의 바이너리 이름도 기대와 일치한다. 현재 플랫폼
바이너리는 `ni --help`와 `ni version`을 정상 실행했고, version output은
`0.4.0`이다.

이 검증은 Homebrew를 Available로 표시하지 않으며 패키지 매니저 claim을
추가하지 않는다.
