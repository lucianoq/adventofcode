package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const Size = 10007

func main() {

	deck := make([]int, Size)
	for i := 0; i < Size; i++ {
		deck[i] = i
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "deal into new stack" {
			deck = dealIntoNewStack(deck)
			continue
		}

		if strings.HasPrefix(line, "cut") {
			ff := strings.Fields(line)
			n, _ := strconv.Atoi(ff[1])

			deck = cutN(deck, n)
			continue
		}

		if strings.HasPrefix(line, "deal with increment") {

			ff := strings.Fields(line)
			n, _ := strconv.Atoi(ff[len(ff)-1])

			deck = dealWithIncrement(deck, n)
			continue
		}
	}

	fmt.Println(find2019(deck))
}

func dealIntoNewStack(deck []int) []int {
	for i := 0; i < Size/2; i++ {
		deck[i], deck[Size-i-1] = deck[Size-i-1], deck[i]
	}
	return deck
}

func cutN(deck []int, n int) []int {
	if n >= 0 {
		return append(deck[n:], deck[:n]...)
	} else {
		return append(deck[len(deck)+n:], deck[:len(deck)+n]...)
	}
}

func dealWithIncrement(deck []int, n int) []int {
	newDeck := make([]int, len(deck))

	for i := 0; i < Size; i++ {
		newDeck[(i*n)%Size] = deck[i]
	}

	return newDeck
}

func find2019(deck []int) int {
	for i := 0; i < Size; i++ {
		if deck[i] == 2019 {
			return i
		}
	}
	return -1
}
