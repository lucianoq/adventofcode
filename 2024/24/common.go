package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rule struct {
	Left, Op, Right string
}

func parseInput() (map[string]bool, map[string]Rule) {
	var facts = map[string]bool{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		ff := strings.Split(line, ": ")
		facts[ff[0]] = ff[1] == "1"
	}

	var rules = map[string]Rule{}
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		rules[ff[4]] = Rule{ff[0], ff[1], ff[2]}
	}

	return facts, rules
}

func GetNumber(letter string, facts map[string]bool, rules map[string]Rule) int {
	var z int
	for i := 0; i < 64; i++ {
		if solve(fmt.Sprintf("%s%02d", letter, i), facts, rules) {
			z |= 1 << i
		}
	}
	return z
}

func solve(target string, facts map[string]bool, rules map[string]Rule) bool {
	if v, ok := facts[target]; ok {
		return v
	}

	rule, ok := rules[target]
	if !ok {
		return false
	}

	left := solve(rule.Left, facts, rules)
	right := solve(rule.Right, facts, rules)

	var result bool
	switch rule.Op {
	case "AND":
		result = left && right
	case "OR":
		result = left || right
	case "XOR":
		result = left != right
	}
	facts[target] = result
	return result
}
