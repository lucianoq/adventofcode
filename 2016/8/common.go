package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	width  = 50
	height = 6
)

var display = [width][height]bool{}

func rotateRow(y, by int) {
	var newRow [width]bool
	for i := 0; i < width; i++ {
		newRow[(i+by)%width] = display[i][y]
	}

	for i := 0; i < width; i++ {
		display[i][y] = newRow[i]
	}
}

func rect(w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			display[i][j] = true
		}
	}
}

func rotateColumn(x, by int) {
	var newColumn [height]bool
	for i := 0; i < height; i++ {
		newColumn[(i+by)%height] = display[x][i]
	}

	for i := 0; i < height; i++ {
		display[x][i] = newColumn[i]
	}
}

func runInput() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		items := strings.Fields(line)
		switch items[0] {
		case "rect":
			ll := strings.Split(items[1], "x")
			w, _ := strconv.Atoi(ll[0])
			h, _ := strconv.Atoi(ll[1])
			rect(w, h)
		case "rotate":
			switch items[1] {
			case "row":
				y, _ := strconv.Atoi(strings.Split(items[2], "=")[1])
				by, _ := strconv.Atoi(items[4])
				rotateRow(y, by)
			case "column":
				x, _ := strconv.Atoi(strings.Split(items[2], "=")[1])
				by, _ := strconv.Atoi(items[4])
				rotateColumn(x, by)
			}
		}
	}
}

func printDisplay(clean bool) {
	if clean {
		fmt.Print("\033[H\033[2J")
	}
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if display[i][j] {
				fmt.Print("█")
				//fmt.Print("#")
			} else {
				fmt.Print(" ")
				//fmt.Print("·")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
