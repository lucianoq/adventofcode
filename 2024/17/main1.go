package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a, b, c, code := parseInput()

	output := run(a, b, c, code)

	fmt.Println(strings.Join(toStr(output), ","))
}

func toStr(x []int) []string {
	s := make([]string, len(x))
	for i := 0; i < len(x); i++ {
		s[i] = strconv.Itoa(x[i])
	}
	return s
}
