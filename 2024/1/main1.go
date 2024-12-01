package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var a, b []int
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())

		n1, _ := strconv.Atoi(ff[0])
		n2, _ := strconv.Atoi(ff[1])

		a = append(a, n1)
		b = append(b, n2)
	}

	sort.Ints(a)
	sort.Ints(b)

	diff := 0
	for i := 0; i < len(a); i++ {
		diff += abs(a[i] - b[i])
	}

	fmt.Println(diff)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
