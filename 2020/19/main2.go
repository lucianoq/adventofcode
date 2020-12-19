package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	rules, messages := parse()

	// <T> ::= a | a<T> ==> one or more `a` ==> a+
	rules["8"] = "<42>+"

	// <T> ::= ab | a<T>b ==> one or more `a` followed by the same amount of `b`
	// This is not a regular grammar, but if we cap the number of occurrences
	// we can write it as regexp.
	// Following is `ab|aabb|aaabbb|aaaabbbb|aaaaabbbbb`
	rules["11"] = "(?:<42><31>|<42>{2}<31>{2}|<42>{3}<31>{3}|<42>{4}<31>{4}|<42>{5}<31>{5})"

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
