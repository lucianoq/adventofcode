package main

import "fmt"

func main() {
	start, rules := parse()

	pairFreq := map[string]int{}
	for i := 1; i < len(start); i++ {
		pair := start[i-1 : i+1]
		pairFreq[pair]++
	}

	for step := 0; step < 40; step++ {
		newPairFreq := map[string]int{}
		for pair, count := range pairFreq {
			leftPair := string(pair[0]) + rules[pair]
			rightPair := rules[pair] + string(pair[1])
			newPairFreq[leftPair] += count
			newPairFreq[rightPair] += count
		}
		pairFreq = newPairFreq
	}

	letterCount := countLetters(pairFreq, start[0])

	min, max := minMax(letterCount)
	fmt.Println(max - min)
}

func countLetters(pairFreq map[string]int, first uint8) map[uint8]int {
	// count the first letter separately
	// (it will never change)
	letters := map[uint8]int{
		first: 1,
	}

	// then count only the second letter of the pair
	for pair, count := range pairFreq {
		letters[pair[1]] += count
	}
	return letters
}
