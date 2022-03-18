package main

import (
	"fmt"

	"github.com/lucianoq/container/set"
)

func main() {
	draws, boards := parse()

	drawn := set.New[int]()

	for _, d := range draws {

		drawn.Add(d)

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
