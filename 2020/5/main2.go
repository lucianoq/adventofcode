package main

import (
	"fmt"
	"sort"
)

func main() {
	ids := parse()

	sort.Ints(ids)

	for i := 0; i < len(ids)-1; i++ {
		if ids[i+1]-ids[i] == 2 {
			fmt.Println(ids[i] + 1)
			return
		}
	}
}
