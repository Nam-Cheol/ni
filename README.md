# ni

[English](README.md) | [한국어](README.ko.md)

Project Intent Compiler for AI Agents.

Don't run the agent yet. Compile the intent first.

`ni` turns planning conversations into a locked, versioned, verifiable project
contract before Codex, Claude, Spec Kit, Hyper Run, namba-ai, a generated
harness, or a human team starts execution.

The current product is `ni-kernel`: a deterministic pre-runtime control layer
for intent, not an execution harness.

```text
conversation -> docs/plan + .ni/contract.json -> ni status -> ni end -> locked intent -> ni run
```

Start with the category docs:

- [Positioning](docs/40_POSITIONING.md)
- [Differentiation map](docs/41_DIFFERENTIATION.md)
- [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md)

## What Problem ni Solves

Agents are often asked to begin from prompts that sound actionable but still
hide critical intent:

- Who is the project for?
- What must be true before work is accepted?
- Which risks require mitigation?
- What is explicitly out of scope?
- Which questions should block execution?
- Has the plan changed since it was accepted?

Most tools try to control the agent after a prompt, spec, worklist, or runtime
loop already exists. `ni` moves control earlier. It asks whether
project intent is explicit, accepted, validated, locked, and unchanged before
any downstream actor starts work.

## The ni Answer: Intent Lock Protocol

The [Intent Lock Protocol](docs/42_INTENT_LOCK_PROTOCOL.md) is the core
mechanism of `ni-kernel`. It defines:

1. how planning conversations become a project contract,
2. when the contract is ready to lock,
3. how the accepted plan is hashed,
4. what downstream actors may trust,
5. when execution must stop because intent changed.

The kernel owns:

- `docs/plan/**`
- `.ni/contract.json`
- deterministic readiness validation
- `.ni/plan.lock.json`
- lock hash verification
- bounded prompt compilation
- inert downstream seed exports

After a plan is locked, source-of-truth precedence is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

If locked hashes no longer match, `ni run`, `ni export`, feedback, pressure,
and downstream handoff commands stop with `BLOCKED`.

## 5-Minute Demo: Ambiguous Prompt Blocked

The fastest way to understand `ni` is the
[ambiguous prompt blocked demo](examples/ambiguous-prompt-blocked/).

It starts from a vague request:

```text
Build me a dashboard for my team.
```

A direct-to-agent path would force the agent to invent hidden assumptions about
users, data, workflow, non-goals, and success criteria. The `ni` path records
the request as planning intent, then refuses to treat it as executable while
blocker questions remain open.

Try the blocked workspace:

```bash
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace
go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions
```

Expected result:

```text
BLOCKED
```

That is the point: ambiguous execution is blocked before an agent starts.

## Non-Software Demo

`ni` is not a software spec generator. It compiles project intent for any
product surface.

Try the [Neighborhood Cooling Study Protocol](examples/research-protocol/):

```bash
go run ./cmd/ni status --dir examples/research-protocol
go run ./cmd/ni run --dir examples/research-protocol --target human-team --out examples/research-protocol/generated/human-team.prompt.md
```

This locked example plans a research protocol, not an app. It has
`product_type: research_protocol`, a `document` delivery surface, protocol
review evaluations, and a human-team handoff prompt. It does not collect data,
run analysis, deploy sensors, or execute fieldwork.

Another non-software example is
[Travel Concierge Triage](examples/conversation-product/), a conversation
product that compiles a human concierge handoff without deploying a chatbot or
booking travel.

Verify all public README demos from source:

```bash
bash scripts/demo-check.sh
```

The demo check asserts the intentionally blocked state for the ambiguous prompt
demo and compiles locked example prompts to temporary files only. It does not
execute downstream agents.

## Core Flow

Create a planning workspace:

```bash
go run ./cmd/ni init --dir <path> --profile prototype
```

Use sustained model-user conversation to maintain `docs/plan/**` and
`.ni/contract.json` together. Skills and models are UX. The CLI is authority.

Check readiness:

```bash
go run ./cmd/ni status --dir <path>
go run ./cmd/ni status --dir <path> --proof --next-questions
```

A model may explain the status, but it may not override it. If the status is
`BLOCKED`, execution must not start.

Lock only after readiness passes:

```bash
go run ./cmd/ni end --dir <path>
```

Compile a bounded downstream prompt from the valid lock:

```bash
go run ./cmd/ni run --dir <path> --target codex --max-chars 4000
```

`ni run` prints or writes a prompt. It does not execute that prompt.

## What ni Blocks

`ni` blocks downstream handoff when intent is not trustworthy yet:

- accepted capabilities without linked evaluations,
- high-severity risks without mitigation,
- open blocker questions,
- conflicting accepted decisions,
- missing or invalid required planning records,
- stale locks where current files no longer match `.ni/plan.lock.json`,
- target prompt compilation before a valid lock exists.

The benchmark protocol in
[docs/43_BENCHMARK_PROTOCOL.md](docs/43_BENCHMARK_PROTOCOL.md) describes how
to compare direct-to-agent prompts against locked `ni` intent without running
downstream agents.

## Core Commands

### Help and Version

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
```

### init

`ni init` creates the planning docs and `.ni` skeleton. It does not start a
contract editing session; the intended authoring flow is model-assisted
conversation after workspace creation.

```bash
go run ./cmd/ni init --dir <path>
go run ./cmd/ni init --dir <path> --profile concept
go run ./cmd/ni init --dir <path> --product-type conversation_product --surface conversation --interaction-mode human_to_system
```

Supported readiness profiles:

```text
concept
prototype
mvp
beta
production
```

Supported product types:

```text
software
conversation_product
research_protocol
operations_process
education_program
document_product
physical_product
mixed
```

Supported delivery surfaces:

```text
web
cli
api
conversation
document
workflow
human_service
physical
```

These fields guide planning and status output. They do not create runtime
stages or execution behavior.

### status

`ni status` evaluates readiness from deterministic rules.

```bash
go run ./cmd/ni status --dir <path>
go run ./cmd/ni status --dir <path> --json
go run ./cmd/ni status --dir <path> --proof
go run ./cmd/ni status --dir <path> --proof --json
go run ./cmd/ni status --dir <path> --next-questions
go run ./cmd/ni status --dir <path> --json --next-questions
```

Status values:

```text
BLOCKED
READY_WITH_DEFERRALS
READY
```

When `--proof` is present, `ni status` prints rule-level evidence from the
readiness, docs/contract sync, and accepted-decision conflict checks.

When `--next-questions` is present, `ni status` derives concise planning
questions from readiness rule failures so a planning conversation can address
the next specific gap.

### end

`ni end` locks a ready plan.

```bash
go run ./cmd/ni end --dir <path>
```

It runs the readiness gate, refuses `BLOCKED`, and writes
`.ni/plan.lock.json` with hashes for `.ni/contract.json` and required
`docs/plan/**` files. `.ni/session.json` is not hashed because it is a mutable
planning aid below locked docs.

### run

`ni run` compiles a prompt from a locked plan.

```bash
go run ./cmd/ni run --dir <path>
go run ./cmd/ni run --dir <path> --target codex
go run ./cmd/ni run --dir <path> --target human-team --out <file>
go run ./cmd/ni run --dir <path> --max-chars 2400
```

Prompt output must stay within the configured maximum, initially 4000
characters. `ni run` does not execute Codex, shell commands, agents, queues, or
adapters.

## Targets and Exports

Targets are consumption shapes for a locked plan. They are not integrations
that `ni` executes, runtime adapters that `ni` owns, or lifecycle state that
becomes part of `ni-kernel`.

See the [target story](docs/45_TARGET_STORY.md) for target-by-target
boundaries.

List supported prompt/export targets:

```bash
go run ./cmd/ni targets
go run ./cmd/ni targets --json
```

Built-in targets:

```text
generic     prompt   general downstream implementation prompt
codex       prompt   bounded implementation prompt seed
human-team  handoff  planning handoff for people
hyper-run   seed     seed material, not .hyper/goals runtime packets
namba-ai    seed     planning seed and suggested graph boundaries
ouroboros   seed     upstream intent notes, not Agent OS execution state
spec-kit    seed     upstream intent summary, not Spec Kit workflow state
```

`ni export` writes locked-plan seed packages for supported downstream targets.
Those outputs are derived from a locked plan and remain mutable downstream
artifacts.

```bash
go run ./cmd/ni export --dir <path> --target hyper-run --out <dir>
go run ./cmd/ni export --dir <path> --target namba-ai --out <dir>
go run ./cmd/ni export --dir <path> --target ouroboros --out <dir>
go run ./cmd/ni export --dir <path> --target spec-kit --out <dir>
```

Export requires `.ni/plan.lock.json`, verifies locked hashes, and refuses stale
plans with `BLOCKED`. It writes seed Markdown only. It does not call external
runtimes, create downstream runtime packets, or add target adapters.

## Additional Kernel Commands

These commands preserve the same boundary: they read or write kernel records
and inert proposals, but they do not execute downstream work.

### feedback

`ni feedback` records downstream observations without mutating the contract or
lock.

```bash
go run ./cmd/ni feedback add --dir <path> --file testdata/feedback/codex.json
go run ./cmd/ni feedback list --dir <path>
go run ./cmd/ni feedback list --dir <path> --json
```

Feedback is appended to `.ni/feedback.jsonl` and translated into observed
pressure items. It is evidence for a future planning cycle, not an automatic
contract change.

### pressure

`ni pressure` tracks recurring planning pressure without changing readiness
rules by itself.

```bash
go run ./cmd/ni pressure status --dir <path>
go run ./cmd/ni pressure promote P-001 --dir <path>
go run ./cmd/ni pressure retire P-001 --dir <path>
```

Promotion is explicit and staged:

```text
observed -> repeated -> promotable -> accepted
```

Accepted pressure still requires a human planning decision before it changes a
locked contract.

### amend and relock

Locked planning docs must not be silently edited. Use amendments to explain why
a locked plan changed, then relock.

```bash
go run ./cmd/ni amend create --dir <path> --title "Clarify acceptance criteria"
go run ./cmd/ni amend list --dir <path>
go run ./cmd/ni amend show AMEND-001 --dir <path>
go run ./cmd/ni amend apply AMEND-001 --dir <path>
go run ./cmd/ni relock --dir <path>
```

An applied amendment must include a reason, affected docs or contract IDs,
proposed changes, risk impact, and readiness impact. `ni relock` refuses stale
locks without an applied amendment and refuses blocked readiness.

### diff and conflicts

`ni diff` and `ni conflicts` compare planning states without resolving or
mutating them.

```bash
go run ./cmd/ni diff --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock> --json
```

Inputs may be a project directory, `.ni/contract.json`, or
`.ni/plan.lock.json`. `ni conflicts` exits nonzero when blocking semantic
conflicts are found, including stale locks, conflicting decisions, weakened
accepted requirements, and risk severity reductions without mitigation context.

### graph and harness

`ni graph` and `ni harness` describe optional downstream work from a locked
contract. Despite the command names, these outputs are inert seed/proposal
material. They are not a task runner, evidence runner, queue, adapter, or
kernel-owned execution state.

```bash
go run ./cmd/ni graph --dir <path>
go run ./cmd/ni harness plan --dir <path>
go run ./cmd/ni harness candidates --dir <path>
go run ./cmd/ni harness propose --dir <path> --from-pressure P-001
go run ./cmd/ni harness validate CAND-001 --dir <path> --evidence <path>
go run ./cmd/ni harness accept CAND-001 --dir <path>
go run ./cmd/ni harness retire CAND-001 --dir <path>
```

The kernel may compile work graphs, evaluation-plan proposals, evidence-rule
notes, and downstream handoff material from a valid lock. It must not execute
them.

## What ni Is Not

`ni` is not:

- a task runner,
- a SPEC runner,
- a multi-agent execution layer,
- a Codex adapter,
- a queue,
- a shell adapter,
- release automation,
- PR automation,
- Hyper Run, Spec Kit, Ouroboros, or namba-ai.

Downstream prompts, seed packages, and harness proposals are derived and
mutable. They do not become kernel-owned execution state.

See the [differentiation map](docs/41_DIFFERENTIATION.md) for how `ni`
differs from host enhancers, SDD toolkits, coding-agent operating systems, and
execution growth runtimes.

## Examples

Complete example workspaces live in `examples/`:

- [Ambiguous Prompt Blocked](examples/ambiguous-prompt-blocked/): the core
  blocking demo for vague requests.
- [Neighborhood Cooling Study Protocol](examples/research-protocol/): a
  non-software research protocol with human-team and generic prompt artifacts.
- [Travel Concierge Triage](examples/conversation-product/): a conversation
  product with human-team and Codex prompt artifacts.
- [Conversation Authoring Fixture](examples/conversation-authoring/): an
  end-to-end transcript showing model-maintained docs and contract records
  after `ni init`, with CLI validation, lock, and prompt compilation.
- [Namba AI Upgrade](examples/namba-ai-upgrade/): a software product planning
  example.
- [Benchmark Report Template](examples/benchmark-report/): a manual report
  template for the pre-runtime intent readiness benchmark.

## Development Status

`ni` is currently source-first. Package publishing, Homebrew taps, GoReleaser,
and automated release tooling are outside the current kernel scope.

Use `go run` when trying the CLI directly from source:

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
go run ./cmd/ni status --dir .
```

Build a local binary into `bin/ni`:

```bash
make build
./bin/ni --help
./bin/ni version
```

Install a local binary to `~/.local/bin/ni` by default:

```bash
make install-local
~/.local/bin/ni version
```

Override `PREFIX` or `BINDIR` to choose another install location. See
[docs/22_INSTALL.md](docs/22_INSTALL.md) for installation details. The Korean
companion README is maintained at [README.ko.md](README.ko.md).

## JSON Schemas

Versioned JSON Schemas for NI state files live in `schema/`:

```text
schema/ni.project.v0.json
schema/ni.contract.v0.json
schema/ni.lock.v0.json
schema/ni.readiness-rules.v0.json
schema/ni.readiness-profiles.v0.json
schema/ni.feedback.v0.json
schema/ni.pressure.v0.json
schema/ni.amendment.v0.json
schema/ni.harness-candidate.v0.json
```

Validate the published schemas and current `.ni` state files with:

```bash
python3 scripts/check-schema.py
```

## Validation

When Go code exists, run:

```bash
gofmt -w .
go test ./...
bash scripts/quality.sh
```

For this repository, the main quality entry point is:

```bash
bash scripts/quality.sh
```

`scripts/quality.sh` runs formatting checks, Go tests, JSON checks, Markdown
fence checks, skill metadata checks, prompt budget checks, core-boundary
self-tests, and smoke tests.

Public demo verification is a separate release proof check:

```bash
bash scripts/demo-check.sh
```

Run it when README demos, example workspaces, status output, or prompt
compilation behavior changes. Details live in
[docs/48_DEMO_VERIFICATION.md](docs/48_DEMO_VERIFICATION.md).

## License and Release Status

`ni` is licensed under the [MIT License](LICENSE).

This README does not claim package distribution or a published binary release.
Use source, local build, or local install mode unless a release process says
otherwise.

Release readiness notes live in
[docs/46_RELEASE_READINESS.md](docs/46_RELEASE_READINESS.md). CI is defined in
[.github/workflows/ci.yml](.github/workflows/ci.yml). The project security
policy is [SECURITY.md](SECURITY.md).
