package main

import "fmt"

func main() {
	documents := parse()

	count := 0
	for _, d := range documents {
		if d.Present() && d.Valid() {
			count++
		}
	}

	fmt.Println(count)
}
