captcha = input("Insert a captcha: ")
s = str(captcha)
len = len(s)
s0 = s[len/2:]
s1 = s[:len/2]
v = s0 + s + s1
output = 0
for i in range (0, len) :
        if s[i] == v[i+len] :
                output = output + int(s[i])
print output
