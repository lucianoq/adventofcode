# !/usr/bin/python3

import sys

d = {"(": 1, ")": -1}

acc, idx = 0, 0
for l in sys.stdin:
    for c in l:
        idx += 1
        acc += d[c]
        if acc == -1:
            print(idx)
            exit(0)
