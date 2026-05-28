# v0.3.0 릴리스 검증

날짜: 2026-05-28

범위: 릴리스 바이너리를 Available로 표시하기 전에 게시된 GitHub Release
`v0.3.0`의 자산을 검증한다. 이 검증은 curl 설치 스크립트, Homebrew,
Scoop 또는 패키지 매니저 배포를 Available로 표시하지 않는다.

저장소 릴리스:
`https://github.com/Nam-Cheol/ni/releases/tag/v0.3.0`

릴리스 메타데이터:

- 태그: `v0.3.0`
- 이름: `v0.3.0`
- 초안: `false`
- 프리릴리스: `false`
- 게시 시각: `2026-05-27T15:20:01Z`
- 대상: `main`

## 자산 목록

게시된 릴리스에는 예상 OS/아키텍처 자산이 모두 포함되어 있다.

| 자산 | 크기 | 다이제스트 |
| --- | ---: | --- |
| `ni_0.3.0_darwin_amd64.tar.gz` | 1,166,535 | `sha256:b6d65b177f0a58e7c9457fc562494e8d6dfdc92655aa0b1bb4aa697a8da952e0` |
| `ni_0.3.0_darwin_arm64.tar.gz` | 1,101,915 | `sha256:a41a45afb0e1f11779b28d70f397430773d7ad5f23252771077cc8fafefe0f33` |
| `ni_0.3.0_linux_amd64.tar.gz` | 1,148,197 | `sha256:7032a70dbe8e3824b10c6fa83e315507d8d135c89fe1cf0cc1597ebab19896e9` |
| `ni_0.3.0_linux_arm64.tar.gz` | 1,063,788 | `sha256:e7401a78465f2401c1948a05c2a4c646dfc9e6f0be834e8f0b888a466e3b20f9` |
| `ni_0.3.0_windows_amd64.zip` | 1,206,050 | `sha256:068d3a9ad0a857bf773f4c522f1e1803cc3e11f0d0b49bbef71d3b183f1e1267` |
| `ni_0.3.0_checksums.txt` | 471 | `sha256:b961642164db1b751e62bda0c5d489e23f901c0c2838b5206611dc9fa1557f44` |

예상 플랫폼 범위가 모두 존재한다.

- Linux amd64
- Linux arm64
- macOS amd64
- macOS arm64
- Windows amd64
- 체크섬 파일

## 다운로드

자산은 임시 로컬 디렉터리에 다운로드했다.

```text
/private/tmp/ni-v0.3.0-assets.9wZywv
```

다운로드한 파일:

```text
ni_0.3.0_checksums.txt
ni_0.3.0_darwin_amd64.tar.gz
ni_0.3.0_darwin_arm64.tar.gz
ni_0.3.0_linux_amd64.tar.gz
ni_0.3.0_linux_arm64.tar.gz
ni_0.3.0_windows_amd64.zip
```

## 체크섬 검증

명령:

```bash
shasum -a 256 -c ni_0.3.0_checksums.txt
```

출력:

```text
ni_0.3.0_darwin_amd64.tar.gz: OK
ni_0.3.0_darwin_arm64.tar.gz: OK
ni_0.3.0_linux_amd64.tar.gz: OK
ni_0.3.0_linux_arm64.tar.gz: OK
ni_0.3.0_windows_amd64.zip: OK
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
ni_0.3.0_darwin_arm64.tar.gz
```

`ni --help`는 정상 실행되었고 CLI 도움말을 출력했다.

`ni version` 출력:

```text
0.3.0
```

## 결과

게시된 `v0.3.0` 릴리스 자산은 예상 OS/아키텍처 매트릭스에서 사용 가능하며,
체크섬이 검증되었고, 모든 아카이브가 압축 해제되었으며, 현재 플랫폼
바이너리가 정상 실행되었다.

릴리스 바이너리는 이제 Available로 표시할 수 있다.

curl 설치 스크립트, Homebrew, Scoop, 패키지 매니저 배포는 별도 검증 전까지
Available 상태가 아니다.
