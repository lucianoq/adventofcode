package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const limit = 10000 //use this for input
//const limit = 32 //use this for testinput

type Point struct {
	X int
	Y int
}

func main() {
	coord := make([]Point, 0)

	line := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		var x, y int
		_, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		if err != nil {
			log.Fatal(err)
		}
		coord = append(coord, Point{X: x, Y: y})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// find max
	var maxX, maxY = math.MinInt64, math.MinInt64
	for _, c := range coord {
		if c.X > maxX {
			maxX = c.X
		}
		if c.Y > maxY {
			maxY = c.Y
		}
	}

	// init matrix
	matrix := make([][]int, maxX+1)
	for i := range matrix {
		matrix[i] = make([]int, maxY+1)
	}

	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			for _, c := range coord {
				distance := Abs(c.X-i) + Abs(c.Y-j)
				matrix[i][j] += distance
			}
		}
	}

	count := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] < limit {
				count++
			}
		}
	}

	fmt.Println(count)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
