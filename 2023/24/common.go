package main

import (
	"bufio"
	"fmt"
	"os"
)

type P struct{ x, y, z float64 }

type HailStone struct{ Pos, Vel P }

func parse() []HailStone {
	hail := []HailStone{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ice := HailStone{}
		_, _ = fmt.Sscanf(
			scanner.Text(),
			"%f, %f, %f @ %f, %f, %f",
			&ice.Pos.x,
			&ice.Pos.y,
			&ice.Pos.z,
			&ice.Vel.x,
			&ice.Vel.y,
			&ice.Vel.z,
		)
		hail = append(hail, ice)
	}
	return hail
}
