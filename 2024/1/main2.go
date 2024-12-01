package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var a []int
	var b = map[int]int{}

	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())

		n1, _ := strconv.Atoi(ff[0])
		n2, _ := strconv.Atoi(ff[1])

		a = append(a, n1)
		b[n2]++
	}

	simScore := 0
	for _, n := range a {
		simScore += n * b[n]
	}

	fmt.Println(simScore)
}
