package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	outputs := parse()

	count := 0
	for _, output := range outputs {
		for _, s := range output {
			switch len(s) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	fmt.Println(count)
}

func parse() [][]string {
	scanner := bufio.NewScanner(os.Stdin)
	outputs := make([][]string, 0)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		output := strings.Split(line[1], " ")
		outputs = append(outputs, output)
	}
	return outputs
}
