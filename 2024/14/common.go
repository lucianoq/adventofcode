package main

import (
	"bufio"
	"fmt"
	"os"
)

var max = P{101, 103}

type P struct{ x, y int }
type Robot struct{ Pos, Vel P }

func parseInput() []Robot {
	robots := []Robot{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var x, y, vx, vy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &x, &y, &vx, &vy)
		robots = append(robots, Robot{P{x, y}, P{vx, vy}})
	}
	return robots
}
