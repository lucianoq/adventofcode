package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	packages := make([]int, 0)
	weight := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		packages = append(packages, num)
		weight += num
	}

	weightGoal := weight / 3

	ch := make(chan []int)
	go func() {
		// Wrote 6 after found out no groups smaller than 6 is possible.
		// TODO parametrise it
		generate(6, weightGoal, packages, []int{}, 0, ch)
		close(ch)
	}()

	minEntanglement := math.MaxInt64
	for list := range ch {
		remaining := complementary(list, packages)
		if !checkValid(remaining, weightGoal) {
			continue
		}

		entanglement := entanglement(list)
		if entanglement < minEntanglement {
			minEntanglement = entanglement
		}
	}
	fmt.Println(minEntanglement)
}

func entanglement(list []int) int {
	prod := 1
	for _, x := range list {
		prod *= x
	}
	return prod
}

// if we are able to find a second group with the same goal, so also the third group will satisfy the goal.
func checkValid(list []int, weightGoal int) bool {
	ch := make(chan []int)
	go func() {
		for i := 0; i < len(list); i++ {
			generate(i, weightGoal, list, []int{}, 0, ch)
		}
		close(ch)
	}()

	x := <-ch
	if x != nil {
		return true
	}
	return false
}

func generate(size int, weightGoal int, list []int, accum []int, sumSoFar int, ch chan []int) {
	if size == 0 {
		if sumSoFar == weightGoal {
			ch <- accum
		}
		return
	}

	for i := range list {
		newSumSoFar := sumSoFar + list[i]
		if newSumSoFar > weightGoal {
			continue
		}

		newList := append([]int{}, list[i+1:]...)

		newRemaining := append([]int{}, accum...)
		newRemaining = append(newRemaining, list[i])

		generate(size-1, weightGoal, newList, newRemaining, newSumSoFar, ch)
	}
}

func complementary(list, total []int) []int {
	remaining := make([]int, 0, len(total)-len(list))
	t, l := 0, 0

	for t < len(total) && l < len(list) {
		switch {
		case total[t] == list[l]:
			t++
			l++
		case total[t] < list[l]:
			remaining = append(remaining, total[t])
			t++
		}
	}
	return append(remaining, total[t:]...)
}
