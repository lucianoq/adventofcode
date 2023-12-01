package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += firstNumber(line)*10 + lastNumber(line)
	}

	fmt.Println(sum)
}

func firstNumber(s string) int {
	acc := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc += string(s[i])

		for i, d := range digits {
			if strings.HasSuffix(acc, d) {
				return i + 1
			}
		}
	}
	log.Fatal("not found")
	return 0
}

func lastNumber(s string) int {
	acc := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc = string(s[i]) + acc

		for i, d := range digits {
			if strings.HasPrefix(acc, d) {
				return i + 1
			}
		}
	}
	log.Fatal("not found")
	return 0
}
