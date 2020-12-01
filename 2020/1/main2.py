from itertools import combinations
from sys import stdin

ls = [int(x) for x in stdin.readlines()]

for x in combinations(ls, 3):
    if x[0] + x[1] + x[2] == 2020:
        print(x[0] * x[1] * x[2])
        exit()
