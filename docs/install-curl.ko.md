# Curl Installer

`install.sh`는 Go 없이 released `ni` binary를 설치한다. 이것은 release asset
infrastructure일 뿐이다. Archive를 download하고, release가 checksum을 제공하면
verify한 뒤, `ni`를 local bin directory에 복사하고 next steps를 출력한다. Model
skills를 install하지 않으며 downstream work를 실행하지 않는다. 명시적으로 opt in하면
reversible zsh/bash PATH block을 추가할 수 있다.

Status: verified v0.5.0 GitHub Release assets에 대해 Available이다. `install.sh`는
선택된 archive와 `ni_0.5.0_checksums.txt`를 download하고, local sha256 tool이
있으면 archive를 verify하고, `ni` binary만 install하며 downstream work를 실행하지
않는다.

## 더 안전한 Script 경로

Local install 전에 installer를 먼저 download하고 inspect한다:

```bash
VERSION="0.5.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

기본 install 위치는 `~/.local/bin/ni`다. 다른 위치는 `BINDIR`로 지정한다:

```bash
BINDIR="$HOME/bin" sh install.sh --dry-run --version "$VERSION"
```

`--version`을 생략하면 installer는 GitHub에서 latest release tag를 확인한다.
[v0.5.0 Post-Release Verification](117_V0_5_0_POST_RELEASE_VERIFICATION.ko.md)이
검증한 release를 원하면 `VERSION="0.5.0"`으로 고정한다. 설치 후 새 shell을 열고
global command를 help 또는 version command로 확인한다:

```bash
ni --help
ni version
```

Global command verification 뒤에는 project directory에서 planning workspace를
시작한다:

```bash
mkdir my-project
cd my-project
ni init .
```

Binary와, 추가했다면 ni-managed PATH block을 uninstall한다:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

## Manual Verification Path

Manual install은 trust step을 모두 보이게 유지한다.

1. Release page를 연다:
   <https://github.com/Nam-Cheol/ni/releases>
2. Release에 assets가 있는지 확인한 뒤 platform에 맞는 archive를 고른다:

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `ni_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `ni_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `ni_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `ni_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `ni_<version>_windows_amd64.zip` |

3. 같은 release에서 archive와 `ni_<version>_checksums.txt`를 download한다.
4. Archive checksum을 verify한다.

Linux:

```bash
VERSION="<published-version-without-v>"
grep " ni_${VERSION}_linux_amd64.tar.gz$" "ni_${VERSION}_checksums.txt" | sha256sum -c -
```

macOS:

```bash
VERSION="<published-version-without-v>"
grep " ni_${VERSION}_darwin_arm64.tar.gz$" "ni_${VERSION}_checksums.txt" | shasum -a 256 -c -
```

5. Extract and install:

```bash
mkdir -p "$HOME/.local/bin"
VERSION="<published-version-without-v>"
tar -xzf "ni_${VERSION}_darwin_arm64.tar.gz"
install -m 0755 ni "$HOME/.local/bin/ni"
PATH="$HOME/.local/bin:$PATH" ni --help
PATH="$HOME/.local/bin:$PATH" ni version
```

자신의 platform에 맞는 archive name을 사용한다. Windows에서는 PowerShell 또는
trusted unzip tool로 `.zip` archive를 풀고 `ni.exe`를 User PATH에 두거나 repository
root의 `install.ps1`를 사용한다.

## Script가 하는 일

- `linux`, `darwin`, Windows-compatible shell을 detect한다.
- `amd64` 또는 `arm64`를 detect한다.
- GoReleaser asset
  `ni_<version>_<os>_<arch>.tar.gz` 또는
  `ni_<version>_windows_amd64.zip`를 선택한다.
- 기본적으로 GitHub Releases에서 download한다.
- `ni_<version>_checksums.txt`를 download하고 가능하면 archive를 verify한다.
- `BINDIR`가 없으면 `~/.local/bin`에 install한다.
- `--update-path`가 passed되면 marked zsh/bash PATH block을 추가한다.
- `--uninstall`로 installed binary와 marked PATH block만 제거한다.
- Command-name help/version commands를 next steps로 출력한다.

`ni init`, `ni status`, `ni end`, `ni run`, shell commands, agents, queues,
runtime execution이 아니다.

## Test Release Validation

Repository validation은 network access와 Go 없이 installer를 test할 수 있도록
local fake release asset을 사용한다:

```bash
bash scripts/test-install-sh.sh
```

Future release에서는 public availability wording을 바꾸기 전에 real release
asset verification을 반복한다:

```bash
VERSION="0.5.0"
sh install.sh --dry-run --version "$VERSION"
BINDIR="$(mktemp -d)" sh install.sh --version "$VERSION"
```

v0.5.0 verification은 2026-06-02에 통과했다. Installer는
`Verified checksum for ni_0.5.0_darwin_arm64.tar.gz`를 출력했고 temporary
`BINDIR`에 binary를 install했으며, installed binary는 `ni version`에서 `0.5.0`을
반환했다. Global command-name verification은 이제 `bash scripts/install-check.sh`가
temporary install directory와 fresh shell PATH context로 cover한다.
