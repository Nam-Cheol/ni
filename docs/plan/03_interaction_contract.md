# Interaction contract

## v0.2 primary interaction

```text
ni init
ni status
ni end
ni relock
ni run
```

`ni init` creates the initial structure. After that, planning state is authored through model-user conversation and persisted by `ni-start` into `docs/plan/**` and `.ni/contract.json`. `ni status` reports deterministic readiness gaps. `ni end` or `ni relock` writes the lock only through CLI authority. `ni run` compiles a bounded handoff prompt only.

## Non-primary and later interaction

```text
ni targets
ni export
ni feedback
ni pressure
ni harness
ni amend
ni diff
ni conflicts
```

## Codex skill interaction

```text
$ni-start
$ni-end
$ni-run
```

## User control

The user accepts a plan by invoking `ni end` for the first lock or `ni relock` after an applied amendment. A model may recommend those commands only after `ni status` reports no blockers.

The v0.2 authoring protocol does not add user-facing contract `add`, `list`, or `set` commands. Users should not have to manually edit `.ni/contract.json`; the model-maintained authoring flow keeps docs and contract synchronized while the CLI validates the result.

## Target interaction

`ni run --target <target>` compiles a prompt. `ni export --target <target> --out <dir>` writes seed material. Neither command executes Codex, shells, agents, queues, PR automation, or downstream runtimes.

## Feedback interaction

Downstream feedback is input to review, not accepted truth. It may create pressure and later support an amendment, but accepted criteria change only through planning docs, `.ni/contract.json`, readiness, and relock.
