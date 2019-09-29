package main

import "fmt"

func main() {
	fmt.Println(generate(parse(), []int{0}))
}

func generate(available map[Comp]struct{}, bridge []int) int {
	maxS := -1
	for c := range available {
		if ok, ord := compatible(c, bridge); ok {
			val := generate(removeComp(available, c), addComp(bridge, ord))
			if val > maxS {
				maxS = val
			}
		}
	}

	if maxS == -1 {
		return strength(bridge)
	}

	return maxS
}
