# Homebrew Decision

Date: 2026-05-29

Decision: **B. Homebrew tap implementation은 v0.5로 defer한다.**

Current Homebrew status: Planned.

## Reviewed Inputs

- [Homebrew Tap Plan](72_HOMEBREW_TAP_PLAN.md)
- [Homebrew Tap Plan, Korean companion](72_HOMEBREW_TAP_PLAN.ko.md)
- [Homebrew Tap Setup](74_HOMEBREW_TAP_SETUP.md)
- v0.3.0 release binary status: release asset과 checksum verification 이후
  Available.
- v0.3.0 curl installer status: real release assets에 대한 verification 이후
  Available.

## Rationale

필수 external tap work가 아직 끝나지 않았으므로 이 task에서 Homebrew를 구현하지
않는다. Intended tap은 계속 `Nam-Cheol/homebrew-tap`이지만, tap이 존재하고,
`Formula/ni.rb`가 real checksums와 함께 publish되고, Homebrew audit이 통과하며,
clean Homebrew environment에서 `brew install Nam-Cheol/tap/ni`가 test되기 전까지
Homebrew는 Planned로 남아야 한다.

구현을 defer하면 현재 release story를 factual하게 유지할 수 있다. Source, local
binary, release binary, curl installer paths는 Available이고, package-manager
distribution은 아직 unverified다. 또한 package publishing은 `ni-kernel` runtime
behavior 밖에 남는다.

## What This Decision Does

- Homebrew status를 Planned로 유지한다.
- Existing tap target을 `Nam-Cheol/homebrew-tap`으로 유지한다.
- `docs/74_HOMEBREW_TAP_SETUP.md`를 later implementation task의 setup procedure로
  유지한다.
- Homebrew tap implementation을 v0.5 distribution infrastructure work로
  schedule한다. Kernel runtime behavior가 아니다.

## What This Decision Does Not Do

- Homebrew tap을 create 또는 modify하지 않는다.
- Public install docs에 `brew install` instructions를 추가하지 않는다.
- Homebrew를 Available로 표시하지 않는다.
- Package-manager automation, release publishing, tags, runtime execution
  behavior를 추가하지 않는다.

## Availability Gate

Homebrew는 다음이 모두 true가 된 뒤에만 Available이 될 수 있다:

1. `Nam-Cheol/homebrew-tap`이 intended public tap으로 존재하고 confirm된다.
2. `Formula/ni.rb`가 그 tap에 publish된다.
3. Formula URLs가 official release assets를 가리킨다.
4. Formula checksums가 published release checksum source와 일치한다.
5. `brew audit --strict --online Formula/ni.rb`가 통과한다.
6. `brew install --build-from-source Formula/ni.rb`가 통과한다.
7. Clean Homebrew environment에서 `brew install Nam-Cheol/tap/ni`가 통과한다.
8. Homebrew로 설치된 binary에서 `ni --help`와 `ni version`이 통과한다.

그 전까지 README, install docs, release docs, launch material은 Homebrew를
Planned로 유지해야 한다.

