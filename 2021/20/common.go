package main

import (
	"bufio"
	"os"
)

type C struct{ R, C int }

type Image struct {
	iea  string
	grid map[C]bool
	minR, maxR,
	minC, maxC int
	background bool
}

func (i *Image) Enhance() {
	newGrid := map[C]bool{}

	i.minR--
	i.maxR++
	i.minC--
	i.maxC++

	for row := i.minR; row <= i.maxR; row++ {
		for col := i.minC; col <= i.maxC; col++ {
			c := C{row, col}
			pixel := i.fingerprint(c)
			newGrid[c] = pixel
		}
	}

	i.grid = newGrid

	if i.iea[0b000000000] == '#' && i.iea[0b111111111] == '.' {
		i.background = !i.background
	}
}

func (i Image) CountPixelsOn() int {
	count := 0
	for _, v := range i.grid {
		if v {
			count++
		}
	}
	return count
}

func parse() *Image {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	iea := scanner.Text()
	scanner.Scan()

	maxR, maxC := 0, 0
	grid := map[C]bool{}
	for row := 0; scanner.Scan(); row++ {

		if row > maxR {
			maxR = row
		}

		line := scanner.Text()
		for col, char := range line {
			switch char {
			case '.':
				grid[C{row, col}] = false
			case '#':
				grid[C{row, col}] = true
				if col > maxC {
					maxC = col
				}
			default:
				panic("wrong char")
			}
		}
	}

	return &Image{
		iea:        iea,
		grid:       grid,
		minR:       0,
		maxR:       maxR,
		minC:       0,
		maxC:       maxC,
		background: false,
	}
}

func (i Image) getPixel(c C) bool {
	if pixel, ok := i.grid[c]; ok {
		return pixel
	} else {
		return i.background
	}
}

func (i Image) fingerprint(c C) bool {
	var fingerprint uint
	for row := c.R - 1; row <= c.R+1; row++ {
		for col := c.C - 1; col <= c.C+1; col++ {
			fingerprint <<= 1
			if i.getPixel(C{row, col}) {
				fingerprint |= 1
			}
		}
	}
	return i.iea[fingerprint] == '#'
}
