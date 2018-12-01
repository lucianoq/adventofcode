# !/usr/bin/python3

import sys

lines = []
for l in sys.stdin:
    if l.strip() != '':
        lines.append(l.strip())

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
