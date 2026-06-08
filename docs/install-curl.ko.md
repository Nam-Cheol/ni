# Curl Installer

`install.sh`는 Go 없이 released Namba Intent binary를 설치하는 release asset
infrastructure이다. Current main은 `namba-intent_<version>` release assets를
선택하고, `namba-intent`를 local bin directory에 복사하고 next steps를 출력한다.
Model skills를 install하지 않으며 downstream work를 실행하지 않는다. 명시적으로 opt
in하면 reversible zsh/bash PATH block을 추가할 수 있다.

Status: verified v0.6.0 macOS `namba-intent` path는 Available이다. Verified
v0.5.1 GitHub Release assets는 historical `ni_<version>` names를 사용하며 v0.5.1
post-release verification record에 계속 기록되어 있다.

## 더 안전한 Script 경로

검증된 v0.6.0 macOS install:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --update-path --version 0.6.0
```

Inspect-first 또는 reproducible pinned install을 원할 때는 local install 전에
installer를 먼저 download한다:

```bash
VERSION="0.6.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --update-path --version "$VERSION"
```

기본 install 위치는 `~/.local/bin/namba-intent`다. 다른 위치는 `BINDIR`로 지정한다:

```bash
BINDIR="$HOME/bin" sh install.sh --dry-run --version "$VERSION"
```

`--version`을 생략하면 installer는 GitHub에서 latest release tag를 확인한다.
Verified Namba Intent release path를 재현하려면 `--version 0.6.0`을 사용한다.
Historical script 또는
[v0.5.1 Post-Release Verification](132_V0_5_1_POST_RELEASE_VERIFICATION.ko.md)이
cover한 evidence를 사용할 때만 `VERSION="0.5.1"`을 고정한다. Current-main
`install.sh --version 0.5.1`은 published되지 않은 `namba-intent_0.5.1...` assets를
선택한다. Current dry-run output은 `--version` 없이 latest를 resolve하지 않으므로
dry-run example은 README primary path가 아니라 pinned path에 둔다. Global command를
help 또는 version command로 확인한다:

```bash
namba-intent --help
namba-intent version
```

Global command verification 뒤에는 project directory에서 planning workspace를
시작한다:

```bash
mkdir my-project
cd my-project
namba-intent init .
```

Binary와, 추가했다면 ni-managed PATH block을 uninstall한다:

```bash
curl -fsSL https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh | sh -s -- --uninstall
```

Downloaded script에서 설치했거나 custom `BINDIR`를 사용했다면 같은 `BINDIR`를
전달해 local script에서 uninstall한다:

```bash
BINDIR="$HOME/.local/bin" sh install.sh --uninstall
```

## Manual Verification Path

Manual install은 trust step을 모두 보이게 유지한다.

1. Release page를 연다:
   <https://github.com/Nam-Cheol/ni/releases>
2. Release에 assets가 있는지 확인한 뒤 platform에 맞는 archive를 고른다. v0.6.0
   이후 Namba Intent release에서는 `namba-intent_<version>` names를 사용한다.
   Historical v0.5.1 evidence에는 v0.5.1 release page의 `ni_<version>` names를
   사용한다.

| Platform | Architecture | Archive |
| --- | --- | --- |
| Linux | amd64 | `namba-intent_<version>_linux_amd64.tar.gz` |
| Linux | arm64 | `namba-intent_<version>_linux_arm64.tar.gz` |
| macOS | amd64 | `namba-intent_<version>_darwin_amd64.tar.gz` |
| macOS | arm64 | `namba-intent_<version>_darwin_arm64.tar.gz` |
| Windows | amd64 | `namba-intent_<version>_windows_amd64.zip` |

3. 같은 release에서 archive와 `namba-intent_<version>_checksums.txt`를 download한다.
4. Archive checksum을 verify한다.

Linux:

```bash
VERSION="<published-version-without-v>"
grep " namba-intent_${VERSION}_linux_amd64.tar.gz$" "namba-intent_${VERSION}_checksums.txt" | sha256sum -c -
```

macOS:

```bash
VERSION="<published-version-without-v>"
grep " namba-intent_${VERSION}_darwin_arm64.tar.gz$" "namba-intent_${VERSION}_checksums.txt" | shasum -a 256 -c -
```

5. Extract and install:

```bash
mkdir -p "$HOME/.local/bin"
VERSION="<published-version-without-v>"
tar -xzf "namba-intent_${VERSION}_darwin_arm64.tar.gz"
install -m 0755 namba-intent "$HOME/.local/bin/namba-intent"
PATH="$HOME/.local/bin:$PATH" namba-intent --help
PATH="$HOME/.local/bin:$PATH" namba-intent version
```

자신의 platform에 맞는 archive name을 사용한다. Windows에서는 PowerShell 또는
trusted unzip tool로 `.zip` archive를 풀고 `namba-intent.exe`를 User PATH에 두거나
repository root의 `install.ps1`를 사용한다. `ni -> New-Item` PowerShell alias
cleanup은 legacy v0.5.x guidance이며 `namba-intent.exe`에는 필요하지 않다.
Command-name resolution이 이상하면 `Get-Command namba-intent -All`로 진단한다.

## Script가 하는 일

- `linux`, `darwin`, Windows-compatible shell을 detect한다.
- `amd64` 또는 `arm64`를 detect한다.
- GoReleaser asset
  `namba-intent_<version>_<os>_<arch>.tar.gz` 또는
  `namba-intent_<version>_windows_amd64.zip`를 선택한다.
- 기본적으로 GitHub Releases에서 download한다.
- `namba-intent_<version>_checksums.txt`를 download하고 가능하면 archive를 verify한다.
- `BINDIR`가 없으면 `~/.local/bin`에 install한다.
- `--update-path`가 passed되면 marked zsh/bash PATH block을 추가한다.
- `--uninstall`로 installed binary와 marked PATH block만 제거한다.
- Command-name help/version commands를 next steps로 출력한다.

`namba-intent init`, `namba-intent status`, `namba-intent end`, `namba-intent run`,
shell commands, agents, queues, runtime execution이 아니다.

## Test Release Validation

Repository validation은 network access와 Go 없이 installer를 test할 수 있도록
local fake release asset을 사용한다:

```bash
bash scripts/test-install-sh.sh
```

v0.6.0 macOS installer path는 real release verification을 거쳤다. 이후 release에서는
public availability wording을 바꾸기 전에 그 verification을 반복한다:

```bash
VERSION="0.6.0"
sh install.sh --dry-run --version "$VERSION"
BINDIR="$(mktemp -d)" sh install.sh --version "$VERSION"
```

Historical v0.5.1 verification은 2026-06-08에 old `ni` release assets 기준으로
통과했다. v0.6.0 macOS verification은
[`140_V0_6_0_POST_RELEASE_VERIFICATION.ko.md`](140_V0_6_0_POST_RELEASE_VERIFICATION.ko.md)에
기록되어 있다. Current tree command-name verification은 `bash scripts/install-check.sh`가
temporary install directory와 fresh shell PATH context로 cover한다.
