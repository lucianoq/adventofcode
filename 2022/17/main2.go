package main

import "fmt"

const (
	// After OffsetLen unique rounds,
	// the pattern repeats every CycleLen rounds

	OffsetLen = 97
	CycleLen  = 1715

	Target = 1000000000000
)

func main() {
	dir := getDir(parse())

	peak := 0
	offsetHeight := 0
	heightsGain := [CycleLen]int{}

	for round := 0; round < OffsetLen+CycleLen; round++ {
		peak = runRound(round, peak, dir)

		if round == OffsetLen-1 {
			offsetHeight = peak
		}
		if round > OffsetLen-1 {
			heightsGain[round-OffsetLen] = peak - offsetHeight
		}
	}

	// Target is 1-based
	numCycles := (Target - 1 - OffsetLen) / CycleLen
	deltaGain := heightsGain[(Target-1-OffsetLen)%CycleLen]
	gainPerCycle := heightsGain[CycleLen-1]

	fmt.Println(offsetHeight + numCycles*gainPerCycle + deltaGain)
}
