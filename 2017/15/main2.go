package main

import "fmt"

func main() {
	a := generator(512, 16807, 4)
	b := generator(191, 48271, 8)

	count := 0
	for i := 0; i < 5000000; i++ {
		if a()&0xffff == b()&0xffff {
			count++
		}
	}
	fmt.Println(count)
}

func generator(start, factor, div2 uint) func() uint {
	prev := start
	return func() uint {
		for {
			prev = prev * factor % 2147483647
			if prev&(div2-1) == 0 {
				return prev
			}
		}
	}
}
