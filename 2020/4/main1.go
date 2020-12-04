package main

import "fmt"

func main() {
	documents := parse()

	count := 0
	for _, d := range documents {
		if d.Present() {
			count++
		}
	}

	fmt.Println(count)
}
