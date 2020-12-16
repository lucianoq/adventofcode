package main

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Name  string
	Valid func(x int) bool
}

func parseRules(scanner *bufio.Scanner) map[string]Rule {
	rg := regexp.MustCompile("^([\\w ]+): (\\d+)-(\\d+) or (\\d+)-(\\d+)$")

	rules := map[string]Rule{}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			return rules
		}

		ff := rg.FindStringSubmatch(line)

		a, _ := strconv.Atoi(ff[2])
		b, _ := strconv.Atoi(ff[3])
		c, _ := strconv.Atoi(ff[4])
		d, _ := strconv.Atoi(ff[5])

		rules[ff[1]] = Rule{
			Name: ff[1],
			Valid: func(x int) bool {
				return (x >= a && x <= b) || (x >= c && x <= d)
			},
		}
	}

	return rules
}

func parseMyTicket(scanner *bufio.Scanner) []int {
	scanner.Scan() // skip "your ticket:"

	scanner.Scan() // read my ticket
	ls := toInts(scanner.Text())

	scanner.Scan() // skip empty line

	return ls

}

func parseNearbyTickets(scanner *bufio.Scanner) [][]int {
	scanner.Scan() // skip "nearby tickets:"

	var retList [][]int
	for scanner.Scan() {
		line := scanner.Text()
		retList = append(retList, toInts(line))
	}

	return retList
}

func toInts(line string) []int {
	ff := strings.Split(line, ",")

	var ls []int
	for _, f := range ff {
		num, _ := strconv.Atoi(f)
		ls = append(ls, num)
	}
	return ls
}
