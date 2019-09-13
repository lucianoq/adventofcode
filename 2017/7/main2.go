package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	Name           string
	Weight         int
	Children       []*Program
	SumWeight      int
	stringChildren []string
}

func (n *Program) LinkRecursively(programs map[string]*Program) {
	for _, c := range programs[n.Name].stringChildren {
		programs[c].LinkRecursively(programs)
		n.Children = append(n.Children, programs[c])
	}
}

func (n *Program) FillSumWeightRecursively() {
	w := n.Weight
	for _, c := range n.Children {
		c.FillSumWeightRecursively()
		w += c.SumWeight
	}
	n.SumWeight = w
}

func (n *Program) FindWrongNodeAndFix(mustWeigh int) (wrong *Program, correctWeight int) {
	// Impossible case.
	// No chance to find out which one is wrong (maybe both!?)
	if len(n.Children) == 2 {
		return nil, 0
	}

	if allSameWeight(n.Children) {
		// SubTowers are fine. Current is the wrong
		supposedWeight := mustWeigh - len(n.Children)*n.Children[0].SumWeight
		return n, supposedWeight
	} else {
		// At least one is wrong. FindWrongNodeAndFix it
		wrong, correctWeight = findWrong(n.Children)
		return wrong.FindWrongNodeAndFix(correctWeight)
	}
}

// return whether all the nodes in a list have the same total weight
func allSameWeight(ns []*Program) bool {
	if len(ns) < 2 {
		return true
	}

	for i := 0; i < len(ns)-1; i++ {
		if ns[i].SumWeight != ns[i+1].SumWeight {
			return false
		}
	}
	return true
}

// find the wrong node and the right weight it should be
func findWrong(ns []*Program) (*Program, int) {
	size := len(ns)
	for i := 0; i < size; i++ {
		if ns[i].SumWeight != ns[(i+1)%size].SumWeight && ns[i].SumWeight != ns[(i+2)%size].SumWeight {
			return ns[i], ns[(i+1)%size].SumWeight
		}
	}
	return nil, 0
}

func main() {
	// found in part1
	const start = "eqgvf"

	programs := parse()
	var tree = &Program{
		Name:   programs[start].Name,
		Weight: programs[start].Weight,
	}
	tree.LinkRecursively(programs)
	tree.FillSumWeightRecursively()

	_, correctWeight := tree.FindWrongNodeAndFix(-1)

	fmt.Println(correctWeight)
}

func parse() map[string]*Program {
	result := make(map[string]*Program)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Fields(line)
		left := ff[0]
		weight, _ := strconv.Atoi(strings.Trim(ff[1], "()"))
		right := make([]string, 0)
		if len(ff) > 2 && ff[2] == "->" {
			for _, s := range ff[3:] {
				right = append(right, strings.TrimRight(s, ","))
			}
		}

		result[left] = &Program{
			Name:           left,
			Weight:         weight,
			stringChildren: right,
		}
	}
	return result
}
