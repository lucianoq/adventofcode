package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

const (
	width  = 150
	height = 150
)

// //testinput
//const (
//	width  = 13
//	height = 6
//)

type Cell struct {
	X    int
	Y    int
	Char string
}

var grid [width][height]Cell
var carts []*Cart

func main() {
	readFile()

	for step := 0; ; step++ {
		SortCarts()
		for _, c := range carts {
			c.Step()
		}
		cr, x, y := Crash()
		if cr {
			fmt.Printf("%d,%d\n", x, y)
			return
		}
	}
}

func readFile() {
	lineID := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), "")
		for i, c := range arr {
			grid[i][lineID] = charToCell(i, lineID, c)
		}
		for i := len(arr); i < width; i++ {
			grid[i][lineID] = Cell{
				X:    i,
				Y:    lineID,
				Char: " "}
		}
		lineID++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func charToCell(x, y int, c string) Cell {
	var dir Direction
	switch c {
	case ">":
		dir = E
	case "<":
		dir = W
	case "^":
		dir = N
	case "v":
		dir = S
	}

	if dir != null {
		carts = append(carts, &Cart{
			Direction: dir,
			NextTurn:  Left,
			X:         x,
			Y:         y,
		})
	}

	return Cell{
		Char: c,
		X:    x,
		Y:    y,
	}
}

func Dump() {
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			fmt.Print(grid[i][j].Char)
		}
		fmt.Println()
	}

	for i, c := range carts {
		fmt.Printf("[%d] (%d,%d) %d %d\n", i, c.X, c.Y, c.Direction, c.NextTurn)
	}
}

func Crash() (bool, int, int) {
	found := make(map[int]bool)
	for _, c := range carts {
		if found[c.X+c.Y*width] {
			return true, c.X, c.Y
		}
		found[c.X+c.Y*width] = true
	}
	return false, 0, 0
}

func SortCarts() {
	sort.Slice(carts, func(i, j int) bool {
		if carts[i].Y < carts[j].Y {
			return true
		}
		if carts[i].Y > carts[j].Y {
			return false
		}
		if carts[i].X < carts[j].X {
			return true
		}
		return false
	})
}
