#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
INSTALL_TMP="$(mktemp -d "${TMPDIR:-/tmp}/ni-install-check.XXXXXX")"
LAST_STDOUT="$INSTALL_TMP/stdout.log"
LAST_STDERR="$INSTALL_TMP/stderr.log"

trap 'rm -rf "$INSTALL_TMP"' EXIT

cd "$ROOT"

run_cmd() {
  local label="$1"
  shift
  echo "install-check: $label" >&2
  : >"$LAST_STDOUT"
  : >"$LAST_STDERR"
  if ! "$@" >"$LAST_STDOUT" 2>"$LAST_STDERR"; then
    echo "install-check failed: $label" >&2
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
    echo "install-check failed: expected stdout to contain: $expected" >&2
    echo "--- stdout ---" >&2
    sed -n '1,120p' "$LAST_STDOUT" >&2
    return 1
  fi
}

require_nonempty_output() {
  if [[ ! -s "$LAST_STDOUT" ]]; then
    echo "install-check failed: expected non-empty stdout" >&2
    return 1
  fi
}

run_cmd "go run ./cmd/ni --help" go run ./cmd/ni --help
require_output "ni is a project intent compiler"

run_cmd "go run ./cmd/ni version" go run ./cmd/ni version
require_output "0.0.0-dev"

run_cmd "make build" make build

run_cmd "./bin/ni --help" ./bin/ni --help
require_output "ni is a project intent compiler"

run_cmd "./bin/ni version" ./bin/ni version
require_nonempty_output

run_cmd "make install-local with temporary BINDIR" make install-local BINDIR="$INSTALL_TMP/bin"

run_cmd "temporary installed ni --help" "$INSTALL_TMP/bin/ni" --help
require_output "ni is a project intent compiler"

run_cmd "temporary installed ni version" "$INSTALL_TMP/bin/ni" version
require_nonempty_output

run_cmd "fresh shell resolves temporary ni --help and version" env \
  PATH="$INSTALL_TMP/bin:$PATH" \
  sh -c 'command -v ni && ni --help && ni version'
require_output "ni is a project intent compiler"
require_nonempty_output

run_cmd "install.sh global install and uninstall behavior" bash scripts/test-install-sh.sh

run_cmd "Windows installer static safety check" python3 scripts/check-install-ps1.py

echo "install-check: source, build, temporary global command, installer, and uninstall checks passed"
