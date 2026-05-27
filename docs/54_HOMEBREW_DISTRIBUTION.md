# Homebrew Distribution Plan

This document plans Homebrew distribution for `ni` without claiming that
Homebrew installation works today.

Current status: Planned. There is no published Homebrew formula for `ni`, no
verified `brew install` path, and no package-manager release automation enabled
for Homebrew.

As of 2026-05-28, a direct repository check for
`https://github.com/Nam-Cheol/homebrew-tap.git` returned "Repository not found."
That means `Nam-Cheol/homebrew-tap` must be treated as next work, not as an
available tap.

## Decision

Use `Nam-Cheol/homebrew-tap` as the first planned Homebrew path, after the tap
repository exists and ownership is confirmed.

Do not use a repository-local Formula as the user-facing install path. A
repository-local formula may be useful later as a review fixture or temporary
test artifact, but it should not be documented as an install channel because it
does not match how Homebrew users normally discover and upgrade external CLI
tools.

Do not pursue an official Homebrew tap or Homebrew core submission yet. That is
a later option after release binaries are stable, checksums are published, and
there is enough usage signal to justify the extra maintenance surface.

## Availability

| Item | Status | Evidence |
| --- | --- | --- |
| Homebrew tap repository | Not found | `git ls-remote https://github.com/Nam-Cheol/homebrew-tap.git` returned repository not found |
| Homebrew formula | Not published | No tap exists and this repository does not ship an active Formula |
| `brew install` command | Not available | No formula has been published or validated |
| GoReleaser Homebrew publishing | Not configured | `.goreleaser.yaml` has build, archive, checksum, snapshot, and changelog config only |

## Required Sequence

Homebrew can move from planned to available only after all of these are true:

1. Release binaries exist in GitHub Releases for supported macOS architectures.
2. Checksums are published and match the release archives.
3. `Nam-Cheol/homebrew-tap` exists and is confirmed as the intended owner tap.
4. A `Formula/ni.rb` formula points to the official release assets and checksums.
5. The formula is validated with Homebrew tooling.
6. A real install command is tested from a clean Homebrew environment.
7. README install language is updated only after the tested command works.

Until then, docs must use planned language only.

## Tap Creation Steps

Because the tap repository does not exist yet, Homebrew status remains Planned.
Use these exact next steps when the owner is ready to create the tap:

1. Create `Nam-Cheol/homebrew-tap` as a public GitHub repository.
2. Clone the tap repository locally.
3. Create `Formula/ni.rb` in the tap repository, not in `ni-kernel`.
4. Use [Homebrew Formula Draft](71_HOMEBREW_FORMULA_DRAFT.md) as the starting
   point, replacing every placeholder checksum with checksums from the published
   GitHub Release.
5. Validate the formula from the tap checkout:

   ```bash
   brew audit --strict --online Formula/ni.rb
   brew install --build-from-source Formula/ni.rb
   ni --help
   ni version
   ```

6. Only after validation succeeds, document the tested install command as:

   ```bash
   brew install Nam-Cheol/tap/ni
   ```

Do not add a Homebrew badge, README install command, or package-manager
availability claim before that command works.

## Future GoReleaser Work

Do not add Homebrew publishing to `.goreleaser.yaml` until the tap exists and
the owner explicitly wants GoReleaser to update it.

When those conditions are met, a future change may add documented GoReleaser
Homebrew configuration that targets the external tap and keeps release
publishing as repository infrastructure. That change must be validated with:

```bash
goreleaser check
bash scripts/quality.sh
```

The GoReleaser config must not turn `ni` into a release runtime, task runner, or
package-manager automation command. It remains release infrastructure outside
`ni-kernel`.

## Next Work

- Create or confirm ownership of `Nam-Cheol/homebrew-tap`.
- Decide whether the tap will be updated manually first or through GoReleaser.
- Publish macOS release archives and checksums before any formula references
  them.
- Add and validate `Formula/ni.rb` in the tap.
- Test the exact Homebrew install command before changing public docs from
  planned to available.

## README Rule

README files may say Homebrew is planned. They must not include a `brew install`
command or imply package-manager availability until a formula exists and the
install path has been tested.
