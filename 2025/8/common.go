package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type P struct{ x, y, z int }

type Pair struct{ p1, p2 P }

type Distance struct {
	p    Pair
	dist int
}

func parseInput() []P {
	var list []P
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Split(scanner.Text(), ",")
		x, _ := strconv.Atoi(ff[0])
		y, _ := strconv.Atoi(ff[1])
		z, _ := strconv.Atoi(ff[2])
		list = append(list, P{x, y, z})
	}
	return list
}

func getDistances(list []P) []Distance {
	var d []Distance
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			d = append(d, Distance{Pair{list[i], list[j]}, euclideanDistance(list[i], list[j])})
		}
	}
	sort.Slice(d, func(i, j int) bool { return d[i].dist < d[j].dist })
	return d
}

// This is not returning the Euclidean distance but the square of it.
// It can still be sorted like the correct one, but we get rid of
// the math library import, the sqrt function, the cast to float64 and back to int.
func euclideanDistance(p, q P) int {
	return (p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y) + (p.z-q.z)*(p.z-q.z)
}

type CircuitGroup struct {
	getCircuit      map[P]int   // Point -> CircuitID
	circuits        map[int][]P // CircuitID -> Points
	autoIncrementID int
}

func NewCircuitGroup() *CircuitGroup {
	return &CircuitGroup{
		getCircuit:      make(map[P]int),
		circuits:        make(map[int][]P),
		autoIncrementID: 0,
	}
}

func (cg *CircuitGroup) Connect(pair Pair) {
	c1, p1HasC := cg.getCircuit[pair.p1]
	c2, p2HasC := cg.getCircuit[pair.p2]

	switch {
	case p1HasC && p2HasC:
		if c1 != c2 {
			// transfer all circuit 2 into 1
			for _, p := range cg.circuits[c2] {
				cg.getCircuit[p] = c1
				cg.circuits[c1] = append(cg.circuits[c1], p)
			}
			delete(cg.circuits, c2)
		}

	case p1HasC:
		// assign point 2 to circuit 1
		cg.getCircuit[pair.p2] = c1
		cg.circuits[c1] = append(cg.circuits[c1], pair.p2)

	case p2HasC:
		// assign point 1 to circuit 2
		cg.getCircuit[pair.p1] = c2
		cg.circuits[c2] = append(cg.circuits[c2], pair.p1)

	default:
		// assign both points to a new circuit
		cg.getCircuit[pair.p1], cg.getCircuit[pair.p2] = cg.autoIncrementID, cg.autoIncrementID
		cg.circuits[cg.autoIncrementID] = append(cg.circuits[cg.autoIncrementID], pair.p1, pair.p2)
		cg.autoIncrementID++
	}
}
