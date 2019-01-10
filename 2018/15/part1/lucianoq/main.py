#!/usr/local/bin/python3

SIZE = 32


def read(file):
    with open(file) as f:
        lines = f.readlines()

    y = 0
    for line in lines:
        if line == '':
            continue

        x = 0
        for c in line:
            if c == '#':
                grid.set(x, y, Cell(x, y, c))
            elif c == '.':
                grid.set(x, y, Cell(x, y, c))
            elif c == 'G':
                grid.set(x, y, Cell(x, y, c, Goblin(x, y)))
            elif c == 'E':
                grid.set(x, y, Cell(x, y, c, Elf(x, y)))
            x += 1
        y += 1


class Grid:
    grid = []

    def __init__(self):
        self.grid = [[None for _ in range(SIZE)] for _ in range(SIZE)]

    def dump(self):
        for j in range(SIZE):
            for i in range(SIZE):
                print(self.grid[i][j].char, end='')
            print()

    def set(self, x, y, cell):
        self.grid[x][y] = cell

    def get(self, x, y):
        return self.grid[x][y]

    def adjacents(self, x, y):
        ad = []
        if y - 1 >= 0:
            ad.append(self.grid[x][y - 1])
        if x - 1 >= 0:
            ad.append(self.grid[x - 1][y])
        if x + 1 < SIZE:
            ad.append(self.grid[x + 1][y])
        if y + 1 < SIZE:
            ad.append(self.grid[x][y + 1])
        return [x for x in ad if x.is_free()]

    def reachable(self, x, y, target_x, target_y, visited, deep=0):
        # print(f"reachable started. from {x},{y} to {target_x},{target_y}")
        # print(visited)
        visited.append(self.grid[x][y])
        for a in self.adjacents(x, y):
            # print(f"({a.x},{a.y})")
            if a not in visited:
                # print(f"({a.x},{a.y})")
                if a.x == target_x and a.y == target_y:
                    # print(deep)
                    return True, deep
                # print("recursion")
                reach, depth = self.reachable(a.x, a.y, target_x, target_y, visited, deep + 1)
                if reach:
                    # print(depth)
                    return True, depth
        return False

    def reachable_unit(self, unit1, unit2):
        return self.reachable(unit1.x, unit1.y, unit2.x, unit2.y, [], 0)


class Cell:
    char = ''
    x = 0
    y = 0
    unit = None

    def __init__(self, x, y, char, unit=None):
        self.x = x
        self.y = y
        self.char = char
        self.unit = unit

    def __eq__(self, other):
        return self.x == other.x and self.y == other.y

    def __repr__(self):
        return f"{self.char}({self.x},{self.y})"

    def is_wall(self):
        return self.char == '#'

    def is_free(self):
        return self.char == '.'


class Unit:
    x = 0
    y = 0
    char = ''

    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __repr__(self):
        return f"{self.char}({self.x},{self.y})"

    def adjacents(self):
        return grid.adjacents(self.x, self.y)

    def is_goblin(self):
        return self.char == 'G'

    def is_elf(self):
        return self.char == 'E'

    def can_reach(self, x, y):
        x, _ = grid.reachable(self.x, self.y, x, y, [], 0)
        return x

    def can_reach_unit(self, unit):
        x, _ = grid.reachable(self.x, self.y, unit.x, unit.y, [], 0)
        return x

    def dist(self, unit):
        _, x = grid.reachable(self.x, self.y, unit.x, unit.y, [], 0)
        return x

    def move(self):
        # reachable = []
        # if u.is_elf():
        #     e = u
        #     for g in only_goblins():
        #         for ag in g.adjacents():
        #             if e.can_reach_unit(ag):
        #                 reachable.append(ag)
        #                 print(f'elf {e} can reach cell {ag} adjacent to goblin {g}')
        # if u.is_goblin():
        #     g = u
        #     for e in only_elves():
        #         for ae in e.adjacents():
        #             if g.can_reach_unit(ae):
        #                 reachable.append(ae)
        #                 print(f'goblin {g} can reach cell {ae} adjacent to elf {e}')
        #
        # nearestUnit = None
        # nearestDist = 99999
        # for r in reachable:
        #     dist = u.dist(r)
        #     if dist < nearestDist:
        #         nearestDist = dist
        #         nearestUnit = r
        #
        # print(nearestUnit)
        # print(nearestDist)


class Elf(Unit):
    def __init__(self, x, y):
        self.char = 'E'
        Unit.__init__(self, x, y)


class Goblin(Unit):
    def __init__(self, x, y):
        self.char = 'G'
        Unit.__init__(self, x, y)


def priority_units():
    units = []
    for j in range(SIZE):
        for i in range(SIZE):
            cell = grid.get(i, j)
            if cell.unit is not None:
                units.append(cell.unit)
    return units


def only_elves():
    return [x for x in priority_units() if x.is_elf()]


def only_goblins():
    return [x for x in priority_units() if x.is_goblin()]


if __name__ == '__main__':
    grid = Grid()
    read('testinput')
    grid.dump()

    # for j in range(SIZE):
    #     for i in range(SIZE):
    #         if grid.reachable(0, 0, i, j, []):
    #             print("@", end='')
    #         else:
    #             print('-', end='')
    #     print()

    print(priority_units())
    print(only_elves())
    print(only_goblins())

    # elf = only_elves()[0]
    # gob = only_goblins()[0]
    # print(elf.can_reach_unit(gob))

    # Targets:      In range:     Reachable:    Nearest:      Chosen:
    # #######       #######       #######       #######       #######
    # #E..G.#       #E.?G?#       #E.@G.#       #E.!G.#       #E.+G.#
    # #...#.#  -->  #.?.#?#  -->  #.@.#.#  -->  #.!.#.#  -->  #...#.#
    # #.G.#G#       #?G?#G#       #@G@#G#       #!G.#G#       #.G.#G#
    # #######       #######       #######       #######       #######

    for u in priority_units():
        u.move()
