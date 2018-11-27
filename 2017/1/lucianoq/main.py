#!/usr/bin/python3

import sys

for line in sys.stdin:
    line = line.strip()
    if line != '':
        tot = 0
        line += line[-1]
        for i in range(len(line) - 1):
            if line[i] == line[i + 1]:
                tot += int(line[i])
        print(tot)
