package main

import "fmt"

func main() {
	instructions, left, right := parse()

	curr := "AAA"
	for i := 1; ; i++ {
		switch instructions[(i-1)%len(instructions)] {
		case 'L':
			curr = left[curr]
		case 'R':
			curr = right[curr]
		}

		if curr == "ZZZ" {
			fmt.Println(i)
			return
		}
	}
}
