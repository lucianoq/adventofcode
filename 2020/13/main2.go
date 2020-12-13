package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	ff := strings.Split(scanner.Text(), ",")

	buses := map[int]int{}
	for idx, b := range ff {
		if b == "x" {
			continue
		}

		id, _ := strconv.Atoi(b)
		buses[idx] = id
	}

	timestamp, delta := buses[0], buses[0]

	for idx, bus := range buses {
		for (timestamp+idx)%bus != 0 {
			timestamp += delta
		}
		delta = lcm(delta, bus)
	}

	fmt.Println(timestamp)
}

func gcd(a, b int) int {
	for b != 0 {
		b, a = a%b, b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
