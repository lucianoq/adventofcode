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

	min := 1<<63 - 1 //MaxInt64
	for c := 'A'; c < '['; c++ {
		trimmedLine := strings.Replace(line, string(c), "", -1)
		trimmedLine = strings.Replace(trimmedLine, string(c+32), "", -1)
		reactedLen := FullyReact(trimmedLine)
		if reactedLen < min {
			min = reactedLen
		}
	}
	fmt.Println(min)
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

func FullyReact(line string) (int) {
	found := true
	for found {
		line, found = React(line)
	}
	return len(line)
}
