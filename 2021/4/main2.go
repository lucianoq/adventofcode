package main

import "fmt"

func main() {
	draws, boards := parse()

	drawn := map[int]struct{}{}

	for _, d := range draws {

		drawn[d] = struct{}{}

		if len(boards) == 1 && boards[0].Wins(drawn) {
			fmt.Println(d * boards[0].SumUnmarked(drawn))
			return
		}

		// remove won boards
		filtered := make([]*Board, 0, len(boards)/2)
		for _, b := range boards {
			if !b.Wins(drawn) {
				filtered = append(filtered, b)
			}
		}
		boards = filtered
	}
}
