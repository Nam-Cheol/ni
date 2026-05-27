# Benchmark Protocol

The benchmark protocol measures whether `ni` improves intent readiness before
any downstream agent, team, or harness starts work.

It is not an execution benchmark. It does not run Codex, Claude, Hyper Run,
shell commands, model APIs, or downstream adapters. It compares the intent
quality of a direct prompt against the intent quality available after the
Intent Lock Protocol has captured, validated, locked, and compiled the plan.

## Purpose

Vague product requests often hide missing acceptance criteria, unresolved
blockers, high-severity risks, implicit assumptions, and unclear non-goals.
Those gaps are expensive once execution has started because a downstream actor
must either stop and ask questions or silently invent intent.

This protocol tests the pre-runtime claim:

```text
ni should reduce ambiguous, stale, or unbounded intent before execution starts.
```

## Inputs

The input set is a small corpus of vague product requests. Each request should
be realistic enough that an agent or team could be tempted to begin work, but
underspecified enough that readiness should require clarification.

Seed requests live in:

```text
testdata/benchmark/vague-requests/
```

Requests may cover software, operations, research, content, internal tooling,
physical products, automation, or non-software product surfaces. The benchmark
should not assume that every project is a software implementation.

## Fixture Format

Each vague request fixture is a directory under
`testdata/benchmark/vague-requests/` with this shape:

```text
<fixture-slug>/
  request.md
  expected-hidden-assumptions.md
  expected-readiness-gaps.md
  suggested-ni-questions.md
```

`request.md` is the only input for the direct-to-agent path. It should contain
the vague request text and a category label.

`expected-hidden-assumptions.md` and `expected-readiness-gaps.md` are reviewer
seed notes for manual scoring. They are not benchmark results and must not be
reported as measured counts unless a reviewer actually applies the protocol.

`suggested-ni-questions.md` lists clarification questions that a ni-start
conversation could ask. It is not a required answer key and does not replace
authoritative `ni status` output.

The current fixture set intentionally covers:

- software dashboard,
- conversation product,
- research protocol,
- internal operations process,
- education program,
- document product,
- API or CLI tool,
- physical product planning,
- Namba AI upgrade style project,
- ambiguous automation request.

## Comparison Paths

For each vague request, compare two paths.

### A. Direct-to-agent prompt

Write the request as it would be handed directly to an agent or team. Do not
clarify it through `ni`, and do not run the downstream actor.

Review the prompt as static text and count visible readiness gaps. Hidden
assumptions are the assumptions a reviewer believes a downstream actor would
need to make in order to start.

### B. ni intent-lock path

Run the request through the pre-runtime `ni` flow:

```text
ni-start -> ni status -> ni-end -> ni-run prompt
```

The path may stop early. If `ni status` reports `BLOCKED`, record the blockers
and do not run `ni end` or `ni run`. If the plan reaches `READY` or
`READY_WITH_DEFERRALS`, use `ni end` to lock it, then use `ni run` to compile a
bounded target prompt.

This path measures the planning contract and compiled prompt only. It must not
execute the compiled prompt.

## Metrics

Use the same rubric for both paths.

| Metric | Count or Status | What to Measure |
| --- | --- | --- |
| Missing acceptance criteria count | Integer | Requirements or capabilities that lack concrete pass/fail, review, or protocol criteria. |
| Unmitigated high-risk count | Integer | High-severity risks without an explicit mitigation, owner, deferral, or acceptance rationale. |
| Unresolved blocker count | Integer | Open questions, conflicts, or missing decisions that should prevent trustworthy execution. |
| Hidden assumption count | Integer | Material assumptions a downstream actor would need to invent to begin work. |
| Non-goal coverage | `none`, `partial`, or `explicit` | Whether excluded scope is absent, implied, or documented clearly enough to constrain execution. |
| Stale plan detection | `not_applicable`, `passes`, or `blocked` | Whether changed planning state is detected before handoff when a lock exists. |
| Target prompt boundedness | Character count and `pass`/`fail` | Whether a compiled target prompt exists and stays within the configured maximum, initially 4000 characters. |
| Readiness status before execution | `BLOCKED`, `READY_WITH_DEFERRALS`, or `READY` | The authoritative `ni status` result before any downstream work begins. |

## Manual Procedure

1. Select one vague request fixture from
   `testdata/benchmark/vague-requests/`.
2. Create the direct-to-agent prompt from that fixture's `request.md` text
   only.
3. Score the direct prompt with the metric table.
4. Create a fresh `ni` workspace for the same request.
5. Use `ni-start` conversation to capture intent into `docs/plan/**` and
   `.ni/contract.json`.
6. Run `ni status` and record the authoritative readiness result.
7. If readiness is blocked, record blockers and stop the `ni` path for that
   request.
8. If readiness passes, run `ni end`, then run `ni run` with a 4000-character
   maximum and record the target prompt size.
9. Score the `ni` path with the same metric table.
10. Use the fixture's expectation files as review aids, not automatic scores.
11. Write the results in a benchmark report without claiming aggregate results
    that were not actually measured.

## Reporting Rules

- Do not claim empirical results unless actual benchmark runs exist.
- Do not fake numbers, averages, percentages, or examples.
- Mark unmeasured cells as `not_measured`.
- Keep request text, reviewer notes, status output, and prompt character counts
  auditable.
- Separate observations from conclusions.
- If reviewer judgment is used, name the reviewer role and the scoring date.
- If two reviewers disagree, keep both scores or record the reconciliation
  note.

## Boundary

The benchmark remains inside `ni-kernel`:

- It evaluates planning docs, contract readiness, lock behavior, and prompt
  boundedness.
- It may produce inert report files or fixture requests.
- It must not execute downstream agents.
- It must not call model APIs.
- It must not add telemetry, network dependencies, queues, shell adapters, or
  execution harnesses.

The benchmark is successful when it makes intent readiness visible before
execution. It is not responsible for proving that downstream implementation
quality improved.
