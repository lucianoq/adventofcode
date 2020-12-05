package main

import "fmt"

func main() {
	var max int

	ids := parse()

	for _, id := range ids {
		if id > max {
			max = id
		}
	}

	fmt.Println(max)
}
