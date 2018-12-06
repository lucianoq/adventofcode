import sys
hz_list = list()
hz = 0
i = 0
twice = True

while twice == True:
    for line in sys.stdin :
        hz += int(line)
        hz_list.append(hz)
        if i > 0 :
            if hz_list[i] == hz_list[i-1] :
                print (hz_list[i])
                twice = False
        i+=1
                
