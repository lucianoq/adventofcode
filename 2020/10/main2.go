package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	ls := parse()

	ls = append(ls, 0) // adding jolts source
	sort.Ints(ls)
	ls = append(ls, ls[len(ls)-1]+3) // adding built-in adapter

	groups := split(ls)

	acc := 1
	for _, g := range groups {
		acc *= combinations(g[0], g[1:])
	}

	fmt.Println(acc)
}

// split creates a partition of the input list.
// It creates a separate subgroup whenever it finds a difference of 3
// between elements.
func split(ls []int) [][]int {
	res := make([][]int, 0)
	cut := 0

	for i := 0; i < len(ls)-1; i++ {
		if ls[i+1]-ls[i] == 3 {
			res = append(res, ls[cut:i+1])
			cut = i + 1
		}
	}
	res = append(res, ls[cut:])
	return res
}

func parse() []int {
	ls := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		ls = append(ls, n)
	}
	return ls
}

func combinations(curr int, ls []int) int {
	if len(ls) == 0 {
		return 1
	}

	count := 0
	for i := 0; i < min(3, len(ls)); i++ {

		if ls[i]-curr <= 3 {
			count += combinations(ls[i], ls[i+1:])
		}

	}
	return count
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}
