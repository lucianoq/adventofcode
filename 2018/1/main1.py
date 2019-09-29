# !/usr/bin/python3

import sys

tot = 0
for l in sys.stdin:
    if l.strip() != '':
        tot += int(l.strip())

print(tot)
