import sys

data = list()

for line in sys.stdin:
    data.append(line)

for i in range(len(data)):

    ID = i+1
    
    a = data[i].split(" @ ")
    b = a[1].split(": ")

    position = b[0].split(",")
    dimension = b[1].split("x")

    X_cordinate = position[0]
    Y_cordinate = position[1]
    lenght = dimension[0]
    height = dimension[1]

    ##print ("La matrice #{N} ha origine nel punto con cordinate x={X} ed y={Y} ed ha larghezza {L} e altezza {A}".format (N=ID, X=X_cordinate, Y=Y_cordinate, L=lenght, A=height))
