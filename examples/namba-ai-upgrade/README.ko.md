# namba-ai Upgrade Planning

## 1. 목적

이 예시는 다음 namba-ai upgrade를 위한 Project Intent Compiler 사용을 보여준다.
runtime, SPEC-runner, Codex execution harness, namba-ai implementation work 전에
계획한다.

software-oriented downstream seed material에서도 ni-kernel이 권한을 유지한다는
점을 보여준다. 현재 fixture는 `ni status`로 검증되며, planning-doc body가
빠져 있어 fresh ready 상태가 아니다.

## 2. 증명하는 것

- software-oriented plan도 pre-runtime, contract-first일 수 있다.
- `ni status`는 readiness 권한이며 stale planning docs를 잡아낸다.
- grouped `ni status --proof --next-questions` output은 모델이 broad planning을
  다시 시작하지 않고 sync repair를 식별하게 한다.
- historical `.ni/plan.lock.json`은 existing downstream prompt seed material의
  boundary로 사용할 수 있다.
- existing lock에 대해 `ni run --target codex`는 Codex를 호출하지 않고 seed
  material만 만든다.
- downstream target notes는 derived and mutable이며 kernel-owned execution
  state가 아니다.

## 3. 제품 유형 / 표면

- `product_type`: `software`
- `delivery_surface`: `cli`, `document`, `workflow`
- 예상 `ni status`: `BLOCKED`
- 예상 `ni run` 대상: existing lock에서만 `codex`

## 4. 파일

- `docs/plan/**`: namba-ai upgrade intent를 위한 planning docs.
- `.ni/contract.json`: accepted capabilities, requirements, risks,
  evaluations, non-goals, artifacts, decisions.
- `.ni/plan.lock.json`: authoritative planning files의 hash를 담은 historical
  CLI lock.
- `generated/codex.goal.prompt.txt`: existing lock에서 컴파일된 Codex target
  prompt.

일부 duplicate `* 2.*` 파일은 historical fixture material이다. authoritative
locked planning files가 아니다.

## 5. 명령

Repository root에서:

```bash
go run ./cmd/ni status --dir examples/namba-ai-upgrade
go run ./cmd/ni status --dir examples/namba-ai-upgrade --proof --next-questions
tmpdir="$(mktemp -d)"
go run ./cmd/ni run --dir examples/namba-ai-upgrade --target codex --max-chars 4000 --out "$tmpdir/codex.goal.prompt.txt"
wc -m "$tmpdir/codex.goal.prompt.txt"
rm -rf "$tmpdir"
```

## 6. 예상 출력

예상 상태: `BLOCKED`.

상태 명령은 다음으로 시작해야 한다.

```text
BLOCKED
profile: prototype
product type: software
delivery surfaces: cli, document, workflow
```

현재 fixture blocker도 보여야 한다.

```text
blocker R012: CAP-001
```

`ni run`은 existing lockfile에서 prompt를 컴파일할 수 있지만, `ni status`가
다시 통과하기 전에는 이 예시를 fresh ready 상태로 설명하면 안 된다.

## 7. demo-check coverage

`bash scripts/demo-check.sh`가 이 예시를 검증한다.

demo check는 현재 `BLOCKED` status와 `R012` sync blocker를 확인한 뒤, existing
lock에서만 Codex prompt를 컴파일한다.

## 8. Korean companion

Korean companion docs: `README.ko.md`.

## 9. 실행하지 않는 경계

이 예시는 Codex를 실행하거나 namba-ai를 수정하지 않는다. shell adapter,
queue, model API, SPEC workflow, downstream tool도 실행하지 않는다. `ni status`가
`BLOCKED`를 보고하는 동안 이 fixture에 `ni end`를 호출하지 않는다. ni는
planning state를 검증하고 existing lock에서 inert prompt seed material만
컴파일한다.
