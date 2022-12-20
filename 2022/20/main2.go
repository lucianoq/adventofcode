package main

import "fmt"

const DecryptionKey = 811589153

func main() {
	index, zero := parse()

	// Multiply everything by DecryptionKey
	for _, item := range index {
		item.Val *= DecryptionKey
	}

	// Mix 10 times
	for j := 0; j < 10; j++ {
		mix(index)
	}

	fmt.Println(groveCoordinates(zero))
}
