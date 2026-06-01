# Benchmark Report Example

## 1. 목적

이 directory는 `docs/43_BENCHMARK_PROTOCOL.md`에 정의된 Project Intent Compiler
benchmark method를 수동으로 보고하기 위한 deterministic template이자 작은
pre-runtime case library이다.

template은 empirical result를 포함하지 않는다. 특정 request에 protocol을
실제로 실행한 뒤에만 채운다. 빈 cell의 값을 만들어내지 않는다. case
directory는 manual qualitative readiness drill을 포함할 수 있다. lock, run,
prompt-count evidence가 없으면 반드시 `not_measured`로 남긴다.

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
  수 있는지 실행 전에 보여준다. 이 case는 historical `BLOCKED` proof를 보존하고
  resolved artifact-readiness variant의 isolated lock과 bounded prompt도
  기록한다. `READY`는 benchmark planning-meeting artifact readiness에만
  적용된다.
- research-protocol case는 같은 blocked-to-ready evidence pattern을 software
  밖에서 보여준다. Vague request `BLOCKED` proof를 보존하고, clearly labeled
  synthetic fixture answer를 적용하며, `READY`, lock, 4000-character prompt
  proof를 기록한다. Real research approval, fieldwork authorization, research
  quality는 `not_measured`로 남긴다.

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
  qualitative readiness drill. isolated ni workspace, checked-in blocked
  status proof, blocker analysis, resolved `READY` proof, isolated lock
  evidence, bounded prompt evidence, before/after evidence, lessons를 포함한다.
- `cases/research-protocol/`: vague non-software neighborhood cooling study
  request에 대한 manual qualitative readiness drill. isolated ni workspace,
  checked-in `BLOCKED` status proof, next-question evidence, blocker analysis,
  resolution path, synthetic fixture answer packet, resolved `READY` proof,
  isolated lock evidence, bounded prompt evidence, before/after evidence,
  lessons, explicit `not_measured` research/runtime boundary를 포함한다.
- `../../docs/88_SECOND_BENCHMARK_CASE_SELECTION.ko.md`: 두 번째 v0.5 benchmark
  case를 위한 selection plan. Research-protocol case를 추천하지만 새 benchmark
  result를 보고하지 않는다.
- `../../docs/43_BENCHMARK_PROTOCOL.md`: scoring method를 정의하는 benchmark
  protocol.

## 5. 명령

Repository root에서:

```bash
test -f examples/benchmark-report/README.md
test -f examples/benchmark-report/README.ko.md
test -f examples/benchmark-report/sample-report.md
test -f examples/benchmark-report/cases/internal-dashboard/README.md
test -f examples/benchmark-report/cases/internal-dashboard/04-measurement-table.md
test -f examples/benchmark-report/cases/internal-dashboard/06-ni-status-proof.md
test -f examples/benchmark-report/cases/internal-dashboard/07-ni-next-questions.md
test -f examples/benchmark-report/cases/internal-dashboard/08-blocker-analysis.md
test -f examples/benchmark-report/cases/internal-dashboard/09-resolution-path.md
test -f examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.md
test -f examples/benchmark-report/cases/internal-dashboard/16-lessons.md
test -f examples/benchmark-report/cases/research-protocol/README.md
test -f examples/benchmark-report/cases/research-protocol/04-measurement-table.md
test -f examples/benchmark-report/cases/research-protocol/06-ni-status-proof.md
test -f examples/benchmark-report/cases/research-protocol/07-ni-next-questions.md
test -f examples/benchmark-report/cases/research-protocol/08-blocker-analysis.md
test -f examples/benchmark-report/cases/research-protocol/09-resolution-path.md
test -f examples/benchmark-report/cases/research-protocol/10-answer-packet.md
test -f examples/benchmark-report/cases/research-protocol/11-resolved-status-proof.md
test -f examples/benchmark-report/cases/research-protocol/13-lock-summary.md
test -f examples/benchmark-report/cases/research-protocol/14-bounded-prompt-summary.md
test -f examples/benchmark-report/cases/research-protocol/15-before-after-evidence.md
test -f examples/benchmark-report/cases/research-protocol/16-lessons.md
test -f docs/43_BENCHMARK_PROTOCOL.md
go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions
rg -n "not_measured|must not execute downstream agents|Target prompt boundedness|internal-dashboard|research-protocol|NI Intent Readiness: BLOCKED|NI Intent Readiness: READY" examples/benchmark-report/README.md examples/benchmark-report/README.ko.md examples/benchmark-report/sample-report.md examples/benchmark-report/cases/internal-dashboard/*.md examples/benchmark-report/cases/research-protocol/*.md docs/43_BENCHMARK_PROTOCOL.md
```

## 6. 예상 출력

`test` 명령은 성공해야 한다.

`ni status` 명령은 resolved internal-dashboard artifact workspace에 대해
`NI Intent Readiness: READY`를 보고해야 한다. Historical blocked proof는
`cases/internal-dashboard/06-ni-status-proof.md`에 남아 있다.

Research-protocol `ni status` 명령은 resolved synthetic fixture workspace에 대해
`NI Intent Readiness: READY`를 보고해야 한다. Historical blocked proof는
`cases/research-protocol/06-ni-status-proof.md`에 남아 있다. Lock summary와
bounded prompt summary는 `13-lock-summary.md`, `14-bounded-prompt-summary.md`에
남아 있다.

`rg` 명령은 template, dashboard case, research case의 `not_measured` marker,
checked-in blocked/resolved status proof, blocker/next-question evidence,
before/after evidence, benchmark protocol의 non-execution 및 prompt-boundedness
marker를 보여야 한다.

## 7. demo-check coverage

`bash scripts/demo-check.sh`가 이 예시를 검증한다.

demo check는 required file을 확인하고 isolated internal-dashboard 및
research-protocol workspace에 대해 `ni status`를 실행한다. Historical blocked
proof, resolved READY proof, isolated lock evidence, bounded prompt evidence,
before/after evidence, lessons, 남은 `not_measured` claim boundary가 존재하는지
확인한다. `ni end`, generated prompt, dashboard code, research fieldwork, model
API, downstream agent는 실행하지 않는다.

## 8. Korean companion

Korean companion docs: `README.ko.md`.

## 9. 실행하지 않는 경계

이 report는 intent-focused 상태로 남아야 한다. downstream execution trace,
implementation result, telemetry, model API output, 만들어낸 aggregate claim을
포함하면 안 된다. statistical significance를 주장하면 안 된다.

internal-dashboard case는 product demo가 아니다. dashboard scaffold, runtime
harness, queue, shell adapter, model call, downstream-agent run으로 바꾸면 안
된다.
