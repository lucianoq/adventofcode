package main

import (
	"fmt"
)

const (
	gridSize         = 300
	gridSerialNumber = 9221
)

func main() {
	r := FindSquare(gridSerialNumber)
	//fmt.Printf("power:%d, x:%d, y:%d, size:%d\n", r.PowerLevel, r.X, r.Y, r.SquareSize)
	fmt.Printf("%d,%d,%d\n", r.X, r.Y, r.SquareSize)
}

func PowerLevel(x, y, serial int) int {
	rackId := x + 10
	powerLevel := rackId * y
	powerLevel += serial
	powerLevel *= rackId
	powerLevel /= 100
	powerLevel %= 10
	powerLevel -= 5
	return powerLevel
}

type Result struct {
	PowerLevel int
	X          int
	Y          int
	SquareSize int
}

func FindSquare(serial int) *Result {
	var grid [gridSize + 1][gridSize + 1]int

	for i := 1; i <= gridSize; i++ {
		for j := 1; j <= gridSize; j++ {
			grid[i][j] = PowerLevel(i, j, serial)
		}
	}
	res := make(chan *Result, gridSize)

	for i := 1; i <= gridSize; i++ {
		go func(i int) {
			maxSquareX := -1
			maxSquareY := -1
			maxSquareSize := -1
			maxPower := -1 << 63
			for j := 1; j <= gridSize; j++ {
				for delta := 0; i+delta <= gridSize && j+delta <= gridSize; delta++ {
					squarePower := 0
					for k := 0; k <= delta; k++ {
						for m := 0; m <= delta; m++ {
							squarePower += grid[i+k][j+m]
						}
					}
					if squarePower > maxPower {
						maxPower = squarePower
						maxSquareX = i
						maxSquareY = j
						maxSquareSize = delta + 1
					}
				}
			}
			res <- &Result{maxPower, maxSquareX, maxSquareY, maxSquareSize}
		}(i)
	}

	maxResult := &Result{PowerLevel: -1 << 63}
	for i := 1; i <= gridSize; i++ {
		r := <-res
		if r.PowerLevel > maxResult.PowerLevel {
			maxResult = r
		}
	}

	return maxResult
}
