# v0.6.4 Post-Release Verification

Status: `V0_6_4_RELEASE_EXECUTED_AND_VERIFIED`

## Release Metadata

| Item | Evidence |
| --- | --- |
| Tag | `v0.6.4` |
| Commit | `760fe660fc0b6887d3a2b7f3821a631915e4545f` |
| GitHub Release | https://github.com/Nam-Cheol/ni/releases/tag/v0.6.4 |
| Workflow run | `27285705816` |
| Published at | `2026-06-10T15:08:48Z` |

Release는 draft가 아니며 prerelease도 아니다.

Release workflow는 Go tests, quality checks, release readiness check를
통과했다. GoReleaser는 archive matrix를 build했지만 GitHub Release 생성 단계에서
GoReleaser publish는 `401 Requires authentication`으로 실패했다. Assets는 tagged commit에서 local로 재생성했으며, 같은 OS/arch matrix, archive names, version linker flag를 사용한 뒤 authenticated `gh release create`로 publish했다.

## Assets

`gh release view v0.6.4 --repo Nam-Cheol/ni --json tagName,name,isDraft,isPrerelease,publishedAt,url,assets`
결과:

| Asset | Status |
| --- | --- |
| `namba-intent_0.6.4_checksums.txt` | uploaded |
| `namba-intent_0.6.4_darwin_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.4_darwin_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.4_linux_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.4_linux_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.4_windows_amd64.zip` | uploaded |

Default `gh release view --repo Nam-Cheol/ni`는 `v0.6.4`를 반환하므로 latest
release view는 이 tag를 가리킨다.

Hosted artifact verification은 모든 published asset을
`/private/tmp/namba-intent-v0.6.4-hosted-verify`로 다운로드한 뒤
`shasum -a 256 -c namba-intent_0.6.4_checksums.txt`를 실행했다. 모든 archive가
통과했다.

## Hosted Install Smoke

Command shape:

```bash
tmp_root=$(mktemp -d /private/tmp/namba-intent-v0.6.4-verify.XXXXXX)
bin_dir="$tmp_root/bin"
workspace="$tmp_root/workspace"
mkdir -p "$bin_dir" "$workspace"
BINDIR="$bin_dir" sh install.sh --version 0.6.4
"$bin_dir/namba-intent" version
"$bin_dir/namba-intent" update --version 0.6.4
"$bin_dir/namba-intent" init --dir "$workspace" --yes
"$bin_dir/namba-intent" status --dir "$workspace" --proof --next-questions
BINDIR="$bin_dir" sh install.sh --uninstall
test ! -e "$bin_dir/namba-intent"
```

Observed result:

| Check | Result |
| --- | --- |
| Installer selected | `namba-intent_0.6.4_darwin_arm64.tar.gz` |
| Checksum verification | Passed |
| Installed binary version | `0.6.4` |
| `namba-intent update --version 0.6.4` | Guidance만 출력; download/install 없음 |
| `namba-intent init --dir ... --yes` | Planning docs와 `.ni` skeleton 생성 |
| `namba-intent status --proof --next-questions` | Expected first-run `BLOCKED` status 반환 |
| Uninstall | Temp binary 제거 |

Temp verification root는
`/private/tmp/namba-intent-v0.6.4-verify.XFnSmO`였다.

## Workflow Notes

- Workflow failure는 GitHub Release publish authentication 단계로 isolate되었다.
  Tests, quality, release readiness, local archive generation, hosted checksum
  verification, hosted installer smoke는 모두 통과했다.
- `gh auth status`는 manual publication에 필요한 local `repo`와 `workflow` scope를
  보여줬다.
- GoReleaser publish가 manual recovery 없이 통과하도록 release workflow는 이후
  maintenance pass가 필요하다.

## Boundary Notes

- v0.6.4의 `namba-intent update`는 guidance-only이다. 명시적 update,
  verification, uninstall command를 출력하지만 file download, installer 실행,
  PATH 변경, planning state 수정은 하지 않는다.
- Windows release asset 존재는 published asset list로 확인했다. Windows real-host
  install execution은 Windows transcript가 있을 때까지 deferred이다.
<!-- ni-boundary-allow: explicit negative boundary statement. -->
- ni-kernel boundary는 유지된다. task runner, SPEC runner, shell adapter, Codex
  adapter, queue, PR automation, release automation, downstream execution layer는
  추가하지 않았다.
