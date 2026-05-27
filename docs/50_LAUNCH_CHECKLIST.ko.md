# Public Launch Checklist

이 checklist는 `ni`를 public하게 공유하기 전의 deterministic gate다. Repository
link 공유, announcement draft, maintainer review를 준비하지만 package publish,
GitHub release 생성, tag 생성, social posting, downstream agent 실행은 하지
않는다.

Release-readiness gate 이후, public sharing 전에 실행한다:

```bash
bash scripts/launch-check.sh
```

## Positioning

- [ ] README tagline이 `Project Intent Compiler for AI Agents`를 유지한다.
- [ ] README가 `ni-kernel`을 deterministic pre-runtime control layer로 설명한다.
<!-- ni-boundary-allow: not a claim; explicit negative boundary checklist item. -->
- [ ] README가 `ni`는 execution harness, task runner, SPEC runner,
      multi-agent execution layer, queue, shell adapter, Codex adapter, PR
      automation, release automation이 아니라고 말한다.
- [ ] Positioning links는 [docs/40_POSITIONING.md](40_POSITIONING.md)와
      [docs/42_INTENT_LOCK_PROTOCOL.md](42_INTENT_LOCK_PROTOCOL.md)를 가리킨다.

## README Rendering

- [ ] `README.md`의 Markdown fences가 balanced 상태로 render된다.
- [ ] README의 public demo commands는 `go run`, `bash scripts/demo-check.sh`,
      `bash scripts/quality.sh`, 또는 `bash scripts/install-check.sh`를 사용한다.
- [ ] README가 install, security, release-readiness, benchmark, target story,
      launch checklist docs로 link한다.
- [ ] README가 hosted binaries, package-manager distribution, package publishing을
      claim하지 않는다.

## README.ko Parity

- [ ] Korean companion docs가 유지되는 동안 `README.ko.md`가 존재한다.
- [ ] `README.ko.md`가 Korean release-readiness와 launch checklist docs로 link한다.
- [ ] `README.ko.md`가 `README.md`와 같은 launch-sensitive claims를 유지한다:
      source-first distribution, MIT license, CI, security policy, demo
      verification, install/build verification, benchmark protocol, no execution
      runtime.

## License

- [ ] `LICENSE`가 존재한다.
- [ ] README와 README.ko가 repository가 [MIT License](../LICENSE)에 따라
      licensed된다고 말한다.
- [ ] Launch material이 추가 commercial, hosted, enterprise license를 암시하지
      않는다.

## Security Policy

- [ ] `SECURITY.md`가 존재한다.
- [ ] Korean companion docs가 유지되는 동안 `SECURITY.ko.md`가 존재한다.
- [ ] README가 `SECURITY.md`로 link한다.
- [ ] Security docs가 private reporting channels, runtime sandboxing,
      downstream agent security, enterprise support를 claim하지 않는다.

## CI

- [ ] `.github/workflows/ci.yml`이 존재한다.
- [ ] CI가 `push`와 `pull_request`에서 실행된다.
- [ ] CI가 `go test ./...`, `bash scripts/quality.sh`,
      `bash scripts/smoke.sh`를 실행한다.
- [ ] Public docs가 CI에서 release, publish, package, downstream-runtime
      automation을 제공한다고 claim하지 않는다.

## Demo Verification

- [ ] [docs/48_DEMO_VERIFICATION.md](48_DEMO_VERIFICATION.md)가 존재한다.
- [ ] `bash scripts/demo-check.sh`가 pass한다.
- [ ] Public demos가 downstream agents를 실행하지 않고 `BLOCKED`, `READY`, 또는
      `READY_WITH_DEFERRALS` states를 증명한다.
- [ ] Demo outputs는 source tree에서 재현 가능하며 hosted services를 요구하지
      않는다.

## Benchmark Protocol

- [ ] [docs/43_BENCHMARK_PROTOCOL.md](43_BENCHMARK_PROTOCOL.md)가 존재한다.
- [ ] Benchmark docs가 benchmark는 execution benchmark가 아니라고 말한다.
- [ ] Benchmark docs가 downstream agents를 실행하면 안 된다고 말한다.
- [ ] Benchmark docs가 direct-to-agent comparisons와 runtime performance claims를
      분리한다.

## Release Notes

- [ ] [docs/68_RELEASE_NOTES_v0.3.0.md](68_RELEASE_NOTES_v0.3.0.md)가 존재한다.
- [ ] Release notes draft가 factual and source-first 상태다.
- [ ] Draft가 release publish, tag 생성, binary upload, package-manager
      availability claim을 하지 않는다고 말한다.
- [ ] Draft가 included kernel capability claims와 not-included runtime, adapter,
      automation, binary-package scope를 분리한다.
- [ ] [docs/69_MANUAL_RELEASE_STEPS.md](69_MANUAL_RELEASE_STEPS.md)가 존재하며
      tag push, asset confirmation, checksum confirmation, README install-status
      updates를 manual step으로 유지한다.

## Install/Build Verification

- [ ] [docs/22_INSTALL.md](22_INSTALL.md)가 존재한다.
- [ ] `bash scripts/install-check.sh`가 pass한다.
- [ ] Source, local build, temporary local install commands가 작동한다.
- [ ] Install docs가 Homebrew, GoReleaser, package-manager, hosted binary
      availability를 claim하지 않는다.

## Issue Templates

- [ ] `.github/ISSUE_TEMPLATE/bug_report.md`가 존재한다.
- [ ] `.github/ISSUE_TEMPLATE/feature_request.md`가 존재한다.
- [ ] `.github/ISSUE_TEMPLATE/boundary_question.md`가 존재한다.
<!-- ni-boundary-allow: not a claim; explicit negative boundary checklist item. -->
- [ ] Issue templates가 reporters에게 ni-kernel boundary를 보존하고
      downstream-runtime, task-runner, queue, PR automation, release automation
      requests를 드러내도록 요구한다.

## No False Claims

- [ ] Public docs가 published binary release를 claim하지 않는다.
- [ ] Public docs가 package publishing, package-manager support, Homebrew
      support를 claim하지 않는다.
- [ ] Public docs가 hosted service availability를 claim하지 않는다.
<!-- ni-boundary-allow: not a claim; explicit negative boundary checklist item. -->
- [ ] Public docs가 release automation을 claim하지 않는다.
- [ ] Public docs가 downstream execution, model API execution, shell execution,
      queue execution, multi-agent orchestration을 claim하지 않는다.

## No Execution-Runtime Drift

- [ ] `ni run`은 여전히 prompt compilation only로 documented된다.
- [ ] Target exports는 여전히 inert downstream seed material로 documented된다.
- [ ] 어떤 launch item도 Codex, Claude, model APIs, shells, downstream agents,
      queues, adapters, generated harnesses 실행을 요구하지 않는다.
- [ ] Execution 관련 future public launch blocker는 deterministic validation,
      locking, bounded prompt compilation, inert seed generation으로 환원되지
      않는 한 `ni-kernel` 밖에 속한다.

## Local Gate

Deterministic launch gate를 실행한다:

```bash
bash scripts/launch-check.sh
```

이 gate는 repository에 이미 존재하는 local validation scripts만 호출할 수 있다:
`quality.sh`, `smoke.sh`, `demo-check.sh`, `release-check.sh`,
`install-check.sh`. Publishing, GitHub release 생성, tag 생성, announcements
posting, package upload, downstream agent 실행은 절대 하지 않는다.
