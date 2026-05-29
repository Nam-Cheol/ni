# Benchmark Case Studies: Intent Readiness

This is a manual, qualitative case study report. It applies the benchmark
protocol to two vague requests and checked-in ni planning examples. It is not a
repeated experiment, does not call external model APIs, does not execute
downstream work, and does not claim statistical significance.

## Case 1: Refund Triage Assistant

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

This case uses the same starting intent as the checked-in ni-start dogfood
transcript: a vague request for a refund triage assistant for support agents.
The direct path is reviewed as static prompt text only. The ni path is reviewed
from checked-in planning artifacts and verified CLI output.

## What Was Measured

These case studies measure intent readiness before execution:

- missing acceptance criteria,
- unmitigated high risks,
- unresolved blockers,
- hidden assumptions,
- non-goal coverage,
- bounded target prompt availability.

It does not measure implementation quality, model quality, downstream agent
performance, user satisfaction, runtime safety, cost, latency, statistical
effect size, or production outcomes.

## Case 1 Artifacts: A. Direct-to-Agent Prompt

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

## Case 1 Artifacts: B. ni Intent-Lock Path

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

## Case 1 Manual Readiness Comparison

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

## Case 2: Neighborhood Cooling Study Protocol

- Fixture:
  `testdata/benchmark/vague-requests/community-heat-field-study/`
- ni path artifact:
  `examples/research-protocol/`
- Reviewer role: repository maintainer applying the rubric in
  `docs/43_BENCHMARK_PROTOCOL.md`
- Scoring date: 2026-05-29
- Comparison:
  - A. direct-to-agent prompt readiness
  - B. `ni-start -> ni status -> ni-end -> ni-run` readiness

This second case is a non-software research-protocol request. It is included to
check that the benchmark evidence is not limited to software assistants or
dashboards. The direct path is reviewed as static prompt text only. The ni path
is reviewed from checked-in locked planning artifacts and verified CLI output.

## Case 2 Artifacts: A. Direct-to-Agent Prompt

Source:
`testdata/benchmark/vague-requests/community-heat-field-study/request.md`

```text
Category: research protocol

Plan a field study to learn how residents deal with extreme heat and what
support would help them. Make it practical for the team to run soon.
```

Manual reviewer notes from the fixture:

```text
Expected hidden assumptions include:
- Residents can be recruited quickly and safely.
- Extreme heat exposure can be discussed without creating participant risk.
- The team has authority to collect, store, and analyze resident responses.
- "Support would help" means a program, physical resource, communication plan,
  or policy recommendation.
- The study method can be selected without a formal ethics or consent review.
- "Soon" gives enough time for protocol design, translation, accessibility,
  and field-team training.

Expected readiness gaps include:
- Missing acceptance criteria for research questions, sampling, consent,
  safety, data handling, field readiness, and final artifacts.
- High-risk gaps around participant safety, vulnerable populations, privacy,
  and environmental exposure.
- No explicit non-goals for intervention delivery, medical advice, policy
  commitments, or quantitative claims.
```

## Case 2 Artifacts: B. ni Intent-Lock Path

The checked-in research-protocol example records the locked planning path:

- planning workspace: `examples/research-protocol/`
- contract summary: `examples/research-protocol/contract-summary.md`
- planning docs: `examples/research-protocol/docs/plan/`
- locked contract: `examples/research-protocol/.ni/contract.json`
- lockfile: `examples/research-protocol/.ni/plan.lock.json`
- compiled prompt:
  `examples/research-protocol/generated/human-team.prompt.md`

Suggested ni questions from the fixture:

```text
- What are the exact research questions and expected study artifacts?
- Who may participate, and which groups require extra safeguards?
- What consent, privacy, translation, and accessibility requirements apply?
- What weather, location, or field-team safety conditions must stop the study?
- Which recommendations or interventions are out of scope for this protocol?
- Who reviews and accepts the protocol before fieldwork begins?
```

Verified `ni status` output on 2026-05-29:

```text
$ go run ./cmd/ni status --dir examples/research-protocol
READY
profile: prototype
product type: research_protocol
delivery surfaces: document
interaction mode: human_to_human
guidance: product_type=research_protocol changes planning guidance only; readiness authority remains the shared gate.
guidance: Cover hypothesis, data handling, method, analysis, ethics, and reproducibility evidence.
guidance: document surface: specify audience, structure, review workflow, and publication format.
```

`ni end` was verified on a temporary copy of the workspace to avoid rewriting
checked-in lock data:

```text
$ go run ./cmd/ni end --dir /private/tmp/ni-research-case-study.dYqnEa/research-protocol
locked plan at /private/tmp/ni-research-case-study.dYqnEa/research-protocol/.ni/plan.lock.json
status READY
```

`ni run` was also verified without executing the compiled prompt:

```text
$ go run ./cmd/ni run --dir examples/research-protocol --target human-team --max-chars 4000 --out /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md
compiled prompt at /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md

$ wc -m /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md
    3621 /private/tmp/ni-research-case-study.dYqnEa/human-team.prompt.md
```

Locked contract summary:

```text
Product type: research_protocol
Delivery surface: document
Readiness: READY
Accepted capabilities:
- CAP-001: Define the study question, hypothesis, sampling frame, and inclusion criteria.
- CAP-002: Specify data handling, ethics, and participant-contact boundaries.
- CAP-003: Define reproducible analysis and evidence requirements for the protocol.
Accepted requirements:
- REQ-001: The protocol must state the research question and hypothesis in falsifiable terms.
- REQ-002: The sampling frame must define eligible street segments, observation windows, and exclusion criteria.
- REQ-003: The protocol must separate public environmental observations from any human-subject interaction.
- REQ-004: The protocol must require ethics review before any participant contact or personally identifying data collection.
- REQ-005: The analysis plan must name variables, comparison method, reproducibility evidence, and reviewer acceptance criteria.
High risks:
- RISK-001: Ambiguous sampling criteria could bias the study; mitigated by explicit inclusion and exclusion criteria plus independent reviewer reproduction.
- RISK-002: Participant privacy or ethics boundaries may be crossed; mitigated by separating environmental observation from participant contact and requiring ethics review before personal data collection.
Non-goals:
- No fieldwork, participant data collection, sensor deployment, analysis runtime, or ethics approval from ni.
```

Target prompt excerpt, from the 3,621-character bounded prompt:

```text
Project: Neighborhood Cooling Study Protocol - Plan a documented field research protocol for comparing street-level cooling interventions before any study execution begins.
Readiness: READY
Target: human-team (handoff)

Accepted capabilities:
- CAP-001: Define the study question, hypothesis, sampling frame, and inclusion criteria.
- CAP-002: Specify data handling, ethics, and participant-contact boundaries.
- CAP-003: Define reproducible analysis and evidence requirements for the protocol.

Accepted requirements:
- REQ-001: The protocol must state the research question and hypothesis in falsifiable terms.
- REQ-002: The sampling frame must define eligible street segments, observation windows, and exclusion criteria.
- REQ-003: The protocol must separate public environmental observations from any human-subject interaction.
- REQ-004: The protocol must require ethics review before any participant contact or personally identifying data collection.
- REQ-005: The analysis plan must name variables, comparison method, reproducibility evidence, and reviewer acceptance criteria.
```

## Case 2 Manual Readiness Comparison

These counts are one reviewer's manual rubric application for this single case,
not repeated numeric benchmark data.

| Metric | A. Direct prompt | B. ni intent-lock path |
| --- | --- | --- |
| Missing acceptance criteria | 7 missing categories: research questions, sampling, consent, safety, data handling, field readiness, final artifacts | 0 blocker-grade gaps visible in the locked contract; five accepted requirements and three evaluations are present |
| Unmitigated high risks | 4 high-risk areas visible but unmitigated: participant safety, vulnerable populations, privacy, environmental exposure | 0 unmitigated high risks; `RISK-001` and `RISK-002` are high severity and include mitigations |
| Unresolved blockers | 5 blocker-grade unknowns: participant criteria, locations, study method, consent process, review owner | 0 blockers; status is `READY` |
| Hidden assumptions | 6 material assumptions listed in the fixture notes | 0 counted as hidden for measured scope; assumptions are represented as requirements, risks, evaluations, or non-goals |
| Non-goal coverage | none | explicit: no fieldwork, participant data collection, sensor deployment, analysis runtime, or ethics approval from ni |
| Bounded target prompt availability | unavailable; direct prompt has no lock-verified compiled target prompt | pass; `ni run --max-chars 4000` produced a 3,621-character `human-team` prompt |

## Observations

Both direct prompts are plausible but not ready to hand to a downstream actor.
In the refund case, a downstream actor would have to invent refund authority
boundaries, policy sources, escalation behavior, evaluation evidence, and
non-goals before starting. In the research-protocol case, a downstream actor
would have to invent sampling, consent, field safety, ethics, evidence, and
review boundaries before starting.

The ni paths did not make the work complete in a production sense. They made
intent auditable before execution: the refund readiness check blocked on the
authoritative policy source before reaching `READY_WITH_DEFERRALS`, the
research-protocol case reached `READY`, both accepted plans carried explicit
requirements, high-risk mitigations, evaluations, and non-goals, and both
target handoff prompts stayed under the configured 4,000-character bound.

## Limits

This case study report is intentionally narrow. It does not prove that every
ni plan improves, that downstream implementation succeeds, or that one process
is statistically better than another. It shows two transparent before/after
cases where the Intent Lock Protocol converts a vague prompt into a bounded,
lock-verified handoff without executing downstream work.
