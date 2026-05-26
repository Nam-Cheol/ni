#!/usr/bin/env python3
from pathlib import Path
import sys

root = Path(__file__).resolve().parents[1]
max_chars = 4000
failed = False
for path in sorted((root / '.ni' / 'generated').glob('*.prompt.txt')) if (root / '.ni' / 'generated').exists() else []:
    n = len(path.read_text(encoding='utf-8'))
    if n > max_chars:
        print(f'Prompt too long: {path.relative_to(root)} has {n} chars > {max_chars}')
        failed = True
if failed:
    sys.exit(1)
print('Generated prompt budget check passed')
