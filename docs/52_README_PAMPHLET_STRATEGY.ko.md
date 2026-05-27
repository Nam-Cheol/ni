# README Pamphlet Strategy

이 strategy는 `README.md`의 다음 역할을 정의한다. `README.md`는 technical
specification이 아니라 `ni`의 product landing page처럼 읽혀야 한다.

`README.md`는 canonical public README로 남는다. Korean companion docs가
유지되는 동안 `README.ko.md`는 companion으로 남는다. 상세 concept,
comparison, protocol, benchmark, roadmap, command material은 README 안에서
길게 설명하지 말고 docs로 이동한 뒤 link해야 한다.

## Job

README는 첫 화면에서 다음 다섯 질문에 순서대로 답해야 한다:

1. `ni`는 어떤 문제를 해결하는가?
2. locked intent contract가 사용자에게 어떤 payoff를 주는가?
3. 가장 짧게 어떻게 시작할 수 있는가?
4. 이 도구가 bounded and trustworthy하다는 trust signal은 무엇인가?
5. 더 깊은 설명은 어디서 읽는가?

첫 화면은 product idea를 팔아야 한다:

- agents는 project intent가 accepted and testable 상태가 되기 전에 시작되는
  경우가 많다;
- `ni`는 accepted planning state를 컴파일해 control을 execution 이전으로
  옮긴다;
- 결과물은 locked, versioned, verifiable project contract다;
- downstream actors는 readiness와 lock checks를 통과한 뒤에만 bounded prompts
  또는 inert seed material을 받는다.

Hero에서는 specific downstream harness products를 이름으로 나열하지 않는다.
필요할 때만 Codex나 Claude 같은 model/workspace examples를 언급할 수 있지만,
hero의 주장은 product-level이어야 한다. 즉 어떤 agent나 team이 work를
시작하기 전에 intent를 compile한다는 메시지가 먼저다.

## Proposed README Section Order

1. Hero
2. Badges
3. Why ni
4. Three-step use path
5. Short demo
6. Install and use options
7. Trust signals
8. What ni is not
9. Read next
10. Development and release status

## Section Intent

### Hero

Hero는 짧고 기억하기 쉬워야 한다:

- product name;
- language toggle;
- tagline: `Project Intent Compiler for AI Agents`;
- problem을 설명하는 한 문장;
- payoff를 설명하는 한 문장;
- conversation, readiness, lock, prompt compilation을 보여주는 작은 flow.

Hero에서 downstream product list를 만들지 않는다. Hero가 target documentation처럼
느껴지면 안 된다.

### Badges

Badges는 이미 사실인 trust signals에만 사용한다:

- CI status가 존재하면 CI status;
- license;
- source-first status를 평이하게 표현할 수 있다면 source-first status.

Package-manager distribution, hosted service availability, published binary
release를 해당 경로가 존재하기 전에 암시하면 안 된다.

### Why ni

Mechanism보다 user pain을 먼저 설명한다:

- prompts는 actionable하게 보이지만 users, acceptance criteria, risks,
  non-goals, blocker questions를 숨길 수 있다;
- hidden assumptions는 agent가 시작된 뒤 더 비싸진다;
- `ni`는 intent를 explicit하게 만들고, deterministic하게 check하고, lock하고,
  intent가 바뀌면 handoff를 멈춘다.

여기서는 protocol terms를 가볍게 유지한다. 전체 mechanism은 Intent Lock
Protocol 문서로 link한다.

### Three-Step Use Path

가장 짧은 source-first path를 보여준다:

```text
1. planning workspace를 init한다
2. readiness를 check하고 blockers를 resolve한다
3. intent를 lock하고 prompt를 compile한다
```

구현된 commands만 사용한다:

- `ni init`
- `ni status`
- `ni end`
- `ni run`

Contract authoring CLI commands를 추가하지 않는다. Authoring은 docs와
`.ni/contract.json`을 다루는 model-user conversation으로 유지하고, CLI는
authority로 남는다.

### Short Demo

상단에는 짧은 demo 하나만 둔다. 기본값으로는 ambiguous prompt blocked demo가
좋다. 이 demo가 product value를 가장 빨리 보여주기 때문이다:

- vague request에서 시작한다;
- `ni status`를 실행한다;
- `BLOCKED`를 보여준다;
- execution이 아직 시작되면 안 되므로 이것이 성공이라는 점을 설명한다.

Demo는 downstream agents를 실행하면 안 되며, smoke checks가 cover하는 local
validation과 CLI commands를 벗어난 shell work를 요구하면 안 된다.

### Install and Use Options

구현된 path만 claim한다:

- source-first `go run`;
- local build;
- repository가 지원한다면 local install.

Homebrew, package-manager distribution, hosted binaries, release automation을
claim하지 않는다.

### Trust Signals

Trust signals를 깊은 references보다 앞에 둔다:

- deterministic readiness gate;
- lockfile hash verification;
- prompt output budget;
- `ni run`은 prompts를 compile하며 실행하지 않는다는 점;
- CI, smoke checks, quality checks, demo checks, install checks;
- source-first release status;
- MIT license와 security policy.

이 signals는 `ni`가 execution runtime이 아니라 pre-runtime kernel이라는 점을
reader에게 안심시켜야 한다.

### What ni Is Not

이 부분은 짧은 boundary block 하나로 유지한다. `ni`가 execution harness, task
runner, multi-agent orchestration layer, queue, shell adapter, release system이
아니라는 점을 말한다.

README가 target별 boundary prose를 들고 있으면 안 된다. 더 깊은 target docs로
link한다.

### Read Next

Pamphlet의 끝에는 작은 map을 둔다:

- install details;
- Intent Lock Protocol;
- command reference;
- target story;
- benchmark protocol;
- release readiness;
- launch checklist;
- post-release roadmap.

## Move Out Of README

다음 material은 docs로 이동하고 README에서는 link한다:

- detailed protocol explanation;
- related-work comparison;
- target-by-target boundaries;
- full command reference;
- benchmark methodology;
- post-release roadmap.

README는 이 주제들에 대해 한 문장 summary를 유지할 수 있지만, body가
specification처럼 길어지면 안 된다.

## Keep In README

README는 다음을 유지해야 한다:

- hero;
- badges;
- why ni;
- three-step use path;
- short demo;
- install and use options;
- trust signals;
- 짧은 `What ni is not` block 하나;
- read-next links.

## Rewrite Guardrails

- `README.md`는 canonical, `README.ko.md`는 companion으로 유지한다.
- README hero에서 specific downstream harness products를 이름으로 언급하지
  않는다.
- 구현되지 않은 install paths를 claim하지 않는다.
- execution runtime behavior를 추가하지 않는다.
- contract authoring CLI commands를 추가하지 않는다.
- model execution adapters, shell adapters, queues, agent orchestration을
  추가하지 않는다.
- `ni run`은 prompt compilation only로 문서화한다.

## Success Test

Rewrite 후 새 reader는 다음 질문에 답할 수 있어야 한다:

- `ni`가 왜 존재하는가;
- agent를 시작하기 전에 왜 이것이 중요한가;
- source에서 어떻게 try할 수 있는가;
- kernel boundary를 왜 trust할 수 있는가;
- 다음에는 어떤 deeper doc을 읽어야 하는가.

README 첫 화면이 protocol reference, command manual, target matrix, benchmark
paper, roadmap처럼 느껴진다면 pamphlet rewrite는 실패한 것이다.
