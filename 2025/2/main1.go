package main

import "fmt"

func main() {
	ranges := parseInput()
	
	sum := 0
	for _, r := range ranges {

		// only even digits
		if r.digits%2 != 0 {
			continue
		}

		for n := range r.InvalidIDs(r.digits / 2) {
			sum += n
		}
	}

	fmt.Println(sum)
}
