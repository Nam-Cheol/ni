# Homebrew Tap Plan

Date: 2026-05-28

Scope: verified release binaries와 verified curl installer path 이후의
Homebrew distribution route를 계획하되, package-manager availability를 claim하지
않는다.

Current Homebrew status: Planned.

Evidence: `git ls-remote https://github.com/Nam-Cheol/homebrew-tap.git`는
2026-05-28에 "Repository not found"를 반환했다. Owner-confirmed tap,
published formula, tested `brew install` path가 없다.

## Route Decision

Homebrew distribution route는 별도 tap repository를 사용한다:

```text
Nam-Cheol/homebrew-tap
```

첫 formula는 tap이 실제로 존재하고 ownership이 확인된 뒤 tap repository에
manual로 추가한다. Repository-local
[Homebrew Formula Draft](71_HOMEBREW_FORMULA_DRAFT.md)는 future tap을 위한
non-published draft일 뿐이다.

GoReleaser Homebrew publishing은 아직 추가하지 않는다. GoReleaser-generated
formula update는 tap이 존재하고, owner가 GoReleaser로 tap을 update하기를
명시적으로 확인하고, brew configuration을 포함한 `goreleaser check`가 통과한
뒤의 later option이다.

## Route Matrix

| Route | Decision | Status |
| --- | --- | --- |
| Separate tap repository | Intended external tap으로 `Nam-Cheol/homebrew-tap`을 사용한다. | Planned; repository not found. |
| Manual formula draft | Tap에 copy되기 전까지 non-published draft로만 유지한다. | `docs/71_HOMEBREW_FORMULA_DRAFT.md`의 draft only. |
| GoReleaser-generated formula | Tap이 존재하고 owner가 GoReleaser 관리를 confirm할 때까지 defer한다. | Not configured. |

## Next Steps To Create The Tap

1. Public GitHub repository로 `Nam-Cheol/homebrew-tap`을 만든다.
2. Tap repository를 local에 clone한다.
3. `Formula/ni.rb`를 `ni-kernel`이 아니라 tap repository 안에 만든다.
4. `docs/71_HOMEBREW_FORMULA_DRAFT.md`에서 시작한다.
5. 모든 placeholder checksum을 `docs/70_RELEASE_VERIFICATION_v0.3.0.md` 또는
   published `ni_0.3.0_checksums.txt` release asset의 matching checksum으로
   교체한다.
6. Tap checkout에서 formula를 validate한다:

   ```bash
   brew audit --strict --online Formula/ni.rb
   brew install --build-from-source Formula/ni.rb
   ni --help
   ni version
   ```

7. Local formula validation이 동작한 뒤에만 published tap command를 test한다:

   ```bash
   brew install Nam-Cheol/tap/ni
   ni --help
   ni version
   ```

8. 그 exact published tap command가 동작한 뒤에만 README와 install docs를
   Planned에서 Available로 바꾼다.

## Availability Rule

Homebrew는 다음이 모두 true가 될 때까지 Planned로 남는다:

1. Tap repository가 존재한다.
2. Owner가 intended tap임을 confirm한다.
3. `Formula/ni.rb`가 real release checksums와 함께 tap에 publish된다.
4. Homebrew tooling이 formula를 validate한다.
5. Clean Homebrew environment에서 `brew install Nam-Cheol/tap/ni`가 동작한다.

그 전까지 Homebrew badge, README `brew install` command, package-manager
availability claim, package publishing automation, `ni-kernel` runtime behavior를
추가하지 않는다.
