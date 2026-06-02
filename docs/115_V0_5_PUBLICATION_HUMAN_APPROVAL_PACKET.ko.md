# v0.5 Publication Human Approval Packet

## Current status

- RC decision: RC_READY_WITH_DEFERRALS
- Release notes preflight decision: RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES
- Artifact dry-run decision: ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS
- Publication checklist decision: PUBLICATION_CHECKLIST_READY_WITH_NOTES
- Release binary: verified release assets 기준 Available; 이 packet은 v0.5
  publication을 claim하지 않는다.
- Curl installer: verified release assets 기준 Available; 이 packet은 v0.5 hosted
  install을 claim하지 않는다.
- Homebrew: Planned / v0.5 candidate
- Model workspace packs: Experimental
- No-terminal method: Experimental / assisted
- CLI is authority.
- Skills are UX.
- Skills are UX; CLI is authority.
- `ni`는 AI Agents를 위한 pre-runtime Project Intent Compiler다.
- `ni run`은 bounded handoff prompt만 compile하며 downstream work를 execute하지
  않는다.
- READY는 planning contract readiness만 의미한다.
- LOCK-STALE은 existing lock이 current planning inputs와 더 이상 match하지
  않는다는 뜻이다.
- Fixture relock은 project-root relock과 별개다.
- Benchmark evidence는 `not_measured` boundaries를 유지한다.
- 이 packet은 publish, tag, GitHub release creation, asset upload, release
  workflow, GoReleaser publish, Homebrew formula creation/publication, generated
  prompt execution, v0.5 released marking을 수행하지 않는다.
- 이 packet은 human approval을 grant하지 않는다.

## Approval-packet goal

이 packet은 later, separate, human-controlled v0.5 publication decision을 위한
non-executing maintainer approval record다. Maintainer가 real publication-prep
또는 publication action을 approve하기 전에 검토해야 할 exact questions,
required evidence, manual release gates, abort criteria, known deferrals,
forbidden overclaims를 모은다.

이 packet은 release plan executor가 아니다. Tag selection, tag push, GitHub
release creation, release asset upload, checksum publication, release notes
draft promotion, Homebrew formula publication, public availability claims update를
하지 않는다.

## Approval-packet readiness decision

Decision: HUMAN_APPROVAL_PACKET_READY_WITH_NOTES.

Justification: human approval surface가 documented되었고, docs/110부터 docs/115까지의
release-readiness chain이 visible하며, manual approval choices가 explicit하고, 모든
release actions가 future-gated 상태로 남아 있다. Notes는 publication이 여기서
approved 또는 performed되지 않았고, docs/114와 docs/115가 later commit에 포함되기
전까지 new 상태이며, Homebrew는 Planned / v0.5 candidate, model workspace packs는
Experimental, no-terminal은 Experimental / assisted, hosted v0.5 artifacts와
checksums는 publication 전까지 unavailable, cross-platform hosted verification은
deferred, external user validation은 limited이기 때문에 남는다.

## Release-readiness chain

| Doc | Decision / role | Status | Notes |
| --- | --- | --- | --- |
| docs/110 | RC_READY_WITH_DEFERRALS | Present and tracked | v0.5 release completion을 claim하지 않는 release-candidate readiness audit. |
| docs/111 | Release notes draft | Present and tracked | Draft-only release notes; publication, tag, upload, Homebrew Available claim 없음. |
| docs/112 | RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES | Present and tracked | Final release-note preflight가 wording, validation, no-release boundaries를 보존한다. |
| docs/113 | ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS | Present and tracked | Dry-run/check-only artifact readiness passed with explicit publication deferrals. |
| docs/114 | PUBLICATION_CHECKLIST_READY_WITH_NOTES | New in prior checklist task | Non-executing publication checklist; release actions는 listed but not run. |
| docs/115 | HUMAN_APPROVAL_PACKET_READY_WITH_NOTES | New in this task | Later explicit maintainer decision을 위한 human approval packet; approval은 granted되지 않았다. |

## Human approval questions

Maintainer는 later real publication-prep 또는 publication task 전에 아래 질문에
답해야 한다. 이 packet은 maintainer 대신 "yes"를 답하지 않는다.

| Question | Required answer before later publication-prep work | Packet evidence | Notes |
| --- | --- | --- | --- |
| RC_READY_WITH_DEFERRALS에서 separate publication task로 진행할 것인가? | Human maintainer의 explicit Yes | Not selected here | 이 task는 packet만 준비한다. |
| docs/110부터 docs/114까지 review했는가? | Yes | Chain listed above | Maintained되는 곳에서는 English/Korean companions까지 review한다. |
| docs/110부터 docs/115까지 next commit에 포함되었거나 포함할 준비가 되었는가? | Yes | Git status / inclusion check | docs/115는 이 task에서 new. |
| Protected `.ni` files가 unchanged인가? | Yes | `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | Any diff blocks publication. |
| `ni status --proof --next-questions`가 blockers, deferrals, warnings 없이 READY인가? | Yes | Required evidence matrix | CLI is authority. |
| 모든 check-only validations가 pass했는가? | Yes | Required evidence matrix | Failing validation blocks publication. |
| Release notes는 publication 전까지 draft-only인가? | Yes | docs/111 and docs/112 | Publication 전 v0.5 released claim 금지. |
| Tag creation, GitHub release creation, asset upload, checksum publication은 아직 not run인가? | Yes | Manual release gates | Later explicit human-approved task가 필요하다. |
| Homebrew는 Planned / v0.5 candidate로 남아 있는가? | Yes | Status and claim-boundary audit | Homebrew Available에는 별도 tap/formula evidence 필요. |
| Model workspace packs는 Experimental로 남아 있는가? | Yes | Status and claim-boundary audit | Global 또는 host-level verification claim 없음. |
| No-terminal은 Experimental / assisted로 남아 있는가? | Yes | Status and claim-boundary audit | Deterministic no-terminal validation claim 없음. |
| Maintainer는 `ni run`이 downstream work를 execute하지 않는다는 점을 알고 있는가? | Yes | Boundary status | `ni run`은 bounded prompt compilation only. |
| Maintainer는 benchmark evidence가 bounded이며 causal이 아니라는 점을 알고 있는가? | Yes | Claim-boundary audit | Implementation, adoption, cost, latency, downstream-quality effect claim 없음. |
| Maintainer는 external user validation이 limited라는 점을 알고 있는가? | Yes | Known deferrals | Adoption proof remains a deferral. |
| Maintainer는 post-approval pre-publication check가 실패하면 stop할 준비가 되어 있는가? | Yes | Abort criteria | Any mismatch, overclaim, validation failure stops publication. |

## Human decision options

Later message에서 maintainer는 아래 option 중 정확히 하나만 선택할 수 있다. 이
packet은 아무 option도 선택하지 않는다.

| Option | Meaning | Selected by this packet? | Allowed next step |
| --- | --- | --- | --- |
| APPROVE_PUBLICATION_PREP_ONLY | Human이 separate publication-prep task를 approve한다. 다음 prompt에 따라 still check-only일 수도 있고 human execution용 commands를 prepare할 수도 있다. | No | Later explicit prompt가 scope와 gates를 정의해야 한다. |
| DO_NOT_APPROVE_FIX_FIRST | Human이 identified issues를 fix하고 validations를 rerun하기 전까지 approval을 거절한다. | No | Fix-only task; publication action 없음. |
| DEFER_PUBLICATION | Human이 publication decision을 defer한다. | No | Wait 또는 unrelated docs/evidence work. |

## Required evidence matrix

| Evidence | Required result | Current packet status | Must be rerun before real publication? | Notes |
| --- | --- | --- | --- | --- |
| docs/110 | RC_READY_WITH_DEFERRALS preserved | Required | Yes | v0.5 released claim 없음. |
| docs/111 | Release notes remain draft-only | Required | Yes | Tag, release, asset, Homebrew Available claim 없음. |
| docs/112 | RELEASE_NOTES_PREFLIGHT_PASS_WITH_NOTES preserved | Required | Yes | Release-note claim boundaries remain explicit. |
| docs/113 | ARTIFACT_DRY_RUN_PASS_WITH_DEFERRALS preserved | Required | Yes | Dry-run/check-only, not publication. |
| docs/114 | PUBLICATION_CHECKLIST_READY_WITH_NOTES preserved | Required | Yes | Publication actions listed as future-gated and not run. |
| docs/115 | HUMAN_APPROVAL_PACKET_READY_WITH_NOTES | New in this task | Yes | Human approval은 granted되지 않았다. |
| `git status --short` | Expected docs-only changes visible; no unexpected files | Required | Yes | docs/110부터 docs/115와 docs/51 review 포함. |
| `git ls-files docs/110_* ... docs/114_*` | docs/110부터 docs/114 tracked as expected | Required | Yes | docs/115는 later commit 전까지 new. |
| Protected `.ni` diff | `.ni/contract.json`, `.ni/session.json`, `.ni/plan.lock.json` diff 없음 | Required | Yes | Any protected root planning diff blocks release actions. |
| `go test ./...` | Pass | Required | Yes | 필요하면 `GOCACHE=/private/tmp/ni-go-cache` 사용. |
| `ni status --dir . --proof --next-questions` | READY; blockers, deferrals, warnings are None | Required | Yes | Model judgment cannot replace this check. |
| `ni --help` | Help renders and preserves project intent compiler framing | Required | Yes | Downstream work execute 없음. |
| `ni version` | Version output is understood and appropriate for the build context | Required | Yes | Source builds may show `0.0.0-dev`; final artifacts need release-version verification. |
| `python3 scripts/check-install-docs.py` | Pass | Required | Yes | Install/distribution claims bounded 유지. |
| `bash scripts/check-skill-packs.sh` | Pass | Required | Yes | Model workspace packs Experimental and CLI-authority bounded 유지. |
| `bash scripts/demo-check.sh` | Pass | Required | Yes | Prompt compilation이 execution으로 변하면 안 된다. |
| `bash scripts/quality.sh` | Pass | Required | Yes | Broad quality wrapper. |
| `bash scripts/smoke.sh` | Pass | Required | Yes | Fixture `ni end` / `ni relock` is not project-root relock. |
| `bash scripts/install-check.sh` | Pass | Required | Yes | Temporary/source install proof only. |
| `bash scripts/release-check.sh` | Pass | Required | Yes | Check-only release gate, not publication. |
| Generated artifacts review | No tracked generated artifacts or unexpected release outputs | Required | Yes | Ignored artifacts가 있을 수 있지만 accidentally staged되면 안 된다. |
| No-overclaim review | Forbidden release, Homebrew, model workspace, no-terminal, ni-run, benchmark claim 없음 | Required | Yes | Claim-boundary audit 참고. |

## Manual release gates - future gated and not run

| Gate | Later action | Run in this task? | Requires later explicit human approval? | Required verification after action |
| --- | --- | --- | --- | --- |
| create release tag | Approved commit에 intended v0.5 tag 생성 | No | Yes | Tag points to intended commit. |
| push release tag | Approved tag push | No | Yes | Remote tag exists and matches local tag. |
| create GitHub release | Release page create 또는 trigger | No | Yes | Release page matches intended tag and notes. |
| upload release assets | Archives upload 또는 publish | No | Yes | Expected assets are downloadable. |
| publish checksums | Checksum file publish | No | Yes | Checksums match all hosted assets. |
| verify hosted assets | Release assets download and verify | No | Yes | Downloads, extraction, help/version, checksums pass. |
| promote release notes from draft | Approved notes를 release/public docs로 이동 | No | Yes | Notes match actual assets and avoid overclaims. |
| update public docs after publication | README/install/distribution surfaces update | No | Yes | Public docs match verified publication evidence. |
| optional Homebrew formula work | Separately approved이면 tap/formula create/update | No | Yes | Formula, sha256, audit, `brew install`, `ni --help`, `ni version` pass. |

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
| no-terminal deterministic validation | Not claimed | Target workspace에 대한 exact trusted-runner CLI output이 필요하다; 이 packet은 제공하지 않는다. |

## Abort criteria

| Condition | Required action | Why |
| --- | --- | --- |
| validation failure | Publication-prep action 전 stop and fix | Release state is not trustworthy. |
| `ni status` is not READY | Stop and repair planning state | CLI authority blocks release action. |
| unexpected blockers, deferrals, or warnings | Stop and investigate | Chain이 이 packet과 더 이상 match하지 않는다. |
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
| generated artifacts | no tracked output expected | ignored or absent | No | Release archives, build output, prompt artifacts를 stage하지 않는다. |
| `.ni/contract.json` | no output | tracked protected file | No | Must have no diff. |
| `.ni/session.json` | no output | tracked protected file | No | Must have no diff. |
| `.ni/plan.lock.json` | no output | tracked protected file | No | Must have no diff. |
| unexpected files | none | n/a | No | Any unexpected file requires review. |

## Claim-boundary audit

| Claim area | Expected boundary | Packet state | Pass? | Notes |
| --- | --- | --- | --- | --- |
| Human approval status | Human이 explicitly grant하지 않는 한 approval 없음 | Not granted | Yes | 이 packet은 later approval만 요청한다. |
| Published/released status | v0.5 published/released claim 금지 | Preserved | Yes | Publication is future-gated. |
| GitHub release status | v0.5 GitHub release exists claim 금지 | Preserved | Yes | Creation is future-gated. |
| Asset upload status | Assets uploaded claim 금지 | Preserved | Yes | Upload is future-gated. |
| Checksum status | Hosted checksums exist claim 금지 | Preserved | Yes | Publication proof is absent. |
| Release binary | Verified release-asset boundaries 안에서만 Available; v0.5 publication claim 없음 | Preserved | Yes | v0.5 hosted assets require later evidence. |
| Curl installer | Verified release-asset boundaries 안에서만 Available; v0.5 hosted install claim 없음 | Preserved | Yes | v0.5 installer verification is future work. |
| Homebrew | Keep Planned / v0.5 candidate | Preserved | Yes | Homebrew Available claim 없음. |
| Model workspace packs | Keep Experimental | Preserved | Yes | Global 또는 host-level availability claim 없음. |
| No-terminal method | Keep Experimental / assisted | Preserved | Yes | Deterministic validation claim 없음. |
| `ni run` | Bounded prompt compilation only | Preserved | Yes | Generated prompt execution 없음. |
| Benchmark evidence | Planning-artifact evidence with `not_measured` limits | Preserved | Yes | Causal, quality, adoption, cost, latency, downstream-performance claim 없음. |
| Fixture relock | Fixture relock is not project-root relock | Preserved | Yes | Validation fixture runs must remain separated. |
| Runtime execution boundary | `ni` is not a task runner, SPEC runner, execution layer, shell adapter, Codex exec adapter, queue, PR automation, release automation | Preserved | Yes | Runtime behavior added 없음. |

## Validation results

| Command | Required result | Packet result | Notes |
| --- | --- | --- | --- |
| `git status --short` | Expected docs-only status | Pass | docs/51과 docs/114 modified, docs/115 untracked가 보이고 unexpected tracked output은 없다. |
| `git ls-files docs/110_* ... docs/114_*` | docs/110 through docs/114 tracked as expected | Pass | docs/110 through docs/114 are tracked; docs/115 remains new until added. |
| `GOCACHE=/private/tmp/ni-go-cache go test ./...` | Pass | Pass | Check-only. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni status --dir . --proof --next-questions` | READY; blockers, deferrals, warnings: None | Pass | CLI reported READY with blockers, deferrals, warnings as None. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni --help` | Help renders | Pass | Downstream work execute 없음. |
| `GOCACHE=/private/tmp/ni-go-cache go run ./cmd/ni version` | Version output understood for build context | Pass | Output: `0.0.0-dev`. |
| `python3 scripts/check-install-docs.py` | Pass | Pass | Install claims bounded 유지. |
| `bash scripts/check-skill-packs.sh` | Pass | Pass | Model workspace packs Experimental 유지. |
| `bash scripts/demo-check.sh` | Pass | Pass | Downstream prompt execution 없음. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/quality.sh` | Pass | Pass | Broad quality wrapper. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/smoke.sh` | Pass | Pass | Fixture-only `ni end` / `ni relock`, not project root. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/install-check.sh` | Pass | Pass | Temporary/source install proof. |
| `GOCACHE=/private/tmp/ni-go-cache bash scripts/release-check.sh` | Pass | Pass | Check-only release gate. |
| `git diff -- .ni/contract.json .ni/session.json .ni/plan.lock.json` | No diff | Pass | Protected root planning files unchanged. |

## Changes made by this packet

- `docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.md`: English human approval
  packet.
- `docs/115_V0_5_PUBLICATION_HUMAN_APPROVAL_PACKET.ko.md`: Korean companion.
- `docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.md`: docs/115로 이어지는 narrow
  follow-up pointer.
- `docs/114_V0_5_RELEASE_PUBLICATION_CHECKLIST.ko.md`: matching Korean pointer.
- `docs/51_POST_RELEASE_ROADMAP.md`: docs/115 roadmap pointer.
- `docs/51_POST_RELEASE_ROADMAP.ko.md`: matching Korean pointer.

## What this packet proves

- exact human approval surface가 documented되었다
- manual release gates가 이 task와 분리되어 있다
- human decision options가 visible and unselected다
- known deferrals와 abort criteria가 explicit하다
- forbidden claims가 bounded로 남아 있다
- 이 packet은 release action을 수행하지 않았다

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

Why: packet이 ready with notes 상태이고, 다음 real step은 Codex 밖의 human decision이어야
한다. Release 또는 publication-prep action에는 explicit maintainer selection of one
human decision option이 필요하므로 여기서는 execution prompt를 제공하지 않는다.

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
