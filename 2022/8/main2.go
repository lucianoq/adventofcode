package main

import "fmt"

func main() {
	s := scenicScore()

	max := 0
	for r := 0; r < Size; r++ {
		for p := 0; p < Size; p++ {
			p := P{r, p}
			if s[p] > max {
				max = s[p]
			}
		}
	}
	fmt.Println(max)
}

func scenicScore() map[P]int {
	s := map[P]int{}

	for r := 0; r < Size; r++ {
		for c := 0; c < Size; c++ {
			p := P{r, c}
			s[p] = view(p, Right) * view(p, Left) * view(p, Down) * view(p, Up)
		}
	}
	return s
}

func view(start, dir P) int {
	score := 0
	for next, valid := start.Add(dir); valid; next, valid = next.Add(dir) {
		score++
		if Map[next] >= Map[start] {
			break
		}
	}
	return score
}
