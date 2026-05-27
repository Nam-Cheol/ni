# Ambiguous Prompt Blocked

## 1. Purpose

This example shows ni's core payoff: ambiguous execution is blocked before an
agent starts implementation.

A vague request such as "build me a dashboard for my team" sounds actionable,
but it hides product decisions that an implementation agent would otherwise
invent silently.

## 2. What this proves

- `ni status` is the readiness authority for captured but incomplete intent.
- Blocker open questions keep the workspace in `BLOCKED`.
- The model may draft planning records, but it may not declare readiness.
- A Codex target prompt is only illustrative here; it is not compiled or
  executed from the blocked workspace.

## 3. Product type / surface

- `product_type`: `software`
- `delivery_surface`: `web`
- Expected `ni status`: `BLOCKED`
- Expected `ni run` target: not applicable while blocked

## 4. Files

- `01-vague-request.md`: the ambiguous user prompt.
- `02-direct-to-agent-risk.md`: assumptions a direct implementation path would
  likely make.
- `03-ni-start-conversation.md`: how `ni-start` captures intent as planning
  records.
- `04-ni-status-blocked.md`: the deterministic blocked status.
- `05-next-questions.md`: focused questions needed before lock.
- `06-user-answers.md`: example answers that could resolve ambiguity.
- `07-locked-contract-summary.md`: the ready-state narrative after answers.
- `08-codex-target-prompt.md`: an illustrative bounded prompt after lock.
- `workspace/docs/plan/**`: the captured blocked planning docs.
- `workspace/.ni/contract.json`: the matching machine-readable contract.

## 5. Commands

From the repository root:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

## 6. Expected output

Expected status: `BLOCKED`.

The status command should include:

```text
BLOCKED
blocker R009: OQ-001 is a blocker open question
blocker R009: OQ-002 is a blocker open question
```

The next-questions command should also start with `BLOCKED` and list the
blocker questions.

## 7. Non-execution boundary

This example does not execute Codex, implement a dashboard, start a shell
adapter, create a queue, call a model API, or run downstream tools. It is a
kernel proof asset for docs, contract capture, readiness blocking, and prompt
handoff boundaries.
