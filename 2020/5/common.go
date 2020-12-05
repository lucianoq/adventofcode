package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse() []int {
	scanner := bufio.NewScanner(os.Stdin)
	ids := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		row := line[:7]
		row = strings.ReplaceAll(row, "F", "0")
		row = strings.ReplaceAll(row, "B", "1")
		r, _ := strconv.ParseUint(row, 2, 8)

		col := line[7:]
		col = strings.ReplaceAll(col, "L", "0")
		col = strings.ReplaceAll(col, "R", "1")
		c, _ := strconv.ParseUint(col, 2, 8)

		id := 8*r + c

		ids = append(ids, int(id))
	}

	return ids
}
