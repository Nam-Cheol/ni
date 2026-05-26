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
```

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
