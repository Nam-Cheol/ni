# Homebrew Decision

Date: 2026-05-29

Decision: **B. Defer Homebrew tap implementation to v0.5.**

Current Homebrew status: Planned.

## Reviewed Inputs

- [Homebrew Tap Plan](72_HOMEBREW_TAP_PLAN.md)
- [Homebrew Tap Plan, Korean companion](72_HOMEBREW_TAP_PLAN.ko.md)
- [Homebrew Tap Setup](74_HOMEBREW_TAP_SETUP.md)
- v0.3.0 release binary status: Available after release asset and checksum
  verification.
- v0.3.0 curl installer status: Available after verification against the real
  release assets.

## Rationale

Homebrew should not be implemented in this task because the required external
tap work has not happened yet. The intended tap is still
`Nam-Cheol/homebrew-tap`, but Homebrew must remain Planned until the tap exists,
`Formula/ni.rb` is published with real checksums, Homebrew audit passes, and
`brew install Nam-Cheol/tap/ni` is tested from a clean Homebrew environment.

Deferring the implementation keeps the current release story factual: source,
local binary, release binary, and curl installer paths are Available, while
package-manager distribution is still unverified. It also keeps package
publishing outside `ni-kernel` runtime behavior.

## What This Decision Does

- Keeps Homebrew status as Planned.
- Keeps the existing tap target as `Nam-Cheol/homebrew-tap`.
- Keeps `docs/74_HOMEBREW_TAP_SETUP.md` as the setup procedure for the later
  implementation task.
- Schedules Homebrew tap implementation for v0.5 as distribution
  infrastructure, not kernel runtime behavior.

## What This Decision Does Not Do

- Does not create or modify a Homebrew tap.
- Does not add `brew install` instructions to public install docs.
- Does not mark Homebrew Available.
- Does not add package-manager automation, release publishing, tags, or runtime
  execution behavior.

## Availability Gate

Homebrew may become Available only after all of these are true:

1. `Nam-Cheol/homebrew-tap` exists and is confirmed as the intended public tap.
2. `Formula/ni.rb` is published in that tap.
3. Formula URLs point to official release assets.
4. Formula checksums match the published release checksum source.
5. `brew audit --strict --online Formula/ni.rb` passes.
6. `brew install --build-from-source Formula/ni.rb` passes.
7. `brew install Nam-Cheol/tap/ni` passes from a clean Homebrew environment.
8. `ni --help` and `ni version` pass for the Homebrew-installed binary.

Until then, README, install docs, release docs, and launch material must keep
Homebrew as Planned.

