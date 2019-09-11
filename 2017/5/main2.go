package main

import (
	"fmt"
)

func main() {
	list := parse()

	pos := 0
	count := 0
	for pos >= 0 && pos < len(list) {
		jump := list[pos]
		oldPos := pos
		pos += jump

		if jump >= 3 {
			list[oldPos]--
		} else {
			list[oldPos]++
		}

		count++
	}
	fmt.Println(count)
}
