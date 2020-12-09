package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const PreambleSize = 25

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ls := make([]int, 0)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		ls = append(ls, n)
	}

	for i := PreambleSize; i < len(ls); i++ {
		if !valid(ls[i], ls[i-PreambleSize:i]) {
			fmt.Println(ls[i])
			return
		}
	}
}

func valid(x int, ls []int) bool {
	for i := 0; i < len(ls)-1; i++ {
		for j := 1; j < len(ls); j++ {
			if ls[i]+ls[j] == x {
				return true
			}
		}
	}
	return false
}
