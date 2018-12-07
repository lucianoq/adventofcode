import sys

hz = list()
seen = dict()
tot = 0
seen[0] = True

for line in sys.stdin :
  hz.append(int(line))

while True :
  for f in hz:
    tot += f
    if tot in seen :
      print(tot)
      sys.exit(0)
    seen[tot] = True
