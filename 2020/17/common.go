package main

import "fmt"

type Dimension map[Point]struct{}

func (d Dimension) Step() Dimension {
	toConsider := map[Point]struct{}{}

	for p := range d {
		toConsider[p] = struct{}{}
		for _, n := range p.Neighbours() {
			toConsider[n] = struct{}{}
		}
	}

	newDim := Dimension{}
	for p := range toConsider {

		countNeighbours := 0
		for _, n := range p.Neighbours() {
			if _, ok := d[n]; ok {
				countNeighbours++
			}
		}

		_, isActive := d[p]

		if isActive && (countNeighbours == 2 || countNeighbours == 3) {
			newDim[p] = struct{}{}
		}

		if !isActive && countNeighbours == 3 {
			newDim[p] = struct{}{}
		}
	}

	return newDim
}

func main() {
	dim := parse()
	for i := 0; i < 6; i++ {
		dim = dim.Step()
	}
	fmt.Println(len(dim))
}
