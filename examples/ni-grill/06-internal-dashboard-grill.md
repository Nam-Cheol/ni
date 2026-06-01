# 06. Internal Dashboard Benchmark Grill

This transcript dogfoods `ni-grill` against the checked-in
`internal-dashboard` benchmark case.

Target workspace:

```text
examples/benchmark-report/cases/internal-dashboard/workspace/
```

Command run against the isolated benchmark workspace:

```bash
go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions
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
approve lock, do not execute generated prompts, and do not create dashboard
product-readiness claims.

## Grill findings

1. GRILL-001 — Critical — claim boundary
   Affected: `RISK-004` /
   `examples/benchmark-report/cases/internal-dashboard/15-before-after-evidence.md`
   Concern: The evidence file says `READY` applies only to benchmark
   planning-meeting artifact readiness, but the transition row still says
   "Safe to hand off as benchmark planning artifact" without naming the
   forbidden downstream interpretations in the same row.
   Why it matters: a downstream reader could quote the transition table without
   the later `not_measured` section and make it sound like dashboard product
   work is ready.
   Question: Should the transition table include an inline note that dashboard
   product readiness, dashboard implementation quality, and downstream agent
   performance remain `not_measured`?
   Answer shape: yes/no plus the exact row wording or a reason the existing
   adjacent scope note is sufficient.
   Suggested action: clarify
   Blocks ni-end: maybe

2. GRILL-002 — High — acceptance evidence
   Affected: `EVAL-002` / `OQ-004` /
   `examples/benchmark-report/cases/internal-dashboard/workspace/docs/plan/07_evaluation_contract.md`
   Concern: The accepted evidence says every required answer field must be
   filled and pass/fail criteria must be testable, but the specific planning
   meeting date remains unassigned.
   Why it matters: `READY` can be valid for the benchmark artifact, yet a
   future planning handoff may need to know whether "next scheduled planning
   meeting" is a placeholder or an intentionally deferred operational detail.
   Question: Should the benchmark mark the meeting date as explicitly
   unassigned-but-non-blocking for artifact readiness?
   Answer shape: accepted non-blocking note, deferred operational follow-up, or
   explanation that the planning owner role is enough evidence.
   Suggested action: clarify
   Blocks ni-end: maybe

3. GRILL-003 — High — risk and non-goal clarity
   Affected: `RISK-002` / `NG-002` /
   `examples/benchmark-report/cases/internal-dashboard/workspace/docs/plan/05_constraints.md`
   Concern: Privacy boundaries prohibit private customer data and sensitive
   source data, while `OQ-003` also allows "approved internal dashboard source
   material required to validate the benchmark case."
   Why it matters: this is defensible, but a reader may not know who approves
   that source material or how much may be copied into evidence files.
   Question: Should the case name the approval role and maximum allowed source
   excerpt shape for internal dashboard source material?
   Answer shape: approval role plus allowed reference-only, summary-only, or
   short excerpt rule.
   Suggested action: clarify
   Blocks ni-end: maybe

4. GRILL-004 — Medium — handoff boundary
   Affected: `CAP-004` / `ART-007` /
   `examples/benchmark-report/cases/internal-dashboard/14-bounded-prompt-summary.md`
   Concern: The bounded prompt summary confirms a 4000-character prompt was
   compiled, but it does not say whether the prompt text itself was reviewed
   for product-readiness overclaiming after truncation.
   Why it matters: a truncated prompt can preserve the lock boundary while
   still cutting away useful warning text or leaving an ambiguous closing
   instruction.
   Question: Should the benchmark add a manual prompt-boundary review note for
   the generated 4000-character prompt?
   Answer shape: reviewer role plus checklist for no execution instruction, no
   dashboard build claim, no product-readiness claim, and preserved
   source-of-truth warning.
   Suggested action: clarify
   Blocks ni-end: no

5. GRILL-005 — Low — docs/contract sync claim
   Affected: `.ni/contract.json` / `DEC-004` /
   `examples/benchmark-report/cases/internal-dashboard/workspace/docs/plan/11_decision_log.md`
   Concern: `DEC-004` keeps approval explicit even when the named person is
   unassigned, but the contract and docs rely on roles rather than a concrete
   owner.
   Why it matters: role-based acceptance is acceptable for a benchmark, but a
   downstream planning meeting may need to avoid treating an unnamed owner as
   real approval.
   Question: Should `DEC-004` be cited in the evidence summary when explaining
   that approval is role-defined, not person-completed?
   Answer shape: cite `DEC-004` in the evidence summary, add a non-blocking
   follow-up, or leave as-is with rationale.
   Suggested action: keep as note
   Blocks ni-end: no

## Grill result

`ni-grill` found no CLI-readiness blockers. The strongest hardening question is
`GRILL-003`, because the phrase "approved internal dashboard source material"
could benefit from a clearer approval role and excerpt rule. The remaining
findings mainly protect claim boundaries after `READY`, especially around
dashboard product readiness and bounded prompt interpretation.
