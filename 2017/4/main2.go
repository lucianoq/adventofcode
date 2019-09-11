package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

		f = sortString(f)

		if _, ok := set[f]; ok {
			return false
		}
		set[f] = struct{}{}
	}
	return true
}

func sortString(s string) string {
	slice := []byte(s)

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})

	return string(slice)
}
