package main

import "fmt"

func main() {
	draws, boards := parse()

	for _, d := range draws {
		for _, b := range boards {
			b.Mark(d)

			if b.Wins() {
				fmt.Println(d * b.SumUnmarked())
				return
			}
		}
	}
}
