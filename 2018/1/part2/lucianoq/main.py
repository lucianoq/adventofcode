# !/usr/bin/python3

import sys

lines = []
for l in sys.stdin:
    l = l.strip()
    if l != '':
        lines.append(l)

shown = {}
tot = 0
shown[0] = True
while True:
    for l in lines:
        tot += int(l)

        if tot in shown:
            print(tot)
            exit(0)

        shown[tot] = True
