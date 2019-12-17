package main

import "fmt"

type C struct{ x, y int }

const (
	_ = iota
	North
	South
	West
	East
)

const (
	HitWall = iota
	OneStep
	OneStepOxygenHere
)

const (
	Unknown = iota
	Wall
	Free
	Oxygen
)

var (
	right = map[int]int{
		North: East,
		South: West,
		West:  North,
		East:  South,
	}

	left = map[int]int{
		North: West,
		South: East,
		West:  South,
		East:  North,
	}

	offset = map[int]C{
		North: {0, -1},
		South: {0, 1},
		West:  {-1, 0},
		East:  {1, 0},
	}
)

func buildGrid(input chan<- int, output <-chan int) map[C]int {
	pos := C{}
	grid := map[C]int{pos: Free}
	dir := North

	for len(grid) < 1654 {
		input <- dir
		response := <-output

		switch response {
		case HitWall:
			grid[C{pos.x + offset[dir].x, pos.y + offset[dir].y}] = Wall
			dir = right[dir]
		case OneStep:
			pos = C{pos.x + offset[dir].x, pos.y + offset[dir].y}
			grid[pos] = Free
			dir = left[dir]
		case OneStepOxygenHere:
			pos = C{pos.x + offset[dir].x, pos.y + offset[dir].y}
			grid[pos] = Oxygen
			dir = left[dir]
		}
	}

	return grid
}

func print(grid map[C]int) {
	fmt.Print("\033[2J\033[H")
	for j := -21; j < 20; j++ {
		for i := -21; i < 20; i++ {
			switch grid[C{i, j}] {
			case Unknown:
				fmt.Print(" ")
			case Wall:
				fmt.Print("#")
			case Free:
				fmt.Print(".")
			case Oxygen:
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}
