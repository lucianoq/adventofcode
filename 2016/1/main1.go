package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	items := parse()
	path := make(map[byte]int)

	face := 0
	for _, i := range items {
		switch i[0] {
		case 'R':
			face = (face + 1 + 4) % 4
		case 'L':
			face = (face - 1 + 4) % 4
		}
		steps, err := strconv.Atoi(i[1:])
		if err != nil {
			log.Fatal(err)
		}
		path[dirs[face]] += steps
	}

	fmt.Println(abs(path['N']-path['S']) + abs(path['E']-path['W']))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
