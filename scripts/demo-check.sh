#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
DEMO_TMP="$(mktemp -d "${TMPDIR:-/tmp}/ni-demo-check.XXXXXX")"

trap 'rm -rf "$DEMO_TMP"' EXIT

cd "$ROOT"

run_demo() {
  local label="$1"
  shift
  echo "demo-check: $label" >&2
  "$@"
}

require_output() {
  local expected="$1"
  local file="$2"
  if ! grep -Fq "$expected" "$file"; then
    echo "demo-check failed: expected output to contain: $expected" >&2
    sed -n '1,120p' "$file" >&2
    return 1
  fi
}

require_first_line() {
  local expected="$1"
  local file="$2"
  local actual
  actual="$(sed -n '1p' "$file")"
  if [[ "$actual" != "$expected" ]]; then
    echo "demo-check failed: expected first line '$expected', got '$actual'" >&2
    sed -n '1,120p' "$file" >&2
    return 1
  fi
}

require_doc_status() {
  local example_dir="$1"
  local expected="$2"
  require_output "Expected status: \`$expected\`." "$example_dir/README.md"
}

require_file() {
  local path="$1"
  if [[ ! -f "$path" ]]; then
    echo "demo-check failed: required file is missing: $path" >&2
    return 1
  fi
}

check_benchmark_report_docs() {
  require_file examples/benchmark-report/README.md
  require_file examples/benchmark-report/README.ko.md
  require_file examples/benchmark-report/sample-report.md
  require_file examples/benchmark-report/cases/internal-dashboard/README.md
  require_file examples/benchmark-report/cases/internal-dashboard/01-original-request.md
  require_file examples/benchmark-report/cases/internal-dashboard/02-direct-to-agent-risk.md
  require_file examples/benchmark-report/cases/internal-dashboard/03-ni-path.md
  require_file examples/benchmark-report/cases/internal-dashboard/04-measurement-table.md
  require_file examples/benchmark-report/cases/internal-dashboard/05-not-measured.md
  require_file examples/benchmark-report/cases/internal-dashboard/06-ni-status-proof.md
  require_file examples/benchmark-report/cases/internal-dashboard/07-ni-next-questions.md
  require_file examples/benchmark-report/cases/internal-dashboard/workspace/.ni/contract.json
  require_file docs/43_BENCHMARK_PROTOCOL.md
}

check_no_terminal_assisted_docs() {
  require_file examples/no-terminal-assisted/README.md
  require_file examples/no-terminal-assisted/README.ko.md
  require_file examples/no-terminal-assisted/docs/plan/00_project_brief.md
  require_file examples/no-terminal-assisted/.ni/contract.json
  require_output "Expected \`ni status\`: not claimed" "examples/no-terminal-assisted/README.md"
  require_output "docs-only example" "examples/no-terminal-assisted/README.md"
  require_output "It does not run" "examples/no-terminal-assisted/README.md"
}

run_if_locked() {
  local example_dir="$1"
  local target="$2"
  local out_path="$3"

  if [[ ! -f "$example_dir/.ni/plan.lock.json" ]]; then
    echo "demo-check: skipping $example_dir run; example is not locked" >&2
    return 0
  fi

  go run ./cmd/ni run --dir "$example_dir" --target "$target" --out "$out_path"
  if [[ ! -s "$out_path" ]]; then
    echo "demo-check failed: compiled prompt is missing or empty: $out_path" >&2
    return 1
  fi
}

export_if_locked() {
  local example_dir="$1"
  local target="$2"
  local out_dir="$3"

  if [[ ! -f "$example_dir/.ni/plan.lock.json" ]]; then
    echo "demo-check: skipping $example_dir export; example is not locked" >&2
    return 0
  fi

  go run ./cmd/ni export --dir "$example_dir" --target "$target" --out "$out_dir"
  python3 scripts/check-target-conformance.py --target "$target" --dir "$out_dir"
}

run_demo "ambiguous prompt demo remains blocked" bash -c '
  go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace >"$1/ambiguous-status.out"
' bash "$DEMO_TMP"
require_first_line "BLOCKED" "$DEMO_TMP/ambiguous-status.out"
require_output "blocker R009" "$DEMO_TMP/ambiguous-status.out"

run_demo "ambiguous prompt next questions render" bash -c '
  go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --proof --next-questions >"$1/ambiguous-next-questions.out"
' bash "$DEMO_TMP"
require_output "NI Intent Readiness: BLOCKED" "$DEMO_TMP/ambiguous-next-questions.out"
require_output "Open blockers:" "$DEMO_TMP/ambiguous-next-questions.out"
require_output "OQ-001: OQ-001 is blocking readiness" "$DEMO_TMP/ambiguous-next-questions.out"

run_demo "research protocol status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/research-protocol >"$1/research-status.out"
  go run ./cmd/ni status --dir examples/research-protocol --proof --next-questions >"$1/research-proof.out"
' bash "$DEMO_TMP"
require_doc_status "examples/research-protocol" "READY"
require_first_line "READY" "$DEMO_TMP/research-status.out"
require_output "NI Intent Readiness: READY" "$DEMO_TMP/research-proof.out"

run_demo "research protocol human-team prompt compiles if locked" \
  run_if_locked "examples/research-protocol" "human-team" "$DEMO_TMP/ni-research-human-team.prompt.md"

run_demo "conversation product status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/conversation-product >"$1/conversation-status.out"
  go run ./cmd/ni status --dir examples/conversation-product --proof --next-questions >"$1/conversation-proof.out"
' bash "$DEMO_TMP"
require_doc_status "examples/conversation-product" "READY"
require_first_line "READY" "$DEMO_TMP/conversation-status.out"
require_output "NI Intent Readiness: READY" "$DEMO_TMP/conversation-proof.out"

run_demo "conversation product human-team prompt compiles if locked" \
  run_if_locked "examples/conversation-product" "human-team" "$DEMO_TMP/ni-conversation-human-team.prompt.md"

run_demo "ni-start dogfood status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace >"$1/ni-start-dogfood-status.out"
  go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace --proof --next-questions >"$1/ni-start-dogfood-proof.out"
' bash "$DEMO_TMP"
require_doc_status "examples/ni-start-dogfood" "READY_WITH_DEFERRALS"
require_first_line "READY_WITH_DEFERRALS" "$DEMO_TMP/ni-start-dogfood-status.out"
require_output "NI Intent Readiness: READY_WITH_DEFERRALS" "$DEMO_TMP/ni-start-dogfood-proof.out"
require_output "Handoff deferrals:" "$DEMO_TMP/ni-start-dogfood-proof.out"

run_demo "ni-start dogfood human-team prompt compiles if locked" \
  run_if_locked "examples/ni-start-dogfood/workspace" "human-team" "$DEMO_TMP/ni-start-dogfood-human-team.prompt.md"

run_demo "conversation authoring status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/conversation-authoring >"$1/conversation-authoring-status.out"
  go run ./cmd/ni status --dir examples/conversation-authoring --proof --next-questions >"$1/conversation-authoring-proof.out"
' bash "$DEMO_TMP"
require_doc_status "examples/conversation-authoring" "BLOCKED"
require_first_line "BLOCKED" "$DEMO_TMP/conversation-authoring-status.out"
require_output "blocker R012" "$DEMO_TMP/conversation-authoring-status.out"
require_output "NI Intent Readiness: BLOCKED" "$DEMO_TMP/conversation-authoring-proof.out"
require_output "Sync repairs:" "$DEMO_TMP/conversation-authoring-proof.out"

run_demo "conversation authoring human-team prompt compiles from existing lock" \
  run_if_locked "examples/conversation-authoring" "human-team" "$DEMO_TMP/conversation-authoring-human-team.prompt.md"

run_demo "namba-ai upgrade status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/namba-ai-upgrade >"$1/namba-ai-upgrade-status.out"
  go run ./cmd/ni status --dir examples/namba-ai-upgrade --proof --next-questions >"$1/namba-ai-upgrade-proof.out"
' bash "$DEMO_TMP"
require_doc_status "examples/namba-ai-upgrade" "BLOCKED"
require_first_line "BLOCKED" "$DEMO_TMP/namba-ai-upgrade-status.out"
require_output "blocker R012: CAP-001" "$DEMO_TMP/namba-ai-upgrade-status.out"
require_output "NI Intent Readiness: BLOCKED" "$DEMO_TMP/namba-ai-upgrade-proof.out"
require_output "Sync repairs:" "$DEMO_TMP/namba-ai-upgrade-proof.out"

run_demo "namba-ai upgrade codex prompt compiles from existing lock" \
  run_if_locked "examples/namba-ai-upgrade" "codex" "$DEMO_TMP/namba-ai-upgrade-codex.prompt.md"

run_demo "benchmark report internal dashboard remains blocked pre-runtime" bash -c '
  go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace >"$1/internal-dashboard-status.out"
  go run ./cmd/ni status --dir examples/benchmark-report/cases/internal-dashboard/workspace --proof --next-questions >"$1/internal-dashboard-proof.out"
' bash "$DEMO_TMP"
check_benchmark_report_docs
require_first_line "BLOCKED" "$DEMO_TMP/internal-dashboard-status.out"
require_output "blocker R009: OQ-001 is a blocker open question" "$DEMO_TMP/internal-dashboard-status.out"
require_output "NI Intent Readiness: BLOCKED" "$DEMO_TMP/internal-dashboard-proof.out"
require_output "Open blockers:" "$DEMO_TMP/internal-dashboard-proof.out"
require_output "OQ-001: OQ-001 is blocking readiness" "$DEMO_TMP/internal-dashboard-proof.out"
require_output "Expected \`ni status\`: not applicable" "examples/benchmark-report/README.md"
require_output "not_measured" "examples/benchmark-report/README.md"
require_output "not_measured" "examples/benchmark-report/README.ko.md"
require_output "not_measured" "examples/benchmark-report/sample-report.md"
require_output "not_measured" "examples/benchmark-report/cases/internal-dashboard/04-measurement-table.md"
require_output "No downstream agent was executed" "examples/benchmark-report/cases/internal-dashboard/05-not-measured.md"
require_output "NI Intent Readiness: BLOCKED" "examples/benchmark-report/cases/internal-dashboard/06-ni-status-proof.md"
require_output "must not execute downstream agents" "docs/43_BENCHMARK_PROTOCOL.md"
require_output "Target prompt boundedness" "docs/43_BENCHMARK_PROTOCOL.md"

run_demo "no-terminal assisted remains docs-only" check_no_terminal_assisted_docs

for export_target in hyper-run namba-ai ouroboros spec-kit; do
  run_demo "conversation product $export_target export stays seed-only if locked" \
    export_if_locked "examples/conversation-product" "$export_target" "$DEMO_TMP/conversation-product-export-$export_target"
done

echo "demo-check: public demos verified without downstream runtime execution"
