package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0
	idx := 1
	left, right := []any{}, []any{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			idx++
			left, right = left[:0], right[:0]
			continue
		}

		left = parse(line)
		scanner.Scan()
		line = scanner.Text()
		right = parse(line)

		if compare(left, right) == -1 {
			sum += idx
		}
	}

	fmt.Println(sum)
}
