package main

import "fmt"

func main() {
	m := NewMap()

	res := descent(m, 1, 1) *
		descent(m, 1, 3) *
		descent(m, 1, 5) *
		descent(m, 1, 7) *
		descent(m, 2, 1)

	fmt.Println(res)
}
