# Readiness gate

`ni status` is the first critical command.

It must answer this question:

```text
Can this plan be locked for execution?
```

## Status outputs

```text
BLOCKED
  Required information is missing or contradictory.

READY_WITH_DEFERRALS
  Execution can start, but explicit low-risk deferrals remain.

READY
  No blocking issue is known.
```

## v0 blocking rules

The first implementation should enforce these rules:

```text
R001 required planning docs exist
R002 .ni/contract.json is valid JSON
R003 at least one capability exists
R004 every accepted capability has at least one linked evaluation
R005 every evaluation has a method
R006 every high-severity risk has mitigation
R007 every accepted capability has at least one artifact or requirement
R008 decisions use accepted, deferred, rejected, or not_applicable
R009 blocker open questions prevent lock
R010 at least one non-goal exists
R011 readiness profile definitions are valid
R012 planning docs and contract are synchronized
R013 accepted decisions do not contradict each other
```

## Proof report

`ni status --proof` prints a deterministic proof of the readiness status from
the same rule failures that produce readiness issues. It does not call a model,
does not edit docs, and does not start execution.

Use proof output when a person or downstream tool needs rule-level evidence:

```bash
ni status --dir . --proof
ni status --dir . --proof --json
```

See `docs/44_STATUS_PROOF.md` for the output contract.

## Next questions

`ni status --next-questions` derives focused planning questions from the same
rule failures that produce readiness issues. The questions are deterministic:
they do not call an LLM, do not resolve gaps, and do not add execution behavior.

Use JSON when another tool needs the questions:

```bash
ni status --dir . --json --next-questions
```

`next_questions` is included in JSON only when requested. See
`docs/34_READINESS_INTERVIEW.md` for the interview contract.

## Planning profiles

`ni status` evaluates issues under the selected `readiness_profile` in `.ni/contract.json`.

Profiles are planning readiness profiles only. They are not implementation stages and do not imply runtime packets, task management, or agent execution.

The supported profiles are:

```text
concept
prototype
mvp
beta
production
```

Profile definitions live in `.ni/readiness.profiles.json`. Each profile maps issue IDs to `blocker` or `deferral`. `prototype` is the default profile.

## Model boundary

The model may explain why a rule failed. It may not ignore a failed rule.
