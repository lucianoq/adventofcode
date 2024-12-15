package main

import "fmt"

func main() {
	m, moves, robot := parseInput()
	for _, move := range moves {
		robot = apply1(m, move, robot)
	}
	fmt.Println(gps(m, 'O'))
}

func apply1(m map[P]byte, move byte, robot P) P {
	delta := getDelta(move)

	nextEmpty, err := findNextEmpty(m, robot, delta)
	if err != nil {
		return robot
	}

	closest := P{robot.x + delta.x, robot.y + delta.y}
	if m[closest] == 'O' {
		m[closest] = '.'
		m[nextEmpty] = 'O'
	}
	return closest
}
