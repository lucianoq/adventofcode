captcha = input("Insert a captcha: ")
string = str(captcha)
max = len(string)-1
output = 0
i = 0

while i < max :
        if string[i] == string[i+1] :
                output += int(string[i])
                i += 1
        else :
                i += 1
                
if string[0] == string[max] :
        output = output+ int(string[max])
                
print (output)
