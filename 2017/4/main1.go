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
		if valid(scanner.Text()) {
			count++
		}
	}
	fmt.Println(count)
}

func valid(line string) bool {
	set := make(map[string]struct{})

	ff := strings.Fields(line)
	for _, f := range ff {
		if _, ok := set[f]; ok {
			return false
		}
		set[f] = struct{}{}
	}
	return true
}
