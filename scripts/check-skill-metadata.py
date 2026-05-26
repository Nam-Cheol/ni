#!/usr/bin/env python3
from pathlib import Path
import sys

root = Path(__file__).resolve().parents[1]
skills_root = root / '.agents' / 'skills'
failed = False
if skills_root.exists():
    for skill in sorted(skills_root.iterdir()):
        if not skill.is_dir():
            continue
        md = skill / 'SKILL.md'
        if not md.exists():
            print(f'Missing SKILL.md: {skill.relative_to(root)}')
            failed = True
            continue
        text = md.read_text(encoding='utf-8')
        if 'name:' not in text or 'description:' not in text:
            print(f'Missing name/description metadata: {md.relative_to(root)}')
            failed = True
if failed:
    sys.exit(1)
print('Skill metadata check passed')
