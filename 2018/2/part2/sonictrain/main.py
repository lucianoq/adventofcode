import sys

ids = list()
for line in sys.stdin :
  ids.append(line.strip())

def leven(x, y):
  distance = 0
  for i in range(len(x)):
    if x[i] != y[i]:
    	distance += 1
  return distance

def solution(x, y):
  for i in range(len(x)):
    if x[i] == y[i]:
      print(x[i], end="")
  print()

for i in range(len(ids)):
  for j in range(i+1, len(ids)):
    if leven(ids[i], ids[j]) == 1:
      solution(ids[i], ids[j])
      sys.exit(0)

##import sys
##
##ids = list()
##for line in sys.stdin :
##  ids.append(line.strip())
##
##for i in range(len(ids)):
##  for j in range(i+1, len(ids)):
##    word1 = ids[i]
##    word2 = ids[j]
##    
##    distance = 0
##    for k in range(len(word1)):
##      if word1[k] != word2[k]:
##        distance += 1
##
##    if distance == 1:
##      string = ""
##      for k in range(len(x)):
##        if word1[k] == word2[k]:
##          string += word1[k]
##      print(string)
##      sys.exit(0)
    
