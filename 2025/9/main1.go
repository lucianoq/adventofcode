package main

import "fmt"

func main() {
	list := parseInput()

	maxArea := 0
	for i := 0; i < len(list)-1; i++ {
		for j := i; j < len(list); j++ {
			a := NewRectangle(list[i], list[j]).Area()
			if a > maxArea {
				maxArea = a
			}
		}
	}
	fmt.Println(maxArea)
}
