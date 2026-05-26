#!/usr/bin/env bash
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
cd "$ROOT"

python3 scripts/check-json.py
python3 scripts/check-markdown-fences.py
python3 scripts/check-skill-metadata.py
python3 scripts/check-prompt-budget.py
python3 scripts/check-core-boundary.py --self-test

if [[ -f go.mod ]]; then
  gofmt -w .
  go test ./...
fi

bash scripts/smoke.sh

echo "quality checks passed"
