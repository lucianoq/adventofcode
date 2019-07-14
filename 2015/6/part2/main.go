package part2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const size = 1000

var grid [size][size]int

func main() {
	r, _ := regexp.Compile("(turn on|turn off|toggle) (\\d+),(\\d+) through (\\d+),(\\d+)")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			ss := r.FindStringSubmatch(line)

			op := ss[1]
			x1, _ := strconv.Atoi(ss[2])
			y1, _ := strconv.Atoi(ss[3])
			x2, _ := strconv.Atoi(ss[4])
			y2, _ := strconv.Atoi(ss[5])

			switch op {
			case "toggle":
				apply(toggle, x1, y1, x2, y2)
			case "turn on":
				apply(turnOn, x1, y1, x2, y2)
			case "turn off":
				apply(turnOff, x1, y1, x2, y2)
			}
		}
	}

	fmt.Println(brightness())
}

func apply(fn func(int, int), x1 int, y1 int, x2 int, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			fn(i, j)
		}
	}
}

func toggle(x, y int) {
	grid[x][y] += 2
}

func turnOn(x, y int) {
	grid[x][y] += 1
}

func turnOff(x, y int) {
	grid[x][y] -= 1
	if grid[x][y] < 0 {
		grid[x][y] = 0
	}
}

func brightness() int {
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			count += grid[i][j]
		}
	}
	return count
}
