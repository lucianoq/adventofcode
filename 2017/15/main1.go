package main

import "fmt"

func main() {
	a := generator(512, 16807)
	b := generator(191, 48271)

	count := 0
	for i := 0; i < 40000000; i++ {
		if a()&0xffff == b()&0xffff {
			count++
		}
	}
	fmt.Println(count)
}

func generator(start, factor uint) func() uint {
	prev := start
	return func() uint {
		prev = prev * factor % 2147483647
		return prev
	}
}
