# Homebrew decision

Date: 2026-05-29

## Current distribution state

- Release binary: Available for the verified v0.4.0 GitHub Release archives and
  checksums.
- Curl installer: Available for the verified v0.4.0 release assets.
- Homebrew: Planned. No tap or formula is published or tested.
- Model workspace packs: Experimental as a broad product path; source packs,
  manual copy, zip packaging, and the Claude target-directory dry-run path are
  available, but global host discovery is unverified.
- No-terminal method: Experimental / assisted. It can support drafting, but
  deterministic readiness, lock, hash, and prompt claims still require CLI
  proof from a trusted runner.

## Decision

Choose **B. Defer Homebrew to v0.5.**

Homebrew should remain Planned for v0.4.1 and should be implemented no earlier
than v0.5 as distribution infrastructure.

## Rationale

Homebrew is useful, but it is not the next adoption bottleneck. The release
binary and curl installer already serve users who want to try `ni` without Go,
and both paths have verification evidence. Homebrew would improve convenience
for macOS and developer users who prefer package-manager installation, but it
does not change the core product experience of planning, readiness proof,
locking, or prompt compilation.

The blocking work for Homebrew is external and operational: the tap repository
is not available, the formula is not published, and the clean `brew install`
path has not been tested. Implementing it now would compete with higher-impact
v0.4.1 stabilization work: conversation authoring UX, readiness proof clarity,
model workspace pack guidance, no-terminal proof capture, and benchmark
evidence.

Because `ni` is a pre-runtime Project Intent Compiler, Homebrew must stay
outside `ni-kernel` behavior. Package distribution should help users obtain the
CLI; it must not become release automation, execution state, adapter behavior,
or a reason to expand `ni run` beyond prompt compilation.

## User impact

Homebrew would benefit users who already trust Homebrew, want upgrades through
a package manager, or expect `brew install Nam-Cheol/tap/ni` for a macOS CLI.
It would reduce install friction for that audience after the tap and formula
are verified.

Users who are willing to inspect a script or manually verify checksums are
already served by the curl installer and release binary paths. Developers,
evaluators, and contributors are already served by source and local binary
paths. Users who need help authoring plans are better served right now by
improvements to conversation authoring, model workspace packs, no-terminal
assisted proof capture, and benchmark evidence.

## Required implementation work if Homebrew is chosen

If a later task chooses to implement Homebrew, it must do this exact work:

1. Create or identify the public tap repository.
2. Choose the formula name.
3. Define the formula source URL.
4. Define the sha256 source from published release checksums.
5. Test `brew install`.
6. Update README/docs only after the tested install path works.

The current intended tap remains:

```text
Nam-Cheol/homebrew-tap
```

## Risks

- False package availability claims: README or install docs could imply
  Homebrew works before a formula exists and has been tested.
- Stale formula checksums: a release asset or checksum update could leave the
  formula pointing at mismatched metadata.
- Release asset naming drift: future archives may change names or platforms,
  breaking formula URLs unless the naming contract stays stable.
- Package-manager maintenance burden: once published, users expect upgrades,
  checksums, audit fixes, and working install commands.
- Confusing Homebrew with ni-kernel behavior: package distribution is
  repository infrastructure, not Intent Lock Protocol behavior, runtime
  execution, or kernel-owned state.

## Status wording

Use this exact status wording:

```text
Homebrew: Planned
```

Acceptable supporting wording:

```text
No tap or formula is published or tested.
```

```text
Package-manager distribution, including Homebrew, is not available yet.
```

Do not use `Homebrew: Experimental` or `Homebrew: Available` until a tap,
formula, checksums, audit, install test, published tap install test, and
`ni --help` / `ni version` verification all exist.

## Next task

Improve conversation-authoring UX proof capture for v0.4.1 stabilization.
