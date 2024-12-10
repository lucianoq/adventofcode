package main

import (
	"bufio"
	"fmt"
	"os"
)

const Empty = -1

func main() {
	memory := readMemory()
	defrag(memory)
	fmt.Println(checksum(memory))
}

func readMemory() []int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	memory := make([]int, 0, len(line)*5)

	for i, c := range line {
		num := int(c - '0')
		if i&1 == 0 {
			for j := 0; j < num; j++ {
				memory = append(memory, i/2)
			}
		} else {
			for j := 0; j < num; j++ {
				memory = append(memory, Empty)
			}
		}
	}

	return memory
}

func defrag(list []int) {
	left, right := 0, len(list)-1

	for {
		for list[left] != Empty {
			left++
		}

		for list[right] == Empty {
			right--
		}

		if left >= right {
			break
		}

		list[left], list[right] = list[right], list[left]
	}
}

func checksum(blocks []int) int {
	sum := 0
	for i, b := range blocks {
		if b != Empty {
			sum += i * b
		}
	}
	return sum
}
