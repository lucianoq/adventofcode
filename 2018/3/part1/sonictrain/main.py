SIZE = 1000

class Claim:
    def __init__(self, x, y, w, h):
        self.x = x
        self.y = y
        self.w = w
        self.h = h

claims = []
with open("input") as f:
    for line in f:
        a = line.split(" @ ")
        b = a[1].split(": ")
        position = b[0].split(",")
        dimension = b[1].split("x")

        claims.append(Claim(int(position[0]), int(position[1]), int(dimension[0]), int(dimension[1])))

# max_x = -1
# max_y = -1
# for c in claims:
#     new_x_max =c.x + c.w
#     if new_x_max > max_x:
#         max_x = new_x_max
#     new_y_max = c.y + c.h
#     if new_y_max > max_y:
#         max_y = new_y_max

matrix = [[0 for y in range(SIZE)] for x in range(SIZE)] 

for c in claims:
    for x in range(c.x, c.x + c.w):
        for y in range(c.y, c.y + c.h):
            matrix[x][y] += 1

count = 0
for x in range(max_x):
    for y in range(max_y):
        if matrix[x][y] > 1:
            count += 1
print(count)
