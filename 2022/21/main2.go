package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	Name        string
	Val         int
	HasVal      bool
	Left, Right *Monkey
	Op          string
}

func main() {
	index := parse()

	// set "humn" as unknown
	index["humn"].HasVal = false

	// new operator: == returns 0 if equals
	//                          1 otherwise
	index["root"].Op = "=="

	fmt.Println(index["root"].Expect(0))
}

func (m *Monkey) Solve() (int, bool) {
	if m.HasVal {
		return m.Val, true
	}

	if m.Left != nil && m.Right != nil {
		left, lOk := m.Left.Solve()
		right, rOk := m.Right.Solve()

		if lOk && rOk {
			switch m.Op {
			case "+":
				return left + right, true
			case "-":
				return left - right, true
			case "*":
				return left * right, true
			case "/":
				return left / right, true
			case "==":
				if left == right {
					return 0, true
				} else {
					return 1, true
				}
			}
		}
	}
	return 0, false
}

func (m *Monkey) Expect(x int) int {
	if m.Name == "humn" {
		return x
	}

	left, lOk := m.Left.Solve()
	right, rOk := m.Right.Solve()

	if !lOk {
		switch m.Op {
		case "+":
			return m.Left.Expect(x - right)
		case "-":
			return m.Left.Expect(x + right)
		case "*":
			return m.Left.Expect(x / right)
		case "/":
			return m.Left.Expect(x * right)
		case "==":
			return m.Left.Expect(right)
		}
	}

	if !rOk {
		switch m.Op {
		case "+":
			return m.Right.Expect(x - left)
		case "-":
			return m.Right.Expect(left - x)
		case "*":
			return m.Right.Expect(x / left)
		case "/":
			return m.Right.Expect(left / x)
		case "==":
			return m.Right.Expect(left)
		}
	}

	panic("impossible")
}

func parse() map[string]*Monkey {
	index := map[string]*Monkey{}

	initMonkey := func(s string) {
		if _, ok := index[s]; !ok {
			index[s] = &Monkey{Name: s}
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ff := strings.Split(line, ": ")
		goal := ff[0]
		initMonkey(goal)

		num, err := strconv.Atoi(ff[1])
		if err == nil {
			index[goal].Val = num
			index[goal].HasVal = true
			continue
		}

		r := strings.Split(ff[1], " ")
		left, op, right := r[0], r[1], r[2]

		initMonkey(left)
		initMonkey(right)

		index[goal].Left = index[left]
		index[goal].Op = op
		index[goal].Right = index[right]
	}
	return index
}
