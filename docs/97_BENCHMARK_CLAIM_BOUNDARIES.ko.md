# Benchmark Claim Boundaries

## 목적

ni의 benchmark evidence는 pre-runtime intent-readiness evidence다. Product
quality, downstream agent quality, adoption, statistical effect를 증명하지
않는다. Benchmark case의 `READY`는 선언된 planning-artifact scope가 readiness
gate를 통과했다는 뜻이다. Product, research work, downstream implementation이
ready라는 뜻이 아니다.

## ni benchmark evidence가 뒷받침할 수 있는 것

| Claim | Supported? | Evidence type | Example |
| --- | --- | --- | --- |
| ni exposes unclear intent | Yes | BLOCKED status proof | OQ blockers |
| ni prevents premature handoff | Yes | no lock / no prompt before readiness | internal-dashboard, research-protocol |
| ni can lock artifact-readiness after answers | Yes | isolated workspace lock | benchmark workspace only |
| ni can compile bounded handoff after lock | Yes | ni-run prompt count | 4000-character prompt |
| ni tracks not_measured boundaries | Yes | benchmark docs and demo checks | not_measured tables |

## ni benchmark evidence가 아직 뒷받침할 수 없는 것

| Claim | Supported? | Why not |
| --- | --- | --- |
| Product implementation is correct | No | no product was implemented |
| Downstream agent performs well | No | no downstream agent was run |
| Rework is reduced | No | no repeated trial was measured |
| Adoption improved | No | no external usage data |
| Cost or latency improved | No | no runtime measurement |
| Research approval exists | No | synthetic fixture only |
| Fieldwork is authorized | No | no real review/authorization |
| Dashboard product is ready | No | artifact-readiness only |

## Required labels

Benchmark docs는 claim, transition table, case summary 근처에 다음 label을
보이게 유지해야 한다.

- measured
- not_measured
- artifact-readiness only
- synthetic fixture
- isolated workspace only
- no downstream execution
- no implementation claim
- no statistical claim

## Status vocabulary

- `BLOCKED`: readiness gap이 명시되어 있으며 lock 또는 prompt가 생성되면 안 된다.
- `READY`: 선언된 scope에 대해서만 readiness gate를 통과했다.
- `READY_WITH_DEFERRALS`: 명시적 deferral이 있을 때 lock이 가능할 수 있다.
- `not_measured`: 해당 claim에 대한 evidence를 수집하지 않았다.
- `artifact-readiness`: planning artifact는 handoff할 수 있지만 product 자체가
  증명된 것은 아니다.

## Case-specific boundaries

### Internal dashboard

모든 internal-dashboard benchmark summary는 다음을 말해야 한다.

- `READY`는 benchmark planning-meeting artifact readiness만 의미한다.
- Dashboard product readiness를 증명하지 않는다.
- Dashboard implementation quality를 증명하지 않는다.
- Downstream agent performance를 증명하지 않는다.

### Research protocol

모든 research-protocol benchmark summary는 다음을 말해야 한다.

- `READY`는 synthetic benchmark protocol planning artifact readiness만 의미한다.
- Real research approval을 증명하지 않는다.
- Fieldwork를 authorize하지 않는다.
- Research quality 또는 intervention effectiveness를 증명하지 않는다.

## Benchmark claim 검토 방법

1. 무엇을 측정했는가?
2. 어떤 command 또는 file이 그것을 증명하는가?
3. 무엇이 `not_measured`로 남아 있는가?
4. Workspace는 isolated 상태인가?
5. Lock이 생성되었는가?
6. Prompt가 compile되었는가?
7. Prompt가 실행되었는가?
8. Synthetic answer가 label되어 있는가?
9. Product/fieldwork/runtime claim을 피했는가?

## ni-grill이 이 문서를 쓰는 방법

`ni-grill`은 benchmark overclaim을 challenge하고, `not_measured` boundary가
보이는지 확인하며, evidence 없이 implementation quality, downstream agent
quality, real research approval, fieldwork authorization, rework reduction,
cost, latency, adoption, statistical effect를 암시하는 benchmark claim을 flag해야
한다.

## Demo-check expectations

`demo-check`는 benchmark docs가 claim-boundary marker를 포함하는지 검증할 수
있다. Generated prompt를 실행하면 안 된다. Downstream agent를 실행하면 안 된다.
