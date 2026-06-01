#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
PACK_ROOT="$ROOT/packages/claude-skills"
TARGET=""
DRY_RUN=0
FORCE=0

usage() {
  cat <<'USAGE'
Usage:
  bash scripts/install-claude-skills.sh --dry-run --target /path/to/skills
  bash scripts/install-claude-skills.sh --target /path/to/skills [--force]

Copies NI Claude skill directories into a user-provided, verified target
directory. This script does not assume a global Claude skill path and does not
install or invoke the ni CLI.
USAGE
}

while [[ $# -gt 0 ]]; do
  case "$1" in
    --target)
      TARGET="${2:-}"
      shift 2
      ;;
    --target=*)
      TARGET="${1#--target=}"
      shift
      ;;
    --dry-run)
      DRY_RUN=1
      shift
      ;;
    --force)
      FORCE=1
      shift
      ;;
    -h|--help)
      usage
      exit 0
      ;;
    *)
      echo "unknown argument: $1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

if [[ -z "$TARGET" ]]; then
  echo "--target is required" >&2
  usage >&2
  exit 1
fi

if [[ ! -d "$PACK_ROOT" ]]; then
  echo "missing package root: packages/claude-skills" >&2
  exit 1
fi

if [[ ! -d "$TARGET" ]]; then
  echo "target directory does not exist: $TARGET" >&2
  echo "Create and verify the target directory first, then rerun this script." >&2
  exit 1
fi

skills=(ni-start ni-grill ni-end ni-run ni-status-review)

for skill in "${skills[@]}"; do
  src="$PACK_ROOT/$skill"
  dest="$TARGET/$skill"

  if [[ ! -f "$src/SKILL.md" ]]; then
    echo "missing skill file: packages/claude-skills/$skill/SKILL.md" >&2
    exit 1
  fi

  if [[ -e "$dest" && "$FORCE" -ne 1 ]]; then
    echo "refusing to overwrite existing target: $dest" >&2
    echo "Use --force after reviewing the target directory." >&2
    exit 1
  fi

  if [[ "$DRY_RUN" -eq 1 ]]; then
    if [[ -e "$dest" ]]; then
      echo "would replace $dest with $src"
    else
      echo "would copy $src to $dest"
    fi
  else
    rm -rf "$dest"
    cp -R "$src" "$dest"
    echo "installed $skill to $dest"
  fi
done

if [[ "$DRY_RUN" -eq 1 ]]; then
  echo "dry run complete; no files changed"
else
  echo "installed NI Claude skills to $TARGET"
fi
