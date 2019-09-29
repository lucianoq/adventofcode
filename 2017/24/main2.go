package main

import "fmt"

func main() {
	_, maxStrength := generate(parse(), []int{0})
	fmt.Println(maxStrength)
}

func generate(available map[Comp]struct{}, bridge []int) (int, int) {
	maxLength, maxStrength := -1, -1

	for c := range available {
		if ok, ord := compatible(c, bridge); ok {
			l, s := generate(removeComp(available, c), addComp(bridge, ord))
			maxLength, maxStrength = updateMax(maxLength, maxStrength, l, s)
		}
	}

	if maxLength == -1 {
		return len(bridge), strength(bridge)
	}

	return maxLength, maxStrength
}

func updateMax(maxL, maxS, l, s int) (int, int) {
	switch {
	case l < maxL:
		return maxL, maxS
	case l > maxL:
		return l, s
	default:
		if s > maxS {
			return l, s
		} else {
			return l, maxS
		}
	}
}
