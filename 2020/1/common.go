package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse() []int {
	scanner := bufio.NewScanner(os.Stdin)

	list := make([]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			i, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			list = append(list, i)
		}
	}
	return list
}
