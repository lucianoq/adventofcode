# !/usr/bin/python3

import sys


def levenshtein_one(str1, str2):
    if len(str1) != len(str2):
        return False

    diff = 0
    for i in range(len(str1)):
        if str1[i] != str2[i]:
            diff += 1

    return diff == 1


def get_sim(str1, str2):
    out = ""
    for i in range(len(str1)):
        if str1[i] == str2[i]:
            out += str1[i]
    return out


lines = []
for l in sys.stdin:
    if l.strip() != '':
        lines.append(l)

for i in range(len(lines)):
    for j in range(i + 1, len(lines)):
        if levenshtein_one(lines[i], lines[j]):
            print(get_sim(lines[i], lines[j]))
            exit(0)

exit(-1)
