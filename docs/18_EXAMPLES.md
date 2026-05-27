# Examples

These examples show how to use `ni` as a pre-runtime contract compiler. They
are not task-runner recipes. Each flow ends at readiness, lock, prompt, export,
or handoff material derived from the lock.

All commands assume the repository form:

```bash
go run ./cmd/ni <command>
```

If you have built or installed the binary, replace `go run ./cmd/ni` with
`ni`.

## Conversation product

Use this for an assistant, interview flow, support workflow, or other product
whose main delivery surface is a conversation.

```bash
go run ./cmd/ni init --dir ./examples/conversation-product \
  --product-type conversation_product \
  --surface conversation \
  --interaction-mode human_to_system
```

See the complete locked example at
[`examples/conversation-product/`](../examples/conversation-product/).

Planning focus:

- turn boundaries,
- memory expectations,
- refusals and escalation,
- transcript evaluation,
- human handoff.

Typical contract IDs:

```text
CAP-001 Define supported conversation turns
REQ-001 The assistant must identify unsupported requests and escalate
EVAL-001 Review transcript fixtures for correct escalation
RISK-001 The assistant may appear ready when handoff behavior is vague
NG-001 Do not connect live support queues before the plan is locked
```

Readiness and lock:

```bash
go run ./cmd/ni status --dir ./examples/conversation-product
go run ./cmd/ni end --dir ./examples/conversation-product
go run ./cmd/ni run --dir ./examples/conversation-product --target human-team
```

The compiled prompt is seed material for implementation. It does not start a
chat agent.

## Conversation authoring fixture

Use this to inspect how planning docs and `.ni/contract.json` are produced from
model-user conversation after `ni init`.

See the complete fixture at
[`examples/conversation-authoring/`](../examples/conversation-authoring/).

The transcript emphasizes:

- the user gives product intent in normal conversation,
- `ni-start` asks focused questions and maintains docs plus contract records,
- the user does not type contract `add`, `list`, or `set` commands,
- `ni status` reports gaps and readiness,
- `ni-end` confirms before the CLI writes the lock,
- `ni-run` compiles a prompt only and does not execute downstream work.

The checked-in example includes generated `docs/plan/**`, `.ni/contract.json`,
bounded `.ni/session.json` state, the CLI-written lockfile, and the
`human-team` prompt for the refund triage assistant plan.

## Ambiguous prompt blocked

Use this to see ni block a vague implementation request before any agent starts
work.

See the demo at
[`examples/ambiguous-prompt-blocked/`](../examples/ambiguous-prompt-blocked/).

The checked-in workspace is intentionally blocked:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
```

Expected payoff:

```text
This is ni's core payoff: it blocks ambiguous execution before the agent starts.
```

The demo does not execute Codex, implement the product, or add downstream
runner behavior. It shows how `ni-start` turns ambiguity into docs and contract
records, how `ni status` blocks on missing intent, and how `ni run` would later
compile a bounded target prompt after lock.

## Software CLI

Use this for a command-line tool, library, API, or normal software project.

```bash
go run ./cmd/ni init --dir ./examples/software-cli \
  --product-type software \
  --surface cli \
  --interaction-mode human_to_system
```

Planning focus:

- command surface,
- input and output contracts,
- error behavior,
- validation commands,
- runtime boundary and non-goals.

Typical contract IDs:

```text
CAP-001 Parse and validate CLI input
REQ-001 Invalid flags must exit nonzero with a specific error
EVAL-001 Unit tests cover valid and invalid flag parsing
RISK-001 CLI behavior could drift from docs
NG-001 Do not add a web UI for v0
```

Lock and compile:

```bash
go run ./cmd/ni status --dir ./examples/software-cli
go run ./cmd/ni end --dir ./examples/software-cli
go run ./cmd/ni run --dir ./examples/software-cli --target generic --out ./examples/software-cli.goal.txt
```

## Research protocol

Use this for a research method, evaluation plan, study protocol, or analysis
workflow where the primary output is documented evidence.

```bash
go run ./cmd/ni init --dir ./examples/research-protocol \
  --product-type research_protocol \
  --surface document \
  --interaction-mode human_to_human
```

See the complete locked example at
[`examples/research-protocol/`](../examples/research-protocol/).

Planning focus:

- hypothesis,
- data handling,
- method,
- analysis plan,
- ethics and reproducibility evidence.

Typical contract IDs:

```text
CAP-001 Define the study protocol
REQ-001 Data inclusion criteria must be explicit
EVAL-001 Independent reviewer can reproduce the sampling procedure
RISK-001 Ambiguous exclusion criteria may bias the result
NG-001 Do not collect participant data until protocol review is complete
```

The readiness gate still uses deterministic contract checks. It does not judge
research quality by model opinion.

```bash
go run ./cmd/ni status --dir ./examples/research-protocol --json
go run ./cmd/ni end --dir ./examples/research-protocol
go run ./cmd/ni run --dir ./examples/research-protocol --target human-team
```

## Human-team handoff

Use the `human-team` target when a locked plan needs to move to a PM, developer,
designer, researcher, reviewer, or operations team.

```bash
go run ./cmd/ni run --dir <locked-project> --target human-team --out ./handoff.prompt.txt
```

Expected handoff content:

- lock authority and hash checks,
- accepted capabilities and requirements,
- evaluation expectations,
- risks and mitigations,
- blocker handling,
- team ownership suggestions.

The handoff prompt may help the team plan work. It must not become an NI-owned
queue or execution ledger.

If downstream work produces feedback, record it inertly:

```bash
go run ./cmd/ni feedback add --dir <locked-project> --file ./team-feedback.json
go run ./cmd/ni pressure status --dir <locked-project>
```

Then create an amendment only when the planning owner accepts a contract change:

```bash
go run ./cmd/ni amend create --dir <locked-project> --title "Revise handoff validation"
go run ./cmd/ni amend apply AMEND-001 --dir <locked-project>
go run ./cmd/ni relock --dir <locked-project>
```

## Hyper Run seed export

Use the `hyper-run` export target when a locked NI plan should seed a downstream
Hyper Run workflow without NI becoming Hyper Run.

```bash
go run ./cmd/ni export --dir <locked-project> --target hyper-run --out ./hyper-run-seed
```

The export writes seed Markdown files such as:

```text
plan.md
ni-context.md
readiness-expectations.md
evidence-requirements.md
first-run-focus.md
```

The export does not create:

```text
.hyper/goals/GOAL-0001/
tasks.md
evidence.md
review.md
next.md
```

The seed package is derived from `.ni/plan.lock.json`. Downstream runtime state
belongs outside NI.
