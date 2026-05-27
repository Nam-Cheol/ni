# Curl Installer

`install.sh`는 Go 없이 released `ni` binary를 설치한다. 이것은 release asset
infrastructure일 뿐이다. Archive를 download하고, release가 checksum을 제공하면
verify한 뒤, `ni`를 local bin directory에 복사하고 next steps를 출력한다. Model
skills를 install하지 않으며 downstream work를 실행하지 않는다.

Status: Release-gated. GitHub Release가 expected archives와 checksum file을
publish하고, `install.sh`가 그 real assets에 대해 verified되기 전까지 public
available install path로 설명하면 안 된다.

## 더 안전한 Script 경로

Release gate가 충족된 뒤 installer를 먼저 download하고 inspect한다:

```bash
VERSION="<published-version-without-v>"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --version "$VERSION"
```

기본 install 위치는 `~/.local/bin/ni`다. 다른 위치는 `BINDIR`로 지정한다:

```bash
BINDIR="$HOME/bin" sh install.sh --dry-run --version "$VERSION"
```

`--version`을 생략하면 installer는 GitHub에서 latest release tag를 확인한다. 이
경로는 적어도 하나의 published release가 필요하다. 설치된 CLI는 help 또는
version command로만 확인한다:

```bash
~/.local/bin/ni --help
~/.local/bin/ni version
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
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

자신의 platform에 맞는 archive name을 사용한다. Windows에서는 PowerShell 또는
trusted unzip tool로 `.zip` archive를 풀고 `ni.exe`를 `PATH`에 둔다.

## Script가 하는 일

- `linux`, `darwin`, Windows-compatible shell을 detect한다.
- `amd64` 또는 `arm64`를 detect한다.
- GoReleaser asset
  `ni_<version>_<os>_<arch>.tar.gz` 또는
  `ni_<version>_windows_amd64.zip`를 선택한다.
- 기본적으로 GitHub Releases에서 download한다.
- `ni_<version>_checksums.txt`를 download하고 가능하면 archive를 verify한다.
- `BINDIR`가 없으면 `~/.local/bin`에 install한다.
- Help/version commands를 next steps로 출력한다.

`ni init`, `ni status`, `ni end`, `ni run`, shell commands, agents, queues,
runtime execution이 아니다.

## Test Release Validation

Repository validation은 network access와 Go 없이 installer를 test할 수 있도록
local fake release asset을 사용한다:

```bash
bash scripts/test-install-sh.sh
```

Public availability wording을 바꾸기 전에는 real release asset도 verify한다:

```bash
VERSION="<published-version-without-v>"
sh install.sh --dry-run --version "$VERSION"
BINDIR="$(mktemp -d)" sh install.sh --version "$VERSION"
```
