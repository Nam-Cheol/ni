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

The release workflow completed successfully. The release is not a draft and not
a prerelease.

## Assets

`gh release view v0.6.3 --repo Nam-Cheol/ni --json tagName,name,isDraft,isPrerelease,publishedAt,url,assets`
reported:

| Asset | Status |
| --- | --- |
| `namba-intent_0.6.3_checksums.txt` | uploaded |
| `namba-intent_0.6.3_darwin_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.3_darwin_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.3_linux_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.3_linux_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.3_windows_amd64.zip` | uploaded |

Default `gh release view --repo Nam-Cheol/ni` returned `v0.6.3`, so the latest
release view points at this tag.

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
| `namba-intent update --version 0.6.3` | Printed guidance only; no download or install |
| `namba-intent init --dir ... --yes` | Created planning docs and `.ni` skeleton |
| `namba-intent status --proof --next-questions` | Returned expected first-run `BLOCKED` status |
| Uninstall | Removed the temp binary |

The temp verification root was
`/private/tmp/namba-intent-v0.6.3-verify.CQziUL`.

## Workflow Warnings

The release workflow succeeded with non-blocking GitHub Actions warnings:

- `actions/checkout@v4`, `actions/setup-go@v5`, and
  `goreleaser/goreleaser-action@v6` emitted Node.js 20 deprecation notices.
- The GoReleaser action warned that its implicit `latest` version resolves to
  the latest v2 series.

These warnings did not block artifact publication, but they should be handled in
a later workflow-maintenance pass.

## Boundary Notes

- `namba-intent update` is guidance-only in v0.6.3. It prints explicit update,
  verification, and uninstall commands, but it does not download files, run an
  installer, change PATH, or modify planning state.
- Windows release asset existence is verified through the published asset list.
  Windows real-host install execution remains deferred until a Windows
  transcript exists.
<!-- ni-boundary-allow: explicit negative boundary statement. -->
- The release preserves the ni-kernel boundary: no task runner, SPEC runner,
  shell adapter, Codex adapter, queue, PR automation, release automation, or
  downstream execution layer was added.
