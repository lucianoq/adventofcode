package main

import "fmt"

func main() {
	sum := 0
	for _, num := range readInput() {
		sum += numChildren(num, Blinks)
	}
	fmt.Println(sum)
}

func blink(n int) []int {
	if n == 0 {
		return []int{1}
	}

	digits := numDigits(n)
	if digits&1 == 0 {
		pow := intPow(10, digits/2)
		return []int{n / pow, n % pow}
	}

	return []int{n * 2024}
}

// table for memoization
var cache = map[args]int{}

func numChildren(n, blinks int) int {

	if res, ok := cache[args{n, blinks}]; ok {
		return res
	}

	if blinks == 0 {
		return 1
	}

	sum := 0
	for _, child := range blink(n) {
		sum += numChildren(child, blinks-1)
	}

	cache[args{n, blinks}] = sum
	return sum
}

type args struct{ a, b int }
