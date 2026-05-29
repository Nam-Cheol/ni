# Homebrew decision

Date: 2026-05-29

## Current distribution state

- Release binary: verified v0.3.0 GitHub Release archives와 checksums에 대해
  Available.
- Curl installer: verified v0.3.0 release assets에 대해 Available.
- Homebrew: Planned. Published 또는 tested tap/formula가 없다.
- Model workspace packs: broad product path로는 Experimental. Source packs,
  manual copy, zip packaging, Claude target-directory dry-run path는
  available이지만 global host discovery는 unverified다.
- No-terminal method: Experimental / assisted. Drafting은 도울 수 있지만
  deterministic readiness, lock, hash, prompt claims는 trusted runner의 CLI
  proof가 필요하다.

## Decision

Choose **B. Defer Homebrew to v0.5.**

Homebrew는 v0.4.1에서 Planned로 유지하고, v0.5보다 이르게 구현하지 않는
distribution infrastructure로 둔다.

## Rationale

Homebrew는 유용하지만 지금의 다음 adoption bottleneck은 아니다. Release binary와
curl installer는 Go 없이 `ni`를 사용해 보려는 사용자를 이미 지원하며, 두 경로는
verification evidence가 있다. Homebrew는 package-manager install을 선호하는
macOS/developer users에게 편의를 더하지만, planning, readiness proof, locking,
prompt compilation이라는 core product experience를 바꾸지는 않는다.

Homebrew의 blocking work는 external operational work다. Tap repository가 아직
available하지 않고, formula가 publish되지 않았으며, clean `brew install` path도
test되지 않았다. 지금 구현하면 conversation authoring UX, readiness proof clarity,
model workspace pack guidance, no-terminal proof capture, benchmark evidence 같은
더 영향이 큰 v0.4.1 stabilization work와 경쟁하게 된다.

`ni`는 pre-runtime Project Intent Compiler이므로 Homebrew는 `ni-kernel` behavior
밖에 있어야 한다. Package distribution은 사용자가 CLI를 얻도록 돕는 것이지 release
automation, execution state, adapter behavior, 또는 `ni run`을 prompt compilation
너머로 확장하는 이유가 아니다.

## User impact

Homebrew는 이미 Homebrew를 신뢰하고 package manager로 upgrade하기를 원하거나,
macOS CLI에 `brew install Nam-Cheol/tap/ni`를 기대하는 사용자에게 도움이 된다.
Tap과 formula가 verified된 뒤에는 그 audience의 install friction을 줄인다.

Script를 inspect하거나 checksum을 직접 verify할 수 있는 사용자는 이미 curl
installer와 release binary paths로 지원된다. Developers, evaluators,
contributors는 source와 local binary paths로 지원된다. Plan authoring 도움이
필요한 사용자는 지금은 conversation authoring, model workspace packs,
no-terminal assisted proof capture, benchmark evidence 개선에서 더 큰 이익을
얻는다.

## Required implementation work if Homebrew is chosen

Later task에서 Homebrew 구현을 선택한다면 반드시 다음 exact work를 수행한다:

1. Create or identify the public tap repository.
2. Choose the formula name.
3. Define the formula source URL.
4. Define the sha256 source from published release checksums.
5. Test `brew install`.
6. Update README/docs only after the tested install path works.

현재 intended tap은 다음으로 유지한다:

```text
Nam-Cheol/homebrew-tap
```

## Risks

- False package availability claims: formula가 존재하고 test되기 전에 README나
  install docs가 Homebrew가 동작한다고 암시할 수 있다.
- Stale formula checksums: release asset 또는 checksum update 뒤 formula가
  mismatched metadata를 가리킬 수 있다.
- Release asset naming drift: future archives의 name 또는 platform이 바뀌면
  naming contract가 안정적이지 않은 한 formula URLs가 깨질 수 있다.
- Package-manager maintenance burden: publish한 뒤에는 upgrades, checksums,
  audit fixes, working install commands를 사용자가 기대한다.
- Confusing Homebrew with ni-kernel behavior: package distribution은 repository
  infrastructure이지 Intent Lock Protocol behavior, runtime execution, 또는
  kernel-owned state가 아니다.

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

Tap, formula, checksums, audit, install test, published tap install test,
`ni --help` / `ni version` verification이 모두 존재하기 전에는
`Homebrew: Experimental` 또는 `Homebrew: Available`을 사용하지 않는다.

## Next task

Improve conversation-authoring UX proof capture for v0.4.1 stabilization.
