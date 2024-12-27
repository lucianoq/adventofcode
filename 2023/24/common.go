package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ x, y, z int }

type Hail struct {
	ID  int
	Pos P
	Vel P
}

func (h Hail) String() string {
	return fmt.Sprintf("Hailstone : (%d, %d, %d) @ (%d, %d, %d)", h.Pos.x, h.Pos.y, h.Pos.z, h.Vel.x, h.Vel.y, h.Vel.z)
}

func parseInput() []Hail {
	hail := []Hail{}

	scanner := bufio.NewScanner(os.Stdin)
	for id := 0; scanner.Scan(); id++ {
		line := scanner.Text()
		var px, py, pz, vx, vy, vz int
		fmt.Sscanf(line, "%d, %d, %d @ %d, %d, %d", &px, &py, &pz, &vx, &vy, &vz)
		hail = append(hail, Hail{
			ID:  id,
			Pos: P{px, py, pz},
			Vel: P{vx, vy, vz},
		})
	}

	return hail
}
