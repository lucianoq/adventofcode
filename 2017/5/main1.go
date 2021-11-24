package main

import "fmt"

func main() {
	list := parse()

	pos := 0
	count := 0
	for pos >= 0 && pos < len(list) {
		list[pos]++
		pos += list[pos] - 1

		count++
	}
	fmt.Println(count)
}
