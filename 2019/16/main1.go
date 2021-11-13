package main

import "fmt"

var pat = [4]int{0, 1, 0, -1}

func pattern(i, j int) int {
	if j < i {
		return 0
	}
	if j >= i && j < 2*i+1 {
		return 1
	}
	return pat[(j+1)/(i+1)%4]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func run100phases(list []int) []int {
	size := len(list)
	for phase := 0; phase < 100; phase++ {
		newList := make([]int, size)
		for i := 0; i < size; i++ {
			var sum int
			for j := i; j < size; j++ {
				sum += pattern(i, j) * list[j]
			}
			newList[i] = abs(sum) % 10
		}
		list = newList
	}
	return list
}

func main() {
	list := parse()
	list = run100phases(list)

	// Only first 8 elements
	fmt.Println(toInt(list[:8]))
}
