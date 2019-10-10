package main

import (
	"bufio"
	"fmt"
	"os"
)

const size = 50

func CountResources(area [size][size]rune) int {
	countTrees, countLumber := 0, 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			switch area[i][j] {
			case '|':
				countTrees++
			case '#':
				countLumber++
			}
		}
	}

	return countTrees * countLumber
}

func Minute(area [size][size]rune) [size][size]rune {
	newArea := [size][size]rune{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			newArea[i][j] = area[i][j]

			adj := neighbour(area, i, j)
			switch area[i][j] {
			case '.':
				if adj['|'] >= 3 {
					newArea[i][j] = '|'
				}
			case '|':
				if adj['#'] >= 3 {
					newArea[i][j] = '#'
				}
			case '#':
				if adj['#'] < 1 || adj['|'] < 1 {
					newArea[i][j] = '.'
				}
			}
		}
	}

	return newArea
}

func parse() [size][size]rune {
	area := [size][size]rune{}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, c := range line {
			area[i][j] = c
		}
	}

	return area
}

func Print(area [size][size]rune) {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(string(area[i][j]))
		}
		fmt.Println()
	}
	fmt.Println()
}

func neighbour(area [size][size]rune, i, j int) map[rune]int {
	res := make(map[rune]int, 0)

	for _, offsetX := range []int{-1, 0, +1} {
		for _, offsetY := range []int{-1, 0, +1} {
			if offsetX|offsetY != 0 {
				x, y := i+offsetX, j+offsetY
				if x >= 0 && x < size {
					if y >= 0 && y < size {
						res[area[x][y]]++
					}
				}
			}
		}
	}
	return res
}
