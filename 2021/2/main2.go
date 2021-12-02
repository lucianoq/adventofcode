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
	var horizontal, depth, aim int

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		ff := strings.Split(line, " ")

		val, err := strconv.Atoi(ff[1])
		if err != nil {
			log.Fatal(err)
		}

		switch ff[0] {
		case "forward":
			horizontal += val
			depth += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		}
	}

	fmt.Println(horizontal * depth)
}
