package main

import (
	"fmt"
	"sort"
)

func main() {
	list := parseInput()
	distances := getDistances(list)
	cg := NewCircuitGroup()

	for k := 0; k < 1000; k++ {
		var d Distance
		d, distances = distances[0], distances[1:]
		cg.Connect(d.p)
	}

	lengths := cg.Lengths()

	fmt.Println(lengths[0] * lengths[1] * lengths[2])
}

func (cg *CircuitGroup) Lengths() []int {
	var lengths []int
	for _, ls := range cg.circuits {
		lengths = append(lengths, len(ls))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lengths)))
	return lengths
}
