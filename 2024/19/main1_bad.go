package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	patterns, designs := parseInput()

	rgxp := regexp.MustCompile(fmt.Sprintf("^(%s)+$", strings.Join(patterns, "|")))

	sum := 0
	for _, d := range designs {
		if rgxp.MatchString(d) {
			sum++
		}
	}
	fmt.Println(sum)
}
