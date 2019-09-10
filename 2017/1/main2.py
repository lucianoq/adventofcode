import sys

for line in sys.stdin:
    line = line.strip()
    if line != '':
        l = len(line)
        tot = 0
        for i in range(l):
            if line[i] == line[(i + l // 2) % l]:
                tot += int(line[i])
        print(tot)
