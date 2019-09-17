package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Scanner int

func (s Scanner) IsZero(pico int) bool {
	return pico%(s.Range()*2-2) == 0
}

func (s Scanner) Range() int {
	return int(s)
}

func parse() (firewall map[int]Scanner, max int) {
	firewall = make(map[int]Scanner)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		layer, _ := strconv.Atoi(strings.TrimRight(ff[0], ":"))
		width, _ := strconv.Atoi(ff[1])
		firewall[layer] = Scanner(width)
		if layer > max {
			max = layer
		}
	}
	return firewall, max
}
