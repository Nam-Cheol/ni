# v0.5 Publication Human Approval Packet

## Current status

- RC decision: RC_READY_WITH_DEFERRALS
- Release notes preflight decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES
- Artifact dry-run decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS
- Publication checklist decision: PUBLICATION_CHECKLIST_READY_WITH_NOTES
- Release binary: Available for verified release assets; v0.5 publication is not
  claimed by this packet.
- Curl installer: Available for verified release assets; v0.5 hosted install is
  not claimed by this packet.
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- `ni` is a pre-runtime Project Intent Compiler for AI Agents.
- `ni run` compiles a bounded handoff prompt only and does not execute
  downstream work.
- READY is planning contract readiness only.
- LOCK-STALE means the existing lock no longer matches current planning inputs.
- Fixture relock is separate from project-root relock.
- Benchmark evidence keeps `not_measured` boundaries.
- This packet does not publish, tag, create a GitHub release, upload assets,
  run a release workflow, run GoReleaser publish, create or publish a Homebrew
  formula, execute generated prompts, or mark v0.5 as released.
- No human approval is granted by this packet.

## Approval-packet goal

This packet is a non-executing maintainer approval record for a later,
separate, human-controlled v0.5 publication decision. It collects the exact
questions, required evidence, manual release gates, abort criteria, known
deferrals, and forbidden overclaims a maintainer must review before approving
any real publication-prep or publication action.

The packet is not a release plan executor. It does not select a tag, push a
tag, create a GitHub release, upload release assets, publish checksums, promote
release notes from draft, publish a Homebrew formula, or update public
availability claims.

## Approval-packet readiness decision

Decision: HUMAN_APPROVAL_PACKET_READY_WITH_NOTES.

Justification: the human approval surface is documented, the release-readiness
chain from docs/110 through docs/115 is visible, the manual approval choices
are explicit, and all release actions remain future-gated. Notes remain because
publication is not approved or performed here, docs/114 and docs/115 are new
until included in a later commit, Homebrew remains Planned / v0.5 candidate,
model workspace packs remain Experimental, no-terminal remains Experimental /
assisted, hosted v0.5 artifacts and checksums remain unavailable until
publication, cross-platform hosted verification remains deferred, and external
user validation remains limited.

## Release-readiness chain

| Doc | Decision / role | Status | Notes |
| --- | --- | --- | --- |
| docs/110 | RC_READY_WITH_DEFERRALS | Present and tracked | Records the release-candidate readiness audit without claiming v0.5 release completion. |
| docs/111 | Release notes draft | Present and tracked | Draft-only release notes; no publication, tag, upload, or Homebrew Available claim. |
| docs/112 | RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES | Present and tracked | Final release-note preflight preserves wording, validation, and no-release boundaries. |
| docs/113 | ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS | Present and tracked | Dry-run/check-only artifact readiness passed with explicit publication deferrals. |
| docs/114 | PUBLICATION_CHECKLIST_READY_WITH_NOTES | New in prior checklist task | Non-executing publication checklist; release actions are listed but not run. |
| docs/115 | HUMAN_APPROVAL_PACKET_READY_WITH_NOTES | New in this task | Human approval packet for a later explicit maintainer decision; no approval is granted here. |

## Human approval questions

The maintainer must answer these before any later real publication-prep or
publication task. This packet does not answer "yes" on the maintainer's behalf.

| Question | Required answer before later publication-prep work | Packet evidence | Notes |
| --- | --- | --- | --- |
| Proceed from RC_READY_WITH_DEFERRALS to a separate publication task? | Yes, explicitly from the human maintainer | Not selected here | This task only prepares the packet. |
| Have docs/110 through docs/114 been reviewed? | Yes | Chain listed above | Review must include English and Korean companions where maintained. |
| Are docs/110 through docs/115 included or ready for the next commit? | Yes | Git status / inclusion check | docs/115 is new in this task. |
| Are protected `.ni` files unchanged? | Yes | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Any diff blocks publication. |
| Does `ni status --proof --next-questions` report READY with no blockers, deferrals, or warnings? | Yes | Required evidence matrix | CLI is authority. |
| Do all check-only validations pass? | Yes | Required evidence matrix | Failing validation blocks publication. |
| Are release notes still draft-only until publication? | Yes | docs/111 and docs/112 | Do not claim v0.5 released before publication. |
| Are tag creation, GitHub release creation, asset upload, and checksum publication still not run? | Yes | Manual release gates | They require a later explicit human-approved task. |
| Does Homebrew remain Planned / v0.5 candidate? | Yes | Status and claim-boundary audit | Homebrew Available requires separate tap/formula evidence. |
| Do model workspace packs remain Experimental? | Yes | Status and claim-boundary audit | No global or host-level verification is claimed. |
| Does no-terminal remain Experimental / assisted? | Yes | Status and claim-boundary audit | No deterministic no-terminal validation is claimed. |
| Is the maintainer aware that `ni run` does not execute downstream work? | Yes | Boundary status | `ni run` remains bounded prompt compilation only. |
| Is the maintainer aware that benchmark evidence is bounded and not causal? | Yes | Claim-boundary audit | No implementation, adoption, cost, latency, or downstream-quality effect is claimed. |
| Is the maintainer aware that external user validation remains limited? | Yes | Known deferrals | Adoption proof remains a deferral. |
| Is the maintainer ready to stop if any post-approval pre-publication check fails? | Yes | Abort criteria | Any mismatch, overclaim, or validation failure stops publication. |

## Human decision options

Exactly one of these options may be selected by the maintainer in a later
message. None is selected by this packet.

| Option | Meaning | Selected by this packet? | Allowed next step |
| --- | --- | --- | --- |
| APPROVE_PUBLICATION_PREP_ONLY | Human approves a separate publication-prep task that may still be check-only or may prepare commands for human execution, depending on the next prompt. | No | A later explicit prompt must define scope and gates. |
| DO_NOT_APPROVE_FIX_FIRST | Human declines approval until identified issues are fixed and validations rerun. | No | Fix-only task; no publication action. |
| DEFER_PUBLICATION | Human defers publication decision. | No | Wait or perform unrelated docs/evidence work. |

## Required evidence matrix

| Evidence | Required result | Current packet status | Must be rerun before real publication? | Notes |
| --- | --- | --- | --- | --- |
| docs/110 | RC_READY_WITH_DEFERRALS preserved | Required | Yes | No v0.5 released claim. |
| docs/111 | Release notes remain draft-only | Required | Yes | No tag, release, asset, or Homebrew Available claim. |
| docs/112 | RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES preserved | Required | Yes | Release-note claim boundaries remain explicit. |
| docs/113 | ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS preserved | Required | Yes | Dry-run/check-only, not publication. |
| docs/114 | PUBLICATION_CHECKLIST_READY_WITH_NOTES preserved | Required | Yes | Publication actions listed as future-gated and not run. |
| docs/115 | HUMAN_APPROVAL_PACKET_READY_WITH_NOTES | New in this task | Yes | No human approval is granted here. |
| `git status --short` | Expected docs-only changes visible; no unexpected files | Required | Yes | Include docs/110 through docs/115 and docs/51 review. |
| `git ls-files docs/110_* ... docs/114_*` | docs/110 through docs/114 tracked as expected | Required | Yes | docs/115 is new until added in a later commit. |
| Protected `.ni` diff | No diff in `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` | Required | Yes | Any protected root planning diff blocks release actions. |
| `go test ./...` | Pass | Required | Yes | Use `GOCACHE=/private/tmp/ni-go-cache` if needed. |
| `ni status --dir . --proof --next-questions` | READY; blockers, deferrals, warnings are None | Required | Yes | Model judgment cannot replace this check. |
| `ni --help` | Help renders and preserves project intent compiler framing | Required | Yes | Does not execute downstream work. |
| `ni version` | Version output is understood and appropriate for the build context | Required | Yes | Source builds may show `0.0.0-dev`; final artifacts need release-version verification. |
| `python3 scripts/check-install-docs.py` | Pass | Required | Yes | Keeps install and distribution claims bounded. |
| `bash scripts/check-skill-packs.sh` | Pass | Required | Yes | Keeps model workspace packs Experimental and CLI-authority bounded. |
| `bash scripts/demo-check.sh` | Pass | Required | Yes | Prompt compilation must not become execution. |
| `bash scripts/quality.sh` | Pass | Required | Yes | Broad quality wrapper. |
| `bash scripts/smoke.sh` | Pass | Required | Yes | Fixture `ni end` / `ni relock` is not project-root relock. |
| `bash scripts/install-check.sh` | Pass | Required | Yes | Temporary/source install proof only. |
| `bash scripts/release-check.sh` | Pass | Required | Yes | Check-only release gate, not publication. |
| Generated artifacts review | No tracked generated artifacts or unexpected release outputs | Required | Yes | Ignored artifacts may exist but must not be staged accidentally. |
| No-overclaim review | No forbidden release, Homebrew, model workspace, no-terminal, ni-run, or benchmark claim | Required | Yes | See claim-boundary audit. |

## Manual release gates - future gated and not run

| Gate | Later action | Run in this task? | Requires later explicit human approval? | Required verification after action |
| --- | --- | --- | --- | --- |
| create release tag | Create the intended v0.5 tag on the approved commit | No | Yes | Tag points to intended commit. |
| push release tag | Push the approved tag | No | Yes | Remote tag exists and matches local tag. |
| create GitHub release | Create or trigger the release page | No | Yes | Release page matches intended tag and notes. |
| upload release assets | Upload or publish archives | No | Yes | Expected assets are downloadable. |
| publish checksums | Publish checksum file | No | Yes | Checksums match all hosted assets. |
| verify hosted assets | Download and verify release assets | No | Yes | Downloads, extraction, help/version, and checksums pass. |
| promote release notes from draft | Move approved notes into release/public docs | No | Yes | Notes match actual assets and avoid overclaims. |
| update public docs after publication | Update README/install/distribution surfaces | No | Yes | Public docs match verified publication evidence. |
| optional Homebrew formula work | Create or update tap/formula if separately approved | No | Yes | Formula, sha256, audit, `brew install`, `ni --help`, and `ni version` pass. |

## Known deferrals

| Deferral | Current boundary | Required future evidence |
| --- | --- | --- |
| actual v0.5 publication | Not approved and not run | Human approval, tag/release workflow, release page proof. |
| GitHub release creation | Not run | Release page exists and matches intended tag. |
| asset upload | Not run | Hosted assets are visible and downloadable. |
| hosted checksum availability | Not available | Checksum file exists and matches hosted assets. |
| Homebrew implementation / availability | Planned / v0.5 candidate | Tap, formula, sha256, audit, `brew install`, `ni --help`, `ni version`. |
| cross-platform hosted install verification | Not verified | Per-platform asset, extraction, help/version, checksum proof. |
| model workspace host verification | Experimental | Host-specific install/discovery transcript. |
| external user validation | Limited | Maintained user-run transcript or scoped external validation notes. |
| additional benchmark breadth if relevant | Deferred | Broader benchmark design with `not_measured` limits preserved. |
| no-terminal deterministic validation | Not claimed | Exact trusted-runner CLI output for the target workspace would be required; this packet does not provide it. |

## Abort criteria

| Condition | Required action | Why |
| --- | --- | --- |
| validation failure | Stop and fix before any publication-prep action | Release state is not trustworthy. |
| `ni status` is not READY | Stop and repair planning state | CLI authority blocks release action. |
| unexpected blockers, deferrals, or warnings | Stop and investigate | The chain no longer matches this packet. |
| protected `.ni` changed | Stop and investigate | Root planning state may have changed. |
| docs/110 through docs/115 missing from git inclusion review | Stop and fix inclusion | Human approval packet is incomplete. |
| tag mismatch | Stop before pushing or publishing | Wrong commit could be released. |
| asset checksum mismatch | Stop, remove or correct assets | Users cannot verify downloads. |
| hosted artifact unavailable | Stop public availability claims | Docs would point to missing artifacts. |
| curl installer mismatch | Stop curl v0.5 availability claim | Installer could fetch the wrong artifact. |
| version/help mismatch | Stop artifact availability claim | Binary may be wrong build/version. |
| overclaim in README/docs | Correct before publication | Public docs could promise unsupported status. |
| Homebrew availability claim without verification | Remove claim and keep Homebrew Planned | Homebrew evidence is absent. |
| generated prompt executed accidentally | Stop and document incident | `ni run` must remain prompt compilation only. |
| downstream execution behavior introduced accidentally | Stop and remove behavior | Violates `ni-kernel` boundary. |
| release action performed without explicit human approval | Stop and document incident | This packet does not grant approval. |

## Forbidden claims

- v0.5 has been released
- a v0.5 GitHub release exists
- v0.5 release assets were uploaded
- v0.5 hosted checksums are available
- Homebrew is Available
- model workspace packs are Available
- global model workspace verification is complete
- host-level model workspace verification is complete, unless actually tested
  and documented
- no-terminal deterministic validation
- no-terminal readiness without exact `ni status` output
- no-terminal lock freshness without exact CLI output from a trusted runner
- no-terminal hash verification without exact CLI output from a trusted runner
- no-terminal relock without exact `ni end` output from the target workspace
- no-terminal bounded handoff compilation without exact `ni run` output
- fixture relock is project-root relock
- validation-script fixture runs are project-root relock
- `ni run` executes downstream work
- benchmark evidence proves implementation quality, adoption, cost, latency, or
  downstream agent performance
- real research approval or fieldwork authorization
- dashboard product readiness
- human approval was granted

## Git status / inclusion check

| Path or group | Expected git status | Tracked-file expectation | Expected in next commit? | Notes |
| --- | --- | --- | --- | --- |
| docs/110_* | no output | tracked | No new change | Must remain present in chain. |
| docs/111_* | no output | tracked | No new change | Draft-only release notes remain bounded. |
| docs/112_* | no output | tracked | No new change | Preflight decision preserved. |
| docs/113_* | `M` from prior pointer work | tracked | Yes | Artifact dry-run audit remains part of release chain. |
| docs/114_* | `??` or tracked depending on commit state | untracked until added | Yes | Publication checklist remains new until committed. |
| docs/115_* | `??` until added | untracked until added | Yes | Human approval packet created by this task. |
| docs/51* | `M` from roadmap pointers | tracked | Yes | Roadmap points to docs/114 and docs/115. |
| generated artifacts | no tracked output expected | ignored or absent | No | Do not stage release archives, build output, or prompt artifacts. |
| `.ni/contract.json` | no output | tracked protected file | No | Must have no diff. |
| `.ni/session.json` | no output | tracked protected file | No | Must have no diff. |
| `.ni/plan.lock.json` | no output | tracked protected file | No | Must have no diff. |
| unexpected files | none | n/a | No | Any unexpected file requires review. |

## Claim-boundary audit

| Claim area | Expected boundary | Packet state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Human approval status | No approval granted unless the human explicitly grants it | Not granted | Yes | This packet asks for later approval only. |
| Published/released status | Do not claim v0.5 is published or released | Preserved | Yes | Publication is future-gated. |
| GitHub release status | Do not claim a v0.5 GitHub release exists | Preserved | Yes | Creation is future-gated. |
| Asset upload status | Do not claim assets were uploaded | Preserved | Yes | Upload is future-gated. |
| Checksum status | Do not claim hosted checksums exist | Preserved | Yes | Publication proof is absent. |
| Release binary | Available only within verified release-asset boundaries; no v0.5 publication claim | Preserved | Yes | v0.5 hosted assets require later evidence. |
| Curl installer | Available only within verified release-asset boundaries; no v0.5 hosted install claim | Preserved | Yes | v0.5 installer verification is future work. |
| Homebrew | Keep Planned / v0.5 candidate | Preserved | Yes | No Homebrew Available claim. |
| Model workspace packs | Keep Experimental | Preserved | Yes | No global or host-level availability claim. |
| No-terminal method | Keep Experimental / assisted | Preserved | Yes | No deterministic validation claim. |
| `ni run` | Bounded prompt compilation only | Preserved | Yes | No generated prompt execution. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits | Preserved | Yes | No causal, quality, adoption, cost, latency, or downstream-performance claim. |
| Fixture relock | Fixture relock is not project-root relock | Preserved | Yes | Validation fixture runs must remain separated. |
| Runtime execution boundary | `ni` is not a task runner, SPEC runner, execution layer, shell adapter, Codex exec adapter, queue, PR automation, or release automation | Preserved | Yes | No runtime behavior added. |

## Validation results

| Command | Required result | Packet result | Notes |
| --- | --- | --- | --- |
| `git status --short` | Expected docs-only status | Pass | Shows docs/51 and docs/114 modified, docs/115 untracked, and no unexpected tracked output. |
| `git ls-files docs/110_* ... docs/114_*` | docs/110 through docs/114 tracked as expected | Pass | docs/110 through docs/114 are tracked; docs/115 remains new until added. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | Pass | Check-only. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | READY; blockers, deferrals, warnings: None | Pass | CLI reported READY with blockers, deferrals, and warnings as None. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Help renders | Pass | Does not execute downstream work. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Version output understood for build context | Pass | Output: `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Pass | Pass | Keeps install claims bounded. |
| `bash scripts/check-skill-packs.sh` | Pass | Pass | Keeps model workspace packs Experimental. |
| `bash scripts/demo-check.sh` | Pass | Pass | No downstream prompt execution. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Pass | Broad quality wrapper. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Pass | Fixture-only `ni end` / `ni relock`, not project root. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Pass | Temporary/source install proof. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Pass | Check-only release gate. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | No diff | Pass | Protected root planning files unchanged. |

## Changes made by this packet

- `docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.md`: English human approval
  packet.
- `docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.ko.md`: Korean companion.
- `docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md`: narrow follow-up pointer to
  docs/115.
- `docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md`: matching Korean pointer.
- `docs/51_POST_RELEASE_ROADMAP.md`: narrow roadmap pointer to docs/115.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer.

## What this packet proves

- the exact human approval surface is documented
- manual release gates are separated from this task
- human decision options are visible and unselected
- known deferrals and abort criteria are explicit
- forbidden claims remain bounded
- no release action was performed by this packet

## What this packet does not prove

- v0.5 has been released
- a v0.5 GitHub release exists
- v0.5 assets were uploaded
- hosted checksums are available
- Homebrew is Available
- cross-platform hosted install works
- external users succeed
- model workspace host behavior is verified
- no-terminal is deterministic
- downstream execution succeeds
- benchmark effect size or causal impact exists
- human approval was granted

## Recommended next task

Selected next task: A. Wait for human approval.

Why: the packet is ready with notes, and the next real step must be a human
decision outside Codex. No execution prompt is provided here because any release
or publication-prep action requires explicit maintainer selection of one human
decision option.

## Human approval request template

```text
Please review docs/110 through docs/115, the final validation report, and the
protected .ni diff result.

Choose exactly one:
- APPROVE_PUBLICATION_PREP_ONLY
- DO_NOT_APPROVE_FIX_FIRST
- DEFER_PUBLICATION

No publication, tag, GitHub release creation, asset upload, checksum
publication, Homebrew work, generated prompt execution, or availability-claim
upgrade should occur unless the next message explicitly approves and scopes that
separate task.
```
