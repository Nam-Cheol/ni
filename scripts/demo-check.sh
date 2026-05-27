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
  go run ./cmd/ni status --dir ./examples/ambiguous-prompt-blocked/workspace --next-questions >"$1/ambiguous-next-questions.out"
' bash "$DEMO_TMP"
require_first_line "BLOCKED" "$DEMO_TMP/ambiguous-next-questions.out"
require_output "question R009" "$DEMO_TMP/ambiguous-next-questions.out"

run_demo "research protocol status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/research-protocol >"$1/research-status.out"
' bash "$DEMO_TMP"
require_doc_status "examples/research-protocol" "READY"
require_first_line "READY" "$DEMO_TMP/research-status.out"

run_demo "research protocol human-team prompt compiles if locked" \
  run_if_locked "examples/research-protocol" "human-team" "$DEMO_TMP/ni-research-human-team.prompt.md"

run_demo "conversation product status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/conversation-product >"$1/conversation-status.out"
' bash "$DEMO_TMP"
require_doc_status "examples/conversation-product" "READY"
require_first_line "READY" "$DEMO_TMP/conversation-status.out"

run_demo "conversation product human-team prompt compiles if locked" \
  run_if_locked "examples/conversation-product" "human-team" "$DEMO_TMP/ni-conversation-human-team.prompt.md"

run_demo "ni-start dogfood status matches docs" bash -c '
  go run ./cmd/ni status --dir examples/ni-start-dogfood/workspace >"$1/ni-start-dogfood-status.out"
' bash "$DEMO_TMP"
require_doc_status "examples/ni-start-dogfood" "READY_WITH_DEFERRALS"
require_first_line "READY_WITH_DEFERRALS" "$DEMO_TMP/ni-start-dogfood-status.out"

run_demo "ni-start dogfood human-team prompt compiles if locked" \
  run_if_locked "examples/ni-start-dogfood/workspace" "human-team" "$DEMO_TMP/ni-start-dogfood-human-team.prompt.md"

for export_target in hyper-run namba-ai ouroboros spec-kit; do
  run_demo "conversation product $export_target export stays seed-only if locked" \
    export_if_locked "examples/conversation-product" "$export_target" "$DEMO_TMP/conversation-product-export-$export_target"
done

echo "demo-check: public demos verified without downstream runtime execution"
