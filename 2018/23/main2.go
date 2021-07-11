package main

import (
	"fmt"
	"sort"
)

func main() {
	nanobots := parse()
	min, max := findMinMax(nanobots)

	cellSize := minRadius(nanobots) / 3

	numPoints := ((max.X - min.X) / cellSize) * ((max.Y - min.Y) / cellSize) * ((max.Z - min.Z) / cellSize)
	// Pre-populate the list of possible points
	// taking samples, equidistant in all three axis,
	// with a distance equal to half the minimum radius
	candidates := make([]P, 0, numPoints)
	for x := min.X; x <= max.X; x += cellSize {
		for y := min.Y; y <= max.Y; y += cellSize {
			for z := min.Z; z <= max.Z; z += cellSize {
				candidates = append(candidates, P{x, y, z})
			}
		}
	}

	for {
		candidates = keepTheBestNPoints(nanobots, candidates, 5)

		// increase the granularity of the research
		cellSize /= 2

		// generation won't make progress when delta is 0.
		// all points will collapse in themselves.
		if cellSize == 0 {
			fmt.Println(distFromOrigin(candidates[0]))
			return
		}

		candidates = generatesPoints(candidates, cellSize)
	}
}

func distFromOrigin(p P) int {
	return manhattan(P{0, 0, 0}, p)
}

// Generates the 27 points  of the cube 3x3 that is around
// the point to consider. In this way we're exploring the space
// around our best candidates.
// Use a set in order to avoid duplicates when points are very
// close and delta is very low.
func generatesPoints(points []P, delta int) []P {
	newPoints := map[P]struct{}{}
	for _, p := range points {
		deltas := []int{-delta, 0, +delta}
		for _, dx := range deltas {
			for _, dy := range deltas {
				for _, dz := range deltas {
					newPoints[P{p.X + dx, p.Y + dy, p.Z + dz}] = struct{}{}
				}
			}
		}
	}

	// overwrite the candidates with the new points
	// old points are always included during generation.
	candidates := make([]P, 0, len(newPoints))
	for k := range newPoints {
		candidates = append(candidates, k)
	}
	return candidates
}

// Order all candidates for the number of nanobots in range and
// for distance from origin.
// Keep the best N of them and discard all the others.
func keepTheBestNPoints(nanobots []Nanobot, toDo []P, n int) []P {
	// for every point we want to consider,
	// find how many nanobots are in range
	inRange := map[P]int{}
	for _, p := range toDo {
		for _, nano := range nanobots {
			if nano.InRange(p) {
				inRange[p]++
			}
		}
	}

	// order the points by the number of nanobots
	// in range.
	// If equal take the closest points to the origin (0,0)
	sort.Slice(toDo, func(i, j int) bool {
		if inRange[toDo[i]] == inRange[toDo[j]] {
			return distFromOrigin(toDo[i]) < distFromOrigin(toDo[j])
		}
		return inRange[toDo[i]] > inRange[toDo[j]]
	})

	// prune the list in order to consider only the best points
	return toDo[:n]
}

// find the minimum radius in a list of nanobots
func minRadius(nanobots []Nanobot) int {
	min := 1 << 31
	for _, nano := range nanobots {
		if nano.R < min {
			min = nano.R
		}
	}
	return min
}

// find the minimum and the maximum coordinates
// on all three axis
func findMinMax(nanobots []Nanobot) (min, max P) {
	min = P{1 << 31, 1 << 31, 1 << 31}
	max = P{-1 << 31, -1 << 31, -1 << 31}
	for _, nano := range nanobots {
		if nano.C.X < min.X {
			min.X = nano.C.X
		}
		if nano.C.Y < min.Y {
			min.Y = nano.C.Y
		}
		if nano.C.Z < min.Z {
			min.Z = nano.C.Z
		}
		if nano.C.X > max.X {
			max.X = nano.C.X
		}
		if nano.C.Y > max.Y {
			max.Y = nano.C.Y
		}
		if nano.C.Z > max.Z {
			max.Z = nano.C.Z
		}
	}
	return
}
