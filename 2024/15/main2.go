package main

import "fmt"

func main() {
	m, moves, robot := parseInput()
	m, robot = grow(m, robot)
	for _, move := range moves {
		robot = apply(m, move, robot)
	}
	fmt.Println(gps(m, '['))
}

func grow(m map[P]byte, robot P) (map[P]byte, P) {
	newM := make(map[P]byte, len(m)*2)
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			p := P{i, j}
			switch m[p] {
			case '#':
				newM[P{p.x, 2 * p.y}] = '#'
				newM[P{p.x, 2*p.y + 1}] = '#'
			case 'O':
				newM[P{p.x, 2 * p.y}] = '['
				newM[P{p.x, 2*p.y + 1}] = ']'
			case '.':
				newM[P{p.x, 2 * p.y}] = '.'
				newM[P{p.x, 2*p.y + 1}] = '.'
			case '@':
				newM[P{p.x, 2 * p.y}] = '@'
				newM[P{p.x, 2*p.y + 1}] = '.'
			}
		}
	}
	Width *= 2
	return newM, P{robot.x, 2 * robot.y}
}

func apply(m map[P]byte, move byte, robot P) P {
	delta := getDelta(move)

	nextEmpty, err := findNextEmpty(m, robot, delta)
	if err != nil {
		return robot
	}

	if move == '<' || move == '>' {
		for curr := nextEmpty; curr != robot; {
			closest := P{curr.x, curr.y - delta.y}
			m[curr], m[closest] = m[closest], m[curr]
			curr = closest
		}
		return P{robot.x, robot.y + delta.y}
	}

	if move == '^' || move == 'v' {
		affected, maxLevel, err := affectedVertically(m, robot, delta.x)
		if err != nil {
			return robot
		}
		for x := maxLevel; x != robot.x; x -= delta.x {
			for col := range affected[x] {
				m[P{x + delta.x, col}], m[P{x, col}] = m[P{x, col}], m[P{x + delta.x, col}]
			}
		}

		return P{robot.x + delta.x, robot.y}
	}

	panic("unreachable")
}

func affectedVertically(m map[P]byte, robot P, deltaX int) (map[int]map[int]struct{}, int, error) {
	affected := map[int]map[int]struct{}{
		robot.x: {robot.y: {}},
	}

	for currX := robot.x; ; currX += deltaX {
		newCols, err := newColumns(m, currX+deltaX, affected[currX])
		if err != nil {
			return nil, 0, err
		}

		if len(newCols) == 0 {
			return affected, currX, nil
		}

		affected[currX+deltaX] = newCols
	}
}

func newColumns(m map[P]byte, nextX int, columns map[int]struct{}) (map[int]struct{}, error) {
	newCols := map[int]struct{}{}
	for col := range columns {
		switch m[P{nextX, col}] {
		case '#':
			return nil, NotFound
		case '[':
			newCols[col] = struct{}{}
			newCols[col+1] = struct{}{}
		case ']':
			newCols[col] = struct{}{}
			newCols[col-1] = struct{}{}
		}
	}
	return newCols, nil
}
