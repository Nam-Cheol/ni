# ni v0.4.0 — Conversation Authoring Hardening

Date: 2026-05-29

Status: draft release plan only. 이 문서는 release를 publish하지 않고, tag를
create하거나 push하지 않고, GitHub Release를 만들지 않고, assets를 upload하지
않고, Homebrew를 Available로 표시하지 않고, runtime execution을 추가하지
않는다.

## Version Decision

현재 repository tag는 `v0.3.0`과 `v0.1.0`이다. `v0.4.0` tag가 없으므로 다음
계획 release는 `v0.4.0`이다. 이 계획을 사용하기 전에 `v0.4.0` tag가 생기면
대신 `v0.4.1`을 준비한다.

## Release Summary

`ni v0.4.0`은 설치 이후 adoption path를 harden한다. 핵심 개선은 runtime
execution이 아니라 conversation-to-contract authoring quality다.

Release story는 다음과 같다:

```text
ni init -> ni-start conversation -> docs/contract/session updates -> ni status --proof --next-questions
```

`ni`는 계속 `ni-kernel`이다: AI Agents를 위한 pre-runtime Project Intent
Compiler. Kernel은 docs contract, readiness gate, lockfile, prompt compiler,
source-of-truth rule을 소유한다. Downstream work를 실행하지 않는다.

## Included Changes

### Conversation authoring

- Fresh `R014`, `R015`, `R016` blockers를 위한 first-run `ni-start`
  conversation card.
- Docs, contract, session update 이후 exact CLI readiness basis를 보여주는
  planning proof capture.
- Planning diff를 minimal, visible, conversation-tied 상태로 유지하는 model
  edit safety 개선. Ambiguous statements는 확인 전까지 assumptions, drafts,
  open questions로 남긴다.

### Readiness proof

- `R014`, `R015`, `R016`에 대한 concrete first-run proof text.
- `ni status --proof`의 blocker `Why it matters`와 `Next` wording 개선.
- `ni status --proof --next-questions`는 blocked planning turn 이후 visible
  explanation loop로 유지된다.

### Docs/contract sync

- Docs와 `.ni/contract.json`의 purpose drift를 나타내는 `SYNC-014`.
- Docs와 contract narrative records의 actor/outcome drift를 나타내는
  `SYNC-015`.
- Docs와 contract의 delivery-surface drift를 나타내는 `SYNC-016`.
- Proof output에 stable `ID`, `Location`, `Problem`, `Why it matters`,
  `Suggested repair`, `Blocks ni-end` fields를 가진 sync diagnostics.
- JSON proof/status output은 affected `R012` `issues[]`와 `proof[]` entries에
  `sync_diagnostic`을 포함할 수 있다.

### Next questions

- 다음 planning turn을 위한 grouped question output.
- Repair questions가 낮은 가치의 generic prompts를 shadow하거나 반복하지
  않도록 prioritization과 deduplication.
- Purpose, actors/outcomes, delivery surface를 위한 first-run card questions.
- `SYNC-014`, `SYNC-015`, `SYNC-016`을 위한 sync repair questions.
- Risk decisions, evaluation evidence, scope boundaries, open blockers groups.

### Model workspace packs

- Grouped questions가 `ni-start` skills와 packaged workspace guidance에
  연결된다.
- 규칙은 유지된다: Skills are UX; CLI is authority.
- Packs는 downstream execution, runtime readiness, second readiness engine을
  claim하면 안 된다.

### Examples

- Example coverage matrix와 demo expectations 개선.
- Companion docs가 유지되는 영역의 Korean companion docs.
- First-run cards, grouped questions, docs/contract/session updates,
  re-status loops를 보여주는 `ni-start` dogfood updates.
- Trusted CLI proof 이전 drafting만 다루는 no-terminal assisted flow hardening.
- Fake empirical claims나 model API calls 없이 intent-readiness benchmark case
  study expansion.

### Distribution

- Release binary는 release asset과 checksum verification 이후 Available로
  남는다.
- Curl installer는 새 release에 대한 verification 이후 Available로 남는다.
- Homebrew는 별도 구현과 테스트가 끝나기 전까지 Planned이며 v0.5로 deferred
  상태를 유지한다.

## Not Included

- No task runner.
- No SPEC runner.
- No execution harness.
- No Codex exec adapter.
- No shell adapter.
- No downstream agents.
- No queue.
- No PR automation.
- No release automation inside `ni-kernel`.
- No Homebrew availability in this release unless separately implemented and
  tested.

이 exclusions는 release boundaries다. `ni-kernel`이 runtime이 아니라 Intent
Lock Protocol에 집중하도록 유지한다.

## Compatibility Notes

- `ni init`, `ni status`, `ni end`, `ni run` CLI command names는 stable하게
  유지된다.
- Human-readable `ni status --proof`와 `--next-questions` output은 바뀌었다:
  proof text가 더 구체적이고, next questions가 grouped이며, first-run 또는
  sync-repair cards가 snapshots/docs expectations에 영향을 줄 수 있다.
- JSON output은 matching flags가 요청될 때 additive shape changes를 가진다:
  `next_questions`는 `--next-questions`가 있을 때만 있고, `proof`는
  `--proof`가 있을 때만 있으며, `sync_diagnostic`은 `R012` `issues[]`와
  `proof[]` entries에 나타날 수 있다.
- Human-only proof grouping, `Why it matters`, `Next` wording은 proof JSON
  schema에 추가되지 않는다.
- `ni run`은 계속 prompt compiler only이며 bounded prompt output을 만든다.
  Shells, agents, queues, adapters, downstream work를 실행하지 않는다.

## Validation Checklist

Tagging 또는 publishing 전에 다음 evidence를 수집한다:

- [ ] `go test ./...`
- [ ] `bash scripts/quality.sh`
- [ ] `bash scripts/smoke.sh`
- [ ] `bash scripts/demo-check.sh`
- [ ] `bash scripts/install-check.sh`
- [ ] `bash scripts/release-check.sh`
- [ ] `bash scripts/fresh-install-check.sh` if present
- [ ] `bash scripts/check-skill-packs.sh`
- [ ] package Claude skills: `bash scripts/package-claude-skills.sh`
- [ ] package Codex skills: `bash scripts/package-codex-skills.sh`
- [ ] release dry run before tag: `bash scripts/release-dry-run.sh`

이 checklist는 evidence collection이다. Release assets를 publish하거나,
tags를 push하거나, Homebrew tap을 update하거나, unverified distribution path를
Available로 표시하지 않는다.

## Manual Release Steps

1. Git tree가 clean한지 확인하고 release plan이 intended release version과
   맞는지 확인한다.
2. Validation checklist를 실행하고 중요한 proof output을 보존한다.
3. Tag를 만들기 전에 release dry run을 실행한다.
4. `v0.4.0` annotated tag를 만든다.
5. Tag를 push한다.
6. GitHub Actions release workflow가 끝날 때까지 기다린다.
7. Release assets와 checksums를 verify한다.
8. Release assets의 current-platform binary를 verify한다.
9. Curl installer를 새 release에 대해 verify한다.
10. Availability status가 바뀌는 경우에만 그 이후 docs를 update한다.

이 plan을 준비하는 동안 manual release steps를 실행하지 않는다.

## Release Risk Notes

- Proof wording changes는 status output을 quote하는 snapshots,
  walkthroughs, docs에 영향을 줄 수 있다.
- Examples와 docs는 grouped next-question output과 계속 aligned되어야 한다.
- Model workspace packs는 UX이지 authority가 아니다. Readiness, lock, prompt
  compilation은 CLI gates만 결정한다.
- Homebrew는 Planned로 남으므로 docs가 package-manager availability를 암시하면
  안 된다.
- No-terminal은 assisted일 뿐 deterministic하지 않다. Readiness에는 여전히
  trusted CLI proof가 필요하다.

## Next Release Direction

`v0.5` candidates:

- 여전히 유용하다면 Homebrew decision과 implementation.
- 더 강한 benchmark evidence와 더 투명한 case-study reporting.
- Software-only planning을 넘어서는 더 많은 product surfaces와 examples.
- Conversation authoring reliability의 지속적 개선.
- Installation과 adoption surfaces가 안정화된 뒤 optional landing page.
- Downstream integrations는 separate packages로만 다루고, `ni-kernel`
  execution behavior로 만들지 않는다.
