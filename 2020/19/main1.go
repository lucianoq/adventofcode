package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	rules, messages := parse()

	line := "^" + rules["0"] + "$"
	for term.MatchString(line) {
		sm := term.FindStringSubmatch(line)
		line = strings.ReplaceAll(line, sm[0], rules[sm[1]])
	}

	regexp0 := regexp.MustCompile(line)

	count := 0
	for _, m := range messages {
		if regexp0.MatchString(m) {
			count++
		}
	}

	fmt.Println(count)
}
