package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rope := [10]P{}
	visited := map[P]struct{}{P{0, 0}: {}}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(ff[1])
		for i := 0; i < num; i++ {

			switch ff[0] {
			case "L":
				rope[0].X--
			case "R":
				rope[0].X++
			case "U":
				rope[0].Y++
			case "D":
				rope[0].Y--
			}

			for i := 1; i < 10; i++ {
				rope[i] = rope[i].Follow(rope[i-1])
			}

			visited[rope[9]] = struct{}{}
		}
	}
	fmt.Println(len(visited))
}
