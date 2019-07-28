package main

import (
	"fmt"
	"math"
)

func main() {
	minCost := math.MaxInt64

	for i := 0; ; i++ {
		foundWinnerList := false

		ch := Generate(i)
		for list := range ch {
			cost := 0
			for _, m := range list {
				cost += m.Cost
			}

			if cost >= minCost {
				continue
			}

			if Game(list) {
				minCost = cost
				foundWinnerList = true
			}
		}

		if foundWinnerList {
			fmt.Println(minCost)
			return
		}
	}
}
