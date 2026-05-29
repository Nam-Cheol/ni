# Benchmark Report Example

## 1. 목적

이 directory는 `docs/43_BENCHMARK_PROTOCOL.md`에 정의된 Project Intent Compiler
benchmark method를 수동으로 보고하기 위한 deterministic template이자 작은
docs-only case library이다.

template은 empirical result를 포함하지 않는다. 특정 request에 protocol을
실제로 실행한 뒤에만 채운다. 빈 cell의 값을 만들어내지 않는다. case
directory는 manual qualitative readiness drill을 포함할 수 있지만 lock,
status, run, prompt-count evidence가 없으면 반드시 `not_measured`로 남긴다.

## 2. 증명하는 것

- benchmark reporting은 pre-runtime이고 evidence-based 상태로 남는다.
- 실제 manual run이 있기 전까지 empty result cell은 `not_measured`로 남는다.
- benchmark는 direct-to-agent prompt readiness와 ni intent-lock path를
  비교하지만 어느 쪽도 실행하지 않는다.
- report format은 status output, prompt boundedness, reviewer notes를
  audit 가능하게 만든다.
- 이것은 case-study method이며 empirical result나 statistical significance를
  주장하지 않는다.
- internal-dashboard case는 그럴듯한 dashboard request가 users, success
  criteria, data boundaries, risks, non-goals, handoff evidence를 어떻게 숨길
  수 있는지 실행 전에 보여준다.

## 3. 제품 유형 / 표면

- `product_type`: 해당 없음. 이것은 report template이다.
- `delivery_surface`: `document`
- 예상 `ni status`: 해당 없음. 이 directory는 ni workspace가 아니다.
- 예상 `ni run` 대상: 해당 없음

## 4. 파일

- `README.md`: report template과 boundary statement.
- `README.ko.md`: Korean companion guide.
- `sample-report.md`: `not_measured` placeholder를 가진 fillable template.
- `cases/internal-dashboard/`: vague dashboard request에 대한 manual
  qualitative readiness drill. docs-only이며 locked ni workspace가 아니다.
- `../../docs/43_BENCHMARK_PROTOCOL.md`: scoring method를 정의하는 benchmark
  protocol.

## 5. 명령

Repository root에서:

```bash
test -f examples/benchmark-report/README.md
test -f examples/benchmark-report/README.ko.md
test -f examples/benchmark-report/sample-report.md
test -f examples/benchmark-report/cases/internal-dashboard/04-measurement-table.md
test -f docs/43_BENCHMARK_PROTOCOL.md
rg -n "not_measured|must not execute downstream agents|Target prompt boundedness|internal-dashboard" examples/benchmark-report/README.md examples/benchmark-report/README.ko.md examples/benchmark-report/sample-report.md examples/benchmark-report/cases/internal-dashboard/*.md docs/43_BENCHMARK_PROTOCOL.md
```

## 6. 예상 출력

`test` 명령은 성공해야 한다.

`rg` 명령은 template과 dashboard case의 `not_measured` marker, 그리고
benchmark protocol의 non-execution 및 prompt-boundedness marker를 보여야 한다.

## 7. demo-check coverage

`bash scripts/demo-check.sh`가 이 예시를 검증한다.

demo check는 이 예시를 docs-only로 다룬다. `ni status`나 `ni run`을 실행하지
않고 required file과 `not_measured` placeholder를 확인한다.

## 8. Korean companion

Korean companion docs: `README.ko.md`.

## 9. 실행하지 않는 경계

이 report는 intent-focused 상태로 남아야 한다. downstream execution trace,
implementation result, telemetry, model API output, 만들어낸 aggregate claim을
포함하면 안 된다. statistical significance를 주장하면 안 된다.

internal-dashboard case는 product demo가 아니다. dashboard scaffold, runtime
harness, queue, shell adapter, model call, downstream-agent run으로 바꾸면 안
된다.
