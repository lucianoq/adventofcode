package main

import (
	"bufio"
	"fmt"
	"os"
)

var pad = [][]string{
	{" ", " ", "1", " ", " "},
	{" ", "2", "3", "4", " "},
	{"5", "6", "7", "8", "9"},
	{" ", "A", "B", "C", " "},
	{" ", " ", "D", " ", " "},
}

func main() {
	pin := ""

	// start from '5'
	x, y := 2, 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		for _, c := range scanner.Text() {
			switch c {
			case 'U':
				x1 := x - 1
				if x1 >= 0 && pad[x1][y] != " " {
					x = x1
				}
			case 'D':
				x1 := x + 1
				if x1 < 5 && pad[x1][y] != " " {
					x = x1
				}
			case 'R':
				y1 := y + 1
				if y1 < 5 && pad[x][y1] != " " {
					y = y1
				}
			case 'L':
				y1 := y - 1
				if y1 >= 0 && pad[x][y1] != " " {
					y = y1
				}
			}
		}
		pin += pad[x][y]
	}

	fmt.Println(pin)
}
