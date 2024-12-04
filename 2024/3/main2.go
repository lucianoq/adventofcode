package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	text := ""
	for scanner.Scan() {
		text += scanner.Text()
	}

	rg := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)|do\\(\\)|don't\\(\\)")

	res := rg.FindAllStringSubmatch(text, -1)

	sum := 0
	enabled := true
	for _, l := range res {
		switch {
		case strings.HasPrefix(l[0], "mul"):
			if enabled {
				n1, _ := strconv.Atoi(l[1])
				n2, _ := strconv.Atoi(l[2])
				sum += n1 * n2
			}

		case l[0] == "do()":
			enabled = true

		case l[0] == "don't()":
			enabled = false
		}
	}

	fmt.Println(sum)
}
