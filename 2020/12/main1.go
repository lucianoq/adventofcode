package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	E = iota * 90
	S
	W
	N
)

func main() {
	x, y := 0, 0
	facing := E

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dir := string(line[0])
		val, _ := strconv.Atoi(line[1:])

		switch dir {
		case "N":
			y += val
		case "S":
			y -= val
		case "W":
			x -= val
		case "E":
			x += val
		case "L":
			facing = (facing - val + 360) % 360
		case "R":
			facing = (facing + val + 360) % 360
		case "F":
			switch facing {
			case N:
				y += val
			case S:
				y -= val
			case W:
				x -= val
			case E:
				x += val
			}
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
