package main

import "fmt"

func main() {
	draws, boards := parse()

	for _, d := range draws {
		for _, b := range boards {
			b.Mark(d)
		}

		if len(boards) == 1 && boards[0].Wins() {
			fmt.Println(d * boards[0].SumUnmarked())
			return
		}

		// remove won boards
		filtered := make([]*Board, 0, len(boards)/2)
		for _, b := range boards {
			if !b.Wins() {
				filtered = append(filtered, b)
			}
		}
		boards = filtered
	}
}
