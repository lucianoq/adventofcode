package main

import (
	"fmt"
	"math"
)

const input = 34000000

func main() {
	for h := 1; ; h++ {
		presents := present(h)
		if presents >= input {
			fmt.Println(h)
			return
		}
	}
}

func present(h int) int {
	presents := 0
	for i := 1; i <= sqrt(h); i++ {
		if h%i == 0 {
			presents += 10 * i
			if h/i != i {
				presents += 10 * h / i
			}
		}
	}
	return presents
}

func sqrt(x int) int {
	return int(math.Floor(math.Sqrt(float64(x))))
}
