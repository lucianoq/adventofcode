package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rules := parseRules(scanner)
	myTicket := parseMyTicket(scanner)
	nearby := parseNearbyTickets(scanner)

	// filter out invalid
	validNearby := filterOutInvalid(nearby, rules)

	// empty slice to be filled, fixed length
	headers := make([]string, len(myTicket))

	// until we didn't assign all the field names
	// to the right place in headers array
	for len(rules) > 0 {

		// for all columns
		for f := 0; f < len(myTicket); f++ {

			// optimization: if we already found the header
			// for this field, there is no need to try others.
			if headers[f] != "" {
				continue
			}

			// try all remaining (to be assigned) rules
			// counting how many are valid and keeping reference of
			// the last valid in case that's the only one
			validRules := 0
			lastValidRule := ""
		Rule:
			for ruleName, rule := range rules {

				// check if every ticket for that column is
				// valid for that rule
				for t := 0; t < len(validNearby); t++ {

					// if at least one is not valid,
					// this rule is not good and
					// we need to check another rule
					if !rule.Valid(validNearby[t][f]) {
						continue Rule
					}
				}

				// rule is valid, count and remember
				validRules++
				lastValidRule = ruleName
			}

			// we take a decision only when it is "forced" that is
			// when only one rule fits for that field in all tickets.
			if validRules == 1 {

				// rule is assigned, put it in the right place
				// and delete it from rule map
				headers[f] = lastValidRule
				delete(rules, lastValidRule)
			}
		}
	}

	// multiply departure fields
	mul := 1
	for i := 0; i < len(headers); i++ {
		if strings.HasPrefix(headers[i], "departure") {
			mul *= myTicket[i]
		}
	}
	fmt.Println(mul)
}

func filterOutInvalid(nearby [][]int, rules map[string]Rule) [][]int {
	var filtered [][]int

Ticket:
	for _, ticket := range nearby {
	Field:
		for _, field := range ticket {
			for _, rule := range rules {
				if rule.Valid(field) {
					continue Field
				}

			}
			continue Ticket
		}
		filtered = append(filtered, ticket)
	}
	return filtered
}
