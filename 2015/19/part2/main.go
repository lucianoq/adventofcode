package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Replacement struct {
	From string
	To   string
}

var (
	input        string
	replacements []Replacement
)

func main() {
	parse()

	rand.Seed(time.Now().UnixNano())

	min := math.MaxInt64

	// Repeating 100 times to find the minimum
	// Actually there is one possible path, so this is useless, but not knowing
	// that, this is better than taking the first
	for i := 0; i < 100; i++ {
		rand.Shuffle(len(replacements), func(i, j int) {
			replacements[i], replacements[j] = replacements[j], replacements[i]
		})

		newVal := guess()
		if newVal < min {
			min = newVal
		}
	}

	fmt.Println(min)
}

func guess() int {
	var s = input
	var steps = 0
	for {
		steps++

		newS, changed := tryToApplyAReplacement(s)
		if newS == "e" {
			return steps
		}

		if !changed {
			// Dead end here. We need to reset, shuffle and restart.
			s = input
			steps = 0
			rand.Shuffle(len(replacements), func(i, j int) {
				replacements[i], replacements[j] = replacements[j], replacements[i]
			})
			continue
		}

		s = newS
	}
}

func tryToApplyAReplacement(s string) (string, bool) {
	for _, r := range replacements {
		newS := strings.Replace(s, r.To, r.From, 1)

		if s != newS {
			return newS, true
		}
	}

	// No replacement is possible, dead end.
	return "", false
}

func parse() (string, []Replacement) {
	scanner := bufio.NewScanner(os.Stdin)

	replacements = make([]Replacement, 0)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			break
		}

		var from, to string
		n, _ := fmt.Sscanf(line, "%s => %s", &from, &to)
		if n != 2 {
			log.Fatal("parsing failed")
		}

		replacements = append(replacements, Replacement{From: from, To: to})
	}

	scanner.Scan()
	input = strings.TrimSpace(scanner.Text())

	return input, replacements
}
