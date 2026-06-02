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
  require_text "$file" "Skills are UX; CLI is authority."
  if [[ "$skill" == "ni-run" ]]; then
    require_text "$file" "ni run"
  else
    require_text "$file" "ni status"
  fi
  if [[ "$skill" == "ni-start" || "$skill" == "ni-grill" || "$skill" == "ni-status-review" ]]; then
    require_text "$file" "Skills may help draft or explain proof-related planning text."
    require_text "$file" "Skills do not determine readiness"
    require_text "$file" 'do not replace `ni status`,'
  fi
  require_text "$file" "BLOCKED"
  require_text "$file" "Do not"
  require_text "$file" "LOCK-STALE"
  require_text "$file" "Skills may help draft amended planning text."
  require_text "$file" "Skills may help explain \`LOCK-STALE\`."
  require_text "$file" "Skills do not lock or relock."
  require_text "$file" "Skills do not update \`.ni/plan.lock.json\`."
  require_text "$file" "review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000"
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

  for skill in ni-start ni-grill ni-status-review ni-end ni-run; do
    check_skill "$pack" "$skill"
  done

  require_text "$pack/README.md" "Skills are UX"
  require_text "$pack/README.md" "Status: Experimental."
  require_text "$pack/README.md" "Not verified: global host install"
  require_text "$pack/README.md" "Boundary: Skills are UX; CLI is authority."
  require_text "$pack/README.md" "CLI remains the authority"
  require_text "$pack/README.md" "Copy This Folder"
  require_text "$pack/README.md" "Verify The Pack"
  require_text "$pack/README.md" "What This Does Not Do"
  require_text "$pack/README.md" "Do not execute downstream work"
  require_text "$pack/README.md" "Does not replace"
  require_text "$pack/README.md" "Skills may help draft or explain proof-related planning text."
  require_text "$pack/README.md" "Skills may help draft amended planning text."
  require_text "$pack/README.md" 'Skills may help explain `LOCK-STALE`.'
  require_text "$pack/README.md" "Skills do not determine readiness."
  require_text "$pack/README.md" "Skills do not lock plans."
  require_text "$pack/README.md" "Skills do not lock or relock."
  require_text "$pack/README.md" 'Skills do not replace `ni status`, `ni end`, or `ni run`.'
  require_text "$pack/README.md" 'Skills do not update `.ni/plan.lock.json`.'
  require_text "$pack/README.md" "review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000"
  require_text "$pack/README.md" "unzip -l dist/"
  require_text "$pack/README.ko.md" "Skills are UX"
  require_text "$pack/README.ko.md" "Status: Experimental."
  require_text "$pack/README.ko.md" "Not verified: global host install"
  require_text "$pack/README.ko.md" "Boundary: Skills are UX; CLI is authority."
  require_text "$pack/README.ko.md" "CLI is authority"
  require_text "$pack/README.ko.md" "Copy This Folder"
  require_text "$pack/README.ko.md" "Verify The Pack"
  require_text "$pack/README.ko.md" "What This Does Not Do"
  require_text "$pack/README.ko.md" "replace하지 않는다"
  require_text "$pack/README.ko.md" "Skills may help draft or explain proof-related planning text."
  require_text "$pack/README.ko.md" "Skills may help draft amended planning text."
  require_text "$pack/README.ko.md" 'Skills may help explain `LOCK-STALE`.'
  require_text "$pack/README.ko.md" "Skills do not determine readiness."
  require_text "$pack/README.ko.md" "Skills do not lock plans."
  require_text "$pack/README.ko.md" "Skills do not lock or relock."
  require_text "$pack/README.ko.md" 'Skills do not replace `ni status`, `ni end`, or `ni run`.'
  require_text "$pack/README.ko.md" 'Skills do not update `.ni/plan.lock.json`.'
  require_text "$pack/README.ko.md" "review changed intent -> ni status --proof --next-questions -> ni end -> ni run --max-chars 4000"
  require_text "$pack/README.ko.md" "unzip -l dist/"
  require_text "$package_script" "$pack"
  require_text "$package_script" "README.md"
  require_text "$package_script" "README.ko.md"
  require_text "$package_script" "ni-start/SKILL.md"
  require_text "$package_script" "ni-grill/SKILL.md"
  require_text "$package_script" "ni-status-review/SKILL.md"
  require_text "$package_script" "ni-end/SKILL.md"
  require_text "$package_script" "ni-run/SKILL.md"
  require_text "$package_script" "zip -qr"
}

check_pack "Claude" "packages/claude-skills" "scripts/package-claude-skills.sh"
check_pack "Codex" "packages/codex-skills" "scripts/package-codex-skills.sh"

echo "checking repo-local Codex-style skills"
for skill in ni-start ni-grill ni-end ni-run; do
  check_skill ".agents/skills" "$skill"
done

require_file "scripts/install-claude-skills.sh"
require_text "scripts/install-claude-skills.sh" "--dry-run"
require_text "scripts/install-claude-skills.sh" "--target"
require_text "scripts/install-claude-skills.sh" "This script does not assume a global Claude skill path"
require_text "README.md" "| Model workspaces | Experimental |"
require_text "README.md" "Host-level/global install remains unverified unless documented"
require_text "README.ko.md" "| Model workspaces | Experimental |"
require_text "README.ko.md" "Host-level/global install은 documented되기 전까지 unverified"
require_file "docs/99_MODEL_WORKSPACE_STATUS.md"
require_file "docs/99_MODEL_WORKSPACE_STATUS.ko.md"
require_file "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md"
require_file "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md"
require_text "docs/99_MODEL_WORKSPACE_STATUS.md" "Model workspace packs are **Experimental**"
require_text "docs/99_MODEL_WORKSPACE_STATUS.md" "Global Claude install | not_verified"
require_text "docs/99_MODEL_WORKSPACE_STATUS.md" "Global Codex install | not_verified"
require_text "docs/99_MODEL_WORKSPACE_STATUS.md" "Provider runtime behavior | not_verified"
require_text "docs/99_MODEL_WORKSPACE_STATUS.md" "Skills are UX; CLI is authority."
require_text "docs/99_MODEL_WORKSPACE_STATUS.ko.md" "Model workspace packs는 broad product path로 **Experimental**"
require_text "docs/99_MODEL_WORKSPACE_STATUS.ko.md" "Global Claude install | not_verified"
require_text "docs/99_MODEL_WORKSPACE_STATUS.ko.md" "Global Codex install | not_verified"
require_text "docs/99_MODEL_WORKSPACE_STATUS.ko.md" "Provider runtime behavior | not_verified"
require_text "docs/99_MODEL_WORKSPACE_STATUS.ko.md" "Skills are UX; CLI is authority."
require_text "docs/83_CONVERSATION_PROOF_CAPTURE.md" "Proof Capture Lifecycle"
require_text "docs/83_CONVERSATION_PROOF_CAPTURE.md" "Conversation proof must not claim implementation correctness"
require_text "docs/83_CONVERSATION_PROOF_CAPTURE.ko.md" "Proof capture lifecycle"
require_text "docs/83_CONVERSATION_PROOF_CAPTURE.ko.md" "Conversation proof는 implementation correctness"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md" "## Proof capture lifecycle"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md" "Skills may help draft or explain proof-related planning text"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md" "Homebrew availability | No"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.md" "Broad model workspace availability | No"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md" "## Proof capture lifecycle"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md" "Skills may help draft or explain proof-related planning text"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md" "Homebrew availability | No"
require_text "docs/101_CONVERSATION_PROOF_CAPTURE_RELIABILITY.ko.md" "Broad model workspace availability | No"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.md" "Overall model workspace pack status: **Experimental**"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.md" "Global install claim"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.md" "not_verified"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.md" "What This Does Not Do"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md" "전체 model workspace pack status는 product path로는 **Experimental**"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md" "Global install claim"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md" "not_verified"
require_text "docs/75_MODEL_PACK_INSTALL_VERIFICATION.ko.md" "What This Does Not Do"

for doc in README.md README.ko.md docs/53_DISTRIBUTION_STRATEGY.md docs/53_DISTRIBUTION_STRATEGY.ko.md packages/claude-skills/README.md packages/claude-skills/README.ko.md packages/codex-skills/README.md packages/codex-skills/README.ko.md; do
  require_no_text "$doc" "| Model workspaces | Available |"
  require_no_text "$doc" "| Model workspace packs | Available |"
  require_no_text "$doc" "| Model workspace mode | Available |"
  require_no_text "$doc" "global host install is verified"
  require_no_text "$doc" "global host install은 verified"
  require_no_text "$doc" "global Codex install is verified"
  require_no_text "$doc" "global Claude install is verified"
done

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

pass "Claude skill pack repository evidence is verified through source files, zip packaging, and dry-run target install"
pass "Codex skill pack repository evidence is verified through repo-local source files and zip packaging; global install remains unverified"
echo "skill pack checks passed"
