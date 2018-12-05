package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	found := true
	for found {
		line, found = React(line)
	}
	fmt.Println(len(line))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func React(line string) (string, bool) {
	for i := 0; i < len(line)-1; i++ {
		if Abs(int(line[i])-int(line[i+1])) == 32 {
			return line[:i] + line[i+2:], true
		}
	}
	return line, false
}
