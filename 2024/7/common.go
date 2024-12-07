package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		ff := strings.Split(scanner.Text(), ":")
		left, _ := strconv.Atoi(ff[0])
		right := []int{}
		for _, v := range strings.Fields(ff[1]) {
			n, _ := strconv.Atoi(v)
			right = append(right, n)
		}

		if valid(left, right[0], right[1:]) {
			sum += left
		}
	}
	fmt.Println(sum)
}
