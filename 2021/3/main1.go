package main

import "fmt"

func main() {
	list := parse()

	numBits := len(list[0])

	sum := make([]int, numBits)
	for _, item := range list {
		for bit := 0; bit < numBits; bit++ {
			if item[bit] == '1' {
				sum[bit]++
			}
		}
	}

	gamma, epsilon := 0, 0
	for bit := 0; bit < numBits; bit++ {
		if sum[bit] < len(list)/2 {
			epsilon += 1 << (numBits - bit - 1)
		} else {
			gamma += 1 << (numBits - bit - 1)
		}
	}

	fmt.Println(epsilon * gamma)
}
