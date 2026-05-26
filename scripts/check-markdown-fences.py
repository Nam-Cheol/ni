#!/usr/bin/env python3
from pathlib import Path
import sys

root = Path(__file__).resolve().parents[1]
failed = False
for path in sorted(root.rglob('*.md')):
    if '.git' in path.parts:
        continue
    ticks = 0
    for line in path.read_text(encoding='utf-8').splitlines():
        if line.strip().startswith('```'):
            ticks += 1
    if ticks % 2 != 0:
        print(f'Unbalanced markdown fence: {path.relative_to(root)}')
        failed = True
if failed:
    sys.exit(1)
print('Markdown fence check passed')
