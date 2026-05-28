# v0.3.0 Curl Installer 검증

Date: 2026-05-28

Scope: Curl installer를 Available로 표시하기 전에 `install.sh`를 실제 `v0.3.0`
GitHub Release assets에 대해 검증한다. 이 검증은 Homebrew, Scoop 또는 package
manager distribution을 Available로 표시하지 않는다.

검증 대상 release:

`https://github.com/Nam-Cheol/ni/releases/tag/v0.3.0`

## 실행한 Commands

Dry-run:

```bash
sh install.sh --dry-run --version v0.3.0
```

Real temporary install:

```bash
BINDIR="/var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi" sh install.sh --version v0.3.0
"/var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni" --help
"/var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni" version
```

Platform override dry-runs:

```bash
NI_INSTALL_OS=linux NI_INSTALL_ARCH=amd64 sh install.sh --dry-run --version v0.3.0
NI_INSTALL_OS=linux NI_INSTALL_ARCH=arm64 sh install.sh --dry-run --version v0.3.0
NI_INSTALL_OS=darwin NI_INSTALL_ARCH=amd64 sh install.sh --dry-run --version v0.3.0
NI_INSTALL_OS=darwin NI_INSTALL_ARCH=arm64 sh install.sh --dry-run --version v0.3.0
NI_INSTALL_OS=windows NI_INSTALL_ARCH=amd64 sh install.sh --dry-run --version v0.3.0
```

Checksum mismatch safety test:

```bash
bash scripts/test-install-sh.sh
```

Mismatch test는 local fake release asset과 intentionally bad local checksum file을
사용한다. Published GitHub Release를 mutate하거나 spoof하지 않는다.

## Dry-Run Output

```text
ni installer
  repository: Nam-Cheol/ni
  platform:   darwin/arm64
  asset:      ni_0.3.0_darwin_arm64.tar.gz
  checksums:  ni_0.3.0_checksums.txt
  install to: /Users/namba/.local/bin/ni
  mode:       dry-run

Would download:
  https://github.com/Nam-Cheol/ni/releases/download/v0.3.0/ni_0.3.0_darwin_arm64.tar.gz
  https://github.com/Nam-Cheol/ni/releases/download/v0.3.0/ni_0.3.0_checksums.txt
```

## Platform Override Results

| Override | Selected asset |
| --- | --- |
| `linux/amd64` | `ni_0.3.0_linux_amd64.tar.gz` |
| `linux/arm64` | `ni_0.3.0_linux_arm64.tar.gz` |
| `darwin/amd64` | `ni_0.3.0_darwin_amd64.tar.gz` |
| `darwin/arm64` | `ni_0.3.0_darwin_arm64.tar.gz` |
| `windows/amd64` | `ni_0.3.0_windows_amd64.zip` |

모든 override는 같은 release URL의 `ni_0.3.0_checksums.txt`를 선택했다.

## Real Temporary Install Output

```text
ni installer
  repository: Nam-Cheol/ni
  platform:   darwin/arm64
  asset:      ni_0.3.0_darwin_arm64.tar.gz
  checksums:  ni_0.3.0_checksums.txt
  install to: /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni
Downloading ni_0.3.0_darwin_arm64.tar.gz
Verified checksum for ni_0.3.0_darwin_arm64.tar.gz
Installed ni to /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni

Next steps:
  1. Ensure /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi is on your PATH.
  2. Check the installed CLI:
     /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni --help
     /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni version

The installer does not install model skills or run downstream work.
```

Installed binary checks:

```text
$ /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni --help
ni is a project intent compiler.

$ /var/folders/p5/dwdgcnk918g9kc4hmctqs3d40000gn/T/tmp.6qBvKt5hGi/ni version
0.3.0
```

## Checksum Behavior

- Real install은 v0.3.0 release에서 `ni_0.3.0_checksums.txt`를 download했다.
- Installer는 `ni_0.3.0_darwin_arm64.tar.gz` checksum entry를 찾았다.
- Installer는 `Verified checksum for ni_0.3.0_darwin_arm64.tar.gz`를 출력했다.
- `bash scripts/test-install-sh.sh`는 local fake release와 bad checksum file을
  사용해 checksum mismatch failure를 확인했다.

## Result

Curl installer status: verified v0.3.0 release assets에 대해 Available.

Homebrew status: Planned.

이 검증은 package-manager availability를 claim하지 않는다.
