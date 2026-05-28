# Landing Page Plan

Task 137은 `ni`에 canonical README를 넘어서는 GitHub Pages landing page가
필요한지 결정한다.

이 문서는 planning note일 뿐이다. Site를 구현하지 않고, frontend stack,
analytics, hosted service, install path 변경을 추가하지 않으며, website를 `ni`
사용의 필수 경로로 만들지 않는다.

## Recommendation

Recommendation: GitHub Pages static HTML.

`README.md`는 canonical quick entry로 유지한다. 다음 public-facing task에서
GitHub repository view보다 더 깔끔한 공유 surface가 필요할 때만 minimal GitHub
Pages page를 추가한다. 그 page는 README와 docs로 돌아가는 얇은 static pamphlet이어야
하며, 두 번째 product manual이 되면 안 된다.

지금 docs site는 만들지 않는다. 현재 product는 작은 source-first landing surface와
repository 안의 명시적인 protocol docs가 가장 잘 맞는다.

## Decision Summary

| Option | Fit | Tradeoff | Decision |
| --- | --- | --- | --- |
| README only | README가 이미 canonical quick entry이자 product pamphlet이므로 strong default다. | GitHub repository frame은 공유, screenshot, first-impression copy에는 덜 집중된 surface다. | Source of truth로 유지한다. |
| GitHub Pages static HTML | Launch, demo, social post link를 위한 lightweight public doorway로 유용하다. | README와 drift하지 않도록 관리해야 하는 두 번째 surface가 생긴다. | Follow-up으로 추천한다. |
| Docs site later | Docs가 navigation problem이 될 때만 유용하다. | 현재 kernel stage에는 너무 무겁고 premature site maintenance risk가 있다. | Deferred. |

## Why Not README Only

README-only는 usage에는 충분하다. 사용자가 `ni`가 무엇인지 이해하고, 설치하고,
더 깊은 docs를 찾는 데에는 계속 README가 충분해야 한다.

그래도 작은 landing page가 유용할 수 있는 이유는 public link의 첫 화면을 더
깔끔하게 만들 수 있기 때문이다:

- hero가 repository chrome 없이 core promise에 집중할 수 있다;
- 첫 demo를 더 쉽게 scan할 수 있다;
- install paths를 README를 더 늘리지 않고 묶을 수 있다;
- screenshots, release posts, social cards가 하나의 간결한 URL을 가리킬 수 있다.

이 이점은 presentation에만 해당한다. 새로운 product behavior나 authority를
추가하면 안 된다.

## Minimal Page Shape

Single static HTML page를 사용하고, 필요하면 작은 CSS file 하나만 둔다. Bundler,
framework, analytics, server component, cookie banner, tracking script,
external runtime dependency는 필요 없다.

Recommended sections:

1. Hero
2. Why ni
3. Start in 60 seconds
4. Install paths
5. Demo
6. Docs links

### Hero

이미 public copy로 정리된 문장을 사용한다:

```text
Don't run the agent yet. Compile the intent first.
```

Category는 다음으로 받친다:

```text
Project Intent Compiler for AI Agents.
```

Hero는 `ni`가 implementation work가 시작되기 전에 planning conversations를
locked project contracts로 바꾼다고 말해야 한다. Hero에서 specific downstream
harness 이름을 나열하지 않는다.

### Why ni

이 section은 짧고 pain-first로 유지한다:

- prompts는 actionable해 보여도 users, acceptance criteria, risks, non-goals,
  blocker questions가 빠져 있을 수 있다;
- early execution은 hidden assumptions를 비싸게 만든다;
- `ni`는 intent를 explicit하게 만들고, readiness를 check하고, accepted state를
  lock하고, intent가 바뀌면 handoff를 멈춘다.

더 깊은 mechanism은 Intent Lock Protocol로 link한다.

### Start In 60 Seconds

구현된 commands만 보여준다:

```bash
go run ./cmd/ni --help
go run ./cmd/ni init --dir ./my-plan --profile prototype
go run ./cmd/ni status --dir ./my-plan
go run ./cmd/ni end --dir ./my-plan
go run ./cmd/ni run --dir ./my-plan --target generic --max-chars 4000
```

Authoring은 `docs/plan/**`와 `.ni/contract.json`을 다루는 conversation을 통해
진행되고, CLI가 readiness, locking, hash validation, prompt compilation의
authority로 남는다는 점을 명확히 한다.

### Install Paths

이미 available 또는 planned로 문서화된 path만 나열한다:

- Source
- Local binary
- Release binary
- Curl installer
- Model workspaces as experimental planning assistance
- No-terminal method as experimental assisted drafting
- Homebrew as planned, not available

각 path는 `README.md`, `docs/22_INSTALL.md`, 또는 관련 source doc으로 link한다.
이미 문서화되지 않은 package-manager distribution, hosted binaries, services를
claim하지 않는다.

### Demo

Blocked ambiguous-prompt demo를 사용한다:

```bash
go run ./cmd/ni status --dir examples/ambiguous-prompt-blocked/workspace
```

```text
BLOCKED
```

Intent가 ready하지 않을 때 `BLOCKED`가 expected success state라는 점을 설명한다.
Downstream agent execution, shell adapters, queues, runtime automation을 추가하지
않는다.

### Docs Links

Link map은 작게 유지한다:

- README
- Install ni
- Intent Lock Protocol
- Command reference
- Why ni exists
- README Visual Wireframe
- Release readiness

README는 canonical quick entry로 남고, landing page는 public doorway일 뿐이다.

## Guardrails

- Heavy frontend stack을 추가하지 않는다.
- Analytics를 추가하지 않는다.
- Hosted service를 추가하지 않는다.
- Website를 `ni` 사용의 필수 경로로 만들지 않는다.
- Runtime execution behavior를 추가하지 않는다.
- Contract authoring CLI commands를 추가하지 않는다.
- GitHub Pages를 commands, install status, protocol rules의 source of truth로
  만들지 않는다.
- README나 docs의 긴 내용을 중복하지 않는다.

## Follow-Up Task

Project가 public doorway를 원할 때만 별도 follow-up task를 만든다:

```text
Task: Implement minimal GitHub Pages landing page.

Scope:
- static HTML page 하나와 작은 CSS file 하나, 또는 repository가 이미 지원하는
  가장 작은 GitHub Pages layout을 추가한다;
- README copy와 existing visual assets를 적절히 재사용한다;
- README를 canonical quick entry로 link한다;
- hero, why ni, start in 60 seconds, install paths, demo, docs links를 포함한다;
- framework, analytics, hosted service, runtime execution behavior를 추가하지
  않는다.

Validation:
- bash scripts/quality.sh
```

## Success Test

Landing page는 새 reader가 한 번 scan해서 `ni`를 이해하고, authoritative next
step으로 README나 install docs를 바로 선택할 수 있을 때만 만들 가치가 있다. Page가
두 번째 README, protocol reference, docs portal이 되기 시작하면 멈추고, later
docs-site task가 정당화될 때까지 README-only를 유지한다.
