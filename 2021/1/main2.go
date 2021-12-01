package main

import "fmt"

func main() {
	list := parse()

	count := 0
	for i := 3; i < len(list); i++ {
		if list[i] > list[i-3] {
			count++
		}
	}

	fmt.Println(count)
}
