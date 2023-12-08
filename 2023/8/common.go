package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parse() (string, map[string]string, map[string]string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan()

	left := map[string]string{}
	right := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.ReplaceAll(line, ",", "")
		line = strings.ReplaceAll(line, ")", "")

		var h, l, r string
		_, _ = fmt.Sscanf(line, "%s = (%s %s", &h, &l, &r)

		left[h] = l
		right[h] = r
	}

	return instructions, left, right
}
