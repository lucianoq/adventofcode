package main

import (
	"fmt"
	"log"
)

const (
	gridSize         = 300
	gridSerialNumber = 9221
)

func main() {
	_, x, y, size := FindSquare(gridSerialNumber)
	fmt.Printf("%d,%d,%d\n", x, y, size)
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

func FindSquare(serial int) (int, int, int, int) {
	var grid [gridSize][gridSize]int

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			grid[i][j] = PowerLevel(i, j, serial)
		}
	}

	//for i:=0; i<gridSize; i++ {
	//	for j:=0; j<gridSize; j++ {
	//		fmt.Printf("%d,", grid[i][j])
	//	}
	//	fmt.Println()
	//}
	//os.Exit(0)

	maxSquareX := -1
	maxSquareY := -1
	maxSquareSize := -1
	maxPower := -1 << 63

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			log.Printf("(%d,%d)", i, j)
			maxSize := min(i, j, gridSize-i, gridSize-j)
			log.Printf("Here I can arrive max at %dx%d", 2*(maxSize-1)+1, 2*(maxSize-1)+1)
			for offset := 0; offset < maxSize; offset++ {
				squarePower := 0
				for k := -offset; k <= offset; k++ {
					for m := -offset; m <= offset; m++ {
						squarePower += grid[i+k][j+m]
					}
				}
				if squarePower > maxPower {
					maxPower = squarePower
					maxSquareX = i - offset
					maxSquareY = j - offset
					maxSquareSize = 2*offset + 1
				}
			}
		}
	}
	return maxPower, maxSquareX, maxSquareY, maxSquareSize
}

func min(xs ...int) int {
	m := 1<<63 - 1
	for _, x := range xs {
		if x < m {
			m = x
		}
	}
	return m
}
