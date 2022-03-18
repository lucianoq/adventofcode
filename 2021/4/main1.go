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

		for _, b := range boards {
			if b.Wins(drawn) {
				fmt.Println(d * b.SumUnmarked(drawn))
				return
			}
		}
	}
}
