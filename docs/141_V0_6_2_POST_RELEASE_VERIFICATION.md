# v0.6.2 Post-Release Verification

Status: `V0_6_2_RELEASE_EXECUTED_AND_VERIFIED`

## Release Metadata

| Item | Evidence |
| --- | --- |
| Tag | `v0.6.2` |
| Commit | `a77cf9393356e4a93fdfc37ccf14250ce104741a` |
| GitHub Release | https://github.com/Nam-Cheol/ni/releases/tag/v0.6.2 |
| Workflow run | `27181080761` |
| Published at | `2026-06-09T03:03:13Z` |

The release workflow completed successfully. The release is not a draft and not
a prerelease.

## Assets

`gh release view v0.6.2 --repo Nam-Cheol/ni --json tagName,name,isDraft,isPrerelease,publishedAt,url,assets`
reported:

| Asset | Status |
| --- | --- |
| `namba-intent_0.6.2_checksums.txt` | uploaded |
| `namba-intent_0.6.2_darwin_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.2_darwin_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.2_linux_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.2_linux_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.2_windows_amd64.zip` | uploaded |

Default `gh release view --repo Nam-Cheol/ni` returned `v0.6.2`, so the latest
release view points at this tag.

## Hosted Install Smoke

Command shape:

```bash
tmp_root=$(mktemp -d /private/tmp/namba-intent-v0.6.2-verify.XXXXXX)
bin_dir="$tmp_root/bin"
workspace="$tmp_root/workspace"
mkdir -p "$bin_dir" "$workspace"
BINDIR="$bin_dir" sh install.sh --version 0.6.2
"$bin_dir/namba-intent" version
"$bin_dir/namba-intent" update --version 0.6.2
"$bin_dir/namba-intent" init --dir "$workspace" --yes
"$bin_dir/namba-intent" status --dir "$workspace" --proof --next-questions
BINDIR="$bin_dir" sh install.sh --uninstall
test ! -e "$bin_dir/namba-intent"
```

Observed result:

| Check | Result |
| --- | --- |
| Installer selected | `namba-intent_0.6.2_darwin_arm64.tar.gz` |
| Checksum verification | Passed |
| Installed binary version | `0.6.2` |
| `namba-intent update --version 0.6.2` | Printed guidance only; no download or install |
| `namba-intent init --dir ... --yes` | Created planning docs and `.ni` skeleton |
| `namba-intent status --proof --next-questions` | Returned expected first-run `BLOCKED` status |
| Uninstall | Removed the temp binary |

The temp verification root was
`/private/tmp/namba-intent-v0.6.2-verify.c0Lb2S`.

## Boundary Notes

- `namba-intent update` is guidance-only in v0.6.2. It prints explicit update,
  verification, and uninstall commands, but it does not download files, run an
  installer, change PATH, or modify planning state.
- Windows release asset existence is verified through the published asset list.
  Windows real-host install execution remains deferred until a Windows
  transcript exists.
<!-- ni-boundary-allow: explicit negative boundary statement. -->
- The release preserves the ni-kernel boundary: no task runner, SPEC runner,
  shell adapter, Codex adapter, queue, PR automation, release automation, or
  downstream execution layer was added.
