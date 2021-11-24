package main

import "fmt"

func main() {
	fmt.Println(toInt(part2(parse())))
}

func multiply10k(list []int) []int {
	newList := make([]int, 0, len(list)*10000)
	for i := 0; i < 10000; i++ {
		newList = append(newList, list...)
	}
	return newList
}

func part2(list []int) []int {
	offset := toInt(list[:7])

	// repeat 10,000 times
	list = multiply10k(list)

	// digits in position n are not affected by
	// digits in position < n,
	// so we can skip completely the part of the list
	// left to the offset, reducing the search space
	// from a 6.5M * 6.5M to a ~500k*500k, by ~99.5%
	list = list[offset:]

	for phase := 0; phase < 100; phase++ {
		newList := make([]int, len(list))

		// To calculate the n-th digit, with n >= offset,
		// the pattern string contains n-1 `0` digits,
		// followed by n `1` digits, [followed by...],
		// but if we know that offset (and so n) > len(list)/2,
		// there will always be enough 1's to never reach
		// the 0's and the -1's.
		// We can remove completely the pattern and just sum
		// the digits (like they're multiplied by 1)
		//
		// Considering we are summing again and again the digits
		// of the same list, we can start from the last digit
		// and keep track of the partial sum, while storing the new digits.
		// In this way the O(n^2) becomes a O(n).
		var sum int
		for i := len(list) - 1; i >= 0; i-- {
			sum += list[i]
			newList[i] = sum % 10
		}

		list = newList
	}

	return list[:8]
}
