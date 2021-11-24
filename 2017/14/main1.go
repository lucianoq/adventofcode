package main

import "fmt"

func main() {
	g := NewGrid()

	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if g[i][j] {
				count++
			}
		}
	}
	fmt.Println(count)
}
