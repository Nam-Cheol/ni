# Authoring protocol

This protocol defines how a model updates NI planning state during sustained
conversation with a user.

## Turn loop

For each planning turn, the model should:

1. Read the current source of truth before changing files.
2. Extract new or corrected planning facts from the conversation.
3. Classify each fact as accepted, rejected, deferred, draft, assumption, or
   open question.
4. Update the relevant `docs/plan/**` files for human review.
5. Update `.ni/contract.json` with matching IDs, statuses, and traceability.
6. Update `.ni/session.json` with bounded carryover context for the next model
   session.
7. Preserve existing accepted criteria unless the user explicitly changes them.
8. Run or recommend `ni status` when readiness may have changed.

The docs and contract should move together. A conversation turn that changes a
capability, requirement, risk, evaluation, non-goal, decision, artifact, or open
question should update both the human-readable doc and the corresponding
contract field.

Session state should summarize the current focus, recent decisions and risks,
pending questions, readiness status, readiness blockers, and docs updated. It
is a planning aid, not authority.

## Classification rules

Accepted records require clear user commitment or a prior accepted decision.
The model may propose accepted wording, but it should not treat its own proposal
as accepted until the user confirms it or the conversation already makes the
acceptance unambiguous.

Draft records are allowed when the user has described likely work but has not
confirmed final scope. Deferred records are for known work intentionally moved
out of the current readiness profile. Rejected and not-applicable records should
capture explicit exclusions so later planning does not rediscover them.

Open questions are required when:

- a statement is ambiguous enough to change scope,
- a capability has no acceptance criteria,
- a capability has no evaluation,
- a high-severity risk lacks mitigation,
- a non-goal conflicts with requested behavior,
- the user asks the model to infer a business, policy, safety, or delivery
  decision without enough evidence.

## Traceability rules

Every accepted capability should trace to at least one requirement, evaluation,
risk, and artifact when the readiness profile requires that detail. Each
traceable item should use the existing ID prefixes:

```text
CAP-001   capability
REQ-001   requirement
DEC-001   decision
RISK-001  risk
EVAL-001  evaluation
ART-001   artifact
OQ-001    open question
NG-001    non-goal
```

IDs should be stable. Do not renumber existing records merely to make a list
look tidy. Add new IDs at the end, and mark obsolete records rejected, deferred,
or not applicable when preserving the history matters.

## Locked plans

After `.ni/plan.lock.json` exists, the source-of-truth order is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

The model must verify lock state before silently changing locked planning docs.
If locked hashes are stale, report `BLOCKED` and use the amendment or relock
flow instead of treating the chat transcript as stronger authority.

## Gate handoff

The model can say that planning appears ready for a gate, but only the CLI can
decide the gate result:

- run `ni status` to check readiness,
- run `ni end` to lock a ready plan,
- run `ni run` to compile the prompt from the locked plan.

`ni run` remains a compiler. It must not execute Codex, shell commands,
adapters, queues, generated harnesses, or downstream tools.
