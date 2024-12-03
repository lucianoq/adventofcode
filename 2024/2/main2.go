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

		if safe(list) || safeByRemoving(list) {
			count++
		}
	}

	fmt.Println(count)
}

func safeByRemoving(list []int) bool {
	for i := 0; i < len(list); i++ {
		newXs := append([]int{}, list[:i]...)
		newXs = append(newXs, list[i+1:]...)

		if safe(newXs) {
			return true
		}
	}
	return false
}
