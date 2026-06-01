#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PACK_ROOT="$ROOT/packages/codex-skills"
DIST_DIR="$ROOT/dist"
OUT="$DIST_DIR/ni-codex-skills.zip"
STAGING_PARENT="${TMPDIR:-/tmp}/ni-codex-skills-package.$$"
STAGING="$STAGING_PARENT/ni-codex-skills"

if [[ ! -d "$PACK_ROOT" ]]; then
  echo "missing package root: packages/codex-skills" >&2
  exit 1
fi

for required in \
  README.md \
  README.ko.md \
  ni-start/SKILL.md \
  ni-grill/SKILL.md \
  ni-end/SKILL.md \
  ni-run/SKILL.md \
  ni-status-review/SKILL.md
do
  if [[ ! -f "$PACK_ROOT/$required" ]]; then
    echo "missing Codex skill pack file: packages/codex-skills/$required" >&2
    exit 1
  fi
done

command -v zip >/dev/null 2>&1 || {
  echo "zip command not found" >&2
  exit 1
}

rm -rf "$STAGING_PARENT"
mkdir -p "$STAGING" "$DIST_DIR"
cp -R "$PACK_ROOT/README.md" "$STAGING/README.md"
cp -R "$PACK_ROOT/README.ko.md" "$STAGING/README.ko.md"
cp -R "$PACK_ROOT/ni-start" "$STAGING/ni-start"
cp -R "$PACK_ROOT/ni-grill" "$STAGING/ni-grill"
cp -R "$PACK_ROOT/ni-end" "$STAGING/ni-end"
cp -R "$PACK_ROOT/ni-run" "$STAGING/ni-run"
cp -R "$PACK_ROOT/ni-status-review" "$STAGING/ni-status-review"

rm -f "$OUT"
(
  cd "$STAGING_PARENT"
  zip -qr "$OUT" ni-codex-skills
)
rm -rf "$STAGING_PARENT"

echo "created $OUT"
