package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ X, Y int }

type Sensor struct {
	Pos    P
	Beacon P

	// derived, but useful to store it
	// to avoid calculating over and over
	Dist int
}

func parse() []Sensor {
	sensors := []Sensor{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		var sx, sy, bx, by int
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		sensor := P{sx, sy}
		beacon := P{bx, by}
		sensors = append(sensors, Sensor{sensor, beacon, manhattan(sensor, beacon)})
	}
	return sensors
}

func manhattan(a, b P) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
