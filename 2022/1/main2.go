package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const Size = 3

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	best := [Size]int{}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			setMax(sum, best[:])
			sum = 0
			continue
		}
		cal, _ := strconv.Atoi(line)
		sum += cal
	}

	sum = 0
	for i := 0; i < Size; i++ {
		sum += best[i]
	}
	fmt.Println(sum)
}

func setMax(num int, best []int) {
	if num < best[0] {
		return
	}

	best[0] = num
	for i := 0; i < Size-1; i++ {
		if best[i] > best[i+1] {
			best[i], best[i+1] = best[i+1], best[i]
		} else {
			return
		}
	}
}
