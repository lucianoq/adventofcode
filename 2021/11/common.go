package main

import (
	"bufio"
	"os"
	"strconv"
)

const Size = 10

type pair struct{ i, j int }

func parse() [Size][Size]uint8 {
	scanner := bufio.NewScanner(os.Stdin)
	grid := [Size][Size]uint8{}

	for i := 0; i < Size; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 0; j < Size; j++ {
			num, _ := strconv.ParseUint(string(line[j]), 10, 8)
			grid[i][j] = uint8(num)
		}
	}
	return grid
}
