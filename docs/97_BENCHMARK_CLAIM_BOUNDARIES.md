# Benchmark Claim Boundaries

## Purpose

Benchmark evidence in ni is pre-runtime intent-readiness evidence. It does not
prove product quality, downstream agent quality, adoption, or statistical
effect. A `READY` benchmark case means the declared planning-artifact scope
passed the readiness gate; it does not mean the product, research work, or
downstream implementation is ready.

## What ni benchmark evidence can support

| Claim | Supported? | Evidence type | Example |
| --- | --- | --- | --- |
| ni exposes unclear intent | Yes | BLOCKED status proof | OQ blockers |
| ni prevents premature handoff | Yes | no lock / no prompt before readiness | internal-dashboard, research-protocol |
| ni can lock artifact-readiness after answers | Yes | isolated workspace lock | benchmark workspace only |
| ni can compile bounded handoff after lock | Yes | ni-run prompt count | 4000-character prompt |
| ni tracks not_measured boundaries | Yes | benchmark docs and demo checks | not_measured tables |

## What ni benchmark evidence cannot support yet

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

Benchmark docs must keep these labels visible near claims, transition tables,
and case summaries:

- measured
- not_measured
- artifact-readiness only
- synthetic fixture
- isolated workspace only
- no downstream execution
- no implementation claim
- no statistical claim

## Status vocabulary

- `BLOCKED`: readiness gaps are explicit; no lock or prompt should be produced.
- `READY`: readiness gates passed for the declared scope only.
- `READY_WITH_DEFERRALS`: lock may be possible with explicit deferrals.
- `not_measured`: no evidence was collected for that claim.
- `artifact-readiness`: the planning artifact can be handed off; the product
  itself is not proven.

## Case-specific boundaries

### Internal dashboard

Every internal-dashboard benchmark summary must say:

- `READY` means benchmark planning-meeting artifact readiness only.
- It does not prove dashboard product readiness.
- It does not prove dashboard implementation quality.
- It does not prove downstream agent performance.

### Research protocol

Every research-protocol benchmark summary must say:

- `READY` means synthetic benchmark protocol planning artifact readiness only.
- It does not prove real research approval.
- It does not authorize fieldwork.
- It does not prove research quality or intervention effectiveness.

## How to review a benchmark claim

1. What was measured?
2. What command or file proves it?
3. What remains not_measured?
4. Is the workspace isolated?
5. Was a lock created?
6. Was a prompt compiled?
7. Was the prompt executed?
8. Are synthetic answers labeled?
9. Are product/fieldwork/runtime claims avoided?

## How ni-grill should use this

`ni-grill` should challenge benchmark overclaims, check whether
`not_measured` boundaries are visible, and flag benchmark claims that imply
implementation quality, downstream agent quality, real research approval,
fieldwork authorization, rework reduction, cost, latency, adoption, or
statistical effect without evidence.

## Demo-check expectations

`demo-check` may verify that benchmark docs contain claim-boundary markers.
It must not execute generated prompts. It must not run downstream agents.
