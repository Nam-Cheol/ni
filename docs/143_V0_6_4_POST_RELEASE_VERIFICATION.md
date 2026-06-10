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

The release is not a draft and not a prerelease.

The release workflow passed Go tests, quality checks, and the release readiness
check. GoReleaser built the archive matrix. GoReleaser publish failed with `401 Requires authentication` while creating the GitHub Release. Assets were regenerated locally from the tagged commit with the same OS/arch matrix, archive names, and version linker flag, then published with authenticated `gh release create`.

## Assets

`gh release view v0.6.4 --repo Nam-Cheol/ni --json tagName,name,isDraft,isPrerelease,publishedAt,url,assets`
reported:

| Asset | Status |
| --- | --- |
| `namba-intent_0.6.4_checksums.txt` | uploaded |
| `namba-intent_0.6.4_darwin_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.4_darwin_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.4_linux_amd64.tar.gz` | uploaded |
| `namba-intent_0.6.4_linux_arm64.tar.gz` | uploaded |
| `namba-intent_0.6.4_windows_amd64.zip` | uploaded |

Default `gh release view --repo Nam-Cheol/ni` returned `v0.6.4`, so the latest
release view points at this tag.

Hosted artifact verification downloaded all published assets to
`/private/tmp/namba-intent-v0.6.4-hosted-verify` and ran
`shasum -a 256 -c namba-intent_0.6.4_checksums.txt`. All archives passed.

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
| `namba-intent update --version 0.6.4` | Printed guidance only; no download or install |
| `namba-intent init --dir ... --yes` | Created planning docs and `.ni` skeleton |
| `namba-intent status --proof --next-questions` | Returned expected first-run `BLOCKED` status |
| Uninstall | Removed the temp binary |

The temp verification root was
`/private/tmp/namba-intent-v0.6.4-verify.XFnSmO`.

## Workflow Notes

- The workflow failure was isolated to GitHub Release publish authentication.
  Tests, quality, release readiness, local archive generation, hosted checksum
  verification, and hosted installer smoke all passed.
- `gh auth status` showed local `repo` and `workflow` scopes for manual
  publication.
- The release workflow still needs a later maintenance pass so GoReleaser
  publish does not require manual recovery.

## Boundary Notes

- `namba-intent update` remains guidance-only in v0.6.4. It prints explicit
  update, verification, and uninstall commands, but it does not download files,
  run an installer, change PATH, or modify planning state.
- Windows release asset existence is verified through the published asset list.
  Windows real-host install execution remains deferred until a Windows
  transcript exists.
<!-- ni-boundary-allow: explicit negative boundary statement. -->
- The release preserves the ni-kernel boundary: no task runner, SPEC runner,
  shell adapter, Codex adapter, queue, PR automation, release automation, or
  downstream execution layer was added.
