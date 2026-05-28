#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

failed=0

fail() {
  echo "FAIL: $*" >&2
  failed=1
}

pass() {
  echo "ok: $*"
}

require_file() {
  local file="$1"
  if [[ ! -f "$file" ]]; then
    fail "missing file: $file"
  fi
}

require_text() {
  local file="$1"
  local pattern="$2"
  if [[ ! -f "$file" ]]; then
    fail "missing file for text check: $file"
    return
  fi
  if ! grep -Fq -- "$pattern" "$file"; then
    fail "missing required text in $file: $pattern"
  fi
}

require_no_text() {
  local file="$1"
  local pattern="$2"
  if [[ -f "$file" ]] && grep -Fq -- "$pattern" "$file"; then
    fail "forbidden text in $file: $pattern"
  fi
}

check_skill() {
  local pack="$1"
  local skill="$2"
  local file="$pack/$skill/SKILL.md"

  require_file "$file"
  require_text "$file" "name: $skill"
  require_text "$file" "description:"
  require_text "$file" "Authority"
  if [[ "$skill" == "ni-run" ]]; then
    require_text "$file" "ni run"
  else
    require_text "$file" "ni status"
  fi
  require_text "$file" "BLOCKED"
  require_text "$file" "Do not"
  require_no_text "$file" "codex exec --"
  require_no_text "$file" "claude "
  require_no_text "$file" "anthropic "
  require_no_text "$file" "openai "
}

check_pack() {
  local label="$1"
  local pack="$2"
  local package_script="$3"

  echo "checking $label skill pack"
  require_file "$pack/README.md"
  require_file "$pack/README.ko.md"
  require_file "$package_script"

  for skill in ni-start ni-status-review ni-end ni-run; do
    check_skill "$pack" "$skill"
  done

  require_text "$pack/README.md" "Skills are UX"
  require_text "$pack/README.md" "CLI remains the authority"
  require_text "$pack/README.md" "Do not execute downstream work"
  require_text "$pack/README.ko.md" "Skills are UX"
  require_text "$pack/README.ko.md" "CLI is authority"
  require_text "$package_script" "$pack"
  require_text "$package_script" "README.md"
  require_text "$package_script" "README.ko.md"
  require_text "$package_script" "ni-start/SKILL.md"
  require_text "$package_script" "ni-status-review/SKILL.md"
  require_text "$package_script" "ni-end/SKILL.md"
  require_text "$package_script" "ni-run/SKILL.md"
  require_text "$package_script" "zip -qr"
}

check_pack "Claude" "packages/claude-skills" "scripts/package-claude-skills.sh"
check_pack "Codex" "packages/codex-skills" "scripts/package-codex-skills.sh"

require_file "scripts/install-claude-skills.sh"
require_text "scripts/install-claude-skills.sh" "--dry-run"
require_text "scripts/install-claude-skills.sh" "--target"
require_text "scripts/install-claude-skills.sh" "This script does not assume a global Claude skill path"

dry_target="$(mktemp -d "${TMPDIR:-/tmp}/ni-skill-pack-check.XXXXXX")"
trap 'rm -rf "$dry_target"' EXIT
dry_output="$(bash scripts/install-claude-skills.sh --dry-run --target "$dry_target")"
if ! grep -Fq "dry run complete; no files changed" <<<"$dry_output"; then
  fail "Claude dry-run installer did not report a no-change dry run"
else
  pass "Claude dry-run installer reports no file changes"
fi

if [[ "$failed" -ne 0 ]]; then
  echo "skill pack checks failed" >&2
  exit 1
fi

pass "Claude skill pack source is Available through manual copy, zip packaging, and dry-run target install"
pass "Codex skill pack source is Available through repo-local/manual copy and zip packaging; global install remains unverified"
echo "skill pack checks passed"
