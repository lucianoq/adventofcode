package main

const NumTiles = 5

func getRisk(grid map[C]int, c C) int {
	val := grid[C{c.x % Size, c.y % Size}] + c.x/Size + c.y/Size
	for val > 9 {
		val -= 9
	}
	return val
}
