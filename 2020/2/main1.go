package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var count int

	for scanner.Scan() {
		var (
			a, b           int
			char, password string
		)

		fmt.Sscanf(scanner.Text(), "%d-%d %1s: %s\n", &a, &b, &char, &password)

		if valid1(a, b, char, password) {
			count++
		}
	}

	fmt.Println(count)
}

func valid1(min, max int, char, password string) bool {
	c := strings.Count(password, char)
	return c >= min && c <= max
}
