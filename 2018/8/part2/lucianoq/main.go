package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput()
	val, _ := readNode(input)
	fmt.Println(val)
}

func readNode(in []int) (int, []int) {
	numChildren, in := in[0], in[1:]
	numMetadata, in := in[0], in[1:]

	valueChildren := make([]int, numChildren)

	var val int
	for i := 0; i < numChildren; i++ {
		val, in = readNode(in)
		valueChildren[i] = val
	}

	if numChildren == 0 {
		var sum int
		for i := 0; i < numMetadata; i++ {
			var md int
			md, in = in[0], in[1:]
			sum += md
		}
		return sum, in
	}

	var sum int
	for i := 0; i < numMetadata; i++ {
		var md int
		md, in = in[0], in[1:]
		if md <= len(valueChildren) {
			sum += valueChildren[md-1]
		}
	}
	return sum, in
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
