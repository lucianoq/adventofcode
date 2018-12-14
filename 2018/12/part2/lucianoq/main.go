package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const huge = 50000000000

var state map[int]bool
var rules []Rule
var left, right int

func main() {
	readInput()

	Dump(0)
	for i := 1; i <= 108; i++ {
		ApplyRules()
		Dump(i)
	}

	fmt.Println(Sum() + (huge-108)*len(state))
}

type Rule struct {
	In  [5]bool
	Out bool
}

func readInput() {
	state = make(map[int]bool)
	rules = make([]Rule, 0)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())

	init := line[15:]

	left, right = 0, len(init)

	for i, pot := range strings.Split(init, "") {
		if pot == "#" {
			state[i] = true
		}
	}
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

		rules = append(rules, Rule{
			In:  in,
			Out: string(line[9]) == "#",
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ApplyRules() {
	ll, l, this, r, rr := state[left-5], state[left-4], state[left-3], state[left-2], state[left-1]
POTS:
	for i := left - 2; i <= right+2; i++ {
		ll, l, this, r, rr = l, this, r, rr, state[i+2]
		for _, rule := range rules {
			if ll == rule.In[0] &&
				l == rule.In[1] &&
				this == rule.In[2] &&
				r == rule.In[3] &&
				rr == rule.In[4] {
				if rule.Out {
					if i < left {
						left = i
					} else if i > right {
						right = i
					}
					state[i] = rule.Out //true
				} else {
					delete(state, i)
				}
				continue POTS
			}
		}
	}
	for i := left - 2; i < right+2; i++ {
		if state[i] {
			left = i - 1
			break
		}
	}
}

func Dump(idx int) {
	fmt.Printf("%d: [%d]->", idx, left-2)
	for i := left - 2; i <= right+2; i++ {
		if i == 0 {
			fmt.Print(" ")
		}
		if state[i] {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func Sum() int {
	sum := 0
	for k, v := range state {
		if v {
			sum += k
		}
	}
	return sum
}
