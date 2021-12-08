package main

import (
	"bufio"
	"fmt"
	"math/bits"
	"os"
	"strings"
)

// Using uint8 as bitset
// (MSB unused, LSB is `a`)
//
//  7 6 5 4 3 2 1 0
//  - g f e d c b a
//
// e.g.
//   pattern = `abcd`
//     00001111
//     ----dcba
//
//   pattern = `aceg`
//     01010101
//     -g-e-c-a

func toBitSet(s string) uint8 {
	var bs uint8
	for _, c := range s {
		bs |= 1 << (c - 'a')
	}
	return bs
}

func contains(a, b uint8) bool {
	return a&b == b
}

type Task struct {
	patterns []uint8
	dict     map[uint8]int
	revDict  map[int]uint8
	output   []uint8
}

func (t *Task) Found(pattern uint8, val int) {
	t.dict[pattern] = val
	t.revDict[val] = pattern
}

func (t *Task) Solve() {
	for _, p := range t.patterns {
		switch bits.OnesCount8(p) {
		case 2:
			t.Found(p, 1)
		case 3:
			t.Found(p, 7)
		case 4:
			t.Found(p, 4)
		case 7:
			t.Found(p, 8)
		}
	}

	topLeftL := t.revDict[4] &^ t.revDict[1]

	for _, p := range t.patterns {
		switch bits.OnesCount8(p) {

		case 5:
			if contains(p, t.revDict[7]) {
				t.Found(p, 3)
				continue
			}
			if contains(p, topLeftL) {
				t.Found(p, 5)
				continue
			}
			t.Found(p, 2)

		case 6:
			if !contains(p, t.revDict[7]) {
				t.Found(p, 6)
				continue
			}
			if contains(p, t.revDict[4]) {
				t.Found(p, 9)
				continue
			}
			t.Found(p, 0)
		}
	}
}

func (t *Task) Solution() int {
	sum := 0
	for i := 0; i < 4; i++ {
		sum = sum*10 + t.dict[t.output[i]]
	}
	return sum
}

func main() {
	tasks := parse()

	sum := 0
	for _, t := range tasks {
		t.Solve()
		sum += t.Solution()
	}
	fmt.Println(sum)
}

func parse() []Task {
	scanner := bufio.NewScanner(os.Stdin)

	tasks := []Task{}

	for scanner.Scan() {
		ff := strings.Split(scanner.Text(), " | ")

		var patterns []uint8
		var output []uint8
		for _, p := range strings.Split(ff[0], " ") {
			patterns = append(patterns, toBitSet(p))
		}
		for _, c := range strings.Split(ff[1], " ") {
			output = append(output, toBitSet(c))
		}

		tasks = append(tasks, Task{
			patterns: patterns,
			dict:     make(map[uint8]int),
			revDict:  make(map[int]uint8),
			output:   output,
		})
	}
	return tasks
}
