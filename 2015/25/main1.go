package main

import "fmt"

var targetRow = 2947
var targetColumn = 3029

func main() {
	var code uint64 = 20151125
	i, j := 1, 1

	for {
		i, j = nextPoint(i, j)
		code = nextCode(code)

		if i == targetRow && j == targetColumn {
			fmt.Println(code)
			return
		}
	}
}

func nextCode(i uint64) uint64 {
	return i * 252533 % 33554393
}

func nextPoint(i, j int) (int, int) {
	if i == 1 {
		return j + 1, 1
	}
	return i - 1, j + 1
}
