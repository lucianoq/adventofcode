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
	scanner := bufio.NewScanner(os.Stdin)

	list := make([]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, n)
	}

	bigSum := 0

	for i := 1; i <= len(list); i++ {
		comb(len(list), i, func(c []int) {
			sum := 0
			for _, j := range c {
				sum += list[j]
			}
			if sum > 150 {
				return
			}
			if sum == 150 {
				bigSum++
			}
		})

		if bigSum >0 {
			break
		}
	}

	fmt.Println(bigSum)
}

func comb(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
