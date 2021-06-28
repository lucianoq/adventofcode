package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type C struct {
	x, y, z, t int
}

func main() {
	constellations := parse()

Step:
	for {
		rand.Shuffle(len(constellations), func(i, j int) {
			constellations[i], constellations[j] = constellations[j], constellations[i]
		})

		l := len(constellations)
		for i := 0; i < l-1; i++ {
			for j := i + 1; j < l; j++ {
				for c1 := range constellations[i] {
					for c2 := range constellations[j] {
						if manhattan(c1, c2) <= 3 {
							constellations = merge(constellations, i, j)
							continue Step
						}
					}
				}
			}
		}

		fmt.Println(l)
		return
	}
}

func parse() []map[C]struct{} {
	scanner := bufio.NewScanner(os.Stdin)

	constellations := []map[C]struct{}{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		ff := strings.Split(line, ",")

		c := C{}

		n, _ := strconv.Atoi(ff[0])
		c.x = n
		n, _ = strconv.Atoi(ff[1])
		c.y = n
		n, _ = strconv.Atoi(ff[2])
		c.z = n
		n, _ = strconv.Atoi(ff[3])
		c.t = n

		constellations = append(constellations, map[C]struct{}{c: {}})
	}
	return constellations
}

func merge(list []map[C]struct{}, i int, j int) []map[C]struct{} {
	for c := range list[j] {
		list[i][c] = struct{}{}
	}
	list[j] = list[len(list)-1]
	return list[:len(list)-1]
}

func manhattan(c1, c2 C) int {
	return abs(c1.x-c2.x) + abs(c1.y-c2.y) + abs(c1.z-c2.z) + abs(c1.t-c2.t)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
