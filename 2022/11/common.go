package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	ID        int
	Items     []int
	Operation func(int) int
	Div       int
	ToTrue    int
	ToFalse   int
}

func parse() []Monkey {
	monkeys := []Monkey{}
	monkey := Monkey{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Fields(line)

		switch {
		case strings.HasPrefix(line, "Monkey"):
			id, _ := strconv.Atoi(strings.TrimRight(strings.Fields(line)[1], ":"))
			monkey.ID = id

		case strings.HasPrefix(line, "  Starting"):
			items := []int{}
			for i := 2; i < len(ff); i++ {
				num, _ := strconv.Atoi(strings.TrimRight(ff[i], ","))
				items = append(items, num)
			}
			monkey.Items = items

		case strings.HasPrefix(line, "  Operation") && ff[4] == "+":
			if ff[5] == "old" {
				monkey.Operation = func(old int) int {
					return old + old
				}
				continue
			}
			num, _ := strconv.Atoi(ff[5])
			monkey.Operation = func(old int) int {
				return old + num
			}

		case strings.HasPrefix(line, "  Operation") && ff[4] == "*":
			if ff[5] == "old" {
				monkey.Operation = func(old int) int {
					return old * old
				}
				continue
			}
			num, _ := strconv.Atoi(ff[5])
			monkey.Operation = func(old int) int {
				return old * num
			}

		case strings.HasPrefix(line, "  Test"):
			num, _ := strconv.Atoi(ff[3])
			monkey.Div = num

		case strings.HasPrefix(line, "    If true"):
			num, _ := strconv.Atoi(ff[5])
			monkey.ToTrue = num

		case strings.HasPrefix(line, "    If false"):
			num, _ := strconv.Atoi(ff[5])
			monkey.ToFalse = num

		case line == "":
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
		}
	}
	monkeys = append(monkeys, monkey)

	return monkeys
}

func monkeyBusiness(monkeys []Monkey, rounds int, manageStress func(worryLevel int) int) int {
	inspects := make([]int, len(monkeys))
	for r := 1; r <= rounds; r++ {
		for _, m := range monkeys {
			for _, worryLevel := range m.Items {
				inspects[m.ID]++
				worryLevel = m.Operation(worryLevel)
				worryLevel = manageStress(worryLevel)
				if worryLevel%m.Div == 0 {
					monkeys[m.ToTrue].Items = append(monkeys[m.ToTrue].Items, worryLevel)
				} else {
					monkeys[m.ToFalse].Items = append(monkeys[m.ToFalse].Items, worryLevel)
				}
			}
			monkeys[m.ID].Items = []int{}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(inspects)))
	return inspects[0] * inspects[1]
}
