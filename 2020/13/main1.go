package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")

	delays := map[int]int{}

	for _, b := range buses {
		if b == "x" {
			continue
		}

		id, _ := strconv.Atoi(b)
		delays[id] = id - t%id
	}

	minID, minDelay := 0, 1<<63-1

	for id, delay := range delays {
		if delay < minDelay {
			minDelay = delay
			minID = id
		}
	}

	fmt.Println(minID * minDelay)
}
