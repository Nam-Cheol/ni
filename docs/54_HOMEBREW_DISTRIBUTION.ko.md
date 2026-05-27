# Homebrew Distribution Plan

이 문서는 `ni`의 Homebrew distribution을 계획한다. 아직 Homebrew installation이
동작한다고 claim하지 않는다.

현재 상태: planned. `ni`용 published Homebrew formula는 없고, verified
`brew install` path도 없다. Homebrew package-manager release automation은 아니다.

2026-05-27 기준으로 `https://github.com/Nam-Cheol/homebrew-tap.git`에 대한
direct repository check는 "Repository not found."를 반환했다. 따라서
`Nam-Cheol/homebrew-tap`은 available tap이 아니라 next work로 다뤄야 한다.

## Decision

첫 Homebrew path는 `Nam-Cheol/homebrew-tap`을 사용한다. 단, tap repository가
실제로 존재하고 ownership이 확인된 뒤에만 그렇게 한다.

Repository-local Formula를 user-facing install path로 쓰지 않는다.
Repository-local formula는 나중에 review fixture나 temporary test artifact로는
쓸 수 있지만, Homebrew users가 external CLI tools를 discover하고 upgrade하는
일반 경로와 맞지 않으므로 install channel로 문서화하지 않는다.

Official Homebrew tap 또는 Homebrew core submission은 아직 추진하지 않는다.
Release binaries가 안정화되고, checksums가 publish되고, 추가 maintenance
surface를 정당화할 usage signal이 생긴 뒤의 later option이다.

## Availability

| Item | Status | Evidence |
| --- | --- | --- |
| Homebrew tap repository | Not found | `git ls-remote https://github.com/Nam-Cheol/homebrew-tap.git`가 repository not found를 반환했다 |
| Homebrew formula | Not published | Tap이 없고 이 repository도 active Formula를 ship하지 않는다 |
| `brew install` command | Not available | Formula가 publish 또는 validate되지 않았다 |
| GoReleaser Homebrew publishing | Not configured | `.goreleaser.yaml`에는 build, archive, checksum, snapshot, changelog config만 있다 |

## Required Sequence

Homebrew는 다음이 모두 true일 때만 planned에서 available로 바뀔 수 있다:

1. Supported macOS architectures용 release binaries가 GitHub Releases에 존재한다.
2. Checksums가 publish되고 release archives와 일치한다.
3. `Nam-Cheol/homebrew-tap`이 존재하며 intended owner tap으로 확인된다.
4. `Formula/ni.rb` formula가 official release assets와 checksums를 가리킨다.
5. Formula가 Homebrew tooling으로 validate된다.
6. Clean Homebrew environment에서 실제 install command가 test된다.
7. Tested command가 동작한 뒤에만 README install language를 update한다.

그 전까지 docs는 planned language만 사용해야 한다.

## Future GoReleaser Work

Tap이 존재하고 owner가 GoReleaser update를 명시적으로 원하기 전까지
`.goreleaser.yaml`에 Homebrew publishing을 추가하지 않는다.

그 조건이 충족되면 future change에서 external tap을 target으로 하는 documented
GoReleaser Homebrew configuration을 추가할 수 있다. Release publishing은 계속
repository infrastructure로 남아야 한다. 그 change는 다음으로 validate해야 한다:

```bash
goreleaser check
bash scripts/quality.sh
```

GoReleaser config는 `ni`를 release runtime, task runner, package-manager
automation command로 바꾸면 안 된다. 이는 `ni-kernel` 밖의 release
infrastructure다.

## Next Work

- `Nam-Cheol/homebrew-tap`을 create하거나 ownership을 confirm한다.
- Tap을 먼저 manual로 update할지 GoReleaser로 update할지 결정한다.
- Formula가 참조하기 전에 macOS release archives와 checksums를 publish한다.
- Tap에 `Formula/ni.rb`를 add and validate한다.
- Public docs를 planned에서 available로 바꾸기 전에 정확한 Homebrew install
  command를 test한다.

## README Rule

README files는 Homebrew가 planned라고 말할 수 있다. Formula가 존재하고 install
path가 test되기 전까지 `brew install` command를 포함하거나 package-manager
availability를 imply하면 안 된다.
