from sys import stdin

mass: [int] = [int(x) // 3 - 2 for x in stdin.readlines()]

total_fuel: int = 0
for m in mass:
    fuel, x = 0, m
    while x > 0:
        fuel += x
        x = x // 3 - 2
    total_fuel += fuel

print(total_fuel)
