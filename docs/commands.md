# Command Reference

This page is the detailed source-first command reference for `ni`.

`ni` is a Project Intent Compiler for AI Agents. Commands create and validate a
planning contract, lock accepted intent, and compile bounded downstream prompts
or inert seed material. They do not execute downstream agents, shell commands,
queues, adapters, PR automation, or release automation.

## Boundary

The kernel is authoritative for:

- `docs/plan/**`
- `.ni/contract.json`
- deterministic readiness validation
- `.ni/plan.lock.json`
- lock hash verification
- bounded prompt compilation
- inert downstream seed exports and proposals

After `.ni/plan.lock.json` exists, source-of-truth precedence is:

```text
.ni/plan.lock.json > .ni/contract.json > docs/plan/** > .ni/session.json > chat history
```

If locked hashes no longer match, target handoff commands stop with `BLOCKED`.

## Source Usage

Run the CLI from source:

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
[docs/22_INSTALL.md](22_INSTALL.md) for installation details.

## Command Reference Table

None of these commands execute downstream work. They do not run Codex, shell
commands, agents, adapters, queues, PR automation, release automation, target
runtimes, work graphs, or harness evidence.

| Command | Purpose | Example | Authority boundary | Lock behavior | Mutates kernel state? | Requires valid lock? | Can return `BLOCKED`? |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `ni --help` | Print implemented command usage. | `ni --help` | Informational CLI surface only. | Does not read, create, or verify a lock. | No. | No. | No. |
| `ni version` | Print the source version. | `ni version` | Informational CLI surface only. | Does not read, create, or verify a lock. | No. | No. | No. |
| `ni init` | Create a planning workspace with `docs/plan/**`, `.ni/contract.json`, readiness config, and bounded session state. | `ni init --dir ./my-plan --profile prototype` | Starts kernel planning state; it is not an interactive contract editor and does not add contract `add`, `list`, or `set` authoring commands. | Does not require or create `.ni/plan.lock.json`. | Yes, creates or preserves kernel planning files. | No. | No. |
| `ni status` | Evaluate deterministic readiness from docs, contract, sync, decisions, risks, evaluations, and blocker questions. | `ni status --dir .` | The authoritative readiness gate; model judgment cannot override it. | Does not create or verify lock hashes. | No. | No. | Yes, as the readiness result. |
| `ni status --proof` | Print the same readiness result with rule-level proof evidence. | `ni status --dir . --proof` | Authoritative readiness proof for planning conversations and release checks. | Does not create or verify lock hashes. | No. | No. | Yes, as the readiness result. |
| `ni status --next-questions` | Derive focused planning questions from readiness failures. | `ni status --dir . --next-questions` | Interview aid for model-user planning; questions do not change readiness by themselves. | Does not create or verify lock hashes. | No. | No. | Yes, as the readiness result. |
| `ni end` | Lock a ready accepted plan. | `ni end --dir .` | CLI authority for first lock; it runs readiness and refuses model-only readiness claims. | Writes `.ni/plan.lock.json` only when readiness is not `BLOCKED`; hashes `.ni/contract.json` and required `docs/plan/**`. | Yes, writes the lockfile. | No existing lock required. | Yes, when readiness is `BLOCKED`. |
| `ni run` | Compile a bounded downstream prompt from a locked plan. | `ni run --dir . --target codex --max-chars 4000` | Prompt compilation only; output is handoff text, not execution. | Verifies `.ni/plan.lock.json`; refuses stale hashes. | No kernel mutation; `--out` writes a prompt artifact only. | Yes. | Yes, on lock hash mismatch. |
| `ni targets` | List supported prompt and export targets. | `ni targets --json` | Target registry only; targets are consumption shapes, not runtime adapters. | Does not read, create, or verify a lock. | No. | No. | No. |
| `ni export` | Write downstream seed Markdown from a locked plan. | `ni export --dir . --target hyper-run --out ./seed` | Seed export only; exported files are derived and mutable downstream artifacts. | Verifies `.ni/plan.lock.json`; refuses stale hashes. | No kernel mutation; writes seed Markdown in `--out`. | Yes. | Yes, on lock hash mismatch. |
| `ni feedback` | Record or list inert downstream observations. | `ni feedback add --dir . --file ./feedback.json` | Evidence for future planning cycles; does not change the contract or lock. | If a lock exists, verifies it before reading or writing feedback. | `add` mutates `.ni/feedback.jsonl` and observed pressure; `list` does not. | Only when a lock exists. | Yes, on existing stale lock. |
| `ni pressure` | Track observed planning pressure and explicit promotion state. | `ni pressure status --dir .` | Pressure is advisory until a human planning decision changes docs and contract. | If a lock exists, verifies it before reading or writing pressure. | `promote` and `retire` mutate `.ni/pressure.json`; `status` does not. | Only when a lock exists. | Yes, on existing stale lock. |
| `ni amend` | Create, inspect, and apply explicit amendment records for locked-plan changes. | `ni amend create --dir . --title "Clarify acceptance criteria"` | Amendment records explain intent changes; they do not edit the contract, docs, or lock by themselves. | `create` records the current lock hash when one exists; `apply` refuses an amendment tied to a different source lock. | `create` and `apply` mutate `.ni/amendments/**`; `list` and `show` do not. | No. | Yes, when applying an amendment against a different source lock. |
| `ni relock` | Create a new lock after an explicitly amended plan. | `ni relock --dir .` | CLI authority for relocking; it preserves the amendment gate and readiness gate. | Requires an existing lock; archives the previous lock; refuses stale docs unless an applied amendment exists for the current lock; refuses blocked readiness. | Yes, archives the previous lock and writes a new `.ni/plan.lock.json`. | Requires an existing lock; current hashes may be stale only with an applied amendment. | Yes, on blocked readiness or stale lock without applied amendment. |
| `ni diff` | Show contract-level changes between two planning states. | `ni diff --base ./base --head ./head --json` | Informational collaboration check only; it does not resolve or apply changes. | May read lockfiles to orient inputs, but does not require valid locks or enforce lock gates. | No. | No. | No. |
| `ni conflicts` | Report semantic planning conflicts between two planning states. | `ni conflicts --base ./base --head ./head` | Collaboration guardrail only; it reports conflicts but does not merge, resolve, or mutate. | Reads lockfiles when present and reports lock hash mismatches as semantic conflicts. | No. | No. | Yes, as conflict severity and nonzero exit when blocking conflicts exist. |
| `ni graph` | Print a read-only work graph proposal from contract capabilities and artifacts. | `ni graph --dir . --json` | Inert proposal material only; not a task runner, queue, scheduler, or execution graph. | If a lock exists, verifies it before proposing the graph. | No. | Only when a lock exists. | Yes, on existing stale lock. |
| `ni harness` | Print or manage generated-harness proposal records. | `ni harness plan --dir .` | Inert proposal material only; not an evidence runner, adapter, queue, or kernel-owned execution state. | `plan`, `candidates`, `propose`, `validate`, `accept`, and `retire` require and verify a valid lock. | `propose`, `validate`, `accept`, and `retire` mutate `.ni/harness.candidates.json`; `plan` and `candidates` do not. | Yes. | Yes, on lock hash mismatch. |

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

If the status is `BLOCKED`, execution must not start.

Lock only after readiness passes:

```bash
go run ./cmd/ni end --dir <path>
```

Compile a bounded downstream prompt from the valid lock:

```bash
go run ./cmd/ni run --dir <path> --target codex --max-chars 4000
```

`ni run` prints or writes a prompt. It does not execute that prompt.

## Help and Version

```bash
go run ./cmd/ni --help
go run ./cmd/ni version
```

`ni --help` lists implemented commands and options. `ni version` prints the
source version.

## init

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

`--interaction-mode` accepts a lowercase identifier such as `human_to_system`
or `human_to_human`.

These fields guide planning and status output. They do not create runtime
stages or execution behavior.

## status

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

## end

`ni end` locks a ready plan.

```bash
go run ./cmd/ni end --dir <path>
```

It runs the readiness gate, refuses `BLOCKED`, and writes
`.ni/plan.lock.json` with hashes for `.ni/contract.json` and required
`docs/plan/**` files. `.ni/session.json` is not hashed because it is a mutable
planning aid below locked docs.

## run

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

## targets

Targets are consumption shapes for a locked plan. They are not integrations
that `ni` executes, runtime adapters that `ni` owns, or lifecycle state that
becomes part of `ni-kernel`.

See the [Target Story](45_TARGET_STORY.md) for target-by-target boundaries.

List supported prompt and export targets:

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

## export

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

## feedback

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

## pressure

`ni pressure` tracks recurring planning pressure without changing readiness
rules by itself.

```bash
go run ./cmd/ni pressure status --dir <path>
go run ./cmd/ni pressure status --dir <path> --json
go run ./cmd/ni pressure promote P-001 --dir <path>
go run ./cmd/ni pressure retire P-001 --dir <path>
```

Promotion is explicit and staged:

```text
observed -> repeated -> promotable -> accepted
```

Accepted pressure still requires a human planning decision before it changes a
locked contract.

## amend and relock

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

## diff and conflicts

`ni diff` and `ni conflicts` compare planning states without resolving or
mutating them.

```bash
go run ./cmd/ni diff --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni diff --base <path-or-lock> --head <path-or-lock> --json
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock>
go run ./cmd/ni conflicts --base <path-or-lock> --head <path-or-lock> --json
```

Inputs may be a project directory, `.ni/contract.json`, or
`.ni/plan.lock.json`. `ni conflicts` exits nonzero when blocking semantic
conflicts are found, including stale locks, conflicting decisions, weakened
accepted requirements, and risk severity reductions without mitigation context.

## graph and harness

`ni graph` and `ni harness` describe optional downstream work as inert
seed/proposal material. `ni graph` can propose from draft contract state and
verifies hashes if a lock is present. `ni harness` requires a valid lock for its
proposal and candidate lifecycle commands. Despite the command names, these
outputs are not a task runner, evidence runner, queue, adapter, or kernel-owned
execution state.

```bash
go run ./cmd/ni graph --dir <path>
go run ./cmd/ni graph --dir <path> --json
go run ./cmd/ni harness plan --dir <path>
go run ./cmd/ni harness plan --dir <path> --json
go run ./cmd/ni harness candidates --dir <path>
go run ./cmd/ni harness candidates --dir <path> --json
go run ./cmd/ni harness propose --dir <path> --from-pressure P-001
go run ./cmd/ni harness validate CAND-001 --dir <path> --evidence <path>
go run ./cmd/ni harness accept CAND-001 --dir <path>
go run ./cmd/ni harness retire CAND-001 --dir <path>
```

The kernel may compile work graphs, evaluation-plan proposals, evidence-rule
notes, and downstream handoff material from a valid lock. It must not execute
them.

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
[docs/48_DEMO_VERIFICATION.md](48_DEMO_VERIFICATION.md).
