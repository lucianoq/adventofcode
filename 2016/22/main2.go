package main

import "fmt"

func main() {
	//initialGrid := parseGrid()
	//Print(initialGrid)

	moves := 0

	// This is my initial situation, where:
	// G is my goal
	// D is my data load
	// E is the empty node I can use to carry data around

	// +-----------------------------------------------+
	// |G                                             D|
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |          -------------------------------------|
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |              E                                |
	// +-----------------------------------------------+

	// First thing I need to carry the Empty where I need it.
	// Now D is one step closer to the Goal.

	// +-----------------------------------------------+
	// |G       +----------------------------------D-> |
	// |        |                33                    |
	// |        |                                      |
	// |        |                                      |
	// |        |                                      |
	// |        |                                      |
	// |      28|                                      |
	// |        |                                      |
	// |        |                                      |
	// |        | -------------------------------------|
	// |        |                                      |
	// |        |                                      |
	// |        |                                      |
	// |        |  7                                   |
	// |        +----+E                                |
	// +-----------------------------------------------+

	// This requires 68 moves.
	moves += 7 + 28 + 33

	// And the new situation, after 68 moves is the following

	// +-----------------------------------------------+
	// |G                                            DE|
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |          -------------------------------------|
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// |                                               |
	// +-----------------------------------------------+

	// Now, to shift D left by 1 position, I need to
	// carry around E until I swap with D from left.
	//
	//       +-----+-----+-----+
	//       |     |     |     |
	//       |     |  D  |  E  |
	//       |     |     |     |
	//       |  +----->  |  +  |
	//       |  |  |     |  |  |
	//       +--|--+-----+--|--+
	//       |  |  |     |  |  |
	//       |  +-----+-----+  |
	//       |     |     |     |
	//       +-----+-----+-----+
	//
	// This requires exactly 5 moves per shift.

	// As I need 33 shifts:
	moves += 33 * 5

	fmt.Println(moves)
}

func Print(grid Grid) {
	maxX, maxY := -1, -1
	for k := range grid {
		if k.X > maxX {
			maxX = k.X
		}
		if k.Y > maxY {
			maxY = k.Y
		}
	}

	fmt.Print("    |  ")
	for i := 0; i <= maxX; i++ {
		fmt.Printf("%3d  ", i)
	}
	fmt.Println()

	fmt.Print("====+")
	for i := 0; i <= maxX; i++ {
		fmt.Printf("=====")
	}
	fmt.Println()

	for j := 0; j <= maxY; j++ {
		fmt.Printf("%3d | ", j)
		for i := 0; i <= maxX; i++ {
			d := grid[P{i, j}]
			if d.Used == 0 {
				fmt.Printf(" ____")
			} else {
				if d.Used > 92 {
					fmt.Printf(" ####")
				} else {
					fmt.Printf(" %3d ", d.Used)
				}
			}
		}
		fmt.Println()
	}
}
