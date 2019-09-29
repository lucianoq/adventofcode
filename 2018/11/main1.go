package main

import "fmt"

const (
	size             = 300
	gridSerialNumber = 9221
)

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

func main() {
	maxSquareX := -1
	maxSquareY := -1
	maxPower := -1 << 63

	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			squarePower := 0
			for k := -1; k < 2; k++ {
				for m := -1; m < 2; m++ {
					squarePower += PowerLevel(i+k, j+m, gridSerialNumber)
				}
			}
			if squarePower > maxPower {
				maxPower = squarePower
				maxSquareX = i - 1
				maxSquareY = j - 1
			}
		}
	}
	fmt.Printf("%d,%d\n", maxSquareX, maxSquareY)
}
