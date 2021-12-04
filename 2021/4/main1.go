package main

import "fmt"

func main() {
	draws, boards := parse()

	drawn := map[int]struct{}{}

	for _, d := range draws {

		drawn[d] = struct{}{}

		for _, b := range boards {
			if b.Wins(drawn) {
				fmt.Println(d * b.SumUnmarked(drawn))
				return
			}
		}
	}
}
