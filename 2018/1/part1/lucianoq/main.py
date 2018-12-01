# !/usr/bin/python3

import sys

lines = []
for l in sys.stdin:
    if l.strip() != '':
        lines.append(l.strip())

tot = 0
for l in lines:
    tot += int(l)

print(tot)
