package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	divider0 = parse("[[2]]")
	divider1 = parse("[[6]]")
)

func main() {
	list := [][]any{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		list = append(list, parse(line))
	}
	list = append(list, divider0, divider1)

	// order all packets
	sort.Slice(list, func(i, j int) bool {
		return compare(list[i], list[j]) == -1
	})

	// multiply indexes (1-based) of dividers
	mul := 1
	for i := 0; i < len(list); i++ {
		if compare(list[i], divider0) == 0 || compare(list[i], divider1) == 0 {
			mul *= i + 1
		}
	}
	fmt.Println(mul)
}
