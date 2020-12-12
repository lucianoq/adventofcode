package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	x, y := 0, 0
	wx, wy := 10, 1

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dir := string(line[0])
		val, _ := strconv.Atoi(line[1:])

		switch dir {

		case "N":
			wy += val

		case "S":
			wy -= val

		case "W":
			wx -= val

		case "E":
			wx += val

		case "L":
			times := val / 90
			for i := 0; i < times; i++ {
				wx, wy = -wy, wx
			}

		case "R":
			times := val / 90
			for i := 0; i < times; i++ {
				wx, wy = wy, -wx
			}

		case "F":
			x += wx * val
			y += wy * val
		}

	}

	fmt.Println(abs(x) + abs(y))
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}
