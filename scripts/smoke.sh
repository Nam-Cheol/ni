#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
SMOKE_TMP="$(mktemp -d "${TMPDIR:-/tmp}/ni-smoke.XXXXXX")"
NI_BIN="$SMOKE_TMP/bin/ni"
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
    "purpose": "Exercise public ni commands.",
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
      "risks": [],
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
      "title": "Smoke tests stay inside ni-kernel boundaries.",
      "status": "accepted"
    }
  ],
  "risks": [],
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
  run_cmd "ni init ($name)" "$NI_BIN" init --dir "$dir" --profile prototype
  write_ready_contract "$dir"
  echo "$dir"
}

make_locked_workspace() {
  local name="$1"
  local dir
  dir="$(make_ready_workspace "$name")"
  run_cmd "ni end ($name)" "$NI_BIN" end --dir "$dir"
  echo "$dir"
}

cd "$ROOT"

echo "smoke: building ni" >&2
go build -o "$NI_BIN" ./cmd/ni

run_cmd "ni --help" "$NI_BIN" --help
require_output "ni is a project intent compiler"

run_cmd "ni version" "$NI_BIN" version
require_output "0.0.0-dev"

init_ws="$SMOKE_TMP/workspaces/init"
run_cmd "ni init" "$NI_BIN" init --dir "$init_ws" --profile concept
require_output "initialized ni planning workspace"

run_cmd "ni status" "$NI_BIN" status --dir "$init_ws"
require_output "BLOCKED"

ready_ws="$(make_ready_workspace "ready")"
run_cmd "ni end on ready fixture" "$NI_BIN" end --dir "$ready_ws"
require_output "locked plan"

locked_ws="$(make_locked_workspace "locked")"
for prompt_target in generic codex human-team; do
  run_cmd "ni run --target $prompt_target" "$NI_BIN" run --dir "$locked_ws" --target "$prompt_target"
  require_output "Target: $prompt_target"
done

run_cmd "ni targets" "$NI_BIN" targets
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
  run_cmd "ni export --target $export_target" "$NI_BIN" export --dir "$locked_ws" --target "$export_target" --out "$SMOKE_TMP/exports/$export_target"
  require_output "exported $export_target seed package"
done

feedback_ws="$SMOKE_TMP/workspaces/feedback"
run_cmd "ni init (feedback)" "$NI_BIN" init --dir "$feedback_ws"
run_cmd "ni feedback add" "$NI_BIN" feedback add --dir "$feedback_ws" --file "$ROOT/testdata/feedback/codex.json"
require_output "recorded feedback from codex"
run_cmd "ni feedback list" "$NI_BIN" feedback list --dir "$feedback_ws"
require_output "codex"

pressure_ws="$SMOKE_TMP/workspaces/pressure"
run_cmd "ni init (pressure)" "$NI_BIN" init --dir "$pressure_ws"
run_cmd "ni feedback add (pressure fixture)" "$NI_BIN" feedback add --dir "$pressure_ws" --file "$ROOT/testdata/feedback/codex.json"
run_cmd "ni pressure status" "$NI_BIN" pressure status --dir "$pressure_ws"
require_output "P-001"
run_cmd "ni pressure promote" "$NI_BIN" pressure promote P-001 --dir "$pressure_ws"
require_output "promoted P-001"
run_cmd "ni pressure retire" "$NI_BIN" pressure retire P-001 --dir "$pressure_ws"
require_output "retired P-001"

amend_ws="$(make_locked_workspace "amend")"
run_cmd "ni amend create" "$NI_BIN" amend create --dir "$amend_ws" --title "Clarify smoke fixture"
require_output "created amendment AMEND-001"
complete_amendment "$amend_ws/.ni/amendments/AMEND-001.json"
run_cmd "ni amend list" "$NI_BIN" amend list --dir "$amend_ws"
require_output "AMEND-001"
run_cmd "ni amend show" "$NI_BIN" amend show AMEND-001 --dir "$amend_ws"
require_output '"status": "draft"'
run_cmd "ni amend apply" "$NI_BIN" amend apply AMEND-001 --dir "$amend_ws"
require_output "applied amendment AMEND-001"

run_cmd "ni relock" "$NI_BIN" relock --dir "$amend_ws"
require_output "relocked plan"

run_cmd "ni diff" "$NI_BIN" diff --base "$ROOT/internal/core/collab/testdata/base.json" --head "$ROOT/internal/core/collab/testdata/non_conflicting_parallel_head.json"
require_output "contract diff"

run_cmd "ni conflicts" "$NI_BIN" conflicts --base "$ROOT/internal/core/collab/testdata/base.json" --head "$ROOT/internal/core/collab/testdata/non_conflicting_parallel_head.json"
require_output "no collaboration conflicts"

graph_ws="$(make_ready_workspace "graph")"
run_cmd "ni graph" "$NI_BIN" graph --dir "$graph_ws"
require_output "work graph proposal"

harness_ws="$(make_locked_workspace "harness")"
run_cmd "ni harness plan" "$NI_BIN" harness plan --dir "$harness_ws"
require_output "generated harness proposal"
run_cmd "ni harness candidates" "$NI_BIN" harness candidates --dir "$harness_ws"
require_output "no harness candidates"

write_harness_pressure "$harness_ws"
run_cmd "ni harness propose" "$NI_BIN" harness propose --dir "$harness_ws" --from-pressure P-001
require_output "proposed harness candidate HC-001"
printf 'validated by smoke test\n' >"$harness_ws/harness-evidence.txt"
run_cmd "ni harness validate" "$NI_BIN" harness validate HC-001 --dir "$harness_ws" --evidence harness-evidence.txt
require_output "validated harness candidate HC-001"
run_cmd "ni harness accept" "$NI_BIN" harness accept HC-001 --dir "$harness_ws"
require_output "accepted harness candidate HC-001"
run_cmd "ni harness retire" "$NI_BIN" harness retire HC-001 --dir "$harness_ws"
require_output "retired harness candidate HC-001"

echo "smoke checks passed"
