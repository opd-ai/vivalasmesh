import sys

with open('PLAN.md', 'r') as f:
    lines = f.readlines()

# Find the line index after which to insert
insert_after = None
for i, line in enumerate(lines):
    if line.strip().endswith('release candidate.'):
        insert_after = i
        break

if insert_after is None:
    print('Could not find insertion point')
    sys.exit(1)

# Prepare new lines to insert
new_lines = [
    '\n',
    '## Progress Tracking\n',
    '\n',
    '- [ ] Milestone 1: Core Architecture & Environment Setup\n',
    '- [ ] Milestone 2: Simulation Foundation & Deterministic Engine\n',
    '- [ ] Milestone 3: Asset Pipeline & World Building\n',
    '- [ ] Milestone 4: Quality Assurance, Balancing & Optimization\n',
    '- [ ] Milestone 5: Production Readiness & Deployment\n',
    '\n'
]

# Insert after insert_after (i.e., at position insert_after+1)
lines = lines[:insert_after+1] + new_lines + lines[insert_after+1:]

with open('PLAN.md', 'w') as f:
    f.writelines(lines)
