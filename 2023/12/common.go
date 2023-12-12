package main

import (
	"strconv"
	"strings"
)

func toList(s string) []int {
	list := strings.Split(s, ",")
	groups := make([]int, 0, len(list))
	for _, val := range list {
		n, _ := strconv.ParseUint(val, 10, 8)
		groups = append(groups, int(n))
	}
	return groups
}
