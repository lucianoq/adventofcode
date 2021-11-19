package main

import "fmt"

type Space map[int]*[Square]bool

func main() {

	input := parse()

	space := Space{
		0: &input,
	}

	// after 200 minutes
	for i := 0; i < 200; i++ {
		space = next2(space)
	}

	// how many bugs are present
	count := 0
	for _, grid := range space {
		for i := 0; i < Square; i++ {
			if grid[i] {
				count++
			}
		}
	}
	fmt.Println(count)
}

func next2(space Space) Space {
	newSpace := Space{}

	minLevel, maxLevel := minMaxLevel(space)

	for level := minLevel - 1; level <= maxLevel+1; level++ {
		newSpace[level] = &[Square]bool{}

		for cell := 0; cell < Square; cell++ {

			// skip central cell
			if cell == 12 {
				continue
			}

			// +------------------------+
			// |  0 |  1 |  2 |  3 |  4 |
			// |----+----+----+----+----|
			// |  5 |  6 |  7 |  8 |  9 |
			// |----+----+----+----+----|
			// | 10 | 11 | 12 | 13 | 14 |
			// |----+----+----+----+----|
			// | 15 | 16 | 17 | 18 | 19 |
			// |----+----+----+----+----|
			// | 20 | 21 | 22 | 23 | 24 |
			// +------------------------+

			row, col := cell/Side, cell%Side
			neighbours := 0

			// TOP outside
			if row == 0 {
				if infested(space, level-1, 7) {
					neighbours++
				}
			}
			// LEFT outside
			if col == 0 {
				if infested(space, level-1, 11) {
					neighbours++
				}
			}
			// RIGHT outside
			if col == 4 {
				if infested(space, level-1, 13) {
					neighbours++
				}
			}
			// BOTTOM outside
			if row == 4 {
				if infested(space, level-1, 17) {
					neighbours++
				}
			}
			// inside from TOP
			if cell == 7 {
				for i := 0; i < Side; i++ {
					if infested(space, level+1, i) { // 0 1 2 3 4
						neighbours++
					}
				}
			}
			// inside from LEFT
			if cell == 11 {
				for i := 0; i < Side; i++ {
					if infested(space, level+1, 5*i) { // 0 5 10 15 20
						neighbours++
					}
				}
			}
			// inside from RIGHT
			if cell == 13 {
				for i := 0; i < Side; i++ {
					if infested(space, level+1, 5*i+Side-1) { // 4 9 14 19 24
						neighbours++
					}
				}
			}
			// inside from BOTTOM
			if cell == 17 {
				for i := 0; i < Side; i++ {
					if infested(space, level+1, (Side-1)*Side+i) { // 20 21 22 23 24
						neighbours++
					}
				}
			}

			// add normal top adj
			if row > 0 && cell != 17 {
				if infested(space, level, cell-Side) {
					neighbours++
				}
			}
			// add normal left adj
			if col > 0 && cell != 13 {
				if infested(space, level, cell-1) {
					neighbours++
				}
			}
			// add normal right adj
			if col < Side-1 && cell != 11 {
				if infested(space, level, cell+1) {
					neighbours++
				}
			}
			// add normal bottom adj
			if row < Side-1 && cell != 7 {
				if infested(space, level, cell+Side) {
					neighbours++
				}
			}

			if infested(space, level, cell) && neighbours != 1 {
				newSpace[level][cell] = false
				continue
			}

			if !infested(space, level, cell) && (neighbours == 1 || neighbours == 2) {
				newSpace[level][cell] = true
				continue
			}

			newSpace[level][cell] = infested(space, level, cell)
		}
	}

	clean(newSpace)

	return newSpace
}

func clean(space Space) {
	min, max := minMaxLevel(space)

	countMin, countMax := 0, 0
	for cell := 0; cell < Square; cell++ {
		if space[min][cell] {
			countMin++
		}
		if space[max][cell] {
			countMax++
		}
	}
	if countMin == 0 {
		delete(space, min)
	}
	if countMax == 0 {
		delete(space, max)
	}
}

func infested(space Space, level, cell int) bool {
	if space[level] == nil {
		return false
	}
	return space[level][cell]
}

func minMaxLevel(space Space) (int, int) {
	min, max := +999999, -999999
	for level := range space {
		if level < min {
			min = level
		}
		if level > max {
			max = level
		}
	}
	return min, max
}
