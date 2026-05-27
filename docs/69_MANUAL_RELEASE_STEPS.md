# Manual Release Steps

These steps prepare and publish a public GitHub Release manually. They are not
`ni-kernel` behavior, and they must not be turned into release automation inside
the product.

Planned version: `v0.3.0`

## Before Tagging

1. Verify the git tree is clean:

   ```bash
   git status --short
   ```

2. Run the non-publishing release dry run:

   ```bash
   bash scripts/release-dry-run.sh
   ```

3. Run the normal quality gate:

   ```bash
   bash scripts/quality.sh
   go test ./...
   ```

4. Re-read the release notes at
   [docs/68_RELEASE_NOTES_v0.3.0.md](68_RELEASE_NOTES_v0.3.0.md) and confirm
   they do not claim hosted binaries, curl installer availability, package
   manager distribution, runtime execution, adapters, queues, or PR/release
   automation inside `ni-kernel`.

## Tag And Publish Assets

1. Create an annotated tag locally:

   ```bash
   git tag -a v0.3.0 -m "ni v0.3.0"
   ```

2. Push only the tag:

   ```bash
   git push origin v0.3.0
   ```

3. Wait for the GitHub Actions release workflow to finish.

4. Confirm the GitHub Release exists and contains all expected archives:

   - `ni_<version>_linux_amd64.tar.gz`
   - `ni_<version>_linux_arm64.tar.gz`
   - `ni_<version>_darwin_amd64.tar.gz`
   - `ni_<version>_darwin_arm64.tar.gz`
   - `ni_<version>_windows_amd64.zip`
   - `ni_<version>_checksums.txt`

5. Confirm checksums match the uploaded archives.

6. Only after assets and checksums exist, update README install status from
   release-gated/planned wording to available wording for the release binary
   path. Do not mark curl installer availability as available unless the
   installer has been verified against the published release assets.

## Stop Conditions

Stop the release if any check fails, if the tag points at the wrong commit, if
assets are missing, if checksums do not match, or if release notes imply that
`ni` is a task runner, SPEC runner, multi-agent execution layer, Codex exec
adapter, shell adapter, queue, PR automation system, release automation system,
or downstream execution runtime.
