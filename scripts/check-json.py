#!/usr/bin/env python3
from pathlib import Path
import json
import sys

root = Path(__file__).resolve().parents[1]
failed = False
for path in sorted(root.rglob('*.json')):
    if any(part in {'.git', 'tmp', 'generated', 'runs'} for part in path.parts):
        continue
    try:
        json.loads(path.read_text(encoding='utf-8'))
    except Exception as exc:
        print(f'JSON invalid: {path.relative_to(root)}: {exc}')
        failed = True
if failed:
    sys.exit(1)
print('JSON check passed')
