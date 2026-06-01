# Example Coverage

This matrix records what each public example proves, how it is verified, and
where the non-execution boundary sits. The examples are Project Intent
Compiler assets: they validate planning contracts, grouped readiness proof,
locks, and prompt compilation boundaries. They do not implement the example
products or execute downstream agents.

## Verification command

Run the public demo check from the repository root:

```bash
bash scripts/demo-check.sh
```

Run the broader repository check:

```bash
bash scripts/quality.sh
```

For v0.5 example work, use the lane-specific acceptance criteria in
[`95_V0_5_ACCEPTANCE_EVIDENCE.md`](95_V0_5_ACCEPTANCE_EVIDENCE.md) before
claiming that an example lane is complete.

For benchmark examples, keep the claim-boundary labels from
[`97_BENCHMARK_CLAIM_BOUNDARIES.md`](97_BENCHMARK_CLAIM_BOUNDARIES.md) next to
every `READY` transition and `not_measured` table.

## Coverage matrix

| Example | Proves | Product type | Delivery surface | Expected status | demo-check coverage | Docs-only? | Korean companion |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `examples/ambiguous-prompt-blocked/` | Vague execution is blocked before handoff; grouped open-blocker questions guide the next turn. | `software` | `web` | `BLOCKED` | Runs `ni status` and grouped next-question rendering. | No, blocked workspace fixture. | Yes |
| `examples/research-protocol/` | ni is not software-only; a research protocol can lock before fieldwork. | `research_protocol` | `document` | `READY` | Runs status and compiles `human-team` prompt from the existing lock. | No, locked workspace fixture. | Yes |
| `examples/conversation-product/` | Conversation-surface planning can lock without becoming a chatbot runner. | `conversation_product` | `conversation` | `READY` | Runs status, compiles `human-team` prompt, and checks seed-only exports. | No, locked workspace fixture. | Yes |
| `examples/conversation-authoring/` | Sustained model-user authoring updates docs, contract, and session while CLI proof catches stale sync. | `conversation_product` | `conversation`, `document` | `BLOCKED` | Runs status, checks `R012`, and compiles from the historical lock only. | No, blocked fixture with historical lock material. | Yes |
| `examples/namba-ai-upgrade/` | ni can plan upstream of an existing harness/workflow project without becoming that harness. | `software` | `cli`, `document`, `workflow` | `BLOCKED` | Runs status, checks `R012`, and compiles Codex prompt from the historical lock only. | No, blocked fixture with historical lock material. | Yes |
| `examples/ni-start-dogfood/` | First-run card, grouped next questions, docs/contract/session update, and re-status loop. | `conversation_product` | `conversation`, `document` | `READY_WITH_DEFERRALS` | Runs status, grouped proof, and compiles `human-team` prompt from the existing lock. | No, locked workspace fixture. | Yes |
| `examples/ni-grill/` | Pre-lock planning challenge UX: deterministic blockers first, then severity-labeled, budgeted `GRILL-*` questions against accepted or nearly accepted content. Also dogfoods benchmark grills and ni's own current planning state to test claim boundaries, roadmap specificity, and `not_measured` visibility without making new empirical claims. | draft `conversation_product`; benchmark evidence review; ni planning review | draft `conversation`, `document`; benchmark workspaces remain isolated | Not claimed | Verifies required docs, benchmark grill files, ni-project dogfood file, severity-labeled `GRILL-*` findings, lessons, and non-execution boundary wording only. | Yes, transcript fixture plus benchmark and ni-project review transcripts. | Yes |
| `examples/benchmark-report/` | Benchmark/case-study reporting method with `not_measured` boundaries, plus measured dashboard and research-protocol cases. The dashboard case packages a blocked-to-ready artifact-readiness transition; the research-protocol case preserves initial `BLOCKED` readiness, then applies synthetic fixture answers and reaches isolated-workspace `READY` with lock and bounded prompt proof. Claim boundaries follow `docs/97_BENCHMARK_CLAIM_BOUNDARIES.md`. | Dashboard artifact case: `document_product`; research case: `research_protocol` | Dashboard: `document`; research: `document`, `workflow`, `human_service` | Dashboard case: `READY` for benchmark artifact readiness only; research case: `READY` for synthetic benchmark fixture readiness only | Verifies required docs, dashboard historical blocked/resolved proof, research initial blocked proof, research resolved proof, blocker/next-question evidence, blocker analysis, resolution path, answer packets, lock/prompt summaries, before/after evidence, lessons, and remaining `not_measured` boundaries for product/runtime/research claims. It also checks the shared claim-boundary marker document. | Partial: report template is docs-only; dashboard case has a locked ni workspace; research case now has a locked isolated benchmark workspace after synthetic fixture answers. | Yes |
| `examples/no-terminal-assisted/` | Assisted planning can draft docs and contract before local CLI validation, show a model-workspace start flow, and hand off to later CLI proof without deterministic readiness claims. | draft `workflow` | draft `document` | Not claimed | Verifies required files, docs-only status, and boundary wording only. | Yes, assisted draft. | Yes |

## Grouped next-question coverage

The grouped `ni status --proof --next-questions` UX is shown directly in:

- `examples/ambiguous-prompt-blocked/05-next-questions.md`
- `examples/benchmark-report/cases/internal-dashboard/06-ni-status-proof.md`
- `examples/benchmark-report/cases/internal-dashboard/07-ni-next-questions.md`
- `examples/benchmark-report/cases/internal-dashboard/11-resolved-status-proof.md`
- `examples/benchmark-report/cases/internal-dashboard/12-resolved-next-questions.md`
- `examples/benchmark-report/cases/research-protocol/06-ni-status-proof.md`
- `examples/benchmark-report/cases/research-protocol/07-ni-next-questions.md`
- `examples/benchmark-report/cases/research-protocol/11-resolved-status-proof.md`
- `examples/benchmark-report/cases/research-protocol/12-resolved-next-questions.md`
- `examples/conversation-authoring/transcript.md`
- `examples/conversation-authoring/session-resume.md`
- `examples/ni-start-dogfood/03-model-summary-and-questions.md`
- `examples/ni-start-dogfood/06-status-proof.md`
- `examples/ni-start-dogfood/07-second-round-questions.md`
- `examples/ni-start-dogfood/README.md`
- `examples/ni-grill/02-grill-questions.md`
- `examples/ni-grill/05-status-after-grill.md`
- `examples/ni-grill/06-internal-dashboard-grill.md`
- `examples/ni-grill/07-research-protocol-grill.md`
- `examples/ni-grill/09-ni-project-grill.md`

The expected model behavior is to preserve group labels, ask the
highest-priority group first, use CLI answer shapes, update `docs/plan/**`,
`.ni/contract.json`, and `.ni/session.json` after the user answers, and run or
request `ni status --dir . --proof --next-questions` again.

## Non-execution boundary

The examples do not:

- implement dashboards, assistants, research studies, travel workflows, or
  namba-ai changes;
- call Codex, Claude APIs, model APIs, downstream agents, or shell adapters;
- create queues, runtime execution, release automation, PR automation, or
  evidence runners;
- add user-facing contract `add`, `list`, or `set` authoring commands;
- fake benchmark results or claim statistical significance.

Locked examples may compile inert prompts through `ni run` because the CLI
verifies existing lock material first. Blocked examples remain blocked even
when historical lock or generated prompt material exists.

The no-terminal assisted example remains docs-only by design. It demonstrates
drafting, team handoff, and graduation to full `ni`, but it does not run
`ni status`, `ni end`, or `ni run` because no-terminal mode is not a trusted
CLI workspace.

The ni-grill benchmark and ni-project dogfood transcripts are also
non-execution examples. They record `ni status` proof and challenge planning or
benchmark claim boundaries, but they do not run generated prompts, call model
APIs, perform fieldwork, implement products, publish releases, edit lockfiles,
or create empirical claims.
