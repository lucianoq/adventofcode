package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInput() map[string][]string {
	m := map[string][]string{}
	for scanner := bufio.NewScanner(os.Stdin); scanner.Scan(); {
		ff := strings.Split(scanner.Text(), ": ")
		m[ff[0]] = strings.Split(ff[1], " ")
	}
	return m
}
