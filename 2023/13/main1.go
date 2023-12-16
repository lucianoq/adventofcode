package main

import "fmt"

func main() {
	patterns := parse()

	sum := 0
	for _, p := range patterns {
		if v, found := p.Transpose().FindHorizontalSymmetry(); found {
			sum += v + 1
		}
		if h, found := p.FindHorizontalSymmetry(); found {
			sum += 100 * (h + 1)
		}
	}
	fmt.Println(sum)
}

func (p Pattern) FindHorizontalSymmetry() (int, bool) {
NextLine:
	for line := 0; line < len(p)-1; line++ {
		for delta := 0; ; delta++ {
			up, down := line-delta, line+delta+1
			if up < 0 || down >= len(p) {
				return line, true
			}

			if p[up] != p[down] {
				continue NextLine
			}
		}
	}

	return 0, false
}
