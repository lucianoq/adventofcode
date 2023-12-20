package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Part [4]int

func parseParts(scanner *bufio.Scanner) []Part {
	var parts []Part
	for scanner.Scan() {
		eqs := strings.Split(strings.Trim(scanner.Text(), "{}"), ",")
		p := Part{}
		for _, eq := range eqs {
			val, _ := strconv.Atoi(eq[2:])
			p[idx[eq[0]]] = val
		}
		parts = append(parts, p)
	}
	return parts
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wfs := parseWorkflows(scanner)
	parts := parseParts(scanner)

	sum := 0
	for _, p := range parts {
		if p.Accepted(wfs) {
			sum += p[0] + p[1] + p[2] + p[3]
		}
	}
	fmt.Println(sum)
}

func (p Part) Accepted(wfs Workflows) bool {
	curr := "in"
	for {
		for _, r := range wfs[curr] {
			if p.Satisfy(r) {
				switch r.To {
				case "A":
					return true
				case "R":
					return false
				}
				curr = r.To
				break
			}
		}
	}
}

func (p Part) Satisfy(r Rule) bool {
	if r.Else {
		return true
	}
	switch r.Op {
	case '<':
		return p[r.Field] < r.Val
	case '>':
		return p[r.Field] > r.Val
	}
	return false
}
