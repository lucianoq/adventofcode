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
	changes := make([]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		n, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		changes = append(changes, n)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var tot int
	dict := make(map[int]bool)
	dict[0] = true
	for i := 0; true; i++ {
		tot += changes[i%len(changes)]
		if dict[tot] {
			fmt.Println(tot)
			return
		}
		dict[tot] = true
	}
}
