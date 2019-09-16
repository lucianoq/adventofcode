package main

import "fmt"

func main() {
	input := parse()

	// Cube coordinates system
	// ref: https://www.redblobgames.com/grids/hexagons/#coordinates
	//
	// 🡖 x = from north-west to south-east
	// 🡕 y = from south-west to north-east
	// 🡒 z = from west to east
	//
	//          \    /
	//      x-   +--+   y+
	//          /    \
	//   z-  --+      +--  z+
	//          \    /
	//      y-   +--+   x+
	//          /    \
	//
	var x, y, z int

	var maxDistance int

	for _, step := range input {
		switch step {
		case "n":
			x--
			y++
		case "s":
			x++
			y--
		case "ne":
			y++
			z++
		case "sw":
			y--
			z--
		case "nw":
			x--
			z--
		case "se":
			x++
			z++
		}

		distance := (abs(x) + abs(y) + abs(z)) / 2

		if distance > maxDistance {
			maxDistance = distance
		}
	}

	fmt.Println(maxDistance)
}
