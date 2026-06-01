# Distribution Strategy

이 문서는 사용자가 Go를 몰라도 `ni`를 채택할 수 있는 경로와, 나중에 terminal을
직접 쓰지 않아도 되는 경로를 정의한다.

이 strategy는 방향성 문서다. 모든 경로가 오늘 available하다는 claim이 아니다.
현재 availability는 source, local build, local install, verified v0.4.0 release
archives, verified v0.4.0 curl installer path, repo-local model workspace
assistance다. Package managers는 external assets가 존재하고 verified되기 전까지
planned다. Future distribution work는 repository infrastructure, packaging,
documentation에 속한다. `ni-kernel` runtime execution behavior가 아니다.

`ni-kernel`의 authority는 계속 다음에 있다:

- docs contract 생성과 validation;
- `ni status`를 통한 readiness;
- `ni end`를 통한 lock 생성;
- `ni run`을 통한 prompt compilation.

Distribution automation은 사용자가 CLI를 얻거나 호출하는 것을 도울 수 있다.
하지만 `ni`를 task runner, shell adapter, package release runtime, hosted
execution service, multi-agent execution layer로 바꾸면 안 된다.

## Current Factual Status

| Path | Status | Notes |
| --- | --- | --- |
| Source / Go | Available | Developer path. |
| Local binary | Available | 이 checkout에서 build되는 local install path. |
| Release binary | Available | Verified v0.4.0 assets. |
| Curl installer | Available | v0.4.0 release assets에 대해 verified. |
| Model workspace packs | Experimental | UX layer; CLI remains authority; global host install은 overclaim하지 않는다. |
| No-terminal method | Experimental / assisted | Drafting only; deterministic validation에는 trusted runner의 CLI proof가 필요하다. |
| Homebrew | Planned / v0.5 candidate | Deferred이며 guaranteed가 아니다. Tap/formula가 존재하고 `brew install`, `ni --help`, `ni version`이 tested되기 전까지 Available이 아니다. |
| Runtime execution, shell adapters, Codex exec, queues, PR automation | Not included | `ni-kernel`의 일부가 아니다. Future downstream integration은 separate packages, target exports, seed formats로만 존재해야 한다. |

## Distribution Matrix

| Track | Status | User type | Required dependency | Trust model | Implementation work needed |
| --- | --- | --- | --- | --- | --- |
| Source mode | Available | Developers, early evaluators, contributors, Go에 익숙한 users | Go 1.22 이상; version metadata에는 Git optional | Checked-out source, local Go toolchain, repository tests, quality checks를 trust한다 | `go run`, `make build`, `make test`, `make quality`가 계속 documented and working 상태여야 한다 |
| Local binary mode | Available | 이 checkout에서 `./bin/ni` 또는 local install을 원하는 users | Go 1.22 이상; local shell | Checked-out source, local build, temporary install checks를 trust한다 | `make build`, `make install-local`, `bash scripts/install-check.sh`가 계속 working 상태여야 한다 |
| Release binary mode | Available | Terminal은 편하지만 Go는 모르는 users | Verified v0.4.0 release에서 OS별 downloaded `ni` binary | Verification 뒤 GitHub Releases assets와 published checksums를 trust한다 | Manual release asset process, OS/arch별 build, checksums, verification과 rollback docs를 최신으로 유지한다 |
| Curl installer mode | Available | One command를 원하는 terminal users | `curl` 또는 equivalent downloader; supported platforms의 POSIX shell | Installer script를 inspect한 뒤 trust한다. Installer는 checksum이 있으면 real release asset을 verify한다 | `install.sh`를 작고 auditable하게 유지하고, 새로운 availability claim 전에는 real release assets에 대해 다시 verify한다 |
| Package manager mode | Planned | Platform package managers를 선호하는 users | Homebrew first; Windows demand가 있으면 Scoop later | Official release assets를 가리키는 package manager metadata, formulas/manifests, checksums를 trust한다 | Release binaries가 안정화된 뒤 Homebrew tap 또는 formula 생성; Scoop은 later; publishing은 `ni-kernel` 밖에 둔다 |
| Model workspace mode | Experimental | Model workspace에서 planning을 author하는 Codex/Claude users | Repository docs를 읽고 authority로 `ni` CLI를 호출할 수 있는 model workspace | Model이 아니라 CLI gates를 trust한다; skills는 docs와 `.ni/contract.json` 위의 UX다 | Portable skill packs packaging과 docs; skill behavior는 `ni status`, `ni end`, `ni run`과 aligned해야 한다 |
| No-terminal mode | Experimental | Direct terminal use 없이 docs-first planning을 원하는 non-technical users, product leads, researchers, teams | Assisted docs-first workflow; trusted runner가 deterministic validation을 위해 `ni` gates를 호출해야 한다 | Visible docs, lockfile hashes, CLI-generated status/lock/run outputs를 trust한다; model은 readiness를 단독 선언할 수 없다 | Assisted no-terminal docs를 factual하게 유지한다; CLI proof 없이 deterministic validation을 claim하지 않는다 |

## Track Details

### 1. Source mode

Source mode는 available now이며, first-time local use를 위한 primary supported
distribution path다.

Users는 repository를 clone한 뒤 다음을 실행한다:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
```

Local binary도 build할 수 있다:

```bash
make build
./bin/ni --help
```

이 경로는 contributors, evaluators, 이미 Go가 설치된 users에게 적합하다. Trust
model은 transparent source와 local validation이다.

### 2. Local binary mode

Local binary mode는 available이다.

Users는 checked-out source에서 `ni`를 build and install할 수 있다:

```bash
make build
make install-local
```

이 경로는 release assets에 의존하지 않는다.

### 3. Release binary mode

Release binary mode는 verified v0.4.0 GitHub Release assets에 대해 available이다.

Users는 Go 없이 GitHub Releases에서 `ni`를 download할 수 있다. Documented trust
path는 matching archive와 checksum file을 같은 release에서 download하고,
checksum을 verify하고, binary를 unpack한 뒤 `ni --help`와 `ni version`을
실행하는 것이다.

이것은 repository distribution infrastructure다. `ni-kernel`에 release
automation을 추가하면 안 되며, `ni run`을 executor로 바꾸면 안 된다.

### 4. Curl installer mode

Curl installer mode는 verified v0.4.0 release assets에 대해 available이다.

Installer는 verified release asset을 download하는 작은 `install.sh`다.
기본적으로 source build를 하거나 downstream work를 실행하거나 trust boundary를
숨기지 않는다. Script는 무엇을 download하는지, 어디에 `ni`를 install하는지,
asset을 어떻게 verify하는지 설명한다.

이 track은 release binary mode에 의존한다. Script는 local fake release assets로
test되며 checksum mismatch behavior도 확인한다. 또한 2026-05-29에 real v0.4.0
darwin/arm64 release archive와 checksum file에 대해 verified되었다.

### 5. Package manager mode

Package manager mode는 planned and deferred이며, 아직 available하지 않다.

Homebrew가 첫 package manager 후보이다. Initial developer audience와 local macOS
usage에 맞기 때문이다. Scoop은 Windows demand가 나타나면 later에 검토한다.
Package definitions는 official release assets와 checksums를 가리켜야 한다.

Package publishing은 external repository infrastructure다. Intent Lock
Protocol의 일부가 아니며 kernel-owned execution state가 되면 안 된다.
Homebrew는 tap/formula가 존재하고 해당 package path에서 `brew install`,
`ni --help`, `ni version`이 tested된 뒤에만 Available이 될 수 있다.

### 6. Model workspace mode

Model workspace mode는 오늘 repo-local form의 experimental path다. Repository에는
planning, locking, prompt compilation을 위한 model-facing skill material이
있지만 CLI가 계속 authority다.

대상 user는 Codex, Claude 또는 비슷한 model workspace에서 작업하는 사람이다.
Model은 `docs/plan/**`과 `.ni/contract.json` authoring을 돕고, 그 뒤 deterministic
gates로 다음을 사용한다:

```bash
ni status
ni end
ni run
```

또는 source equivalents를 사용한다.

Portable skill packs는 planned distribution work다. UX와 instructions를
packaging해야 하며 readiness, lock, hash verification을 우회하면 안 된다.

### 7. No-terminal mode

No-terminal mode는 assisted docs-first method로 experimental이다.

Intended shape은 downloadable model pack과 docs-first workflow다. Terminal을
직접 쓰지 않는 user도 plan, status proof, lock proof, compiled handoff를 볼 수
있어야 한다. 오늘은 assisted planning method일 뿐이다. Trusted local 또는
workspace runner가 deterministic validation을 위해 `ni`를 호출해야 하며,
user-facing experience는 어떤 output이 CLI에서 왔는지 명확해야 한다.

이것은 terminal-less web service, hosted execution service, hidden agent runner가
아니다. Core contract는 계속 다음과 같다:

```text
planning conversation -> docs contract -> readiness gate -> lockfile -> prompt
```

## Availability Rules

- Source mode, local binary build, local install, repo-local model workspace
  assistance는 오늘 available 또는 experimental이라고 설명할 수 있다.
- Release binaries는 verified GitHub Release assets와 checksums에 대해
  available하다고 설명할 수 있으며, 현재 대상은 v0.4.0이다.
- Curl install은 `install.sh`로 real release assets와 checksums를 verify한
  대상에 대해 available하다고 설명할 수 있으며, 현재 대상은 v0.4.0이다.
- Package manager install은 packages 또는 formulas가 published되기 전까지
  available하다고 설명하면 안 된다.
- No-terminal mode는 downloadable model pack과 proof-oriented workflow가 있기
  전까지 assisted 또는 experimental로만 설명할 수 있다. CLI proof 없이
  deterministic validation으로 설명하면 안 된다.

## Boundary Rules

Distribution work는 scripts, release checklists, package metadata, checksums,
installer tests, docs, downloadable model packs를 추가할 수 있다.

Distribution work는 다음을 추가하면 안 된다:

- `ni run`의 shell adapter behavior;
- `ni-kernel` 내부의 Codex 또는 Claude execution adapters;
<!-- ni-boundary-allow: explicit negative boundary list item. -->
- kernel behavior로서의 release automation;
<!-- ni-boundary-allow: explicit negative boundary list item. -->
- queues, PR automation, execution evidence loops;
- 이 strategy의 일부인 terminal-less hosted service.

Repository는 나중에 builds와 publishing을 infrastructure로 automate할 수 있다.
그 automation은 `ni`의 runtime behavior 밖에 있어야 한다.
