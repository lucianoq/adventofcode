from itertools import combinations
from sys import stdin

ls = [int(x) for x in stdin.readlines()]

for x in combinations(ls, 2):
    if x[0] + x[1] == 2020:
        print(x[0] * x[1])
        exit()
