# Interaction contract

## CLI interaction

```text
ni init
ni status
ni end
ni relock
ni run
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

## Target interaction

`ni run --target <target>` compiles a prompt. `ni export --target <target> --out <dir>` writes seed material. Neither command executes Codex, shells, agents, queues, PR automation, or downstream runtimes.

## Feedback interaction

Downstream feedback is input to review, not accepted truth. It may create pressure and later support an amendment, but accepted criteria change only through planning docs, `.ni/contract.json`, readiness, and relock.
