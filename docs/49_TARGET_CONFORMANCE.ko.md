# 대상 출력 적합성

대상 출력은 downstream seed material이다. Codex, Spec Kit, Hyper Run,
Ouroboros, namba-ai, human team이 잠긴 NI plan에서 시작하도록 도울 수는
있지만, downstream tool이 소유하는 runtime execution state를 만들면 안 된다.

NI는 pre-runtime intent lock layer로 남는다. kernel은 readiness를 검증하고,
`.ni/plan.lock.json`을 쓰고, hash를 확인하고, 제한된 handoff artifact를
컴파일한다. 실행, queue, slash-command workflow, lifecycle state,
tool-owned packet은 NI 밖에 있어야 한다.

## 공통 규칙

모든 target output은 다음을 지켜야 한다:

- `.ni/plan.lock.json`, `.ni/contract.json`, locked `docs/plan/**` content에서
  파생될 것,
- output 생성 전에 `.ni/plan.lock.json`을 요구할 것,
- output을 쓰기 전에 lock hash를 검증할 것,
- stale lock이면 `BLOCKED`로 거부할 것,
- accepted requirements, risk mitigations, non-goals, blocker handling,
  source-of-truth order를 보존할 것,
- prompt, handoff, notes, maps, seed documents 같은 inert artifact로 남을 것.

모든 target output은 다음을 하면 안 된다:

- downstream CLI나 external binary 호출,
- downstream runtime packet directory 생성,
- queue, task-runner state, slash-command state, lifecycle state, PR
  automation, adapter execution state 생성,
- target output을 usable하게 보이도록 locked contract를 약화.

## Prompt targets

`codex` output은 제한된 prompt만이다. `ni run --target codex`는 artifact가
`prompt`인 prompt를 출력하거나 쓸 수 있지만, Codex를 호출하거나,
`.codex/commands`, queue state, execution script를 만들면 안 된다.

`human-team` output은 제한된 handoff prompt만이다. `ni run --target
human-team`은 PM/dev/design/research coordination을 위한 handoff artifact를
출력하거나 쓸 수 있지만, owner database, team workflow state, queue,
NI-owned orchestration을 만들면 안 된다.

두 prompt target은 prompt budget 안에 있어야 하며 기본값은 4000자다.

## Seed export targets

### hyper-run

Hyper Run export는 다음 seed files를 만들 수 있다:

```text
plan.md
ni-context.md
readiness-expectations.md
evidence-requirements.md
first-run-focus.md
```

다음 Hyper Run runtime packet paths를 만들면 안 된다:

```text
.hyper/
.hyper/goals/
.hyper/goals/GOAL-0001/
tasks.md
evidence.md
review.md
next.md
```

### namba-ai

namba-ai export는 다음 seed files를 만들 수 있다:

```text
planning.md
ni-lock-summary.md
capability-map.md
evaluation-map.md
risk-map.md
suggested-spec-boundaries.md
```

`suggested-spec-boundaries.md`는 graph-oriented seed material로 남아야 한다.
candidate boundary와 `depends_on` edge를 말할 수는 있지만, mandatory
sequential SPEC execution chain을 만들면 안 된다.

금지되는 namba-ai runtime 또는 chain state는 다음과 같다:

```text
.namba/
.namba/specs/
SPEC-001.md
SPEC-002.md
SPEC_SEQUENCE.md
specs/
tasks.md
run.md
sync.md
pr.md
land.md
```

### ouroboros

Ouroboros export는 다음 seed file을 만들 수 있다:

```text
ouroboros-seed-notes.md
```

다음 execute, evaluate, evolve runtime state를 만들면 안 된다:

```text
.ouroboros/
.ouroboros/runtime/
execute
execute.md
evaluate
evaluate.md
evolve
evolve.md
runtime/
```

### spec-kit

Spec Kit export는 다음 seed file을 만들 수 있다:

```text
spec-kit-seed-notes.md
```

다음 slash-command workflow state를 만들면 안 된다:

```text
.specify/
.specify/specs/
.specify/memory/
.github/prompts/
.claude/commands/
.codex/commands/
slash-commands.md
commands/
specify.md
plan.md
tasks.md
```

## Enforcement

Conformance는 세 층에서 확인한다:

- `go test ./...`는 stale lock refusal, seed-only export paths, prompt-only
  또는 handoff-only target output을 검증한다.
- `scripts/smoke.sh`는 각 seed export target에 대해
  `scripts/check-target-conformance.py`를 실행한다.
- `scripts/demo-check.sh`는 public demo가 downstream runtime을 호출하지 않고
  검증되는지 확인한다.

checker는 의도적으로 path-oriented이다. seed content는 발전할 수 있지만,
runtime-owned packet name과 directory는 NI output 밖에 있어야 한다.
