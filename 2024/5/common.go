package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/lucianoq/container/set"
)

func parseInput() ([][]string, func(a, b string) int) {
	scanner := bufio.NewScanner(os.Stdin)

	lessThan := map[string]set.Set[string]{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ff := strings.Split(line, "|")
		if lessThan[ff[0]] == nil {
			lessThan[ff[0]] = set.New[string]()
		}
		lessThan[ff[0]].Add(ff[1])
	}

	updates := [][]string{}
	for scanner.Scan() {
		updates = append(updates, strings.Split(scanner.Text(), ","))
	}

	// slices.SortFunc
	cmpFn := func(a, b string) int {
		if lessThan[a].Contains(b) {
			return -1
		}
		if lessThan[b].Contains(b) {
			return 1
		}
		// 0 doesn't necessarily mean they're equal, but
		// as slices.SortFunc documentation says:
		// | The function should return 0 for incomparable items.
		return 0
	}

	return updates, cmpFn
}
