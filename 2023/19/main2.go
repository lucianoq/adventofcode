package main

import (
	"bufio"
	"fmt"
	"os"
)

type Range [4]struct {
	Min, Max int
}

func (r Range) And(rule Rule) (Range, bool) {
	switch rule.Op {
	case '<':
		if rule.Val <= r[rule.Field].Min {
			return Range{}, false
		}

		if rule.Val > r[rule.Field].Max {
			return r, true
		}

		r[rule.Field].Max = rule.Val - 1
		return r, true

	case '>':
		if rule.Val < r[rule.Field].Min {
			return r, true
		}

		if rule.Val >= r[rule.Field].Max {
			return Range{}, false
		}

		r[rule.Field].Min = rule.Val + 1
		return r, true
	}
	return Range{}, false
}

func (r Range) AndNot(rule Rule) (Range, bool) {
	switch rule.Op {
	case '<':
		if rule.Val <= r[rule.Field].Min {
			return r, true
		}

		if rule.Val > r[rule.Field].Max {
			return Range{}, false
		}

		r[rule.Field].Min = rule.Val
		return r, true
	case '>':
		if rule.Val < r[rule.Field].Min {
			return Range{}, false
		}

		if rule.Val >= r[rule.Field].Max {
			return r, true
		}

		r[rule.Field].Max = rule.Val
		return r, true
	}
	return Range{}, false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	wfs := parseWorkflows(scanner)

	solutions := make(chan Range)
	go func() {
		rec(wfs, "in", Range{{1, 4000}, {1, 4000}, {1, 4000}, {1, 4000}}, solutions)
		close(solutions)
	}()

	sum := 0
	for r := range solutions {
		sum += (r[0].Max - r[0].Min + 1) * (r[1].Max - r[1].Min + 1) *
			(r[2].Max - r[2].Min + 1) * (r[3].Max - r[3].Min + 1)
	}
	fmt.Println(sum)
}

// recursive DFS
func rec(wfs Workflows, name string, rg Range, solutions chan<- Range) {
	if name == "R" {
		return
	}

	if name == "A" {
		solutions <- rg
		return
	}

	for _, r := range wfs[name] {

		// if it is "else", it's always true
		// and there is no alternative
		if r.Else {
			rec(wfs, r.To, rg, solutions)
			continue
		}

		newRg, ok := rg.And(r)
		if ok {
			// check branch when condition is true
			rec(wfs, r.To, newRg, solutions)
		}

		newRg, ok = rg.AndNot(r)
		if !ok {
			return
		}

		// if we can negate the rule,
		// use it and check the next
		rg = newRg
	}
}
