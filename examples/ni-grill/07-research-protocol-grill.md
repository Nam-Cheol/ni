# 07. Research Protocol Benchmark Grill

This transcript dogfoods `ni-grill` against the checked-in
`research-protocol` benchmark case.

Target workspace:

```text
examples/benchmark-report/cases/research-protocol/workspace/
```

Command run against the isolated benchmark workspace:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/research-protocol/workspace --proof --next-questions
```

Observed status:

```text
NI Intent Readiness: READY

Blockers:
- None.

Deferrals:
- None.

Warnings:
- None.
```

No deterministic blockers or deferrals were returned, so these findings are
pre-handoff hardening questions. They do not replace `ni status`, do not
approve lock, do not execute generated prompts, do not perform fieldwork, and
do not create real research approval or research-quality claims.

## Grill findings

1. GRILL-001 — Critical — claim boundary
   Affected: `RISK-004` /
   `examples/benchmark-report/cases/research-protocol/15-before-after-evidence.md`
   Concern: The case strongly says `READY` is synthetic benchmark protocol
   planning artifact readiness only, but the phrase "Safe to hand off as
   benchmark planning artifact" could be reused without the adjacent
   non-approval context.
   Why it matters: research-protocol handoff language can be mistaken for
   fieldwork authorization or ethics approval if quoted out of context.
   Question: Should the readiness transition table include inline wording that
   real research approval, fieldwork authorization, research quality, and
   intervention effectiveness remain `not_measured`?
   Answer shape: yes/no plus exact table wording or a reason the existing
   critical scope note is sufficient.
   Suggested action: clarify
   Blocks ni-end: maybe

2. GRILL-002 — High — acceptance evidence
   Affected: `EVAL-002` / `OQ-005` /
   `examples/benchmark-report/cases/research-protocol/workspace/docs/plan/07_evaluation_contract.md`
   Concern: Acceptance evidence requires a planning owner and privacy/safety
   reviewer, but the fixture does not distinguish synthetic review ownership
   from real institutional approval.
   Why it matters: the benchmark can be `READY`, yet a human reader may need a
   stronger cue that reviewer names are fixture roles rather than real
   governance approval.
   Question: Should the protocol add a fixture-reviewer note saying reviewer
   roles are synthetic benchmark roles and not IRB, ethics board, legal, or
   city approval?
   Answer shape: explicit non-approval note, reviewer role clarification, or
   rationale that existing `not_measured` sections are sufficient.
   Suggested action: clarify
   Blocks ni-end: maybe

3. GRILL-003 — High — risk and non-goal clarity
   Affected: `RISK-003` / `NG-001` /
   `examples/benchmark-report/cases/research-protocol/workspace/docs/plan/05_constraints.md`
   Concern: Field safety rules are detailed enough for benchmark planning, but
   they read like operational rules that someone could follow in the field.
   Why it matters: specific safety rules reduce ambiguity, but in a benchmark
   artifact they must remain planning constraints, not authorization to deploy
   a field team.
   Question: Should the safety section add a sentence that these rules are
   checklist criteria for a future protocol review and must not be used to
   start fieldwork from this benchmark?
   Answer shape: one sentence added to constraints or delivery operation, or a
   reason the existing non-goals already cover it.
   Suggested action: clarify
   Blocks ni-end: maybe

4. GRILL-004 — Medium — synthetic fixture label
   Affected: `DEC-004` / `OQ-001` through `OQ-005` /
   `examples/benchmark-report/cases/research-protocol/10-answer-packet.md`
   Concern: The answer packet clearly labels synthetic fixture answers at the
   top, but individual OQ sections contain realistic operational details that
   could be copied without the header.
   Why it matters: copied sections could lose the synthetic-fixture boundary
   and look like real approved study instructions.
   Question: Should each `OQ-*` section repeat "Synthetic benchmark fixture
   answer" in the required answer block or section lead?
   Answer shape: repeat label per section, add a copy-safety note, or explain
   why the top-level label is enough.
   Suggested action: clarify
   Blocks ni-end: no

5. GRILL-005 — Medium — prompt boundary
   Affected: `CAP-004` / `ART-008` /
   `examples/benchmark-report/cases/research-protocol/14-bounded-prompt-summary.md`
   Concern: The bounded prompt summary proves character count and
   non-execution, but not that truncation preserved the research non-approval
   warnings.
   Why it matters: a 4000-character seed prompt can be valid as kernel output
   while still needing human review for overclaiming before external use.
   Question: Should the case record a prompt-boundary review checklist for no
   fieldwork authorization, no participant recruitment, no research approval,
   no model API call, and preserved source-of-truth warning?
   Answer shape: reviewer role plus checklist result, or explicit decision
   that prompt text is not reviewed beyond boundedness in this benchmark.
   Suggested action: clarify
   Blocks ni-end: no

## Grill result

`ni-grill` found no CLI-readiness blockers. The strongest hardening question is
`GRILL-003`, because detailed safety rules are useful but must not be confused
with fieldwork authorization. The rest of the findings reinforce synthetic
fixture labeling, acceptance evidence, and `not_measured` claim boundaries
after `READY`.
