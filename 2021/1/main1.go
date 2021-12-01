package main

import "fmt"

func main() {
	list := parse()

	count := 0
	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			count++
		}
	}

	fmt.Println(count)
}
