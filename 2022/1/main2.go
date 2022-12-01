package main

import (
	"fmt"
	"sort"
)

func main() {
	elves := parse()

	list := []int{}
	for _, v := range elves {
		list = append(list, v)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(list)))
	sum := list[0] + list[1] + list[2]
	fmt.Println(sum)
}
