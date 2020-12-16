package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	rules := parseRules(scanner)
	_ = parseMyTicket(scanner)
	nearby := parseNearbyTickets(scanner)

	errorRate := 0
	for _, ticket := range nearby {
	Field:
		for _, field := range ticket {
			for _, rule := range rules {
				if rule.Valid(field) {
					continue Field
				}
			}
			errorRate += field
		}
	}
	fmt.Println(errorRate)
}
