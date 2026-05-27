# Distribution Strategy

이 문서는 사용자가 Go를 몰라도 `ni`를 채택할 수 있는 경로와, 나중에 terminal을
직접 쓰지 않아도 되는 경로를 정의한다.

이 strategy는 방향성 문서다. 모든 경로가 오늘 available하다는 claim이 아니다.
현재 availability는 source-first다. Future distribution work는 repository
infrastructure, packaging, documentation에 속한다. `ni-kernel` runtime execution
behavior가 아니다.

`ni-kernel`의 authority는 계속 다음에 있다:

- docs contract 생성과 validation;
- `ni status`를 통한 readiness;
- `ni end`를 통한 lock 생성;
- `ni run`을 통한 prompt compilation.

Distribution automation은 사용자가 CLI를 얻거나 호출하는 것을 도울 수 있다.
하지만 `ni`를 task runner, shell adapter, package release runtime, hosted
execution service, multi-agent execution layer로 바꾸면 안 된다.

## Distribution Matrix

| Track | Status | User type | Required dependency | Trust model | Implementation work needed |
| --- | --- | --- | --- | --- | --- |
| Source mode | Available | Developers, early evaluators, contributors, Go에 익숙한 users | Go 1.22 이상; version metadata에는 Git optional | Checked-out source, local Go toolchain, repository tests, quality checks를 trust한다 | `go run`, `make build`, `make test`, `make quality`가 계속 documented and working 상태여야 한다 |
| Release binary mode | Next | Terminal은 편하지만 Go는 모르는 users | Terminal; OS별 downloaded `ni` binary | 구현된 뒤 GitHub Releases assets와 published checksums 또는 signatures를 trust한다 | Manual release asset process 정의, OS/arch별 build, checksums publish, verification과 rollback 문서화 |
| Curl installer mode | Planned | Go setup 없이 one command를 원하는 terminal users | `curl` 또는 equivalent downloader; supported platforms의 POSIX shell | Installer script가 downloaded release assets를 published checksums 또는 signatures로 verify할 때만 trust한다 | Release binaries 이후에만 `install.sh` 추가; script는 작고 auditable해야 하며 installer checks로 보호해야 한다 |
| Package manager mode | Planned | Platform package managers를 선호하는 users | Homebrew first; Windows demand가 있으면 Scoop later | Official release assets를 가리키는 package manager metadata, formulas/manifests, checksums를 trust한다 | Release binaries가 안정화된 뒤 Homebrew tap 또는 formula 생성; Scoop은 later; publishing은 `ni-kernel` 밖에 둔다 |
| Model workspace mode | Available in repo-local form; portable packs planned | Model workspace에서 planning을 author하는 Codex/Claude users | Repository docs를 읽고 authority로 `ni` CLI를 호출할 수 있는 model workspace | Model이 아니라 CLI gates를 trust한다; skills는 docs와 `.ni/contract.json` 위의 UX다 | Portable skill packs packaging과 docs; skill behavior는 `ni status`, `ni end`, `ni run`과 aligned해야 한다 |
| No-terminal mode | Planned | Direct terminal use 없이 docs-first planning을 원하는 non-technical users, product leads, researchers, teams | Downloadable model pack과 docs-first workflow; trusted runner가 뒤에서 `ni` gates를 호출해야 한다 | Visible docs, lockfile hashes, CLI-generated status/lock/run outputs를 trust한다; model은 readiness를 단독 선언할 수 없다 | Downloadable model pack, guided docs workflow, proof display 설계; 이 task에서는 hosted service나 terminal-less web runtime을 추가하지 않는다 |

## Track Details

### 1. Source mode

Source mode는 available now이며, first-time local use를 위한 유일한 fully
supported distribution path다.

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

### 2. Release binary mode

Release binary mode는 next이며, 아직 available하지 않다.

목표는 users가 Go를 설치하지 않고 GitHub Releases에서 `ni`를 download할 수 있게
하는 것이다. Release assets는 plain OS/architecture binaries 또는 archives와
checksums여야 한다. Verification이 문서화되기 전에는 supported path라고
표현하지 않는다.

이것은 repository distribution infrastructure다. `ni-kernel`에 release
automation을 추가하면 안 되며, `ni run`을 executor로 바꾸면 안 된다.

### 3. Curl installer mode

Curl installer mode는 planned이며, 아직 available하지 않다.

Installer는 verified release asset을 download하는 작은 `install.sh`여야 한다.
기본적으로 source build를 하거나 downstream work를 실행하거나 trust boundary를
숨기면 안 된다. Script는 무엇을 download하는지, 어디에 `ni`를 install하는지,
asset을 어떻게 verify하는지 설명해야 한다.

이 track은 release binary mode에 의존한다. Release assets와 verification
metadata가 있기 전에는 추가하지 않는다.

### 4. Package manager mode

Package manager mode는 planned이며, 아직 available하지 않다.

Homebrew가 첫 package manager 후보이다. Initial developer audience와 local macOS
usage에 맞기 때문이다. Scoop은 Windows demand가 나타나면 later에 검토한다.
Package definitions는 official release assets와 checksums를 가리켜야 한다.

Package publishing은 external repository infrastructure다. Intent Lock
Protocol의 일부가 아니며 kernel-owned execution state가 되면 안 된다.

### 5. Model workspace mode

Model workspace mode는 오늘 repo-local form으로만 available하다. Repository에는
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

### 6. No-terminal mode

No-terminal mode는 planned이며, 아직 available하지 않다.

Intended shape은 downloadable model pack과 docs-first workflow다. Terminal을
직접 쓰지 않는 user도 plan, status proof, lock proof, compiled handoff를 볼 수
있어야 한다. Trusted local 또는 workspace runner가 `ni`를 호출할 수 있지만,
user-facing experience는 어떤 output이 CLI에서 왔는지 명확해야 한다.

이것은 terminal-less web service, hosted execution service, hidden agent runner가
아니다. Core contract는 계속 다음과 같다:

```text
planning conversation -> docs contract -> readiness gate -> lockfile -> prompt
```

## Availability Rules

- Source mode, local binary build, local install, repo-local model workspace
  usage만 오늘 available하다고 설명할 수 있다.
- Release binaries는 supported platforms용 GitHub Releases assets가 있기 전까지
  available하다고 설명하면 안 된다.
- Curl install은 `install.sh`가 존재하고 release assets를 verify하기 전까지
  available하다고 설명하면 안 된다.
- Package manager install은 packages 또는 formulas가 published되기 전까지
  available하다고 설명하면 안 된다.
- No-terminal mode는 downloadable model pack과 proof-oriented workflow가 있기
  전까지 available하다고 설명하면 안 된다.

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
