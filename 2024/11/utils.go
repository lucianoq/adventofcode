package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func numDigits(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}

	if m == 1 {
		return n
	}

	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func readInput() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nums := []int{}
	for _, num := range strings.Fields(scanner.Text()) {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}
	return nums
}
