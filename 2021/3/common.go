package main

import (
	"bufio"
	"os"
	"strings"
)

func parse() []string {
	scanner := bufio.NewScanner(os.Stdin)
	list := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, strings.TrimSpace(line))
	}
	return list
}
