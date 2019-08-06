package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

const id = "ABCDEFGHIJKLMNOPQRSTUVWXYZΑΒΓΔΕΖΗΘΙΚΛΜΝΞΟΠΡΣΤΥΦΧΨΩ"

type Cell struct {
	X      int
	Y      int
	Label  string
	Senior bool
}

func main() {
	coord := make([]Cell, 0)

	line := ""
	i := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		var x, y int
		_, err := fmt.Sscanf(line, "%d, %d", &x, &y)
		if err != nil {
			log.Fatal(err)
		}
		coord = append(coord, Cell{X: x, Y: y, Label: string([]rune(id)[i]), Senior: true})
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// find min and max
	var minX, minY = math.MaxInt64, math.MaxInt64
	var maxX, maxY = math.MinInt64, math.MinInt64
	for _, p := range coord {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	// autocrop
	for i := range coord {
		coord[i].X -= minX
		coord[i].Y -= minY
	}
	maxX -= minX
	maxY -= minY

	// init matrix
	matrix := make([][]Cell, maxX+1)
	for i := range matrix {
		matrix[i] = make([]Cell, maxY+1)
	}

	// Fill matrix with coordinates
	for _, c := range coord {
		matrix[c.X][c.Y] = c
	}

	// populate with influence of seniors cell
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if !matrix[i][j].Senior {
				var found bool
				for k := 1; !found; k++ {
					neighbors := generateIndexes(k)

					seniorClose := make([]Cell, 0)
					for _, off := range neighbors {
						if i+off.X >= 0 && i+off.X < maxX+1 && j+off.Y >= 0 && j+off.Y < maxY+1 {
							if matrix[i+off.X][j+off.Y].Senior {
								seniorClose = append(seniorClose, matrix[i+off.X][j+off.Y])
							}
						}
					}

					switch len(seniorClose) {
					case 0:
						continue
					case 1:
						found = true
						matrix[i][j].Label = strings.ToLower(seniorClose[0].Label)
					default:
						found = true
						matrix[i][j].Label = "."
					}
				}
			}
		}
	}

	// count area
	count := make(map[string]int)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			lb := matrix[i][j].Label
			if lb != "." {
				count[strings.ToUpper(lb)]++
			}
		}
	}

	// find max area
	maxNum := math.MinInt64
	maxCoord := ""
	for k, v := range count {
		if v > maxNum {
			maxNum = v
			maxCoord = k
		}
	}

	_ = maxCoord
	//fmt.Printf("Largest area is %s with %d\n", maxCoord, maxNum)
	fmt.Println(maxNum)
}

func generateIndexes(dst int) []Cell {
	points := make([]Cell, 0)
	for i := 0; i < dst+1; i++ {
		for _, j := range []int{+1, -1} {
			switch {
			case i == 0:
				points = append(points, Cell{X: 0, Y: j * dst})
			case i == dst:
				points = append(points, Cell{X: j * dst, Y: 0})
			default:
				points = append(points, Cell{X: j * i, Y: j * (dst - i)})
				points = append(points, Cell{X: -j * i, Y: j * (dst - i)})
			}
		}
	}
	return points
}
