package main

import (
	"fmt"
	"math"
)

const input = 34000000

func main() {
	for h := 1; ; h++ {
		if presents(h) >= input {
			fmt.Println(h)
			return
		}
	}
}

func presents(h int) int {
	presents := 0
	for i := 1; i <= sqrt(h); i++ {
		if h%i == 0 {

			if h < 50*i {
				presents += 11 * i
			}

			mirrorDiv := h / i
			if mirrorDiv != i {
				if h < 50*mirrorDiv {
					presents += 11 * mirrorDiv
				}
			}
		}
	}
	return presents
}

func sqrt(x int) int {
	return int(math.Floor(math.Sqrt(float64(x))))
}
