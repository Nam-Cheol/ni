# Public Announcement Kit

이 문서는 `ni` public launch material에 붙여 넣을 수 있는 announcement copy를
모아 둔다. Social channel에 post하지 않고, release를 publish하지 않고, tag를
만들지 않고, asset을 upload하지 않고, agent를 실행하지 않고, distribution을
자동화하지 않는다.

## Core Copy

One-liner:

> Don't run the agent yet. Compile the intent first.

Short product description:

> ni turns planning conversations into locked project contracts before AI agents
> or teams start work.

Category:

`ni`는 AI Agents를 위한 Project Intent Compiler다. Downstream execution이
시작되기 전에 planning contract를 만들고, validate하고, lock하고, compile하는
pre-runtime intent contract layer다.

## Boundary Warnings

Public launch material에는 아래 warning을 유지한다:

- `ni`는 downstream work를 실행하지 않는다.
- `ni`는 task runner가 아니다.
- `ni`는 agent orchestration runtime이 아니다.
- `ni run`은 bounded handoff prompt만 compile한다.
- Model workspace packs는 UX layer다. Readiness, lock creation, hash
  validation, prompt compilation의 authority는 CLI에 있다.

## Install Paths

현재 public install paths는 [Install ni](22_INSTALL.md)와
[Curl Installer](install-curl.md)에 정리되어 있다.

### Release Binary

Status: verified v0.3.0 GitHub Release archives와 checksums에 대해 Available.

Platform에 맞는 archive와 `ni_0.3.0_checksums.txt`를 아래 release에서 받는다:

<https://github.com/Nam-Cheol/ni/releases/tag/v0.3.0>

Checksum을 verify하고 binary를 unpack한 뒤 아래 명령으로 확인한다:

```bash
ni --help
ni version
```

### Curl Installer

Status: verified v0.3.0 release assets에 대해 Available.

실행 전에 installer를 inspect한다:

```bash
VERSION="0.3.0"
curl -fsSLO https://raw.githubusercontent.com/Nam-Cheol/ni/main/install.sh
sed -n '1,320p' install.sh
sh install.sh --dry-run --version "$VERSION"
BINDIR="$HOME/.local/bin" sh install.sh --version "$VERSION"
"$HOME/.local/bin/ni" --help
"$HOME/.local/bin/ni" version
```

### Source

Status: Available.

`ni`를 개발하거나 checkout에서 CLI를 바로 확인할 때 source mode를 사용한다:

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
go run ./cmd/ni status --dir .
```

Source와 local build path에는 Go 1.22 이상이 필요하다.

### Model Workspace Packs

Status: source packs, verified source path에서 manual copy, Codex zip
packaging, Claude zip packaging, Claude target-directory dry-run install은
Available. Broad product path로는 Experimental이다. Global host discovery와
global install locations가 아직 verified 상태가 아니기 때문이다.

자세한 내용은
[Model Pack Install Verification](75_MODEL_PACK_INSTALL_VERIFICATION.md)을 본다.
Packs는 `ni`를 설치하지 않고, model API를 호출하지 않고, downstream execution을
실행하지 않고, CLI proof를 대체하지 않는다.

## Announcement Drafts

### GitHub Release Note Summary

`ni` v0.3.0은 AI Agents를 위한 Project Intent Compiler를 소개한다. Agent나 팀이
작업을 시작하기 전에 planning conversation을 checked, locked project contract로
바꾸는 작은 CLI다.

Tagline: Don't run the agent yet. Compile the intent first.

이 release에는 `ni init`, `ni status`, `ni end`, `ni run`, lockfile hash
validation, status proof output, install documentation, verified release binary
및 curl installer path, experimental model workspace packs가 포함된다.

Boundary: `ni`는 downstream work를 실행하지 않는다. Task runner도 아니고 agent
orchestration runtime도 아니다. `ni run`은 bounded handoff prompt만 compile한다.

### X/Twitter Short Post

Don't run the agent yet. Compile the intent first.

`ni`는 AI agents나 teams가 work를 시작하기 전에 planning conversations을 locked
project contracts로 바꿉니다.

v0.3.0에는 core CLI, verified release binary와 curl install paths, experimental
model workspace packs가 포함됩니다.

Task runner가 아닙니다. Orchestration runtime도 아닙니다.

### LinkedIn Post

`ni`를 공유합니다. `ni`는 AI Agents를 위한 Project Intent Compiler입니다.

아이디어는 단순합니다. Coding agent, workflow, team이 작업을 시작하기 전에
planning conversation을 checked and locked project contract로 바꿉니다.

`ni`는 이 흐름에 deterministic pre-runtime layer를 둡니다. Planning workspace를
init하고, readiness를 review하고, accepted intent를 lock하고, hash를 validate하고,
bounded handoff prompt를 compile합니다.

One-liner: Don't run the agent yet. Compile the intent first.

`ni` v0.3.0에는 verified release binary와 curl installer paths, source usage,
experimental model workspace packs가 있습니다. `ni`는 downstream work를 실행하지
않고, task runner나 agent orchestration runtime이 아닙니다.

### Hacker News / Reddit Style Post

`ni`는 execution이 시작되기 전에 planning을 compile 대상으로 다루는 작은 CLI다.

Planning conversation을 project contract로 바꾸고, contract가 ready인지
확인하고, accepted intent를 hash와 함께 lock하고, downstream agent나 team에 넘길
짧은 handoff prompt를 compile한다.

Product boundary는 의도적으로 좁다. `ni`는 downstream work를 실행하지 않는다.
Task runner도 아니고 orchestration runtime도 아니다. 첫 release는 kernel인
`init`, `status`, `end`, `run`에 집중한다.

One-liner: Don't run the agent yet. Compile the intent first.

### Korean Community Post

`ni`를 공개합니다.

`ni`는 AI agent나 팀이 작업을 시작하기 전에 planning conversation을 locked project
contract로 바꾸는 Project Intent Compiler입니다.

핵심 문장은 이겁니다:

> Don't run the agent yet. Compile the intent first.

`ni` v0.3.0은 `init`, `status`, `end`, `run` 중심의 작은 CLI입니다. Readiness를
확인하고, accepted intent를 lock하고, hash를 검증하고, downstream agent나 팀에
넘길 짧은 handoff prompt를 compile합니다.

주의할 점도 명확합니다. `ni`는 downstream work를 실행하지 않습니다. Task runner가
아니고, agent orchestration runtime도 아닙니다.

## FAQ

### Is ni a replacement for coding agents?

아니다. `ni`는 coding agents나 human teams 앞단에 놓인다. Downstream actors가
trust할 수 있는 intent를 알 수 있도록 planning contracts를 만들고, validate하고,
lock하고, compile한다.

### Can I use it without Go?

가능하다. Verified v0.3.0 release binary나 verified curl installer path를 사용하면
Go 없이 사용할 수 있다. Source usage와 local build에는 Go 1.22 이상이 필요하다.

### Can I use it without terminal?

부분적으로 가능하다. Model workspace packs는 model workspace 안에서 planning
docs를 draft하는 데 도움을 줄 수 있다. 하지만 deterministic readiness, lock
creation, hash validation, prompt compilation은 현재 CLI proof가 필요하다.

### What does ni lock?

`ni end`는 accepted planning state를 `.ni/plan.lock.json`에 lock한다. Lock은
contract, planning docs, accepted plan material의 hashes를 기록해서 downstream
actors가 intent change를 감지할 수 있게 한다.

### What happens after ni-run?

`ni run`은 locked plan에서 4000자 이하의 handoff prompt를 compile한다. 그 다음
downstream agent, team, workflow가 그 prompt를 사용할 수 있지만, `ni`는 그 work를
실행하지 않는다.
