import sys

string = list()
alpha = ('a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z')
twice = 0
triple = 0
check2 = 0
check3 = 0

for line in sys.stdin :
    string.append(line)

for i in range (0,len(string)):
    
    check_twice = 0
    check_triple = 0

    for j in range (0,len(alpha)):
        
        counter = string[i].count(alpha[j])
        if counter == 2 and check_twice == 0:
            twice += 1
            check_twice = 1
        elif counter == 3 and check_triple == 0:
            triple += 1
            check_triple = 1
            
checksum = twice*triple
print (checksum)
