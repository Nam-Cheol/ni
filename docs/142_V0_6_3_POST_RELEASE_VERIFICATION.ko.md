# v0.6.3 Post-Release Verification

Status: `V0_6_3_RELEASE_EXECUTED_AND_VERIFIED`

## Release Metadata

| Item | Evidence |
| --- | --- |
| Tag | `v0.6.3` |
| Commit | `d048158a91f64888a71304ee1547ff6c4bbebe0e` |
| GitHub Release | https://github.com/Nam-Cheol/ni/releases/tag/v0.6.3 |
| Workflow run | `27244128964` |
| Published at | `2026-06-10T00:11:42Z` |

Release workflow는 성공했고, release는 draft가 아니며 prerelease도 아니다.

## Assets

`gh release view v0.6.3 --repo Nam-Cheol/ni --json tagName,name,isDraft,isPrerelease,publishedAt,url,assets`
결과:

| Asset | Status |
| --- | --- |
| `namba-intent_0.6.3_checksums.txt` | uploaded |
| `namba-intent_0.6.3_darwin_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.3_darwin_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.3_linux_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.3_linux_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.3_windows_amd64.zip` | uploaded |

Default `gh release view --repo Nam-Cheol/ni`는 `v0.6.3`를 반환하므로 latest
release view는 이 tag를 가리킨다.

## Hosted Install Smoke

Command shape:

```bash
tmp_root=$(mktemp -d /private/tmp/namba-intent-v0.6.3-verify.XXXXXX)
bin_dir="$tmp_root/bin"
workspace="$tmp_root/workspace"
mkdir -p "$bin_dir" "$workspace"
BINDIR="$bin_dir" sh install.sh --version 0.6.3
"$bin_dir/namba-intent" version
"$bin_dir/namba-intent" update --version 0.6.3
"$bin_dir/namba-intent" init --dir "$workspace" --yes
"$bin_dir/namba-intent" status --dir "$workspace" --proof --next-questions
BINDIR="$bin_dir" sh install.sh --uninstall
test ! -e "$bin_dir/namba-intent"
```

Observed result:

| Check | Result |
| --- | --- |
| Installer selected | `namba-intent_0.6.3_darwin_arm64.tar.gz` |
| Checksum verification | Passed |
| Installed binary version | `0.6.3` |
| `namba-intent update --version 0.6.3` | Guidance만 출력; download/install 없음 |
| `namba-intent init --dir ... --yes` | Planning docs와 `.ni` skeleton 생성 |
| `namba-intent status --proof --next-questions` | Expected first-run `BLOCKED` status 반환 |
| Uninstall | Temp binary 제거 |

Temp verification root는
`/private/tmp/namba-intent-v0.6.3-verify.CQziUL`였다.

## Workflow Warnings

Release workflow는 성공했지만 non-blocking GitHub Actions warning이 있었다.

- `actions/checkout@v4`, `actions/setup-go@v5`,
  `goreleaser/goreleaser-action@v6`에서 Node.js 20 deprecation notice가 나왔다.
- GoReleaser action은 implicit `latest` version이 latest v2 series로 resolve된다고
  warning했다.

이 warning은 artifact publication을 막지 않았지만, 이후 workflow-maintenance pass에서
정리해야 한다.

## Boundary Notes

- v0.6.3의 `namba-intent update`는 guidance-only이다. 명시적 update,
  verification, uninstall command를 출력하지만 file download, installer 실행,
  PATH 변경, planning state 수정은 하지 않는다.
- Windows release asset 존재는 published asset list로 확인했다. Windows real-host
  install execution은 Windows transcript가 있을 때까지 deferred이다.
<!-- ni-boundary-allow: explicit negative boundary statement. -->
- ni-kernel boundary는 유지된다. task runner, SPEC runner, shell adapter, Codex
  adapter, queue, PR automation, release automation, downstream execution layer는
  추가하지 않았다.
