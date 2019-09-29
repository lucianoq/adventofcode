package main

import (
	"fmt"
	"math"
)

func main() {
	particles := parse()

	// Candidates to be the closest to 0 are the particles
	// with the smallest acceleration

	// find the smallest acceleration in the set
	minAcceleration := math.MaxInt64
	for _, p := range particles {
		sum := Manhattan(p.Acc)
		if sum < minAcceleration {
			minAcceleration = sum
		}
	}

	// find the particles that have it
	candidates := make([]Particle, 0)
	for _, p := range particles {
		sum := Manhattan(p.Acc)
		if sum == minAcceleration {
			candidates = append(candidates, p)
		}
	}

time:
	// actually I can start with a big number and saves a lot of iteration
	for t := 0; ; t++ {

		// This loop skips all the iterations until position
		// is consistent with acceleration for all candidate particles
		for _, c := range candidates {
			pos := c.PositionAfter(t)

			// In long term, particles, in all axis, will be in the same side
			// of their acceleration.
			if pos.X*c.Acc.X < 0 || pos.Y*c.Acc.Y < 0 || pos.Z*c.Acc.Z < 0 {
				continue time
			}
		}

		// find the closest particle in that moment (t)
		minParticle := -1
		minDistance := math.MaxInt64
		for _, c := range candidates {
			distance := Manhattan(c.PositionAfter(t))
			if distance < minDistance {
				minDistance = distance
				minParticle = c.ID
			}
		}

		// fmt.Printf("After t=%d\n", t)
		fmt.Println(minParticle)
		return
	}
}
