# 01. init

This transcript is illustrative, but the checked-in workspace under
`workspace/` is a real NI example state that can be validated by the CLI.

## User

Initialize a planning workspace for a support-assistant idea.

## Model

I will let the CLI create the project skeleton first. After that, planning
state will be maintained through conversation by updating `docs/plan/**` and
`.ni/contract.json` together.

```bash
ni init --dir ./refund-triage-plan
```

## CLI result

```text
initialized NI workspace at ./refund-triage-plan
created docs/plan/** templates
created .ni/contract.json
created .ni/session.json
```

## Boundary shown

The user has not typed contract `add`, `list`, or `set` commands. `ni init`
creates the workspace; sustained conversation becomes the authoring interface.
