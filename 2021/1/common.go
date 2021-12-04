package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func parse() []int {
	scanner := bufio.NewScanner(os.Stdin)

	list := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, num)
	}
	return list
}
