package part2

import (
	"fmt"
	"log"
	"strconv"
)

func extract(s string) (string, []string) {
	var op, res, a1, a2 string

	var n int
	var err error

	n, err = fmt.Sscanf(s, "%s -> %s", &a1, &res)
	if err == nil && n == 2 {
		return ":=", []string{a1, res}
	}

	n, err = fmt.Sscanf(s, "NOT %s -> %s", &a1, &res)
	if err == nil && n == 2 {
		return "NOT", []string{a1, res}
	}

	n, err = fmt.Sscanf(s, "%s %s %s -> %s", &a1, &op, &a2, &res)
	if err == nil && n == 4 {
		return op, []string{a1, a2, res}
	}

	log.Fatalf("error with %s", s)

	return "", nil
}

func isNum(s string) bool {
	_, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		return false
	}
	return true
}

func available(s string) bool {
	_, ok := mem[s]
	return ok
}
