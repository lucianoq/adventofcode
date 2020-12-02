package main

import (
	"bufio"
	"fmt"
	"os"
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

		if valid2(a, b, char, password) {
			count++
		}
	}

	fmt.Println(count)
}

func valid2(i, j int, char, password string) bool {
	r := char[0]
	return (password[i-1] == r || password[j-1] == r) && password[i-1] != password[j-1]
}
