package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	solve := map[string]func() int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Split(scanner.Text(), ": ")

		num, err := strconv.Atoi(ff[1])
		if err == nil {
			solve[ff[0]] = func() int { return num }
			continue
		}

		r := strings.Split(ff[1], " ")
		switch r[1] {
		case "+":
			solve[ff[0]] = func() int { return solve[r[0]]() + solve[r[2]]() }
		case "-":
			solve[ff[0]] = func() int { return solve[r[0]]() - solve[r[2]]() }
		case "*":
			solve[ff[0]] = func() int { return solve[r[0]]() * solve[r[2]]() }
		case "/":
			solve[ff[0]] = func() int { return solve[r[0]]() / solve[r[2]]() }
		}
	}

	fmt.Println(solve["root"]())
}
