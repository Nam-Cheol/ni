#!/usr/bin/env python3
from pathlib import Path
import sys

root = Path(__file__).resolve().parents[1]
skills_root = root / '.agents' / 'skills'
failed = False

required_phrases = {
    'ni-start': [
        'docs/plan/**',
        '.ni/contract.json',
        '.ni/session.json',
        'Do not create `.ni/plan.lock.json`',
        'Do not edit files outside `docs/plan/**`, `.ni/contract.json`, and',
    ],
    'ni-end': [
        '`ni status` and `ni end` are the authority',
        'ni status --dir .',
        'ni end --dir .',
        'Do not write `.ni/plan.lock.json` manually',
    ],
    'ni-run': [
        '`ni run` is a prompt compiler in v0',
        'Do not reimplement prompt compilation in the skill',
        'ni run --dir . --target <target> --max-chars 4000',
        'State clearly that `ni` compiled a prompt only and did not execute',
        'Do not execute Codex or shell commands as part of v0 `ni run`',
    ],
}

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
        for phrase in required_phrases.get(skill.name, []):
            if phrase not in text:
                print(f'Missing required skill phrase in {md.relative_to(root)}: {phrase}')
                failed = True
if failed:
    sys.exit(1)
print('Skill metadata check passed')
