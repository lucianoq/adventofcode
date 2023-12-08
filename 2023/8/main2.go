package main

import "fmt"

func main() {
	instructions, left, right := parse()

	curr := []string{}
	for k := range left {
		if k[2] == 'A' {
			curr = append(curr, k)
		}
	}

	periods := make([]int, len(curr))

	for i := 1; ; i++ {
		for j, k := range curr {

			switch instructions[(i-1)%len(instructions)] {
			case 'L':
				curr[j] = left[k]
			case 'R':
				curr[j] = right[k]
			}

			if curr[j][2] == 'Z' && periods[j] == 0 {
				periods[j] = i
			}
		}

		if complete(periods) {
			fmt.Println(lcmm(periods))
			return
		}
	}
}

func complete(xs []int) bool {
	for _, x := range xs {
		if x == 0 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcmm(xs []int) int {
	lcm := func(a, b int) int { return a * b / gcd(a, b) }

	result := 1
	for _, n := range xs {
		result = lcm(result, n)
	}
	return result
}
