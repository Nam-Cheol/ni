# v0.5.1 Patch Release Plan

## Current status

State:
- v0.5.0 publication: verified.
- v0.5.0 release binary: Available.
- v0.5.0 curl installer: Available.
- published v0.5.0 install lane: `ni --help` and `ni version` passed.
- published v0.5.0 `ni version`: `0.5.0`.
- published v0.5.0 `ni init .`: failed with `unknown init option: .`.
- public install parity decision from docs/126: PUBLIC_INSTALL_PARITY_MISMATCH_V0_5_1_PATCH_NEEDED.
- current-tree `ni init .` Bubble Tea v2 / Lip Gloss v2 TUI: implemented.
- current-tree first-user onboarding smoke after TUI: FIRST_USER_ONBOARDING_SMOKE_PASS_WITH_NOTES.
- Windows real-host execution: deferred on the macOS-only development host.
- Homebrew: Planned / v0.5 candidate.
- Model workspace packs: Experimental.
- No-terminal method: Experimental / assisted.
- Skills are UX; CLI is authority.
- `ni run` compiles a bounded handoff prompt and does not execute downstream work.

This document is a release plan only. It does not publish, tag, create a GitHub
Release, upload assets, run release workflows, run GoReleaser publish, create
or publish a Homebrew formula, run `ni end` on the project root, relock the
project root, execute generated prompts, or add runtime execution behavior.

## Decision

V0_5_1_PATCH_PLAN_READY_WITH_NOTES

Justification: the patch scope is clear and current-tree validation can prove
the behavior that public v0.5.0 lacks. The plan still carries notes because
Windows real-host execution, Homebrew availability, external user validation,
and post-publication hosted artifact parity must remain deferred until their
own transcripts exist.

## Release rationale

The public onboarding docs now teach a first-success path that includes
`ni init .`. Task 205 showed that published v0.5.0 can be installed and can run
`ni --help` and `ni version`, but fails the current-directory init step with
`unknown init option: .`. Current-tree work after the `v0.5.0` tag added
positional `ni init .`, a guided TUI, non-interactive fallback behavior, and
install-path handling that the public first-user path depends on.

The right release action is therefore a small v0.5.1 patch that packages the
already implemented onboarding and install-parity fixes. The patch must not
become a new execution layer or a Homebrew availability release.

## Patch scope

Required v0.5.1 contents:

- `ni init .` positional target support.
- Bubble Tea v2 / Lip Gloss v2 interactive init TUI.
- Non-interactive fallback for CI and non-TTY contexts.
- Domain init logic separated from TUI rendering.
- Existing file protection; init must not silently overwrite planning files.
- `.ni/plan.lock.json` protection; init must not modify an existing lockfile.
- Post-TUI plain text summary.
- README macOS / Windows two-path onboarding.
- Install docs aligned to command-name verification through `ni --help` and
  `ni version`.
- Current checker alignment for install docs, Windows installer static safety,
  README surface, demo, smoke, install, and release checks.
- No downstream execution behavior.

Optional current-tree contents that may be included because they are already
present and bounded:

- Cleanup of accidental duplicate `examples/namba-ai-upgrade/docs/plan/* 2.md`
  files.
- Deterministic README SVG assets and asset drift checks.
- macOS/Linux global install path handling updates in `install.sh`.
- `install.sh --update-path` and `install.sh --uninstall` improvements.
- `install.ps1` User PATH safety and uninstall behavior.

Explicit exclusions:

The patch must not include these kernel-boundary items:

- Task runner, SPEC runner, execution harness, shell adapter, Codex exec
  adapter, queue, PR automation, release automation inside `ni-kernel`, or
  downstream execution layer.
- Homebrew Available claim, tap publication, or formula publication.
- Windows real-host verified claim without a Windows transcript.
- Model workspace packs as broad Available status.
- No-terminal deterministic validation claim.
- Any claim that `ni run` executes downstream work.

## Current-tree comparison

Commits after `v0.5.0` that belong in the v0.5.1 patch scope:

| Commit | Role in patch |
| --- | --- |
| `34a613a` Simplify README install paths and add init TUI onboarding | Two-path README and initial `ni init .` onboarding implementation. |
| `0745621` Implement global install path handling | macOS/Linux installer PATH handling needed by command-name verification. |
| `9e587d8` Update README install docs for v0.5.0 release | Release/install doc baseline for current public path. |
| `c943a53` Add README visual asset pass references | Deterministic README visual support, optional but already in current tree. |
| `4ae4a85` Redesign README and add Bubble Tea init TUI | Main TUI, domain separation, tests, and duplicate cleanup. |
| `93dc241` Document first-user onboarding smoke after TUI | Current-tree onboarding smoke evidence. |

Current uncommitted documentation that belongs in the release-planning bundle:

| Path | Role |
| --- | --- |
| `README.md` | Public parity note linking docs/126. |
| `README.ko.md` | Korean companion parity note. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md` | Public v0.5.0 mismatch evidence. |
| `docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.ko.md` | Korean companion mismatch evidence. |
| `docs/127_V0_5_1_PATCH_RELEASE_PLAN.md` | This release plan. |
| `docs/127_V0_5_1_PATCH_RELEASE_PLAN.ko.md` | Korean companion release plan. |

Current-tree surfaces that support the patch:

| Surface | Evidence |
| --- | --- |
| CLI init routing | `cmd/ni/main.go` accepts `ni init [.]` and routes interactive positional init to TUI when safe. |
| Init domain logic | `internal/core/docstore` builds file plans, skips existing files, and protects `.ni/plan.lock.json`. |
| Init TUI | `internal/tui/init` uses Bubble Tea v2 and Lip Gloss v2 for model/view rendering. |
| Tests | `cmd/ni/main_test.go`, `internal/core/docstore/docstore_test.go`, and `internal/tui/init/model_test.go` cover positional init, fallback, existing files, lockfile safety, and TUI model behavior. |
| Installers | `install.sh` supports `--update-path` / `--uninstall`; `install.ps1` updates User PATH only. |
| Release versioning | `.goreleaser.yaml` injects `ni/internal/version.Version={{ .Version }}` for release builds. |
| Release workflow | `.github/workflows/release.yml` runs tests, quality, release-check, and GoReleaser on `v*` tags. |

Surfaces that should not drive v0.5.1 scope:

| Surface | Treatment |
| --- | --- |
| Homebrew draft formula | Keep as Planned / candidate evidence only; no Available claim or tap publication. |
| Windows installer real execution | Keep deferred unless a real Windows install/new-session/uninstall transcript exists. |
| Model workspace packs | Keep Experimental and CLI-authority bounded. |
| No-terminal method | Keep assisted only; no deterministic validation claim. |
| Downstream seed material | May remain derived/mutable documentation, not kernel-owned execution state. |

## Release criteria

Before v0.5.1 can be published, all required criteria must pass:

| Criterion | Required result |
| --- | --- |
| `go test ./...` | Passes on current tree. |
| `ni status --proof --next-questions` on project root | Reports the actual current state; only CLI output may support readiness wording. |
| Local `ni --help` | Works and includes `ni init [.]`. |
| Local `ni version` | Works; source/local build may report `0.0.0-dev` unless linker flags are set. |
| First-user current-tree smoke | Passes with notes from docs/125 or a newer equivalent. |
| Temporary `ni init . --yes` | Creates `.ni/contract.json`, `.ni/session.json`, and `docs/plan/**`. |
| Repeated `ni init . --yes` | Does not silently overwrite existing planning files. |
| Lockfile safety | Existing `.ni/plan.lock.json` is not modified by init. |
| Install docs checks | `scripts/check-install-docs.py` passes. |
| Windows static installer checks | `scripts/check-install-ps1.py` passes; this is not real-host proof. |
| Release check | `scripts/release-check.sh` passes. |
| Protected `.ni` files | Project-root `.ni/contract.json`, `.ni/session.json`, and `.ni/plan.lock.json` remain unchanged unless an explicitly authorized amendment flow is used. |
| Release notes delta | Accurately states the v0.5.1 patch delta and avoids overclaims. |
| Homebrew wording | No Homebrew Available claim. |
| Windows wording | No Windows real-host verified claim unless transcript exists. |
| `ni run` wording | Prompt compilation only; no downstream execution claim. |

## Publication readiness gates

The maintainer must complete these gates before actual v0.5.1 publication:

1. Confirm the release commit contains only the approved patch scope.
2. Bump or set the release version if required by the project release process.
3. Verify release build linker flags make `ni version` report `0.5.1`.
4. Run full validation: `gofmt -w .`, `go test ./...`, and
   `bash scripts/quality.sh`.
5. Run `bash scripts/release-check.sh`.
6. Create a release notes delta focused on public install parity.
7. Run an artifact dry-run where tooling is available.
8. Verify checksums for generated archives.
9. Confirm `install.sh` retrieves v0.5.1, not stale v0.5.0 artifacts.
10. Confirm current-platform release binary runs `ni --help` and `ni version`
    by command name.
11. Confirm isolated curl installer install for v0.5.1 in temporary
    `HOME` / `BINDIR`.
12. Confirm isolated `ni init .` and `ni status --proof --next-questions`
    with the installed v0.5.1 binary.
13. Update README only if needed to keep public install instructions honest.
14. Keep Windows real-host verification deferred unless transcript exists.
15. Keep Homebrew Planned / v0.5 candidate unless full tap/formula/checksum/
    audit/install/help/version/uninstall proof passes in a separate task.

## Validation matrix

| Phase | Command or check | Expected |
| --- | --- | --- |
| planning task | `git status --short` | Shows only expected docs/README planning changes; no staging. |
| planning task | `git log --oneline --decorate -20` | Shows post-`v0.5.0` onboarding and install commits. |
| planning task | `git diff --name-only v0.5.0..HEAD` | Identifies candidate patch files. |
| planning task | `git diff --stat v0.5.0..HEAD` | Confirms patch size and surfaces. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Reports actual project-root status. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Pass. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Pass with source version, normally `0.0.0-dev`. |
| planning task | `python3 scripts/check-install-docs.py` | Pass. |
| planning task | `python3 scripts/check-install-ps1.py` | Pass as static Windows safety only. |
| planning task | `bash scripts/check-skill-packs.sh` | Pass; preserves CLI authority wording. |
| planning task | `bash scripts/demo-check.sh` | Pass. |
| planning task | `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass. |
| planning task | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Empty. |
| pre-publication | Release build `ni version` | Reports `0.5.1`. |
| pre-publication | Artifact dry-run and checksums | Archives and `ni_0.5.1_checksums.txt` are internally consistent. |
| post-publication | Isolated installed `ni --help` / `ni version` | Command-name checks pass; version reports `0.5.1`. |
| post-publication | Isolated installed `ni init .` | Positional init works with the release binary. |
| post-publication | Isolated installed `ni status --proof --next-questions` | Runs after init and reports actual readiness state. |

## Rollback or docs correction path

If v0.5.1 publication is not approved or validation fails:

- Do not publish or tag.
- Keep docs/126 as the public v0.5.0 parity mismatch record.
- Keep README parity wording that warns v0.5.0 lacks positional `ni init .`, or
  adjust the README first-project path so it no longer implies v0.5.0 parity.
- Record the failed gate and exact command output in a follow-up audit.
- Re-run this plan only after the blocker is resolved.

If v0.5.1 is published but hosted parity fails:

- Stop claiming public install parity for `ni init .`.
- Add a post-publication mismatch note before broadening install claims.
- Prefer docs correction over hiding the kernel boundary or weakening
  acceptance criteria.

## Protected files

This planning task must not edit, relock, or silently regenerate:

- `.ni/contract.json`
- `.ni/session.json`
- `.ni/plan.lock.json`

If any protected file diff appears outside an explicitly authorized amendment
or relock flow, the release-plan decision becomes
`V0_5_1_PATCH_PLAN_BLOCKED`.

## Next executable prompt

```text
Proceed in /Users/namba/Documents/project/ni.

Goal:
Run the authorized v0.5.1 release preflight for public install parity. This is
still pre-publication unless the user separately authorizes tagging and release.

Read:
- AGENTS.md
- docs/126_PUBLIC_INSTALL_PARITY_AND_PATCH_READINESS.md
- docs/127_V0_5_1_PATCH_RELEASE_PLAN.md
- README.md
- README.ko.md
- docs/22_INSTALL.md
- docs/install-curl.md
- install.sh
- install.ps1
- .goreleaser.yaml
- .github/workflows/release.yml
- cmd/ni/
- internal/
- scripts/release-check.sh

Required checks:
- git status --short
- git diff --name-only v0.5.0..HEAD
- git diff --stat v0.5.0..HEAD
- gofmt -w .
- GOCACHE=/private/tmp/ni-go-cache go test ./...
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version
- GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions
- temporary current-tree binary smoke: ni init . --yes, repeated init, lockfile protection
- python3 scripts/check-install-docs.py
- python3 scripts/check-install-ps1.py
- bash scripts/check-skill-packs.sh
- bash scripts/demo-check.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh
- GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh
- git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json

Release-preflight output:
- exact release delta for v0.5.1
- validation transcript summary
- release notes draft or update if needed
- artifact dry-run result if GoReleaser is available
- explicit decision: V0_5_1_PREFLIGHT_READY,
  V0_5_1_PREFLIGHT_READY_WITH_NOTES, or V0_5_1_PREFLIGHT_BLOCKED

Rules:
- Do not publish, tag, create a GitHub Release, upload assets, run release
  workflows, run goreleaser publish, create or publish a Homebrew formula, run
  ni end on the project root, relock the project root, execute generated
  prompts, or add runtime execution behavior unless the user separately
  authorizes that exact release action.
- Do not mark Homebrew Available.
- Do not claim Windows real-host execution verified unless a Windows transcript
  exists.
- Keep Skills are UX; CLI is authority.
- Keep ni run as bounded prompt compilation only.

Final response:
Report changed files, preflight decision, validation results, protected .ni
diff, remaining notes, and confirmation that no publish/tag/release/upload/root
relock/generated prompt execution occurred.
```

## Validation run in this task

| Command | Result |
| --- | --- |
| `git status --short` | To be rechecked after validation. |
| `git log --oneline --decorate -20` | Checked; post-`v0.5.0` onboarding and install commits are visible. |
| `git diff --name-only v0.5.0..HEAD` | Checked; candidate patch files identified. |
| `git diff --stat v0.5.0..HEAD` | Checked; candidate patch size identified. |
| Required ripgrep scans from the task prompt | Checked; no release-scope expansion needed. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | Passed; project root reported `NI Intent Readiness: READY` with no blockers, deferrals, or warnings. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Passed; help includes `ni init [.]`. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Passed; source build output `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Passed. |
| `python3 scripts/check-install-ps1.py` | Passed. |
| `bash scripts/check-skill-packs.sh` | Passed. |
| `bash scripts/demo-check.sh` | Passed. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Passed after adding explicit `must not include` boundary context to the excluded-scope list. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Passed after the same boundary-context fix. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Passed; empty diff. |

Note: the first `quality.sh` and `release-check.sh` attempts failed because the
new excluded-scope list used boundary phrases without enough local `must not`
context for `scripts/check-core-boundary.py`. The docs were clarified; no
release scope was broadened.
