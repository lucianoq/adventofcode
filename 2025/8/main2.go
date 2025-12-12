package main

import "fmt"

func main() {
	list := parseInput()
	distances := getDistances(list)
	cg := NewCircuitGroup()

	var lastPair Pair
	for cg.AddedPoints() < len(list) {
		var d Distance
		d, distances = distances[0], distances[1:]
		lastPair = d.p
		cg.Connect(d.p)
	}

	fmt.Println(lastPair.p1.x * lastPair.p2.x)
}

func (cg *CircuitGroup) Len() int {
	return len(cg.circuits)
}

func (cg *CircuitGroup) AddedPoints() int {
	return len(cg.getCircuit)
}

func (cg *CircuitGroup) GetCircuit() []P {
	for _, v := range cg.circuits {
		return v
	}
	return nil
}
