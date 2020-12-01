package main

import "fmt"

func main() {
	list := parse()
	size := len(list)

	for i := 0; i < size-1; i++ {

		for j := i + 1; j < size; j++ {

			if list[i]+list[j] == 2020 {
				fmt.Println(list[i] * list[j])
				return
			}
		}
	}
}
