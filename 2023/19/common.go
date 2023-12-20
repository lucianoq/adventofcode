package main

import (
	"bufio"
	"strconv"
	"strings"
)

type Rule struct {
	Else  bool
	Field uint8
	Op    byte
	Val   int
	To    string
}

var idx = map[byte]int{'x': 0, 'm': 1, 'a': 2, 's': 3}

type Workflows map[string][]Rule

func parseWorkflows(scanner *bufio.Scanner) Workflows {
	wfs := Workflows{}

	for scanner.Scan() {
		line := strings.TrimSuffix(scanner.Text(), "}")

		if line == "" {
			break
		}

		ff := strings.Split(line, "{")

		rulesStr := strings.Split(ff[1], ",")
		var rules []Rule

		for _, rule := range rulesStr {
			ifThen := strings.Split(rule, ":")

			if len(ifThen) == 1 {
				rules = append(rules, Rule{
					Else: true,
					To:   ifThen[0],
				})
				continue
			}

			cond := ifThen[0]
			val, _ := strconv.Atoi(cond[2:])
			rules = append(rules, Rule{
				Field: uint8(idx[cond[0]]),
				Op:    cond[1],
				Val:   val,
				To:    ifThen[1],
			})
		}

		wfs[ff[0]] = rules
	}
	return wfs
}
