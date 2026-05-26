# v0.1.0 Release Readiness Checklist

Use this checklist before tagging `v0.1.0`. The release gate is deterministic:
run the local script, verify CI is green for the commit being tagged, then tag
manually.

Do not publish packages, create a GitHub release automatically, add GoReleaser,
or add release automation as part of this checklist.

## Local Gate

```bash
bash scripts/release-check.sh
```

The script must pass on the exact commit intended for the tag.

## Required Checks

- [ ] CI passes for the commit being tagged.
- [ ] Smoke passes through `bash scripts/smoke.sh`.
- [ ] Golden tests pass through the Go test suite.
- [ ] Schemas validate through `python3 scripts/check-schema.py`.
- [ ] README quickstart verified for source, built binary, and local install modes.
- [ ] No stale roadmap references remain in release-facing roadmap docs.
- [ ] No core-boundary violations are reported.
- [ ] No release automation claims are present.
- [ ] No runtime execution claims are present.
- [ ] All public commands listed in README have smoke coverage.

## Manual Tagging

After the local gate passes and CI is green, tag manually:

```bash
git tag v0.1.0
git push origin v0.1.0
```

This checklist does not publish packages or create hosted release artifacts.
