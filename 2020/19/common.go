package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var (
	char        = regexp.MustCompile(`^(\d+): "(\w)"$`)
	alternative = regexp.MustCompile(`^(\d+): ([\d ]+) \| ([\d ]+)$`)
	concat      = regexp.MustCompile(`^(\d+): ([\d ]+)$`)
	term        = regexp.MustCompile(`<(\d+)>`)
)

func parse() (map[string]string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	rules := map[string]string{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		switch {
		case char.MatchString(line):
			sm := char.FindStringSubmatch(line)
			rules[sm[1]] = sm[2]

		case concat.MatchString(line):
			sm := concat.FindStringSubmatch(line)

			ls := strings.Fields(sm[2])
			for _, x := range ls {
				rules[sm[1]] += "<" + x + ">"
			}

		case alternative.MatchString(line):
			sm := alternative.FindStringSubmatch(line)

			alt1 := strings.Fields(sm[2])
			alt2 := strings.Fields(sm[3])

			s := "(?:" // non-capturing group
			for _, x := range alt1 {
				s += "<" + x + ">"
			}
			s += "|"
			for _, x := range alt2 {
				s += "<" + x + ">"
			}
			s += ")"

			rules[sm[1]] = s
		}
	}

	var messages []string
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}

	return rules, messages
}
