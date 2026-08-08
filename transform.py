import sys
import re

def main():
    with open('PLAN.md', 'r') as f:
        lines = f.readlines()
    
    out_lines = []
    i = 0
    while i < len(lines):
        line = lines[i].rstrip('\n')
        # Determine transformation
        new_line = line
        # Title line
        if line == 'Viva Las Mesh: Master Production Release Plan and Technical Roadmap':
            new_line = '# ' + line
        # Milestone headings
        elif re.match(r'^Milestone \d+: .+', line):
            new_line = '## ' + line
        # Specification headings
        elif re.match(r'.+Specifications$', line) and line not in ['Execution Checklist: Core Architecture & Setup', 'Execution Checklist: Simulation Foundation & Deterministic Engine', 'Execution Checklist: Asset Pipeline & World Building', 'Execution Checklist: Quality Assurance & Optimization', 'Execution Checklist: Production Readiness & Deployment']:
            # Avoid matching the execution checklist lines (they end with 'Setup' not 'Specifications')
            new_line = '### ' + line
        # Execution Checklist headings
        elif line.startswith('Execution Checklist:'):
            new_line = '### ' + line
        # Deployment Pipeline & Release Artifacts (under Milestone 5)
        elif line == 'Deployment Pipeline & Release Artifacts':
            new_line = '### ' + line
        # Technical Summary & Critical Path Dependencies
        elif line == 'Technical Summary & Critical Path Dependencies':
            new_line = '## ' + line
        # else keep as is
        out_lines.append(new_line)
        i += 1
    
    # Write back
    with open('PLAN.md', 'w') as f:
        for line in out_lines:
            f.write(line + '\n')

if __name__ == '__main__':
    main()
