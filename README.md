# ni

`ni` is a pre-runtime project intent compiler.

It turns planning work into a machine-readable contract, checks whether that
contract is ready, locks the accepted plan with hashes, and compiles short
handoff prompts or seed material from the locked plan.

The current product is `ni-kernel`, not an execution harness.

```text
conversation -> docs/plan + .ni/contract.json -> ni status -> ni end -> ni run
```

## Authoring model

`ni init` creates the planning workspace. After that, authoring happens through
conversation with a planning model and NI skills. The model maintains
`docs/plan/**` and `.ni/contract.json` from the ongoing conversation; users
should not have to manually edit contract JSON.

The CLI is still authoritative, but it is not the primary planning interface.
Use `ni status` to validate readiness, `ni end` to lock a ready plan, and
`ni run` to compile a bounded prompt from a valid lock.

Authoring rules are documented in:

- [Conversation authoring](docs/28_CONVERSATION_AUTHORING.md)
- [Authoring protocol](docs/29_AUTHORING_PROTOCOL.md)
- [Document update rules](docs/30_DOC_UPDATE_RULES.md)
- [ni-start behavior](docs/31_NI_START_BEHAVIOR.md)
- [Model edit safety](docs/37_MODEL_EDIT_SAFETY.md)

## JSON schemas

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

## What ni is

`ni` is the authority for:

- planning docs and the `.ni/contract.json` schema,
- deterministic readiness checks,
- `.ni/plan.lock.json`,
- source-of-truth ordering,
- prompt compilation from a locked plan,
- `.ni/session.json` as non-authoritative planning continuity state,
- downstream seed exports derived from a locked plan,
- inert feedback, pressure, amendment, and collaboration records.

After a plan is locked, the source-of-truth order is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

If a locked file hash no longer matches, `ni run`, `ni export`, feedback, and
pressure commands stop with `BLOCKED`.

## What ni is not

`ni` is not:

- a task runner,
- a SPEC runner,
- a multi-agent execution layer,
- a queue,
- a shell or Codex adapter,
- release automation,
- PR automation,
- a Hyper Run clone.

Generated harnesses, prompts, and export packages are derived and mutable. They
do not become kernel-owned execution state.

## Quickstart

### Go run mode

Use `go run` when trying the CLI directly from source:

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
```

Create a planning workspace:

```bash
tmp="$(mktemp -d)"
go run ./cmd/ni init --dir "$tmp/plan" --profile prototype
go run ./cmd/ni status --dir "$tmp/plan"
```

A new template workspace is expected to print `BLOCKED` because it still has
TODO values and an open blocker question.

Continue planning with `ni-start` in Codex or another model environment:

```text
User: Invoke ni-start for $tmp/plan. Help me finish the plan.
Model: Reads docs/plan/** and .ni/contract.json, summarizes state, asks focused questions.
User: Answers the focused questions.
Model: Updates docs/plan/** and .ni/contract.json together, then runs or requests ni status.
```

Keep the conversation going until the model has persisted the answers into the
planning docs and contract. Then check readiness again:

```bash
go run ./cmd/ni status --dir "$tmp/plan"
```

Only after `ni status` reports `READY` or `READY_WITH_DEFERRALS`, lock and
compile:

```bash
go run ./cmd/ni end --dir "$tmp/plan"
go run ./cmd/ni run --dir "$tmp/plan" --target codex --out "$tmp/codex.prompt.txt"
```

To inspect this repository's already locked plan:

```bash
go run ./cmd/ni status --dir .
go run ./cmd/ni run --dir . --target generic --max-chars 4000
```

## Examples

Complete locked example workspaces live in `examples/`:

- [Travel Concierge Triage](examples/conversation-product/): a non-software conversation product with human-team and Codex prompt artifacts.
- [Neighborhood Cooling Study Protocol](examples/research-protocol/): a non-software research protocol with human-team and generic prompt artifacts.
- [Namba AI Upgrade](examples/namba-ai-upgrade/): a software product planning example.

The non-software examples demonstrate that ni compiles product planning
contracts before implementation, not only software specs.

### Built binary mode

Build a local binary into `bin/ni`:

```bash
make build
./bin/ni --help
./bin/ni version
```

`make build` injects a git-derived version when git metadata is available.

### Local install mode

Install a local binary to `~/.local/bin/ni` by default:

```bash
make install-local
~/.local/bin/ni version
```

Override `PREFIX` or `BINDIR` to choose another install location. See
[docs/22_INSTALL.md](docs/22_INSTALL.md) for installation details.

## Core commands

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
go run ./cmd/ni status --dir <path> --next-questions
go run ./cmd/ni status --dir <path> --json --next-questions
```

Status values:

```text
BLOCKED
READY_WITH_DEFERRALS
READY
```

A model may explain the status, but it may not override it.
When `--next-questions` is present, `ni status` derives concise planning
questions from readiness rule failures so `ni-start` can ask about the next
specific gap.

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

`ni run` compiles a 4000-character-or-less prompt from a locked plan.

```bash
go run ./cmd/ni run --dir <path>
go run ./cmd/ni run --dir <path> --target codex
go run ./cmd/ni run --dir <path> --target human-team --out <file>
go run ./cmd/ni run --dir <path> --max-chars 2400
```

`ni run` does not execute Codex, shell commands, agents, queues, or adapters.

## Targets

List supported prompt/export targets:

```bash
go run ./cmd/ni targets
go run ./cmd/ni targets --json
```

Built-in targets:

```text
generic     prompt   general downstream implementation prompt
codex       prompt   Codex-oriented prompt seed
human-team  handoff  team handoff prompt
hyper-run   seed     Hyper Run seed material
namba-ai    seed     namba-ai seed material
ouroboros   seed     Ouroboros seed notes
spec-kit    seed     Spec Kit seed notes
```

Targets are downstream shapes. They do not change kernel authority.

## Export

`ni export` writes locked-plan seed packages for supported downstream targets.

```bash
go run ./cmd/ni export --dir <path> --target hyper-run --out <dir>
go run ./cmd/ni export --dir <path> --target namba-ai --out <dir>
go run ./cmd/ni export --dir <path> --target ouroboros --out <dir>
go run ./cmd/ni export --dir <path> --target spec-kit --out <dir>
```

Export requires `.ni/plan.lock.json`, verifies locked hashes, and refuses stale
plans with `BLOCKED`. It writes seed Markdown only. It does not call external
runtimes or create downstream runtime packets.

## Feedback

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

## Pressure

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

## Amend and relock

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

## Collaboration conflict checks

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

## Generated harness

`ni graph` and `ni harness` describe possible downstream work from a locked
contract. They remain inert planning artifacts.

```bash
go run ./cmd/ni graph --dir <path>
go run ./cmd/ni harness plan --dir <path>
go run ./cmd/ni harness candidates --dir <path>
go run ./cmd/ni harness propose --dir <path> --from-pressure P-001
go run ./cmd/ni harness validate CAND-001 --dir <path> --evidence <path>
go run ./cmd/ni harness accept CAND-001 --dir <path>
go run ./cmd/ni harness retire CAND-001 --dir <path>
```

The kernel may propose work graphs, evaluation plans, and evidence rules. It
must not execute them.

## Development validation

When Go code exists, run:

```bash
make test
make quality
make smoke
make build
```

`make quality` runs `scripts/quality.sh`, which already runs formatting, Go
tests, JSON checks, Markdown fence checks, skill metadata checks, prompt budget
checks, core-boundary self-tests, and smoke tests.

## License

TODO: choose an open-source license with the project owner, then add a root
`LICENSE` file. Until then, this repository should not be treated as legally
reusable open-source software.
