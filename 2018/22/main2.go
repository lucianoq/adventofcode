package main

import "fmt"

func main() {
	g := makeGridEL(20, 100)
	// Print(g, 5, 10)
	fmt.Println(bfs(g))
}

const (
	Neither = iota
	Torch
	Gear
)

func (c C) Adjacents(overX, overY int) []C {
	res := make([]C, 0, 4)
	for _, off := range [][]int{{0, +1}, {0, -1}, {-1, 0}, {+1, 0}} {
		newP := C{c.X + off[0], c.Y + off[1]}
		if newP.Y >= 0 && newP.Y <= TargetRow+overY {
			if newP.X >= 0 && newP.X <= TargetCol+overX {
				res = append(res, newP)
			}
		}
	}
	return res
}

type Step struct {
	Pos     C
	Holding int
}

func bfs(grid map[C]int) int {
	appendIfBetter := func(toDo []Step, discovered map[Step]int, step Step, elapsed int) []Step {
		t, ok := discovered[step]
		if ok && elapsed >= t {
			return toDo
		}
		discovered[step] = elapsed
		return append(toDo, step)
	}

	start := Step{
		Pos:     C{0, 0},
		Holding: Torch,
	}

	toDo := []Step{start}
	discovered := map[Step]int{start: 0}

	for len(toDo) > 0 {
		var curr Step
		curr, toDo = toDo[0], toDo[1:]

		minTimeArrivedHere := discovered[curr]

		for _, to := range curr.Pos.Adjacents(20, 100) {
			// if the tool I'm holding is fine for the destination area type
			// Reminder:
			// Neither(0) !== Rocky(0)
			// Torch(1)   !== Wet(1)
			// Gear(2)    !== Narrow(2)
			if curr.Holding != grid[to]%3 {
				// This branch includes moving from:
				//
				// Neither(0) -> Wet(1)
				// Neither(0) -> Narrow(2)
				// Torch(1) -> Rocky(0)
				// Torch(1) -> Narrow(2)
				// Gear(2) -> Rocky(0)
				// Gear(2) -> Wet(1)
				//
				// that are always two different constants
				nextStep := Step{
					Pos:     to,           // move
					Holding: curr.Holding, // keep tool
				}

				toDo = appendIfBetter(toDo, discovered, nextStep, minTimeArrivedHere+1)
			}
		}

		// Change tool
		//
		// I'm   Torch(1)   in  Rocky(0)  --> need  Gear(2)
		// I'm   Gear(2)    in  Rocky(0)  --> need  Torch(1)
		// I'm   Torch(1)   in  Narrow(2) --> need  Neither(0)
		// I'm   Neither(0) in  Wet(1)    --> need  Gear(2)
		// I'm   Neither(0) in  Narrow(2) --> need  Torch(1)
		// I'm   Gear(2)    in  Wet(1)    --> need  Neither(0)
		//
		// New tool will always be in {0,1,2} but different from
		// `pos` region type and what I'm holding now.
		// `3 - areaType - holdingNow` should do the trick.
		nextStep := Step{
			Pos:     curr.Pos,                                // stay here
			Holding: 3 - (grid[curr.Pos] % 3) - curr.Holding, // change tool
		}

		toDo = appendIfBetter(toDo, discovered, nextStep, minTimeArrivedHere+7)
	}

	goalStep := Step{
		Pos:     C{TargetCol, TargetRow},
		Holding: Torch,
	}

	if elapsed, ok := discovered[goalStep]; ok {
		return elapsed
	}
	return -1
}
