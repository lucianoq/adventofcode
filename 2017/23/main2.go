package main

import (
	"fmt"
	"math"
)

func main() {
	count := 0
	for n := 107900; n <= 124900; n += 17 {
		if !Prime(n) {
			count++
		}
	}
	fmt.Println(count)
}

func Prime(n int) bool {
	for i := 2; i <= sqrt(n); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Optimization
func sqrt(i int) int {
	return int(math.Floor(math.Sqrt(float64(i))))
}
