package main

import "fmt"

func main() {
	runInput()
	fmt.Println(count())
}

func count() int {
	count := 0
	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if display[i][j] {
				count++
			}
		}
	}
	return count
}
