# Homebrew Tap Plan

Date: 2026-05-28

Scope: Plan the Homebrew distribution route after verified release binaries and
the verified curl installer path, without claiming package-manager availability.

Current Homebrew status: Planned.

Evidence: `git ls-remote https://github.com/Nam-Cheol/homebrew-tap.git`
returned "Repository not found" on 2026-05-28. There is no owner-confirmed tap,
no published formula, and no tested `brew install` path.

## Route Decision

Use a separate tap repository as the Homebrew distribution route:

```text
Nam-Cheol/homebrew-tap
```

The first formula should be added manually to the tap repository after the tap
exists and ownership is confirmed. The repository-local
[Homebrew Formula Draft](71_HOMEBREW_FORMULA_DRAFT.md) is only a
non-published draft for that future tap.

Do not add GoReleaser Homebrew publishing yet. GoReleaser-generated formula
updates are a later option after the tap exists, the owner explicitly confirms
that GoReleaser should update the tap, and `goreleaser check` passes with the
brew configuration included.

## Route Matrix

| Route | Decision | Status |
| --- | --- | --- |
| Separate tap repository | Use `Nam-Cheol/homebrew-tap` as the intended external tap. | Planned; repository not found. |
| Manual formula draft | Keep only as a non-published draft until copied into the tap. | Draft only in `docs/71_HOMEBREW_FORMULA_DRAFT.md`. |
| GoReleaser-generated formula | Defer until the tap exists and owner confirms GoReleaser should maintain it. | Not configured. |

## Next Steps To Create The Tap

1. Create `Nam-Cheol/homebrew-tap` as a public GitHub repository.
2. Clone the tap repository locally.
3. Create `Formula/ni.rb` inside the tap repository, not in `ni-kernel`.
4. Start from `docs/71_HOMEBREW_FORMULA_DRAFT.md`.
5. Replace every placeholder checksum with the matching checksum from
   `docs/70_RELEASE_VERIFICATION_v0.3.0.md` or the published
   `ni_0.3.0_checksums.txt` release asset.
6. Validate the formula from the tap checkout:

   ```bash
   brew audit --strict --online Formula/ni.rb
   brew install --build-from-source Formula/ni.rb
   ni --help
   ni version
   ```

7. Only after the local formula validation works, test the published tap command:

   ```bash
   brew install Nam-Cheol/tap/ni
   ni --help
   ni version
   ```

8. Update README and install docs from Planned to Available only after that
   exact published tap command works.

## Availability Rule

Homebrew remains Planned until all of these are true:

1. The tap repository exists.
2. The owner confirms it is the intended tap.
3. `Formula/ni.rb` is published in the tap with real release checksums.
4. Homebrew tooling validates the formula.
5. `brew install Nam-Cheol/tap/ni` works from a clean Homebrew environment.

Until then, do not add a Homebrew badge, README `brew install` command,
package-manager availability claim, package publishing automation, or any
`ni-kernel` runtime behavior.
