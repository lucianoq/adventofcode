package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	initial, rules := readInput()
	fmt.Printf("%d:\t", 0)
	Dump(initial)

	current := initial
	for i := 1; i <= 20; i++ {
		current = ApplyRules(current, rules)
		fmt.Printf("%d:\t", i)
		Dump(current)
	}

	fmt.Println(Sum(current))
}

type Rule struct {
	In  [5]bool
	Out bool
}

func readInput() (map[int]bool, []Rule) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	initialState := make(map[int]bool)
	init := line[15:]
	for i, pot := range strings.Split(init, "") {
		initialState[i] = pot == "#"
	}
	rules := make([]Rule, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		in := [5]bool{
			string(line[0]) == "#",
			string(line[1]) == "#",
			string(line[2]) == "#",
			string(line[3]) == "#",
			string(line[4]) == "#",
		}

		r := Rule{
			In:  in,
			Out: string(line[9]) == "#",
		}
		rules = append(rules, r)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return initialState, rules
}

func ApplyRules(state map[int]bool, rules []Rule) map[int]bool {
	min, max := MinMax(state)

	newState := make(map[int]bool)
POTS:
	for i := min - 2; i <= max+2; i++ {
		for _, r := range rules {
			if state[i-2] == r.In[0] &&
				state[i-1] == r.In[1] &&
				state[i] == r.In[2] &&
				state[i+1] == r.In[3] &&
				state[i+2] == r.In[4] {
				newState[i] = r.Out
				continue POTS
			}
		}
		newState[i] = false
	}

	min, max = MinMax(newState)
	for i := min; i < 0; i++ {
		if !newState[i] {
			delete(newState, i)
		}
	}
	for i := max; !newState[i-1]; i-- {
		if !newState[i] {
			delete(newState, i)
		}
	}
	return newState
}

func Dump(state map[int]bool) {
	min, max := MinMax(state)
	for i := min; i <= max; i++ {
		if i == 0 {
			fmt.Print("\t")
		}
		if state[i] {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

	}
	fmt.Println()
}

func MinMax(state map[int]bool) (int, int) {
	min := 1<<63 - 1
	max := -1 << 63
	for k := range state {
		if k > max {
			max = k
		}
		if k < min {
			min = k
		}
	}
	return min, max
}

func Sum(state map[int]bool) int {
	min, max := MinMax(state)
	sum := 0
	for i := min; i <= max; i++ {
		if state[i] {
			sum += i
		}
	}
	return sum
}
