package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
)

func parse() ([]int, int) {
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

	return packages, weight
}

func entanglement(list []int) int {
	prod := 1
	for _, x := range list {
		prod *= x
	}
	return prod
}

func minEntanglement(packages []int, weightGoal int, ch <-chan []int) int {
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
	return minEntanglement
}

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
