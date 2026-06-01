# Conversation proof capture

Conversation proof capture is the short audit trail that `ni-start` should
show after each meaningful authoring update. It explains how the user's latest
answer changed planning state, which files and contract records were touched,
and what the CLI readiness gate said before and after the edit.

It exists so a user can inspect planning progress without trusting model vibes.
The proof is about intent authoring. It is not execution evidence, does not run
downstream work, and does not make the model a readiness authority.

## What Planning Proof Is

Planning proof is a user-visible summary of one authoring turn:

- the user's answer, paraphrased briefly,
- the model's planning interpretation by record type,
- the planning artifacts that actually changed,
- the contract fields or IDs that actually changed,
- decisions, assumptions, non-goals, and open questions affected by the turn,
- the `ni status --proof --next-questions` result before and after the edit,
- the next highest-priority question group from the CLI.

It should be concise. It should not expose hidden chain-of-thought. It should
not claim file changes, contract IDs, readiness, or lock state that did not
actually happen.

## How It Differs From Execution Evidence

Planning proof records the movement from conversation to docs and contract. It
answers "what intent changed?" and "what did the readiness gate say?".

Execution evidence would answer whether implementation ran correctly. That is
outside `ni-kernel`. Do not add runtime execution, downstream agents, shell
adapters, queues, or an execution evidence loop to produce planning proof.

## Required ni-start Block

After a meaningful authoring update, `ni-start` should run or request
`ni status --dir . --proof --next-questions` again, then report a block in this
shape:

```text
Planning proof:
- User input captured:
  "<short paraphrase of user answer>"
- Interpreted planning records:
  - Purpose: ...
  - Actors/outcomes: ...
  - Delivery surface: ...
  - Capabilities: CAP-001 ...
  - Requirements: REQ-001 ...
  - Risks: RISK-001 ...
  - Evaluations: EVAL-001 ...
  - Decisions: DEC-001 accepted/deferred/rejected if applicable
  - Assumptions: ASM-001 or open question if applicable
  - Non-goals: NG-001 if applicable
  - Open questions: OQ-001 ...
- Updated planning artifacts:
  - docs/plan/00_project_brief.md: purpose clarified
  - docs/plan/01_actors_outcomes.md: primary actors added
  - docs/plan/03_interaction_contract.md: delivery surface recorded
  - .ni/contract.json: project.purpose, actors/outcomes, delivery_surfaces updated
  - .ni/session.json: active focus and pending questions updated
- Status result:
  - before: BLOCKED because R014/R015/R016
  - after: BLOCKED/READY_WITH_DEFERRALS/READY because ...
- Remaining blockers:
  - OQ-001 ...
- Next question group:
  - Sync repairs / Risk decisions / Evaluation evidence / Open blockers / none
```

If a record type did not change, omit that line or say `none`. If no files were
changed, write `No planning artifacts were updated.` If command execution is
unavailable, the proof block should say that the before or after status is
pending exact CLI output from the user or a trusted runner.

## ni-grill Proof Use

`ni-grill` uses the same planning proof shape after the user answers grill
questions and the model updates planning artifacts. The proof should name the
`GRILL-*` findings addressed, the records changed, the before/after
`ni status --proof --next-questions` result, and the next question group.

If `ni-grill` only critiques and asks questions without changing files, it
should say `No planning artifacts were updated.` It must not present critique
as readiness proof or lock approval.

## How Users Should Read It

Users should check the proof in this order:

1. Does the paraphrase match what they meant?
2. Were tentative statements kept as assumptions or open questions?
3. Were clear exclusions captured as non-goals?
4. Do the changed files and contract fields match the stated interpretation?
5. Did the after-status come from `ni status --proof --next-questions`?
6. Is the next question group the highest-priority CLI group?

If the proof says docs and contract disagree, the next turn should repair the
stale side or keep the disagreement as a blocker. Do not proceed to `ni-end`
while sync diagnostics block readiness.

## What Not To Trust Without CLI Validation

Do not trust a model-only proof block as readiness, lock, or handoff authority.
The model may summarize edits, but only:

- `ni status` determines `BLOCKED`, `READY_WITH_DEFERRALS`, or `READY`,
- `ni end` creates `.ni/plan.lock.json`,
- `ni run` verifies lock hashes and compiles the bounded prompt.

In no-terminal mode, planning proof is a draft audit trail. It is useful for
reviewing what the model attempted to change, but it becomes trusted only after
a CLI run validates the drafted docs and contract.
