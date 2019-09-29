# !/usr/bin/python3

import sys


def contains23(str):
    c2, c3 = False, False
    count = {}
    for ch in str:
        if ch not in count:
            count[ch] = 1
        else:
            count[ch] += 1
    for k, v in count.items():
        if v == 3:
            c3 = True
        if v == 2:
            c2 = True
        if c2 and c3:
            return True, True
    return c2, c3


num2, num3 = 0, 0
for l in sys.stdin:
    if l.strip() != '':
        contains2, contains3 = contains23(l)
        if contains2:
            num2 += 1
        if contains3:
            num3 += 1

print(num2 * num3)
