#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SMOKE_TMP="$(mktemp -d "${TMPDIR:-/tmp}/ni-smoke.XXXXXX")"
NI_BIN="$SMOKE_TMP/bin/namba-intent"
LAST_STDOUT="$SMOKE_TMP/stdout.log"
LAST_STDERR="$SMOKE_TMP/stderr.log"

trap 'rm -rf "$SMOKE_TMP"' EXIT

mkdir -p "$SMOKE_TMP/bin" "$SMOKE_TMP/workspaces"

run_cmd() {
  local label="$1"
  shift
  echo "smoke: $label" >&2
  : >"$LAST_STDOUT"
  : >"$LAST_STDERR"
  if ! "$@" >"$LAST_STDOUT" 2>"$LAST_STDERR"; then
    echo "smoke failed: $label" >&2
    echo "--- stdout ---" >&2
    sed -n '1,120p' "$LAST_STDOUT" >&2
    echo "--- stderr ---" >&2
    sed -n '1,120p' "$LAST_STDERR" >&2
    return 1
  fi
}

require_output() {
  local expected="$1"
  if ! grep -Fq "$expected" "$LAST_STDOUT"; then
    echo "smoke failed: expected stdout to contain: $expected" >&2
    echo "--- stdout ---" >&2
    sed -n '1,120p' "$LAST_STDOUT" >&2
    return 1
  fi
}

write_ready_contract() {
  local dir="$1"
  cat >"$dir/.ni/contract.json" <<'JSON'
{
  "schema": "ni.contract.v0",
  "readiness_profile": "prototype",
  "product_type": "software",
  "delivery_surfaces": ["cli"],
  "interaction_mode": "human_to_system",
  "project": {
    "id": "smoke-fixture",
    "name": "Smoke Fixture",
    "purpose": "Exercise public namba-intent commands.",
    "status": "draft"
  },
  "non_goals": [
    {
      "id": "NG-001",
      "title": "Do not execute external tools."
    }
  ],
  "capabilities": [
    {
      "id": "CAP-001",
      "title": "Public command smoke coverage",
      "status": "accepted",
      "requirements": ["REQ-001"],
      "evaluations": ["EVAL-001"],
      "risks": ["RISK-001"],
      "artifacts": ["ART-001"]
    }
  ],
  "requirements": [
    {
      "id": "REQ-001",
      "title": "Every public command has a success-path smoke invocation.",
      "status": "accepted"
    }
  ],
  "decisions": [
    {
      "id": "DEC-001",
      "title": "Smoke tests stay inside Namba Intent kernel boundaries.",
      "status": "accepted"
    }
  ],
  "risks": [
    {
      "id": "RISK-001",
      "title": "Smoke fixture risk",
      "severity": "high",
      "status": "accepted",
      "mitigation": "The fixture stays inside local Namba Intent commands."
    }
  ],
  "evaluations": [
    {
      "id": "EVAL-001",
      "title": "Run scripts/smoke.sh",
      "method": "local shell smoke test"
    }
  ],
  "artifacts": [
    {
      "id": "ART-001",
      "path": "scripts/smoke.sh",
      "kind": "test_script"
    }
  ],
  "open_questions": []
}
JSON
  cat >"$dir/docs/plan/00_project_brief.md" <<'MD'
# Project brief

## Product type

software

## Delivery surfaces

- cli

## Purpose

Exercise public namba-intent commands.
MD
  cat >"$dir/docs/plan/01_actors_outcomes.md" <<'MD'
# Actors and outcomes

## Actors

- User: reviews public command smoke coverage.
- CLI: validates readiness and lock state.

## Outcomes

- Public ni commands have success-path smoke coverage.
MD
  cat >"$dir/docs/plan/08_delivery_operation.md" <<'MD'
# Delivery and operation

## Delivery surfaces

- cli

## Initial delivery

The smoke fixture is reviewed before lock.
MD
  cat >"$dir/docs/plan/10_open_questions.md" <<'MD'
# Open questions

No open questions are listed in this fixture.
MD
}

complete_amendment() {
  local path="$1"
  python3 - "$path" <<'PY'
import json
import sys

path = sys.argv[1]
with open(path, "r", encoding="utf-8") as f:
    payload = json.load(f)

payload.update({
    "reason": "Smoke fixture amendment exercises the public amendment lifecycle.",
    "affected_docs": ["docs/plan/02_capabilities.md"],
    "affected_contract_ids": ["CAP-001", "REQ-001"],
    "proposed_changes": ["Clarify smoke fixture planning text without weakening readiness."],
    "risk_impact": "No external execution or new high-severity risk.",
    "readiness_impact": "Readiness remains READY under deterministic validation.",
    "created_from_feedback_refs": [],
    "created_from_pressure_refs": []
})

with open(path, "w", encoding="utf-8") as f:
    json.dump(payload, f, indent=2)
    f.write("\n")
PY
}

write_harness_pressure() {
  local dir="$1"
  cat >"$dir/.ni/pressure.json" <<'JSON'
{
  "schema": "ni.pressure.v0",
  "items": [
    {
      "id": "P-001",
      "kind": "harness_candidate",
      "status": "accepted",
      "evidence_refs": ["smoke:test"],
      "related_capabilities": ["CAP-001"],
      "related_risks": [],
      "proposed_action": "Create an inert downstream harness proposal for smoke coverage.",
      "requires_user_acceptance": true
    }
  ]
}
JSON
}

make_ready_workspace() {
  local name="$1"
  local dir="$SMOKE_TMP/workspaces/$name"
  run_cmd "namba-intent init ($name)" "$NI_BIN" init --dir "$dir" --profile prototype
  write_ready_contract "$dir"
  echo "$dir"
}

make_locked_workspace() {
  local name="$1"
  local dir
  dir="$(make_ready_workspace "$name")"
  run_cmd "namba-intent end ($name)" "$NI_BIN" end --dir "$dir"
  echo "$dir"
}

cd "$ROOT"

echo "smoke: building namba-intent" >&2
go build -o "$NI_BIN" ./cmd/namba-intent

run_cmd "namba-intent --help" "$NI_BIN" --help
require_output "Namba Intent is a Project Intent Compiler for AI Agents."

run_cmd "namba-intent version" "$NI_BIN" version
require_output "0.0.0-dev"

init_ws="$SMOKE_TMP/workspaces/init"
run_cmd "namba-intent init" "$NI_BIN" init --dir "$init_ws" --profile concept
require_output "initialized Namba Intent planning workspace"

run_cmd "namba-intent status" "$NI_BIN" status --dir "$init_ws"
require_output "BLOCKED"

run_cmd "namba-intent status proof" "$NI_BIN" status --dir "$init_ws" --proof --next-questions
require_output "NI Intent Readiness: BLOCKED"
require_output "Blockers:"
require_output "Next: answer or defer the blocker question"
require_output "Passed checks:"
require_output "Execution must not start."
require_output "Next questions:"

ready_ws="$(make_ready_workspace "ready")"
run_cmd "namba-intent end on ready fixture" "$NI_BIN" end --dir "$ready_ws"
require_output "locked plan"

locked_ws="$(make_locked_workspace "locked")"
for prompt_target in generic codex human-team; do
  run_cmd "namba-intent run --target $prompt_target" "$NI_BIN" run --dir "$locked_ws" --target "$prompt_target"
  require_output "Target: $prompt_target"
done

run_cmd "namba-intent targets" "$NI_BIN" targets
require_output "generic"
require_output "spec-kit"

for example_dir in "$ROOT/examples/conversation-product" "$ROOT/examples/research-protocol"; do
  example_name="$(basename "$example_dir")"
  run_cmd "example status ($example_name)" "$NI_BIN" status --dir "$example_dir"
  require_output "READY"
  run_cmd "example human-team prompt ($example_name)" "$NI_BIN" run --dir "$example_dir" --target human-team
  require_output "Target: human-team"
done

for export_target in hyper-run namba-ai ouroboros spec-kit; do
  run_cmd "namba-intent export --target $export_target" "$NI_BIN" export --dir "$locked_ws" --target "$export_target" --out "$SMOKE_TMP/exports/$export_target"
  require_output "exported $export_target seed package"
  run_cmd "target conformance ($export_target)" python3 "$ROOT/scripts/check-target-conformance.py" --target "$export_target" --dir "$SMOKE_TMP/exports/$export_target"
  require_output "$export_target export conforms to seed-only boundary"
done

feedback_ws="$SMOKE_TMP/workspaces/feedback"
run_cmd "namba-intent init (feedback)" "$NI_BIN" init --dir "$feedback_ws"
run_cmd "namba-intent feedback add" "$NI_BIN" feedback add --dir "$feedback_ws" --file "$ROOT/testdata/feedback/codex.json"
require_output "recorded feedback from codex"
run_cmd "namba-intent feedback list" "$NI_BIN" feedback list --dir "$feedback_ws"
require_output "codex"

pressure_ws="$SMOKE_TMP/workspaces/pressure"
run_cmd "namba-intent init (pressure)" "$NI_BIN" init --dir "$pressure_ws"
run_cmd "namba-intent feedback add (pressure fixture)" "$NI_BIN" feedback add --dir "$pressure_ws" --file "$ROOT/testdata/feedback/codex.json"
run_cmd "namba-intent pressure status" "$NI_BIN" pressure status --dir "$pressure_ws"
require_output "P-001"
run_cmd "namba-intent pressure promote" "$NI_BIN" pressure promote P-001 --dir "$pressure_ws"
require_output "promoted P-001"
run_cmd "namba-intent pressure retire" "$NI_BIN" pressure retire P-001 --dir "$pressure_ws"
require_output "retired P-001"

amend_ws="$(make_locked_workspace "amend")"
run_cmd "namba-intent amend create" "$NI_BIN" amend create --dir "$amend_ws" --title "Clarify smoke fixture"
require_output "created amendment AMEND-001"
complete_amendment "$amend_ws/.ni/amendments/AMEND-001.json"
run_cmd "namba-intent amend list" "$NI_BIN" amend list --dir "$amend_ws"
require_output "AMEND-001"
run_cmd "namba-intent amend show" "$NI_BIN" amend show AMEND-001 --dir "$amend_ws"
require_output '"status": "draft"'
run_cmd "namba-intent amend apply" "$NI_BIN" amend apply AMEND-001 --dir "$amend_ws"
require_output "applied amendment AMEND-001"

run_cmd "namba-intent relock" "$NI_BIN" relock --dir "$amend_ws"
require_output "relocked plan"

run_cmd "namba-intent diff" "$NI_BIN" diff --base "$ROOT/internal/core/collab/testdata/base.json" --head "$ROOT/internal/core/collab/testdata/non_conflicting_parallel_head.json"
require_output "contract diff"

run_cmd "namba-intent conflicts" "$NI_BIN" conflicts --base "$ROOT/internal/core/collab/testdata/base.json" --head "$ROOT/internal/core/collab/testdata/non_conflicting_parallel_head.json"
require_output "no collaboration conflicts"

graph_ws="$(make_ready_workspace "graph")"
run_cmd "namba-intent graph" "$NI_BIN" graph --dir "$graph_ws"
require_output "work graph proposal"

harness_ws="$(make_locked_workspace "harness")"
run_cmd "namba-intent harness plan" "$NI_BIN" harness plan --dir "$harness_ws"
require_output "generated harness proposal"
run_cmd "namba-intent harness candidates" "$NI_BIN" harness candidates --dir "$harness_ws"
require_output "no harness candidates"

write_harness_pressure "$harness_ws"
run_cmd "namba-intent harness propose" "$NI_BIN" harness propose --dir "$harness_ws" --from-pressure P-001
require_output "proposed harness candidate HC-001"
printf 'validated by smoke test\n' >"$harness_ws/harness-evidence.txt"
run_cmd "namba-intent harness validate" "$NI_BIN" harness validate HC-001 --dir "$harness_ws" --evidence harness-evidence.txt
require_output "validated harness candidate HC-001"
run_cmd "namba-intent harness accept" "$NI_BIN" harness accept HC-001 --dir "$harness_ws"
require_output "accepted harness candidate HC-001"
run_cmd "namba-intent harness retire" "$NI_BIN" harness retire HC-001 --dir "$harness_ws"
require_output "retired harness candidate HC-001"

echo "smoke checks passed"
