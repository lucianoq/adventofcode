package main

import "fmt"

func main() {
	patterns := parse()

	sum := 0
	for _, p := range patterns {
		if v, found := p.Transpose().FindHorizontalSymmetry(1); found {
			sum += v + 1
		}
		if h, found := p.FindHorizontalSymmetry(1); found {
			sum += 100 * (h + 1)
		}
	}
	fmt.Println(sum)
}

func (p Pattern) FindHorizontalSymmetry(smudges int) (int, bool) {
NextLine:
	for line := 0; line < len(p)-1; line++ {
		changes := smudges
		for delta := 0; ; delta++ {
			up, down := line-delta, line+delta+1
			if up < 0 || down >= len(p) {
				// only if all changes have been done
				if changes == 0 {
					return line, true
				}
				continue NextLine
			}

			diff := levenshtein(p[up], p[down])
			if diff > changes {
				continue NextLine
			}

			// apply changes
			changes -= diff
		}
	}

	return 0, false
}
