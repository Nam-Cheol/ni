<p align="center">
  <img src="assets/hero.svg" alt="ni hero banner: Project Intent Compiler for AI Agents" width="100%">
</p>

<h1 align="center">agent를 아직 실행하지 마라. 먼저 의도를 컴파일하라.</h1>

<p align="center"><strong>ni는 AI agents나 teams가 work를 시작하기 전에 planning conversation을 locked project contract로 바꾼다.</strong></p>

<p align="center">
  <a href="README.md" aria-label="Read in English"><img alt="English" src="assets/badge-english.svg" width="84" height="28"></a>
  <a href="README.ko.md" aria-label="한국어로 읽기"><img alt="한국어" src="assets/badge-korean.svg" width="84" height="28"></a>
</p>

<p align="center">
  <a href="LICENSE"><img alt="License MIT" src="https://img.shields.io/badge/license-MIT-f4b860"></a>
  <a href=".github/workflows/ci.yml"><img alt="CI workflow exists" src="https://img.shields.io/badge/CI-workflow%20exists-25334a"></a>
  <a href="SECURITY.md"><img alt="Security policy exists" src="https://img.shields.io/badge/security-policy%20exists-2d5a52"></a>
  <a href="docs/00_START_HERE.md"><img alt="Docs index exists" src="https://img.shields.io/badge/docs-index%20exists-5b8def"></a>
</p>

<p align="center">
  <sub>Trust signals: MIT license; CI workflow; security policy; docs index.</sub>
</p>

<p align="center">
  <a href="#why-ni"><img src="assets/card-why.svg" alt="Why ni: prompt가 users, risks, non-goals, acceptance, blockers를 숨길 수 있다." width="32%"></a>
  <a href="#60초-시작"><img src="assets/card-start.svg" alt="Start path: initialize, readiness check, intent lock, prompt compile." width="32%"></a>
  <a href="#다음에-읽을-것"><img src="assets/card-docs.svg" alt="Docs map: protocol, commands, target boundaries, benchmark, launch notes." width="32%"></a>
</p>

## Why ni

Agents는 code ability 부족보다 unclear intent 때문에 더 자주 실패한다.

`ni`는 Project Intent Compiler다. Execution이 시작되기 전, vague goals가 hidden
assumptions로 바뀌는 지점에 선다:

```text
planning conversation -> explicit contract -> readiness gate -> locked plan -> bounded prompt or seed
```

1. AI agents는 너무 일찍 실행된다.
2. `ni`는 ambiguous execution을 block한다.
3. `ni`는 intent를 locked project contract로 compile한다.
4. 그 뒤 humans, models, tools가 그 contract를 기준으로 work할 수 있다.

Payoff: `ni`는 unclear intent를 visible하게 만들고, unsafe handoff를 block하며,
locked plan에서 bounded prompt 또는 seed를 만든다.

더 깊은 product story는 [Why ni exists](docs/product-story.ko.md)에 있다.

## 60초 시작

`ni`는 현재 source-first다. Repository를 checkout한 뒤 실행한다:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
```

이제 conversation으로 `./my-plan/docs/plan/**`과
`./my-plan/.ni/contract.json`을 채운다. Readiness authority는 model이 아니라
CLI다:

```bash
go run ./cmd/ni status --dir ./my-plan --next-questions
go run ./cmd/ni end --dir ./my-plan
go run ./cmd/ni run --dir ./my-plan --target generic --max-chars 4000
```

`ni run`은 prompt를 compile한다. Shell commands, queues, agents, downstream
work를 실행하지 않는다.

## Choose your path

Go user가 아니어도 Intent Lock method를 시작할 수 있다. 다만 full deterministic
`ni` gate는 여전히 CLI에서 나온다.

| Path | Status | Best for | Boundary |
| --- | --- | --- | --- |
| Source | Available | `go run ./cmd/ni ...`를 실행할 수 있는 contributors와 early users. | Full deterministic `status`, `end`, `run`. |
| Release binary | Next | Go 설치 없이 `ni`를 쓰고 싶은 terminal users. | Published GitHub Release assets와 checksums를 기다린다. |
| Curl installer | Next | Release assets가 생긴 뒤 one-command install을 원하는 users. | `install.sh`는 있고 local test도 있지만, public install은 release assets를 기다린다. |
| Homebrew | Planned | Package manager를 선호하는 macOS users. | Published tap이나 formula가 없다. |
| Claude skill pack | Available | Packaged Claude skills로 model-assisted planning을 시작하는 users. | Drafting UX일 뿐이며 readiness와 lock authority는 CLI다. |
| Codex skill pack | Available | Packaged Codex skills로 model-assisted planning을 시작하는 users. | Drafting UX일 뿐이며 readiness와 lock authority는 CLI다. |
| No-terminal model workflow | Available as assisted method | CLI를 누군가 실행하기 전에 intent를 draft하고 review하려는 teams. | Deterministic full `ni`가 아니며 authoritative validation에는 CLI output이 필요하다. |

Source와 local build path는 [Install ni](docs/22_INSTALL.md)를 참고하라.
Release-gated installer는 [Curl Installer](docs/install-curl.ko.md)를,
assisted workflow는 [터미널 없이 계획하기](docs/no-terminal.ko.md)를 참고하라.
Planned distribution tracks는
[Distribution Strategy](docs/53_DISTRIBUTION_STRATEGY.ko.md)와
[Homebrew Distribution Plan](docs/54_HOMEBREW_DISTRIBUTION.ko.md)에 있다.
Distribution automation은 repository infrastructure이지 `ni` runtime execution이
아니다.

이 README는 package distribution이나 published binary release를 claim하지 않는다.
Deterministic CLI 사용에는 GitHub Release가 verified release assets를 실제로
포함하기 전까지 source, local build, local install mode를 사용한다. Skill packs와
assisted no-terminal planning은 그 validation 전에 intent draft를 도울 수 있다.

## Locked되는 것

Kernel은 pre-runtime control layer를 소유한다:

- `docs/plan/**` planning docs;
- `.ni/contract.json`;
- `ni status`의 deterministic readiness;
- `.ni/plan.lock.json`;
- `ni run`의 bounded prompt compilation.

Lock이 생긴 뒤에는 lockfile이 source of truth다. Current plan이 locked hashes와
더 이상 일치하지 않으면 handoff는 `BLOCKED`로 멈춘다.

## ni가 아닌 것

`ni`는 task runner, spec runner, multi-agent execution layer, queue, shell
adapter, PR automation system, release automation system, downstream work
runtime이 아니다. Seed material은 derived and mutable이며, locked plan이
authority다.

## 다음에 읽을 것

| Read | Why |
| --- | --- |
| [Why ni exists](docs/product-story.ko.md) | Compile-before-run 뒤의 product story. |
| [Why ni](docs/why-ni.md) | Product argument, boundary, benchmark framing. |
| [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) | Readiness, locking, hash trust, blocked handoff 규칙. |
| [터미널 없이 계획하기](docs/no-terminal.ko.md) | CLI 설치 전 method를 사용하되 validation claim을 하지 않는 방법. |
| [Command reference](docs/commands.ko.md) | Implemented CLI surface. |
| [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/) | Vague intent가 execution을 올바르게 멈추는 small demo. |

## License

`ni`는 [MIT License](LICENSE)로 licensed된다.

Security policy와 reporting scope는 [SECURITY.md](SECURITY.md)에 문서화되어 있다.
