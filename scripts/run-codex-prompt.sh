#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

if [[ -n "$(git status --porcelain=v1)" ]]; then
  echo "refusing to run: git tree is not clean" >&2
  git status --short >&2
  exit 1
fi

if ! command -v codex >/dev/null 2>&1; then
  echo "refusing to run: codex CLI was not found on PATH" >&2
  exit 1
fi

PROMPT=".ni/generated/goal.prompt.txt"

go run ./cmd/namba-intent run --dir "$ROOT" --out "$PROMPT"

exec codex exec --sandbox workspace-write --cd "$ROOT" - < "$PROMPT"
