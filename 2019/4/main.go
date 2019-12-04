package main

import "fmt"

const (
	min = 171309
	max = 643603
)

func main() {
	count := 0
	for i := min; i <= max; i++ {
		if valid(i) {
			count++
		}
	}
	fmt.Println(count)
}
