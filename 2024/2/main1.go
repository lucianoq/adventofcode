package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	count := 0

	for scanner.Scan() {
		list, _ := toInt(strings.Fields(scanner.Text()))

		if safe(list) {
			count++
		}
	}

	fmt.Println(count)
}
