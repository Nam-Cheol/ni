# Security Policy

`ni`는 초기 source-first project다. 현재 제품은 `ni-kernel`이며, planning
contract를 validate, lock, compile하는 Project Intent Compiler다. `ni-kernel`은
downstream agents, shell commands, queues, runtime adapters를 실행하지 않는다.

## Supported Versions

아직 stable supported release channel은 없다. Source tags가 존재할 수 있지만,
release documentation이 달리 말하지 않는 한 stable security support channel은
아니다.

## Reporting Security Issues

아직 private vulnerability reporting channel은 publish되지 않았다.

Secrets, credentials, private prompts, proprietary planning contracts, sensitive
vulnerability details를 public issues에 올리지 마라. GitHub issues는
non-sensitive security reports, documentation corrections, 이 policy에 관한
questions에만 사용한다.

Future policy update에서 private security contact가 추가될 수 있다.

## Scope

이 policy는 다음을 포함한 `ni-kernel` project files와 behavior에 적용된다:

- planning docs and contract schemas,
- `.ni/contract.json` validation,
- `.ni/plan.lock.json` lock validation,
- source-of-truth and stale-lock checks,
- prompt and export compilation.

## Out of Scope

이 policy는 다음 downstream tools 또는 execution environments를 포함하지 않는다:

- Codex,
- Hyper Run,
- namba-ai,
- Spec Kit,
- Ouroboros,
- shell commands,
- generated prompts after they are executed outside `ni`.

## Secret Handling

Secrets를 `docs/plan/**`, `.ni/contract.json`, generated prompts, examples,
issues, documentation에 넣지 마라. Planning contracts와 prompts는 project owner가
별도의 private workflow를 명시적으로 제공하지 않는 한 source-visible project
artifacts로 취급해야 한다.

## Runtime Boundary

`ni-kernel`은 downstream execution 전에 intent를 compile하고 validate한다.
Downstream runtimes, agents, shell commands, generated prompts를 실행하지 않는다.
