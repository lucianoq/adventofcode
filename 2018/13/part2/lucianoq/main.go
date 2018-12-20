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

//testinput
//const (
//	width  = 7
//	height = 7
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

		crashed := false
		for _, c := range carts {
			c.Step()
			if Crash(c) {
				crashed = true
			}
		}
		if crashed {
			Clean()
			if len(carts) == 1 {
				fmt.Printf("%d,%d\n", carts[0].X, carts[0].Y)
				return
			}
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
		c = "-"
	case "<":
		dir = W
		c = "-"
	case "^":
		dir = N
		c = "|"
	case "v":
		dir = S
		c = "|"
	}

	if dir != null {
		carts = append(carts, NewCart(x, y, dir))
	}

	return Cell{
		Char: c,
		X:    x,
		Y:    y,
	}
}

//func Dump() {
//	for j := 0; j < height; j++ {
//	CELL:
//		for i := 0; i < width; i++ {
//			for _, c := range carts {
//				if c.X == grid[i][j].X && c.Y == grid[i][j].Y {
//					color.Set(color.BgRed)
//					fmt.Print(c.String())
//					color.Unset()
//					continue CELL
//				}
//			}
//			fmt.Print(grid[i][j].Char)
//		}
//		fmt.Println()
//	}
//
//	for i, c := range carts {
//		fmt.Printf("[%d] (%d,%d) %d %d\n", i, c.X, c.Y, c.Direction, c.NextTurn)
//	}
//}

func Crash(c *Cart) bool {
	found := false
	for _, car := range carts {
		if c.ID != car.ID {
			if c.X == car.X && c.Y == car.Y {
				found = true
				c.Crashed = true
				car.Crashed = true
			}
		}
	}
	return found
}

func Clean() {
	newCarts := carts[:0]
	for _, c := range carts {
		if !c.Crashed {
			newCarts = append(newCarts, c)
		}
	}
	carts = newCarts
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
