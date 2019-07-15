package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

func parseInput() map[string]map[string]int {
	happiness := make(map[string]map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line != "" {
			var (
				from, to, sign string
				h              int
			)

			n, _ := fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s.\n", &from, &sign, &h, &to)
			if n != 4 {
				log.Fatal("not 4")
			}
			to = strings.TrimRight(to, ".")

			if happiness[from] == nil {
				happiness[from] = make(map[string]int)
			}
			switch sign {
			case "gain":
				happiness[from][to] = h
			case "lose":
				happiness[from][to] = -h
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return happiness
}

func main() {
	happiness := parseInput()

	var keys []string
	for k := range happiness {
		keys = append(keys, k)
	}

	// Add myself to the table
	happiness["luciano"] = make(map[string]int)
	for _, k := range keys {
		happiness[k]["luciano"] = 0
		happiness["luciano"][k] = 0
	}
	keys = append(keys, "luciano")

	// run async generator
	ch := make(chan int)
	go perm(keys, func(rs []string) {
		total := 0
		for i := 0; i < len(rs); i++ {
			total += happiness[rs[i]][rs[(i+1)%len(rs)]]
			total += happiness[rs[(i+1)%len(rs)]][rs[i]]
		}
		ch <- total
	}, 0)

	// collect data and find max
	max := math.MinInt32
	for i := uint64(0); i < factorial(len(keys)); i++ {
		x := <-ch
		if x > max {
			max = x
		}
	}

	fmt.Println(max)
}

func perm(a []string, f func([]string), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

func factorial(n int) uint64 {
	if n < 0 {
		log.Fatal("negative number")
	}

	factVal := uint64(1)
	for i := 1; i <= n; i++ {
		factVal *= uint64(i)
	}
	return factVal
}
