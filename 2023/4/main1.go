package main

import "fmt"

func main() {
	cards := parse()

	sum := 0
	for _, c := range cards {
		sum += c.Points()
	}
	fmt.Println(sum)
}

func (c Card) Points() int {
	pt := 0
	for _, h := range c.Have {
		if _, ok := c.Win[h]; ok {
			if pt == 0 {
				pt = 1
				continue
			}
			pt *= 2
		}
	}
	return pt
}
