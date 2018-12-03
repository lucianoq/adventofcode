# !/usr/bin/python3

import sys


def count_for(str):
    two, three = False, False
    count = {}
    for c in l:
        if c not in count:
            count[c] = 1
        else:
            count[c] += 1
    for k, v in count.items():
        if v == 3:
            three = True
        if v == 2:
            two = True
        if two and three:
            return True, True
    return two, three


count_two, count_three = 0, 0
for l in sys.stdin:
    if l.strip() != '':
        two, three = count_for(l)
        if two:
            count_two += 1
        if three:
            count_three += 1

print(count_two * count_three)
