package main

import "fmt"

func (c Card) Matching() int {
	match := 0
	for _, h := range c.Have {
		if _, ok := c.Win[h]; ok {
			match++
		}
	}
	return match
}

func main() {
	cards := parse()

	// ID -> count
	pocket := map[int]int{}

	for i := 1; i <= Cards; i++ {
		pocket[i] = 1
	}

	for i := 1; i <= Cards; i++ {
		match := cards[i].Matching()
		for j := 1; j <= match; j++ {
			pocket[i+j] += pocket[i]
		}
	}

	sum := 0
	for _, count := range pocket {
		sum += count
	}
	fmt.Println(sum)
}
