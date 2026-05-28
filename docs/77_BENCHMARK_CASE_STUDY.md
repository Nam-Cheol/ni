# Benchmark Case Study: Refund Triage Assistant

This is a manual, qualitative case study. It applies the benchmark protocol to
one vague request and one checked-in ni planning example. It is not a repeated
experiment, does not call external model APIs, does not execute downstream
work, and does not claim statistical significance.

## Case

- Fixture:
  `testdata/benchmark/vague-requests/refund-triage-assistant/`
- ni path artifact:
  `examples/ni-start-dogfood/`
- Reviewer role: repository maintainer applying the rubric in
  `docs/43_BENCHMARK_PROTOCOL.md`
- Scoring date: 2026-05-28
- Comparison:
  - A. direct-to-agent prompt readiness
  - B. `ni-start -> ni status -> ni-end -> ni-run` readiness

The case uses the same starting intent as the checked-in ni-start dogfood
transcript: a vague request for a refund triage assistant for support agents.
The direct path is reviewed as static prompt text only. The ni path is reviewed
from checked-in planning artifacts and verified CLI output.

## What Was Measured

The case study measures intent readiness before execution:

- missing acceptance criteria,
- unmitigated high risks,
- unresolved blockers,
- hidden assumptions,
- non-goal coverage,
- bounded target prompt availability.

It does not measure implementation quality, model quality, downstream agent
performance, user satisfaction, runtime safety, cost, latency, statistical
effect size, or production outcomes.

## Artifacts: A. Direct-to-Agent Prompt

Source:
`testdata/benchmark/vague-requests/refund-triage-assistant/request.md`

```text
I want a refund triage assistant for support agents.
```

Manual reviewer notes from the fixture:

```text
Expected hidden assumptions include:
- The assistant only drafts recommendations and does not approve, issue, or
  initiate refunds.
- Support agents, not customers, are the primary users.
- A specific internal refund policy source exists and is authoritative.
- Ambiguous policy cases should escalate to a supervisor instead of inventing a
  policy interpretation.

Expected readiness gaps include:
- Missing acceptance criteria for recommendation scope, policy citation,
  escalation behavior, privacy handling, transcript review, and runtime
  boundary.
- High-risk gaps around refund authority and stale or unclear policy
  interpretation.
- No explicit non-goals for refund approval, customer contact, supervisor
  replacement, or live integration work.
```

## Artifacts: B. ni Intent-Lock Path

The checked-in ni-start dogfood transcript records the planning path:

- vague user idea: `examples/ni-start-dogfood/02-user-vague-idea.md`
- confirmed answers: `examples/ni-start-dogfood/04-user-answers.md`
- docs and contract delta:
  `examples/ni-start-dogfood/05-docs-contract-delta.md`
- status proof: `examples/ni-start-dogfood/06-status-proof.md`
- end confirmation: `examples/ni-start-dogfood/08-ni-end-confirmation.md`
- run handoff: `examples/ni-start-dogfood/09-ni-run-handoff.md`
- planning workspace: `examples/ni-start-dogfood/workspace/`
- compiled prompt:
  `examples/ni-start-dogfood/workspace/generated/human-team.prompt.txt`

Initial `ni status --next-questions` result before the final blocker was
answered:

```text
BLOCKED
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document

blocker R009: OQ-001 is a blocker open question

question R009 OQ-001: Which refund policy source is authoritative for refund eligibility?
```

After the blocker was answered and the plan was updated, the verified command
output on 2026-05-28 was:

```text
$ go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace
READY_WITH_DEFERRALS
profile: prototype
product type: conversation_product
delivery surfaces: conversation, document
interaction mode: human_to_system
guidance: product_type=conversation_product changes planning guidance only; readiness authority remains the shared gate.
guidance: Cover conversation turns, failure handling, transcript evaluation, and human handoff.
guidance: conversation surface: specify turn boundaries, memory expectations, refusals, and escalation.
guidance: document surface: specify audience, structure, review workflow, and publication format.
deferral D001: DEC-004 is deferred
deferral D002: OQ-002 remains open
```

`ni end` was verified on a temporary copy of the workspace to avoid rewriting
checked-in lock data:

```text
$ go run ./cmd/ni end --dir /private/tmp/ni-benchmark-case-study.HlL2Rc/workspace
locked plan at /private/tmp/ni-benchmark-case-study.HlL2Rc/workspace/.ni/plan.lock.json
status READY_WITH_DEFERRALS
```

`ni run` was also verified on the temporary copy:

```text
$ go run ./cmd/ni run --dir /private/tmp/ni-benchmark-case-study.HlL2Rc/workspace --target human-team --max-chars 4000 --out /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt
compiled prompt at /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt

$ wc -m /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt
    3805 /private/tmp/ni-benchmark-case-study.HlL2Rc/human-team.prompt.txt
```

Actual compiled prompt artifact:

```text
Human-team handoff

Goal: Hand this locked NI plan to a PM/dev/design/research team for coordinated implementation planning.

Project: Refund Triage Assistant Plan - Plan a support-agent assistant that drafts refund recommendations from tickets and the internal policy page, escalates ambiguity, and excludes refund approval, customer contact, and runtimes.
Readiness: READY_WITH_DEFERRALS
Target: human-team (handoff)
Locked at: 2026-05-26T13:48:04Z

Authoritative sources:
- .ni/plan.lock.json is authoritative for lock state, hashes, and source-of-truth order.
- .ni/contract.json carries accepted CAP/REQ/EVAL/RISK IDs and acceptance criteria.
- docs/plan/ contains locked planning context; use only when hashes match.
- .ni/session.json is a planning aid below locked docs; it must not override contract or docs.
Source of truth: .ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history

Rules:
- Treat .ni/plan.lock.json as authoritative over .ni/contract.json, docs/plan, session state, and chat.
- Verify locked file hashes before using planning docs; on lock mismatch report BLOCKED and stop.
- If there are conflicting requirements, report BLOCKED with conflicting sources or IDs and stop.
- Do not weaken acceptance criteria, risk mitigations, or blocker handling to proceed.
- Every work item must trace to CAP/REQ/EVAL/RISK IDs.
- Keep ni at the pre-runtime boundary; this prompt is seed material, not kernel-owned execution state.
- Do not make ni run execute shell, Codex, adapters, queues, PR automation, or agent teams.

Target instructions:
- PM: maintain scope against accepted capabilities, non-goals, and open blocker questions.
- Dev: identify implementation packets and validation mapped to CAP/REQ/EVAL/RISK IDs.
- Design/research: confirm user-facing assumptions, evidence needs, and unresolved risks without changing acceptance criteria.
- Execution responsibility stays outside ni; this is a handoff artifact, not orchestration.

Process:
1. Review .ni/plan.lock.json first, then .ni/contract.json and docs/plan.
2. Assign next ownership only after confirming no lock mismatch or conflicting requirements.
3. Record validation evidence and blockers for the team before implementation proceeds.

Accepted capabilities:
- CAP-001: Draft refund recommendations for support agents.
- CAP-002: Escalate ambiguous or conflicting refund policy cases.
- CAP-003: Maintain synchronized human docs and machine contract records from conversation.

Accepted requirements:
- REQ-001: Draft recommendations only; never approve, issue, or initiate refunds.
- REQ-002: Cite ticket facts and the internal refund policy section used.
- REQ-003: Escalate policy ambiguity or ticket-policy conflict to a supervisor.
- REQ-004: Escalation language must not invent policy or customer commitments.
- REQ-005: Model-authored planning turns must update docs/plan/** and .ni/contract.json together.
- REQ-006: Readiness, lock, and prompt compilation must be delegated to ni status, ni end, and ni run rather than model judgment.

Risks:
- RISK-001 (high): Require recommendation language, agent review, and transcript checks for approval or customer-contact wording.
- RISK-002 (medium): Use only ticket facts needed for triage and keep generated handoffs inside the support-agent workflow.
- RISK-003 (high): Cite policy, escalate conflicts or unclear sections, and keep policy ownership visibly deferred.
- RISK-004 (medium): Name changed docs and affected IDs each authoring turn, then run ni status.

Open questions:
- OQ-002 blocker=false: Which support dashboard will eventually display the recommendation draft?

Expected output: team handoff with owners, next packets, validation evidence, risks, decisions needed, and BLOCKED if the lock or requirements conflict.
```

## Manual Readiness Comparison

These counts are one reviewer's manual rubric application for this single case,
not repeated numeric benchmark data.

| Metric | A. Direct prompt | B. ni intent-lock path |
| --- | --- | --- |
| Missing acceptance criteria | 6 missing categories: recommendation authority, policy citation, escalation behavior, privacy handling, transcript evaluation, runtime boundary | 0 blocker-grade gaps visible in the locked contract; six accepted requirements and three evaluations are present |
| Unmitigated high risks | 2 high-risk areas visible but unmitigated: implied refund authority and stale or unclear policy interpretation | 0 unmitigated high risks; `RISK-001` and `RISK-003` are high severity and include mitigations |
| Unresolved blockers | 5 blocker-grade unknowns: authority, policy source, escalation threshold, evaluation evidence, handoff target | 0 blockers; status is `READY_WITH_DEFERRALS` with two visible non-blocking deferrals |
| Hidden assumptions | 7 material assumptions listed in the fixture notes | 0 counted as hidden for measured scope; remaining uncertainty is visible as `DEC-004` and `OQ-002` deferrals |
| Non-goal coverage | none | explicit: no refund issuing or approval, no customer contact, no supervisor replacement, no runtime or live integration |
| Bounded target prompt availability | unavailable; direct prompt has no lock-verified compiled target prompt | pass; `ni run --max-chars 4000` produced a 3,805-character `human-team` prompt |

## Observations

The direct prompt is plausible but not ready to hand to an implementation
actor. A downstream actor would have to invent refund authority boundaries,
policy sources, escalation behavior, evaluation evidence, and non-goals before
starting.

The ni path did not make the work complete in a production sense. It made the
intent auditable before execution: the first readiness check blocked on the
authoritative policy source, the accepted plan carried explicit requirements,
high-risk mitigations, evaluations, non-goals, and visible deferrals, and the
target handoff prompt stayed under the configured 4,000-character bound.

## Limits

This case study is intentionally narrow. It does not prove that every ni plan
improves, that downstream implementation succeeds, or that one process is
statistically better than another. It shows one transparent before/after case
where the Intent Lock Protocol converts a vague prompt into a bounded,
lock-verified handoff without executing downstream work.
