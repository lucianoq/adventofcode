package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var h, t P
	visited := map[P]struct{}{t: {}}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		num, _ := strconv.Atoi(ff[1])
		for i := 0; i < num; i++ {
			switch ff[0] {
			case "L":
				h.X--
			case "R":
				h.X++
			case "U":
				h.Y++
			case "D":
				h.Y--
			}
			t = t.Follow(h)
			visited[t] = struct{}{}
		}
	}
	fmt.Println(len(visited))
}
