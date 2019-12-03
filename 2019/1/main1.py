from sys import stdin

print(sum([int(x) // 3 - 2 for x in stdin.readlines()]))
