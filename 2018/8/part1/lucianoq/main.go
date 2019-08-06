package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var sum int

func main() {
	input := readInput()
	readNode(input)
	fmt.Println(sum)
}

func readNode(in []int) []int {
	numChildren := in[0]
	numMetadata := in[1]

	leaving := in[2:]
	for i := 0; i < numChildren; i++ {
		leaving = readNode(leaving)
	}

	for i := 0; i < numMetadata; i++ {
		sum += leaving[i]
	}
	return leaving[numMetadata:]
}

func readInput() []int {
	input := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			nums := strings.Split(line, " ")
			for _, n := range nums {
				i, err := strconv.Atoi(n)
				if err != nil {
					log.Fatal(err)
				}
				input = append(input, i)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return input
}
