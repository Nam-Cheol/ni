# Fresh User Install Verification

Date: 2026-05-28

Scope: Verify that a new user can install and run `ni` from the published
`v0.3.0` release without knowing or installing Go.

This verification covers the release binary path and curl installer path only.
It does not test Homebrew, Scoop, package-manager distribution, model skills,
Codex, Claude, downstream agents, or runtime execution behavior.

## Verification Script

Run:

```bash
bash scripts/fresh-install-check.sh
```

The script uses temporary directories only. It downloads the current platform's
published `v0.3.0` archive and `ni_0.3.0_checksums.txt`, verifies the archive
checksum, extracts the binary into a temporary install directory, and runs:

```bash
ni --help
ni version
ni init --dir <temporary-project> --profile prototype
ni status --dir <temporary-project>
```

It then downloads the public curl installer into a temporary directory, installs
the same `v0.3.0` binary into a temporary `BINDIR`, and repeats the same
help/version/init/status checks.

## Expected Output

The output should include:

```text
fresh-install-check: manual release binary path passed
fresh-install-check: curl installer path passed
fresh-install-check: ni --help, ni version, ni init, and ni status passed without Go
```

`ni status` is expected to print `BLOCKED` for the new temporary project because
`ni init` creates a draft planning workspace with open intent gaps. That blocked
result is successful fresh-user proof: the installed CLI can initialize a
project and run the deterministic readiness gate without requiring Go.

## Paths Verified

| Path | Status | Proof |
| --- | --- | --- |
| Manual release binary | Verified by script | Downloads the release archive and checksum file, verifies SHA-256, extracts, and runs the binary from a temporary directory. |
| Curl installer | Verified by script | Downloads `install.sh`, uses a temporary `BINDIR`, verifies checksum output, and runs the installed binary. |
| Homebrew | Not tested | Homebrew remains Planned and must not be marked Available by this check. |

## Result

A non-Go user can install `ni` from the verified `v0.3.0` release binary path or
the verified curl installer path, then run `ni --help`, `ni version`, `ni init`,
and `ni status`.

This verification does not publish a new release, push tags, execute downstream
work, call agents, or add runtime execution behavior.
