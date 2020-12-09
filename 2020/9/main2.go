package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ls := make([]int, 0)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		ls = append(ls, n)
	}

	const target = 26796446

	l, r := findRange(target, ls)

	min, max := target, 0
	for i := l; i < r+1; i++ {
		if ls[i] < min {
			min = ls[i]
		}
		if ls[i] > max {
			max = ls[i]
		}
	}

	fmt.Println(min + max)
}

func findRange(target int, ls []int) (int, int) {
L:
	for l := 0; l < len(ls)-2; l++ {
		sum := ls[l]
		for r := l + 1; r < len(ls); r++ {
			sum += ls[r]

			if sum == target {
				return l, r
			}

			if sum > target {
				continue L
			}
		}
	}
	panic("not found")
}
